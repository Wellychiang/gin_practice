import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link} from 'react-router-dom';
import {getTypeList, getBlogList} from '../webAPI'

const Home = (props: {setBlogId: (blogId: string) => void}) =>{
    // const [page, setPage] = useState("")
    const [bloglist, setBlogList] = useState([])

    useEffect(() =>{
        getBlogList().then(data =>{
            setBlogList(data.data)
            console.log(bloglist)
        })

    }, [])


    if (!localStorage.getItem('token')){
        return (
            <div>
                not login page
                
                <ul>
                    {
                    bloglist.map(item=>
                        <li key={item}>
                            <Link to="/blog" onClick={() =>props.setBlogId(item['id'])}>{item['title']}</Link>
                            <nav>點擊次數: {item['clickhit']}</nav>
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
                <br></br>

                {/* TODO: 轉到 post blog page, 才 post */}
                <Link to="/postblog">Post</Link>
                <ul>
                    {
                    bloglist.map(item=>
                        <li key={item}>
                            <Link to="/blog" onClick={() =>props.setBlogId(item['id'])}>{item['title']}</Link>
                            <nav>點擊次數: {item['clickhit']}</nav>
                        </li>
                    )
                    }
                </ul>
            </div>
        )
    }
}


export default Home;