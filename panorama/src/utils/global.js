import * as echarts from 'echarts';
let a = 1;
export default {
  install(Vue) {
    Vue.prototype.$echarts = echarts;
    Vue.prototype.$a = a;
  },
};
