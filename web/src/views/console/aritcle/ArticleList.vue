<template>

  <!-- 面包屑 -->
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>
      <a href="/">博文管理</a>
    </el-breadcrumb-item>
    <el-breadcrumb-item>博文列表</el-breadcrumb-item>
  </el-breadcrumb>


  <!-- 卡片 -->
  <el-card shadow="never">
    <template #header>
      <div class="card-header">
        <span>文章列表</span>
        <el-button class="button" type="primary" @click="handPost">撰写文章</el-button>
      </div>
    </template>
    <!-- 检索表单 -->
    <el-form :inline="true" :model="searchForm">
      <el-form-item label="标题">
        <el-input v-model="searchForm.title" placeholder="请输入需要检索的标题" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">检索</el-button>
      </el-form-item>
    </el-form>

    <el-divider />

    <!-- 数据表格 -->
    <el-table v-loading="tableDataLoading" :data="tableData.list" border style="width: 100%" height="600">
      <el-table-column type="index" label="#" width="50" align="center" />
      <el-table-column prop="title" label="标题" width="180" :resizable="false" />
      <el-table-column prop="cover" label="封面" width="130" :resizable="false">
        <template #default="scope">
          <el-tooltip effect="light" placement="right">
            <template #content>
              <el-image style="width: 200px; height: 150px" :key="scope.row.id" :src="scope.row.cover" lazy />
            </template>
            <div style=" display: flex; align-items: center; justify-content: center;">
              <el-image style="width: 100px; height: 75px" :key="scope.row.id" :src="scope.row.cover" lazy />
            </div>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="简介" min-width="300" :resizable="false">
        <template #default="scope">
          <el-tooltip effect="light" placement="top-end">
            <template #content>
              <div style="width: 250px">
                {{ scope.row.description }}
              </div>
            </template>
            <div style="
                overflow: hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 2;
                -webkit-box-orient: vertical;
              ">
              {{ scope.row.description }}
            </div>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="view_num" label="阅读次数" width="90" :resizable="false" />
      <el-table-column prop="updated_at" label="状态" width="85" :resizable="false">
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.status == 1">已发布</el-tag>
          <el-tag type="danger" v-if="scope.row.status == 0">未发布</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="updated_at" label="更新时间" width="160" :resizable="false" />
      <el-table-column prop="created_at" label="创建时间" width="160" :resizable="false" />

      <el-table-column label="操作" fixed="right" width="300" :resizable="false">
        <template #default="scope">
          <el-button size="small" @click="handleEdit(scope.row.id)">编辑</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>

          <el-button size="small" type="success" v-if="scope.row.status == 0"
            @click="handlePublic(scope.row.id)">发布</el-button>
          <el-button size="small" type="wanring" v-if="scope.row.status == 1"
            @click="handleCancelPublic(scope.row.id)">取消发布
          </el-button>

          <el-button size="small" type="primary" @click="handlePreview(scope.row.id)">预览</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页组件 -->
    <el-pagination style="margin: 18px 0;float: right;" v-model:current-page="searchForm.page_num"
      v-model:page-size="searchForm.page_size" :small="false" :hide-on-single-page="true" :background="true"
      :page-sizes="[10, 20, 50]" layout="total, sizes, prev, pager, next, jumper" :total="tableData.total" />

  </el-card>
</template>

<script setup>
import { fetchArticleList, deleteArticle } from "@/api/aritcleApi";
import { ref, reactive, watch } from "vue";

// 表格数据
const tableData = ref({});

// 表格数据加载状态
const tableDataLoading = ref(false)

// 检索表单
const searchForm = reactive({
  title: "",
  page_num: 1,
  page_size: 10,
});

// 监听页面参数
watch(searchForm, (newVal, oldVal) => {
  initTableData()
})

// 初始化表格数据
const initTableData = async () => {
  tableDataLoading.value = true
  const res = await fetchArticleList(searchForm);
  if (res.code === 2000) {
    tableData.value = res["data"];
    tableDataLoading.value = false;
    return
  }
  ElMessage({
    message: '加载失败：' + res.msg,
    type: 'error',
  })
  tableDataLoading.value = false;
};

// 初始化表格数据
initTableData();

// 编辑
const handleEdit = (id) => {
  window.open("/console/article/edit/" + id)
};

// 删除
const handleDelete = async (id) => {
  const res = await deleteArticle(id)
  console.log(res);
  if (res.code === 2000) {
    initTableData()
    return
  }
};


const handPost = () => {
  window.open("/console/article/post")
}

// 发布
const handlePublic = (id) => { };

// 取消发布
const handleCancelPublic = (id) => {
};


const handlePreview = (id) => {
  window.open("/article/" + id + "?action=preview")
};


// 提交检索请求
const onSubmit = () => {
  initTableData()
}

</script>

<style scoped>
.el-card {
  margin-top: 18px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>