import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link} from 'react-router-dom';
import { json } from 'stream/consumers';
import { resolveModuleNameFromCache } from 'typescript';
import {getTypeList, } from '../webAPI'


const Blog = (props: {blogId: any, bloggerId: any}) =>{
    const [blogId, setBlogId] = useState('')
    const [blogCreator, setBlogCreator] = useState('')
    const [blogTitle, setBlogTitle] = useState('')
    const [blogContent, setBlogContent] = useState('')
    const [blogClickhit, setBlogClickhit] = useState('')
    const [blogNextId, setBlogNextId] = useState('')
    const [blogPreviousId, setBlogPreviousId] = useState('')
    const [blogComment, setBlogComment] = useState([])

    const [comment, setUploadComment] = useState('')

    useEffect(() =>{
        getBlogContent(props.blogId).then(data =>{
            // setBlogList(data)
            // console.log(data.data.next.id)
            // console.log(data.data.previous.id)
            setBlogId(props.blogId)
            setBlogCreator(data.data.blog_content.blogger)
            setBlogTitle(data.data.blog_content.title)
            setBlogContent(data.data.blog_content.content)
            setBlogClickhit(data.data.blog_content.clickhit)

            if (data.data.next){
                setBlogNextId(data.data.next.id)
            }
            if (data.data.previous){
                setBlogPreviousId(data.data.previous.id)
            }

            setBlogComment(data.data.blog_comment)
        })

    }, [])


    const getBlogContent = async(blogId: any) =>{
        return await fetch(`http://localhost:8080/api/v1/blog?id=${blogId}`, {
            method: 'GET',
            headers: {'Content-Type': 'application/json'}
        }).then((res) => res.json())
    }
    const nextBlogContent= async(blogId: any) =>{
        getBlogContent(blogId).then(data =>{
            // setBlogList(data)
            // console.log(data.data.next.id)
            // console.log(data.data.previous.id)
            setBlogId(blogId)
            setBlogCreator(data.data.blog_content.blogger)
            setBlogTitle(data.data.blog_content.title)
            setBlogContent(data.data.blog_content.content)
            setBlogClickhit(data.data.blog_content.clickhit)
            setBlogComment(data.data.blog_comment)

            // 這裡用 ! 能判斷 null 為 false
            if (!data.data.next){
                console.log("im in condition")
                setBlogPreviousId(data.data.previous.id)
            }
            else{
                setBlogNextId(data.data.next.id)
                setBlogPreviousId(data.data.previous.id)
            }
        })
    }
    const previousBlogContent= async(blogId: any) =>{
        getBlogContent(blogId).then(data =>{
            // setBlogList(data)
            // console.log(data.data.next.id)
            // console.log(data.data.previous.id)
            setBlogId(blogId)
            setBlogCreator(data.data.blog_content.blogger)
            setBlogTitle(data.data.blog_content.title)
            setBlogContent(data.data.blog_content.content)
            setBlogClickhit(data.data.blog_content.clickhit)
            setBlogComment(data.data.blog_comment)
            // 這裡用 ! 能判斷 null 為 false
            if (!data.data.previous){
                console.log("im in condition")
                setBlogNextId(data.data.next.id)
            }
            else{
                setBlogNextId(data.data.next.id)
                setBlogPreviousId(data.data.previous.id)
            }
        })
    }
    const uploadComment = async(e: SyntheticEvent) =>{
        e.preventDefault()

        return await fetch('http://localhost:8080/api/v1/admin/blog/comment', {
            method: 'POST',
            headers: {'content-type': 'application/json', 'token': `${localStorage.getItem('token')}`},
            body: JSON.stringify({
                bloggerid: props.bloggerId,
                blogid: blogId,
                addtime: new Date(Date.now()),
                content: comment
            })
        })
    }

    return (
        <div>
            <div className="creator">Creator: {blogCreator}</div>
            <div className="title">Title: {blogTitle}</div>
            <div className="blogContent">Content: {blogContent}</div>
            <div className="clickhit">Click count: {blogClickhit}</div>
            {/* <Link to="/blog">Next</Link> */}
            {/* 目前卡在 無法點擊 next 刷新同一頁面 */}


            {/* TODO:  <button onClick={() => }>Reply</button> */}
            <br></br>
            <br></br>

            <div>{blogComment.length} comment</div>
            <form onSubmit={uploadComment}>
                <input placeholder="add comment" required onChange={e => setUploadComment(e.target.value)}></input><button>add</button>
            </form>
            <ul>
                {
                blogComment.map(item=>
                    <li key={item} className="comment">
                        <div className="comment_nickname">Nickname: {item['nickname']}</div>
                        <div className="comment_content">Content: {item['content']}</div>
                    </li>
                )
                }
            </ul>
            <button onClick={() => nextBlogContent(blogNextId)}>Next</button><button onClick={() => previousBlogContent(blogPreviousId)}>Previous</button>
            <br></br>

        </div>
    )
}




export default Blog;