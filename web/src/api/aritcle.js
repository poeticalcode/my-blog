import request from "@/util/request.js";

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


const addArticle = (data) => {
    return request({
        url: '/article',
        method: 'post',
        data: data,
        headers: {'content-type': 'application/json'},
    })
}

const updateArticle = (data) => {
    return request({
        url: '/article',
        method: 'put',
        data: data,
        headers: {'content-type': 'application/json'},
    })
}

export {fetchArticleList, fetchArticleDetail, addArticle, updateArticle}

