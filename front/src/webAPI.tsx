
export const login = async(username: string, password: string) =>{
    return await fetch('http://localhost:8080/api/v1/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            // credentials: 'include', credentials 是用來用 cookie 的, 會在 response 裡的 set-cookie 取出(需要 server 設置)
            body: JSON.stringify({
                username: username,
                password: password
            })
    }).then((res) => res.json())
}

export const getTypeList = async(page: string, size: string) =>{
    return await fetch('http://localhost:8080/api/v1/admin/type/list', {
            method: 'GET',
            headers: {'Content-Type': 'application/json', 'token': `${localStorage.getItem('token')}`},
            body: JSON.stringify({
                page: page,
                size: size
            })
    }).then((res) => res.json())
}