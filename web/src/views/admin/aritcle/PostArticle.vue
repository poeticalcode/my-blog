<template>
  <div style=" display: flex;flex-direction: column;height: 100%; margin: 0 auto;">
    <el-form :inline="true" :model="formInline" label-position="left" size="large">
      <el-form-item>
        <el-input v-model="formInline.title" input-style="flex:1" placeholder="文章标题"/>
      </el-form-item>
      <el-form-item>
        <el-popover
            placement="top-start"
            title="文章属性"
            :width="300"
            trigger="click"
        >
          <p>Are you sure to delete this?</p>
          <div style="text-align: right; margin: 0">
            <el-button size="small" text @click="visible = false">cancel</el-button>
            <el-button size="small" type="primary" @click="handleAddArticle">确认发布</el-button>
          </div>
          <template #reference>
            <el-button type="primary" >发布</el-button>
          </template>
        </el-popover>
      </el-form-item>
    </el-form>
    <Editor ref="myEditor" v-model="formInline.md_text"></Editor>
  </div>
</template>

<script setup>
import Editor from "@/components/markdown/editor/Index.vue";

import {addArticle} from "@/api/aritcleApi";

import {ref, reactive, watch, onMounted} from "vue";
import { useRouter } from 'vue-router'
const router = useRouter()


const formInline = reactive({
  title: "",
  md_text: ""
});


const handleAddArticle = async () => {
  formInline.description = formInline.md_text
  const {code, data, msg} = await addArticle(formInline)
  if (code === 2000) {
    ElMessage.success(msg)
    setTimeout(() => {
      router.push("/admin/article/edit?id=" + data.id)
    }, 1000)
    return
  }
  ElMessage.error(msg)
}


</script>

<style lang="scss" >
html,
body,
#app {
  width: 100%;
  height: 100%;
}
.el-form {
  padding: 0 18px;
  display: flex;
  height: 64px;
  align-items: center;
  gap: 18px;
}

.el-form .el-form-item:first-child {
  flex: 1;
}

.editor-md {
  margin: unset;

  .CodeMirror.cm-s-default.CodeMirror-wrap {
    height: 100%;
  }
}

.el-form .el-form-item {
  margin: unset;
}
</style>