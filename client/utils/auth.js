import { httpGet } from './http'

const API_BASE = '/api/oauth'
const CHECK = API_BASE + '/check'


export function check() {
  return httpGet(CHECK)
}
