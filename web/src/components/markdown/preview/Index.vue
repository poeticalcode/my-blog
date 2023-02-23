<template>
  <div :id="id" class="markdown-body mac-header" v-html="compiledMarkdown"></div>
</template>

<script setup>
import {v4 as UUID4} from "uuid"
import {ref, defineProps, computed, h, onMounted, render} from "vue";
import {marked, Renderer} from 'marked'
import "highlight.js/styles/intellij-light.css";
import "@/components/markdown/preview/themes/github.scss"
import highlight from "highlight.js"

import ToolBar from "./ToolBar.vue";


// UUID
const id = ref(UUID4())

// 定义需要接收的 props
const props = defineProps(["height", "value", "toc"])


const toc = []


let anchor = 0;
let tocList = [];

// Override function
const renderer = {
  heading(text, level) {
    anchor += 1
    tocList.push({level: level, id: text, 'anchor': "toc-" + anchor})
    return `<h${level} id="toc-${anchor}"> ${text} </h${level}>`;
  },
  // 修改链接为新打开标签页
  link: function (href, title, text) {
    return '<a href="' + href + '" title="' + text + '" target="_blank">' + text + '</a>';
  },
  code: function (code, lang) {
    lang = (lang || '').match(/\S*/)[0];
    code = highlight.highlightAuto(code).value;
    code = code.replace(/\n$/, '') + '\n';
    return '<pre lang="' + lang + '"><code class="'
        + this.options.langPrefix
        + lang
        + '">'
        + code
        + '</code></pre>\n';
  }
};

marked.use({renderer});

//  md 转 html
const compiledMarkdown = computed(() => {
  return marked(props["value"])
})


/**
 * 真实 dom 转为 虚拟
 * @param el
 * @returns {VNode}
 */
const realDom2VirtualDom = (el) => {
  const tag = el.nodeName;
  const attributes = el.attributes;
  const childNodes = el.childNodes;
  // 遍历属性，生成props
  const props = {};
  for (const node of attributes) {
    props[node.nodeName] = node.nodeValue
  }
  // 遍历子级，生成children
  const children = [];
  childNodes.forEach(child => {
    if (child.nodeType === 1) {   // 元素节点
      child = realDom2VirtualDom(child)
    } else {                      // 文本节点
      child = child.nodeValue;
    }
    children.push(child)
  })
  return h(tag, props, children);
}


const appendToolBar = () => {
  // 获取渲染 markdown 的实例
  let markdown = document.getElementById(id.value)
  // 获取实例内的 pre，用来丰富代码块
  markdown.querySelectorAll("pre").forEach(pre => {
    // 获取最终的代码块样式
    let preComputedStyle = window.getComputedStyle(pre)
    const {backgroundColor, borderRadius, borderBottom, margin, border, paddingLeft, paddingRight} = preComputedStyle
    // 消除左上右上圆角
    pre.style.borderTopLeftRadius = pre.style.borderTopRightRadius = "0";
    // 边框取消
    pre.style.border = "0"
    // margin 取消
    pre.style.margin = "0"
    let toolBar = h(ToolBar, {
      code: pre.innerText,
      lang: pre.lang || "未指定语言",
      style: {
        borderBottom,
        paddingLeft,
        paddingRight,
        "paddingTop": "0.8rem",
        "paddingBottom": "0.8rem"
      }
    })
    let wrapStyle = {backgroundColor, borderRadius, margin, border}
    let wrap = h("div",
        {
          style: wrapStyle
        },
        [
          toolBar,
          h(realDom2VirtualDom(pre))
        ])
    let mountNode = document.createElement("div");
    render(wrap, mountNode)
    pre.before(wrap.el)
    pre.remove()
  })
}

onMounted(appendToolBar)
</script>

<style lang="scss">

</style>