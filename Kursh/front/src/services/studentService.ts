// src/services/studentService.ts
import { api } from './api';
import type { Student } from '../types';

export async function fetchStudents(): Promise<Student[]> {
    return api.get<Student[]>('/students');
}

export async function fetchStudent(id: string): Promise<Student> {
    return api.get<Student>(`/students/${id}`);
}

// alias под старое имя
export const fetchStudentById = fetchStudent;

export async function createStudent(
    fullName: string,
    email: string,
): Promise<Student> {
    return api.post<Student>('/students', { fullName, email });
}

export async function updateStudent(
    id: string,
    fullName: string,
    email: string,
): Promise<Student> {
    return api.put<Student>(`/students/${id}`, { fullName, email });
}
