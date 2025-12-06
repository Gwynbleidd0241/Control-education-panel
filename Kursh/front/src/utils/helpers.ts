import type { ApiError } from '../types';

export function formatPrice(value: number | string): string {
    const num = typeof value === 'string' ? Number(value) : value;

    if (Number.isNaN(num)) {
        return '—';
    }

    return new Intl.NumberFormat('ru-RU', {
        style: 'currency',
        currency: 'RUB',
        maximumFractionDigits: 0,
    }).format(num);
}

export function handleApiError(err: unknown): string {
    if (!err) return 'Неизвестная ошибка';

    if (typeof err === 'object' && err !== null && 'message' in err) {
        const apiErr = err as ApiError;
        return apiErr.message || 'Ошибка при запросе к серверу';
    }

    if (err instanceof Error) {
        return err.message || 'Ошибка при запросе к серверу';
    }

    return 'Ошибка при запросе к серверу';
}
