import request from "@/axios";

const articleList = () => {
    return request({
        url: '/article/list',
        method: 'get',
        params: {
            "page_num": 1,
            "page_size": 1
        }
    })
}


export {articleList}

