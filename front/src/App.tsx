import React, {useState} from 'react';
import './App.css';
import Login from './pages/Login'
import Home from './pages/Home'
import Blog from './pages/Blog'
import Register from './pages/Register'
import Nav from './components/Nav'
import { BrowserRouter, Route, Routes } from 'react-router-dom';



function App() {
  const [name, setName] = useState('');
  const [blogId, setBlogId] = useState('');


  return (
    <div className="App">
      <BrowserRouter>
        <Nav/>

        <main className="form-signin">
            <Routes>
              <Route path="/" element={<Home setBlogId={setBlogId}/>} /> 
              <Route path="/login" element={<Login setName={setName}/>}/>
              <Route path="/register" element={<Register/>}/>

              <Route path="/blog" element={<Blog blogId={blogId}/>}/>
            </Routes>

        </main>

      </BrowserRouter>
    </div>
  );
}

export default App;
