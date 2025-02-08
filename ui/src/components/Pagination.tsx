import React, { useEffect, useState } from 'react';

interface PaginationProps {
    currentPage: number;
    totalPages: number;
    onPageChange: (page: number) => void;
    from?: number;
    to?: number;
    total?: number;
}

export default function Pagination({
    currentPage,
    totalPages,
    onPageChange,
    from,
    to,
    total,
}: PaginationProps) {
    const [pages, setPages] = useState<(number | 'ellipsis')[]>([]);

    useEffect(() => {
        const getPages = (): (number | 'ellipsis')[] => {
            if (totalPages <= 7) {
                return Array.from({ length: totalPages }, (_, i) => i + 1);
            }

            const pages: (number | 'ellipsis')[] = [];
            if (currentPage <= 4) {
                for (let i = 1; i <= 5; i++) {
                    pages.push(i);
                }
                pages.push('ellipsis');
                pages.push(totalPages);
            } else if (currentPage >= totalPages - 3) {
                pages.push(1);
                pages.push('ellipsis');
                for (let i = totalPages - 4; i <= totalPages; i++) {
                    pages.push(i);
                }
            } else {
                pages.push(1);
                pages.push('ellipsis');
                pages.push(currentPage - 1, currentPage, currentPage + 1);
                pages.push('ellipsis');
                pages.push(totalPages);
            }
            return pages;
        };

        setPages(getPages());
    }, [currentPage, totalPages]);

    const handlePageChange = (page: number, e?: React.MouseEvent) => {
        if (e) e.preventDefault();
        if (page >= 1 && page <= totalPages && page !== currentPage) {
            onPageChange(page);
        }
    };

    return (
        <div className="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6">
            {/* Mobile navigation */}
            <div className="flex flex-1 justify-between sm:hidden">
                <a
                    href="#"
                    onClick={(e) => handlePageChange(currentPage - 1, e)}
                    className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
                >
                    Previous
                </a>
                <a
                    href="#"
                    onClick={(e) => handlePageChange(currentPage + 1, e)}
                    className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50"
                >
                    Next
                </a>
            </div>

            {/* Desktop navigation */}
            <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
                {from !== undefined && to !== undefined && total !== undefined && (
                    <div>
                        <p className="text-sm text-gray-700">
                            Showing <span className="font-medium">{from}</span> to{' '}
                            <span className="font-medium">{to}</span> of{' '}
                            <span className="font-medium">{total}</span> results
                        </p>
                    </div>
                )}
                <div>
                    <nav
                        className="isolate inline-flex -space-x-px rounded-md shadow-xs"
                        aria-label="Pagination"
                    >
                        {/* Previous button */}
                        <a
                            href="#"
                            onClick={(e) => handlePageChange(currentPage - 1, e)}
                            className="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-gray-300 ring-inset hover:bg-gray-50 focus:z-20 focus:outline-offset-0 disabled:opacity-50"
                        >
                            <span className="sr-only">Previous</span>
                            <svg
                                className="h-5 w-5"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                                aria-hidden="true"
                            >
                                <path
                                    fillRule="evenodd"
                                    d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z"
                                    clipRule="evenodd"
                                />
                            </svg>
                        </a>

                        {/* Page number buttons */}
                        {pages.map((page, index) =>
                            page === 'ellipsis' ? (
                                <span
                                    key={`ellipsis-${index}`}
                                    className="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-gray-300 ring-inset"
                                >
                                    ...
                                </span>
                            ) : (
                                <a
                                    key={page}
                                    href="#"
                                    aria-current={page === currentPage ? 'page' : undefined}
                                    onClick={(e) => handlePageChange(page as number, e)}
                                    className={
                                        page === currentPage
                                            ? 'relative z-10 inline-flex items-center bg-indigo-600 px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600'
                                            : 'relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-gray-300 ring-inset hover:bg-gray-50 focus:z-20 focus:outline-offset-0'
                                    }
                                >
                                    {page}
                                </a>
                            )
                        )}

                        {/* Next button */}
                        <a
                            href="#"
                            onClick={(e) => handlePageChange(currentPage + 1, e)}
                            className="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-gray-300 ring-inset hover:bg-gray-50 focus:z-20 focus:outline-offset-0 disabled:opacity-50"
                        >
                            <span className="sr-only">Next</span>
                            <svg
                                className="h-5 w-5"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                                aria-hidden="true"
                            >
                                <path
                                    fillRule="evenodd"
                                    d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z"
                                    clipRule="evenodd"
                                />
                            </svg>
                        </a>
                    </nav>
                </div>
            </div>
        </div>
    );
}
