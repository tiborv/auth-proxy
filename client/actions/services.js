import * as api from '../utils/sertice'

export const LIST_SUCESS = 'SERVICE.LIST_SUCESS'
export const list = () => (dispatch) => {
  api
    .list()
    .then(services => dispatch({
      type: LIST_SUCESS,
      services,
    }))
}

export const DELETE_SUCESS = 'SERVICE.DELETE_SUCESS'
export const del = (id) => (dispatch) => {
  api
    .del(slug)
    .then(() => dispatch({
      type: DELETE_SUCESS,
      id
    }))
}

export const UPDATE_SUCESS = 'SERVICE.UPDATE_SUCESS'
export const update = (service) => (dispatch) => {
  api
    .update(service)
    .then(() => dispatch({
      type: UPDATE_SUCESS,
      service
    }))
}

export const CREATE_SUCESS = 'SERVICE.CREATE_SUCESS'
export const create = (service) => (dispatch) => {
  api
    .create(service)
    .then(() => dispatch({
      type: CREATE_SUCESS,
      service
    }))
}
