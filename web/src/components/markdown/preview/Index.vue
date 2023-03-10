<template>
  <div :id="id" class="markdown-body mac-header" v-html="compiledMarkdown"></div>
</template>

<script setup>
import {v4 as UUID4} from "uuid"
import {ref, defineProps, computed, h, onMounted, render, defineExpose} from "vue";
import {marked} from 'marked'
import "./styles/theme.scss"
import highlight from "highlight.js"
import ToolBar from "./ToolBar.vue";
// UUID
const id = ref(UUID4())
const tocTree = ref([]);

// 定义需要接收的 props
const props = defineProps(["height", "value", "toc"])

let tocList = [];
const toc = []

let anchor = 0;


marked.setOptions({
  langPrefix: "hljs "
})

// Override function
const renderer = {
  heading(text, level) {
    anchor += 1
    let tocAnchor = "toc-anchor-" + anchor
    tocList.push({level: level, title: text, 'anchor': tocAnchor})
    return `<h${level} id="${tocAnchor}"> ${text} </h${level}>`;
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
const dom2VNode = (el) => {
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
      child = dom2VNode(child)
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
    let {padding} = preComputedStyle
    let codeDom = pre.querySelector("code")
    const code = codeDom.innerText
    let codeVNode = dom2VNode(codeDom)
    let mountDom = document.createElement("div")
    let paddingSize = padding.slice(0, -2)
    let halfPadding = paddingSize / 3 + "px"
    const lang = pre.lang
    render(h('pre', {
      lang,
      style: {
        padding: "unset",
        border: "unset",
        overflowX: "hidden"
      }
    }, [
      h(ToolBar, {
        code: code,
        lang: lang || "未指定语言",
        style: {
          padding: padding,
          paddingTop: halfPadding,
          paddingBottom: halfPadding,
          borderBottom: "1px solid rgb(221 221 221)"
        }
      }),
      h(codeVNode, {
        style: {
          padding: padding,
          display: "block",
          overflowX: "auto"
        }
      })
    ]), mountDom)
    pre.before(mountDom.querySelector("pre"))
    pre.remove()
  })
}

onMounted(appendToolBar)
const genTocTree = function () {
  let root = {level: 0, title: "", children: []}
  let curr = root;
  for (let toc of tocList) {
    let level = toc.level
    let node = Object.assign(toc, {"children": []})

    // level 越大说明标题越小 ,curr.level >= level 说明遇到了一个比自己大的标题，要回退到上一个节点
    while (curr.level >= level) {
      curr = curr.parent;
    }
    node.parent = curr;
    curr.children.push(node)
    curr = node
  }
  return root.children
};
onMounted(() => {
  tocTree.value = genTocTree()
})

// 将目录树暴露出去
defineExpose(tocTree)
</script>

<style lang="scss">


</style>