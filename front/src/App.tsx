import React, {useState} from 'react';
import './App.css';
import Login from './pages/Login'
import Home from './pages/Home'
import Register from './pages/Register'
import Nav from './components/Nav'
import { BrowserRouter, Route, Routes } from 'react-router-dom';



function App() {
  const [name, setName] = useState('');
  const token = localStorage.getItem('token');


  return (
    <div className="App">
      <BrowserRouter>
        <Nav/>

        <main className="form-signin">
            <Routes>
              <Route path="/" element={<Home/>} /> 
              <Route path="/login" element={<Login setName={setName}/>}/>
              <Route path="/register" element={<Register/>}/>
            </Routes>

            {/* <Route path="/">{Home}</Route> */}
            {/* <Route path="/login">{Login}</Route> */}
            {/* <Route path="/register">{Register}</Route> */}

        </main>

      </BrowserRouter>
    </div>
  );
}

export default App;
