<template>
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false"
    @select="handleSelect">
    <el-menu-item index="0">He'Blog</el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1">首页</el-menu-item>
    <el-menu-item index="2">归档</el-menu-item>
    <el-menu-item index="3">标签</el-menu-item>
    <el-menu-item index="4">关于我</el-menu-item>
  </el-menu>
</template>

<script setup>

import { ref } from 'vue'

const activeIndex = ref('1')
const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.el-menu-demo{
  width: 100%;
  padding: 0 100px;
}
.flex-grow {
  flex-grow: 1;
}
</style>