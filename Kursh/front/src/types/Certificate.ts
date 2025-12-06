export interface Certificate {
    id: string;
    code: string;
    studentId: string;
    studentName: string;
    courseId: string;
    courseTitle: string;
    issuedAt: string;
    grade?: string | null;
}
