<template>
  <div
    style="
      display: flex;
      flex-direction: column;
      height: 100%;
      width: 95%;
      margin: 0 auto;
    "
  >
    <el-form
      :inline="true"
      :model="formInline"
      label-position="left"
      size="large"
    >
      <el-form-item label="文章标题">
        <el-input
          v-model="formInline.user"
          input-style="flex:1"
          placeholder="文章标题"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary">保存</el-button>
      </el-form-item>
    </el-form>
    <Editor ref="myEditor" v-model="formInline.md"></Editor>
  </div>
</template>

<script setup>
import "@toast-ui/editor/dist/toastui-editor.css";

import Editor from "@/components/markdown-editor/Index.vue";

import { ref, reactive, watch, onMounted } from "vue";

const formInline = reactive({
  user: "",
  region: "",
  md: "年后",
});
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