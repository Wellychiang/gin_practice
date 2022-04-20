import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';
import {login} from '../webAPI'


const Login = (props:{setName:(name:string)=>void}) =>{
    const [username, setName] = useState("")
    const [password, setPassword] = useState("")
    const [redirect, setRedirect] = useState(false)

    const submit = async(e:SyntheticEvent) =>{
        e.preventDefault();

        login(username, password).then(data =>{
            if (data['msg'] === ""){
                console.log(data)
                localStorage.setItem('token', data['data'])
                
                props.setName('qq')
                setRedirect(true)
            }
            else{
                console.log(data)
                return data['msg']
            }
        })

    }

    if (redirect){
        // TODO: 用 reload 會沒辦法 nav 到 /, 可是不用的話登入不會改變頂部欄位的顯示, 需要手動刷新
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