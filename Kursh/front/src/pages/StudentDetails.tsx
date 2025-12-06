import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { fetchStudentById, updateStudent, UpdateStudentPayload } from '../services/studentService';
import { fetchCourseById } from '../services/courseService';
import { Student, Course } from '../types';
import Loader from '../components/Loader';
import '../styles/studentDetails.css';

const StudentDetails = () => {
    const { id } = useParams<{ id: string }>();
    const [student, setStudent] = useState<Student | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [showModal, setShowModal] = useState(false);
    const [saving, setSaving] = useState(false);
    const [form, setForm] = useState<UpdateStudentPayload | null>(null);

    useEffect(() => {
        const loadStudent = async () => {
            try {
                if (!id) return;

                let data = await fetchStudentById(id);

                if (!data.course && data.courseId) {
                    try {
                        const course: Course = await fetchCourseById(data.courseId);
                        data = { ...data, course };
                    } catch {
                        // оставляем без курса, если не нашли
                    }
                }

                setStudent(data);
                setForm({
                    fullName: data.fullName,
                    email: data.email,
                    age: data.age,
                    performance: data.performance,
                    photoUrl: data.photoUrl,
                    city: data.city,
                    phone: data.phone,
                    format: data.format,
                    progress: data.progress,
                    courseId: data.courseId ?? null,
                });
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

    const handleChange = (key: keyof UpdateStudentPayload, value: string | number | null) => {
        if (!form) return;
        setForm({ ...form, [key]: value });
    };

    const handleSave = async () => {
        if (!id || !form) return;
        setSaving(true);
        try {
            const updated = await updateStudent(id, form);
            // сохраняем курс, если он был подтянут ранее
            setStudent({ ...updated, course: student.course });
            setShowModal(false);
        } catch {
            setError('Не удалось сохранить изменения');
        } finally {
            setSaving(false);
        }
    };

    return (
        <div className="student-details-wrapper">
            <div className="student-card">
                <div className="student-header">
                    <h1>{student.fullName}</h1>
                    <div className="header-actions">
                    <Link to="/students" className="btn secondary">
                        <span>Назад к списку</span>
                    </Link>
                        <button
                            className="btn primary"
                            type="button"
                            onClick={() => setShowModal(true)}
                        >
                            Редактировать
                        </button>
                    </div>
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
                    {!student.course ? (
                        <p>Курс не назначен</p>
                    ) : (
                        <ul>
                            <li><strong>Название:</strong> {student.course.title}</li>
                            <li><strong>Уровень:</strong> {student.course.level}</li>
                            <li><strong>Стоимость:</strong> {student.course.price}</li>
                        </ul>
                    )}
                </div>
            </div>

            {showModal && form && (
                <div className="modal-overlay" onClick={() => setShowModal(false)}>
                    <div className="modal" onClick={(e) => e.stopPropagation()}>
                        <div className="modal-header">
                            <h3>Редактировать студента</h3>
                            <button className="close-btn" onClick={() => setShowModal(false)}>×</button>
                        </div>
                        <div className="modal-body">
                            <label>
                                ФИО
                                <input
                                    value={form.fullName}
                                    onChange={(e) => handleChange('fullName', e.target.value)}
                                />
                            </label>
                            <label>
                                Email
                                <input
                                    type="email"
                                    value={form.email}
                                    onChange={(e) => handleChange('email', e.target.value)}
                                />
                            </label>
                            <label>
                                Телефон
                                <input
                                    value={form.phone}
                                    onChange={(e) => handleChange('phone', e.target.value)}
                                />
                            </label>
                            <label>
                                Город
                                <input
                                    value={form.city}
                                    onChange={(e) => handleChange('city', e.target.value)}
                                />
                            </label>
                            <div className="modal-row">
                                <label>
                                    Возраст
                                    <input
                                        type="number"
                                        value={form.age}
                                        onChange={(e) => handleChange('age', Number(e.target.value))}
                                    />
                                </label>
                                <label>
                                    Прогресс (%)
                                    <input
                                        type="number"
                                        min={0}
                                        max={100}
                                        value={form.progress}
                                        onChange={(e) => handleChange('progress', Number(e.target.value))}
                                    />
                                </label>
                            </div>
                            <div className="modal-row">
                                <label>
                                    Формат
                                    <select
                                        value={form.format}
                                        onChange={(e) => handleChange('format', e.target.value)}
                                    >
                                        <option value="Онлайн">Онлайн</option>
                                        <option value="Очно">Очно</option>
                                    </select>
                                </label>
                                <label>
                                    Успеваемость
                                    <select
                                        value={form.performance}
                                        onChange={(e) => handleChange('performance', e.target.value)}
                                    >
                                        <option value="учиться легко">учиться легко</option>
                                        <option value="учиться средне">учиться средне</option>
                                        <option value="учиться тяжело">учиться тяжело</option>
                                    </select>
                                </label>
                            </div>
                        </div>
                        <div className="modal-actions">
                            <button
                                type="button"
                                className="btn secondary"
                                onClick={() => setShowModal(false)}
                                disabled={saving}
                            >
                                Отмена
                            </button>
                            <button
                                type="button"
                                className="btn primary"
                                onClick={handleSave}
                                disabled={saving}
                            >
                                {saving ? 'Сохранение…' : 'Сохранить'}
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default StudentDetails;
