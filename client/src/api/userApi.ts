import { apiCall } from './api';

export interface User {
    id: string;
    email: string;
    name: string;
}

export interface LoginResponse {
    token: string;
    user: User;
}

export interface Credentials {
    email: string;
    password: string;
}

export const login = async (credentials: Credentials): Promise<LoginResponse> => {
    return apiCall<LoginResponse>('/users/login', 'POST', credentials);
};

export const register = async (userData: Credentials & { name: string }): Promise<User> => {
    return apiCall<User>('/users/register', 'POST', userData);
};