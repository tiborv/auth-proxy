import * as api from '../utils/auth'

export const NOT_AUTH = 'AUTH.NOT_AUTH'
export const notAuthorized = () => (dispatch) => {
  dispatch({
    type: NOT_AUTH,
    auth: false,
  })
}

export const AUTH = 'AUTH.AUTH'
export const authorize = ({ dispatch }) => {
  api
    .check()
    .then(res => dispatch({
      type: AUTH,
      status: res.msg === 'true',
    }))
}
