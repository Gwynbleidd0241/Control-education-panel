import { useLocation, useParams, Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { fetchCourseById } from '../services/courseService';
import { Course } from '../types';
import Loader from '../components/Loader';
import { formatPrice } from '../utils/helpers';

const CourseDetails = () => {
    const { id } = useParams();
    const [course, setCourse] = useState<Course | null>(null);
    const [loading, setLoading] = useState(true);

    const location = useLocation();
    const from = location.state?.from || '/courses';

    useEffect(() => {
        const loadData = async () => {
            try {
                if (id) {
                    const data = await fetchCourseById(id);
                    setCourse(data);
                }
            } catch (error) {
                console.error('Ошибка загрузки данных курса:', error);
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, [id]);

    if (loading) return <Loader />;
    if (!course) return <div className="error-message">Курс был удален. Обучение остановлено.</div>;

    return (
        <div className="course-details-wrapper">
            <div className="course-card">
                <div className="course-header">
                    <h1>{course.title}</h1>
                    <div className="course-buttons mobile-priority">
                        <Link to={`/courses/${course.id}/edit`} className="btn primary">
                            ✎ Редактировать
                        </Link>
                        <Link to={from} className="btn secondary">
                            ← Назад
                        </Link>
                    </div>
                </div>



                <div className="course-section">
                <h3>Основная информация</h3>
                <ul className="course-info-list">
                    <li><strong>Уровень:</strong> {course.level}</li>
                    <li><strong>Стоимость:</strong> {formatPrice(course.price)}</li>
                    <li><strong>Продолжительность:</strong> {course.duration} ч</li>
                    <li><strong>Рейтинг:</strong> {course.rating} ★</li>
                </ul>
            </div>

            <div className="course-section">
                <h3>Описание</h3>
                <p>{course.full_description}</p>
            </div>

            <div className="course-section">
                <h3>Для кого курс</h3>
                <p>{course.targetAudience}</p>
            </div>

            <div className="course-section">
                <h3>Начальные требования</h3>
                <p>{course.prerequisites}</p>
            </div>

            <div className="course-section">
                <h3>Преподаватель</h3>
                <p>{course.instructor}</p>
            </div>

            <div className="course-section">
                <h3>Материалы курса</h3>
                <a href={course.materialsUrl} target="_blank" rel="noopener noreferrer">
                    Перейти к материалам
                </a>
            </div>

        </div>
        </div>
    );
};

export default CourseDetails;
