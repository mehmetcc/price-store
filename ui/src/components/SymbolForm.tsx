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
        <>
            <form onSubmit={handleSubmit} className="mb-4">
                <div className="flex">
                    <input
                        type="text"
                        value={symbol}
                        onChange={(e) => setSymbol(e.target.value)}
                        placeholder="Enter symbol"
                        className="flex-1 border p-2 mr-2"
                    />
                    <button type="submit" className="bg-indigo-600 text-white p-2" disabled={loading}>
                        {loading ? 'Adding...' : 'Add Symbol'}
                    </button>
                </div>
            </form>

            {error && (
                <div className="fixed inset-0 flex items-center justify-center z-50">
                    <div className="fixed inset-0 bg-black opacity-50"></div>
                    <div className="bg-white p-6 rounded shadow-lg w-80 z-10">
                        <h2 className="text-xl font-bold mb-4 text-red-500">Error</h2>
                        <p className="mb-4">{error}</p>
                        <button
                            onClick={() => setError(null)}
                            className="bg-indigo-600 text-white px-4 py-2 rounded"
                        >
                            Close
                        </button>
                    </div>
                </div>
            )}
        </>
    );
}
