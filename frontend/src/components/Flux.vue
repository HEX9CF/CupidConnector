<script setup lang="ts">
import { onMounted, ref, onUnmounted } from 'vue'
import * as echarts from 'echarts/core';
import { GaugeChart, GaugeSeriesOption } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';
import { GetInfo } from '../../wailsjs/go/application/App'
import { model } from '../../wailsjs/go/models'
import { EventsOn, EventsOff } from "../../wailsjs/runtime";

// 注册必须的组件
echarts.use([GaugeChart, CanvasRenderer]);
type EChartsOption = echarts.ComposeOption<GaugeSeriesOption>;

const info = ref<model.Info>();
const chartDom = ref<HTMLElement | null>(null);
const myChart = ref<echarts.ECharts | null>(null);
let option: EChartsOption;

const getInfo = async () => {
  await GetInfo().then((res) => {
    console.log(res)
    if (res.code === 1) {
      info.value = res.data;
    } else {
      info.value = {
        user_name: "未知用户",
        overall: 0,
        used: 0,
        expiration_time: "未知",
        account_status: "未知",
      };
    }
    if (info.value == null || info.value?.user_name == "") {
      info.value = {
        user_name: "未知用户",
        overall: 0,
        used: 0,
        expiration_time: "未知",
        account_status: "未知",
      };
    }
  })
}

const updateOption = async () => {
  if (myChart.value) {
    option = {
      series: [
        {
          min: 0,
          max: info.value?.overall,
          splitNumber: 8,
          type: 'gauge',
          itemStyle: {
            color: '#aaa',
            shadowColor: 'rgba(0,138,255,0.45)',
            shadowBlur: 10,
            shadowOffsetX: 2,
            shadowOffsetY: 2,
          },
          axisLine: {
            lineStyle: {
              width: 18,
              color: [
                [0.25, '#FF6E76'],
                [0.5, '#FDDD60'],
                [0.75, '#58D9F9'],
                [1, '#7CFFB2']
              ]
            }
          },
          axisTick: {
            show: true
          },
          splitLine: {
            length: 10,
            lineStyle: {
              width: 2,
              color: '#fff'
            }
          },
          axisLabel: {
            distance: 20,
            color: '#fff',
            fontSize: 5
          },
          anchor: {
            show: true,
            showAbove: true,
            size: 15,
            itemStyle: {
              borderWidth: 10
            }
          },
          title: {
            offsetCenter: [0, '60%'],
          },
          detail: {
            valueAnimation: true,
            fontSize: 16,
            offsetCenter: [0, '85%'],
            formatter: '{value}' + ' MB',
            color: '#fff'
          },
          data: [
            {
              value: parseFloat(((info.value?.overall ?? 0) - (info.value?.used ?? 0)).toFixed(2) ?? 0),
              name: '剩余流量',
            }
          ]
        }
      ]
    };
    option && myChart.value.setOption(option);
  }
}

const initChart = () => {
  if (chartDom.value && !myChart.value) {
    myChart.value = echarts.init(chartDom.value);
    window.addEventListener('resize', handleResize);
  }
}

const handleResize = () => {
  if (myChart.value) {
    myChart.value.resize();
  }
}

const updateInfo = async () => {
  await getInfo();
  initChart();
  await updateOption();
}

onMounted(async () => {
  // 监听信息更新事件
  EventsOn("updateInfo", updateInfo);

  // 初始化图表和数据
  await updateInfo();
})

onUnmounted(() => {
  // 清理事件监听和图表实例
  EventsOff("updateInfo");
  window.removeEventListener('resize', handleResize);
  if (myChart.value) {
    myChart.value.dispose();
  }
})
</script>

<template>
  <div class="flux-chart">
    <div class="statistics-area-chart" ref="chartDom"></div>
  </div>
  <div class="flux" v-if="info">
    已用流量: {{ info?.used.toFixed(2) }} MB
    <br/> 流量总额: {{ info?.overall.toFixed(2) }} MB
  </div>
</template>

<style scoped>
.statistics-area-chart {
  margin-top: 50px;
  width: 200px;
  height: 200px;
  text-align: center;
}

.flux-chart {
  display: flex;
  justify-content: center;
  text-align: center;
}

.flux {
  font-size: 12px;
  color: #aaa;
}
</style>