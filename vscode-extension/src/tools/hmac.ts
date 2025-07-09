import * as vscode from 'vscode';
import * as crypto from 'crypto';

export async function handleHmac() {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }

    const text = editor.document.getText(editor.selection);
    if (!text) {
        vscode.window.showInformationMessage('Please select the text you want to sign.');
        return;
    }

    const secretKey = await vscode.window.showInputBox({ prompt: 'Enter your HMAC secret key' });
    if (!secretKey) { return; }

    const hmac = crypto.createHmac('sha256', secretKey).update(text).digest('hex');

    vscode.window.showInformationMessage(`HMAC-SHA256: ${hmac}`, 'Copy').then(choice => {
        if (choice === 'Copy') {
            vscode.env.clipboard.writeText(hmac);
            vscode.window.showInformationMessage('HMAC copied to clipboard!');
        }
    });
}