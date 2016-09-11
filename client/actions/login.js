import * as api from '../utils/authAPI'

const LOGIN_SUCESS = 'LOGIN_SUCESS'
const LOGIN_FAILED = 'LOGIN_FAILED'

export const login = (username, password) => (dispatch) => {
  api
    .login(username, password)
    .then(() => dispatch({
      type: LOGIN_SUCESS,
    }))
    .catch(() => dispatch({
      type: LOGIN_FAILED,
    }))
}
