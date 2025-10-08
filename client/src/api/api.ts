const API_BASE_URL = 'http://localhost:8081/api';

interface ApiResult<T> {
    data: T | null;
    error: string | null;
}

export const apiCall = async <T>(
    endpoint: string,
    method: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH',
    body: object | null = null,
    token: string | null = null
): Promise<T> => {
    const headers: HeadersInit = { 'Content-Type': 'application/json' };

    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    }

    const config: RequestInit = {
        method,
        headers,
        body: body ? JSON.stringify(body) : null,
    };

    try {
        const url = `${API_BASE_URL}${endpoint}`;
        const response = await fetch(url, config);

        if (!response.ok) {
            let errorDetails = response.statusText;

            try {
                const errorData = await response.json();
                errorDetails = errorData.message || response.statusText;
            } catch {}

            throw new Error(`Ошибка ${response.status}: ${errorDetails}`);
        }

        if (response.status === 201 || response.status === 204) {
            return {} as T;
        }

        return await response.json() as T;

    } catch (err) {
        throw new Error(err instanceof Error ? err.message : "Неизвестная ошибка API");
    }
};