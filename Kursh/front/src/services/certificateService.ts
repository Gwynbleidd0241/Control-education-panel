import { api } from './api';
import type { Certificate } from '../types';

export async function fetchCertificates(): Promise<Certificate[]> {
    const res = await api.get<Certificate[]>('/certificates');
    return res ?? [];
}

export async function fetchCertificateById(id: string): Promise<Certificate | null> {
    try {
        return await api.get<Certificate>(`/certificates/${id}`);
    } catch (err) {
        try {
            const list = await fetchCertificates();
            return list.find((c) => c.id === id) ?? null;
        } catch {
            return null;
        }
    }
}

export async function verifyCertificateByCode(code: string): Promise<Certificate | null> {
    const all = await fetchCertificates();
    const found = all.find((c) => c.code.toLowerCase() === code.toLowerCase());
    return found ?? null;
}

export interface IssueCertificatePayload {
    studentId: string;
    courseId: string;
    grade?: string;
}

export async function issueCertificate(
    payload: IssueCertificatePayload,
): Promise<Certificate> {
    const res = await api.post<Certificate>('/certificates/issue', payload);
    if (!res) {
        throw new Error('Failed to issue certificate');
    }
    return res;
}