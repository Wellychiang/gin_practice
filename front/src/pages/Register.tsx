import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';

const Register = () =>{
    const [username, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [msg, setMsg] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) =>{
        e.preventDefault();
        
        await fetch('http://localhost:8080/api/v1/register',{
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                username,
                email,
                password
            })
        }).then((res) => res.json()).then(data =>{
            console.log(data)
            if (data['code'] === 1001){
                setMsg(data['msg'])
                return
            }
            else{
                setRedirect(true);
            }
        })
    }
    if (redirect){
        return <Navigate to='/login'/>
    }

    return (
        <form onSubmit={submit}>
            <a>Username</a>
            <input type="name" className="form-control"  placeholder="input username" required
                onChange={e => setName(e.target.value)}/>
            <br></br>
            <a>Password</a>
            <input type="password" className="form-control"  placeholder="input password" required
                onChange={e => setPassword(e.target.value)}/>
            <br></br>
            <a>Email</a>
            <input type="email" className="form-control"  placeholder="name@example.com" required
                onChange={e => setEmail(e.target.value)}/>
            <h2>{msg}</h2>

            <br></br>
            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    )
}


export default Register;