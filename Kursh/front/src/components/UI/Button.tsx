import React from 'react';

type ButtonProps = {
    children: React.ReactNode;
    type?: 'button' | 'submit' | 'reset';
    onClick?: () => void;
    fullWidth?: boolean;
    variant?: 'primary' | 'secondary' | 'danger';
};

const Button: React.FC<ButtonProps> = ({
                                           children,
                                           type = 'button',
                                           onClick,
                                           fullWidth = false,
                                           variant = 'primary'
                                       }) => (
    <button
        className={`btn ${variant} ${fullWidth ? 'full-width' : ''}`}
        type={type}
        onClick={onClick}
    >
        {children}
    </button>
);

export default Button;