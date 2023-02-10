import request from "@/axios";

const articleList = (pagingParam) => {
    return request({
        url: '/article/list',
        method: 'get',
        params: pagingParam
    })
}


export {articleList}

