<template>
  <el-row :gutter="10">
    <el-col :md="24" :lg="18" :xl="18">
      <!-- 下拉刷新 -->
      <!-- <div v-infinite-scroll="initTableData"> -->
      <!-- 渲染文章 -->
      <el-card :body-style="{ padding: '0px' }" v-for="item in articleList" :key="item.id"
        @click="toArticleDetail(item.id)">
        <el-row class="card-inner">
          <el-col :md="10" :lg="10" :xl="10">
            <!-- 封面 -->
            <el-image style="width: 360px;height: 220px;" :src="item.cover" lazy fit="fill" />
          </el-col>
          <el-col :md="14" :lg="14" :xl="14">
            <!-- 文章信息 -->
            <div class="article-text-info">
              <div class="description">
                <p> {{ item.description }}</p>
              </div>
              <div class="footer">
                <span>发布时间：{{ item.created_at }}</span>
                <span>阅读数量：{{ item.view_num }}</span>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-card>
      <!-- </div> -->
    </el-col>
    <!-- 左侧功能列表 -->
    <el-col class="hidden-md-and-down" :lg="6" :xl="6">
      <el-row>
        <el-col>
          <el-card>
            <el-calendar>
              <template #header="{ date }">
                <div></div>
              </template>
              <template #date-cell="{ data }">
                <div style="text-align: center;" :title="data.day">
                  <span>{{ data.day.split('-')[2] }}</span>
                </div>
              </template>
            </el-calendar>
          </el-card>
        </el-col>
      </el-row>
    </el-col>
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
  window.open("/article?id=" + id)
}



</script>

<style>
.el-calendar {
  --el-calendar-border: var(--el-table-border, 1px solid var(--el-border-color-lighter));
  --el-calendar-header-border-bottom: var(--el-calendar-border);
  --el-calendar-selected-bg-color: var(--el-color-primary-light-9);
  --el-calendar-cell-width: 45px;
  background-color: var(--el-fill-color-blank);
}

.el-calendar .el-calendar__body {
  border: 1px solid #ebeef5;
  border-radius: 5px;
  padding: unset;
}

.el-calendar .el-calendar__header {
  display: none;
}

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
  margin: 0px 0px 16px;
}

.article-text-info {
  padding: 14px;
  flex: 1;
  height: 220px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.article-text-info .footer {
  display: flex;
  gap: 8px;
  color: #8590a6;
}

.main-inner {
  display: flex;
  justify-content: center;
}

.el-card .description {
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 6;
  -webkit-box-orient: vertical;
}
</style>