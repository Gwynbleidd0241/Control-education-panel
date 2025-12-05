import { useEffect } from 'react';

interface NotificationProps {
    type: 'success' | 'error' | 'info';
    message: string;
    onClose?: () => void;
}

const Notification = ({ type, message, onClose }: NotificationProps) => {
    useEffect(() => {
        const timer = setTimeout(() => {
            if (onClose) onClose();
        }, 3000);
        return () => clearTimeout(timer);
    }, [onClose]);

    return (
        <div className={`notification ${type}`}>
            <span>{message}</span>
        </div>
    );
};

export default Notification;