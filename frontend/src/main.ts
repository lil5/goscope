import Vue from "vue";
import App from "./App.vue";
import router from "./router";

import {
  faSync,
  faServer,
  faClipboardList,
  faEye
} from "@fortawesome/free-solid-svg-icons";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(faSync, faServer, faClipboardList, faEye);
Vue.component("font-awesome-icon", FontAwesomeIcon);

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
