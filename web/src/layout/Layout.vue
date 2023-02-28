<template>
  <div class="container-fluid" style="padding:0;">
    <Header style="background-color: white;"></Header>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup>
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
import Header from "./components/Header.vue";
</script>

<style lang="scss">

</style>