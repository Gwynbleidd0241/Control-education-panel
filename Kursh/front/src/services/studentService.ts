import { api } from './api';
import type { Student } from '../types';

export async function fetchStudents(): Promise<Student[]> {
    const res = await api.get<Student[]>('/students');
    return res ?? [];
}

export async function fetchStudentById(id: string): Promise<Student> {
    const res = await api.get<Student>(`/students/${id}`);
    if (!res) {
        throw new Error('Student not found');
    }
    return res;
}

export type UpdateStudentPayload = {
    fullName: string;
    email: string;
    age: number;
    performance: Student['performance'];
    photoUrl: string;
    city: string;
    phone: string;
    format: Student['format'];
    progress: number;
    courseId?: string | null;
};

export async function updateStudent(
    id: string,
    payload: UpdateStudentPayload,
): Promise<Student> {
    const res = await api.put<Student>(`/students/${id}`, payload);
    if (!res) {
        throw new Error('Failed to update student');
    }
    return res;
}
