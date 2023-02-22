<template>
  <div class="row">
    <div class="col-xl-9">
      <div class="card article" v-if="articleDetail">
        <div class="card-header bg-transparent ">
          {{ articleDetail.title }}
        </div>
        <div class="card-body">
          <MarkdownPreview v-viewer v-code-copy v-target="'_blank'" :value="articleDetail.md_text" toc="toc" />
        </div>
        <div class="card-footer bg-transparent">
          <small class="text-muted" style="font-size: 1.2rem;text-align: right;padding: 0.4rem;">发布于: {{
            articleDetail.created_at }}</small>
        </div>
      </div>
    </div>
    <div class="col-xl-3">
      <div v-if="articleDetail">
        <div class="card" id="toc"></div>
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


<style lang="scss" scoped>
.card.article {
  padding: 18px;
  box-shadow: 0px 0px 25px 0px rgba(0, 0, 0, .09);

  .card-header {
    font-size: 2.4rem;
  }

  .card-body {
    font-size: 1.4rem;
  }

  .card-footer {
    font-size: 1.2rem;
  }
}
</style>