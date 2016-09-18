import React, { Component } from 'react'
import Form from 'react-jsonschema-form'

const ServiceSchema = {
  title: 'Service',
  type: 'object',
  required: ['url', 'host'],
  properties: {
    slug: { type: 'string', title: 'Slug' },
    url: { type: 'string', title: 'Url' },
    host: { type: 'string', title: 'Host' },
    scheme: { type: 'string', title: 'Scheme', default: 'http' },
    clients: {
      title: 'Clients',
      type: 'array',
      items: {
        type: 'string',
        properties: {
          description: {
            name: 'string',
            token: 'string'
          }
        }
      }
    }
  }
}

class Service extends Component {
  constructor(props) {
    super(props)
    this.state = {
      edit: this.props.edit || false
    }
  }

  save(form) {
    if (this.props.service) {
      this.props.update(form.formData)
      this.setState({ edit: false })
      return
    }
    this.props.create(form.formData)
    this.props.addNew()
  }

  toogleEdit() {
    this.setState({ edit: !this.state.edit })
    if (this.props.addNew) this.props.addNew()
  }

  delete() {
    this.props.del(this.props.service)
    this.setState({ edit: false })
  }

  render() {
    const { service, clients, addNew } = this.props
    const uiSchema = {
      slug: { 'ui:readonly': !addNew }
    }
    ServiceSchema.properties.clients.items.enum = clients.map(c => c.token)
    ServiceSchema.properties.clients.items.enumNames = clients.map(c => c.name)
    return this.state.edit ? (
      <div>
      <Form schema={ServiceSchema}
        uiSchema={uiSchema}
        onSubmit={::this.save}
        formData={service}/>
      <button className="btn btn-info" onClick={::this.toogleEdit}>Cancel</button>
      { addNew ? (
        <div></div>
      ) : (
        <button className="btn btn-danger" onClick={::this.delete}>Delete</button>
        )
      }
      </div>
    ) : (
      <div>
      Service {service.slug}
      <button className="btn btn-info" onClick={::this.toogleEdit}> Edit </button>
      </div>
    )
  }
}


export default Service
