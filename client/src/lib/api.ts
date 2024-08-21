import axios from 'axios';

const url = {
  development: 'http://localhost:4000/api/v1',
  production: 'https://api.example.com/api/v1',
  test: 'http://localhost:4000/api/v1',
};

const api = axios.create({
  baseURL: url[process.env.NODE_ENV],
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
  withCredentials: true,
});

export default api;
