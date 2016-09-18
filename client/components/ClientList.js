import React, { Component } from 'react'
import Client from './Client'

class ClientList extends Component {
  constructor(props) {
    super(props)
    this.state = {
      addNew: false
    }
  }

  addNew() {
    this.setState({ addNew: true })
  }

  render() {
    const { addNew } = this.state
    const { clients, ...actions } = this.props
    return (
      <div>
        {clients.map((client, i) =>
          <Client
            key={i}
            client={client}
            {...actions }
          />
        )}
        {addNew ? (
          <Client
            edit={true}
            {...actions }
          />
        ) : (
          <button onClick={::this.addNew}>Create Client</button>
        )}
      </div>
    )
  }
}

export default ClientList
