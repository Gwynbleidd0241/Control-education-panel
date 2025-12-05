// src/pages/Dashboard.tsx
import React, { useEffect, useMemo, useState } from 'react';
import { fetchCourses } from '../services/courseService';
import { fetchStudents } from '../services/studentService';
import Loader from '../components/Loader';
import StatsCard from '../components/StatsCard';
import LogoutButton from '../components/LogoutButton';
import type { Course, Student } from '../types';

const Dashboard: React.FC = () => {
    const [courses, setCourses] = useState<Course[]>([]);
    const [students, setStudents] = useState<Student[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loadData = async () => {
            try {
                const [coursesData, studentsData] = await Promise.all([
                    fetchCourses(),
                    fetchStudents(),
                ]);

                // сохраняем массивы в стейт
                setCourses(coursesData ?? []);
                setStudents(studentsData ?? []);
            } catch (error) {
                console.error('Ошибка загрузки данных:', error);
                // на всякий случай очистим, чтобы не было null
                setCourses([]);
                setStudents([]);
            } finally {
                setLoading(false);
            }
        };

        loadData();
    }, []);

    // если очень хочется useMemo — считаем длины через него
    const coursesCount = useMemo(() => courses.length, [courses]);
    const studentsCount = useMemo(() => students.length, [students]);

    if (loading) {
        return <Loader />;
    }

    return (
        <div className="dashboard-container">
            <LogoutButton />
            <h1>Административная панель Gwynbleidd</h1>
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
            </div>
        </div>
    );
};

export default React.memo(Dashboard);
