import { logout } from '../utils/auth';

const Navbar = () => {
    const handleLogout = () => {
        logout();
        window.location.reload();
    };

    return (
        <nav className="navbar" style={{ justifyContent: 'flex-end' }}>
            <button onClick={handleLogout} className="logout-btn">
                Выйти
            </button>
        </nav>
    );
};

export default Navbar;
