import React, { useState } from 'react';
import Loader from '../components/Loader';
import { verifyCertificateByCode } from '../services/certificateService';
import { Certificate } from '../types';
import '../styles/certificates.css';

const VerifyCertificate: React.FC = () => {
    const [code, setCode] = useState('');
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState<Certificate | null | undefined>(undefined);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        try {
            const cert = await verifyCertificateByCode(code.trim());
            setResult(cert ?? null);
        } catch (e) {
            console.error('Ошибка проверки сертификата', e);
            setResult(null);
        } finally {
            setLoading(false);
        }
    };

    return (
        <main className="certificates-page">
            <header className="page-header">
                <h1>Проверка сертификата</h1>
                <p className="page-subtitle">
                    Введите номер сертификата, чтобы проверить его действительность.
                </p>
            </header>

            <section className="table-wrapper" style={{ padding: '1.5rem' }}>
                <form onSubmit={handleSubmit} className="filters-row">
                    <input
                        type="text"
                        className="search-input"
                        placeholder="Например: CERT-001"
                        value={code}
                        onChange={(e) => setCode(e.target.value)}
                    />
                    <button type="submit" className="btn primary">
                        Проверить
                    </button>
                </form>

                {loading && <Loader />}

                {!loading && result !== undefined && (
                    <div style={{ marginTop: '1.5rem' }}>
                        {result ? (
                            <div>
                                <h2>Сертификат найден</h2>
                                <p>
                                    <strong>Код:</strong> {result.code}
                                </p>
                                <p>
                                    <strong>Владелец:</strong> {result.studentName}
                                </p>
                                <p>
                                    <strong>Курс:</strong> {result.courseTitle}
                                </p>
                                <p>
                                    <strong>Дата выдачи:</strong> {result.issuedAt}
                                </p>
                            </div>
                        ) : (
                            <p className="empty-state">
                                Сертификат с номером <strong>{code}</strong> не найден.
                            </p>
                        )}
                    </div>
                )}
            </section>
        </main>
    );
};

export default VerifyCertificate;
