import {createVNode, render} from "vue";
import CodeCopy from "@/components/markdown/components/codecopy/Index.vue";

const callback = (preList) => {
  preList.forEach(el => {
    if (el.classList.contains('code-copy-added')) return
    //   https://cn.vuejs.org/v2/api/index.html#Vue-extend
    /* 使用基础 Vue 构造器，创建一个“子类”。参数是一个包含组件选项的对象 */
    let copy = createVNode(CodeCopy, {
      code: el.innerText
    })
    copy.parent = el
    el.style.position = "relative"
    let mountNode = document.createElement("div");
    render(copy, mountNode)
    el.classList.add('code-copy-added')
    el.appendChild(copy.el)
  })
}

export default {
  name: "code-copy",
  handle: (el, bind) => {
    let preList = []
    let flag = false
    let count = 0
    let tim = setInterval(() => {
      if (flag) {
        clearInterval(tim)
        callback(preList)
      }
      preList = el.querySelectorAll('pre')
      flag = preList.length > 0 || count++ > 40
    }, 500)
  }
};