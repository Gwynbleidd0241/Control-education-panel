import React, { useEffect, useState } from 'react';
import Navbar from '../components/Navbar';
import Loader from '../components/Loader';
import { CertificateTemplate } from '../types';
import { fetchCertificateTemplates } from '../services/certificateService';
import '../styles/certificates.css';

const CertificateTemplates: React.FC = () => {
    const [templates, setTemplates] = useState<CertificateTemplate[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const load = async () => {
            try {
                const data = await fetchCertificateTemplates();
                setTemplates(data);
            } catch (e) {
                console.error('Ошибка загрузки шаблонов', e);
            } finally {
                setLoading(false);
            }
        };
        load();
    }, []);

    if (loading) return <Loader />;

    return (
        <>
            <Navbar />
            <main className="certificates-page">
                <header className="page-header">
                    <h1>Шаблоны сертификатов</h1>
                </header>

                <section className="table-wrapper">
                    {templates.length === 0 ? (
                        <p className="empty-state">Шаблоны сертификатов не найдены.</p>
                    ) : (
                        <table className="certificates-table">
                            <thead>
                            <tr>
                                <th>Название</th>
                                <th>Курс</th>
                                <th>Срок действия</th>
                            </tr>
                            </thead>
                            <tbody>
                            {templates.map((tpl) => (
                                <tr key={tpl.id}>
                                    <td>{tpl.name}</td>
                                    <td>{tpl.courseTitle ?? tpl.courseId}</td>
                                    <td>
                                        {tpl.validityDays
                                            ? `${tpl.validityDays} дней`
                                            : 'Бессрочный'}
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

export default CertificateTemplates;
