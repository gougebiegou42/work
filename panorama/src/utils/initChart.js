import * as echart from 'echarts';
// 初始化echart  并设置响应式大小
const init = (dom, option) => {
  const myChart = echart.init(document.querySelector(dom));
  myChart.setOption(option);

  window.addEventListener('resize', () => {
    myChart.resize();
  });
  // debounce2(function () {
  //   return window.addEventListener('resize', () => {
  //     myChart.resize();
  //   });
  // }, 2000);
};

function debounce2(fn, wait) {
  let timeout;
  return function () {
    let context = this; // 获取this
    let arg = arguments; // 获取event对象
    clearTimeout(timeout);
    timeout = setTimeout(function () {
      fn.apply(context, arg); // 改变函数的this指向及接收event参数
      console.log('666');
    }, wait);
  };
}
export default init;
