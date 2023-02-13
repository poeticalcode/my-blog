<template>
  <div class="nav">
    <div>
      <span>He'Blog</span>
    </div>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false"
      @select="handleSelect" :router="true">
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/archive">归档</el-menu-item>
      <el-menu-item index="/tags">标签</el-menu-item>
      <el-menu-item index="/about">关于我</el-menu-item>
    </el-menu>
  </div>
</template>

<script setup>

import { ref } from 'vue'

const activeIndex = ref('/')
const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.el-menu {
  border-bottom: unset;
}

.nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>