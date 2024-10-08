import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import en from 'element-plus/es/locale/lang/en'

import '/src/static/css/base.css'

import i18n from "@/i18n/index.js";

import ECharts from 'vue-echarts'
import { use } from "echarts/core"
import { CanvasRenderer } from 'echarts/renderers'
import { BarChart, LineChart, PieChart } from 'echarts/charts'
import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent
} from 'echarts/components'

use([
  CanvasRenderer,
  BarChart,
  LineChart,
  PieChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent
])

const app = createApp(App)

const savedLanguage = localStorage.getItem('lang') || 'en'
i18n.global.locale.value = savedLanguage

app.use(router)
app.use(i18n)
app.use(ElementPlus, {
  locale: i18n.global.locale.value === 'en' ? en : zhCn,
})
app.component('v-chart', ECharts)
app.mount('#app')
