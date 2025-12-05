import { Link } from 'react-router-dom';
import { Student } from '../types';
import { renderLoadByLevel } from '../utils/performanceHelpers';

interface Props {
    student: Student;
}

const StudentTableRow: React.FC<Props> = ({ student }) => {
    const course = student.course; // может быть undefined/null

    return (
        <tr>
            <td>{student.fullName}</td>
            <td>{student.email}</td>

            <td>
                {course ? (
                    <Link to={`/courses/${course.id}`} state={{ from: '/students' }}>
                        {course.title}
                    </Link>
                ) : (
                    'Без курса'
                )}
            </td>

            <td>
                {course ? renderLoadByLevel(course.level) : '—'}
            </td>

            <td>
                <Link to={`/students/${student.id}`} className="action-link">
                    Просмотр
                </Link>
            </td>
        </tr>
    );
};

export default StudentTableRow;
