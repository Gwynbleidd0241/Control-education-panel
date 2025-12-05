// src/services/api.ts

const API_BASE =
    process.env.REACT_APP_API_BASE || 'http://localhost:8081/api';

async function request<T>(
    path: string,
    options: RequestInit = {},
): Promise<T> {
    const res = await fetch(`${API_BASE}${path}`, {
        headers: {
            'Content-Type': 'application/json',
            ...(options.headers || {}),
        },
        credentials: 'include',
        ...options,
    });

    if (!res.ok) {
        const text = await res.text();
        throw new Error(`API error ${res.status}: ${text || res.statusText}`);
    }

    if (res.status === 204) {
        // No Content
        // @ts-expect-error
        return null;
    }

    return (await res.json()) as T;
}

export const api = {
    get:  <T>(path: string) => request<T>(path),
    post: <T>(path: string, body: unknown) =>
        request<T>(path, {
            method: 'POST',
            body: JSON.stringify(body),
        }),
    put:  <T>(path: string, body: unknown) =>
        request<T>(path, {
            method: 'PUT',
            body: JSON.stringify(body),
        }),
    del:  (path: string) =>
        request<null>(path, {
            method: 'DELETE',
        }),
};
