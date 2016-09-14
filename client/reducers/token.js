import { LIST_SUCESS } from '../actions/token'

const tokens = (state = [], action) => {
  switch (action.type) {
    case LIST_SUCESS:
      return action.tokens
    default:
      return state
  }
}
export default tokens
