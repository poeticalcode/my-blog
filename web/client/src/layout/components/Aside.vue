<template>
  <el-menu
      default-active="2"
      :router="true"
      background-color="rgb(58, 118, 191)"
      text-color="#fff"
      active-text-color="#ffd04b"
      :collapse="layoutConfig.isCollapse"
  >

    <el-menu-item index="/">
      <el-icon>
        <HomeFilled/>
      </el-icon>
      <template #title>首页</template>
    </el-menu-item>

    <el-sub-menu index="/article">
      <template #title>
        <el-icon>
          <Document/>
        </el-icon>
        <span>博文管理</span>
      </template>
      <el-menu-item index="/article/list">博文列表</el-menu-item>
      <el-menu-item index="/article/post">发表文章</el-menu-item>
    </el-sub-menu>

    <el-menu-item index="3">
      <el-icon>
        <Picture/>
      </el-icon>
      <template #title>图片管理</template>
    </el-menu-item>

    <el-menu-item index="4">
      <el-icon>
        <Tools/>
      </el-icon>
      <template #title>系统管理</template>
    </el-menu-item>

  </el-menu>

</template>

<script setup>
import {inject} from "vue";

const layoutConfig = inject("layoutConfig")

</script>

<style scoped>

.el-menu {
  height: 100%;
}

/*菜单未折叠的宽度 */
.el-menu:not(.el-menu--collapse) {
  width: 200px;
}


</style>