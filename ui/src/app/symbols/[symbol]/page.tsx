'use client';

import { useEffect, useState } from 'react';
import { fetchPriceUpdatesBySymbol, fetchFilteredPriceCount } from '@/api';
import { PriceUpdate } from '@/types';
import Pagination from '@/components/Pagination';
import config from '@/config/config';
import Link from 'next/link';

interface SymbolPageProps {
    params: {
        symbol: string;
    };
}

export default function SymbolPage({ params }: SymbolPageProps) {
    const [prices, setPrices] = useState<PriceUpdate[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    const loadPrices = async (page: number) => {
        setLoading(true);
        setError(null);
        try {
            const [updates, count] = await Promise.all([
                fetchPriceUpdatesBySymbol(params.symbol, page, config.pageSize),
                fetchFilteredPriceCount(params.symbol)
            ]);
            setPrices(updates);
            setTotalPages(Math.ceil(count / config.pageSize));
            setCurrentPage(page);
        } catch (err: any) {
            setError(err.message || 'Error fetching price updates');
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadPrices(currentPage);
    }, [currentPage, params.symbol]);

    return (
        <div>
            <div className="container mx-auto p-4">
                <div className="flex justify-between items-center mb-4">
                    <h1 className="text-2xl font-bold">Price History for {params.symbol}</h1>
                    <Link
                        href="/prices"
                        className="px-4 py-2 text-sm bg-gray-200 rounded hover:bg-gray-300"
                    >
                        ← Back to All Prices
                    </Link>
                </div>
            </div>

            {loading && <p>Loading price updates...</p>}
            {error && <p className="text-red-500">{error}</p>}

            <table className="min-w-full border mt-4">
                <thead>
                    <tr>
                        <th className="px-4 py-2 border">ID</th>
                        <th className="px-4 py-2 border">Price</th>
                        <th className="px-4 py-2 border">Timestamp</th>
                    </tr>
                </thead>
                <tbody>
                    {prices.map((price) => (
                        <tr key={price.id}>
                            <td className="px-4 py-2 border">{price.id}</td>
                            <td className="px-4 py-2 border">{price.price.toFixed(2)}</td>
                            <td className="px-4 py-2 border">
                                {new Date(price.timestamp).toLocaleString()}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            <div className="mt-4 flex justify-center">
                <div className="w-full max-w-4xl">
                    <Pagination
                        currentPage={currentPage}
                        totalPages={totalPages}
                        onPageChange={loadPrices}
                    />
                </div>
            </div>
        </div>
    );
}