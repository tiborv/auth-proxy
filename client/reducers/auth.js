import { NOT_AUTH, AUTH } from '../actions/auth'


const auth = (state = true, action) => {
  switch (action.type) {
    case NOT_AUTH:
      return false
    case AUTH:
      return action.status
    default:
      return state
  }
}
export default auth
