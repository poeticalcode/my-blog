<template>
  <header :style="{position: 'relative',height:'400px', background:'url('+articleDetail.cover+')'}"
          v-if="articleDetail">
    <p> {{ articleDetail.title }}</p>
  </header>
  <article id="main-inner" class="container" style="margin-top: 1.8rem;max-width: 1140px;" v-if="articleDetail">
    <div class="row">
      <div class="col-md-12">
        <div class="card article">
          <div class="card-body">
            <MarkdownPreview v-viewer v-target="'_blank'" :value="articleDetail.md_text" toc="toc"/>
          </div>
          <div class="card-footer bg-transparent">
            <small class="text-muted" style="font-size: 1.2rem;text-align: right;padding: 0.4rem;">发布于:
              {{ articleDetail.created_at }}</small>
          </div>
        </div>
      </div>
      <div class="col-hidden-down-xll col-xl-3">
        <div v-if="articleDetail">
          <div class="card" id="toc"></div>
        </div>
      </div>
    </div>
  </article>
</template>


<script setup>
import {useRoute} from 'vue-router'
import {fetchArticleDetail} from "@/api/aritcleApi";
import MarkdownPreview from "@/components/markdown/preview/Index.vue";
import {ref} from "vue";

const route = useRoute()
const {articleId} = route.params
const articleDetail = ref()

const initTableData = async () => {
  const res = await fetchArticleDetail(articleId);
  if (res.code === 2000) {
    articleDetail.value = res.data
  }
};
initTableData()

</script>


<style lang="scss">

nav.navbar{
  margin-bottom: 0 !important;
}
#main-inner {
  max-width: 960px !important;
}

header {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2.8rem;
  color: #41464b;
  font-weight: bold;
}

header p {
  letter-spacing: 1.6rem;
}

article{
  position: relative;
  top: -10rem;
}

.card.article {
  user-select: text;
  padding: 0 1.6rem;
  box-shadow: 0 0 25px 0 rgba(0, 0, 0, .09);
  margin-bottom: 18px;
  border: none;

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

@media (max-width: 767px) {

  body {
    background-color: white;

    #app nav {
      display: none !important;
    }

    .card.article {
      box-shadow: none;
      padding: 0;

      .card-header {
        margin: 1.2rem 0;
      }

      .card-body {
        padding: 0;
      }
    }
  }
}
</style>