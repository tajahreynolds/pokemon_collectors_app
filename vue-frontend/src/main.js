import './assets/main.css';

import { createApp } from 'vue';
import { createMemoryHistory, createRouter } from 'vue-router';

import App from './App.vue';
import Home from './components/Home.vue';
import Login from './components/Login.vue';
import PageNotFound from './components/PageNotFound.vue';

const routes = [
  { path: "/", component: Login },
  { path: "/home", component: Home },
  { path: "/:pathmatch(.*)", component: PageNotFound },
];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
});

createApp(App)
  .use(router)
  .mount('#app');
