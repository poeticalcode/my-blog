<template>
  <div class="tool-bar">
    <div class="copy-to-clipboard" aria-describedby="tool-bar-tooltip" @click="copyMessage">
      <!-- 复制按钮 -->
      <svg t="1609826359524" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
           p-id="2955">
        <path
            d="M770.63802083 933.875H216.92708332c-44.82421875 0-79.1015625-34.27734375-79.10156249-79.1015625V195.59374999h553.7109375c44.82421875 0 79.1015625 34.27734375 79.1015625 79.10156251v659.1796875zM190.55989583 248.328125v606.4453125c0 15.8203125 10.546875 26.3671875 26.36718751 26.3671875h500.97656249V274.6953125c0-15.8203125-10.546875-26.3671875-26.3671875-26.3671875H190.55989583z"
            p-id="2956"/>
        <path
            d="M612.43489583 424.98828125H296.02864583c-13.18359375 0-26.3671875-10.546875-26.3671875-26.3671875 0-13.18359375 10.546875-26.3671875 26.3671875-26.3671875h316.40625c13.18359375 0 26.3671875 10.546875 26.36718751 26.3671875 0 13.18359375-13.18359375 26.3671875-26.36718751 26.3671875z m0 131.8359375H296.02864583c-13.18359375 0-26.3671875-10.546875-26.3671875-26.3671875 0-13.18359375 10.546875-26.3671875 26.3671875-26.3671875h316.40625c13.18359375 0 26.3671875 10.546875 26.36718751 26.3671875 0 13.18359375-13.18359375 26.3671875-26.36718751 26.3671875z m0 131.8359375H296.02864583c-13.18359375 0-26.3671875-10.546875-26.3671875-26.3671875 0-13.18359375 10.546875-26.3671875 26.3671875-26.3671875h316.40625c13.18359375 0 26.3671875 10.546875 26.36718751 26.3671875 0 13.18359375-13.18359375 26.3671875-26.36718751 26.3671875z"
            p-id="2957"/>
        <path
            d="M828.64583333 90.125h-527.34375001c-15.8203125 0-26.3671875 10.546875-26.36718749 26.3671875s10.546875 26.3671875 26.36718751 26.3671875h527.34374999c15.8203125 0 26.3671875 10.546875 26.3671875 26.3671875v606.4453125H823.37239583v52.73437499h84.375V169.2265625c0-44.82421875-36.9140625-79.1015625-79.1015625-79.1015625z"
            p-id="2958"/>
        <path
            d="M797.00520833 802.0390625a26.3671875 26.3671875 0 1 0 52.73437501 0 26.3671875 26.3671875 0 1 0-52.73437501 0z"
            p-id="2959"/>
      </svg>
      <textarea style="display:none" v-model="code"></textarea>

      <div class="tool-bar-tooltip" role="tool-bar-tooltip">
        {{ success ? "Copied!" : "Copy to clipboard" }}
        <div class="arrow" data-popper-arrow></div>
      </div>
    </div>


  </div>


</template>

<script setup>

import ClipboardJS from 'clipboard'
import {ref, defineProps, onMounted} from "vue";
import {createPopper} from '@popperjs/core';


const props = defineProps(["code"])
//复制插件
const code = ref(props["code"])
const success = ref(false)

const copyMessage = function (value) {
  success.value = false
  let clipboard = new ClipboardJS('.copy-to-clipboard', {
    text: () => {
      return code.value
    }
  })
  clipboard.on('success', function (e) {
    success.value = true
    clipboard.destroy() // 销毁,避免多次点击重复出现
  })
  clipboard.on('error', function (e) {
  })
}

onMounted(() => {
  const copyToClipboard = document.querySelector('.copy-to-clipboard');
  const toolBarTooltip = document.querySelector('.tool-bar-tooltip');

  const popperInstance = createPopper(copyToClipboard, toolBarTooltip, {
    placement: 'top',
    modifiers: [
      {
        name: 'offset',
        options: {
          offset: [0, 8],
        },
      },
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
})


</script>

<style lang="scss" scoped>

.tool-bar {
  position: absolute;
  top: 0;
  right: 1.5rem;

  .copy-to-clipboard {
    cursor: pointer;

    .icon {
      width: 1.2rem;
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