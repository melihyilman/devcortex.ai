import * as vscode from 'vscode';
import { quicktype, InputData, jsonInputForTargetLanguage } from 'quicktype-core';

async function jsonToCSharp(jsonString: string): Promise<string> {
    const jsonInput = jsonInputForTargetLanguage('csharp');
    
    await jsonInput.addSource({
        name: 'Root',
        samples: [jsonString],
    });

    const inputData = new InputData();
    inputData.addInput(jsonInput);

    const { lines } = await quicktype({
        inputData,
        lang: 'csharp',
        rendererOptions: {
            'namespace': 'Devcortex.Generated',
            'features': 'just-types'
        }
    });

    return lines.join('\n');
}

export async function convertJsonToCsharp() {
    const editor = vscode.window.activeTextEditor;
    
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }

    const selection = editor.selection;
    const jsonText = selection.isEmpty 
        ? editor.document.getText() 
        : editor.document.getText(selection);

    if (!jsonText) {
        vscode.window.showErrorMessage('No JSON text found. Select some JSON or open a JSON file.');
        return;
    }

    try {
        const csharpCode = await jsonToCSharp(jsonText);
        const doc = await vscode.workspace.openTextDocument({
            content: csharpCode,
            language: 'csharp'
        });
        
        await vscode.window.showTextDocument(doc, { preview: false });
    } catch (e: any) {
        vscode.window.showErrorMessage(`Failed to convert JSON to C#: ${e.message}`);
    }
}