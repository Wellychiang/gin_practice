import React, {SyntheticEvent, useState, useEffect} from 'react';
import {Navigate} from 'react-router-dom';
import {getTypeList, getBlogList} from '../webAPI'

const PostBlog = (props: {bloggerid: any}) =>{
    // const [page, setPage] = useState("")
    const [bloglist, setBlogList] = useState([])
    const [content, setContent] = useState('')
    const [title, setTitle] = useState('')
    const [redirect, setRedirect] = useState(false)

    // useEffect(() =>{
    //     getBlogList().then(data =>{
    //         setBlogList(data.data)
    //         console.log(bloglist)
    //     })

    // }, [])

    const submitt = async(e:SyntheticEvent) =>{
        e.preventDefault();
        console.log('bloggerid: ', props.bloggerid)

        await fetch('http://localhost:8080/api/v1/admin/blog', {
            method: 'POST',
            headers: {'content-type': 'application/json', 'token': `${localStorage.getItem('token')}`},
            body: JSON.stringify({
                title: title,
                content: content,
                bloggerid: props.bloggerid,
                addtime: new Date(Date.now()),
                updatetime: new Date(Date.now())
            })
        }).then( res => res.json()).then(data =>{
            if (data['msg'] === ''){
                console.log(data)
                setRedirect(true)
            }
        })
    }

    if (redirect){
        return <Navigate to="/"/>
    }


    return (
        // 現在到這, 還沒完成
        <form onSubmit={submitt}>
          <h1 className="h3 mb-3 fw-normal">blog info</h1>

            <input type="title" className="form-control"  placeholder="blog title" required
                onChange={e => setTitle(e.target.value)}
            />
            <input type="content" className="form-control"  placeholder="blog content" required
                onChange={e => setContent(e.target.value)}
            />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>
    )
}


export default PostBlog;