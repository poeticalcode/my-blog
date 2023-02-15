// 修改被修饰元素内部的 a 标签的 target
function install(Vue, options = {}) {
  Vue.directive(options.name || "target", (el, binding) => {
    let target = binding.value || ""
    setTimeout(() => {
      let aList = el.getElementsByTagName("a")
      Array.from(aList).map(item => {
        item.target = target
      })
    }, 1000);
  })
}
export default { install }