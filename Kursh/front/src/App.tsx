import React, { Suspense, lazy } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Loader from './components/Loader';
import ProtectedRoute from './components/ProtectedRoute';
import './styles/index.css';

const Login = lazy(() => import('./pages/Login'));
const Dashboard = lazy(() => import('./pages/Dashboard'));
const Courses = lazy(() => import('./pages/Courses'));
const CourseDetails = lazy(() => import('./pages/CourseDetails'));
const CourseForm = lazy(() => import('./pages/CourseForm'));
const Students = lazy(() => import('./pages/Students'));
const StudentDetails = lazy(() => import('./pages/StudentDetails'));
const Certificates = lazy(() => import('./pages/Certificates'));
const CertificateView = lazy(() => import('./pages/CertificateView'));
const VerifyCertificate = lazy(() => import('./pages/VerifyCertificate'));

const App: React.FC = () => (
    <Router>
        <Suspense fallback={<Loader />}>
            <Routes>
                <Route path="/login" element={<Login />} />
                <Route element={<ProtectedRoute />}>
                    <Route path="/" element={<Dashboard />} />
                    <Route path="/courses" element={<Courses />} />
                    <Route path="/courses/new" element={<CourseForm />} />
                    <Route path="/courses/:id" element={<CourseDetails />} />
                    <Route path="/courses/:id/edit" element={<CourseForm />} />

                    <Route path="/students" element={<Students />} />
                    <Route path="/students/:id" element={<StudentDetails />} />

                    <Route path="/certificates" element={<Certificates />} />
                    <Route path="/certificates/:id" element={<CertificateView />} />
                    <Route
                        path="/verify-certificate"
                        element={<VerifyCertificate />}
                    />
                </Route>

                <Route path="*" element={<Navigate to="/" replace />} />
            </Routes>
        </Suspense>
    </Router>
);

export default App;
