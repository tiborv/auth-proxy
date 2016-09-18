import { httpGet, httpPost } from './http'

const API_BASE = '/api/client'
const LIST = API_BASE + '/list'
const CREATE = API_BASE + '/create'
const UPDATE = API_BASE + '/update'
const DELETE = API_BASE + '/delete'


export function list() {
  return httpGet(LIST)
}

export function create(token) {
  return httpPost(CREATE, token)
}

export function update(token) {
  return httpPost(UPDATE, token)
}

export function del(client) {
  return httpPost(DELETE, client)
}
