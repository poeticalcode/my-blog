<template>
  <el-row :gutter="8">
    <!-- 内容区 -->
    <el-col :xs="24" :xm="24" :md="18" :lg="18" :xl="18">
      <el-card v-if="articleDetail">
        <div style="font-size: 1.4rem;font-weight: bold;margin: 18px 0;">
          {{ articleDetail.title }}
        </div>
        <span>发布于 {{ articleDetail.created_at}}</span>
        <MarkdownPreview style="font-size: 1.6rem;" v-viewer v-code-copy v-target="'_blank'" :value="articleDetail.md_text" toc="toc" />
      </el-card>
    </el-col>

    <!-- 目录导航 -->
    <el-col class="hidden-xm-and-down" :lg="6" :xl="6">
      <el-card v-if="articleDetail">
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
import { fetchArticleDetail } from "@/api/aritcleApi";
import MarkdownPreview from "@/components/markdown/preview/Index.vue";
import { ref } from "vue";

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

:deep(.el-row) {
  margin: 0;
}
</style>