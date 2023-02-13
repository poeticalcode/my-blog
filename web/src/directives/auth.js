import {check} from "../util/auth.js";

function install(Vue, options = {}) {
    Vue.directive(options.name || "auth", {
        inserted(el, binding) {
            if (!check(binding.value)) { // 如果没有对应权限则移除
                el.parentNode && el.parentNode.remove()
            }
        }
    })
}

export default {install}