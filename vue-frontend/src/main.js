import './assets/main.css';

import { createApp } from 'vue';
import { createWebHistory, createRouter } from 'vue-router';

import App from './App.vue';
import Home from './components/Home.vue';
import Login from './components/Login.vue';
import Profile from './components/Profile.vue';
import PageNotFound from './components/PageNotFound.vue';

const routes = [
    { path: '/', component: Home, name: Home },
    { path: '/login', component: Login, name: Login },
    { path: '/profile', component: Profile, name: Profile },
    { path: '/:pathmatch(.*)', component: PageNotFound, name: PageNotFound },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach(async (to, from) => {
    if (to.name !== 'Login') {
        // open login form
    }
});

createApp(App).use(router).mount('#app');
