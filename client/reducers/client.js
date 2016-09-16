import { LIST_SUCESS } from '../actions/client'

const clients = (state = [], action) => {
  switch (action.type) {
    case LIST_SUCESS:
      return action.clients
    default:
      return state
  }
}
export default clients
