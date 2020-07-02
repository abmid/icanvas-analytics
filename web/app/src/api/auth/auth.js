import request from '@/helpers/request';

export function login(email, password) {
    const data = {
        email,
        password,
      };
      return request({
        url: '/auth/login',
        method: 'post',
        data,
        withCredentials: true
      });    
}

export function logout() {
    return request({
        url: "/auth/logout",
        method: "post",
        withCredentials: true
    })
}