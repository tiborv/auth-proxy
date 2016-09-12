import { httpPost } from './http'

const API_LOGIN = '/api/auth/login'
const API_LOGOUT = '/api/auth/logout'

export function login(username, password) {
  return httpPost(API_LOGIN, { username, password })
}

export function logout() {
  return httpPost(API_LOGOUT)
}
