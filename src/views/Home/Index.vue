<template>
	<dv-full-screen-container>
		<el-container id="home">
			<el-container>
				<el-header id="header">
					<el-row>
						<!-- logo -->
						<el-col :span="6">
							<!-- <dv-decoration-1 :color="['white', '#112d4e']"
                style="width:100px;height:50px;position:absolute;left10px" /> -->
							<a href="/"><img src="../../../public/sologon.png" style="width:80%;margin-left:0%"
									alt=""></a>
							<div v-if="!maliao_back_falg">
								<img class="pic" src="../../assets/bg/maliao.gif"
									:style="'z-index:-2;position:absolute;top:'+(maliao_height-15)+'%;width:3%;margin-left:'+maliao_position+';'"
									alt="">
								<img src="../../assets/bg/floor.png"
									:style="'z-index:-1;position:absolute;left:'+-screenWidth*0.20+'px;top:10px;height:80%;width:'+screenWidth*1.8+'px;'"
									alt="">
							</div>
							<div v-else>
								<img class="pic" src="../../assets/bg/maliao_back.gif"
									:style="'z-index:-1;position:absolute;top:'+(maliao_height-15)+'%;width::3%;margin-left:'+maliao_position+';'"
									alt="">
								<img src="../../assets/bg/floor.png"
									:style="'z-index:-2;position:absolute;left:'+-screenWidth*0.20+'px;top:10px;height:80%;width:'+screenWidth*1.8+'px;'"
									alt="">
							</div>
						</el-col>
						<!-- 首页 -->
						<el-col :span="8">
							<el-button  v-if="this.$root.sys_status === 1" type="text" style="color:black">
								<i class="el-icon-cpu" @click="sys_info">服务器在线</i>
							</el-button>
							<el-button  v-if="this.$root.sys_status === -1" type="text" style="color:black">
								<i class="el-icon-cpu" @click="sys_info">OJ离线</i>
							</el-button>
							<el-button  v-if="this.$root.sys_status === 0" type="text" style="color:black">
								<i class="el-icon-cpu" @click="sys_info">服务器离线</i>
							</el-button>
							<el-button type="text" style="color:black">
								<i class="el-icon-cpu" @click="page_navigator('/home/nav1')">首页</i>
							</el-button>
							<el-button type="text" style="color:black">
								<i class="el-icon-cpu" @click="page_navigator('/home/nav2')">题库</i>
							</el-button>
							<el-button type="text" style="color:black">
								<i class="el-icon-cpu" @click="page_navigator('/home/nav3')">作业</i>
							</el-button>
							<el-button type="text" style="color:black">
								<i class="el-icon-cpu" @click="page_navigator('/home/nav4')">提交</i>
							</el-button>
							<el-button type="text" style="color:black">
								<i class="el-icon-cpu" @click="page_navigator('/home/nav5')">排名</i>
							</el-button>
						</el-col>
						<!-- 页面跳转 -->
						
						<!-- 天气 -->
						
						<el-col v-if="wether.basic" :span="4">
							<el-button type="text" style="color:black">
								<i  class="el-icon-location-information">{{ wether.basic.city }}</i>&nbsp;
								<!-- <i class="el-icon-timer">{{ wether.daily_forecast[0].date }}</i>&nbsp; -->
								<i
									class="el-icon-wind-power">{{ wether.daily_forecast[0].wind.dir }}&nbsp;{{ wether.daily_forecast[0].wind.spd }}</i>&nbsp;
								<i
									class="el-icon-sunny">{{ wether.daily_forecast[0].tmp.min }}­°C-{{ wether.daily_forecast[0].tmp.max }}­°C</i>&nbsp;
							</el-button>
						</el-col>
						<!-- 账户管理 -->
						<el-col :span="3" style="paddingTop:2px">
							<el-dropdown trigger="click" @command="handleCommand">
								<el-button type="text">
									<span>
										<div v-if="user_profile === null">
											<i class="el-icon-user-solid">&nbsp;<span
													style="color:black">账户管理</span></i>
										</div>
										<div v-else style="display:flex">
											<img :src="user_profile.avatar" style="width:30px" />
											<div v-if="user_profile.user.admin_type==='Admin'" style="paddingTop:6px">
												&nbsp;<span style="color:black">账户管理</span>&nbsp;<span
													style="color:black">{{real_name}}老师</span></div>
											<div v-else style="paddingTop:6px">&nbsp;<span
													style="color:black">账户管理</span>&nbsp;<span
													style="color:black">{{real_name}}</span>
											</div>
										</div>
									</span>
								</el-button>
								<el-dropdown-menu slot="dropdown">
									<div v-if="user_profile === null">
										<el-dropdown-item command="logIn">
											<span>登录</span>
										</el-dropdown-item>
										<el-dropdown-item command="register">
											<span>注册</span>
										</el-dropdown-item>
									</div>
									<div v-if="user_profile !== null">
										<el-dropdown-item command="user_setting">
											<span>个人信息</span>
										</el-dropdown-item>
										<el-dropdown-item command="weeklyclass">
											<span>本周课表</span>
										</el-dropdown-item>
									</div>
									<el-dropdown-item command="contact">
										<span>联系我们</span>
									</el-dropdown-item>
									<div v-if="user_profile !== null">
										<el-dropdown-item command="logout">
											<span>退出登录</span>
										</el-dropdown-item>
									</div>
								</el-dropdown-menu>
							</el-dropdown>
						</el-col>
						<!-- 教师入口 -->
						<el-col :span="2">
							<div v-if="user_profile!==null && user_profile.user.admin_type==='Admin'"
								style="marginTop:-3px">
								<el-button type="text" @click="gotoTeacherSide"><span style="color:black">教师入口</span>
								</el-button>
							</div>
							<div v-else-if="user_profile!==null && user_profile.user.admin_type==='Super Admin'"
								style="marginTop:-8px">
								<el-button type="text" @click="gotoAdminSide"><span style="color:black">管理入口</span>
								</el-button>
							</div>
						</el-col>
						<!-- 系统通知 -->
						<el-col :span="1.5" style="fontSize: 14px;color:black;marginTop: 18px;">
							<el-badge :value="sys_info_num" :max="99">
								<span style="color:black"><i class="el-icon-message" style="color:black"
										@click="showSysNotify">系统通知</i></span>
							</el-badge>
						</el-col>
					</el-row>
				</el-header>

				<el-main style="padding:0">
					<div id="main">
						<router-view />
					</div>
				</el-main>
			</el-container>
		</el-container>
		<!-- 联系我们抽屉 -->
		<el-drawer title="联系我们" :visible.sync="dialog" append-to-body direction="ltr" size="50%"
			custom-class="demo-drawer" ref="drawer">
			<el-popover style="position:absolute;right:100px;top:10px" placement="right-end" title="联系我们？" width="200"
				trigger="hover" content="我们的系统刚上线存在诸多问题，欢迎您向我们提出意见或者联系我们进行账号处理，如果需要回复请务必登录并更改默认邮箱">
				<el-button slot="reference"><span style="color:green;margin-left:10%"><i
							class=" el-icon-question"></i></span>
				</el-button>
			</el-popover>

			<div class="demo-drawer__content">
				<el-form :model="form">
					<el-form-item label="主题" :label-width="formLabelWidth" style="width:80%">
						<el-select v-model="form.subject" placeholder="请选择">
							<el-option v-for="item in form.options" :key="item.value" :label="item.label"
								:value="item.value">
							</el-option>
						</el-select>
						<!-- <el-input v-model="form.subject" autocomplete="off"></el-input> -->
					</el-form-item>
					<el-form-item label="内容" :label-width="formLabelWidth">
						<quill-editor :content="form.content" style="width:80%" />
						<!-- <textarea style="height:400px;width:80%" v-model="form.content" autocomplete="off"></textarea> -->
					</el-form-item>
					<el-form-item label="是否匿名" :label-width="formLabelWidth">
						<el-switch v-model="form.annon" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
					</el-form-item>
				</el-form>
				<div style="float:right;marginRight:20%">
					<el-button @click="cancelForm">取 消</el-button>
					<el-button type="primary" @click="handleClose2" :loading="loading">{{ loading ? '提交中 ...' : '确 定' }}
					</el-button>
				</div>
			</div>
		</el-drawer>
		<!-- 本周课表抽屉 -->
		<el-drawer title="本周课表" :visible.sync="class_visble" append-to-body direction="rtl" size="60%"
			custom-class="demo-drawer" ref="drawer">
			第{{thisweek}}周
			<div class="demo-drawer__content">
				<el-table :data="weekClass" stripe>
					<el-table-column property="c_name" label="课程" width="100"></el-table-column>
					<el-table-column property="kksd" label="开课周次"></el-table-column>
					<el-table-column property="xqj" label="日期" width="100">
						<template slot-scope="scope">星期{{scope.row.xqj}}</template>
					</el-table-column>
					<el-table-column property="ksjc" label="节次"></el-table-column>
					<el-table-column property="room_name" label="教室"></el-table-column>
					<el-table-column property="t_name" label="教师"></el-table-column>
				</el-table>
			</div>
		</el-drawer>
		<!-- 系统通知抽屉 -->
		<el-drawer title="系统通知" :visible.sync="sysyinfo_visble" append-to-body direction="rtl" size="25%"
			custom-class="demo-drawer" ref="drawer">
			<el-badge type="primary" :value="annoucement_num" :max="99">
				<span style="color:black"><i class="el-icon-message" style="color:black;"
						@click="show_announcement">公告</i></span>
			</el-badge>
			<el-divider direction="vertical"></el-divider>
			<el-badge v-if="user_profile!==null" :value="sys_info_num" :max="99">
				<span style="color:black"><i class="el-icon-message" style="color:black"
						@click="showSysNotify_button">系统信息</i></span>
			</el-badge>
			<div class="demo-drawer__content">
				<div v-if="sysinfo_annoucement==='sysinfo' && user_profile!==null">
					<div v-for="(info, index) of sysinfo_arr" :key="index" style="margin:10px;">
						<dv-border-box-8 :dur=6 style="height:100px"><br>&nbsp;&nbsp;{{info.msg}}
							<el-col :span="20">
								<dv-border-box-13>
									<el-button style="float: right; padding: 3px 0" type="primary"
										@click="set_info_read(info.id)">已读
									</el-button>
								</dv-border-box-13>
								<!-- <el-card shadow="hover">
                {{info.msg}}
                <el-button style="float: right; padding: 3px 0" type="text" @click="set_info_read(info.id)">已读
                </el-button>
              </el-card> -->
							</el-col>
						</dv-border-box-8>
					</div>
				</div>
				<div v-else>
					<div v-for="(info, index) of annoncement_arr" :key="index" style="margin:10px;">
						<el-col :span="22">
							<el-card class="box-card" shadow="always">
								<div slot="header" class="clearfix">
									<span>{{info.title}}</span>
								</div>
								<div v-html="info.content" style="color:green"></div>
							</el-card>
						</el-col>
					</div>
				</div>
			</div>
		</el-drawer>
	</dv-full-screen-container>
