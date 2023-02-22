<template>
  <div :id="id" class="markdown-body mac-header" v-html="compiledMarkdown"></div>
</template>

<script setup>

import {v4 as UUID4} from "uuid"
import {ref, defineProps, computed, watch, onMounted, createVNode, render} from "vue";
import {marked} from 'marked'
import h from "highlight.js"
import "highlight.js/styles/intellij-light.css";

import "@/components/markdown/preview/themes/github.scss"
import ToolBar from "@/components/markdown/preview/ToolBar.vue";
// UUID
const id = ref(UUID4())

const props = defineProps(["height", "value", "toc"])

// 设置选项
marked.setOptions({
  // 代码高亮
  highlight: function (code) {
    return h.highlightAuto(code).value;
  }
})

const compiledMarkdown = computed(() => {
  return marked(props["value"])
})

onMounted(() => {
  let markdown = document.getElementById(id.value)
  markdown.querySelectorAll("pre").forEach(pre => {
    if (pre.classList.contains('code-copy-added')) return
    let copy = createVNode(ToolBar, {
      code: pre.innerText
    })
    copy.parent = pre
    pre.style.position = "relative"
    let mountNode = document.createElement("div");
    render(copy, mountNode)
    pre.classList.add('code-copy-added')
    pre.appendChild(copy.el)
  })
})


</script>

<style lang="scss">
.code-copy-added:hover {
  .copy-btn,
  .copy-success-text {
    background: rgb(3 0 0 / 58%);
  }
}

.markdown-body {
  // 图片居中
  p:has(img) {
    text-align: center;
  }
}


// mac 风格代码块
$pre-padding: 1rem;

.mac-header pre {
  padding: $pre-padding *3 $pre-padding $pre-padding;
  position: relative;
  overflow-x: hidden;
  box-shadow: 0 0 $pre-padding 0 #8A7B7B66;
}

.mac-header pre::before {
  content: " ";
  position: absolute;
  background: #FF6057;
  margin-top: - 2*$pre-padding;
  width: $pre-padding;
  height: $pre-padding;
  left: $pre-padding;
  -webkit-border-radius: 50%;
  border-radius: 50%;
  -webkit-box-shadow: $pre-padding 0 #FFBD2F, $pre-padding 0 #28C93F;
  box-shadow: 2 * $pre-padding 0 #FFBD2F, 4 * $pre-padding 0 #28C93F;
}


</style>