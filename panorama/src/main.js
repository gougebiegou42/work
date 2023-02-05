import Vue from 'vue';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
// import install from '@/utils/global';
import * as echarts from 'echarts';
import App from './App.vue';
import store from './store';
import router from './router';
import 'animate.css';
import '@/style/index.less';
import '@/assets/main.css';

import '@/utils/flexible';

Vue.prototype.$echarts = echarts;
// console.log(import.meta.env);
Vue.use(ElementUI);
// Vue.use(install);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
