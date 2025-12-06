import { api } from './api';
import type { Course, CoursePayload } from '../types';

export async function fetchCourses(): Promise<Course[]> {
    const res = await api.get<Course[]>('/courses');
    return res ?? [];
}

export async function fetchCourseById(id: string): Promise<Course> {
    const res = await api.get<Course>(`/courses/${id}`);
    if (!res) {
        throw new Error('Course not found');
    }
    return res;
}

export async function createCourse(payload: CoursePayload): Promise<Course> {
    const res = await api.post<Course>('/courses', payload);
    if (!res) {
        throw new Error('Failed to create course');
    }
    return res;
}

export async function updateCourse(
    id: string,
    payload: Partial<CoursePayload>,
): Promise<Course> {
    const res = await api.put<Course>(`/courses/${id}`, payload);
    if (!res) {
        throw new Error('Failed to update course');
    }
    return res;
}

export async function deleteCourse(id: string): Promise<void> {
    await api.del(`/courses/${id}`);
}
