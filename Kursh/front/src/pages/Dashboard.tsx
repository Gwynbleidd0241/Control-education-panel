import React, { useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchCourses } from '../services/courseService';
import { fetchStudents } from '../services/studentService';
import { fetchCertificates } from '../services/certificateService';
import Loader from '../components/Loader';
import StatsCard from '../components/StatsCard';
import type { Course, Student, Certificate } from '../types';
import { logout } from '../utils/auth';

const Dashboard: React.FC = () => {
    const navigate = useNavigate();
    const [courses, setCourses] = useState<Course[]>([]);
    const [students, setStudents] = useState<Student[]>([]);
    const [certificates, setCertificates] = useState<Certificate[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loadData = async () => {
            try {
                const [coursesData, studentsData, certificatesData] = await Promise.all([
                    fetchCourses(),
                    fetchStudents(),
                    fetchCertificates(),
                ]);

                setCourses(coursesData ?? []);
                setStudents(studentsData ?? []);
                setCertificates(certificatesData ?? []);
            } catch (error) {
                console.error('Ошибка загрузки данных:', error);
                setCourses([]);
                setStudents([]);
                setCertificates([]);
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, []);

    const coursesCount = useMemo(() => courses.length, [courses]);
    const studentsCount = useMemo(() => students.length, [students]);
    const certificatesCount = useMemo(() => certificates.length, [certificates]);

    if (loading) {
        return <Loader />;
    }

    return (
        <div className="dashboard-container">
            <div
                className="page-header"
                style={{
                    position: 'relative',
                    justifyContent: 'center',
                    textAlign: 'center',
                    paddingRight: '4rem',
                }}
            >
                <h1 style={{ marginBottom: 0 }}>Административная панель Gwynbleidd</h1>
                <button
                    type="button"
                    className="btn secondary"
                    style={{
                        position: 'absolute',
                        top: 0,
                        right: 0,
                    }}
                    onClick={() => {
                        logout();
                        navigate('/login', { replace: true });
                    }}
                >
                    Выйти
                </button>
            </div>
            <div className="stats-grid">
                <StatsCard
                    title="Курсов на платформе"
                    count={coursesCount}
                    link="/courses"
                />
                <StatsCard
                    title="Студентов обучается"
                    count={studentsCount}
                    link="/students"
                />
                <StatsCard
                    title="Сертификатов выдано"
                    count={certificatesCount}
                    link="/certificates"
                />
            </div>
        </div>
    );
};

export default React.memo(Dashboard);
