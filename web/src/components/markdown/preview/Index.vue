<template>
  <div :id="id" class="editor-md"></div>
</template>

<script setup>
import asyncLoadJs from '@/util/fetchScript'

import {v4 as UUID4} from 'uuid'
import {onMounted, ref, defineProps} from "vue";

const props = defineProps(["height", "value", "toc"])

const id = ref(UUID4())

const editor_ = ref(null)


// 初始化编辑器
const initEditor = async function () {
  // 请求资源
  await asyncLoadJs('/plugins/editor.md/lib/jquery.js')
  await asyncLoadJs('/plugins/editor.md/lib/marked.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/prettify.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/raphael.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/underscore.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/sequence-diagram.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/flowchart.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/jquery.flowchart.min.js')
  await asyncLoadJs('/plugins/editor.md/lib/marked.min.js')
  await asyncLoadJs('/plugins/editor.md/editormd.min.js')


  $('head').append($('<link rel="stylesheet" type="text/css" href="/plugins/editor.md/css/editormd.css" />'));

  editor_.value = editormd.markdownToHTML(id.value, {
    markdown: props["value"],//+ "\r\n" + $("#append-test").text(),
    //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
    htmlDecode: "style,script,iframe",  // you can filter tags decode
    atLink    : false,
    toc: true,
    tocm: true,    // Using [TOCM]
    tocContainer: "#" + props["toc"], // 自定义 ToC 容器层
    //gfm             : false,
    tocDropdown: true,
    // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
    emoji: true,
    taskList: true,
    tex: true,  // 默认不解析
    flowChart: true,  // 默认不解析
    sequenceDiagram: true,  // 默认不解析
  })





}

onMounted(initEditor)


</script>


<style lang="scss">

.editor-md.markdown-body.editormd-html-preview {
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
}

.code-copy-added:hover {
  .copy-btn ,.copy-success-text{
    background: rgb(3 0 0 / 58%);
  }
}
</style>