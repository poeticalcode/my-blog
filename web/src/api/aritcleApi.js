import request from "@/util/request.js";

/**
 * 获取文章列表
 * @param {*} pagingParam
 * @returns
 */
export const fetchArticleList = (pagingParam) => {
  return request({
    url: '/article/list',
    method: 'get',
    params: pagingParam
  })
}


/**
 * 获取文章详情
 * @param {*} id
 * @returns
 */
export const fetchArticleDetail = (id) => {
  return request({
    url: '/article/' + id,
    method: 'get'
  })
}

/**
 * 添加文章
 * @param {*} data
 * @returns
 */
export const addArticle = (data) => {
  return request({
    url: '/article',
    method: 'post',
    data: data,
    headers: {'content-type': 'application/json'},
  })
}

/**
 * 更新文章
 * @param {*} data
 * @returns
 */
export const updateArticle = (data) => {
  return request({
    url: '/article',
    method: 'put',
    data: data,
    headers: {'content-type': 'application/json'},
  })
}
/**
 * 删除文章
 * @param {*} id
 * @returns
 */
export const deleteArticle = (id) => {
  return request({
    url: '/article/' + id,
    method: 'delete'
  })
}


/**
 * 获取每日发布文章的数据
 * @returns {*}
 */
export const fetchDailyArticleCount = () => {
  return request({
    url: '/article/count/daily-public',
    method: 'get'
  })
}