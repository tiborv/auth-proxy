const auth = (state = {}, action) => {
  switch (action.type) {
    case 'LOGIN_SUCESS':
      return {
        loggedIn: true
      }
    case 'LOGIN_FAILED':
      return {
        loggedIn: false
      }
    default:
      return state
  }
}
export default auth
