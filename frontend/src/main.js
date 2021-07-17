import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import '@mdi/font/css/materialdesignicons.css'

// Vuesax!
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
Vue.use(Vuetify, {
  iconfont: 'md'
})

export default new Vuetify({})

Vue.config.errorHandler = function(err, vm, info) {
  console.log(`Error: ${err.toString()}\nInfo: ${info}`);
}

// Vuex!
import Vuex from 'vuex'
Vue.use(Vuex)

import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
  new Vue({
	vuetify: new Vuetify(),
    render: h => h(App)
  }).$mount('#app');
});
