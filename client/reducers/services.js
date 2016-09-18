import { LIST_SUCESS, CREATE_SUCESS, UPDATE_SUCESS, DELETE_SUCESS } from '../actions/service'

const services = (state = [], action) => {
  switch (action.type) {
    case LIST_SUCESS:
      return action.services || state
    case CREATE_SUCESS:
      return [...state, action.service]
    case UPDATE_SUCESS:
      return [...state.map(s => {
        if (s.slug === action.service.slug) return action.service
        return s
      })]
    case DELETE_SUCESS:
      return [...state.filter(s => s.slug !== action.service.slug)]
    default:
      return state
  }
}

export default services
