import React, { ButtonHTMLAttributes } from 'react';
import { Loader, LucideIcon } from 'lucide-react';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    icon?: LucideIcon;
    loading?: boolean;
    children: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({ children, icon: Icon, type = 'button', loading = false, disabled = false, ...props }) => (
    <button
        type={type}
        disabled={disabled || loading}
        className={`w-full flex justify-center items-center px-4 py-2 text-white font-semibold rounded-lg shadow-md transition duration-200 ease-in-out transform 
            ${disabled || loading
            ? 'bg-gray-500 cursor-not-allowed'
            : 'bg-teal-600 hover:bg-teal-500 hover:scale-[1.01] focus:outline-none focus:ring-4 focus:ring-teal-500/50'
        }`}
        {...props}
    >
        {loading ? <Loader className="w-5 h-5 mr-2 animate-spin" /> : Icon && <Icon className="w-5 h-5 mr-2" />}
        {children}
    </button>
);

export default Button;