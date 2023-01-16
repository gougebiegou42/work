import Vue from 'vue';
import * as echarts from 'echarts';
import App from './App.vue';
import store from './store';
import router from './router';
import '@/style/index.less';
import '@/assets/main.css';

import '@/utils/flexible';

Vue.prototype.$echarts = echarts;
// console.log(import.meta.env);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
