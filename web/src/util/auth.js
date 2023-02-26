// 返回当前权限
export function getCurrentAuthority() {
    return ["console"]
}

export function check(authority) {
    const currAuth = getCurrentAuthority()
    return currAuth.some(item => authority.includes(item))
}

// 是否登录
export function isLogin() {
    const current = getCurrentAuthority()
    return current && current[0] !== "guest"
}