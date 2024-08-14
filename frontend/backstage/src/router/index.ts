import { createRouter, createWebHashHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import { sendHttpRequest, ResponseStatus } from "frontend-resources";
import HomeView from "@/views/HomeView.vue";
import LoginView from "@/views/LoginView.vue";

type RouteRecordExtra = {
  label?: string;
};

type RouteRecord = RouteRecordExtra & RouteRecordRaw;

export const routes: RouteRecord[] = [
  {
    path: "/",
    name: "home",
    label: "Home",
    component: HomeView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/pages",
    name: "page",
    label: "Pages",
    component: () => import("@/views/PagesView.vue"),
  },
  {
    path: "/resources",
    name: "resource",
    label: "Resources",
    component: () => import("@/views/ResourcesView.vue"),
  },
  {
    path: "/articles",
    name: "article",
    label: "Articles",
    component: () => import("@/views/ArticlesView.vue"),
  },
  {
    path: "/files",
    name: "files",
    label: "Files",
    component: () => import("@/views/FilesView.vue"),
  },
  {
    path: "/tags",
    name: "tag",
    label: "Tags",
    component: () => import("@/views/TagsView.vue"),
  },
];

export const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach(async (to) => {
  const response = await sendHttpRequest<undefined, undefined, undefined>("", "/site/user-auth-challenge");
  if (response.status !== ResponseStatus.OK && to.name !== "login") {
    return { name: "login" };
  }
});
