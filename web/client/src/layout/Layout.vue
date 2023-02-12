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

<style scoped>
.el-header {
  padding: unset;
}

.el-main{
  overflow: unset;
}
.main-inner{
  padding: 0 96px;
}
</style>