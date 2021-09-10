<template>
	<dv-full-screen-container class="bg"
		v-bind:class="{ bg1: bg[0], bg2: bg[1], bg3: bg[2], bg4: bg[3],bg5: bg[4], bg6: bg[5],bg7: bg[6], bg8: bg[7],bg9: bg[8], bg10: bg[9],bg11: bg[10], bg12: bg[11]}">
		<div class="title">
			<dv-decoration-11 style="width:20%;height:10%;position:absolute;left:40%;top:5%;font-size:200%">
				<a href="/"><img src="../../public/sologon.png" style="width:100%;marginTop:10%" alt=""></a>
			</dv-decoration-11>
		</div>
		<!-- 普通注册/老师注册二选一 -->
		<div class="panel">
			<el-form label-position="left" label-width="80px" :model="form" :rules="rules" ref="ruleForm">
				<el-form-item label="用户名称">
					<el-input prefix-icon="el-icon-user" v-model="form.username" placeholder="请输入用户名称"></el-input>
				</el-form-item>
				<el-form-item label="登录密码" prop="password">
					<el-input type="password" v-model="form.password" placeholder="请输入登录密码" show-password
						prefix-icon="el-icon-lock"></el-input>
				</el-form-item>
				<el-form-item label="确认密码" prop="checkPass">
					<el-input type="password" v-model="form.checkPass" placeholder="请再次输入登录密码" show-password
						prefix-icon="el-icon-lock"></el-input>
				</el-form-item>
				<el-form-item label="邮箱" prop="email">
					<el-input prefix-icon="el-icon-s-data" v-model="form.email" placeholder="请输入邮箱"></el-input>
				</el-form-item>
				<el-form-item label="验证码">
					<el-input prefix-icon="el-icon-mobile" v-model="form.captcha_str" placeholder="请输入验证码"
						style="width:70%;">
					</el-input>
					<img :src="captcha" @click="get_captcha_img" class="captchaImg" />
				</el-form-item>
				<!-- <el-form-item label="注册类型" size="medium" style="width:290px">
					<el-radio-group v-model="form.registerType">
						<el-radio border label="学生">学生</el-radio>
						<el-radio border label="老师" @change="register_msg(form.registerType)">老师</el-radio>
					</el-radio-group>
				</el-form-item> -->
				<el-button style="width:80%;" type="primary" @click="submitForm('ruleForm')" :loading="form.loading">注册
				</el-button>
			</el-form>
			<el-link type="primary" @click="gotoLoginPage">已注册? 立即登录</el-link>
		</div>

		<!-- <footer>
      <p>xxxxx</p>
      </footer>-->
		<!-- </dv-border-box-3> -->
	</dv-full-screen-container>
</template>
<script>
	export default {
		name: 'Register',

		data() {
			let isEmail = s => {
				return /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,3}){1,2})$/.test(
					s
				)
			}
			let validateEmail = (rule, value, callback) => {
				if (!isEmail(value)) {
					callback(new Error('邮箱格式错误'))
				} else {
					callback()
				}
			}
			let validatePass = (rule, value, callback) => {
				if (value === '') {
					callback(new Error('请输入密码'))
				} else {
					if (this.form.checkPass !== '') {
						this.$refs.ruleForm.validateField('checkPass')
					}
					callback()
				}
			}
			let validatePass2 = (rule, value, callback) => {
				if (value === '') {
					callback(new Error('请再次输入密码'))
				} else if (value !== this.form.password) {
					callback(new Error('两次输入密码不一致!'))
				} else {
					callback()
				}
			}
			return {
				bg: [false, false, false, false, false, false, false, false, false, false, false, false],
				form: {
					username: '',
					password: '',
					checkPass: '',
					email: '',
					captcha_str: '',
					loading: false
				},
				rules: {
					password: [{
						validator: validatePass,
						trigger: 'blur'
					}],
					checkPass: [{
						validator: validatePass2,
						trigger: 'blur'
					}],
					email: [{
							required: true,
							message: '邮箱不能为空',
							trigger: 'blur'
						},
						{
							validator: validateEmail,
							trigger: 'blur'
						}
					]
				},
				captcha: '',
				captchaId: ''
			}
		},
		beforeMount() {
			let month = new Date().getMonth() //获取当前月份(0-11,0代表1月)
			this.bg[month] = true
			// 获取验证码图片
			this.get_captcha_img()
		},
		mounted() {

		},
		methods: {
			gotoLoginPage() {
				this.$router.push('/login')
			},
			get_captcha_img() {
				this.$captcha_axios.NewCaptId().then(res => {
					this.captcha = 'http://192.168.128.0:8666/captcha/' + res.CaptchaId + '.png'
					this.captchaId = res.CaptchaId
					this.selfLog(this.captchaId, '----', res.CaptchaId)
				})
			},
			// 先校验一把表单的字段合法性  和  验证码是不是对的 
			submitForm(formName) {
				this.form.loading = true
				this.$refs[formName].validate(valid => {
					if (valid) {
						this.selfLog({
							'captchaId': this.captchaId,
							'captchaSolution': this.form.captcha_str
						})
						this.$captcha_axios.VerifyCaptcha({
							'captchaId': this.captchaId,
							'captchaSolution': this.form.captcha_str
						}).then(res => {
							this.selfLog(res)
							if (res.errcode === 200) {
								this.register()
							} else if (res.errcode === 205) {
								this.$message({
									message: '验证码错误',
									type: 'warning'
								})
								this.form.loading = false
							}
						})
					} else {
						this.$message({
							message: '请检查字段合法性',
							type: 'warning'
						})
						this.form.loading = false
						return false
					}
				})
			},
			register(type) {
				this.get_captcha_img()
				this.$user_axios.Register({
					'UserName': this.form.username,
					'Password': this.form.password,
					'Email': this.form.email
				}).then(res => {
					this.selfLog(res)
					if (res.errcode === 205) {
						this.$message({
							message: '该用户名已经被注册',
							type: 'warning'
						})
					} else if (res.errcode === 206) {
						this.$message({
							message: '该email已经被注册',
							type: 'warning'
						})
					}
				})
				this.form.loading = false
			},
			register_msg(type) {
				if (type === '老师') {
					this.$notify({
						title: '审核通知',
						message: '尊敬的老师您好，注册老师账号需审核，请确保您的邮箱正确',
						type: 'warning'
					})
				}
			}
		}
	}
</script>

<style lang="scss" scoped>
	.bg {
		.title {
			height: 16%;
		}

		.panel {
			width: 25%;
			height: 55%;
			margin: auto;
			background-color: white;
			opacity: 0.95;
			padding: 18px;

			.captchaImg {
				width: 30%;
				// height: 20%;
				margin-top: 5px;
				margin-bottom: -15px;
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
		background: url("../assets/uestc/calendar/3.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg5 {
		background: url("../assets/uestc/calendar/5.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg6 {
		background: url("../assets/uestc/calendar/6.jpeg") no-repeat;
		background-size: 100% 100%;
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

	.bg1 {
		background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg2 {
		background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg1 {
		background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
		background-size: 100% 100%;
	}

	.bg2 {
		background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
		background-size: 100% 100%;
	}
</style>
