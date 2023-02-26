<template>
  <div id="hot" style="width: 829px; height: 256px;">
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
import {onMounted} from "vue";

const initHotMap = () => {
  let chartDom = document.getElementById('hot');
  let myChart = echarts.init(chartDom);
  let option;

  function getVirtualData(year) {
    const date = +echarts.time.parse(year + '-01-01');
    const end = +echarts.time.parse(+year + 1 + '-01-01');
    const dayTime = 3600 * 24 * 1000;
    const data = [];
    for (let time = date; time < end; time += dayTime) {
      data.push([
        echarts.time.format(time, '{yyyy}-{MM}-{dd}', false),
        Math.floor(Math.random() * 10000)
      ]);
    }
    return data;
  }

  option = {
    title: {
      top: 30,
      left: 'center',
      text: '文章发布日历'
    },
    tooltip: {},
    visualMap: {
      min: 0,
      max: 10000,
      type: 'piecewise',
      orient: 'horizontal',
      left: 'center',
      top: 65
    },
    calendar: {
      top: 120,
      left: 30,
      right: 30,
      cellSize: ['auto', 13],
      range: '2023',
      itemStyle: {
        borderWidth: 0.5
      },
      yearLabel: {show: true}
    },
    series: {
      type: 'heatmap',
      coordinateSystem: 'calendar',
      data: getVirtualData('2023')
    }
  };

  option && myChart.setOption(option);
}

onMounted(() => {
  initHotMap()
})

</script>