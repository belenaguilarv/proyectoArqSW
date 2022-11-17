import React, { useContext, useEffect, useState} from "react";

import Cookies from "universal-cookie";
import swal from "sweetalert2";
const Cookie = new Cookies();
const id_user = setUser()

function setUser (){
  let cookieUser = Cookie.get("username")

  if(cookieUser!=undefined){
  let array = cookieUser.split(",")
  return array[0]
  }else{
    return "undefined"
  }
}
function vaciarCarrito(){
   
    document.cookie = "cart"+id_user+"=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}
export const CrearOrden = ()=>{
  

    let cookie = Cookie.get("username")
    let id_user = parseInt(cookie.split(",")[0]);
    let OrderDetail ={'product_id':0,'quantity':0}
    let details = [];

    let a = Cookie.get("cart"+id_user).split(";")
 

    for (let i = 0; i < a.length; i++){
      let item = a[i];
      if(item != ""){
        let array = item.split(",")
        OrderDetail ={'product_id':0,'quantity':0}
         OrderDetail.product_id = parseInt(array[0])
         OrderDetail.quantity =parseInt(array[1])
         
        details.push(OrderDetail)
        
      }
    }
    
    
    const crearOrden = async()=>{
      fetch('http://localhost:8090/neworder',requestOptions)
      .then(response => {if (response.status != 201) {
        swal.fire({icon: 'error', text:"No se pudo realizar la compra"}
        ).then((result) => {
          if (result.isConfirmed) {
            window.location.reload();
            return response.json()
          }}) 
       }else{
          swal.fire({icon: 'success', text:"Compra realizada con éxito"}
          ).then((result) => {
            if (result.isConfirmed) {
              window.location.replace("/")
              vaciarCarrito()
              return response.json()
            }})
         
        }})
      };
      const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({
             id_user: user_id,
             details: details
    })
    };
     
    crearOrden();

}