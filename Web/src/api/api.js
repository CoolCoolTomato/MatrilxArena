import axios from 'axios';
import config from '@/config/config.js';

export const apiClient = axios.create({
    baseURL: config.baseURL,
    withCredentials: false,
    headers: {
        Accept: 'application/json',
    }
});

apiClient.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  if (!(config.data instanceof FormData)) {
    config.headers['Content-Type'] = 'application/json';
  }
  return config;
});

export const handleResponse = (response) => {
    return response.data;
};

export const handleError = (error) => {
    console.error('API call error:', error);
    throw error;
};
