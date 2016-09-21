import * as api from '../utils/service'

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
export const del = (service) => (dispatch) => {
  api
    .del(service)
    .then(() => dispatch({
      type: DELETE_SUCESS,
      service
    }))
}

export const UPDATE_SUCESS = 'SERVICE.UPDATE_SUCESS'
export const update = (service) => (dispatch) => {
  api
    .update(service)
    .then(updatedService => dispatch({
      type: UPDATE_SUCESS,
      updatedService
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
