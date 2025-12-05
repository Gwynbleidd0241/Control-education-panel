// src/pages/Courses.tsx
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { fetchCourses, deleteCourse } from '../services/courseService';
import { Course } from '../types';
import Loader from '../components/Loader';
import CourseCard from '../components/CourseCard';
import Notification from '../components/Notification';
import { handleApiError, formatPrice } from '../utils/helpers';

const Courses: React.FC = () => {
    const [courses, setCourses] = useState<Course[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const loadCourses = async () => {
            try {
                const data = await fetchCourses();
                setCourses(data);
            } catch (err) {
                setError(handleApiError(err));
            } finally {
                setLoading(false);
            }
        };

        loadCourses();
    }, []);

    const handleDelete = async (id: string) => {
        if (!window.confirm('Точно удалить этот курс?')) return;

        try {
            await deleteCourse(id);
            setCourses((prev) => prev.filter((c) => c.id !== id));
        } catch (err) {
            setError(handleApiError(err));
        }
    };

    if (loading) return <Loader />;

    return (
        <div className="page-container">
            <div className="page-header">
                <h1>Курсы</h1>

                <Link to="/" className="btn secondary">
                    ← В дашборд
                </Link>

                <Link to="/courses/new" className="btn primary">
                    + Новый курс
                </Link>
            </div>

            {error && (
                <Notification
                    type="error"
                    message={error}
                    onClose={() => setError('')}
                />
            )}

            {courses.length === 0 ? (
                <p>Пока нет ни одного курса.</p>
            ) : (
                <div className="courses-grid">
                    {courses.map((course) => (
                        <CourseCard
                            key={course.id}
                            course={course}
                            formattedPrice={formatPrice(course.price)}
                            onDelete={() => handleDelete(course.id)}
                        />
                    ))}
                </div>
            )}
        </div>
    );
};

export default Courses;
