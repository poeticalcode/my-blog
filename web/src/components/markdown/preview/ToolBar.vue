<template>
  <div class="tool-bar">
    <div class="mac-header"></div>
    <!-- 语言类型-->
    <div class="mac-title">{{ props["lang"] }}</div>
    <div :id="copyId" class="copy-to-clipboard" id="copy-" aria-describedby="tool-bar-tooltip"
         @click="handleCopyMessage">
      <!-- 复制按钮 -->
      <svg t="1677148441267" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
           p-id="8607" width="16" height="16">
        <path
            d="M763.392 543.744h-157.184V386.048a31.744 31.744 0 1 0-62.976 0v157.696H385.536a31.744 31.744 0 0 0 0 62.976h157.696v157.696a31.744 31.744 0 0 0 62.976 0v-157.696h157.696a31.744 31.744 0 0 0 0-62.976z"
            fill="" p-id="8608"></path>
        <path
            d="M802.816 240.128H348.16a113.664 113.664 0 0 0-113.664 113.664v454.144A113.664 113.664 0 0 0 348.16 921.6h454.144a113.664 113.664 0 0 0 113.664-113.664V353.792a113.664 113.664 0 0 0-113.152-113.664zM870.4 807.936a68.096 68.096 0 0 1-68.096 68.096H348.16a68.096 68.096 0 0 1-68.096-68.096V353.792A68.096 68.096 0 0 1 348.16 285.696h454.144A68.096 68.096 0 0 1 870.4 353.792z"
            fill="" p-id="8609"></path>
        <path
            d="M699.904 162.816a25.6 25.6 0 0 0 0-51.2H245.248A139.264 139.264 0 0 0 106.496 250.88v454.144a25.6 25.6 0 0 0 51.2 0V250.88a88.064 88.064 0 0 1 87.552-88.064z"
            fill="" p-id="8610"></path>
      </svg>
      <div :id="copyTipId" class="tool-bar-tooltip">
        {{ success ? "Copied!" : "Copy to clipboard" }}
        <div class="arrow" data-popper-arrow></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import ClipboardJS from 'clipboard'
import {createPopper} from '@popperjs/core';
import {ref, defineProps, onMounted} from "vue";
import {v4 as UUID4} from "uuid"
// UUID
const id = UUID4()
const copyId = ref("copy-" + id)
const copyTipId = ref("copy-tip-" + id)

const props = defineProps(["code", "lang"])
// 复制状态
const success = ref(false)

// 复制代码功能
const handleCopyMessage = function (value) {
  success.value = false
  let clipboard = new ClipboardJS("#" + copyId.value, {
    text: () => {
      return props["code"]
    }
  })
  clipboard.on('success', function (e) {
    success.value = true
    clipboard.destroy() // 销毁,避免多次点击重复出现
  })
  clipboard.on('error', function (e) {
    success.value = false
    clipboard.destroy() // 销毁,避免多次点击重复出现
  })
}

// 复制图标上的提示
const handleTooltip = () => {
  const copyToClipboard = document.querySelector("#" + copyId.value);
  const toolBarTooltip = document.querySelector("#" + copyTipId.value);

  const popperInstance = createPopper(copyToClipboard, toolBarTooltip, {
    placement: 'top',
    modifiers: [
      {
        name: 'offset',
        options: {
          offset: [0, 8],
        }
      }
    ],
  });

  const showEvents = ['mouseenter', 'focus'];
  const hideEvents = ['mouseleave', 'blur'];
  showEvents.forEach((event) => {
    copyToClipboard.addEventListener(event, () => {
      // 使工具提示可见
      toolBarTooltip.setAttribute('data-show', '');

      // 启用事件侦听器
      popperInstance.setOptions((options) => ({
        ...options,
        modifiers: [
          ...options.modifiers,
          {name: 'eventListeners', enabled: true},
        ],
      }));
      // 更新位置
      popperInstance.update();
    });
  });

  hideEvents.forEach((event) => {
    copyToClipboard.addEventListener(event, () => {
      // 隐藏工具提示
      toolBarTooltip.removeAttribute('data-show');
      // 禁用事件侦听器
      popperInstance.setOptions((options) => ({
        ...options,
        modifiers: [
          ...options.modifiers,
          {name: 'eventListeners', enabled: false},
        ],
      }));
    });
  });

}

onMounted(handleTooltip)


</script>

<style lang="scss" scoped>


$circle-size: 1rem;
$pre-padding: 1rem;
$bei: 2;
.tool-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;

  // mac 风格代码块
  .mac-header {
    //padding: $pre-padding * $bei * 2 $pre-padding * $bei $pre-padding * $bei;
    background: #FF6057;
    width: $circle-size;
    height: $circle-size;
    -webkit-border-radius: 50%;
    border-radius: 50%;
    -webkit-box-shadow: $pre-padding 0 #FFBD2F, $pre-padding 0 #28C93F;
    box-shadow: 2 * $circle-size 0 #FFBD2F, 4 * $circle-size 0 #28C93F;
  }

  .mac-title {
    font-size: 1.2rem;
    color: rgba(0, 0, 0, 0.6);
  }

  .copy-to-clipboard {
    cursor: pointer;

    .icon {
      width: 1.6rem;
      height: 1.6rem;
    }
  }
}


.tool-bar-tooltip {
  background: #333;
  color: white;
  padding: 0.4rem 0.8rem;
  font-size: 1.2rem;
  border-radius: 0.4rem;
  display: none;
}

.tool-bar-tooltip[data-show] {
  display: block;
}

.arrow,
.arrow::before {
  position: absolute;
  width: 0.8rem;
  height: 0.8rem;
  background: inherit;
}

.arrow {
  visibility: hidden;
}

.arrow::before {
  visibility: visible;
  content: '';
  transform: rotate(45deg);
}

.tool-bar-tooltip[data-popper-placement^='top'] > .arrow {
  bottom: -0.4rem;
}

.tool-bar-tooltip[data-popper-placement^='bottom'] > .arrow {
  top: -0.4rem;
}

.tool-bar-tooltip[data-popper-placement^='left'] > .arrow {
  right: -0.4rem;
}

.tool-bar-tooltip[data-popper-placement^='right'] > .arrow {
  left: -0.4rem;
}


</style>