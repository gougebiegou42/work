import Vue from 'vue';
import * as echarts from 'echarts';
import App from './App.vue';
import store from './store';
import router from './router';
import '@/style/index.less';
import '@/assets/main.css';

import '@/utils/flexible';
// import '@/utils/china';

Vue.prototype.$echarts = echarts;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
