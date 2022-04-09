import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';
import {login} from '../webAPI'


const Login = () =>{
    const [username, setName] = useState("")
    const [password, setPassword] = useState("")
    const [redirect, setRedirect] = useState(false)

    const submit = async(e:SyntheticEvent) =>{
        e.preventDefault();

        login(username, password).then(data =>{
            if (data['msg'] === ""){
                console.log(data)
                localStorage.setItem('token', data['data'])
                setRedirect(true)
            }
            else{
                console.log(data)
                return data['msg']
            }
        })
        // const response = await fetch('http://localhost:8080/api/v1/login', {
        //     method: 'POST',
        //     headers: {'Content-Type': 'application/json'},
        //     // credentials: 'include',
        //     body: JSON.stringify({
        //         username: username,
        //         password: password
        //     })
        // }).then((res) => res.json())

    }

    if (redirect){
        return <Navigate to='/'/>
    }


    return (
        <form onSubmit={submit}>
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

            {/* <input type="email" className="form-control"  placeholder="name@example.com" required/> */}
            <input type="username" className="form-control"  placeholder="username" required
                onChange={e => setName(e.target.value)}
            />
            <input type="password" className="form-control"  placeholder="Password" required
                onChange={e => setPassword(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    )
}

export default Login;