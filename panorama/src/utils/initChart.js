import * as echart from 'echarts';
// 初始化echart  并设置响应式大小
const init = (dom, option) => {
  const myChart = echart.init(document.querySelector(dom));
  myChart.setOption(option);

  window.addEventListener(
    'resize',
    throttle2(function () {
      myChart.resize();
    }, 100),
  );
};

// 移除echart监听
const end = (dom, option) => {
  const myChart = echart.init(document.querySelector(dom));
  window.removeEventListener(
    'resize',
    throttle2(function () {
      myChart.resize();
    }, 100),
  );
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
function throttle2(fn, wait) {
  let timeout;
  return function () {
    let context = this;
    let args = arguments;
    if (!timeout) {
      timeout = setTimeout(function () {
        timeout = null;
        fn.apply(context, args);
      }, wait);
    }
  };
}

export default init;
