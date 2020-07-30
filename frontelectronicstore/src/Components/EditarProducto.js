import React from 'react';

import {
  Link,
  withRouter
} from "react-router-dom";

class EditarProductoComponent extends React.Component{

  constructor(props) {
      super(props);
      this.state = {
        done : false,
        producto: { 
          nombre: "",
          precio: 0,
          caracteristicas: ""
        }
      };

      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
  }
  
  componentDidMount() {
        const codigo = this.props.match.params.codigo;
        fetch("http://localhost:8081/productos/"+codigo, {
          method: "GET"
        })
        .then(result=>result.json())
        .then(resultado=>this.setState({
            done: true,
            producto: resultado
         }))
        .catch(() => {
            this.setState({
                done: true,
                success: false
            })
        })
    }

  handleChange(event) {
    const target = event.target;
    const name = target.name;
    const value = name != "precio" ? target.value : parseInt(target.value);
    this.setState({
      producto: {
        [name]: value
      }
    });
  }

  handleSubmit(event) {
    event.preventDefault();
    const codigo = this.props.match.params.codigo;
    fetch("http://localhost:8081/productos/"+codigo,{
            method: "PUT",
            body: JSON.stringify(this.state.producto)
        })
        .then(result=>result.json())
        .then(items=>this.setState({
            done: true,
            items
         }))
        .catch(() => {
            this.setState({
                done: true,
                success: false
            })
        })
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="col-lg-8" id="form-producto">
        <label>
          Nombre
    </label>
    <input type="text" name="nombre"  value={this.state.producto.nombre} onChange={this.handleChange} className="form-control" required />
    <label>
          Precio
    </label>
    <input type="text" name="precio"  value={this.state.producto.precio} onChange={this.handleChange} className="form-control" required />
    <label>
          caracteristicas:
          
        </label>
        <textarea onChange={this.handleChange} value={this.state.producto.caracteristicas} name="caracteristicas" />
        <Link to="/" className="btn btn-md btn-primary">Volver</Link>
        <input type="submit" value="Actualizar Producto" className="btn btn-success"/>
      </form>
    );
  }
}

export default withRouter(EditarProductoComponent)