import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { fetchStudents } from '../services/studentService';
import { fetchCourses } from '../services/courseService';
import { Student, Course } from '../types';
import Loader from '../components/Loader';
import StudentTableRow from '../components/StudentTableRow';
import { renderLoadByLevel } from '../utils/performanceHelpers';

const useIsMobile = (query = '(max-width: 767px)') => {
    const [isMobile, setIsMobile] = useState(false);

    useEffect(() => {
        if (typeof window === 'undefined' || !window.matchMedia) return;

        const mediaQuery = window.matchMedia(query);
        const onChange = (e: MediaQueryListEvent | MediaQueryList) => {
            setIsMobile('matches' in e ? e.matches : (e as MediaQueryList).matches);
        };

        onChange(mediaQuery);
        mediaQuery.addEventListener?.('change', onChange);
        mediaQuery.addListener?.(onChange);

        return () => {
            mediaQuery.removeEventListener?.('change', onChange);
            mediaQuery.removeListener?.(onChange);
        };
    }, [query]);

    return isMobile;
};

const Students = () => {
    const [students, setStudents] = useState<Student[]>([]);
    const [courses, setCourses] = useState<Record<string, Course>>({});
    const [loading, setLoading] = useState(true);

    const isMobile = useIsMobile();

    useEffect(() => {
        const loadData = async () => {
            try {
                const [studentsData, coursesData] = await Promise.all([
                    fetchStudents(),
                    fetchCourses().catch(() => []),
                ]);

                const coursesSafe: Course[] = Array.isArray(coursesData) ? coursesData : [];
                const coursesMap = coursesSafe.reduce<Record<string, Course>>((acc, course) => {
                    acc[course.id] = course;
                    return acc;
                }, {} as Record<string, Course>);

                const enriched = studentsData.map((student) => {
                    const course =
                        student.course ??
                        (student.courseId ? coursesMap[student.courseId] : null);

                    return { ...student, course };
                });

                setCourses(coursesMap);
                setStudents(enriched);
            } catch (error) {
                console.error('Ошибка загрузки студентов:', error);
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, []);

    if (loading) return <Loader />;

    return (
        <div className="page-container">
            <div className="page-header">
                <h1 className="page-title">Список студентов</h1>
                <Link to="/" className="back-button"> {isMobile ? '←' : '← Назад'}</Link>
            </div>

            {!isMobile && (
                <div className="students-table-wrapper">
                    <table className="students-table">
                        <thead>
                        <tr>
                            <th>ФИО</th>
                            <th>Email</th>
                            <th>Курс</th>
                            <th>Нагрузка</th>
                            <th>Подробнее</th>
                        </tr>
                        </thead>
                        <tbody>
                        {students.map(student => (
                            <StudentTableRow key={student.id} student={student} />
                        ))}
                        </tbody>
                    </table>
                </div>
            )}

            {isMobile && (
                <div className="students-card-list">
                    {students.map(student => {
                        const course = student.course ?? (student.courseId ? courses[student.courseId] : null);
                        return (
                        <div key={student.id} className="student-card">
                            <h3>{student.fullName}</h3>
                            <p><strong>Email:</strong> {student.email}</p>
                            <p>
                                <strong>Курс:</strong>{' '}
                                {course?.title ?? 'Без курса'}
                            </p>
                            <p>
                                <strong>Нагрузка:</strong>{' '}
                                {course
                                    ? renderLoadByLevel(course.level)
                                    : '—'}
                            </p>
                            <Link to={`/students/${student.id}`}>Подробнее</Link>
                        </div>
                    );})}
                </div>
            )}
        </div>
    );
};

export default Students;
