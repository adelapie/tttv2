import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Cell from "./components/Cell";
import Tictactoeboard from "./components/Tictactoeboard"

Vue.config.productionTip = false;

Vue.component('tic-tac-toe', Tictactoeboard);
Vue.component('cell', Cell);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
