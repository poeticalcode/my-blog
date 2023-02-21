<template>

  <div class="row">
    <div class="col-xl-9">
      <div v-if="articleDetail">
        <div style="font-size: 2.4rem;font-weight: bold;margin: 18px 0;">
          {{ articleDetail.title }}
        </div>
        <span>发布于 {{ articleDetail.created_at}}</span>
        <MarkdownPreview style="font-size: 1.2rem;" v-viewer v-code-copy v-target="'_blank'" :value="articleDetail.md_text" toc="toc" />
      </div>
    </div>
    <div class="col-xl-3">
      <div v-if="articleDetail">
        <div id="toc"></div>
      </div>
    </div>
  </div>
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
</style>