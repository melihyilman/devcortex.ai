import * as vscode from 'vscode';
import { v4 as uuidv4 } from 'uuid';

export function generateUuid() {
    const newUuid = uuidv4();
    vscode.env.clipboard.writeText(newUuid);
    vscode.window.showInformationMessage(`UUID Copied: ${newUuid}`);
}