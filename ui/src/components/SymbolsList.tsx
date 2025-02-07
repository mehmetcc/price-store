'use client';

import { useEffect, useState } from 'react';
import { fetchSymbols } from '../api';

export default function SymbolsList() {
    const [symbols, setSymbols] = useState<string[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const loadSymbols = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await fetchSymbols();
            setSymbols(data);
        } catch (err: any) {
            setError(err.message || 'Error fetching symbols');
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadSymbols();
    }, []);

    return (
        <div>
            {loading && <p>Loading symbols...</p>}
            {error && <p className="text-red-500">{error}</p>}
            <ul className="list-disc ml-5">
                {symbols.map((symbol) => (
                    <li key={symbol}>{symbol}</li>
                ))}
            </ul>
        </div>
    );
}
