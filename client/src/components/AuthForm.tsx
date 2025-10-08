"use client";

import React, { useState } from 'react';
import { LogIn, UserPlus, AlertTriangle } from 'lucide-react';
import Input from './ui/Input';
import Button from './ui/Button';
import { useAuth } from '@/contexts/AuthContext';

interface AuthFormProps {
    isRegister: boolean;
    onSwitchMode: () => void;
}

const AuthForm: React.FC<AuthFormProps> = ({ isRegister, onSwitchMode }) => {
    const { login, register, error, isLoading, setError } = useAuth();

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');

        if (isRegister) {
            if (name && email && password) {
                await register(name, email, password);
            }
        } else {
            if (email && password) {
                await login(email, password);
            }
        }
    };

    return (
        <div className="bg-gray-800 p-8 rounded-xl shadow-2xl w-full max-w-md border border-gray-700">
            <div className="flex justify-between items-center mb-6">
                <h2 className="text-3xl font-bold text-white flex items-center">
                    {isRegister ? <UserPlus className="w-7 h-7 mr-3 text-teal-400" /> : <LogIn className="w-7 h-7 mr-3 text-teal-400" />}
                    {isRegister ? 'Регистрация' : 'Авторизация'}
                </h2>

            </div>
            {error && (
                <div className="mb-4 p-3 bg-red-800/50 border border-red-600 text-red-300 rounded-lg flex items-center text-sm">
                    <AlertTriangle className="w-5 h-5 mr-2" />
                    {error}
                </div>
            )}
            <form onSubmit={handleSubmit}>
                {isRegister && (
                    <Input
                        id="name"
                        label="Имя"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                    />
                )}
                <Input
                    id="email"
                    label="Email"
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <Input
                    id="password"
                    label="Пароль"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <div className="mt-6 flex flex-col items-center ">
                    <Button type="submit" icon={isRegister ? UserPlus : LogIn} loading={isLoading}>
                        {isRegister ? 'Зарегистрироваться' : 'Войти'}
                    </Button>
                    <button
                        onClick={onSwitchMode}
                        className="text-sm mt-4 text-teal-400 hover:text-teal-300 transition duration-150"
                    >
                        {isRegister ? 'Уже есть аккаунт? Войти' : 'Нет аккаунта? Зарегистрироваться'}
                    </button>
                </div>
            </form>
        </div>
    );
};

export default AuthForm;
