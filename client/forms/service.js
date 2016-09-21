import React, { Component } from 'react'
import Form from 'react-jsonschema-form'

const ServiceSchema = {
  type: 'object',
  required: ['url', 'host'],
  properties: {
    slug: { type: 'string', title: 'Slug' },
    url: { type: 'string', title: 'Url' },
    host: { type: 'string', title: 'Host' },
    scheme: {
      type: 'string',
      title: 'Scheme',
      default: 'http',
      enum: ['http', 'https'],
    },
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


class ServiceForm extends Component {
  render() {
    ServiceSchema.properties.clients.items.enum = this.props.enum
    ServiceSchema.properties.clients.items.enumNames = this.props.enumNames
    ServiceSchema.title = this.props.new ? 'Create a new Service' : 'Edit Service'
    const uiSchema = {
      slug: { 'ui:readonly': !this.props.new }
    }
    return (
      <Form schema={ServiceSchema} uiSchema={uiSchema} {...this.props} />
    )
  }
}


export default ServiceForm
