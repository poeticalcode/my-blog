import request from "@/axios";

const articleList = () => {
    return request({
        url: '/article/list',
        method: 'get',
        data: {
            "page_num": 1,
            "page_size": 100
        }
    })
}


export {articleList}

