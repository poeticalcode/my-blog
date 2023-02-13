<template>
  <el-row style="gap: 24px;">
    <!-- 内容区 -->
    <el-col :xs="24" :xm="24" :md="24" :lg="18" :xl="17">
      <el-card v-if="articleDetail" style=" padding: 18px;">
        <template #header>
          <h2>{{ articleDetail.title }}</h2>
        </template>
        <MmarkdownPreview v-model="articleDetail.md_text" toc="toc"></MmarkdownPreview>
        <div class="footer"
          style="color: rgb(133, 144, 166);display: flex;align-items: center;gap: 10px;text-align: right;justify-content: flex-end;">
          <span>发布于 {{ articleDetail.created_at }}</span>
          <span>最近更新于 {{ articleDetail.updated_at }}</span>
          <span>阅读数量 {{ articleDetail.view_num }}</span>
        </div>
      </el-card>
    </el-col>
    
    <!-- 左侧功能区 -->
    <el-col class="hidden-md-and-down"  :lg="5" :xl="6">
      <el-card v-if="articleDetail" :body-style="{ padding: '0px' }">
        <template #header>
          <div class="card-header">
            <span>{{ articleDetail.title }}</span>
          </div>
        </template>
        <!-- 目录内容 -->
        <div id="toc"></div>
      </el-card>
    </el-col>
  </el-row>
</template>


<script setup>
import { useRoute } from 'vue-router'
import { fetchArticleDetail } from "@/api/aritcle";
import MmarkdownPreview from "@/components/markdown-preview/Index.vue";
import { ref, reactive, watch } from "vue";
const route = useRoute()
const { id: id } = route.query


const articleDetail = ref()

const initTableData = async () => {
  const res = await fetchArticleDetail(id);
  if (res.code === 2000) {
    articleDetail.value = res.data
  }
};
initTableData()

</script>


<style scoped>
.el-card {
  border-radius: 5px;
}

.editormd-preview-container,
.editormd-html-preview {
  padding: 0;
}
</style>