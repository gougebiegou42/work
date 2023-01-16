import request from '@/utils/request';

// 登录
export function login(username, password, code, uuid) {
  const data = {
    username,
    password,
    verifyCode: code,
    verifyKey: uuid,
  };
  return request({
    url: '/system/login',
    method: 'post',
    data: data,
  });
}

// 退出登录
export function logout() {
  return request({
    url: '/system/logout',
    method: 'post',
  });
}

// 获取验证码 （获取uuid）
export function getCodeImg() {
  return request({
    url: '/captcha/get',
    method: 'get',
  });
}
