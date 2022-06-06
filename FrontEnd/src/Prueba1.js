
import React, { useState } from "react";
import logo from './logo.svg';
import './Prueba1.css';
//import './Menu.js';

/*async function getUserByID(id) {
 return fetch('http://127.0.0.1:8090/user/' + id, {
   method: 'GET',
   headers: {
     'Content-Type': 'application/json'
   }
 })

 .then((Response) => {
  if (Response.status == 200){
    console.log(Response)
    return Response.json()
  }
  else{
   throw `error with status ${Response.status}`;
  }
})
}*/

async function getUserByID(id) {
  return fetch('http://127.0.0.1:8090/user/' + id, {
    crossDomain:true,
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(data => data.json())

 }




function Prueba1() {

  const [userData, setUserData] = useState({});
  const [isUser, setIsUser] = useState(false);
  
  const handleSubmit = async event => {
    //Prevent page reload
    event.preventDefault();

    // Obtenemos Textos
    var {uid, upass}  = document.forms[0];

    // Find user login info
    try{
      const user = await getUserByID(uid.value);
    } catch(e){
      console.log(e);
    }
    
    setUserData(user);
    setIsUser(true);
    
  };

  
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Username </label>
          <input type="text" name="uid" required />
          <label> Password </label>
          <input type="password" name="upass" required />
        </div>
        <div className="button-container">
          <input type="submit" value="Ingresar" />
        </div>
      </form>
    </div>
  );

  return (
    <div className="app">
      <div className="login-form">
        <div className="title">THE HANDBAG STORE {userData.id}</div>
          {isUser? <div>Usuario: {userData.id} </div> : renderForm}
      </div>
    </div>
  );
}
    // llamar a la funcion de aca arriba
export default Prueba1;
