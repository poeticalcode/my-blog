<template>
  <el-breadcrumb separator="/">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>
      <a href="/">博文管理</a>
    </el-breadcrumb-item>
    <el-breadcrumb-item>博文列表</el-breadcrumb-item>
  </el-breadcrumb>

  <el-card shadow="never">
    <el-table :data="tableData.list" border style="width: 100%" height="600">
      <el-table-column type="index" label="#" width="50" align="center" />
      <el-table-column
        prop="title"
        label="标题"
        width="180"
        align="center"
        :resizable="false"
      />
      <el-table-column
        prop="cover"
        label="封面"
        width="130"
        align="center"
        :resizable="false"
      >
        <template #default="scope">
          <el-tooltip effect="light" placement="right">
            <template #content>
              <el-image
                style="width: 200px; height: 150px"
                :key="scope.row.id"
                :src="scope.row.cover"
                lazy
              />
            </template>
            <div
              style="
                display: flex;
                align-items: center;
                justify-content: center;
              "
            >
              <el-image
                style="width: 100px; height: 75px"
                :key="scope.row.id"
                :src="scope.row.cover"
                lazy
              />
            </div>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        prop="description"
        label="简介"
        align="center"
        width="300"
        :resizable="false"
      >
        <template #default="scope">
          <el-tooltip effect="light" placement="right">
            <template #content>
              <div style="width: 250px">
                {{ scope.row.description }}
              </div>
            </template>
            <div
              style="
                overflow: hidden;
                text-overflow: ellipsis;
                display: -webkit-box;
                -webkit-line-clamp: 2;
                -webkit-box-orient: vertical;
              "
            >
              {{ scope.row.description }}
            </div>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column
        prop="view_num"
        label="阅读次数"
        align="center"
        width="90"
        :resizable="false"
      />
      <el-table-column
        prop="updated_at"
        label="状态"
        align="center"
        width="85"
        :resizable="false"
      >
        <template #default="scope">
          <el-tag type="success" v-if="scope.row.status == 1">已发布</el-tag>
          <el-tag type="danger" v-if="scope.row.status == 0">未发布</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        prop="updated_at"
        label="更新时间"
        align="center"
        width="160"
        :resizable="false"
      />
      <el-table-column
        prop="created_at"
        label="创建时间"
        align="center"
        width="160"
        :resizable="false"
      />

      <el-table-column
        label="操作"
        fixed="right"
        width="200"
        align="center"
        :resizable="false"
      >
        <template #default="scope">
          <el-button size="small" @click="handleEdit(scope.row.id)"
            >编辑</el-button
          >
          <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.row.id)"
            >删除</el-button
          >

          <el-button
            size="small"
            type="success"
            v-if="scope.row.status == 0"
            @click="handlePublic(scope.row.id)"
            >发布</el-button
          >
          <el-button
            size="small"
            type="wanring"
            v-if="scope.row.status == 1"
            @click="handleCancelPublic(scope.row.id)"
            >取消发布</el-button
          >
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { articleList } from "@/api/aritcle";
import { ref, reactive } from "vue";

const tableData = ref([]);

const pagingParam = reactive({
  page_num: 1,
  page_size: 10,
});

// 初始化表格数据
const initTableData = async () => {
  const res = await articleList(pagingParam);
  if (res.code === 2000) {
    tableData.value = res["data"];
  }
};
// 初始化表格数据
initTableData();
// 编辑
const handleEdit = (id) => {
  console.log(id);
};

// 删除
const handleDelete = (id) => {
  console.log(id);
};

// 发布
const handlePublic = (id) => {};
// 取消发布
const handleCancelPublic = (id) => {};
</script>

<style scoped>
.el-card {
  margin-top: 18px;
}
</style>