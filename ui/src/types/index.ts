export interface SymbolResponse {
    symbols: string[];
}

export interface PriceUpdate {
    id: number;
    symbol: string;
    price: number;
    timestamp: string;
}

export interface PriceUpdatesResponse {
    price_updates: PriceUpdate[];
}

export interface PriceCountResponse {
    count: number;
}
