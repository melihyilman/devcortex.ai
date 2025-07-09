import * as vscode from 'vscode';
import * as crypto from 'crypto';

type HashAlgorithm = 'md5' | 'sha1' | 'sha256' | 'sha512';

async function performHash(algorithm: HashAlgorithm) {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }

    const text = editor.document.getText(editor.selection);
    if (!text) {
        vscode.window.showInformationMessage('Please select the text you want to hash.');
        return;
    }

    const hash = crypto.createHash(algorithm).update(text).digest('hex');
    
    vscode.window.showInformationMessage(`${algorithm.toUpperCase()} Hash: ${hash}`, 'Copy').then(choice => {
        if (choice === 'Copy') {
            vscode.env.clipboard.writeText(hash);
            vscode.window.showInformationMessage('Hash copied to clipboard!');
        }
    });
}

export function handleHashing() {
    const items: vscode.QuickPickItem[] = [
        { label: 'MD5', description: 'Generate an MD5 hash' },
        { label: 'SHA1', description: 'Generate a SHA1 hash' },
        { label: 'SHA256', description: 'Generate a SHA256 hash' },
        { label: 'SHA512', description: 'Generate a SHA512 hash' },
    ];

    vscode.window.showQuickPick(items, { placeHolder: 'Select a hash algorithm' }).then(selection => {
        if (!selection) {
            return;
        }
        performHash(selection.label.toLowerCase() as HashAlgorithm);
    });
}