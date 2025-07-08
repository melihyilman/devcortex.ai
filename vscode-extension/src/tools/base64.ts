import * as vscode from 'vscode';

async function performBase64Action(action: 'Encode' | 'Decode') {
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

    try {
        let result: string;
        if (action === 'Encode') {
            result = Buffer.from(text, 'utf8').toString('base64');
        } else { // Decode
            result = Buffer.from(text, 'base64').toString('utf8');
        }

        editor.edit(editBuilder => {
            editBuilder.replace(selection, result);
        });

        vscode.window.showInformationMessage(`Text successfully ${action.toLowerCase()}d.`);

    } catch (e: any) {
        vscode.window.showErrorMessage(`Failed to ${action.toLowerCase()} Base64: ${e.message}`);
    }
}

export function handleBase64Conversion() {
    const items: vscode.QuickPickItem[] = [
        { label: 'Encode', description: 'Convert selected text to Base64' },
        { label: 'Decode', description: 'Convert selected Base64 text to UTF-8' }
    ];

    vscode.window.showQuickPick(items, { placeHolder: 'Select a Base64 action' }).then(selection => {
        if (!selection) {
            return;
        }
        performBase64Action(selection.label as 'Encode' | 'Decode');
    });
}