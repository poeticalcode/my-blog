<template>
  <div style="
      display: flex;
      flex-direction: column;
      height: 100%;
      width: 95%;
      margin: 0 auto;
    ">
    <el-form :inline="true" :model="formInline" label-position="left" size="large">
      <el-form-item label="文章标题">
        <el-input v-model="formInline.title" input-style="flex:1" placeholder="文章标题"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleAddArticle">保存</el-button>
      </el-form-item>
    </el-form>
    <Editor ref="myEditor" v-model="formInline.md_text"></Editor>
  </div>
</template>

<script setup>
import Editor from "@/components/markdown/editor/Index.vue";

import {addArticle} from "@/api/aritcle"

import {ref, reactive, watch, onMounted} from "vue";

const formInline = reactive({
  title: "",
  md_text: ""
});


const handleAddArticle = async () => {
  formInline.description = formInline.md_text
  const {code, msg} = await addArticle(formInline)
  if (code === 2000) {
    ElMessage.success(msg)
    return
  }
  ElMessage.error(msg)
}


</script>

<style scoped>
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