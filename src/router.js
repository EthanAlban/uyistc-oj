import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        // ==========================================    LOGIN
        {
            path: '/login',
            name: 'login',
            component (resolve) {
                require(['@/views/Login.vue'], resolve)
            }
            // component: () => import('@/views/Login.vue')
        },
        // ==========================================    Register
        {
            path: '/register',
            name: 'register',
            component (resolve) {
                require(['@/views/Register.vue'], resolve)
            }
            // component: () => import('@/views/Register.vue')
        },
        {
            path: '/',
            redirect: '/home/nav1'
        },

        {
            path: '/home',
            name: 'home',
            component: () => import('@/views/Home/Index.vue'),
            children: [{
                path: '/home/nav1',
                name: 'nav1',
                component (resolve) {
                    require(['@/views/Home/Nav1.vue'], resolve)
                }
                // component: () => import('@/views/Home/Nav1.vue'),
            },
            {
                path: '/home/nav2',
                name: 'nav2',
                component (resolve) {
                    require(['@/views/Home/Nav2.vue'], resolve)
                }
                // component: () => import('@/views/Home/Nav2.vue'),
            },
            {
                path: '/home/nav3',
                name: 'nav3',
                component (resolve) {
                    require(['@/views/Home/Nav3.vue'], resolve)
                }
                // component: () => import('@/views/Home/Nav3.vue'),
            },
            {
                path: '/home/nav4',
                name: 'nav4',
                component (resolve) {
                    require(['@/views/Home/Nav4.vue'], resolve)
                }
                // component: () => import('@/views/Home/Nav4.vue'),
            },
            {
                path: '/home/nav5',
                name: 'nav5',
                component (resolve) {
                    require(['@/views/Home/Nav5.vue'], resolve)
                }
                // component: () => import('@/views/Home/Nav5.vue'),
            },
            {
                path: '/problemDetail/:id',
                name: 'problemDetail',
                component (resolve) {
                    require(['@/views/Home/problemDetail.vue'], resolve)
                }
                // component: () => import('@/views/Home/problemDetail.vue')
            },
            {
                path: '/problemDetail_Contest/:contest_id/:problem_id',
                name: 'problemDetail_Contest',
                component (resolve) {
                    require(['@/views/Home/problemDetail_Contest.vue'], resolve)
                }
                // component: () => import('@/views/Home/problemDetail_Contest.vue')
            }
            ]
        },
        // ===========================================    USER
        {
            path: '/user',
            redirect: '/user/profile'
        },
        {
            path: '/user',
            name: 'user',
            component: () => import('@/views/User/index.vue'),
            children: [{
                path: '/user/profile',
                name: 'profile',
                component (resolve) {
                    require(['@/views/User/01_profile.vue'], resolve)
                }
                // component: () => import('@/views/User/01_profile.vue'),
            },
            {
                path: '/user/updateInfo',
                name: 'updateInfo',
                component (resolve) {
                    require(['@/views/User/02_updateInfo.vue'], resolve)
                }
                // component: () => import('@/views/User/02_updateInfo.vue'),
            },
            {
                path: '/user/ModifyPwd',
                name: 'ModifyPwd',
                component (resolve) {
                    require(['@/views/User/03_1modifyPwd.vue'], resolve)
                }
                // component: () => import('@/views/User/03_1modifyPwd.vue'),
            },
            {
                path: '/user/ModifyEmail',
                name: 'ModifyEmail',
                component (resolve) {
                    require(['@/views/User/03_2modifyEmail.vue'], resolve)
                }
                // component: () => import('@/views/User/03_2modifyEmail.vue'),
            },
            {
                path: '/user/loginHistory',
                name: 'loginHistory',
                component (resolve) {
                    require(['@/views/User/04_loginHistory.vue'], resolve)
                }
                // component: () => import('@/views/User/04_loginHistory.vue'),
            },

            ]
        },
        // ===========================================    TEACHER    ==============================================
        {
            path: '/teacher',
            redirect: '/teacher/class'
        },
        {
            path: '/teacher',
            name: 'teacher',
            // component: () => import('@/views/Teacher/index.vue'),
            component (resolve) {
                require(['@/views/Teacher/index.vue'], resolve)
            },
            children: [{
                path: '/teacher/question',
                name: 'profile',
                component (resolve) {
                    require(['@/views/Teacher/question.vue'], resolve)
                }
                // component: () => import('@/views/Teacher/question.vue'),
            },
            {
                path: '/teacher/assign',
                name: 'profile',
                component (resolve) {
                    require(['@/views/Teacher/assign.vue'], resolve)
                }
                // component: () => import('@/views/Teacher/assign.vue'),
            },
            {
                path: '/teacher/class',
                name: 'profile',
                component (resolve) {
                    require(['@/views/Teacher/class.vue'], resolve)
                }
                // component: () => import('@/views/Teacher/class.vue'),
            }

            ]
        },
        // =========================================================  admin  =================================================
        {
            path: '/admin',
            redirect: '/admin/login'
        },
        {
            path: '/admin',
            name: 'admin',
            component (resolve) {
                require(['@/views/Admin/index.vue'], resolve)
            },
            // component: () => import('@/views/Admin/index.vue'),
            children: [{
                path: '/admin/login',
                name: 'login',
                component (resolve) {
                    require(['@/views/Admin/00_login.vue'], resolve)
                }
                // component: () => import('@/views/Admin/00_login.vue'),
            },
            {
                path: '/admin/sysinfo',
                name: 'sysinfo',
                component (resolve) {
                    require(['@/views/Admin/01_sysinfo.vue'], resolve)
                }
                // component: () => import('@/views/Admin/01_sysinfo.vue'),
            },
            {
                path: '/admin/annoncement',
                name: 'annoncement',
                component (resolve) {
                    require(['@/views/Admin/02_annoncement.vue'], resolve)
                }
                // component: () => import('@/views/Admin/02_annoncement.vue'),
            },
            {
                path: '/admin/usrlist',
                name: 'usrlist',
                component (resolve) {
                    require(['@/views/Admin/03_userlist.vue'], resolve)
                }
                // component: () => import('@/views/Admin/03_userlist.vue'),
            },
            {
                path: '/admin/check_tea_list',
                name: 'check_tea_list',
                component (resolve) {
                    require(['@/views/Admin/04_check_teacher.vue'], resolve)
                }
                // component: () => import('@/views/Admin/04_check_teacher.vue'),
            },
            {
                path: '/admin/question',
                name: 'problem_list',
                component (resolve) {
                    require(['@/views/Admin/05_problem_list.vue'], resolve)
                }
                // component: () => import('@/views/Admin/05_problem_list.vue'),
            },
            {
                path: '/admin/assign',
                name: 'task_list',
                component (resolve) {
                    require(['@/views/Admin/06_task_list.vue'], resolve)
                }
                // component: () => import('@/views/Admin/06_task_list.vue')
            },
            ]
        },

        // ===========================================    信息    ==============================================
        {
            path: '/info',
            redirect: '/info/admitted'
        },
        {
            path: '/info',
            name: 'info',
            component: () => import('@/views/Info/index.vue'),
            children: [{
                path: '/info/admitted',
                name: 'admitted',
                component (resolve) {
                    require(['@/views/Info/01_datas.vue'], resolve)
                }
                // component: () => import('@/views/User/01_profile.vue'),
            },
            ]
        },


        // 未登录的时候显示首页而非登录页
        {
            path: "/404",
            name: "notFound",
            component (resolve) {
                require(['@/views/404.vue'], resolve)
            }
            // component: () => import('@/views/404.vue')

        },

        //testcase
        {
            path: '/testcase',
            name: 'content',
            component: () => import('@/views/testcase.vue')
        },
        {
            path: "/testcase",
            redirect: "/testcase"
        },
        // {
        //     path: "*", // 此处需特别注意置于最底部
        //     redirect: "/404"
        // },
    ]
})
