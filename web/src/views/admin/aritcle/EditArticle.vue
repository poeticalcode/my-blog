<template>
  <div style="
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 95%;
        margin: 0 auto;
      ">
    <el-form :inline="true" :model="articleDetail" label-position="left" size="large">
      <el-form-item label="文章标题">
        <el-input v-model="articleDetail.title" input-style="flex:1" placeholder="文章标题" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleUpdateArticle">更新</el-button>
      </el-form-item>
    </el-form>
    <Editor ref="myEditor" v-model="articleDetail.md_text"></Editor>
  </div>
</template>

<script setup>
import Editor from "@/components/markdown/editor/Index.vue";
import { ref } from "vue";
import { useRoute, useRouter } from 'vue-router'

import { updateArticle, fetchArticleDetail } from "@/api/aritcleApi";


const router = useRouter()
const route = useRoute()
const { id: id } = route.query

const articleDetail = ref({})
const initTableData = async () => {
  const res = await fetchArticleDetail(id);
  if (res.code === 2000) {
    articleDetail.value = res.data
  }
};

if (id == undefined) {
  ElMessage({
    message: '传入 ID 错误，跳转至发布页面',
    type: 'error',
  })
  router.push("/admin/article/post")
}
initTableData()


const handleUpdateArticle = async () => {
  const { id, description, title, md_text } = articleDetail.value
  const { code, msg } = await updateArticle({
    id,
    description,
    title,
    md_text
  })
  if (code === 2000) {
    ElMessage.success(msg)
    return
  }
  ElMessage.error(msg)
}


</script>

<style >
html,
body,
#app {
  width: 100%;
  height: 100%;
}

.el-form {
  display: flex;
  justify-content: space-between;
  height: 64px;
  align-items: center;
  gap: 8px;
}

.el-form .el-form-item:first-child {
  flex: 1;
}

.el-form .el-form-item {
  margin: unset;
}
</style>