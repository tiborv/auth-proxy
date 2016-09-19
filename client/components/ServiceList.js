import React, { Component } from 'react'
import Service from './Service'

class ServiceList extends Component {
  constructor(props) {
    super(props)
    this.state = {
      newService: false
    }
  }

  toggleNew() {
    this.setState({ newService: !this.state.newService })
  }

  render() {
    const { newService } = this.state
    const { services, ...props } = this.props
    return (
      <div>
        <div className="list-group">
        {services.map((service, i) =>
          <Service
            key={i}
            service={service}
            {...props }
          />
        )}
      </div>

      {newService ? (
        <Service
          edit={true}
          addNew={::this.toggleNew}
          {...props }
        />
      ) : (
        <button className="btn btn-info" onClick={::this.toggleNew}>Create Service</button>
      )}
      </div>
    )
  }
}

export default ServiceList
