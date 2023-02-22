<template>
  <div class="row">
    <div class="col-md-12 col-lg-9 col-xl-9">
      <!-- 没有数据就显示骨架屏 -->
      <div class="card" aria-hidden="true" v-for="i in 3" :key="'c' + i" v-if="!articleData.list">
        <div class="row g-0">
          <div class="col-md-5">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="180" xmlns="http://www.w3.org/2000/svg"
              role="img" aria-label="占 位 符" preserveAspectRatio="xMidYMid slice" focusable="false" _mstaria-label="177515"
              _mstHash="178">
              <title _mstTextHash="7742150" _mstHash="179">占 位 符</title>
              <rect width="100%" height="100%" fill="#868e96"></rect>
            </svg>
          </div>
          <div class="col-md-7">
            <div class="card-body">
              <h5 class="card-title text-truncate placeholder-glow"></h5>
              <p class="card-text text-truncate placeholder-glow"></p>
              <p class="card-text placeholder-glow"><small class="text-muted">发布时间：
                </small><small class="text-muted">
                  阅读数量：</small></p>
            </div>
          </div>
        </div>
      </div>
      <!-- 渲染文章 -->
      <div v-else class="card mb-3" v-for="item in articleData.list" :key="item.id" @click="toArticleDetail(item.id)">
        <div class="row g-0">
          <div class="col-md-5">
            <img :src="item.cover" class="card-img-top" alt="...">
          </div>
          <div class="col-md-7">
            <div class="card-body">
              <div>
                <p class="card-title">{{ item.title }}</p>
                <p class="card-desc">{{ item.description }}</p>
              </div>
              <p class="card-text">
                <small class="text-muted">
                  发布时间：{{ item.created_at }}
                </small>
                <small class="text-muted">
                  阅读数量：{{ item.view_num }}
                </small>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 左侧功能列表 -->
    <div class="col-hidden-down-md col-lg-3 col-xl-3">
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
    </div>
  </div>
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
.card {
  cursor: pointer;
  box-shadow: 0px 0px 10px 0px #e9e7e7;

  .card-body {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .card-title {
    font-size: 2rem;
    font-size: bold;
    margin-bottom: 0.8rem;
  }

  .card-desc {
    max-height: 124px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
}

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