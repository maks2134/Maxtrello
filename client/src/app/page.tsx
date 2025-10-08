"use client";

import React, { useState } from 'react';
import AuthForm from '../components/AuthForm';
import Dashboard from '../components/Dashboard';
import { useAuth } from '@/contexts/AuthContext';

const HomePage: React.FC = () => {
    const { user } = useAuth();
    const [isRegisterMode, setIsRegisterMode] = useState(false);

    return (
        <div className="min-h-screen bg-gray-900 flex flex-col items-center justify-center p-4">
            <header className="mb-10 text-center">
                <h1 className="text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-teal-400 to-blue-500 tracking-tight">
                    MaxTrello
                </h1>
            </header>

            <main className="w-full flex justify-center">
                {user ? (
                    <Dashboard />
                ) : (
                    <AuthForm
                        isRegister={isRegisterMode}
                        onSwitchMode={() => setIsRegisterMode(prev => !prev)}
                    />
                )}
            </main>
        </div>
    );
};

export default HomePage;
