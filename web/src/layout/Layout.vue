<template>
  <el-container style="height: 100%;">
    <el-header>
      <Header></Header>
    </el-header>

    <el-main>
      <el-config-provider :locale="locale">
        <router-view></router-view>
      </el-config-provider>
    </el-main>
  </el-container>
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
.el-header {
  $bgColor: white;
  border-bottom: solid 1px var(--el-menu-border-color);
  background-color: $bgColor;
  box-shadow: 0px 0px 8px 0px #c3c3c3;
}

.el-main {
  width: 100%;
  padding: 15px 0;
  margin-right: auto;
  margin-left: auto;
  overflow: hidden;
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