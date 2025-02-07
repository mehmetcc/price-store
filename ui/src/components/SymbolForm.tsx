'use client';

import { useState } from 'react';
import { addSymbol } from '../api';

interface SymbolFormProps {
    onSymbolAdded: () => void;
}

export default function SymbolForm({ onSymbolAdded }: SymbolFormProps) {
    const [symbol, setSymbol] = useState('');
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!symbol.trim()) {
            setError('Symbol is required');
            return;
        }
        setLoading(true);
        setError(null);
        try {
            await addSymbol(symbol.trim());
            setSymbol('');
            onSymbolAdded();
        } catch (err: any) {
            setError(err.message || 'Error adding symbol');
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="mb-4">
            <div className="flex">
                <input
                    type="text"
                    value={symbol}
                    onChange={(e) => setSymbol(e.target.value)}
                    placeholder="Enter symbol"
                    className="flex-1 border p-2 mr-2"
                />
                <button type="submit" className="bg-blue-500 text-white p-2" disabled={loading}>
                    {loading ? 'Adding...' : 'Add Symbol'}
                </button>
            </div>
            {error && <p className="text-red-500 mt-2">{error}</p>}
        </form>
    );
}
