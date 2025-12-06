export function renderLoadByLevel(level: string): string {
    switch (level) {
        case 'Начальный':
            return 'Лёгкая';
        case 'Средний':
            return 'Средняя';
        case 'PRO':
        case 'Продвинутый':
            return 'Высокая';
        default:
            return 'Неизвестна';
    }
}
