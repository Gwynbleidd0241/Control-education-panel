import { Course } from './Course';

export interface Student {
    id: string;
    fullName: string;
    email: string;
    age: number;
    course?: Course | null;
    performance: 'учиться легко' | 'учиться средне' | 'учиться тяжело';
    photoUrl: string;
    city: string;
    phone: string;
    format: 'Онлайн' | 'Очно';
    progress: number;
}

export type StudentFormData = Omit<Student, 'id' | 'course'> & {
    courseId: string;
};
