<template>
  <div class="container-fluid" style="padding:0;">
    <Header style="margin-bottom: 18px;"></Header>
    <div class="container" style="margin-top: 1.8rem;">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>


import Header from "./components/Header.vue";

import { provide, reactive, } from "vue";

import zhCn from 'element-plus/lib/locale/lang/zh-cn'

let locale = zhCn

const layoutConfig = reactive({
  isCollapse: true
});
provide("layoutConfig", layoutConfig)

</script>

<style lang="scss" scoped>
.el-main {
  width: 100%;
  margin-right: auto;
  margin-left: auto;
  overflow: hidden;
  padding: 15px 5px;
}

// xs
@media only screen and (min-width: 576px) {
  .el-main {
    max-width: 540px;
  }
}

// sm
@media only screen and (min-width: 768px) {
  .el-main {
    max-width: 720px;
  }
}

// md
@media only screen and (min-width: 992px) {
  .el-main {
    max-width: 960px;
  }
}

// lg
@media only screen and (min-width: 1200px) {
  .el-main {
    max-width: 1140px;
  }
}

// xl
@media only screen and (min-width: 1920px) {
  .el-main {
    max-width: 1840;
  }
}
</style>