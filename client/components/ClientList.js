import React, { Component } from 'react'
import Client from './Client'

class ClientList extends Component {
  constructor(props) {
    super(props)
    this.state = {
      newClient: false
    }
  }

  toggleNew() {
    this.setState({ newClient: !this.state.newClient })
  }

  render() {
    const { newClient } = this.state
    const { clients, ...actions } = this.props
    return (
      <div>
      <div className="list-group">
        {clients.sort((a, b) => b.stats.count - a.stats.count).map((client, i) =>
          <Client
            key={i}
            client={client}
            {...actions }
          />
        )}
      </div>

      {newClient ? (
        <Client
          edit={true}
          addNew={::this.toggleNew}
          {...actions }
        />
      ) : (
        <button className="btn btn-info" onClick={::this.toggleNew}>Create Client</button>
      )}
      </div>
    )
  }
}
export default ClientList
