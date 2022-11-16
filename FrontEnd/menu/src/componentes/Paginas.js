import React from "react";
import { Routes, Route } from "react-router-dom";
import { CategoryLista } from "./Inicio";
import { ProductosLista } from "./Productos/index"
import GetUserByLogin from "./Login/Prueba1";


export const Paginas = () => {
    return (
        <section>
            <Routes>
                <Route path="/" exact element={<GetUserByLogin/>} />
                <Route path="/inicio" exact element={<CategoryLista/>} />
                <Route path="/productos" exact element={<ProductosLista/>} />
            </Routes>
        </section>
    )
}