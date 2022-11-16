/*import React, {useContext} from "react";
import IMG from "../../images/img01.jpg";


export const ProductoItem = ({title, image, category, price, id}) => {
  return (
    <div className="producto">
    <a href="#">
        <div className="producto__img">
            <img src={image} alt={title}></img>
            </div>
    </a>
    <div className="producto__footer">
        <h1> {title}</h1> 
        <p>{category}</p>
        <p className="price">${price}</p>
    </div>
    <div className="buttom">
        <button className="btn">Añadir al carrito</button>
        <div>
         <a href="#" className="btn">Vista</a>
        </div>
    </div>
    </div>
  )
};
*/
import React from "react";
//import {addToCart} from "../cookies/add2car"
import Cookies from "universal-cookie";
import swal from "sweetalert";

function notLogin(){
    swal("Debe loguearse")
    window.location.replace("/login")
}
const Cookie = new Cookies();
let cookie = Cookie.get("user")
    let id_user;
    if(cookie!=undefined){
    let array = cookie.split(",")
     id_user = array[0]
    }
    else{
         id_user = "undefined"
    }
export const ProductoItem =(
    {
    product_id,
    name,
    picture_url,
    product_unit_price,
    stock,
    description,
    Category_id
})=>{


    return(
        <div className="producto">
        <a href="#">
        <div className="producto_img">
            <img className="image" src={picture_url} alt=""/>
        </div>
        </a>
        <div className="producto_footer">
            <h1>{name}</h1>
            {stock>0?<p>Stock: {stock}</p>:
            <p>Sin Stock</p>}
            
            <p className="price">$ {product_unit_price}</p>
        </div>
        {/*
        <div className="buttom">
        {stock>0?
            id_user!="undefined"?
            <button className="btn" onClick={()=> addToCart(id)}>
                Agregar
            </button>:
            <button className="btn" onClick={()=>notLogin()}>
            Agregar
        </button>:
            <p>Este Producto no esta Disponible</p>}
        </div>
        */}
        </div>
    )
}