// src/services/certificateService.ts
import { api } from './api';
import type { Certificate, CertificateTemplate } from '../types';

// все сертификаты
export async function fetchCertificates(): Promise<Certificate[]> {
    return api.get<Certificate[]>('/certificates');
}

// все шаблоны сертификатов
export async function fetchCertificateTemplates(): Promise<CertificateTemplate[]> {
    return api.get<CertificateTemplate[]>('/certificate-templates');
}

// проверка сертификата по коду
export async function verifyCertificateByCode(code: string): Promise<Certificate | null> {
    const all = await fetchCertificates();
    const found = all.find((c) => c.code.toLowerCase() === code.toLowerCase());
    return found ?? null;
}

// выдача сертификата
export async function issueCertificate(
    payload: Omit<Certificate, 'id'>,
): Promise<Certificate> {
    return api.post<Certificate>('/certificates', payload);
}

// отзыв сертификата
export async function revokeCertificate(id: string): Promise<Certificate> {
    return api.put<Certificate>(`/certificates/${id}`, { status: 'REVOKED' });
}
