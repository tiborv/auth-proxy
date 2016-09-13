import * as api from '../utils/token'

export const LIST_SUCESS = 'TOKEN.LIST_SUCESS'
export const list = () => (dispatch) => {
  api
    .list()
    .then(tokens => dispatch({
      type: LIST_SUCESS,
      tokens,
    }))
}

export const DELETE_SUCESS = 'TOKEN.DELETE_SUCESS'
export const del = (id) => (dispatch) => {
  api
    .del(id)
    .then(() => dispatch({
      type: DELETE_SUCESS,
      id
    }))
}

export const UPDATE_SUCESS = 'TOKEN.UPDATE_SUCESS'
export const update = (token) => (dispatch) => {
  api
    .update(token)
    .then(() => dispatch({
      type: UPDATE_SUCESS,
      token
    }))
}

export const CREATE_SUCESS = 'TOKEN.CREATE_SUCESS'
export const create = (token) => (dispatch) => {
  api
    .create(token)
    .then(() => dispatch({
      type: CREATE_SUCESS,
      token
    }))
}
