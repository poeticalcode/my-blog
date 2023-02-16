<template>

  <el-container style="height: 100%;">
    <el-header>
      <Header></Header>
    </el-header>
    <el-main>
      <div class="main-inner">
        <el-config-provider :locale="locale">
          <router-view></router-view>
        </el-config-provider>
      </div>
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

<style lang="scss">
.el-header {
  $bgColor: white;
  border-bottom: solid 1px var(--el-menu-border-color);
  background-color: $bgColor;
  box-shadow: 0px 0px 8px 0px #c3c3c3;
  padding: 0 100px;

  .el-menu {
    border-bottom: unset;
    background-color: $bgColor;
  }
}

.el-main {
  padding: 20px 0;
}

@media only screen and (min-width: 960) {
  .main-inner {
    padding: 0 96px;
  }

  .el-card {
    padding: 0 !important;
  }
}



@media only screen and (max-width: 1920px) {
  .main-inner {
    width: 1100px;
    margin: 0 auto;
  }

  .el-header {
    padding: 0px 250px;
  }

}


/* 在超过 1920px 的屏幕上怎么显示 */
@media only screen and (min-width: 1920px) {
  .main-inner {
    width: 1100px;
    margin: 0 auto;
  }

  .el-header {
    padding: 0px 250px;
  }
}
</style>