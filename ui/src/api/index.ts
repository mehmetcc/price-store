import config from '@/config/config';
import { SymbolResponse, PriceUpdatesResponse, PriceCountResponse, PriceUpdate } from '../types';

export async function fetchSymbols(): Promise<string[]> {
    const response = await fetch(`${config.baseUrl}/symbol`);
    if (!response.ok) {
        throw new Error('Error fetching symbols');
    }
    const data: SymbolResponse = await response.json();
    return data.symbols;
}

export async function addSymbol(symbol: string): Promise<void> {
    const response = await fetch(`${config.baseUrl}/symbol`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ symbol }),
    });
    if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || 'Error adding symbol');
    }
    // Parse and ignore the response (contains status and symbol)
    await response.json();
}

export async function deleteSymbol(symbol: string): Promise<void> {
    const response = await fetch(`${config.baseUrl}/symbol`, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ symbol }),
    });
    if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || 'Error deleting symbol');
    }
    // Parse and ignore the response (contains status and symbol)
    await response.json();
}

export async function fetchPriceUpdates(page: number, pageSize: number): Promise<PriceUpdate[]> {
    const response = await fetch(`${config.baseUrl}/price?page=${page}&pageSize=${pageSize}`);
    if (!response.ok) {
        throw new Error('Error fetching price updates');
    }
    const data: PriceUpdatesResponse = await response.json();
    return data.price_updates;
}

export async function fetchPriceCount(): Promise<number> {
    const response = await fetch(`${config.baseUrl}/price/count`);
    if (!response.ok) {
        throw new Error('Error fetching price count');
    }
    const data: PriceCountResponse = await response.json();
    return data.count;
}
