import React, {SyntheticEvent, useState, useEffect} from 'react';
import Redirect, {Link, Navigate} from 'react-router-dom';
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

    const [redirect, setRedirect] = useState(false)

    useEffect(() =>{

        var storageBlogId = Number(localStorage.getItem('blogId'))
        var storageBloggerId = Number(localStorage.getItem('bloggerId'))
        if (!storageBlogId){
            storageBlogId = props.blogId
        }
        if (!storageBloggerId){
            storageBloggerId = props.bloggerId
        } 

        getBlogContent(storageBlogId).then(data =>{
            localStorage.setItem('blogId', String(storageBlogId))
            localStorage.setItem('bloggerId', String(storageBloggerId))
            setBlogId(`${storageBlogId}`)
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

            localStorage.setItem('blogId', String(blogId))

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
            localStorage.setItem('blogId', String(blogId))

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

        console.log('im trying')
        if (!localStorage.getItem('token')){
            setRedirect(true)
        }
        else{
            await fetch('http://localhost:8080/api/v1/admin/blog/comment', {
                method: 'POST',
                headers: {'content-type': 'application/json', 'token': `${localStorage.getItem('token')}`},
                body: JSON.stringify({
                    bloggerid: Number(localStorage.getItem('bloggerId')),
                    blogid: Number(blogId),
                    addtime: new Date(Date.now()),
                    content: comment
                })
            })
        window.location.reload()
        }
    }

    if (redirect){
        return <Navigate to='/login'/>
    }

    return (



        <div>

          <div className="blog-post">
            <h2 className="blog-post-title">{blogTitle}</h2>
            <p className="blog-post-meta">by <a href="#">{blogCreator}</a></p>

            <p>{blogContent}</p>

        </div>
            <br></br>

            <div>{blogComment.length} comment</div>
            <form onSubmit={uploadComment}>
                <input placeholder="add comment" required onChange={e => setUploadComment(e.target.value)}></input><button>add</button>
            </form>
            <ul>
                {
                blogComment.map(item=>

                    <a key={item}>
                        <div className="comment_nickname"> <a href="#">{item['nickname']}</a></div>
                        <div className="comment_content">{item['content']}</div>
                        <hr></hr>
                    </a>
                )
                }
            </ul>
            <button onClick={() => nextBlogContent(blogNextId)}>Next</button><button onClick={() => previousBlogContent(blogPreviousId)}>Previous</button>
            <br></br>

        </div>
    )
}




export default Blog;