import React, { useState, useEffect } from 'react';
import { fetchSymbols, fetchPriceUpdatesBySymbol, fetchPriceCount, fetchFilteredPriceCount } from '@/api';
import { PriceUpdate } from '@/types';

interface SymbolFilterProps {
  onPriceUpdatesChange: (updates: PriceUpdate[], totalCount: number, symbol: string) => void;
  currentPage: number;
  pageSize: number;
}

export const SymbolFilter: React.FC<SymbolFilterProps> = ({
  onPriceUpdatesChange,
  currentPage,
  pageSize
}) => {
  const [symbols, setSymbols] = useState<string[]>([]);
  const [selectedSymbol, setSelectedSymbol] = useState<string>('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    loadSymbols();
  }, []);

  const loadSymbols = async () => {
    try {
      const availableSymbols = await fetchSymbols();
      setSymbols(availableSymbols);
    } catch (error) {
      console.error('Failed to load symbols:', error);
    }
  };

  const handleSymbolChange = async (event: React.ChangeEvent<HTMLSelectElement>) => {
    const symbol = event.target.value;
    setSelectedSymbol(symbol);

    if (symbol) {
      setLoading(true);
      try {
        const updates = await fetchPriceUpdatesBySymbol(symbol, currentPage, pageSize);
        const count = await fetchFilteredPriceCount(symbol);
        onPriceUpdatesChange(updates, count, symbol);
      } catch (error) {
        console.error('Failed to fetch price updates:', error);
      } finally {
        setLoading(false);
      }
    }
  };



  return (
    <div className="flex items-center gap-4">
      <select
        value={selectedSymbol}
        onChange={handleSymbolChange}
        className="px-4 py-2 border rounded-md"
        disabled={loading}
      >
        <option value="">Select a symbol</option>
        {symbols.map((symbol) => (
          <option key={symbol} value={symbol}>
            {symbol}
          </option>
        ))}
      </select>
      {loading && <span>Loading...</span>}
    </div>
  );
};