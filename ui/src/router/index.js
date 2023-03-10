import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import MainFeed from '../views/MainFeed.vue'
import LoginView from '../views/LoginView.vue'
import CalendarView from '../views/CalendarView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/feed',
    name: 'feed',
    component: MainFeed
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/calendar',
    name: 'logicalendern',
    component: CalendarView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
