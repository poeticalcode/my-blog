<template>


  <el-card shadow="never">

    <el-table>

    </el-table>
  </el-card>

</template>

<script setup>

import {articleList} from "@/api/aritcle"
import {onMounted} from "vue";

onMounted(() => {
  articleList()
})


</script>

<style scoped>

</style>