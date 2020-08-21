import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Requests",
    component: () => import("../views/RequestList.vue")
  },
  {
    path: "/logs",
    name: "Logs",
    component: () => import("../views/LogsList.vue")
  },
  {
    path: "/logs/:id",
    name: "Log Details",
    component: () => import("../views/LogDetails.vue")
  },
  {
    path: "/info",
    name: "System Information",
    component: () => import("../views/SystemInfo.vue")
  }
];

const router = new VueRouter({
  routes
});

export default router;
