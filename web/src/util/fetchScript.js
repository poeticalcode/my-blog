import asyncLoadJs from 'scriptjs'

const fetchScript = (url) => {
    return new Promise((resolve) => {
        asyncLoadJs(url, () => {
            resolve()
        })
    })
}

export default fetchScript
