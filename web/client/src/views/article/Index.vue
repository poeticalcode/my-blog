<template>
  <el-card v-if="articleDetail">
    <h2>{{ articleDetail.title }}</h2>
    <el-divider />

    {{ articleDetail.md_text }}

    <el-divider />
    <div class="footer" style="color: rgb(133, 144, 166);
    display: flex;
    align-items: center;
    gap: 10px;">
      <span>发布于 {{ articleDetail.created_at }}</span>
      <span>最近更新于 {{ articleDetail.updated_at }}</span>
      <span>阅读数量 {{ articleDetail.view_num }}</span>
    </div>
  </el-card>
</template>


<script setup>
import { useRoute } from 'vue-router'
import { fetchArticleDetail } from "@/api/aritcle";
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

</style>