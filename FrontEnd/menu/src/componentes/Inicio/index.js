import React, {useEffect, useState } from "react";
import { ProductoItem } from "../Productos/ProductoItem";
import swal from "sweetalert2";


async function GetProductByIdCategory(Id) {
    return fetch('http://localhost:8090/productsByCategory/' + Id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
   }

export const CategoryLista =()=>{   
    const [categorias,setCategorias] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/categories')
    .then((response) => response.json());
    setCategorias(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])

    const [productos,setProductos]=useState([]);
    async function Handle (Id) {
    const response = await GetProductByIdCategory(Id)
    if (response.status == 400) {
      swal.fire({
        icon: 'error',
        text: "No hay productos en esta categoria",
      })
   }else{
    setProductos(response)
    console.log(response);
   }
    };
const Render =(
    <div className="productos">
    {
            productos.map(producto =>(
              <ProductoItem 
              key={producto.id}
              product_id={producto.product_id}
              name={producto.name}
              product_unit_price={producto.product_unit_price}
              Category_id={producto.Category_id}
              stock={producto.stock}
              picture_url={producto.picture_url}
              description={producto.description}
              />  
            ))
    }
    </div> 
)
    return(
        <>
        <div className="categorias">
            {
                categorias.map(categoria =>(
                  <button className="category"onClick={()=>Handle(categoria.Id)}>
                    <span>{categoria.Name}</span>
                  </button>
                ))
            }
        </div>
          <div>
              {Render}
          </div>
        </>
    );
}
