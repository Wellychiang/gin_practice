import React, {useState} from 'react';
import {Link} from 'react-router-dom';

import {logout} from '../webAPI'

const Nav = (props:{username:any}) => {

    if (!localStorage.getItem('username')){
        localStorage.setItem('username', props.username)
    }
    

    const logout = async() =>{
        await fetch('http://localhost:8080/api/v1/logout', {
            method: 'POST',
            headers: {'Content-Type': 'application/json', 'token': `${localStorage.getItem('token')}`},
            // credentials: 'include', credentials 是用來用 cookie 的, 會在 response 裡的 set-cookie 取出(需要 server 設置)
        })
        localStorage.setItem('token', '')
        window.location.reload();
    }


    if (localStorage.getItem('token') != ''){
        return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
            <Link to="/" className="navbar-brand">Home</Link>

            <div>
                <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/member_center" className="nav-link" >{localStorage.getItem('username')}</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/login" className="nav-link" onClick={logout}>Logout</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/register" className="nav-link">Register</Link>
                </li>
                </ul>
            </div>
            </div>
        </nav>
        )
    } else {
        return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
            <Link to="/" className="navbar-brand">Home</Link>

            <div>
                <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link">Login</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/register" className="nav-link">Register</Link>
                </li>
                </ul>
            </div>
            </div>
        </nav>
        )
    }
}

export default Nav;