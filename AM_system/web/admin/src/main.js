import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from '@/plugins/axiosInstance.js';
// src/main.js
createApp(App).use(router).mount('#app');

const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.use(router)
app.mount('#app')