import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect from 'react-router-dom';
import {getTypeList, blogList} from '../webAPI'

const Home = () =>{
    // const [page, setPage] = useState("")
    const [size, setSize] = useState([])


    useEffect(() =>{
        blogList().then(da =>{
            console.log(da.data[0])
            setSize(da.data)
        })

    }, [])
    

    if (!localStorage.getItem('token')){
        return (
            <div>
                not login page
                
                <ul>{
                    size.map(item=>
                        <li key={item}>
                            <div>{item['title']}</div>
                        </li>
                    )
                    }
                </ul>
            </div>
        )
    }
    else{
        return (
            <div>
                login home page
                {size}
            </div>
        )
    }
}


export default Home;