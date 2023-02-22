<template>
  <div :id="id" class="markdown-body" v-html="compiledMarkdown"></div>
</template>

<script setup>

import { v4 as UUID4 } from 'uuid'
import { ref, defineProps, computed } from "vue";
import { marked } from 'marked'
import hljs from 'highlight.js'//引用
import 'highlight.js/styles/intellij-light.css';
import './themes/github.scss'
const props = defineProps(["height", "value", "toc"])

const id = ref(UUID4())

marked.setOptions({
  highlight: function (code) {
    return hljs.highlightAuto(code).value;
  }
})

const compiledMarkdown = computed(() => {
  return marked(props["value"])
})

</script>

<style lang="scss">
.code-copy-added:hover {

  .copy-btn,
  .copy-success-text {
    background: rgb(3 0 0 / 58%);
  }
}

p:has(img) {
  text-align: center;
}

// 添加相关样式
pre {
  padding: 32px 10px 10px 10px !important;
  overflow-x: auto !important;
  position: relative;
}

pre:before {
  content: " ";
  position: absolute;
  -webkit-border-radius: 50%;
  border-radius: 50%;
  background: #fc625d;
  margin-top: -2.2rem;
  width: 1.2rem;
  height: 1.2rem;
  -webkit-box-shadow: 2rem 0 #fdbc40, 4rem 0 #35cd4b;
  box-shadow: 2rem 0 #fdbc40, 4rem 0 #35cd4b;
  z-index: 2;
}


</style>