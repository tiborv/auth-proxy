import React, { Component } from 'react'
import Form from 'react-jsonschema-form'

const ClientSchema = {
  type: 'object',
  required: ['name'],
  properties: {
    name: { type: 'string', title: 'Name' },
    token: { type: 'string', title: 'Token' },
  }
}

const uiSchema = {
  token: { 'ui:readonly': true }
}

class ClientForm extends Component {
  render() {
    ClientSchema.title = this.props.new ? 'Create a new Client' : 'Edit Client'

    return (
      <Form schema={ClientSchema} uiSchema={uiSchema} {...this.props} />
    )
  }
}

export default ClientForm
