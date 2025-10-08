"use client";

import React from 'react';
import { LogOut, Home } from 'lucide-react';
import Button from './ui/Button';
import { useAuth } from '@/contexts/AuthContext';

const Dashboard: React.FC = () => {
    const { user, token, logout } = useAuth();

    if (!user) return null;

    return (
        <div className="bg-gray-800 p-8 rounded-xl shadow-2xl w-full max-w-xl border border-teal-600">
            <h2 className="text-4xl font-extrabold text-white mb-6 border-b border-teal-700 pb-3 flex items-center">
                <Home className="w-8 h-8 mr-3 text-teal-400" />
                Личный кабинет
            </h2>
            <div className="space-y-4 text-lg text-gray-300">
                <p className='text-2xl'>
                    <span className="font-semibold text-teal-400">Привет,</span> {user.name}!
                </p>
                <div className="p-4 bg-gray-700 rounded-lg">
                    <p>
                        <span className="font-medium text-gray-400 block text-sm">Email:</span>
                        <span className="text-white font-mono">{user.email}</span>
                    </p>
                </div>
                <div className="p-4 bg-gray-700 rounded-lg">
                    <p>
                        <span className="font-medium text-gray-400 block text-sm">Ваш ID:</span>
                        <span className="text-white font-mono break-all">{user.id}</span>
                    </p>
                </div>
                <div className="p-4 bg-gray-700 rounded-lg">
                    <p>
                        <span className="font-medium text-gray-400 block text-sm">Ваш JWT-токен:</span>
                        <span className="text-sm text-yellow-400 font-mono break-all">
                            {token ? `${token.substring(0, 40)}...` : 'N/A'}
                        </span>
                    </p>
                </div>
            </div>
            <div className="mt-8">
                <Button onClick={logout} icon={LogOut}>
                    Выйти
                </Button>
            </div>
        </div>
    );
};

export default Dashboard;
