// src/pages/Certificates.tsx
import React, { useEffect, useMemo, useState } from 'react';
import Navbar from '../components/Navbar';
import Loader from '../components/Loader';
import { Certificate, CertificateStatus } from '../types';
import { fetchCertificates } from '../services/certificateService';
import '../styles/certificates.css';

type StatusFilter = 'all' | CertificateStatus;

const statusLabels: Record<CertificateStatus, string> = {
    valid: 'Действует',
    expired: 'Истёк',
    revoked: 'Аннулирован',
};

const formatDate = (iso?: string | null) =>
    iso ? new Date(iso).toLocaleDateString('ru-RU') : '—';

const Certificates: React.FC = () => {
    const [certificates, setCertificates] = useState<Certificate[]>([]);
    const [loading, setLoading] = useState(true);
    const [statusFilter, setStatusFilter] = useState<StatusFilter>('all');
    const [search, setSearch] = useState('');

    useEffect(() => {
        const load = async () => {
            try {
                const data = await fetchCertificates();
                setCertificates(data);
            } catch (e) {
                console.error('Ошибка загрузки сертификатов:', e);
            } finally {
                setLoading(false);
            }
        };

        load();
    }, []);

    const filtered = useMemo(
        () =>
            certificates.filter((cert) => {
                if (statusFilter !== 'all' && cert.status !== statusFilter) {
                    return false;
                }

                if (!search.trim()) return true;

                const q = search.toLowerCase();

                return (
                    cert.studentName.toLowerCase().includes(q) ||
                    cert.courseTitle.toLowerCase().includes(q) ||
                    cert.code.toLowerCase().includes(q)
                );
            }),
        [certificates, statusFilter, search],
    );

    if (loading) {
        return (
            <>
                <Navbar />
                <Loader />
            </>
        );
    }

    return (
        <>
            <Navbar />

            <main className="certificates-page">
                <header className="page-header">
                    <h1>Сертификаты</h1>
                    <div className="actions-group">
                        <button type="button" className="btn secondary">
                            Экспорт в CSV
                        </button>
                        <button type="button" className="btn primary">
                            Выдать сертификат
                        </button>
                    </div>
                </header>

                <section className="filters-row">
                    <div className="filters-group">
                        <button
                            type="button"
                            className={`filter-chip ${statusFilter === 'all' ? 'active' : ''}`}
                            onClick={() => setStatusFilter('all')}
                        >
                            Все
                        </button>
                        <button
                            type="button"
                            className={`filter-chip ${statusFilter === 'valid' ? 'active' : ''}`}
                            onClick={() => setStatusFilter('valid')}
                        >
                            Действующие
                        </button>
                        <button
                            type="button"
                            className={`filter-chip ${statusFilter === 'expired' ? 'active' : ''}`}
                            onClick={() => setStatusFilter('expired')}
                        >
                            Истёкшие
                        </button>
                        <button
                            type="button"
                            className={`filter-chip ${statusFilter === 'revoked' ? 'active' : ''}`}
                            onClick={() => setStatusFilter('revoked')}
                        >
                            Аннулированные
                        </button>
                    </div>

                    <div className="search-wrapper">
                        <input
                            type="text"
                            className="search-input"
                            placeholder="Поиск по студенту, курсу или коду…"
                            value={search}
                            onChange={(e) => setSearch(e.target.value)}
                        />
                    </div>
                </section>

                <section className="table-wrapper">
                    {filtered.length === 0 ? (
                        <p className="empty-state">
                            По текущим фильтрам сертификатов не найдено.
                        </p>
                    ) : (
                        <table className="certificates-table">
                            <thead>
                            <tr>
                                <th>Код</th>
                                <th>Студент</th>
                                <th>Курс</th>
                                <th>Оценка</th>
                                <th>Выдан</th>
                                <th>Действителен до</th>
                                <th>Статус</th>
                            </tr>
                            </thead>
                            <tbody>
                            {filtered.map((cert) => (
                                <tr key={cert.id}>
                                    <td>{cert.code}</td>
                                    <td>{cert.studentName}</td>
                                    <td>{cert.courseTitle}</td>
                                    <td>{cert.grade ?? '—'}</td>
                                    <td>{formatDate(cert.issuedAt)}</td>
                                    <td>{formatDate(cert.expiresAt)}</td>
                                    <td>
                                            <span
                                                className={`status-badge status-${cert.status}`}
                                            >
                                                {statusLabels[cert.status]}
                                            </span>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    )}
                </section>
            </main>
        </>
    );
};

export default Certificates;
