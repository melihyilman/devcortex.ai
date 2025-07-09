import * as vscode from 'vscode';
import * as bcrypt from 'bcryptjs';

async function generateBcryptHash() {
    const textToHash = await vscode.window.showInputBox({ prompt: 'Enter the string to hash' });
    if (!textToHash) { return; }

    const saltRoundsStr = await vscode.window.showInputBox({ prompt: 'Enter salt rounds (e.g., 10)', value: '10' });
    const saltRounds = parseInt(saltRoundsStr || '10', 10);

    if (isNaN(saltRounds)) {
        vscode.window.showErrorMessage('Invalid salt rounds.');
        return;
    }

    const hash = await bcrypt.hash(textToHash, saltRounds);
    vscode.window.showInformationMessage(`Bcrypt Hash: ${hash}`, 'Copy').then(choice => {
        if (choice === 'Copy') {
            vscode.env.clipboard.writeText(hash);
            vscode.window.showInformationMessage('Bcrypt hash copied to clipboard!');
        }
    });
}

async function verifyBcryptHash() {
    const textToVerify = await vscode.window.showInputBox({ prompt: 'Enter the original string' });
    if (!textToVerify) { return; }

    const hashToVerify = await vscode.window.showInputBox({ prompt: 'Enter the bcrypt hash to verify against' });
    if (!hashToVerify) { return; }

    const isMatch = await bcrypt.compare(textToVerify, hashToVerify);
    vscode.window.showInformationMessage(isMatch ? '✅ The string and hash match.' : '❌ The string and hash DO NOT match.');
}

export function handleBcrypt() {
    const items: vscode.QuickPickItem[] = [
        { label: 'Generate Hash', description: 'Create a new bcrypt hash' },
        { label: 'Verify Hash', description: 'Compare a string against a bcrypt hash' },
    ];

    vscode.window.showQuickPick(items, { placeHolder: 'Select a bcrypt action' }).then(selection => {
        if (!selection) { return; }
        if (selection.label === 'Generate Hash') {
            generateBcryptHash();
        } else {
            verifyBcryptHash();
        }
    });
}