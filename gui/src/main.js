import "@/plugins/buefy";
import "@/plugins/axios";
import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import "normalize.css";
import "@/registerServiceWorker";
import "@/assets/pace-progressbar";
import "@/assets/pace-progressbar/themes/blue/pace-theme-minimal.css";

Vue.config.productionTip = false;

new Vue({
  store,
  render: h => h(App)
}).$mount("#app");
