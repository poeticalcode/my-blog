<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary fixed-top">
    <div class="container">
      <a class="navbar-brand" href="/">He'Blog</a>
      <div>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse right" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <a class="nav-link active" aria-current="page" href="/">首页</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="archive">归档</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="tags">标签</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="about">关于我</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>

import { useRouter } from 'vue-router'
const router = useRouter()
</script>

<style lang="scss" scoped>
.navbar {
  height: 6rem;
  box-shadow: 0px 0px 12px rgba(0, 0, 0, 0.12);
  font-size: 1.4rem;
}

</style>