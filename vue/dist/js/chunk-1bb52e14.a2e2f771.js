(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1bb52e14"],{"14a9":function(t,e,s){},a38b:function(t,e,s){"use strict";s("14a9")},b413:function(t,e,s){"use strict";s.r(e);var a=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("el-form",{staticClass:"demo-form-inline",staticStyle:{margin:"10px 0 0 20px",postion:"flex",float:"left"},attrs:{inline:!0,model:t.formInline}},[s("el-form-item",{attrs:{label:"作业查询"}},[s("el-input",{attrs:{placeholder:"输入关键字"},model:{value:t.formInline.keyword,callback:function(e){t.$set(t.formInline,"keyword",e)},expression:"formInline.keyword"}})],1),s("el-form-item",{attrs:{label:""}},[s("el-select",{attrs:{placeholder:"作业选项"},model:{value:t.formInline.status,callback:function(e){t.$set(t.formInline,"status",e)},expression:"formInline.status"}},[s("el-option",{attrs:{label:"全部",value:""}}),s("el-option",{attrs:{label:"正在进行",value:"0"}}),s("el-option",{attrs:{label:"尚未开始",value:"1"}}),s("el-option",{attrs:{label:"已经结束",value:"-1"}})],1)],1),s("el-form-item",[s("el-button",{attrs:{type:"primary"},on:{click:t.search_task}},[t._v("查询")])],1)],1),s("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData}},[s("el-table-column",{attrs:{type:"expand"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-form",{staticClass:"demo-table-expand",attrs:{"label-position":"left",inline:""}},[s("el-form-item",{attrs:{label:"发布时间"}},[s("span",[t._v(t._s(e.row.create_time))])]),s("el-form-item",{attrs:{label:"开始时间"}},[s("span",[t._v(t._s(e.row.start_time))])]),s("el-form-item",{attrs:{label:"最近修改时间"}},[s("span",[t._v(t._s(e.row.last_update_time))])])],1)]}}])}),s("el-table-column",{attrs:{label:"作业",prop:"title"}}),s("el-table-column",{attrs:{label:"作业描述",prop:"description"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("div",{domProps:{innerHTML:t._s(e.row.description)}})]}}])}),s("el-table-column",{attrs:{label:"截止时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("i",{staticClass:"el-icon-time"}),s("span",{staticStyle:{"margin-left":"10px"}},[t._v(t._s(t._f("formatDate")(e.row.end_time)))])]}}])}),s("el-table-column",{attrs:{label:"创建人",prop:"created_by.username"}}),s("el-table-column",{attrs:{label:"可见状态",prop:"visible"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949"},on:{change:function(s){return t.hide_task(e.row)}},model:{value:e.row.visible,callback:function(s){t.$set(e.row,"visible",s)},expression:"scope.row.visible"}})]}}])}),s("el-table-column",{attrs:{label:"管理"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-button",{attrs:{type:"text",size:"small"},on:{click:function(s){return t.EnterContest(e.row.id)}}},[t._v("查看")]),s("el-button",{attrs:{type:"text",size:"small"}},[t._v("编辑")])]}}])})],1),s("el-dialog",{attrs:{title:"验证",visible:t.dialogFormVisible,"append-to-body":""},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[s("i",{staticClass:"el-icon-lock"}),s("el-form",{attrs:{model:t.form}},[s("el-form-item",{attrs:{label:"作业密码","label-width":t.formLabelWidth}},[s("el-input",{attrs:{type:"password",autocomplete:"off"},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1)],1),s("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[s("el-button",{on:{click:function(e){t.dialogFormVisible=!1}}},[t._v("取 消")]),s("el-button",{attrs:{type:"primary"},on:{click:t.EnterContest_password}},[t._v("确 定")])],1)],1),s("el-drawer",{attrs:{title:"作业题目列表",visible:t.table,direction:"rtl",size:"50%","append-to-body":""},on:{"update:visible":function(e){t.table=e}}},[s("el-table",{staticStyle:{width:"100%"},attrs:{data:t.problems}},[s("el-table-column",{attrs:{type:"expand"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-form",{staticClass:"demo-table-expand",attrs:{"label-position":"left",inline:""}},[s("el-form-item",{attrs:{label:"习题描述"}},[s("div",{staticStyle:{color:"blue"},domProps:{innerHTML:t._s(e.row.description)}})]),s("el-form-item",{attrs:{label:"可使用语言"}},t._l(e.row.languages,(function(e,a){return s("div",{key:a,staticStyle:{color:"blue"}},[t._v(t._s(e)+"、")])})),0),s("el-form-item",{attrs:{label:""}},[s("el-button",{attrs:{type:"primary",plain:""},on:{click:function(s){return t.doProblem(e.row._id)}}},[t._v("完成习题")])],1)],1)]}}])}),s("el-table-column",{attrs:{label:"ID",prop:"_id"}}),s("el-table-column",{attrs:{label:"标题",prop:"title"}}),s("el-table-column",{attrs:{label:"已完成人数",prop:"accepted_number"}})],1)],1),s("div",{staticClass:"block"},[s("el-pagination",{attrs:{"current-page":t.page,"page-sizes":[10,20,40,80],"page-size":100,layout:"total, sizes, prev, pager, next, jumper",total:t.total},on:{"size-change":t.handleSizeChange,"current-change":t.handleCurrentChange}})],1)],1)},l=[],o={beforeMount:function(){var t=this;this.Get_Contests_List(),this.$axios.user_profile().then((function(e){t.selfLog(e),t.user_profile=e["data"],t.user_profile.avatar=t.$OJIP+t.user_profile.avatar,t.selfLog(t.user_profile)}))},data:function(){return{user_profile:null,page:1,offset:0,limit:10,total:0,contest_id:0,tableData:[],dialogFormVisible:!1,formLabelWidth:"80px",form:{password:""},formInline:{keyword:"",status:""},problems:{},table:!1}},methods:{Get_Contests_List:function(){var t=this;this.$axios.admin_contests_list(this.offset,this.limit).then((function(e){t.total=e.data.total,t.selfLog(e),t.tableData=e.data.results}))},handleSizeChange:function(t){this.limit=t,this.offset=(this.page-1)*this.limit,this.Get_Contests_List(),this.selfLog("每页 ".concat(t," 条"))},handleCurrentChange:function(t){this.page=t,this.offset=(this.page-1)*this.limit,this.Get_Contests_List(),this.selfLog("当前页: ".concat(t))},EnterContest:function(t){var e=this;if(null===this.user_profile)return this.$message({message:"查看习题请先登录",type:"warning"}),!1;this.selfLog(t),this.contest_id=t,this.$axios.access_contest(t).then((function(t){e.selfLog(t),t.data.access?(e.$message({message:"拥有权限，可进入",type:"success"}),e.$axios.contest_problem_list(e.contest_id).then((function(t){e.selfLog(t),"Contest has not started yet."===t.data?e.$message({message:"作业尚未到达开始时间",type:"warning"}):(e.problems=t.data,e.dialogFormVisible=!1,e.table=!0)}))):(e.$message({message:"暂无权限，请输入密码",type:"warning"}),e.dialogFormVisible=!0)}))},EnterContest_password:function(t){var e=this;if(null===this.user_profile)return this.$message({message:"查看习题请先登录",type:"warning"}),!1;if(""===this.form.password)return this.$message({message:"密码不能为空",type:"warning"}),!1;var s={contest_id:this.contest_id,password:this.form.password};this.$axios.contest_passowrd_check(s).then((function(t){e.selfLog(t),"Wrong password or password expired"===t.data?e.$message({message:"密码错误或者密码已过期",type:"warning"}):null===t.error&&e.$axios.contest_problem_list(e.contest_id).then((function(t){"Contest has not started yet."===t.data?e.$message({message:"作业尚未到达开始时间",type:"warning"}):(e.selfLog(t),e.problems=t.data,e.dialogFormVisible=!1,e.table=!0)}))}))},doProblem:function(t){this.selfLog(t),this.$router.push("/problemDetail_Contest/".concat(t,"/").concat(this.contest_id))},search_task:function(){var t=this;""===this.formInline.status?this.$axios.search_task_ignore_status(this.offset,this.limit,this.formInline.keyword).then((function(e){t.total=e.data.total,t.tableData=e.data.results,t.selfLog(e)})):this.$axios.search_task(this.offset,this.limit,this.formInline.status,this.formInline.keyword).then((function(e){t.total=e.data.total,t.tableData=e.data.results,t.selfLog(e)}))},hide_task:function(t){var e=this,s={allowed_ip_ranges:t.allowed_ip_ranges,contest_type:t.contest_type,create_time:t.create_time,created_by:t.created_by,description:t.description,end_time:t.end_time,id:t.id,last_update_time:t.last_update_time,password:t.password,real_time_rank:t.real_time_rank,rule_type:t.rule_type,start_time:t.start_time,status:t.status,title:t.title,visible:t.visible};this.$axios.hide_task(s).then((function(t){e.selfLog(t),null===t.error&&e.$message({message:"修改成功",type:"success"})}))}},computed:{}},i=o,n=(s("a38b"),s("2877")),r=Object(n["a"])(i,a,l,!1,null,"2bc8f116",null);e["default"]=r.exports}}]);
//# sourceMappingURL=chunk-1bb52e14.a2e2f771.js.map