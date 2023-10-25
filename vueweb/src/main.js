import { createApp } from 'vue'
import './style.css'
import router from "./router/index";
import App from './App.vue'
import axios from 'axios'

// axios.defaults.baseURL='http://10.96.208.18:8080/'
axios.defaults.baseURL='http://127.0.0.1:8080/'


const app=createApp(App)
app.provide('$axios', axios)
app.use(router)
app.mount('#app')
