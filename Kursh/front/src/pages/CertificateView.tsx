import React, { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import Loader from '../components/Loader';
import { Certificate, Course, Student } from '../types';
import { fetchCertificateById } from '../services/certificateService';
import { fetchCourseById } from '../services/courseService';
import { fetchStudentById } from '../services/studentService';
import '../styles/certificates.css';

const formatDate = (iso?: string | null) =>
    iso ? new Date(iso).toLocaleDateString('ru-RU') : '—';

const CertificateView: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [certificate, setCertificate] = useState<Certificate | null>(null);
    const [student, setStudent] = useState<Student | null>(null);
    const [course, setCourse] = useState<Course | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const load = async () => {
            if (!id) {
                setError('Сертификат не найден');
                setLoading(false);
                return;
            }

            try {
                const cert = await fetchCertificateById(id);
                if (!cert) {
                    setError('Сертификат не найден');
                    return;
                }

                setCertificate(cert);

                const [studentData, courseData] = await Promise.all([
                    fetchStudentById(cert.studentId).catch(() => null),
                    fetchCourseById(cert.courseId).catch(() => null),
                ]);

                if (studentData) setStudent(studentData);
                if (courseData) setCourse(courseData);
            } catch (e) {
                console.error('Не удалось загрузить сертификат', e);
                setError('Не удалось загрузить сертификат');
            } finally {
                setLoading(false);
            }
        };

        load();
    }, [id]);

    if (loading) return <Loader />;
    if (error) return <div className="error-message">{error}</div>;
    if (!certificate) return <div className="error-message">Сертификат не найден</div>;

    return (
        <main className="certificate-view">
            <div className="actions-group" style={{ justifyContent: 'space-between', marginBottom: '1.5rem' }}>
                <Link to="/certificates" className="btn secondary">
                    ← Назад к списку
                </Link>
                <span className="certificate-code">Код: {certificate.code}</span>
            </div>

            <section className="certificate-hero">
                <div className="certificate-body">
                    <div className="certificate-headline">
                        <div>
                            <p style={{ color: 'rgba(0,0,0,0.6)', marginBottom: '0.35rem' }}>Сертификат о прохождении курса</p>
                            <h1>{course?.title || certificate.courseTitle || 'Курс'}</h1>
                        </div>
                        <div className="certificate-meta">
                            <span>Дата выдачи: {formatDate(certificate.issuedAt)}</span>
                        </div>
                    </div>

                    <div className="certificate-grid">
                        <div className="certificate-tile">
                            <strong>Студент</strong>
                            <span>{student?.fullName || certificate.studentName || certificate.studentId}</span>
                        </div>
                        <div className="certificate-tile">
                            <strong>Курс</strong>
                            <span>{course?.title || certificate.courseTitle || certificate.courseId}</span>
                        </div>
                        <div className="certificate-tile">
                            <strong>Прогресс</strong>
                            <span>
                                {certificate.grade
                                    ? certificate.grade
                                    : typeof student?.progress === 'number'
                                        ? `${student.progress}%`
                                        : '—'}
                            </span>
                        </div>
                        <div className="certificate-tile">
                            <strong>Инструктор</strong>
                            <span>{course?.instructor || '—'}</span>
                        </div>
                    </div>

                        <div className="certificate-footer">
                            <div className="certificate-code">Номер: {certificate.code}</div>
                        </div>
                </div>
            </section>
        </main>
    );
};

export default CertificateView;

