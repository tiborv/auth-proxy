import { httpGet, httpPost } from './http'

const API_BASE = '/api/service'
const LIST = API_BASE + '/list'
const CREATE = API_BASE + '/create'
const UPDATE = API_BASE + '/update'
const DELETE = API_BASE + '/delete'

export function list() {
  return httpGet(LIST)
}

export function create(service) {
  return httpPost(CREATE, service)
}

export function update(service) {
  return httpPost(UPDATE, service)
}

export function del(service) {
  return httpPost(DELETE, service)
}
