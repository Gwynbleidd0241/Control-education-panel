import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { fetchStudents } from '../services/studentService';
import { Student } from '../types';
import Loader from '../components/Loader';
import StudentTableRow from '../components/StudentTableRow';
import { renderLoadByLevel } from '../utils/performanceHelpers';
import { useMediaQuery } from 'react-responsive';

const Students = () => {
    const [students, setStudents] = useState<Student[]>([]);
    const [loading, setLoading] = useState(true);

    const isMobile = useMediaQuery({ maxWidth: 767 });

    useEffect(() => {
        const loadData = async () => {
            try {
                const data = await fetchStudents();
                setStudents(data);
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
                    {students.map(student => (
                        <div key={student.id} className="student-card">
                            <h3>{student.fullName}</h3>
                            <p><strong>Email:</strong> {student.email}</p>
                            <p>
                                <strong>Курс:</strong>{' '}
                                {student.course?.title ?? 'Без курса'}
                            </p>
                            <p>
                                <strong>Нагрузка:</strong>{' '}
                                {student.course
                                    ? renderLoadByLevel(student.course.level)
                                    : '—'}
                            </p>
                            <Link to={`/students/${student.id}`}>Подробнее</Link>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default Students;
