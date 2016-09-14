import { connect } from 'react-redux'
import { list } from '../actions/token'
import TokenList from '../components/TokenList'

const mapStateToProps = (state, ownProps) => ({
  tokens: state.tokens
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  onClick: () => {
    dispatch(list())
  }
})

const TokenContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(TokenList)

export default TokenContainer
