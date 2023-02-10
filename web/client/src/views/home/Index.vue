<template>


  <el-row :gutter="10">

    <el-col :xs="8" :sm="6" :md="4" :lg="14" :xl="12" :offset="1">
      <!-- 下拉刷新 -->
      <!-- <div v-infinite-scroll="initTableData"> -->
        <!-- 渲染文章 -->
        <el-card style="margin: 16px 0;" :body-style="{ padding: '0px' }" v-for="item in articleList" :key="item.id"
          @click="toArticleDetail(item.id)">
          <div class="card-inner">
            <!-- 封面 -->
            <el-image style="width: 360px; " :src="item.cover" lazy fit="fill" />
            <!-- 文章信息 -->
            <div style="padding: 14px;flex: 1;height: 220px;">
              <p> {{ item.description }}</p>
            </div>
          </div>
        </el-card>
      <!-- </div> -->
    </el-col>

    <!--    <el-col :xs="8" :sm="6" :md="4" :lg="6" :xl="6">
      <div class="grid-content ep-bg-purple-light" />
    </el-col> -->
  </el-row>

</template>

<script setup>
import { fetchArticleList } from "@/api/aritcle";
import { ref, reactive, watch } from "vue";

let articleList = ref([])

let startNum = 1
let pageSize = 10
let currentPage = -1
let totalPage = 0

// 初始化表格数据
const initTableData = async () => {
  if (currentPage === totalPage && totalPage > 0) {
    console.log("已经到底");
    return
  }
  const res = await fetchArticleList({
    page_num: startNum,
    page_size: pageSize
  });

  if (res.code === 2000) {
    const { list: list, next_page: next_page, total_page: total_page, current_page: current_page } = res["data"]
    startNum = next_page
    totalPage = total_page
    currentPage = current_page
    articleList.value = list
    return
  }
};
initTableData()

// 去文章详情页面
const toArticleDetail = (id) => {
  window.open("/#/article?id=" + id)
}

</script>

<style scoped>
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}


.el-card .card-inner {
  display: flex;
  height: 220px;
}

.el-card {
  cursor: pointer;
}


.main-inner {
  display: flex;
  justify-content: center;
}
</style>