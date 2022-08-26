import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link, Navigate} from 'react-router-dom';
import { json } from 'stream/consumers';

const MemberCenter = (props:{username: string}) =>{
    const [username, setUserName] = useState('')
    const [nickName, setNickName] = useState('')
    const [profile, setProfile] = useState('')
    const [sign, setSign] = useState('')

    const [redirect, setRedirect] = useState(false)


    useEffect(() =>{
        getBloggerInfo(localStorage.getItem('username')).then(data =>{
            console.log(data)
            setUserName(data.data.username)
            setNickName(data.data.nickname)
            setProfile(data.data.profile)
            setSign(data.data.sign)
        })

    }, [])

    const getBloggerInfo = async(username: any) =>{
    return await fetch(`http://localhost:8080/api/v1/admin/blogger?username=${username}`, {
        method: 'GET',
        headers: {'Content-Type': 'application/json', 'token': `${localStorage.getItem('token')}`},
    }).then(res => res.json())
    }

    const submit = async(e: SyntheticEvent) =>{
        e.preventDefault()
        changeUserInfo()
    }

    const changeUserInfo = async() =>{
        return await fetch('http://localhost:8080/api/v1/admin/info', {
            method: 'PUT',
            headers: {'content-type': 'application/json', 'token': `${localStorage.getItem('token')}`},
            body: JSON.stringify({
                username: username,
                nickname: nickName,
                profile: profile,
                sign: sign
            })
        }).then((res) => res.json()).then(data =>{
            console.log(data)
            setRedirect(true)
        })
    }

    if (redirect){
        return <Navigate to='/'/>
    }

    return (
        <form onSubmit={submit}>

          <h1 className="h3 mb-3 fw-normal">User's info</h1>

            {/* <input type="email" className="form-control"  placeholder="name@example.com" required/> */}
            <div>Username: </div>
            <div className="form-control"  placeholder="username">{username}</div>
            <br></br>

            <div>Nickname: </div>
            <input type="nickname" className="form-control"  placeholder="Nickname" required
                onChange={e => setNickName(e.target.value)} value={nickName}
            />
            <br></br>

            <div>Profile: </div>
            <input type="profile" className="form-control"  placeholder="Profile" required
                onChange={e => setProfile(e.target.value)} value={profile}
            />
            <br></br>

            <div>Sign: </div>
            <input type="sign" className="form-control"  placeholder="Sign" required
                onChange={e => setSign(e.target.value)} value={sign}
            />

            <br></br>
            <button className="w-100 btn btn-lg btn-primary" type="submit">Change info</button>

        </form>
        // <div>
        //     <div>Username: {username}</div>
        //     <div>Nick name: {nickName}</div>
        //     <div>Profile: {profile}</div>
        //     <div>Sign: {sign}</div>
        // </div>
    )
}

export default MemberCenter;