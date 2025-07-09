import { v4 as uuidv4 } from 'uuid';

const firstNames = ['James', 'Mary', 'Robert', 'Patricia', 'John', 'Jennifer', 'Michael', 'Linda', 'David', 'Elizabeth', 'William', 'Barbara', 'Richard', 'Susan', 'Joseph', 'Jessica', 'Thomas', 'Sarah', 'Ahmet', 'Mehmet', 'Elif', 'Ayşe', 'Mustafa', 'Fatma'];
const lastNames = ['Smith', 'Johnson', 'Williams', 'Brown', 'Jones', 'Garcia', 'Miller', 'Davis', 'Rodriguez', 'Martinez', 'Hernandez', 'Lopez', 'Gonzalez', 'Wilson', 'Anderson', 'Thomas', 'Taylor', 'Moore', 'Yılmaz', 'Kaya', 'Demir', 'Şahin', 'Çelik', 'Arslan'];
const cities = ['New York', 'London', 'Tokyo', 'Paris', 'Dubai', 'Singapore', 'Los Angeles', 'Berlin', 'Sydney', 'Toronto', 'Chicago', 'Istanbul', 'Ankara', 'Izmir', 'Bursa', 'Antalya', 'Adana', 'Gaziantep', 'Konya', 'Trabzon', 'Diyarbakır'];
const companies = ['Devcortex', 'Acme Inc', 'Stark Industries', 'Wayne Enterprises', 'Cyberdyne Systems'];
const domains = ['example.com', 'test.dev', 'mail.io', 'devcortex.ai'];
const sampleWords = ['lorem', 'ipsum', 'dolor', 'sit', 'amet', 'consectetur', 'adipiscing', 'elit'];

// --- Helper Functions ---
function getRand<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateSmartString(key: string, value: string): string {
    const lowerKey = key.toLowerCase();
    const lowerValue = typeof value === 'string' ? value.toLowerCase() : '';

    if (lowerValue === 'uuid' || lowerKey === 'uuid' || (lowerKey.endsWith('id') && lowerValue !== 'number')) return uuidv4();
    if (lowerValue === 'email' || lowerKey.includes('email')) return `${getRand(firstNames).toLowerCase()}.${getRand(lastNames).toLowerCase()}@${getRand(domains)}`;
    if (lowerValue === 'fullname' || lowerKey.includes('name')) return `${getRand(firstNames)} ${getRand(lastNames)}`;
    if (lowerValue === 'firstname' || lowerKey.includes('firstname')) return getRand(firstNames);
    if (lowerValue === 'lastname' || lowerKey.includes('lastname')) return getRand(lastNames);
    if (lowerValue === 'city' || lowerKey.includes('city')) return getRand(cities);
    if (lowerValue === 'company' || lowerKey.includes('company')) return getRand(companies);
    if (lowerValue === 'phone' || lowerKey.includes('phone')) return `+${Math.floor(Math.random() * 90 + 10)} ${Math.floor(Math.random() * 900 + 100)} ${Math.floor(Math.random() * 9000 + 1000)}`;
    if (lowerValue === 'url' || lowerKey.includes('url') || lowerKey.includes('website')) return `https://${getRand(companies).toLowerCase().replace(/ /g, '-')}.com`;
    if (lowerValue === 'image' || lowerValue === 'avatar' || lowerKey.includes('image') || lowerKey.includes('avatar') || lowerKey.includes('picture')) return `https://placehold.co/600x400?text=${getRand(firstNames)}`;

    // Default: generate a random sentence
    const wordCount = Math.floor(Math.random() * 5) + 3;
    let sentence = Array.from({ length: wordCount }, () => getRand(sampleWords)).join(' ');
    return sentence.charAt(0).toUpperCase() + sentence.slice(1);
}

// --- Main Generator Function ---
export function generateRandomData(template: any, key: string = ''): any {
    if (template === null || template === undefined) return null;

    const type = typeof template;

    switch (type) {
        case 'string':
            return generateSmartString(key, template);
        case 'number':
            return Math.floor(Math.random() * 1000);
        case 'boolean':
            return Math.random() > 0.5;
        case 'object':
            if (Array.isArray(template)) {
                const arrayTemplate = template[0];
                if (arrayTemplate === undefined) return [];
                const arrayLength = Math.floor(Math.random() * 5) + 1;
                return Array.from({ length: arrayLength }, () => generateRandomData(arrayTemplate, key));
            }
            const newObj: { [key: string]: any } = {};
            for (const objKey in template) {
                if (Object.prototype.hasOwnProperty.call(template, objKey)) {
                    newObj[objKey] = generateRandomData(template[objKey], objKey);
                }
            }
            return newObj;
        default:
            return null;
    }
}