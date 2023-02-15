// 修改被修饰元素内部的 a 标签的 target
export default {
  name: "target",
  handle: (el, binding) => {
    let target = binding.value || ""
    setTimeout(() => {
      let aList = el.getElementsByTagName("a")
      Array.from(aList).map(item => {
        item.target = target
      })
    }, 1000);
  }
} 