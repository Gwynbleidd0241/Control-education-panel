import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { fetchStudentById } from '../services/studentService';
import { Student } from '../types';
import Loader from '../components/Loader';
import '../styles/studentDetails.css';

const StudentDetails = () => {
    const { id } = useParams<{ id: string }>();
    const [student, setStudent] = useState<Student | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const loadStudent = async () => {
            try {
                if (id) {
                    const data = await fetchStudentById(id);
                    setStudent(data);
                }
            } catch (error) {
                setError('Не удалось загрузить данные студента');
            } finally {
                setLoading(false);
            }
        };
        loadStudent();
    }, [id]);

    if (loading) return <Loader />;
    if (error) return <div className="error-message">{error}</div>;
    if (!student) return <div className="error-message">Студент не найден</div>;

    return (
        <div className="student-details-wrapper">
            <div className="student-card">
                <div className="student-header">
                    <h1>{student.fullName}</h1>
                    <Link to="/students" className="btn secondary">
                        <span>Назад к списку</span>
                    </Link>
                </div>

                {student.photoUrl && (
                    <div className="student-photo">
                        <img src={student.photoUrl} alt={student.fullName} />
                    </div>
                )}

                <div className="student-section">
                    <h3>Контактная информация</h3>
                    <ul>
                        <li><strong>Email:</strong> {student.email}</li>
                        <li><strong>Телефон:</strong> {student.phone}</li>
                        <li><strong>Город:</strong> {student.city}</li>
                        <li><strong>Возраст:</strong> {student.age}</li>
                    </ul>
                </div>

                <div className="student-section">
                    <h3>Обучение</h3>
                    <ul>
                        <li><strong>Формат:</strong> {student.format}</li>
                        <li><strong>Успеваемость:</strong> {student.performance}</li>
                        <li><strong>Прогресс прохождения курса:</strong> {student.progress}%</li>
                    </ul>
                </div>

                <div className="student-section">
                    <h3>Курс</h3>
                    <ul>
                        <li><strong>Название:</strong> {student.course.title}</li>
                        <li><strong>Уровень:</strong> {student.course.level}</li>
                        <li><strong>Стоимость:</strong> {student.course.price}</li>
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default StudentDetails;
