import React from "react";

export const OrdenItem =(
    {order_id,
        date,
        total_price,
        user_id,
        OrdersDetalle   //Hace falta poner esto para el nuestro?
})=>{
    return(
      <div className="ordenes">
        <div className="orden">
          <table>
         <thead>
           <tr>
             <th>Total Price</th>
             <th>Date</th>
           </tr>
         </thead>
         <tbody>
            <tr>
               <td>{total_price}</td>
               <td>{date}</td>
             </tr>
         </tbody>
       </table>
       <h2>Detalle de la compra</h2>
       <table>
        <thead>
           <tr>
             <th>Producto</th>
             <th>Precio Unitario</th>
             <th>Cantidad</th>
             <th>Total</th>
           </tr>
         </thead>
         <tbody>
         {
           OrdersDetalle.map((ordenDetalle) =>(      //Hace falta poner esto para el nuestro?
                <tr>
                <td>{ordenDetalle.nombre}</td>
                <td>{ordenDetalle.precio_unitario}</td>
                <td>{ordenDetalle.cantidad}</td>
                <td>{ordenDetalle.total}</td>
              </tr>
                ))
           }
         </tbody>
       </table>
       <br></br>
        </div>
        <div>
        </div>
        </div>
    )
}