'use client';

import { useEffect, useState } from 'react';
import { fetchPriceUpdates, fetchPriceCount } from '../api';
import { PriceUpdate } from '../types';
import Pagination from './Pagination';
import { SymbolFilter } from './SymbolFilter';
import config from '@/config/config';

export default function PastPricesList() {
    const [prices, setPrices] = useState<PriceUpdate[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [isFiltered, setIsFiltered] = useState(false);

    const loadPrices = async (page: number) => {
        if (isFiltered) return; // Don't load paginated prices when filtering by symbol
        setLoading(true);
        setError(null);
        try {
            const [priceData, count] = await Promise.all([
                fetchPriceUpdates(page, config.pageSize),
                fetchPriceCount(),
            ]);
            setPrices(priceData);
            setTotalPages(Math.ceil(count / config.pageSize));
            setCurrentPage(page);
        } catch (err: any) {
            setError(err.message || 'Error fetching price updates');
        } finally {
            setLoading(false);
        }
    };

    const handleSymbolPriceUpdates = (updates: PriceUpdate[]) => {
        setPrices(updates);
        setIsFiltered(true);
        setTotalPages(1); // Reset pagination when filtering
    };

    const handleSymbolClear = () => {
        setIsFiltered(false);
        loadPrices(1);
    };

    useEffect(() => {
        if (!isFiltered) {
            loadPrices(currentPage);
        }
    }, [currentPage, isFiltered]);

    return (
        <div>
            <div className="mb-4 flex items-center justify-end">
                <SymbolFilter onPriceUpdatesChange={handleSymbolPriceUpdates} />
                {isFiltered && (
                    <button
                        onClick={handleSymbolClear}
                        className="ml-4 px-3 py-1 text-sm bg-gray-200 rounded hover:bg-gray-300"
                    >
                        Clear Filter
                    </button>
                )}
            </div>

            {loading && <p>Loading price updates...</p>}
            {error && <p className="text-red-500">{error}</p>}
            
            <table className="min-w-full border mt-4">
                {/* ...existing table header... */}
                <tbody>
                    {prices.map((price) => (
                        <tr key={price.id}>
                            <td className="px-4 py-2 border">{price.id}</td>
                            <td className="px-4 py-2 border">{price.symbol}</td>
                            <td className="px-4 py-2 border">{price.price.toFixed(2)}</td>
                            <td className="px-4 py-2 border">
                                {new Date(price.timestamp).toLocaleString()}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            {!isFiltered && (
                <div className="mt-4 flex justify-center">
                    <div className="w-full max-w-4xl">
                        <Pagination
                            currentPage={currentPage}
                            totalPages={totalPages}
                            onPageChange={loadPrices}
                        />
                    </div>
                </div>
            )}
        </div>
    );
}