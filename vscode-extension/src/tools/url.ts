import * as vscode from 'vscode';

function performUrlAction(action: 'Encode' | 'Decode') {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }

    const selection = editor.selection;
    const text = editor.document.getText(selection);

    if (!text) {
        vscode.window.showInformationMessage('Please select the text you want to ' + action.toLowerCase() + '.');
        return;
    }

    let result: string;
    if (action === 'Encode') {
        result = encodeURIComponent(text);
    } else { // Decode
        result = decodeURIComponent(text);
    }

    editor.edit(editBuilder => {
        editBuilder.replace(selection, result);
    });
}

export function handleUrlEncoding() {
    const items: vscode.QuickPickItem[] = [
        { label: 'Encode', description: 'Encode selected text for URL' },
        { label: 'Decode', description: 'Decode selected URL-encoded text' }
    ];

    vscode.window.showQuickPick(items, { placeHolder: 'Select a URL action' }).then(selection => {
        if (!selection) {
            return;
        }
        performUrlAction(selection.label as 'Encode' | 'Decode');
    });
}