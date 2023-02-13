<template>
  <div class="header">
    <el-icon style="cursor: pointer;font-size: 24px" @click="isExpand = false" v-show="isExpand">
      <Expand />
    </el-icon>
    <el-icon style="cursor: pointer;font-size: 24px" @click="isExpand = true" v-show="!isExpand">
      <Fold />
    </el-icon>

    <div>
      <el-button size="small" @click="toFornt">回到前台</el-button>
      <el-dropdown>
        <el-icon style="margin-right: 8px; margin-top: 1px">
          <setting />
        </el-icon>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>View</el-dropdown-item>
            <el-dropdown-item>Add</el-dropdown-item>
            <el-dropdown-item>Delete</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <span>Tom</span>
    </div>
  </div>
</template>

<script setup>

import { inject, ref, watch } from "vue";

const layoutConfig = inject("layoutConfig")

let isExpand = ref(false)

watch(isExpand, (newVal) => {
  layoutConfig.isCollapse = newVal;
}, {
  immediate: true
})



const toFornt = () => {
  window.open('/')
}

</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex: 1;
}
</style>