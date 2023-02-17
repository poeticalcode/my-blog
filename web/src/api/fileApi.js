import request from "@/util/request.js";

/**
 * 上传图片
 * @param {*} formData 
 * @returns 
 */
export const uploadImg = (formData) => {
    return request({
        url: '/file/image',
        method: 'post',
        data: formData
    })
}