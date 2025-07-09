import * as vscode from 'vscode';
import { generateUuid } from './tools/generateUuid';
import { convertJsonToCsharp } from './tools/jsonToCsharp';
import { handleBase64Conversion } from './tools/base64';
import { handleHashing } from './tools/hashing';
import { handleBcrypt } from './tools/bcrypt';
import { handleHmac } from './tools/hmac';
import { decodeJwt } from './tools/jwt';
import { handleUrlEncoding } from './tools/url';
import { pasteJsonAsCode } from './tools/pasteJsonAsCode';
import { startServer, stopServer } from './tools/mockApi';

let mockApiStatusBarItem: vscode.StatusBarItem;

interface ToolItem extends vscode.QuickPickItem {
    id: string;
    action: () => void;
}

const tools: (ToolItem | { kind: vscode.QuickPickItemKind.Separator; label: string })[] = [
    { label: 'Generators', kind: vscode.QuickPickItemKind.Separator },
    {
        id: 'generate-uuid',
        label: '$(symbol-key) Generate Random UUID',
        description: 'Generates a new v4 UUID and copies it to the clipboard',
        action: generateUuid
    },
    {
        id: 'start-mock-api',
        label: '$(server-process) Start Mock API from Selection',
        description: 'Starts a mock API server based on selected JSON',
        action: handleStartMockApi
    },
    { label: 'Converters', kind: vscode.QuickPickItemKind.Separator },
    {
        id: 'paste-json-as-code',
        label: '$(file-code) Paste JSON as Code',
        description: 'Converts JSON in clipboard to C#, Java, TS, or Go',
        action: pasteJsonAsCode
    },
    {
        id: 'json-to-csharp',
        label: '$(json) JSON to C# (Selection)',
        description: 'Converts selected JSON to C# classes',
        action: convertJsonToCsharp
    },
    { label: 'Encoders / Decoders', kind: vscode.QuickPickItemKind.Separator },
    {
        id: 'base64',
        label: '$(lock) Base64 Encoder / Decoder',
        description: 'Encode or decode selected text using Base64',
        action: handleBase64Conversion
    },
    {
        id: 'url-encoder-decoder',
        label: '$(link) URL Encoder / Decoder',
        description: 'Encode or decode selected text for URLs',
        action: handleUrlEncoding
    },
    {
        id: 'jwt-decoder',
        label: '$(key) JWT Decoder',
        description: 'Decode a JSON Web Token (JWT)',
        action: decodeJwt
    },
    { label: 'Hashing & Cryptography', kind: vscode.QuickPickItemKind.Separator },
    {
        id: 'hashing',
        label: '$(hash) Hash Generator',
        description: 'Generate MD5, SHA1, SHA256, SHA512 hashes',
        action: handleHashing
    },
    {
        id: 'hmac',
        label: '$(key) HMAC Generator',
        description: 'Generate HMAC-SHA256 signature',
        action: handleHmac
    },
    {
        id: 'bcrypt',
        label: '$(shield) Bcrypt',
        description: 'Generate and verify bcrypt hashes',
        action: handleBcrypt
    },
    { label: 'Help', kind: vscode.QuickPickItemKind.Separator },
    {
        id: 'show-documentation',
        label: '$(book) Show Documentation',
        description: 'Opens the tool documentation',
        action: () => vscode.commands.executeCommand('devcortex.showDocumentation')
    },
];

const sampleTemplate = `{
    "/users": {
        "GET": {
            "response": [
                {
                    "userId": "uuid",
                    "fullName": "fullName",
                    "email": "email",
                    "company": "company",
                    "isActive": true,
                    "profileImage": "avatar"
                }
            ]
        },
        "POST": {
            "request": {
                "fullName": "string",
                "email": "string"
            },
            "response": {
                "userId": "uuid",
                "status": "created"
            }
        }
    },
    "/products/{id}": {
        "GET": {
            "response": {
                "productId": "number",
                "productName": "string",
                "price": 123.45,
                "productImage": "image",
                "supplier_website": "url"
            }
        }
    }
}`;

async function showSampleTemplate() {
    const doc = await vscode.workspace.openTextDocument({ content: sampleTemplate, language: 'json' });
    await vscode.window.showTextDocument(doc, { preview: false });
    vscode.window.showInformationMessage('No valid template found. Here is a sample to get you started. Modify it, select the text, and run the command again.');
}

