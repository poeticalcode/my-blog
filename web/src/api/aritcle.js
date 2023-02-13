import request from "@/axios";

const fetchArticleList = (pagingParam) => {
    return request({
        url: '/article/list',
        method: 'get',
        params: pagingParam
    })
}


const fetchArticleDetail = (id) => {
    return request({
        url: '/article/' + id,
        method: 'get'
    })
}

export { fetchArticleList, fetchArticleDetail }

