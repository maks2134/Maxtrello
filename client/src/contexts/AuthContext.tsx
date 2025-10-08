"use client";

import React, { createContext, useContext, useState, useEffect, useCallback, ReactNode } from 'react';
import { login as apiLogin, register as apiRegister, User, LoginResponse } from '../api/userApi';

interface AuthContextType {
    user: User | null;
    token: string | null;
    error: string;
    isLoading: boolean;
    login: (email: string, password: string) => Promise<boolean>;
    register: (name: string, email: string, password: string) => Promise<boolean>;
    logout: () => void;
    setError: React.Dispatch<React.SetStateAction<string>>;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth должен использоваться внутри AuthContextProvider');
    }
    return context;
};

export const AuthContextProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);
    const [token, setToken] = useState<string | null>(null);
    const [error, setError] = useState<string>('');
    const [isLoading, setIsLoading] = useState<boolean>(true);

    const updateAuth = useCallback((newToken: string | null, newUser: User | null) => {
        if (newToken && newUser) {
            localStorage.setItem('jwtToken', newToken);
            localStorage.setItem('user', JSON.stringify(newUser));
            setToken(newToken);
            setUser(newUser);
        } else {
            localStorage.removeItem('jwtToken');
            localStorage.removeItem('user');
            setToken(null);
            setUser(null);
        }
        setError('');
    }, []);

    useEffect(() => {
        const savedToken = localStorage.getItem('jwtToken');
        const savedUser = localStorage.getItem('user');

        if (savedToken && savedUser) {
            try {
                setToken(savedToken);
                setUser(JSON.parse(savedUser) as User);
            } catch (e) {
                updateAuth(null, null);
            }
        }
        setIsLoading(false);
    }, [updateAuth]);


    const login = async (email: string, password: string): Promise<boolean> => {
        setIsLoading(true);
        setError('');
        try {
            const data: LoginResponse = await apiLogin({ email, password });
            if (data && data.token && data.user) {
                updateAuth(data.token, data.user);
                return true;
            }
            return false;
        } catch (e) {
            setError(e instanceof Error ? e.message : "Ошибка входа");
            return false;
        } finally {
            setIsLoading(false);
        }
    };

    const register = async (name: string, email: string, password: string): Promise<boolean> => {
        setIsLoading(true);
        setError('');
        try {
            await apiRegister({ name, email, password });

            // Автоматический логин после успешной регистрации
            const success = await login(email, password);
            if (success) {
                return true;
            } else {
                setError('Регистрация успешна, но не удалось войти. Попробуйте войти вручную.');
                return false;
            }

        } catch (e) {
            setError(e instanceof Error ? e.message : "Ошибка регистрации");
            return false;
        } finally {
            if (!user) setIsLoading(false);
        }
    };

    const logout = () => {
        updateAuth(null, null);
    };

    return (
        <AuthContext.Provider value={{ user, token, error, isLoading, login, register, logout, setError }}>
            {!isLoading && children}
        </AuthContext.Provider>
    );
};
