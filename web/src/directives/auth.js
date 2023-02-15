import { check } from "../util/auth.js";

export default {
    name: "auth",
    handle: {
        inserted(el, binding) {
            if (!check(binding.value)) { // 如果没有对应权限则移除
                el.parentNode && el.parentNode.remove()
            }
        }
    }
}