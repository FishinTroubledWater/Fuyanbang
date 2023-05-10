import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import './assets/global.css'
import Table from "@/components/Table";
import Breadcrumb from "@/components/Breadcrumb";
import Drawer from "@/components/Drawer";
import Pagination from "@/components/Pagination";

axios.defaults.baseURL = 'http://localhost:8088/v1/backend/'
Vue.config.productionTip = false
//封装组件的全局声明
Vue.component("Pagination",Pagination);
Vue.component("Table",Table);
Vue.component("Breadcrumb",Breadcrumb);
Vue.component("Drawer",Drawer);
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
