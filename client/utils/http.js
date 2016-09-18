import * as request from 'superagent'
import Promise from 'bluebird'

import { notAuthorized } from '../actions/auth'

const StatusUnauthorized = 401

export function httpPost(url, body) {
  return new Promise((resolve, reject) => {
    request
      .post(url)
      .send(body)
      .end((err, res) => {
        if (err) return reject(err)
        return resolve(res.body)
      })
  })
}


export function httpGet(url, query) {
  return new Promise((resolve, reject) => {
    request
      .get(url)
      .query(query)
      .end((err, res) => {
        if (err) return reject(err)
        return resolve(res.body)
      })
  })
}
