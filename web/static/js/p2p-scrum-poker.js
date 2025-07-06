document.addEventListener('DOMContentLoaded', () => {
    // UI Elements
    const roomSetup = document.getElementById('room-setup');
    const pokerRoom = document.getElementById('poker-room');
    const createRoomBtn = document.getElementById('createRoomBtn');
    const joinRoomBtn = document.getElementById('joinRoomBtn');
    const userNameInput = document.getElementById('userName');
    const roomIdInput = document.getElementById('roomIdInput');
    const roomIdDisplay = document.getElementById('roomIdDisplay');
    const adminControls = document.getElementById('admin-controls');
    const startVoteBtn = document.getElementById('startVoteBtn');
    const revealCardsBtn = document.getElementById('revealCardsBtn');
    const participantsList = document.getElementById('participants-list');
    const participantCountSpan = document.getElementById('participant-count');
    const votingCards = document.querySelectorAll('.card-vote');
    const resultsSummary = document.getElementById('results-summary');
    const averageVoteSpan = document.getElementById('average-vote');

    // State
    let peer;
    let myId = '';
    let myName = '';
    let myVote = null;
    let isHost = false;
    let connections = {};
    let participants = {}; // The single source of truth for the room state, only the host modifies it directly.
    let cardsRevealed = false;

    // --- URL Handling ---
    const urlParams = new URLSearchParams(window.location.search);
    const roomFromUrl = urlParams.get('room');
    if (roomFromUrl) {
        roomIdInput.value = roomFromUrl;
    }

    // --- Lifecycle Handling ---
    window.addEventListener('beforeunload', () => {
        // Proactively notify host that we are leaving. This is a "best effort" attempt.
        if (!isHost && peer && !peer.destroyed) {
            const hostConn = Object.values(connections)[0];
            if (hostConn && hostConn.open) {
                hostConn.send({ type: 'user-leave' });
            }
        }
    });

    // --- UI Functions ---
    function showPokerRoom() {
        roomSetup.classList.add('d-none');
        pokerRoom.classList.remove('d-none');
    }

    function updateParticipantsUI() {
        participantsList.innerHTML = '';
        const participantIds = Object.keys(participants);
        participantCountSpan.textContent = participantIds.length;

        let allVoted = participantIds.length > 0;

        participantIds.forEach(id => {
            const participant = participants[id];
            if (participant.vote === null) allVoted = false;

            let voteDisplay = '';

            if (cardsRevealed) {
                voteDisplay = participant.vote ?
                    `<span class="badge bg-info text-dark fs-6">${participant.vote}</span>` :
                    `<span class="badge bg-secondary">N/A</span>`;
            } else {
                voteDisplay = participant.vote ?
                    `<span class="badge bg-primary"><i class="bi bi-check-lg"></i> Voted</span>` :
                    `<span class="badge bg-secondary">Waiting...</span>`;
            }

            const isSelf = id === myId;
            const selfClass = isSelf ? 'active' : '';

            const li = document.createElement('li');
            li.className = `list-group-item list-group-item-dark d-flex justify-content-between align-items-center ${selfClass}`;
            li.innerHTML = `
                <span>${participant.name} ${isSelf ? '(You)' : ''}</span>
                ${voteDisplay}
            `;
            participantsList.appendChild(li);
        });

        if (isHost) {
            revealCardsBtn.disabled = !allVoted || cardsRevealed;
        }
    }

    function displayResults() {
        if (!cardsRevealed) {
            resultsSummary.classList.add('d-none');
            return;
        }

        const votes = Object.values(participants)
            .map(p => p.vote)
            .filter(v => v !== null && !isNaN(parseFloat(v)));

        if (votes.length > 0) {
            const numericVotes = votes.map(v => parseFloat(v));
            const sum = numericVotes.reduce((acc, val) => acc + val, 0);
            const average = (sum / numericVotes.length).toFixed(2);
            averageVoteSpan.textContent = average;
        } else {
            averageVoteSpan.textContent = 'N/A';
        }
        resultsSummary.classList.remove('d-none');
    }

    function resetLocalUI() {
        myVote = null;
        votingCards.forEach(c => {
            c.classList.remove('btn-light', 'text-dark');
            c.classList.add('btn-outline-light');
        });
        resultsSummary.classList.add('d-none');
        if (isHost) {
            revealCardsBtn.disabled = true;
        }
    }

    // --- P2P Communication ---
    function broadcast(data) {
        Object.values(connections).forEach(conn => conn.send(data));
    }

    function handleStateUpdate(state) {
        const wasRevealed = cardsRevealed; // Capture old state
        participants = state.participants;
        cardsRevealed = state.revealed;

        // If the state changed from revealed to not revealed, it's a new round.
        if (wasRevealed && !cardsRevealed) {
            resetLocalUI();
        }

        updateParticipantsUI();
        displayResults();
    }

    function setupConnection(conn) {
        conn.on('data', (data) => {
            if (isHost) {
                // Host handles messages from clients
                switch (data.type) {
                    case 'user-join':
                        participants[conn.peer] = { name: data.payload.name, vote: null };
                        connections[conn.peer] = conn;
                        broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
                        updateParticipantsUI();
                        break;
                    case 'user-vote':
                        if (participants[conn.peer]) {
                            participants[conn.peer].vote = data.payload.vote;
                            broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
                            updateParticipantsUI();
                            displayResults(); // BUG FIX: Recalculate results on host
                        }
                        break;
                    case 'user-leave':
                        if (participants[conn.peer]) {
                            delete participants[conn.peer];
                            delete connections[conn.peer];
                            broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
                            updateParticipantsUI();
                        }
                        break;
                }
            } else {
                // Client handles messages from host
                switch (data.type) {
                    case 'state-update':
                        handleStateUpdate(data.payload);
                        break;
                }
            }
        });

        conn.on('close', () => {
            if (isHost) {
                delete participants[conn.peer];
                delete connections[conn.peer];
                broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
                updateParticipantsUI();
            } else {
                // Client lost connection to host
                alert('Connection to the host was lost. Please refresh and rejoin.');
                roomSetup.classList.remove('d-none');
                pokerRoom.classList.add('d-none');
            }
        });

        // If I'm a client, send my info to the host once connected
        if (!isHost) {
            conn.on('open', () => {
                conn.send({ type: 'user-join', payload: { name: myName } });
            });
        }
    }

    function initialize(isCreatingRoom) {
        myName = userNameInput.value.trim();
        if (!myName) {
            alert('Please enter your name.');
            return;
        }

        peer = new Peer();
        isHost = isCreatingRoom;

        peer.on('open', (id) => {
            myId = id;
            const url = new URL(window.location);
            if (isHost) {
                roomIdDisplay.textContent = myId;
                participants[myId] = { name: myName, vote: null };
                adminControls.classList.remove('d-none');
                showPokerRoom();
                updateParticipantsUI();
                url.searchParams.set('room', myId);
                history.pushState({}, '', url);
            } else {
                const roomId = roomIdInput.value.trim();
                if (!roomId) {
                    alert('Please enter a Room ID to join.');
                    peer.destroy();
                    return;
                }
                roomIdDisplay.textContent = roomId;
                const conn = peer.connect(roomId);
                connections[roomId] = conn; // Store host connection
                setupConnection(conn);
                showPokerRoom();
                url.searchParams.set('room', roomId);
                history.pushState({}, '', url);
            }
        });

        peer.on('connection', (conn) => {
            if (isHost) {
                setupConnection(conn);
            }
        });

        peer.on('error', (err) => {
            console.error(err);
            alert('An error occurred: ' + err.message);
        });
    }

    // --- Event Listeners ---
    createRoomBtn.addEventListener('click', () => initialize(true));
    joinRoomBtn.addEventListener('click', () => initialize(false));

    votingCards.forEach(card => {
        card.addEventListener('click', (e) => {
            myVote = e.target.dataset.value;
            // UI feedback
            votingCards.forEach(c => c.classList.remove('btn-light', 'text-dark'));
            e.target.classList.add('btn-light', 'text-dark');

            if (isHost) {
                participants[myId].vote = myVote;
                broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
                updateParticipantsUI();
                displayResults(); // BUG FIX: Recalculate results on host
            } else {
                const hostConn = Object.values(connections)[0];
                if (hostConn) {
                    hostConn.send({ type: 'user-vote', payload: { vote: myVote } });
                }
            }
        });
    });

    revealCardsBtn.addEventListener('click', () => {
        if (!isHost) return;
        cardsRevealed = true;
        broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
        updateParticipantsUI();
        displayResults();
    });

    startVoteBtn.addEventListener('click', () => {
        if (!isHost) return;
        cardsRevealed = false;
        Object.keys(participants).forEach(id => {
            participants[id].vote = null;
        });
        resetLocalUI();
        broadcast({ type: 'state-update', payload: { participants, revealed: cardsRevealed } });
        updateParticipantsUI();
    });
});