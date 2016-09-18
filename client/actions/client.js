import * as api from '../utils/client'

export const LIST_SUCESS = 'CLIENT.LIST_SUCESS'
export const list = ({ dispatch }) => {
  api
    .list()
    .then(clients => dispatch({
      type: LIST_SUCESS,
      clients,
    }))
}

export const DELETE_SUCESS = 'CLIENT.DELETE_SUCESS'
export const del = (client) => (dispatch) => {
  api
    .del(client)
    .then(() => dispatch({
      type: DELETE_SUCESS,
      client,
    }))
}

export const UPDATE_SUCESS = 'CLIENT.UPDATE_SUCESS'
export const update = (updateClient) => (dispatch) => {
  api
    .update(updateClient)
    .then(client => dispatch({
      type: UPDATE_SUCESS,
      client
    }))
}

export const CREATE_SUCESS = 'CLIENT.CREATE_SUCESS'
export const create = (newClient) => (dispatch) => {
  api
    .create(newClient)
    .then(client => dispatch({
      type: CREATE_SUCESS,
      client
    }))
}
