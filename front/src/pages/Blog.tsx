import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link} from 'react-router-dom';
import { resolveModuleNameFromCache } from 'typescript';
import {getTypeList, } from '../webAPI'


const Blog = (props: {blogId: any}) =>{
    const [blogCreator, setBlogCreator] = useState('')
    const [blogTitle, setBlogTitle] = useState('')
    const [blogContent, setBlogContent] = useState('')
    const [blogClickhit, setBlogClickhit] = useState('')
    const [blogNextId, setBlogNextId] = useState('')
    const [blogPreviousId, setBlogPreviousId] = useState('')
    const [blogComment, setBlogComment] = useState([])

    useEffect(() =>{
        getBlogContent(props.blogId).then(data =>{
            // setBlogList(data)
            // console.log(data.data.next.id)
            // console.log(data.data.previous.id)
            setBlogCreator(data.data.blog_content.blogger)
            setBlogTitle(data.data.blog_content.title)
            setBlogContent(data.data.blog_content.content)
            setBlogClickhit(data.data.blog_content.clickhit)
            setBlogNextId(data.data.next.id)
            setBlogPreviousId(data.data.previous.id)

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


    return (
        <div>
            <div className="creator">Creator: {blogCreator}</div>
            <div className="title">Title: {blogTitle}</div>
            <div className="blogContent">Content: {blogContent}</div>
            <div className="clickhit">Click count: {blogClickhit}</div>
            {/* <Link to="/blog">Next</Link> */}
            {/* 目前卡在 無法點擊 next 刷新同一頁面 */}
            <br></br>
            <br></br>
            <br></br>
            <div>Comment:</div>
            <ol>
                {
                blogComment.map(item=>
                    <li key={item}>
                        <div>Comment title: {item['blogtitle']}</div>
                        <div>Comment content: {item['content']}</div>
                    </li>
                )
                }
            </ol>
            <button onClick={() => nextBlogContent(blogNextId)}>Next</button><button onClick={() => previousBlogContent(blogPreviousId)}>Previous</button>




        </div>
    )
}




export default Blog;