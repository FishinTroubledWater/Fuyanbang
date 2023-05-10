import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import './assets/global.css'
import Table from "@/components/Table";
import Breadcrumb from "@/components/Breadcrumb";

axios.defaults.baseURL = 'http://localhost:8088/v1/backend/'
Vue.config.productionTip = false

Vue.component("Table",Table);
Vue.component("Breadcrumb",Breadcrumb);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
