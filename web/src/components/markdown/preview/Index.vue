<template>
  <div :id="id" class="markdown-body" v-html="compiledMarkdown"></div>
</template>

<script setup>

import { v4 as UUID4 } from 'uuid'
import { ref, defineProps, computed } from "vue";
import { marked } from 'marked'
import hljs from 'highlight.js'//引用
import 'highlight.js/styles/intellij-light.css';

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
.markdown-body {
  padding: 0;
  box-sizing: border-box;
  user-select: text;

  img {
    cursor: zoom-in;
  }

  h2,
  h3,
  h4,
  h5,
  h6 {
    border-left: 4px solid #3E9DEA;
    padding-left: 4px;
  }

  pre {
    padding: 10px;
    border: 1px solid #ddd;
    white-space: pre-wrap;
    word-wrap: break-word;
    font-size: 1.4rem;
    border-radius: 5px;
  }
}

.code-copy-added:hover {

  .copy-btn,
  .copy-success-text {
    background: rgb(3 0 0 / 58%);
  }
}
</style>