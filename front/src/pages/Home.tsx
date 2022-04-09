import React, {SyntheticEvent, useState} from 'react';
import Redirect from 'react-router-dom';
import {getTypeList} from '../webAPI'

const Home = () =>{
    // const [page, setPage] = useState("")
    // const [size, setSize] = useState("")

    const submit = async() =>{
        const typedata = getTypeList('1', '25')
        console.log(typedata)
        return typedata
    }

    if (!localStorage.getItem('token')){
        return (
            <div>
                not login home page
            </div>
        )
    }
    else{
        return (
            <div onSubmit={submit}>
                login home page
            </div>
        )
    }
}


export default Home;