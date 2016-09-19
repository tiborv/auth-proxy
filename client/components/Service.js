import React, { Component } from 'react'
import ServiceForm from '../forms/service'


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

  toggleEdit() {
    this.setState({ edit: !this.state.edit })
    if (this.props.addNew) this.props.addNew()
  }

  delete(e) {
    e.preventDefault()
    this.props.del(this.props.service)
    this.setState({ edit: false })
  }

  render() {
    const { service, clients, addNew } = this.props
    return this.state.edit ? (
      <ServiceForm
        onSubmit={::this.save}
        formData={service}
        enum={clients.map(c => c.token)}
        enumNames={clients.map(c => c.name)}
        new={addNew}>

        <div className="btn-group" role="group">
          <button type="submit" className="btn btn-info">Save</button>
          <button className="btn btn-info">Cancel</button>

          { addNew ? (
              <div/>
            ) : (
              <button className="btn btn-danger" onClick={::this.delete}>Delete</button>
            )
          }
        </div>
      </ServiceForm>
    ) : (
      <button type="button" className="list-group-item" onClick={::this.toggleEdit}>
        {service.slug}
      </button>
    )
  }
}


export default Service
