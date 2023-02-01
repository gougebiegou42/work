export default [
  {
    path: '/aaa',
    name: 'home',
    component: () => import('@/views/Panorama.vue'),
  },
  {
    path: '/map',
    name: 'map1',
    component: () => import('@/views/echartdemo.vue'),
  },
  {
    path: '/',
    name: 'map',
    component: () => import('@/views/map.vue'),
  },
];
