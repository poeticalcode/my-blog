import copy from "./copy.js";

import {watermark} from '@hewenyao/vue-directives'

// import watermark from "./watermark.js";
import target from "./target.js";
import auth from "./auth.js";

export default {
  install: (Vue, options = {}) => {
    [target, copy, watermark, auth].forEach(item => {
      Vue.directive(item.name, item.handle)
    })
  }
}