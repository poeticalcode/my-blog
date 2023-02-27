<template>
  <div id="main-inner">
    <div id="hot" style="width: 829px; height: 256px;"></div>
  </div>
</template>

<script setup>

import * as echarts from 'echarts/core';
import {
  TitleComponent,
  CalendarComponent,
  TooltipComponent,
  VisualMapComponent
} from 'echarts/components';
import {HeatmapChart} from 'echarts/charts';
import {CanvasRenderer} from 'echarts/renderers';

echarts.use([
  TitleComponent,
  CalendarComponent,
  TooltipComponent,
  VisualMapComponent,
  HeatmapChart,
  CanvasRenderer
]);
import {onMounted, ref} from "vue";

import {fetchDailyArticleCount} from "@/api/aritcleApi.js";


const dailyArticleCountData = ref([])


const getBeforeDate = (n) => {
  let date = new Date();
  let year, month, day;
  date.setDate(date.getDate() - n);
  year = date.getFullYear();
  month = date.getMonth() + 1;
  day = date.getDate();
  return year + '-' + (month < 10 ? ('0' + month) : month) + '-' + (day < 10 ? ('0' + day) : day);
}

const initDailyArticleCount = async () => {
  const {code, data, msg} = await fetchDailyArticleCount()
  if (code === 2000) {
    let dayList = []
    let i = 365;
    while (i >= 0) {
      dayList.push([getBeforeDate(i)])
      i--
    }
    for (let dayListElement of dayList) {
      for (let activeListElement of data) {
        if (dayListElement[0] === activeListElement["date"]) {
          dayListElement.splice(1)
          dayListElement.push(activeListElement["count"])
        } else if (dayListElement.length === 1) {
          dayListElement.push(0)
        }
      }
    }
    dailyArticleCountData.value = dayList
    initHotMap()
  }
}


const initHotMap = () => {
  let chartDom = document.getElementById('hot');
  let myChart = echarts.init(chartDom, null, {locale: 'ZH'});
  let option;

  option = {
    title: {
      top: 30,
      left: 'center',
      text: '文章发布日历'
    },
    tooltip: {},
    visualMap: {
      type: 'piecewise',
      orient: 'horizontal',
      left: 'center',
      top: 65,
      pieces: [
        // {gte: 15, color: 'blue'}, // 不指定 max，表示 max 为无限大（Infinity）。
        {gte: 1, color: 'rgb(98,155,223)'},
        // {gte: 5, lte: 10, color: 'rgb(167,213,255)'},
        // {gte: 1, lte: 5, color: 'rgb(214,233,250)'},
        {lte: 0, color: 'rgb(238,238,238)'}],
    },
    calendar: {
      top: 120,
      left: 30,
      right: 30,
      cellSize: ['auto', 13],
      range: [getBeforeDate(365), getBeforeDate(0)],
      itemStyle: {
        borderWidth: 0.5
      },
      yearLabel: {show: true}
    },
    series: {
      type: 'heatmap',
      coordinateSystem: 'calendar',
      data: dailyArticleCountData.value
    }
  };
  console.log(option);
  option && myChart.setOption(option);
}

onMounted(() => {
  initDailyArticleCount()
})

</script>

<style lang="scss" scoped>
#main-inner {
  display: flex;
  justify-content: center;
}
</style>