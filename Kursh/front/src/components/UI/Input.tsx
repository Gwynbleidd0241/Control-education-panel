import React from 'react';

type InputProps = {
    label?: string;
    type?: React.HTMLInputTypeAttribute | 'textarea';
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
    required?: boolean;
    placeholder?: string;
    min?: string;
};

const Input: React.FC<InputProps> = ({
                                         label,
                                         type = 'text',
                                         value,
                                         onChange,
                                         required = false,
                                         placeholder,
                                         min
                                     }) => {
    return (
        <div className="input-group">
            {label && <label>{label}</label>}
            {type === 'textarea' ? (
                <textarea
                    value={value}
                    onChange={onChange}
                    required={required}
                    placeholder={placeholder}
                    rows={4}
                />
            ) : (
                <input
                    type={type}
                    value={value}
                    onChange={onChange}
                    required={required}
                    placeholder={placeholder}
                    min={min}
                />
            )}
        </div>
    );
};

export default Input;