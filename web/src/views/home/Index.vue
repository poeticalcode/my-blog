<template>
  <el-row :gutter="10">
    <el-col :md="24" :lg="18" :xl="18">
      <!-- 下拉刷新 -->
      <!-- <div v-infinite-scroll="initTableData"> -->
      <!-- 渲染文章 -->
      <el-card class="article-item" :body-style="{ padding: '0px' }" v-for="item in articleData.list" :key="item.id"
        @click="toArticleDetail(item.id)">
        <el-row class="card-inner">
          <el-col :sm="24" :md="10" :lg="9" :xl="10">
            <!-- 封面 -->
            <el-image style="height: 100%;" :src="item.cover" lazy fit="fill" />
          </el-col>
          <el-col :sm="24" :md="14" :lg="15" :xl="14">
            <!-- 文章信息 -->
            <div class="article-text-info">
              <div>
                <div style="font-weight: bold;font-size: 20px;">
                  <span>{{ item.title }}</span>
                </div>
                <div class="description">
                  <p> {{ item.description }}</p>
                </div>
              </div>
              <div class="footer">
                <span>发布时间：{{ item.created_at }}</span>
                <span>阅读数量：{{ item.view_num }}</span>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-card>

      <div style="width: 100%;width: 100%;display: flex;align-items: center;justify-content: center;">
        <!-- 分页组件 -->
        <el-pagination v-model:current-page="pagingParam.page_num" v-model:page-size="pagingParam.page_size"
          :small="false" :hide-on-single-page="true" :background="true" layout="total, prev, pager, next, jumper"
          :total="articleData.total" />
      </div>

      <!-- </div> -->
    </el-col>
    <!-- 左侧功能列表 -->
    <el-col class="hidden-md-and-down" :lg="6" :xl="6">
      <el-space direction="vertical">
        <el-card>
          <el-calendar>
            <template #header="{ data }">
              <div>{{ data }}</div>
            </template>
            <template #date-cell="{ data }">
              <div style="text-align: center;" :title="data.day">
                <span>{{ data.day.split('-')[2] }}</span>
              </div>
            </template>
          </el-calendar>
        </el-card>
      </el-space>

    </el-col>
  </el-row>
</template>

<script setup>

import { fetchArticleList } from "@/api/aritcleApi";
import { ref, reactive, watch } from "vue";

let articleData = ref([])

const pagingParam = reactive({
  page_num: 1,
  page_size: 10
})

// 监听页面参数
watch(pagingParam, (newVal, oldVal) => {
  initTableData()
})

// 初始化表格数据
const initTableData = async () => {
  const res = await fetchArticleList(pagingParam);
  if (res.code === 2000) {
    articleData.value = res["data"]
    return
  }
};
initTableData()

// 去文章详情页面
const toArticleDetail = (id) => {
  window.open("/article?id=" + id)
}



</script>

<style lang="scss" scoped>
.el-calendar {
  --el-calendar-border: var(--el-table-border, 1px solid var(--el-border-color-lighter));
  --el-calendar-header-border-bottom: var(--el-calendar-border);
  --el-calendar-selected-bg-color: var(--el-color-primary-light-9);
  --el-calendar-cell-width: 45px;
  background-color: var(--el-fill-color-blank);

  :deep(.el-calendar__header) {
    display: none;
  }

  :deep(.el-calendar__body) {
    border: 1px solid #ebeef5;
    border-radius: 5px;
    padding: unset;
  }


  .infinite-list {
    height: 300px;
    padding: 0;
    margin: 0;
    list-style: none;
  }

}


.el-card {
  $cardHeight: 200px;
  cursor: pointer;
  margin: 0px 0px 16px;

  .card-inner {
    display: flex;
    min-height: $cardHeight;

    .article-text-info {
      padding: 14px;
      flex: 1;
      height: $cardHeight;
      box-sizing: border-box;
      display: flex;
      flex-direction: column;
      justify-content: space-between;

      .description {
        font-size: 14px;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 4;
        -webkit-box-orient: vertical;
      }

      .footer {
        font-size: 13px;
        display: flex;
        gap: 8px;
        color: #8590a6;
      }
    }
  }
}
</style>