import React, {useState} from 'react';
import './App.css';
import Login from './pages/Login'
import Home from './pages/Home'
import Blog from './pages/Blog'
import PostBlog from './pages/PostBlog'
import Register from './pages/Register'
import MemberCenter from './pages/memberCenter'
import Nav from './components/Nav'
import { BrowserRouter, Route, Routes } from 'react-router-dom';



function App() {
  const [name, setName] = useState('');
  const [bloggerId, setBloggerId] = useState('');
  const [blogId, setBlogId] = useState('');


  return (
    <div className="App">
      <BrowserRouter>
        <Nav username={name}/>

        <main className="form-signin">
            <Routes>
              <Route path="/" element={<Home setBlogId={setBlogId}/>} /> 
              <Route path="/login" element={<Login setName={setName} setBloggerId={setBloggerId}/>}/>
              <Route path="/register" element={<Register/>}/>

              <Route path="/blog" element={<Blog blogId={blogId} bloggerId={bloggerId}/>}/>
              <Route path="/postblog" element={<PostBlog />}/>
              <Route path="/member_center" element={<MemberCenter username={name}/>}/>
            </Routes>

        </main>

      </BrowserRouter>
    </div>
  );
}

export default App;
