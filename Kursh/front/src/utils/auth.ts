// src/utils/auth.ts
const AUTH_KEY = 'authenticated';

export function login(username: string, password: string): boolean {
    if (username === 'admin' && password === 'admin') {
        localStorage.setItem(AUTH_KEY, 'true');
        return true;
    }
    return false;
}

export function logout(): void {
    localStorage.removeItem(AUTH_KEY);
}

export function isAuthenticated(): boolean {
    return localStorage.getItem(AUTH_KEY) === 'true';
}
