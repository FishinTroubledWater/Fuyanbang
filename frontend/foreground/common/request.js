// 此vm参数为页面的实例，可以通过它引用vuex中的变量
module.exports = (vm) => {
  // 初始化请求配置
  uni.$u.http.setConfig((config) => {
    /* config 为默认全局配置*/
    // config.baseURL = 'http://124.222.141.238:8088'; /* 根域名 */
	    config.baseURL = 'http://localhost:8088'; /* 根域名 */
    // http://124.222.141.238:8088/v1/frontend/news/list    http://jsonplaceholder.typicode.com
    return config
  })
}