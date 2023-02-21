import copy from "./copy.js";

// import {watermark} from '@hewenyao/vue-directives'

/* 图片预览指令 */
import 'viewerjs/dist/viewer.css'
import VueViewer from 'v-viewer'

import watermark from "./watermark.js";
import target from "./target.js";
import auth from "./auth.js";

export default {
  install: (Vue, options = {}) => {
    Vue.use(VueViewer);
    [target, copy, watermark, auth].forEach(item => {
      Vue.directive(item.name, item.handle)
    })
  }
}