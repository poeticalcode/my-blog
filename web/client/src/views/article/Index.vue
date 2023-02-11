<template>
  {{ articleDetail }}
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
  if(res.code === 2000){
    articleDetail.value = res.data
  }
};
initTableData()

</script>


<style scoped>

</style>