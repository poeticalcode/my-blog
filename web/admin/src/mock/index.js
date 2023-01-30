import Mock from 'mockjs'

const Random = Mock.Random // 生成随机数


let Result = {
    code: 200,
    data: null,
    msg: "success"
}

// 文章列表
Mock.mock("/api/article/list", "get", () => {
    return Result
})