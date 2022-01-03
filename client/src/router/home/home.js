export default {
    path:'/home',
    name:'home',
    component:() => import('@/views/home/home.vue'),
    children:[
        {
            path:'/home/index',
            name:'homeIndex',
            component:() => import('@/views/home/homeIndex.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/online',
            name:'online',
            component:() => import('@/views/home/onlineTable.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/output',
            name:'output',
            component:() => import('@/views/home/outputTable.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home/center',
            name:'center',
            component:() => import('@/views/home/center.vue'),
            meta:{
                isAuth: true,
            }
        },
        {
            path:'/home',
            redirect: '/home/index'
        }
    ]
}