import { connect } from 'react-redux'
import { list, del, create, update } from '../actions/service'
import ServiceList from '../components/ServiceList'

const mapStateToProps = (state, ownProps) => ({
  services: state.services,
  clients: state.clients
})

const mapDispatchToProps = (dispatch, ownProps) => ({
  create: (service) => {
    dispatch(create(service))
  },
  update: (service) => {
    dispatch(update(service))
  },
  del: (service) => {
    dispatch(del(service))
  }
})

const ClientContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ServiceList)

export default ClientContainer
