import request from "@/axios";

const fetchArticleList = (pagingParam) => {
    return request({
        url: '/article/list',
        method: 'get',
        params: pagingParam
    })
}


export {fetchArticleList}

