// src/services/courseService.ts
import { api } from './api';
import type { Course, CoursePayload } from '../types';

export async function fetchCourses(): Promise<Course[]> {
    return api.get<Course[]>('/courses');
}

export async function fetchCourse(id: string): Promise<Course> {
    return api.get<Course>(`/courses/${id}`);
}

// alias под старое имя, которое уже используется в компонентах
export const fetchCourseById = fetchCourse;

export async function createCourse(payload: CoursePayload): Promise<Course> {
    return api.post<Course>('/courses', payload);
}

export async function updateCourse(
    id: string,
    payload: Partial<CoursePayload>,
): Promise<Course> {
    return api.put<Course>(`/courses/${id}`, payload);
}

export async function deleteCourse(id: string): Promise<void> {
    await api.del(`/courses/${id}`);
}
