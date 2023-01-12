import * as echart from 'echarts';
// 初始化echart  并设置响应式大小
const init = (dom, option) => {
  const myChart = echart.init(document.querySelector(dom));
  myChart.setOption(option);
  window.addEventListener('resize', () => {
    myChart.resize();
  });
};
export default init;
