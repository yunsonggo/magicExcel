import Vue from 'vue'
import VueRouter from 'vue-router'
import UserRouter from './user/user'
import HomeRouter from './home/home.js'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this,location).catch(err=>err)
}

const routes = [
    UserRouter,
  HomeRouter,
  {
    path:'/*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode:'history',
  base:process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (!localStorage.getItem("token") && to.meta.isAuth === true) {
    router.push('/')
    Vue.prototype.$message.error("需要登录以后才能访问,请重新登录")
    return
  }
  next()
})

export default router
