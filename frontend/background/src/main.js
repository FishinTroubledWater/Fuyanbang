//全局引入组件
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Table from "@/components/Table";
import Breadcrumb from "@/components/Breadcrumb";
import Drawer from "@/components/Drawer";
import Pagination from "@/components/Pagination";
import  VueQuillEditor from 'vue-quill-editor'//调用编辑器
import moment from "moment";
//全局引入插件
import './plugins/axios'
import './plugins/element.js'

//全局引入样式表
import './assets/global.css'
import 'quill/dist/quill.core.css' // import styles
import 'quill/dist/quill.snow.css' // for snow theme
import 'quill/dist/quill.bubble.css' // for bubble theme

// axios.defaults.baseURL = 'http://124.222.141.238:8088/v1/backend/'
axios.defaults.baseURL = 'http://localhost:8088/v1/backend/'
Vue.config.productionTip = false

//封装组件的全局引入
Vue.component("Pagination",Pagination);
Vue.component("Table",Table);
Vue.component("Breadcrumb",Breadcrumb);
Vue.component("Drawer",Drawer);
Vue.use(VueQuillEditor)
Vue.prototype.$moment = moment;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