async function handleStartMockApi() {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
        vscode.window.showErrorMessage('No active editor found.');
        return;
    }
    const selection = editor.selection;
    const jsonText = editor.document.getText(selection);

    let jsonTemplate;
    try {
        if (!jsonText) throw new Error('No text selected');
        jsonTemplate = JSON.parse(jsonText);
    } catch (e) {
        showSampleTemplate();
        return;
    }

    try {
        const port = await startServer(jsonTemplate);
        const url = `http://localhost:${port}`;

        mockApiStatusBarItem.text = `$(server-process) Mock API: ${port}`;
        mockApiStatusBarItem.tooltip = `Mock API is running on port ${port}. Click to stop.`;
        mockApiStatusBarItem.command = 'devcortex.stopMockApi';
        mockApiStatusBarItem.show();

        const choice = await vscode.window.showInformationMessage(
            `Mock API running at ${url}`,
            'Stop Server', 'Copy URL', 'Copy cURL'
        );

        if (choice === 'Stop Server') {
            vscode.commands.executeCommand('devcortex.stopMockApi');
        } else if (choice === 'Copy URL') {
            vscode.env.clipboard.writeText(url);
            vscode.window.showInformationMessage('URL copied to clipboard!');
        } else if (choice === 'Copy cURL') {
            vscode.env.clipboard.writeText(`curl ${url}`);
            vscode.window.showInformationMessage('cURL command copied to clipboard!');
        }

    } catch (error: any) {
        vscode.window.showErrorMessage(`Failed to start mock API: ${error.message}`);
    }
}

async function handleStopMockApi() {
    try {
        await stopServer();
        mockApiStatusBarItem.hide();
        vscode.window.showInformationMessage('Mock API stopped successfully.');
    } catch (error: any) {
        vscode.window.showErrorMessage(`Failed to stop mock API: ${error.message}`);
    }
}


function showToolsMenu() {
    const toolItems: ToolItem[] = tools
        .filter(tool => tool.kind !== vscode.QuickPickItemKind.Separator)
        .map(tool => tool as ToolItem);

    vscode.window.showQuickPick(tools, { placeHolder: 'Select a Devcortex tool' }).then(selection => {
        if (!selection || selection.kind === vscode.QuickPickItemKind.Separator) {
            return;
        }
        const selectedTool = toolItems.find(t => t.id === (selection as ToolItem).id);
        if (selectedTool) {
            selectedTool.action();
        }
    });
}

export function activate(context: vscode.ExtensionContext) {
    console.log('Devcortex Development Tools is now active!');

    const showDocumentationCommand = vscode.commands.registerCommand('devcortex.showDocumentation', () => {
        const docPath = vscode.Uri.joinPath(context.extensionUri, 'TOOLS.md');
        vscode.commands.executeCommand('markdown.showPreview', docPath);
    });
    context.subscriptions.push(showDocumentationCommand);

    context.subscriptions.push(vscode.commands.registerCommand('devcortex.showTools', showToolsMenu));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.generateUuid', generateUuid));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.jsonToCsharp', convertJsonToCsharp));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.base64', handleBase64Conversion));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.url', handleUrlEncoding));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.jwtDecode', decodeJwt));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.hashing', handleHashing));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.hmac', handleHmac));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.bcrypt', handleBcrypt));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.pasteJsonAsCode', pasteJsonAsCode));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.startMockApi', handleStartMockApi));
    context.subscriptions.push(vscode.commands.registerCommand('devcortex.stopMockApi', handleStopMockApi));

    // Show welcome page on first install/update
    if (context.globalState.get('devcortex.hasShownWelcome') !== true) {
        setTimeout(() => {
            vscode.commands.executeCommand('devcortex.showDocumentation');
        }, 1000); // 1 second delay to not be too intrusive
        context.globalState.update('devcortex.hasShownWelcome', true);
    }

    // Create the main Status Bar icon.
    // Note: Custom SVGs are not directly supported in the status bar text.
    // A suitable Codicon is used instead to represent the extension.
    const mainStatusBarItem = vscode.window.createStatusBarItem(vscode.StatusBarAlignment.Right, 100);
    mainStatusBarItem.command = 'devcortex.showTools';
    mainStatusBarItem.text = `$(hubot) Devcortex`;
    mainStatusBarItem.tooltip = 'Open Devcortex Tools';
    mainStatusBarItem.show();
    context.subscriptions.push(mainStatusBarItem);

    mockApiStatusBarItem = vscode.window.createStatusBarItem(vscode.StatusBarAlignment.Right, 99);
    mockApiStatusBarItem.hide();
    context.subscriptions.push(mockApiStatusBarItem);
}

export function deactivate() {
    stopServer();
}