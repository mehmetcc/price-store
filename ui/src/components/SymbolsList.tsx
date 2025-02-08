'use client';

import { useEffect, useState } from 'react';
import { fetchSymbols, deleteSymbol } from '../api';
import { TrashIcon } from '@heroicons/react/24/outline';
 
export default function SymbolsList() {
    const [symbols, setSymbols] = useState<string[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const [deleting, setDeleting] = useState<string | null>(null); // symbol being deleted

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

    const handleDelete = async (symbol: string) => {
        setDeleting(symbol);
        setError(null);
        try {
            await deleteSymbol(symbol);
            setSymbols((prev) => prev.filter((s) => s !== symbol));
        } catch (err: any) {
            setError(err.message || 'Error deleting symbol');
        } finally {
            setDeleting(null);
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
                    <li key={symbol} className="flex justify-end items-center">
                        <span className="text-lg">{symbol}</span>
                        <button
                            onClick={() => handleDelete(symbol)}
                            className="text-red-500 hover:text-red-700 ml-2"
                            disabled={deleting === symbol}
                        >
                            {deleting === symbol ? 'Deleting...' : <TrashIcon className="w-5 h-5" />}
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
}
