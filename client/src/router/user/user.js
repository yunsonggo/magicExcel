export default {
    path:'/',
    name:'user',
    component:() => import('@/views/user/index.vue'),
    children:[
        {
            path:'/user/login',
            name:'login',
            component:() => import('@/views/user/login.vue')
        },
        {
            path:'/user/manager/yunsongcailu/register',
            name:'register',
            component:() => import('@/views/user/register.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/',
            redirect: '/user/login'
        }
    ]
}