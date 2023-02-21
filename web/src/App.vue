<template>
  <router-view></router-view>
</template>

<style lang="scss" >
:root {
  font-size: 10px;
}
*{
  margin: 0;
  padding: 0;
}

/* 滚动条优化 start */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-track {
  background: #f6f6f6;
  border-radius: 2px;
}
::-webkit-scrollbar-thumb {
  background: #cdcdcd;
  border-radius: 2px;
}
::-webkit-scrollbar-thumb:hover {
  background: #747474;
}
::-webkit-scrollbar-corner {
  background: #f6f6f6;
}
/* 滚动条优化 end */
</style>