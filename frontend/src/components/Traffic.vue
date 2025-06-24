<template>
      <el-button type="info" @click="handleClick" :icon="Odometer" size="small">流量</el-button>
      <el-dialog v-model="dialogFormVisible" width="500">
          <div class="traffic-container">
            <div class="traffic-chart" ref="chartDom"></div>
          </div>
          <template #footer>
              <div class="dialog-footer">
                  <el-button @click="dialogFormVisible = false">返回</el-button>
              </div>
          </template>
      </el-dialog>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue"
import { ElButton, ElDialog } from "element-plus"
import { Odometer } from "@element-plus/icons-vue";
import * as echarts from "echarts/core";
import { LineChart, LineSeriesOption } from 'echarts/charts';
import { TooltipComponent, GridComponent, TitleComponent, LegendComponent, ToolboxComponent } from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import {EventsOn} from "../../wailsjs/runtime";
import {GetInternetSpeed} from "../../wailsjs/go/application/App";

// 注册必须的组件
echarts.use([
  LineChart,
  CanvasRenderer,
  TooltipComponent,
  GridComponent,
  TitleComponent,
  LegendComponent,
  ToolboxComponent
]);

type EChartsOption = echarts.ComposeOption<LineSeriesOption>;

const dialogFormVisible = ref(false)
const uploadSpeeds = ref<number[]>([]);
const downloadSpeeds = ref<number[]>([]);
const timePoints = ref<string[]>([]);
const size = 60;

const chartDom = ref<HTMLElement | null>(null);
const myChart = ref<echarts.ECharts | null>(null);
let option: EChartsOption;

const handleClick = () => {
    dialogFormVisible.value = true;
    setTimeout(() => {
      if (chartDom.value && !myChart.value) {
        initChart();
      }
      updateOption();
    }, 100);
};

const getInternetSpeed = async () => {
  await GetInternetSpeed().then((res) => {
    console.log(res)
    if (res.code === 1) {
      // 获取当前时间作为X轴标签
      const now = new Date();
      const timeStr = `${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}`;

      // 更新数据
      uploadSpeeds.value.push(res.data.uploadSpeed);
      if (uploadSpeeds.value.length > size) uploadSpeeds.value.shift();

      downloadSpeeds.value.push(res.data.downloadSpeed);
      if (downloadSpeeds.value.length > size) downloadSpeeds.value.shift();

      timePoints.value.push(timeStr);
      if (timePoints.value.length > size) timePoints.value.shift();
    }
  })
}

const initChart = () => {
  if (chartDom.value) {
    myChart.value = echarts.init(chartDom.value);
  }
}

const updateOption = () => {
  if (myChart.value) {
    option = {
      title: {
        text: '实时网络流量',
        left: 'center'
      },
      tooltip: {
        trigger: 'axis',
        formatter: function (params: any) {
          let result = params[0].name + '<br/>';
          params.forEach((param: any) => {
            result += param.seriesName + ': ' + param.value + ' KB/s<br/>';
          });
          return result;
        }
      },
      legend: {
        data: ['上传速度', '下载速度'],
        top: 30
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
      },
      toolbox: {
        feature: {
          saveAsImage: {}
        }
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: timePoints.value.length > 0 ? timePoints.value : Array(size).fill('').map((_, i) => `${i}s`)
      },
      yAxis: {
        type: 'value',
        name: 'KB/s',
        axisLabel: {
          formatter: '{value}'
        }
      },
      series: [
        {
          name: '上传速度',
          type: 'line',
          data: uploadSpeeds.value,
          smooth: true,
          showSymbol: false,
          lineStyle: {
            width: 2
          },
          areaStyle: {
            opacity: 0.2
          },
          itemStyle: {
            color: '#58D9F9'
          }
        },
        {
          name: '下载速度',
          type: 'line',
          data: downloadSpeeds.value,
          smooth: true,
          showSymbol: false,
          lineStyle: {
            width: 2
          },
          areaStyle: {
            opacity: 0.2
          },
          itemStyle: {
            color: '#7CFFB2'
          }
        },
      ]
    };

    myChart.value.setOption(option);
  }
}

const updateInternetSpeed = async () => {
  await getInternetSpeed();
  if (dialogFormVisible.value && myChart.value) {
    updateOption();
  }
}

onMounted(async () => {
  // 初始化数据
  const now = new Date();
  for (let i = 0; i < size; i++) {
    uploadSpeeds.value.push(0);
    downloadSpeeds.value.push(0);
    const time = new Date(now.getTime() - (size - i) * 1000);
    timePoints.value.push(`${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}`);
  }

  EventsOn("updateInternetSpeed", async () => {
    await updateInternetSpeed();
  });

  setTimeout(async () => {
    await updateInternetSpeed();
  }, 1000);
})
</script>

<style scoped>
.traffic-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.traffic-chart {
  width: 450px;
  height: 300px;
}

.dialog-footer {
  text-align: right;
  margin-top: 10px;
}
</style>