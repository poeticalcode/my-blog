<template>
  <el-container style="height: 100%">
    <el-aside>
      <Aside :collapse="layoutConfig.isCollapse"></Aside>
    </el-aside>

    <el-container>
      <el-header>
        <Header></Header>
      </el-header>

      <el-main>
        <el-config-provider :locale="locale">
          <router-view></router-view>
        </el-config-provider>
      </el-main>

    </el-container>
  </el-container>
</template>

<script setup>
import Aside from "./components/Aside.vue";
import Header from "./components/Header.vue";

import { provide, reactive, } from "vue";

import zhCn from 'element-plus/lib/locale/lang/zh-cn'
let locale = zhCn

const layoutConfig = reactive({
  isCollapse: true
});
provide("layoutConfig", layoutConfig)

</script>

<style >
html,
body,
#app {
  width: 100%;
  height: 100%;
  margin: 0;
}

.el-header {
  display: flex;
  align-items: center;
  border-bottom: 1px solid #f6f6f6;
  user-select: none;
}

.el-aside {
  width: auto;
}
</style>