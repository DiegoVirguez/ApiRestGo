import React, { useRef, useState } from 'react';

import {
  Link,
  useRouteMatch,
  useParams
} from "react-router-dom";
import {Modal, ModalHeader, ModalBody, ModalFooter, Button} from 'reactstrap';

class ProductoComponent extends React.Component{
     constructor() {
        super();
        this.state = { 
            done: false,
            modalOpen : false,
            search: "",
            items: []
        };
    }
    

	componentDidMount() {
        fetch('http://localhost:8081/productos',{
        	method: "GET"
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
    
    eliminarProducto(codigo){
        fetch('http://localhost:8081/productos/'+codigo,{
          method: "DELETE"
        })
        .then(result=>result.json())
        .then(items=>this.setState({
            done: true,
         }))
        .catch(() => {
            this.setState({
                done: true,
                success: false
            })
        });
        this.componentDidMount();
    }
  
  onChange = e => {
    this.setState({ 
      search: e.target.value
    });
  };

	render(){
    const productos =  this.state.items.filter(producto => {
                return producto.nombre.toLowerCase().indexOf(this.state.search.toLowerCase()) !== -1;
            }); 
		return(    
              <div className="row" >
              <div className="col-lg-6"></div>
              <div className="col-lg-6 search_producto" >
                 <i className="fas fa-search" aria-hidden="true"></i>
                 <input className="form-control form-control-sm ml-3 w-75" type="text" placeholder="Search"
                    onChange={this.onChange} aria-label="Search" />
               </div>
              {this.state.items.length ? productos.map(producto => <div key={producto.codigo} className="col-lg-3">
			  <div className="thumbnail">
				<div className="caption">
					<h3>{producto.nombre}</h3>
					<p>$ {producto.precio} </p>
					<  div>
            <Link to={`/verProducto/${producto.codigo}`} className="btn btn-md btn-primary">Ver</Link>
						<Link to={`/editarProducto/${producto.codigo}`} className="btn btn-md btn-warning">Editar</Link>
						<a className="btn btn-md btn-danger" onClick={() => this.eliminarProducto(producto.codigo)}>Borrar</a>
					</div>
				</div>
			   </div>

			   </div>) :  <p>No hay productos registrados </p> }
		    </div>
			);
	}
}

export default ProductoComponent