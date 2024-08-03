import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import LoginView from "@/views/LoginView.vue";
import FilesView from "@/views/FilesView.vue";
import ArticlesView from "@/views/ArticlesView.vue";
import PagesView from "@/views/PagesView.vue";
import ResourcesView from "@/views/ResourcesView.vue";
import TagsView from "@/views/TagsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
    },
    {
      path: "/files",
      name: "files",
      component: FilesView,
    },
    {
      path: "/articles",
      name: "article",
      component: ArticlesView,
    },
    {
      path: "/pages",
      name: "page",
      component: PagesView,
    },
    {
      path: "/resources",
      name: "resource",
      component: ResourcesView,
    },
    {
      path: "/tags",
      name: "tag",
      component: TagsView,
    },
  ],
});

export default router;
