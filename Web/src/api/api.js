import axios from 'axios';
import config from '@/config/config.js';

export const apiClient = axios.create({
    baseURL: config.baseURL,
    withCredentials: false,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    }
});

export const handleResponse = (response) => {
    return response.data;
};

export const handleError = (error) => {
    console.error('API call error:', error);
    throw error;
};
