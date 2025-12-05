// src/types/Certificate.ts

// статус выданного сертификата
export type CertificateStatus = 'valid' | 'expired' | 'revoked';

// шаблон сертификата для курса
export interface CertificateTemplate {
    id: string;
    name: string;         // Название шаблона
    courseId: string;     // ID курса
    courseTitle?: string; // Человеческое название курса
    validityDays?: number; // Срок действия в днях (если есть)
}

// выданный сертификат
export interface Certificate {
    id: string;           // Внутренний id записи
    code: string;         // Код/номер сертификата (то, что вводит пользователь)
    studentId: string;
    studentName: string;
    courseId: string;
    courseTitle: string;
    issuedAt: string;     // '2025-02-01'
    expiresAt: string | null;
    status: CertificateStatus;
    grade?: string | null;
}
