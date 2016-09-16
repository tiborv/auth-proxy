import { connect } from 'react-redux'
import { list } from '../actions/client'
import ClientList from '../components/ClientList'

const mapStateToProps = (state, ownProps) => ({
  clients: state.clients
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  onClick: () => {
    dispatch(list())
  }
})

const ClientContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ClientList)

export default ClientContainer
