'use client';

import { useEffect, useState } from 'react';
import { fetchPriceUpdates, fetchPriceCount } from '../api';
import { PriceUpdate } from '../types';
import Pagination from './Pagination';
import config from '@/config/config';

export default function PastPricesList() {
    const [prices, setPrices] = useState<PriceUpdate[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    const loadPrices = async (page: number) => {
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

    useEffect(() => {
        loadPrices(currentPage);
    }, [currentPage]);

    useEffect(() => {
        console.log("Prices", prices)
    }, [prices]);


    return (
        <div>
            {loading && <p>Loading price updates...</p>}
            {error && <p className="text-red-500">{error}</p>}
            <table className="min-w-full border mt-4">
                <thead className="bg-gray-50">
                    <tr>
                        <th className="px-4 py-2 border">ID</th>
                        <th className="px-4 py-2 border">Symbol</th>
                        <th className="px-4 py-2 border">Price</th>
                        <th className="px-4 py-2 border">Timestamp</th>
                    </tr>
                </thead>
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
