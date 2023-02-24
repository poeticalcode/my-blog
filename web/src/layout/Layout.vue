<template>
  <div class="container-fluid" style="padding:0;">
    <Header style="margin-bottom: 18px;background-color: white;"></Header>
    <div id="main" class="container" style="margin-top: 1.8rem;max-width: 1140px;">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
import Header from "./components/Header.vue";
</script>

<style lang="scss">
</style>