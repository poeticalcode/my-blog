import {createVNode, render} from "vue";
import CodeCopy from "@/components/markdown/components/codecopy/Index.vue";
import ClipboardJS from "clipboard";

const callback = (preList) => {
  preList.forEach(el => {
    // el.onclick = ((el) => {
    //   return () => {
    //     console.log(el.classList[0]);
    //     let clipboard = new ClipboardJS(`.${el.classList[0]}`, {
    //       text: () => {
    //         return el.innerText
    //       }
    //     })
    //     clipboard.on('success', function (e) {
    //       console.log("ok");
    //       clipboard.destroy() // 销毁,避免多次点击重复出现
    //     })
    //     clipboard.on('error', function (e) {
    //       console.log('复制失败')
    //     })
    //   }
    // })(el)
    if (el.classList.contains('code-copy-added')) return
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