import { httpGet, httpPost } from './http'

const API_BASE = '/api/token'
const LIST = API_BASE + '/list'
const CREATE = API_BASE + '/create'
const UPDATE = API_BASE + '/update'
const DELETE = API_BASE + '/delete'
const ADD_SERVICE = API_BASE + '/service'


export function list() {
  return httpGet(LIST)
}

export function create(token) {
  return httpPost(CREATE, token)
}

export function update(token) {
  return httpPost(UPDATE, token)
}

export function del(id) {
  return httpPost(DELETE, { id })
}

export function addService(serviceId) {
  return httpPost(ADD_SERVICE, { id: serviceId })
}
