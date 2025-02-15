export default [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('@/views/Auth.vue'),
    },
    {
      path: '/quesbank',
      name: 'quesbank',
      component: () => import('@/views/QuestionBank.vue'),
    },
    {
      path: '/question/:id',
      name: 'question',
      component: () => import('@/views/QuestionDetail.vue'),
    },
    {
      path: '/preview/:id',
      name: 'preview',
      component: () => import('@/views/Preview.vue'),
    }
  ];
