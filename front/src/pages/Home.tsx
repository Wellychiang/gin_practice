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


    if (!bloglist){
        return (
            <div>they're no any blog right now.</div>
        )
    }
    if (!localStorage.getItem('token')){
        return (
            <div className="list-group list-group-checkable d-grid gap-2 border-0 w-auto">
                <ul>
                    {
                        bloglist.map(item=>
                            <label className="list-group-item rounded-3 py-3" key={item}>
                                <Link to="/blog" className="d-block small opacity-50" onClick={() =>props.setBlogId(item['id'])}>{item['title']}</Link>
                                <nav>Click count: {item['clickhit']}</nav>
                            </label>
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
                            <nav>Click count: {item['clickhit']}</nav>
                        </li>
                    )
                    }
                </ul>
            </div>
        )
    }
}


export default Home;