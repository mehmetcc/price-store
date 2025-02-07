import Link from 'next/link';

export default function Navbar() {
    return (
        <nav className="bg-gray-800 h-full p-4">
            <ul className="space-y-4">
                <li>
                    <Link href="/symbols" className="text-white hover:underline">
                        Symbols
                    </Link>
                </li>
                <li>
                    <Link href="/prices" className="text-white hover:underline">
                        Past Prices
                    </Link>
                </li>
            </ul>
        </nav>
    );
}
