import React, { useEffect, useMemo, useState } from 'react';
import { Link } from 'react-router-dom';
import Loader from '../components/Loader';
import { Certificate } from '../types';
import { fetchCertificates, issueCertificate } from '../services/certificateService';
import { fetchStudents } from '../services/studentService';
import { fetchCourses } from '../services/courseService';
import { Student, Course } from '../types';
import '../styles/certificates.css';

const formatDate = (iso?: string | null) =>
    iso ? new Date(iso).toLocaleDateString('ru-RU') : '—';

const Certificates: React.FC = () => {
    const [certificates, setCertificates] = useState<Certificate[]>([]);
    const [students, setStudents] = useState<Student[]>([]);
    const [courses, setCourses] = useState<Course[]>([]);
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState('');
    const [showIssueForm, setShowIssueForm] = useState(false);
    const [issueForm, setIssueForm] = useState({
        studentId: '',
        courseId: '',
        progress: '',
    });
    const [submitting, setSubmitting] = useState(false);

    useEffect(() => {
        const load = async () => {
            try {
                const [certsData, studentsData, coursesData] = await Promise.all([
                    fetchCertificates().catch(() => []),
                    fetchStudents().catch(() => []),
                    fetchCourses().catch(() => []),
                ]);
                setCertificates(certsData || []);
                setStudents(studentsData || []);
                setCourses(coursesData || []);
            } catch (e) {
                console.error('Ошибка загрузки данных:', e);
            } finally {
                setLoading(false);
            }
        };

        load();
    }, []);

    const handleIssue = async () => {
        if (!issueForm.studentId || !issueForm.courseId) {
            alert('Выберите студента и курс');
            return;
        }

        try {
            setSubmitting(true);
            const student = students.find(s => s.id === issueForm.studentId);
            const course = courses.find(c => c.id === issueForm.courseId);
            
            if (!student || !course) {
                alert('Студент или курс не найдены');
                return;
            }

            await issueCertificate({
                studentId: issueForm.studentId,
                courseId: issueForm.courseId,
                grade: issueForm.progress ? `${issueForm.progress}%` : '',
            });

            const data = await fetchCertificates();
            setCertificates(data || []);
            setShowIssueForm(false);
            setIssueForm({ studentId: '', courseId: '', progress: '' });
            alert('Сертификат успешно выдан!');
        } catch (e) {
            console.error('Ошибка выдачи сертификата:', e);
            alert('Не удалось выдать сертификат');
        } finally {
            setSubmitting(false);
        }
    };

    const filtered = useMemo(
        () =>
            certificates.filter((cert) => {
                if (!search.trim()) return true;
                const q = search.toLowerCase();
                const student = students.find(s => s.id === cert.studentId);
                const course = courses.find(c => c.id === cert.courseId);
                return (
                    student?.fullName.toLowerCase().includes(q) ||
                    course?.title.toLowerCase().includes(q) ||
                    cert.code?.toLowerCase().includes(q)
                );
            }),
        [certificates, search, students, courses],
    );

    if (loading) {
        return <Loader />;
    }

    return (
        <main className="certificates-page">
            <header className="page-header">
                <h1>Сертификаты</h1>
                <div className="actions-group">
                    <Link to="/" className="btn secondary">
                        ← В дашборд
                    </Link>
                    <button
                        type="button"
                        className="btn primary"
                        onClick={() => setShowIssueForm(!showIssueForm)}
                    >
                        {showIssueForm ? 'Отмена' : '+ Выдать сертификат'}
                    </button>
                </div>
            </header>

            {showIssueForm && (
                <div className="issue-card">
                    <h3>Выдача сертификата</h3>
                    <div className="issue-grid">
                        <label>
                            Студент
                            <select
                                value={issueForm.studentId}
                                onChange={(e) => setIssueForm({ ...issueForm, studentId: e.target.value })}
                            >
                                <option value="">Выберите студента</option>
                                {students.map((s) => (
                                    <option key={s.id} value={s.id}>
                                        {s.fullName}
                                    </option>
                                ))}
                            </select>
                        </label>
                        <label>
                            Курс
                            <select
                                value={issueForm.courseId}
                                onChange={(e) => setIssueForm({ ...issueForm, courseId: e.target.value })}
                            >
                                <option value="">Выберите курс</option>
                                {courses.map((c) => (
                                    <option key={c.id} value={c.id}>
                                        {c.title}
                                    </option>
                                ))}
                            </select>
                        </label>
                        <label>
                            Прогресс, %
                            <input
                                type="number"
                                min={0}
                                max={100}
                                value={issueForm.progress}
                                onChange={(e) => setIssueForm({ ...issueForm, progress: e.target.value })}
                                placeholder="0–100"
                            />
                        </label>
                    </div>
                    <div className="issue-actions">
                        <button
                            type="button"
                            className="btn primary"
                            onClick={handleIssue}
                            disabled={submitting}
                        >
                            {submitting ? 'Выдача…' : 'Выдать сертификат'}
                        </button>
                    </div>
                </div>
            )}

            <section className="search-section" style={{ marginBottom: '1.5rem' }}>
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Поиск по студенту, курсу или коду…"
                        value={search}
                        onChange={(e) => setSearch(e.target.value)}
                    />
                </section>

            <section className="table-wrapper">
                    {filtered.length === 0 ? (
                        <p className="empty-state" style={{ textAlign: 'center', padding: '2rem', color: 'var(--gray)' }}>
                            {certificates.length === 0 
                                ? 'Сертификатов пока нет. Выдайте первый сертификат!'
                                : 'По текущему поиску сертификатов не найдено.'}
                        </p>
                    ) : (
                        <table className="certificates-table">
                            <thead>
                            <tr>
                                <th>Код</th>
                                <th>Студент</th>
                                <th>Курс</th>
                                <th>Прогресс</th>
                                <th>Выдан</th>
                                <th>Действия</th>
                            </tr>
                            </thead>
                            <tbody>
                            {filtered.map((cert) => {
                                const student = students.find(s => s.id === cert.studentId);
                                const course = courses.find(c => c.id === cert.courseId);
                                return (
                                <tr key={cert.id}>
                                    <td>{cert.code}</td>
                                    <td>{student?.fullName || cert.studentId}</td>
                                    <td>{course?.title || cert.courseId}</td>
                                    <td>
                                        {cert.grade
                                            ? cert.grade
                                            : typeof student?.progress === 'number'
                                                ? `${student.progress}%`
                                                : '—'}
                                    </td>
                                    <td>{formatDate(cert.issuedAt)}</td>
                                    <td>
                                        <div className="table-actions">
                                            <Link to={`/certificates/${cert.id}`} className="btn secondary">
                                                Просмотр
                                            </Link>
                                        </div>
                                    </td>
                                </tr>
                                );
                            })}
                            </tbody>
                        </table>
                    )}
                </section>
            </main>
    );
};

export default Certificates;
