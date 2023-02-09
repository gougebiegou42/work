export default [
  {
    path: '/home',
    name: 'home',
    component: () => import('@/views/Panorama.vue'),
  },
  {
    path: '/',
    name: 'backstage',
    component: () => import('@/views/Backstage.vue'),
    redirect: '/tab1',
    children: [
      {
        path: 'tab1',
        component: () => import('@/components/backStage/tab1.vue'),
      },
    ],
  },
];
