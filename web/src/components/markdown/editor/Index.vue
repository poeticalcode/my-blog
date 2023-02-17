<template><div :id="id" class="editor-md"></div></template>

<script setup>
import asyncLoadJs from '@/util/fetchScript'
import { uploadImg } from '@/api/fileApi'

import { v4 as UUID4 } from 'uuid'
import { onMounted, ref, defineProps } from "vue";

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
    atLink    : false,
    syncScrolling: "single",
    placeholder: "记录你的想法",
    onload: function () {
      this.setMarkdown(props["modelValue"])
    },
    onchange: function () {
      emits("update:modelValue", this.getMarkdown())
    }
  })
  console.log(editor_.value);
  initPasteDragImg(editor_.value, handleUploadImg)
}

onMounted(initEditor)


const toHtml = () => {
  return editor_.value.getHTML()
}

const handleUploadImg = async (file, Editor) => {
  let formData = new FormData()
  formData.append('file', file)
  const { code, data: { url }, msg } = await uploadImg(formData)
  if (code === 2000) {
    console.log(url);
    if (/\.(png|jpg|jpeg|gif|bmp|ico)$/.test(url)) {
      Editor.insertValue("![图片alt](" + url + " '图片title')");
    } else {
      Editor.insertValue("[下载附件](" + msg["url"] + ")");
    }
    ElMessage.success("上传成功")
    return
  }
  ElMessage.error(msg)
}

// 增加拖拽粘贴上传
const initPasteDragImg = (Editor, uploadImg) => {
  let doc = document.getElementById(Editor.id)
  doc.addEventListener('paste', function (event) {
    let items = (event.clipboardData || window.clipboardData).items;
    let file = null;
    if (items && items.length) {
      // 搜索剪切板items
      for (let i = 0; i < items.length; i++) {
        if (items[i].type.indexOf('image') !== -1) {
          file = items[i].getAsFile();
          break;
        }
      }
    } else {
      console.log("当前浏览器不支持");
      return;
    }
    if (!file) {
      console.log("粘贴内容非图片");
      return;
    }
    uploadImg(file, Editor);
  });

  doc.addEventListener("dragover", function (e) {
    e.preventDefault()
    e.stopPropagation()
  })
  doc.addEventListener("dragenter", function (e) {
    e.preventDefault()
    e.stopPropagation()
  })
  doc.addEventListener("drop", function (e) {
    e.preventDefault()
    e.stopPropagation()
    let files = this.files || e.dataTransfer.files;
    uploadImg(files[0], Editor);
  })
}

</script>


<style scoped>
.editor-md {
  box-sizing: border-box;
}
</style>