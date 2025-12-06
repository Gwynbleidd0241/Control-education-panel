import { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { createCourse, updateCourse, fetchCourseById } from '../services/courseService';
import type { Course, CoursePayload } from '../types';

type FormData = {
    title: string;
    description: string;
    full_description: string;
    duration: string;
    price: string;
    level: 'Начальный' | 'Средний' | 'PRO';
    rating: string;
    materialsUrl: string;
    prerequisites: string;
    targetAudience: string;
    instructor: string;
    photoUrl: string;
};

const CourseForm = () => {
    const navigate = useNavigate();
    const { id } = useParams();
    const isEditMode = Boolean(id);

    const [formData, setFormData] = useState<FormData>({
        title: '',
        description: '',
        full_description: '',
        duration: '',
        price: '',
        level: 'Начальный',
        rating: '',
        materialsUrl: '',
        prerequisites: '',
        targetAudience: '',
        instructor: '',
        photoUrl: '',
    });

    const [error, setError] = useState('');
    const [validationErrors, setValidationErrors] = useState({
        price: false,
        duration: false,
        rating: false,
    });

    useEffect(() => {
        if (!isEditMode || !id) {
            return;
        }

        fetchCourseById(id)
            .then((course: Course) => {
                if (!course || !course.id) {
                    console.warn('Курс не найден');
                    navigate('/courses');
                    return;
                }

                setFormData({
                    title: course.title,
                    description: course.description,
                    full_description: course.full_description,
                    duration: course.duration?.toString() ?? '',
                    price: course.price.toString(),
                    level: course.level as FormData['level'],
                    rating: course.rating?.toString() || '',
                    materialsUrl: course.materialsUrl || '',
                    prerequisites: course.prerequisites || '',
                    targetAudience: course.targetAudience || '',
                    instructor: course.instructor || '',
                    photoUrl: course.photoUrl || '',
                });
            })
            .catch((err: unknown) => {
                console.error('Ошибка запроса:', err);
            });
    }, [isEditMode, id, navigate]);

    const handleChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>,
    ) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: value }));
    };

    const handleNumberFieldChange = (name: keyof FormData, value: string) => {
        if (/^\d*(\.\d{0,2})?$/.test(value)) {
            setFormData((prev) => ({ ...prev, [name]: value }));
            setValidationErrors((prev) => ({ ...prev, [name]: false }));
        } else {
            setValidationErrors((prev) => ({ ...prev, [name]: true }));
        }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const { title, description, full_description, duration, price } = formData;
        if (!title || !description || !full_description || !duration || !price) {
            setError('Все поля обязательны для заполнения');
            return;
        }

        if (isNaN(Number(duration)) || isNaN(Number(price))) {
            setError('Цена и продолжительность должны быть числовыми значениями');
            return;
        }

        const ratingNum = Number(formData.rating);
        if (formData.rating && (isNaN(ratingNum) || ratingNum < 0 || ratingNum > 5)) {
            setError('Рейтинг должен быть числом от 0 до 5');
            return;
        }

        if (formData.materialsUrl) {
            try {
                const url = new URL(formData.materialsUrl);
                if (!['http:', 'https:'].includes(url.protocol)) {
                    setError('Ссылка должна начинаться с http:// или https://');
                    return;
                }
            } catch {
                setError('Ссылка на материалы должна быть валидным URL');
                return;
            }
        }

        try {
            const payload: CoursePayload = {
                title: formData.title,
                description: formData.description,
                full_description: formData.full_description,
                duration: Number(formData.duration),
                price: Number(formData.price),
                level: formData.level,
                rating: formData.rating ? Number(formData.rating) : 0,
                materialsUrl: formData.materialsUrl,
                prerequisites: formData.prerequisites,
                targetAudience: formData.targetAudience,
                instructor: formData.instructor,
                photoUrl: formData.photoUrl,
            };

            if (isEditMode && id) {
                await updateCourse(id, payload);
            } else {
                await createCourse(payload);
            }

            navigate('/courses');
        } catch (err) {
            console.error('Ошибка при сохранении курса:', err);
            setError('Ошибка при сохранении курса');
            setTimeout(() => setError(''), 3000);
        }
    };

    return (
        <div className="form-container">
            <h1>{isEditMode ? 'Редактирование курса' : 'Создание нового курса'}</h1>
            <form onSubmit={handleSubmit}>
                <div className="input-group">
                    <label>Название курса</label>
                    <input name="title" value={formData.title} onChange={handleChange} required />
                </div>

                <div className="input-group">
                    <label>Краткое описание</label>
                    <textarea
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="input-group">
                    <label>Полное описание</label>
                    <textarea
                        name="full_description"
                        value={formData.full_description}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className={`input-group ${validationErrors.duration ? 'error' : ''}`}>
                    <label>Продолжительность (в часах)</label>
                    <input
                        name="duration"
                        value={formData.duration}
                        onChange={(e) => handleNumberFieldChange('duration', e.target.value)}
                        required
                    />
                    <span className="hint">Только цифры</span>
                </div>

                <div className={`input-group ${validationErrors.price ? 'error' : ''}`}>
                    <label>Цена (₽)</label>
                    <input
                        name="price"
                        value={formData.price}
                        onChange={(e) => handleNumberFieldChange('price', e.target.value)}
                        required
                    />
                    <span className="hint">Только цифры</span>
                </div>

                <div className="input-group">
                    <label>Уровень сложности</label>
                    <select name="level" value={formData.level} onChange={handleChange}>
                        <option value="Начальный">Начальный</option>
                        <option value="Средний">Средний</option>
                        <option value="PRO">PRO</option>
                    </select>
                </div>

                <div className={`input-group ${validationErrors.rating ? 'error' : ''}`}>
                    <label>Рейтинг</label>
                    <input
                        name="rating"
                        value={formData.rating}
                        onChange={(e) => handleNumberFieldChange('rating', e.target.value)}
                    />
                    <span className="hint">Число от 0 до 5</span>
                </div>

                <div className="input-group">
                    <label>Материалы курса (ссылка)</label>
                    <input
                        name="materialsUrl"
                        value={formData.materialsUrl}
                        onChange={handleChange}
                    />
                </div>

                <div className="input-group">
                    <label>Начальные требования</label>
                    <textarea
                        name="prerequisites"
                        value={formData.prerequisites}
                        onChange={handleChange}
                    />
                </div>

                <div className="input-group">
                    <label>Для кого курс</label>
                    <textarea
                        name="targetAudience"
                        value={formData.targetAudience}
                        onChange={handleChange}
                    />
                </div>

                <div className="input-group">
                    <label>Преподаватель</label>
                    <input
                        name="instructor"
                        value={formData.instructor}
                        onChange={handleChange}
                    />
                </div>

                <div className="input-group">
                    <label>Фото курса (URL)</label>
                    <input
                        name="photoUrl"
                        value={formData.photoUrl}
                        onChange={handleChange}
                    />
                </div>

                <div className="form-actions">
                    <button type="submit" className="primary-btn">
                        {isEditMode ? 'Сохранить изменения' : 'Создать курс'}
                    </button>
                    <button
                        type="button"
                        className="cancel-btn"
                        onClick={() => navigate('/courses')}
                    >
                        Отмена
                    </button>
                </div>

                {error && <div className="error-notification">{error}</div>}
            </form>
        </div>
    );
};

export default CourseForm;
