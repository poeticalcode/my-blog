<template>
  <div class="tool-bar">
    <div class="mac-header"></div>
    <!-- 语言类型-->
    <div class="mac-title">{{ props["lang"] }}</div>
    <div :id="copyId" class="copy-to-clipboard" id="copy-" aria-describedby="tool-bar-tooltip"
         @click="handleCopyMessage">
      <!-- 复制按钮 -->
      <svg t="1677218711606" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
           p-id="5059" width="16" height="16">
        <path
            d="M729.6 588.8v198.4c0 57.6-38.4 102.4-102.4 102.4H236.8c-57.6 0-102.4-38.4-102.4-102.4V396.8c0-57.6 38.4-96 102.4-96h390.4c57.6 0 96 38.4 96 96 6.4 64 6.4 128 6.4 192z m-64 0V390.4c0-25.6-12.8-32-32-32H236.8c-25.6 0-38.4 12.8-38.4 38.4v390.4c0 25.6 12.8 38.4 38.4 38.4h390.4c25.6 0 38.4-12.8 38.4-38.4V588.8z"
            p-id="5060" fill="#ffffff"></path>
        <path
            d="M883.2 435.2v198.4c0 51.2-38.4 89.6-83.2 96-19.2 0-38.4-6.4-38.4-25.6 0-19.2 12.8-32 32-32s32-12.8 32-32V236.8c0-19.2-12.8-32-32-32H390.4c-19.2 0-32 12.8-32 32-6.4 19.2-19.2 25.6-38.4 25.6s-25.6-19.2-25.6-38.4c6.4-51.2 44.8-83.2 96-83.2h403.2c57.6 0 96 38.4 96 96-6.4 64-6.4 128-6.4 198.4z"
            p-id="5061" fill="#ffffff"></path>
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
    color: rgba(255, 255, 255, 0.8);
  }

  .copy-to-clipboard {
    cursor: pointer;

    .icon {
      color: white;
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