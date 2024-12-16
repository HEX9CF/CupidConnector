<template>
  <el-card>
    <ElRow :gutter="20" style="width: 100%;">
        <ElCol :span="12">
          <div>
            <h3>用户信息</h3>
            <p>用户名：{{ info?.user_name }}</p>
            <p>过期时间：{{ info?.expiration_time }}</p>
            <p>用户状态：{{ info?.account_status }}</p>
          </div>
          <ElButton @click="refresh" style="margin-top: 10px; width: 10px;" :icon="RefreshRight" />
        </ElCol>
        <ElCol :span="12" style="text-align: right">
          <div class="statistics-area-chart" ref="chartDom"></div>
        </ElCol>
    </ElRow>
  </el-card>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ElCol, ElRow, ElCard, ElNotification } from 'element-plus'
import { GetInfo } from '../../wailsjs/go/main/App'
import { model } from '../../wailsjs/go/models'
import * as echarts from 'echarts/core';
import { GaugeChart, GaugeSeriesOption } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';
import { RefreshRight } from "@element-plus/icons-vue";
echarts.use([GaugeChart, CanvasRenderer]);
type EChartsOption = echarts.ComposeOption<GaugeSeriesOption>;

const info = ref<model.Info>()
const overall = ref<number>(0)
const used = ref<number>(0)

const getInfo = async () => {
    await GetInfo().then((res) => {
        if (res.code === 1) {
            info.value = res.data
            const overall_unit = res.data.overall.replace(/\d*\.?\d*/g, '');
            const used_unit = res.data.used.replace(/\d*\.?\d*/g, '');

            overall.value = res.data.overall.match(/\d+(\.\d+)?/g)?.join('') || '';
            used.value = res.data.used.match(/\d+(\.\d+)?/g)?.join('') || '';
            if (overall_unit === 'G') {
                overall.value = overall.value * 1024
            }
            if (used_unit === 'G') {
                used.value = used.value * 1024
            }
        } else {
          overall.value = 0;
          used.value = 0;
            ElNotification({
                title: '获取信息失败',
                message: res.msg,
                type: 'error',
            })
        }
    })
}

const chartDom = ref<HTMLElement | null>(null);
const myChart = ref<echarts.ECharts | null>(null);
var option: EChartsOption;

const updateOption = async () => {
    if (myChart.value) {
        option = {
            series: [
                {
                    min: 0,
                    max: overall.value,
                    splitNumber: 8,
                    type: 'gauge',
                    itemStyle: {
                        color: '#58D9F9',
                        shadowColor: 'rgba(0,138,255,0.45)',
                        shadowBlur: 10,
                        shadowOffsetX: 2,
                        shadowOffsetY: 2
                    },
                    progress: {
                        show: true,
                        width: 18
                    },
                    axisLine: {
                        lineStyle: {
                            width: 18
                        }
                    },
                    axisTick: {
                        show: false
                    },
                    splitLine: {
                        length: 10,
                        lineStyle: {
                            width: 2,
                            color: '#fff'
                        }
                    },
                    axisLabel: {
                        distance: 25,
                        color: '#fff',
                        fontSize: 10
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
                        fontSize: 20,
                        offsetCenter: [0, '85%'],
                        formatter: '{value}' + 'MB',
                        color: '#fff'
                    },
                    data: [
                        {
                            value: used.value,
                            name: '已使用',
                        }
                    ]
                }
            ]
        };
        option && myChart.value.setOption(option)
    }
}

const refresh = async () => {
    await getInfo()
    if (chartDom.value) {
        // 初始化 ECharts 实例
        myChart.value = echarts.init(chartDom.value);
        await updateOption();
    }
}

onMounted(async () => {
    await refresh();
})


</script>

<style scoped>
.statistics-area-chart {
  width: 270px;
  height: 270px;
}
</style>