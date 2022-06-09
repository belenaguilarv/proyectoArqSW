import React, {useContext} from "react";
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