</template>
<script type="text/javascript"
	src="https://sf1-scmcdn-tos.pstatp.com/goofy/ee/lark/h5jssdk/lark/js_sdk/h5-js-sdk-1.5.2.js">
	window.h5sdk.ready(() => {
		//  lark.ready参数为回调函数，在环境准备就绪时触发
	});
</script>
<script>
	export default {
		name: "Index",
		data() {
			return {
				// 学号
				stuNum: "",
				dialog: false,
				loading: false,
				real_name: "",
				OJIP: this.$OJIP,
				user_profile: null,
				wether: {},
				weekClass: [],
				thisweek: "1",
				class_visble: false,
				form: {
					annon: false,
					subject: "",
					content: "填写您的意见~",
					options: [{
						value: '选项1',
						label: '账号修改'
					}, {
						value: '选项2',
						label: '功能建议'
					}, {
						value: '选项3',
						label: 'bug上传'
					}, {
						value: '选项4',
						label: '其他~'
					}],
					value: ''
				},
				formLabelWidth: "80px",
				//   系统通知抽屉可见性
				sysyinfo_visble: false,
				//   系统通知数组
				sysinfo_arr: [1, 2, 3],
				sys_info_num: 0,
				annoncement_arr: [],
				// 系统信息抽屉显示公告或者系统信息
				sysinfo_annoucement: "sysinfo",
				annoucement_num: 0,
				maliao_position: "0px",
				maliao_position_int: -20,
				maliao_back_falg: false,
				maliao_height: 10,
				//   屏幕大小
				screenWidth: document.body.clientWidth,
				windowHeight: document.documentElement.clientHeight
			};
		},
		mounted() {
			//   屏幕大小
			const that = this
			window.onresize = () => {
				return (() => {
					window.screenWidth = document.body.clientWidth
					that.screenWidth = window.screenWidth
					that.windowHeight = window.fullHeight;
					this.selfLog("当前屏幕大小：" + this.screenWidth)
				})()
			}

		},
		// 添加钩子函数在页面之前把弹幕加载进来
		beforeMount () {
			// 查天气
			this.wetherToday();
			this.move_maliao()
		},
		methods: {
			move_maliao() {
				setInterval(() => {
					if (this.maliao_position_int === 1000) {
						this.maliao_back_falg = true
					} else if (this.maliao_position_int === -20) {
						this.maliao_back_falg = false
					}
					// 往回走
					if (this.maliao_back_falg === true) {
						if (this.maliao_position_int < 640 && this.maliao_position_int >= 580) {
							this.maliao_height = this.maliao_height - 1
						}
						if (this.maliao_position_int < 580 && this.maliao_position_int >= 520) {
							this.maliao_height = this.maliao_height + 1
						}
						this.maliao_position_int = this.maliao_position_int - 1
						this.maliao_position = this.maliao_position_int + "" + "px"
					}
					// 正向走
					else {
						if (this.maliao_position_int >= 520 && this.maliao_position_int < 580) {
							this.maliao_height = this.maliao_height - 1
						}
						if (this.maliao_position_int >= 580 && this.maliao_position_int < 640) {
							this.maliao_height = this.maliao_height + 1
						}
						this.maliao_position_int = this.maliao_position_int + 1
						this.maliao_position = this.maliao_position_int + "" + "px"
					}
				}, 20)
			},
			gotoTeacherSide() {
				this.$router.push("/teacher");
			},
			gotoAdminSide() {
				this.$router.push("/admin");
			},
			handleOpen(key, keyPath) {
				this.selfLog(key, keyPath);
			},
			handleClose(key, keyPath) {
				this.selfLog(key, keyPath);
			},
			cancelForm() {
				this.loading = false;
				this.dialog = false;
				clearTimeout(this.timer);
			},
			logOut() {
				this.$confirm("确认退出？").then(() => {
					this.user_profile = null;
					this.$axios.logout();
					this.$router.push({
						path: "/login"
					});
				});
			},
			logIn() {
				// 最好做一个弹窗
				this.$router.push("/login");
			},
			user_setting() {
				this.$router.push("/user/profile");
			},
			handleCommand(command) {
				if (command === "logout") this.logOut();
				else if (command === "logIn") {
					this.logIn();
				} else if (command === "user_setting") {
					this.user_setting();
				} else if (command === "contact") {
					this.dialog = true;
				} else if (command === "weeklyclass") {
					this.weeklyclass();
				} else if (command === "register") {
					this.$router.push("/register");
				}
			},
			wetherToday() {
				this.$utils_axios.WeatherToday().then(res => {
					this.selfLog(res.data.result.HeWeather5[0]);
					this.wether = res.data.result.HeWeather5[0];
				})
			},
			handleClose2(done) {
				if (this.loading) {
					return;
				}
				this.$confirm("确定要提交表单吗？")
					.then(_ => {
						this.loading = true;
						let mail = "";
						let userid = "";
						// 匿名提交
						if (this.form.annon) {
							mail = "anonymous@163.com";
							userid = -1;
						} else if (!this.form.annon && this.user_profile === null) {
							this.$message({
								message: "实名建议请先登录",
								type: "warning"
							});
							return false;
						} else if (this.user_profile !== null && !this.form.annon) {
							// 检查eamil是不是默认的 默认的不发邮件
							if (
								this.user_profile.user.email ===
								this.user_profile.user.username + "@oj.org"
							) {
								this.$message({
									message: "您的注册邮箱尚未更新，请修改以便接收系统通知",
									type: "warning"
								});
								email = "";
							} else {
								mail = this.user_profile.user.email;
							}
							userid = this.user_profile.id;
						}
						let params = {
							subject: this.form.subject,
							content: this.form.content,
							mail: mail,
							userid: userid
						};
						this.$axios.contactUS(params).then(res => {
							const h = this.$createElement;
							this.$notify({
								title: "提交成功",
								message: h(
									"i", {
										style: "color: teal"
									},
									"感谢您的建议，我们将尽快进行改进"
								)
							});
						});
						this.timer = setTimeout(() => {
							done();
							// 动画关闭需要一定的时间
							setTimeout(() => {
								this.loading = false;
							}, 200);
						}, 500);
					})
					.catch(_ => {});
			},
			sys_info() {
				if (this.$root.sys_status === 1) {
					this.$notify({
						title: "系统正常",
						message: "服务器在线运行，祝您使用愉快~",
						type: "success"
					});
				} else if (this.$root.sys_status === 0) {
					this.$notify({
						title: "系统异常",
						message: "啊嘞，服务器不在线，请联系管理员~",
						type: "warning"
					});
				} else if (this.$root.sys_status === -1) {
					this.$notify({
						title: "系统异常",
						message: "啊嘞，OJ服务器不在线，请联系管理员~",
						type: "warning"
					});
				}
			},
			// 获取本周课表
			weeklyclass() {
				this.$axios.weeklyclass(this.stuNum).then(res => {
					for (let course of res) {
						course.ksjc = course.ksjc + "-" + course.jsjc + "节";
					}
					this.weekClass = res;
					this.$axios.getWeek().then(res => {
						this.thisweek = res;
						this.class_visble = true;
					});
					// this.weekClass = res
				});
			},
			//展示系统通知
			showSysNotify() {
				this.sysyinfo_visble = true;
				let params = {
					userId: this.user_profile.id
				};
				this.show_announcement()
				this.$axios.get_sys_info(params).then(res => {
					this.selfLog(res);
					this.sysinfo_arr = res.data;
					this.sys_info_num = res.data.length;
				});
			},
			// 已读系统消息
			set_info_read(read_info_id) {
				let param = {
					info_id: read_info_id
				};
				this.$axios.set_info_read(param).then(res => {
					this.showSysNotify();
				});
				this.selfLog(read_info_id);
			},
			// 设置系统显示的信息抽屉显示系统信息
			showSysNotify_button() {
				this.sysinfo_annoucement = "sysinfo"
			},
			show_announcement() {
				this.$axios.obtain_announcement(0, 10).then(res => {
					this.selfLog(res);
					this.annoncement_arr = res.data.results
					this.annoucement_num = this.annoncement_arr.length
				})
				this.sysinfo_annoucement = "annoucement"
			},
			// 页面导航
			page_navigator(page){
				this.selfLog(page)
				this.$router.push(page);
			}
		},
		watch: {},
		computed: {}
	};
</script>


<style lang='scss' scoped>
	#home {
		#header {
			// background-color: rgba(191, 248, 245, 0.2);
			background-color: transparent;
			z-index: 5;
			font-size: 40px;

			.header-logo {
				height: 50px;
				text-align: center;
				line-height: 50px;
			}

			.el-icon-message {
				cursor: pointer;
				color: #409eff;
			}
		}

		#aside {
			height: 1500px;
			background-color: #112d4e;

			.menuNav {
				text-align: left;
			}
		}

		#main {
			height: 100%;
		}

		.item {
			margin-top: 10px;
			margin-right: 40px;
		}
	}
</style>
