import './globals.css';
import { AuthContextProvider } from '@/contexts/AuthContext';
import React, { ReactNode } from 'react';

export const metadata = {
    title: 'Trello Clone - User Auth',
    description: 'Frontend for User Microservice',
};

const RootLayout: React.FC<{ children: ReactNode }> = ({ children }) => {
    return (
        <html lang="ru">
        <body className="font-sans antialiased">
        <AuthContextProvider>
            {children}
        </AuthContextProvider>
        </body>
        </html>
    );
}

export default RootLayout;
