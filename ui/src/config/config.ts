const config = {
    baseUrl: process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:9991',
    pageSize: Number(process.env.NEXT_PUBLIC_PAGE_SIZE) || 20,
};

export default config;
