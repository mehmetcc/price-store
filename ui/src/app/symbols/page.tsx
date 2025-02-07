'use client';

import { useState } from 'react';
import SymbolForm from '../../components/SymbolForm';
import SymbolsList from '../../components/SymbolsList';

export default function SymbolsPage() {
    // A simple way to trigger a refresh of the symbols list after adding one.
    const [refreshCount, setRefreshCount] = useState(0);

    const handleSymbolAdded = () => {
        setRefreshCount((prev) => prev + 1);
    };

    return (
        <div>
            <h1 className="text-2xl font-bold mb-4">Tracked Symbols</h1>
            <SymbolForm onSymbolAdded={handleSymbolAdded} />
            {/* By passing the refresh count as key, the SymbolsList reloads when updated */}
            <SymbolsList key={refreshCount} />
        </div>
    );
}
