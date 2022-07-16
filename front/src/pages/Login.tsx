import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';
import {login} from '../webAPI'


const Login = (props:{setName:(name:string)=>void, setBloggerId:(bloggerId:any)=>void}) =>{
    const [username, setName] = useState("")
    const [password, setPassword] = useState("")
    const [msg, setMsg] = useState("")
    const [redirect, setRedirect] = useState(false)

    const submit = async(e:SyntheticEvent) =>{
        e.preventDefault();

        login(username, password).then(data =>{
            if (data['msg'] === ""){
                console.log(data)
                localStorage.setItem('token', data['token'])
                
                props.setName(username)
                props.setBloggerId(data['data']['userid'])
                setRedirect(true)
            }
            else{
                console.log(data)
                setMsg(data['msg'])
                // return data['msg']
            }
        })

    }

    if (redirect){
        // TODO: 用 reload 會沒辦法 nav 到 /
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
            
            {/* {msg && <h2>{msg}</h2>} 目前看起來這行跟下面那行好像沒有差別? */}
            <h2>{msg}</h2>

            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    )
}

export default Login;