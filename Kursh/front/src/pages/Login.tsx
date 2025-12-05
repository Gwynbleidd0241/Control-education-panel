import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login, isAuthenticated } from '../utils/auth';
import '../styles/login.css';

const Login: React.FC = () => {
    const navigate = useNavigate();

    const [username, setUsername] = useState('admin');
    const [password, setPassword] = useState('admin');
    const [error, setError] = useState('');
    const [submitting, setSubmitting] = useState(false);

    useEffect(() => {
        if (isAuthenticated()) {
            navigate('/', { replace: true });
        }
    }, [navigate]);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setSubmitting(true);

        const ok = login(username.trim(), password.trim());

        if (!ok) {
            setError('Неверный логин или пароль');
            setSubmitting(false);
            return;
        }

        navigate('/', { replace: true });
    };

    return (
        <div className="login-page">
            <div className="login-form-container">
                <h1>Авторизация администратора</h1>

                <form onSubmit={handleSubmit}>
                    <div className="input-group">
                        <label htmlFor="username">Логин</label>
                        <input
                            id="username"
                            type="text"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>

                    <div className="input-group">
                        <label htmlFor="password">Пароль</label>
                        <input
                            id="password"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>

                    {error && <div className="notification error">{error}</div>}

                    <button
                        type="submit"
                        className="login-button"
                        disabled={submitting}
                    >
                        {submitting ? 'Входим…' : 'Войти'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default Login;
