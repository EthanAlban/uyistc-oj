<template>
	<dv-full-screen-container class="bg"
		v-bind:class="{ bg1: bg[0], bg2: bg[1], bg3: bg[2], bg4: bg[3],bg5: bg[4], bg6: bg[5],bg7: bg[6], bg8: bg[7],bg9: bg[8], bg10: bg[9],bg11: bg[10], bg12: bg[11]}">
		<div class="title">
			<dv-decoration-11 style="width:20%;height:10%;position:absolute;left:40%;top:5%;font-size:200%">
				<a href="/"><img src="../../public/sologon.png" style="width:100%;marginTop:10%" alt=""></a>
			</dv-decoration-11>
		</div>
		<div class="third_login">
			<el-tooltip class="item" effect="light" content="成电飞书登录" placement="right">
				<img src="../icons/FeiShuLogo.png" class="FeiShu" @click="FeiShu_authorize" alt="">
			</el-tooltip>
			<el-tooltip class="item" effect="light" content="微博登录" placement="right">
				<img src="../icons/WeiBoLogo.png" class="WeiBo" @click="FeiShu_authorize" alt="">
			</el-tooltip>
		</div>
		<!-- 学号登录/普通账号登录二选一 -->
		<el-tabs type="border-card" class="panel">
			<el-tab-pane label="用户登录">
				<el-form label-position="left" label-width="80px" :model="form">
					<el-form-item label="用户名称">
						<el-input prefix-icon="el-icon-user" v-model="form.username" placeholder="请输入用户名称"></el-input>
					</el-form-item>
					<el-form-item label="登录密码">
						<el-input prefix-icon="el-icon-lock" type="password" v-model="form.password"
							placeholder="请输入登录密码" show-password></el-input>
					</el-form-item>
					<el-button style="width:80%" type="primary" round @click="login" :loading="loading">登录</el-button>
				</el-form>
			</el-tab-pane>


			<el-link type="primary" @click="register" style="float:left;margin:20px 0 0 20px">新用户? 立即注册</el-link>
			<el-link type="primary" @click="visitor_login" style="float:right;margin:20px 20px 0 0">游客登录</el-link>

		</el-tabs>
	</dv-full-screen-container>
</template>

<script>
	import axios from 'axios'
	export default {
		name: 'Login',
		beforeMount() {
			let month = new Date().getMonth() //获取当前月份(0-11,0代表1月)
			// month = 8
			this.bg[month] = true
		},
		mounted() {
			let code = this.$route.query.code
			this.code = code
		},
		data() {
			return {
				bg: [false, false, false, false, false, false, false, false, false, false, false, false],
				loading: false,
				stu_not_registered: true,
				form: {
					username: '',
					password: '',
					captcha_str: ''
				},
				captcha: '',
				// 登录预授权码
				code: '',
				FeiShu_login_check: false,
				//   飞书登录时等待网络请求loading
				FeiShu_login_loadin: false,
				formLabelWidth: '80px',
				FeiShu: {
					tenant_access_token: '',
					app_access_token: ''
				},
				is_new_stu: false
			}
		},


		methods: {
			FeiShu_authorize() {
				window.location.href =
					'https://open.feishu.cn/open-apis/authen/v1/user_auth_page_beta?app_id=cli_a0f809e4ca78100e&redirect_uri=http%3A%2F%2F47.94.135.51%3A8080%2Flogin&state=login'
			},
			//用户登录
			login() {
				this.$user_axios.Login({
					'UserName': this.form.username,
					'Password': this.form.password
				}).then(res => {
					this.selfLog(res)
					if (res.errcode === 205) {
						this.$message.error('用户名密码不匹配，请重试~')
					} else if (res.errcode === 204) {
						this.$message({
							message: '用户名未注册，请检查',
							type: 'warning'
						})
					} else if (res.errcode === 200) {
						this.$message({
							message: '登录成功~',
							type: 'success'
						})
						this.$router.push('/')
					}
				}).catch(err=>{
					this.$message({
						message: '登录失败请稍后重试~',
						type: 'warning'
					})
				})
			},
			register() {
				this.$router.push('/register')
			},
			visitor_login() {
				this.$router.push('/')
			},
		}
	}
</script>

<style lang="scss" scoped>
	.bg {
		.register {
			float: right;
			top: 2px;
			right: 2px;
		}

		.title {
			height: 16%;

			.sys-name {
				font-size: 45px;
				color: #f5f5f5;
			}
		}

		.panel {
			width: 25%;
			height: 40%;
			margin: auto;
			background-color: white;
			opacity: 0.95;

			.captchaImg {
				width: 40%;
				height: 40px;
				margin-bottom: -15px;
			}

			.loginButton {
				width: 200px;
			}
		}
	}

	footer {
		position: absolute;
		bottom: 10%;

		@media screen and (max-width: 1440px) and (min-width: 1280px) {
			bottom: 16%;
		}

		display: flex;
		width: 100%;
		justify-content: center;

		p {
			font-size: 30px;
		}
	}

	.bg1 {
		background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg2 {
		background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg3 {
		background: url("../assets/uestc/calendar/3.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg4 {
		background: url("../assets/uestc/calendar/4.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg5 {
		background: url("../assets/uestc/calendar/5.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg6 {
		//   background: url("../assets/uestc/calendar/6.jpeg") no-repeat;
		background: url("../assets/bg/drangon_boat.png") no-repeat;
		background-size: 100% 90%;
	}

	.bg7 {
		background: url("../assets/uestc/calendar/7.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg8 {
		background: url("../assets/uestc/calendar/8.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg9 {
		background: url("../assets/uestc/calendar/9.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg10 {
		background: url("../assets/uestc/calendar/10.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg11 {
		background: url("../assets/uestc/calendar/11.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg12 {
		background: url("../assets/uestc/calendar/12.jpeg") no-repeat;
		background-size: 100% 100%;
	}



	.third_login {
		margin-top: 20%;
		position: absolute;
		right: 32%;
		top: -15%;
		width: 6%;

		.FeiShu {
			position: felx;
			background-image: url("../icons/FeiShuLogo.png");
			width: 51%;
			height: 51%;
			border-radius: 50%;
			box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
			-webkit-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
			-moz-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
		}

		.WeiBo {
			position: felx;
			margin-left: 5%;
			background-image: url("../icons/FeiShuLogo.png");
			width: 51%;
			height: 51%;
			border-radius: 50%;
			box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
			-webkit-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
			-moz-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
		}
	}
</style>
