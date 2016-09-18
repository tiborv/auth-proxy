import { LIST_SUCESS, CREATE_SUCESS, UPDATE_SUCESS, DELETE_SUCESS } from '../actions/client'

const clients = (state = [], action) => {
  switch (action.type) {
    case LIST_SUCESS:
      return action.clients || state
    case CREATE_SUCESS:
      return [...state, action.client]
    case UPDATE_SUCESS:
      return [...state.map(c => {
        if (c.token === action.client.token) return action.client
        return c
      })]
    case DELETE_SUCESS:
      return [...state.filter(c => c.token !== action.client.token)]
    default:
      return state
  }
}

export default clients
