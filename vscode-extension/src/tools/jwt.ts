import * as vscode from 'vscode';
import { jwtDecode } from 'jwt-decode';

export async function decodeJwt() {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }

    const token = editor.document.getText(editor.selection);
    if (!token) {
        vscode.window.showInformationMessage('Please select the JWT you want to decode.');
        return;
    }

    try {
        const decodedHeader = jwtDecode(token, { header: true });
        const decodedPayload = jwtDecode(token);

        const formattedOutput = 
`// HEADER: ALGORITHM & TOKEN TYPE
${JSON.stringify(decodedHeader, null, 2)}

// PAYLOAD: DATA
${JSON.stringify(decodedPayload, null, 2)}
`;

        const doc = await vscode.workspace.openTextDocument({ content: formattedOutput, language: 'json' });
        await vscode.window.showTextDocument(doc, { preview: false });

    } catch (e: any) {
        vscode.window.showErrorMessage(`Failed to decode JWT: ${e.message}`);
    }
}