import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import './assets/global.css'
import Table from "@/components/Table";

axios.defaults.baseURL = 'http://localhost:8088/v1/backend/'
Vue.config.productionTip = false

new Vue({
  router,
  store,
  Table,
  render: h => h(App)
}).$mount('#app')
