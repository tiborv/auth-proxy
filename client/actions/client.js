import * as api from '../utils/client'

export const LIST_SUCESS = 'CLIENT.LIST_SUCESS'
export const list = () => (dispatch) => {
  api
    .list()
    .then(tokens => dispatch({
      type: LIST_SUCESS,
      tokens,
    }))
}

export const DELETE_SUCESS = 'CLIENT.DELETE_SUCESS'
export const del = (token) => (dispatch) => {
  api
    .del(token)
    .then(() => dispatch({
      type: DELETE_SUCESS,
      token,
    }))
}

export const UPDATE_SUCESS = 'CLIENT.UPDATE_SUCESS'
export const update = (client) => (dispatch) => {
  api
    .update(client)
    .then(() => dispatch({
      type: UPDATE_SUCESS,
      client
    }))
}

export const CREATE_SUCESS = 'CLIENT.CREATE_SUCESS'
export const create = (client) => (dispatch) => {
  api
    .create(client)
    .then(() => dispatch({
      type: CREATE_SUCESS,
      client
    }))
}
