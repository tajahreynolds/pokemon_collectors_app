import './assets/main.css';

import { createApp } from 'vue';
import { createWebHistory, createRouter } from 'vue-router';

import App from './App.vue';
import Home from './components/Home.vue';
import Login from './components/Login.vue';
import Profile from './components/Profile.vue';
import PageNotFound from './components/PageNotFound.vue';

const routes = [
    { path: '/', component: Login },
    { path: '/home', component: Home },
    { path: '/profile', component: Profile },
    { path: '/:pathmatch(.*)', component: PageNotFound },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

createApp(App).use(router).mount('#app');
