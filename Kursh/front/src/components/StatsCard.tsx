import { Link } from 'react-router-dom';

interface StatsCardProps {
    title: string;
    count: number;
    link: string;
}

const StatsCard = ({ title, count, link }: StatsCardProps) => (
    <Link to={link} className="stats-card">
        <div className="stats-card-inner">
            <h3>{title}</h3>
            <div className="stats-content">
                <span className="count">{count}</span>
            </div>
        </div>
    </Link>
);

export default StatsCard;