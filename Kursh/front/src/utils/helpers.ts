// src/utils/helpers.ts
import type { ApiError } from '../types';

/**
 * Форматируем цену в рублях: 1200 -> "1 200 ₽"
 */
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

/**
 * Приводим ошибку из fetch/axios/бэка к человекочитаемой строке
 */
export function handleApiError(err: unknown): string {
    if (!err) return 'Неизвестная ошибка';

    // если наш типизированный ApiError
    if (typeof err === 'object' && err !== null && 'message' in err) {
        const apiErr = err as ApiError;
        return apiErr.message || 'Ошибка при запросе к серверу';
    }

    // стандартный Error
    if (err instanceof Error) {
        return err.message || 'Ошибка при запросе к серверу';
    }

    // что-то странное
    return 'Ошибка при запросе к серверу';
}
