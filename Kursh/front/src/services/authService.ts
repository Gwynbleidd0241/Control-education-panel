// src/services/authService.ts

const AUTH_KEY = 'authenticated';

export const login = async (username: string, password: string): Promise<boolean> => {
    // здесь можно будет когда-нибудь прикрутить настоящий API,
    // сейчас просто заглушка "admin / admin"
    const ok = username === 'admin' && password === 'admin';

    if (ok) {
        localStorage.setItem(AUTH_KEY, 'true');
    }

    return ok;
};

export const logout = (): void => {
    localStorage.removeItem(AUTH_KEY);
};

export const checkAuth = (): boolean => {
    return localStorage.getItem(AUTH_KEY) === 'true';
};
