import { Get, Post, Put, Delete, Get2, Post2, Get3, Post3, Get4, Post4, getCookie, Get5 } from "@/api/request";
import { selfLog, UrlEncode } from '../utils'
import axios from 'axios'
export default {
    // =========================================oj backend Django==============================================
    // 用户登录
    userLogin: (params) => {
        return Post('login', params);
    },
    announcement: (params) => {
        return Get('announcement?offset=0&limit=10', params)
    },
    // 用户详情（读取cookie session确定是哪一个用户）
    user_profile: () => {
        return Get('profile')
    },
    // 获取题目列表
    problem_list: () => {
        return Get("problem?paging=true&offset=0&limit=10&page=1")
    },
    // 获取题目列表
    problem_list_page: (offset, limit, page, tag) => {
        return Get("problem?paging=true&offset=" + offset + "&limit=" + limit + "&page=" + page + "&tag=" + tag)
    },
    // 使用id获取问题详情
    problem_detail_by_id: (id) => {
        return Get("problem?problem_id=" + id)
    },
    //  管理员查询问题详情
    admin_problem_detail_by_id: (id) => {
        return Get("admin/problem?id=" + id)
    },
    // 查询自己的提交记录 
    submissions_myself: (page, limit, offset) => {
        selfLog("submissions?myself=1&result=&username=&page=" + page + "&limit=" + limit + "&offset=" + offset);
        return Get("submissions?myself=1&result=&username=&page=" + page + "&limit=" + limit + "&offset=" + offset)
    },
    // 获取验证码图片
    get_captcha_img: () => {
        return Get('captcha')
    },
    register_new: (param) => {
        return Post('register', param)
    },
    // 退出登录
    logout: () => {
        return Get('/logout')
    },
    // 修改个人信息（非账号信息）
    updateInfo: (params) => {
        return Put('profile', params)
    },
    // 上传头像
    uploadAvatar: (params) => {
        return Post('upload_avatar', params)
    },
    // 改密码
    change_password: (params) => {
        return Post('change_password', params)
    },
    // 改邮件
    change_email: (params) => {
        return Post('change_email', params)
    },
    // 获取登录记录
    login_history: () => {
        return Get('sessions')
    },
    // 提交代码
    submission: (params) => {
        return Post('submission', params)
    },
    // 查询提交在不在
    submission_exists: (problem_id) => {
        return Get('submission_exists?problem_id=' + problem_id)
    },
    // 查询提交的状态
    submission_id: (id) => {
        return Get('submission?id=' + id)
    },
    //上传测试用例
    upload_test_case: (params) => {
        return Post('admin/test_case', params)
    },
    //下载测试用例
    download_test_case: (id) => {
        return Get('admin/test_case?problem_id=' + id, { responseType: 'blob' })
    },
    // 作业列表
    // contests_list: (offset, limit) => {
    //     return Get('admin/problem?offset=' + offset + '&limit=' + limit)
    // },
    contests_list: (offset, limit) => {
        return Get('contests?offset=' + offset + '&limit=' + limit)
    },
    //老师获取作业列表
    admin_contests_list: (offset, limit) => {
        return Get('admin/contest?paging=true&offset=' + offset + '&limit=' + limit)
    },
    // 查看当前登录的用户有没有权限进入作业
    access_contest: (id) => {
        return Get('contest/access?contest_id=' + id)
    },
    // 进入作业列表
    enter_contest_list: (id) => {
        return Get('contest?id=' + id)
    },
    contest_passowrd_check: (params) => {
        return Post('contest/password', params)
    },
    // 查询作业的问题列表
    contest_problem_list: (contest_id) => {
        return Get('contest/problem?contest_id=' + contest_id)
    },
    // 查询作业问题的详情
    problem_contest: (problem_id, contest_id) => {
        return Get("contest/problem?contest_id=" + contest_id + "&problem_id=" + problem_id)
    },
    //老师拿作业列表
    admin_problem_list: (limit, offset) => {
        return Get("admin/problem?limit=" + limit + "&offset=" + offset)
    },
    // 查询作业的所有提交
    assign_submission_list (contest_id) {
        return Get("contest_submissions?myself=0&result=&username=&page=1&contest_id=" + contest_id + "&limit=12&offset=0")
    },
    // 普通用户拿公告信息
    obtain_announcement (offset, limit) {
        return Get('announcement?offset=' + offset + '&limit=' + limit)
    },
    // 管理员拿公告信息
    admin_obtain_announcement () {
        return Get('admin/announcement?paging=true&offset=0&limit=15')
    },
    // 管理员更新公告信息
    update_annocement (param) {
        return Put('admin/announcement', param)
    },
    // 管理员新增公告
    update_annocement_new (param) {
        return Post('admin/announcement', param)
    },
    // 管理员删除公告
    delete_anoncement (id) {
        return Delete('admin/announcement?id=' + id)
    },
    // 管理员拿系统信息
    dashboard_info () {
        return Get('admin/dashboard_info')
    },
    // 管理员拿用户列表
    obtain_users (offset, limit) {
        return Get('admin/user?paging=true&offset=' + offset + '&limit=' + limit)
    },
    // 管理员删除用户
    delete_user (id) {
        return Delete('admin/user?id=' + id)
    },
    // 管理员修改用户
    alerter_user (param) {
        return Put('admin/user', param)
    },
    // 拿用户的登录历史
    login_history () {
        return Get('sessions')
    },
    search_assign (offset, limit, params) {
        if (params.keyword !== "" && params.status !== "") {
            return Get('contests?offset=' + offset + '&limit=' + limit + '&status=' + params.status + '&keyword=' + params.keyword)
        }
        if (params.keyword === "" && params.status !== "") {
            return Get('contests?offset=' + offset + '&limit=' + limit + '&status=' + params.status)
        }
        if (params.keyword !== "" && params.status === "") {
            return Get('contests?offset=' + offset + '&limit=' + limit + '&keyword=' + params.keyword)
        }
        if (params.keyword === "" && params.status === "") {
            return Get('contests?offset=' + offset + '&limit=' + limit)
        }
    },
    // 模糊查询作业列表
    search_task (offset, limit, status, keyword) {
        return Get('admin/contest?paging=true&offset=' + offset + '&limit=' + limit + '&status=' + status + '&keyword=' + keyword)
    },
    // 获取全部状态的模糊查询
    search_task_ignore_status (offset, limit, keyword) {
        return Get('admin/contest?paging=true&offset=' + offset + '&limit=' + limit + '&keyword=' + keyword)
    },
    // 管理员删除问题
    delete_problem (id) {
        return Delete('/admin/problem?id=' + id)
    },
    // 管理员修改问题
    admin_update_problem (param) {
        return Put('/admin/problem', param)
    },
    // 管理员隐藏作业
    // 管理员创建作业
    create_task (params, method) {
        if (method === 'post') {
            // 新建作业
            return Post('admin/contest', params)
        }
        // 更新作业
        else if (method === 'put') {
            return Put('admin/contest', params)
        }

    },
    // 老师查询作业详情
    get_detail_assign_id (id) {
        return Get('admin/contest?id=' + id)
    },
    // 老师添加问题
    admin_add_problem (params) {
        return Post('admin/problem', params)
    },
    // 新增问题标签
    add_tags (keyword) {
        return Get('problem/tags?keyword=' + keyword)
    },
    //为作业添加题目
    add_problem (params) {
        return Post('admin/contest/add_problem_from_public', params)
    },
    //在作业中查找包含的问题
    search_problem_list (limit, offset, contest_id) {
        return Get('admin/contest/problem?limit=' + limit + '1&offset=' + offset + '&contest_id=' + contest_id)
    },
    // 教师删除作业中的问题
    delete_problem_from_assign (id) {
        return Delete('admin/contest/problem?id=' + id)
    },
    // 教师新增作业公告
    publish_annoncement_for_assign (param) {
        return Post('admin/contest/announcement', param)
    },
    // 查询作业中的公告列表
    get_annonce_for_contest (id) {
        return Get('admin/contest/announcement?contest_id=' + id)
    },
    // 学生查询公告列表
    get_annonce_for_contest_stu (id) {
        return Get('contest/announcement?contest_id=' + id)
    },
    // 教师删除公告
    delete_annonce (id) {
        return Delete('admin/contest/announcement?id=' + id)
    },
    // 获取站内排名数据
    get_rank (offset, limit) {
        return Get('user_rank?offset=' + offset + '&limit=' + limit + '&rule=ACM')
    },
    // 获取所有种类tags
    get_tags () {
        return Get('problem/tags')
    },
    // =========================================自己的Django==============================================
    // 分页获取弹幕
    get_bullet_chats: (params) => {
        return Post2('usr/get_bullet_chats/', params)
    },
    // 发送弹幕
    send_single_bullet_chat: (params) => {
        return Post2('usr/addbulletchat/', params)
    },
    // 使用学号登录
    stuLogin: (params) => {
        return Post2('usr/stu_login/', params)
    },
    // 获取天气情况
    wetherToday: () => {
        return Get2('dashboard/wetherToday/')
    },
    // 用学号拿真名
    get_real_name_by_stunum (stuNum) {
        return Post2('usr/get_real_name_by_stunum/', stuNum)
    },
    update_password_stu (form) {
        return Post2('usr/update_password_stu/', form)
    },


    contactUS (params) {
        return Post2('usr/contactus/', params)
    },
    // 心跳包监测
    heartbeat () {
        return Get2('common/heartbeat/')
    },
    // 验证码错误的时候重置密码为0 # 删除用户也需要重置为0
    reset_pass_to_zero (param) {
        return Post2('usr/reset_pass_to_zero/', param)
    },
    // 老师注册发邮件
    send_mail_to_teacher_register (param) {
        return Post2('usr/send_mail_to_teacher_register/', param)
    },
    // 老师注册申请
    apply_to_be_a_teacher (param) {
        return Post2('usr/apply_to_be_a_teacher/', param)
    },
    // 获取系统通知
    get_sys_info (userId) {
        return Post2('common/get_sys_info/', userId)
    },
    // 存系统通知send_sys_info
    send_sys_info (param) {
        return Post2('common/send_sys_info/', param)
    },
    // 读系统的消息 设为已读
    set_info_read (param) {
        return Post2('common/set_info_read/', param)
    },
    // 根据教师工号查询自己的课程
    get_teacher_courses (param) {
        return Post2('usr/get_teacher_courses/', param)
    },
    // 注册用户名的时候先看一下是不是学号
    check_username_is_stunum (username) {
        return Post2('usr/check_username_is_stunum/', username)
    },
    // 查学号注册了没
    check_if_stunum_registered (username) {
        return Post2('usr/check_if_stunum_registered/', username)
    },
    // 学号第一次登录之后把数据库中的学生对应的userid改成实际的userId
    update_new_stu_userId (stunum, userId) {
        let param = {
            stunum: stunum,
            userId: userId
        }
        return Post2("usr/update_new_stu_userId/", param)
    },
    // 拿所有没有通过审核的教师列表
    obtain_teacher_apply_list () {
        return Get2('usr/obtain_teacher_apply_list/')
    },
    // 拿所有被拒绝了的列表
    obtain_teacher_reject_list () {
        return Get2('usr/obtain_teacher_reject_list/')
    },
    // 审核通过教师
    pass_tea_check (param) {
        return Post2('usr/pass_tea_check/', param)
    },
    // 拒绝给予教师权限
    reject_tea_check (param) {
        return Post2('usr/reject_tea_check/', param)
    },
    // 删除用户
    deleteUsr (param) {
        return Post2('usr/deleteUsr/', param)
    },

    // 飞书登录插入新学生学号
    insert_feishu_stu (param) {
        return Post2("usr/insert_feishu_stu/", param)
    },
    // 拿test_case内容
    get_test_case (testcase_id) {
        let param = {
            testcase_id: testcase_id
        }
        return Post2("problem/get_test_case/", param)
    },
    // 修改某个test_case
    update_testcase (param) {
        return Post2("problem/update_testcase/", param)
    },
    // 查一个作业所有提交的代码重复率
    submisson_code_similarity (param) {
        return Post2("view_contest/submisson_code_similarity/", param)
    },
    // =========================================================学校的API===========================================
    weeklyclass (userid) {
        return Get3('getThisWeekClassSchedule?userId=' + userid)
    },
    getWeek () {
        return Get3('getWeek')
    },
    // 获取学号，工资号
    GetUserInfoByH5 (eid, code) {
        selfLog(code)
        return Get5("GetUserInfoByH5?eid=" + eid + "&code=" + code)
    },
    // =========================================================飞书API=============================================
    // 获取tenant_access_token
    get_tenant_access_token () {
        let param = {
            "app_id": "cli_a0f809e4ca78100e",
            "app_secret": "u9OWRYtZ1lvVFnVknciFgexszZKdaDBP"
        }
        return Post4("auth/v3/tenant_access_token/internal/", param, "")
    },
    // 获取用户信息
    get_FeiShu_user_profile (tenant_access_token, code) {
        let headers = {
            "Authorization": "Bearer " + tenant_access_token,
            "Content-ype": "application/json; charset=utf-8"
        }
        let param = {
            "grant_type": "authorization_code",
            "code": code
        }
        return Post4("authen/v1/access_token", param, headers)
    },
    // 获取JSAPI
    get_JSAPI (tenant_access_token_) {
        let param = {
            tenant_access_token: tenant_access_token_
        }
        return Post4("jssdk/ticket/get", param)
    },
    // 请求用户身份验证
    request_user_profile () {
        let app_id = "cli_a0f809e4ca78100e"
        // let REDIRECT_URI = UrlEncode("http://127.0.0.1/login")
        let REDIRECT_URI = "http%3A%2F%2Fwww.baidu.com"
        // http://tool.chinaz.com/tools/urlencode.aspx
        let url = "authen/v1/user_auth_page_beta?app_id=cli_a0f809e4ca78100e&redirect_uri=http%3A%2F%2F127.0.0.1%3A8080%2Flogin"
        selfLog(url)
        return url
    },
    // =========================================================工具类==============================================
    getCookie: (name) => {
        return getCookie(name)
    }

}
