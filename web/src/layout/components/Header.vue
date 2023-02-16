<template>
  <div class="nav">
    <div class="brand" @click="router.push('/')">
      <span>He'Blog</span>
    </div>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false" :router="true">
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/archive">归档</el-menu-item>
      <el-menu-item index="/tags">标签</el-menu-item>
      <el-menu-item index="/about">关于我</el-menu-item>
    </el-menu>
</div>
</template>

<script setup>

import { ref } from 'vue'
import { useRouter } from 'vue-router'
const activeIndex = ref('/')
const router = useRouter()
</script>

<style lang="scss" scoped>
.el-menu {
  border-bottom: unset;
}

.nav {
  display: flex;
  justify-content: space-between;
  align-items: center;

  .brand {
    cursor: pointer;
  }
}
</style>