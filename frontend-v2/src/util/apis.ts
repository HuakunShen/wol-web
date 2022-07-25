import axios from 'axios';
import { AxiosResponse } from 'axios';

export type Computer = {
  id: number;
  user_id: number;
  name: string;
  mac: string;
  ip: string;
  port: number;
};

export const getUserCount = (): Promise<
  AxiosResponse<{
    data: {
      num_user_allowed: number;
      user_count: number;
    };
    message: string;
  }>
> => {
  return axios.get('/api/users/count');
};

export const getComputers = (): Promise<AxiosResponse<{ data: Computer[] }>> =>
  axios.get('/api/computers');

export const addComputer = (
  name: string,
  mac: string,
  ip: string,
  port: string
): Promise<AxiosResponse<{ data: Computer }>> =>
  axios.post('/api/computers', {
    name,
    mac,
    ip,
    port,
  });

export const deleteComputer = (
  id: number
): Promise<AxiosResponse<{ data: Computer; message: string }>> =>
  axios.delete(`/api/computers/${id}`);

export const loadAuth = () =>
  axios.get('/api/users/user', {
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  });

export const logout = (): Promise<AxiosResponse<{ message: string }>> =>
  axios.post('/api/users/logout');

export const login = (
  username: string,
  password: string
): Promise<
  AxiosResponse<{
    data: { id: number; username: string };
    message: string;
  }>
> =>
  axios.post('/api/users/login', {
    username: username,
    password: password,
  });

export const wol = (id: number) => axios.post(`/api/wol/${id}`);
