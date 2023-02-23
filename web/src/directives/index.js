import 'viewerjs/dist/viewer.css'
import VueViewer from 'v-viewer'

import watermark from "./watermark.js";
import target from "./target.js";
import auth from "./auth.js";

export default {
  install: (Vue, options = {}) => {
    Vue.use(VueViewer);
    [target, watermark, auth].forEach(item => {
      Vue.directive(item.name, item.handle)
    })
  }
}