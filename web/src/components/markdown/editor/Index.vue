<template>
  <div :id="id" class="editor-md"></div>
</template>

<script setup>
import asyncLoadJs from '@/util/fetchScript'

import {v4 as UUID4} from 'uuid'
import {onMounted, ref, defineProps} from "vue";

const props = defineProps(["height", "modelValue"])

const emits = defineEmits(["update:modelValue"])

const id = ref(UUID4())

const editor_ = ref(null)

const initEditor = async function () {
  // 请求资源
  await asyncLoadJs('/plugins/editor.md/lib/jquery.js')
  await asyncLoadJs('/plugins/editor.md/editormd.js')
  $('head').append($('<link rel="stylesheet" type="text/css" href="/plugins/editor.md/css/editormd.css" />'));

  editor_.value = window.editormd(id.value, {
    path: '/plugins/editor.md/lib/',
    height: props["height"],
    syncScrolling: "single",
    placeholder: "记录你的想法",
    onload: function () {
      this.setMarkdown(props["modelValue"])
    },
    onchange: function () {
      emits("update:modelValue", this.getMarkdown())
    }
  })
}

onMounted(initEditor)


const toHtml = () => {
  return editor_.value.getHTML()
}

</script>


<style scoped>

.editor-md {
  box-sizing: border-box;
}

</style>