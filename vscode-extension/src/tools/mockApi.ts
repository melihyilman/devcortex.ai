import * as http from 'http';
import * as url from 'url';
import { generateRandomData } from './randomizer';

let server: http.Server | null = null;
let routeConfig: any = null;

function findFreePort(): Promise<number> {
    return new Promise((resolve, reject) => {
        const tempServer = http.createServer();
        tempServer.listen(0, () => {
            const port = (tempServer.address() as any).port;
            tempServer.close(() => resolve(port));
        });
        tempServer.on('error', reject);
    });
}

// ex: /products/123 -matches /products/{id} .
function findMatchingRoute(path: string): string | null {
    const pathSegments = path.split('/').filter(Boolean);

    for (const route in routeConfig) {
        const routeSegments = route.split('/').filter(Boolean);
        if (pathSegments.length !== routeSegments.length) continue;

        const isMatch = routeSegments.every((segment, i) => {
            return segment.startsWith('{') || segment === pathSegments[i];
        });

        if (isMatch) return route;
    }
    return null;
}

export async function startServer(template: any): Promise<number> {
    if (server) {
        throw new Error('Server is already running.');
    }

    routeConfig = template;
    const port = await findFreePort();

    server = http.createServer((req, res) => {
        const parsedUrl = url.parse(req.url || '', true);
        const path = parsedUrl.pathname || '/';
        const method = req.method?.toUpperCase() || 'GET';

        res.setHeader('Content-Type', 'application/json');
        res.setHeader('Access-Control-Allow-Origin', '*');
        res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');

        const matchedRoute = findMatchingRoute(path);

        if (matchedRoute && routeConfig[matchedRoute] && routeConfig[matchedRoute][method]) {
            const routeHandler = routeConfig[matchedRoute][method];
            if (routeHandler.response) {
                const randomData = generateRandomData(routeHandler.response);
                res.writeHead(200);
                res.end(JSON.stringify(randomData, null, 2));
            } else {
                res.writeHead(204); // No Content
                res.end();
            }
        } else {
            res.writeHead(404);
            res.end(JSON.stringify({ error: 'Not Found', message: `No mock defined for ${method} ${path}` }));
        }
    });

    return new Promise((resolve, reject) => {
        server?.listen(port, () => resolve(port));
        server?.on('error', reject);
    });
}

export function stopServer(): Promise<void> {
    return new Promise((resolve) => {
        if (server) {
            server.close(() => {
                server = null;
                routeConfig = null;
                resolve();
            });
        } else {
            resolve();
        }
    });
}