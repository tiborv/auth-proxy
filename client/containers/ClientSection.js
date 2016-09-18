import { connect } from 'react-redux'
import { list, del, create, update } from '../actions/client'
import ClientList from '../components/ClientList'

const mapStateToProps = (state, ownProps) => ({
  clients: state.clients
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  create: (client) => {
    dispatch(create(client))
  },
  update: (client) => {
    dispatch(update(client))
  },
  del: (client) => {
    dispatch(del(client))
  }
})

const ClientContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ClientList)

export default ClientContainer
