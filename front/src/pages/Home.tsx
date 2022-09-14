import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link} from 'react-router-dom';
import {getTypeList, getBlogList} from '../webAPI'

const Home = (props: {setBlogId: (blogId: string) => void}) =>{
    // const [page, setPage] = useState("")
    const [bloglist, setBlogList] = useState([])

    useEffect(() =>{
        getBlogList().then(data =>{
            setBlogList(data.data)
            localStorage.setItem('blogId', '')
            console.log(bloglist)
        })

    }, [])


    if (!bloglist || bloglist.length === 0){
        return (
            <div>
                <div>they're no any blog right now.</div>
                <Link to="/postblog">Post first</Link>
            </div>
        )
    }
    if (!localStorage.getItem('token')){
        return (
            <div className="list-group list-group-checkable d-grid gap-2 border-0 w-auto">
                {
                    bloglist.map(item=>
                        <label className="list-group-item rounded-3 py-3" key={item}>
                            <Link to="/blog" className="d-block small opacity-50" onClick={() =>props.setBlogId(item['Id'])}>{item['title']}</Link>
                            <nav>Click count: {item['clickhit']}</nav>
                        </label>
                    )
                }
            </div>

        )
    }
    else{
        return (
            // <div>
            <div className="list-group list-group-checkable d-grid gap-2 border-0 w-auto">
                <Link to="/postblog">Post</Link>
                {/* <ul> */}
                    {
                    bloglist.map(item=>
                        <label className="list-group-item rounded-3 py-3" key={item}>
                            <Link to="/blog" onClick={() =>props.setBlogId(item['Id'])}>{item['title']}</Link>
                            <nav>Click count: {item['clickhit']}</nav>
                        </label>
                    )
                    }
                {/* </ul> */}
            </div>
        )
    }
}


export default Home;