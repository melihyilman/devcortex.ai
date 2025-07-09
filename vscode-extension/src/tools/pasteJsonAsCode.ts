import * as vscode from 'vscode';
import { quicktype, InputData, jsonInputForTargetLanguage, TargetLanguage } from 'quicktype-core';

interface LanguageQuickPickItem extends vscode.QuickPickItem {
    lang: TargetLanguage;
    rendererOptions?: any;
}

const supportedLanguages: LanguageQuickPickItem[] = [
    { label: "C#", description: "Paste JSON as C# classes", lang: "csharp" as unknown as TargetLanguage, rendererOptions: { namespace: "Devcortex.Generated", features: 'just-types' } },
    { label: "Java", description: "Paste JSON as Java classes", lang: "java" as unknown as TargetLanguage, rendererOptions: { package: "com.devcortex.generated", features: 'just-types' } },
    { label: "TypeScript", description: "Paste JSON as TypeScript interfaces", lang: "ts" as unknown as TargetLanguage, rendererOptions: { features: 'just-types' } },
    { label: "Go", description: "Paste JSON as Go structs", lang: "go" as unknown as TargetLanguage, rendererOptions: { features: 'just-types' } },
];

async function convert(jsonString: string, language: LanguageQuickPickItem) {
    const jsonInput = jsonInputForTargetLanguage(language.lang);
    await jsonInput.addSource({ name: "Root", samples: [jsonString] });

    const inputData = new InputData();
    inputData.addInput(jsonInput);

    const { lines } = await quicktype({
        inputData,
        lang: language.lang,
        rendererOptions: language.rendererOptions
    });

    return lines.join('\n');
}

export async function pasteJsonAsCode() {
    const clipboardText = await vscode.env.clipboard.readText();
    if (!clipboardText) {
        vscode.window.showErrorMessage('Clipboard is empty.');
        return;
    }

    try {
        JSON.parse(clipboardText);
    } catch (error) {
        vscode.window.showErrorMessage('The text in your clipboard is not a valid JSON.');
        return;
    }

    const selectedLanguage = await vscode.window.showQuickPick(supportedLanguages, {
        placeHolder: "Select the target language to paste JSON as"
    });

    if (!selectedLanguage) { return; }

    try {
        const code = await convert(clipboardText, selectedLanguage);
        const doc = await vscode.workspace.openTextDocument({ content: code, language: selectedLanguage.label.toLowerCase() });
        await vscode.window.showTextDocument(doc, { preview: false });
    } catch (e: any) {
        vscode.window.showErrorMessage(`Failed to convert JSON: ${e.message}`);
    }
}