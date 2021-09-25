<template>
	<div id="main">
		<div class="q_desc">
			<div class="header">
				<div class="title">{{ problbemDetail.Pid }}. {{ problbemDetail.Title }}</div>
				<div style="display:flex;">
					<p>难度</p>
					<p style="marginLeft:10px"
						:class="{low:problbemDetailDifficulty==='简单',mid:problbemDetailDifficulty==='中等',high:problbemDetailDifficulty==='困难'}">
						{{ problbemDetailDifficulty }}
					</p>
					<el-tag style="margin:10px 0 10px 10px" :type="tag_types[idx%5]" v-for="(tag,idx) in tags"
						:key="tag">{{tag}}</el-tag>
				</div>
			</div>
			<el-divider>
				<i class="el-icon-sunny"></i>
			</el-divider>
			<div class="body">
				<div v-html="problbemDetail.Content"></div>
				<span style="fontWeight:bold">输入描述:</span>
				<div v-html="problbemDetail.InputDescription"></div>
				<span style="fontWeight:bold">输出描述:</span>
				<div v-html="problbemDetail.OutputDescription"></div>
				<span style="fontWeight:bold">提示:</span>
				<div v-html="problbemDetail.Hint"></div>
				<div v-for="sample in samples" :key="sample.id">
					<div>
						<p style="fontWeight:bold">示例{{ sample.id }}:</p>
						<div class="sampleBox">
							<div class="input">输入:{{ sample.input }}</div>
							<div class="output">输出:{{ sample.output }}</div>
						</div>
					</div>
				</div>
				<div style="marginTop:20px">
					<span>通过次数:{{ acceptSub }}</span>
					<el-divider direction="vertical"></el-divider>
					<span>提交次数:{{ totalSub }}</span>
					<!-- 使用echarts画图 -->
					<div id="echart-pass-rate"></div>
				</div>
			</div>
		</div>
		<div class="ans_editor">
			<div class="demo-inline">
				<p class="title">语言</p>
				<el-select v-model="language" placeholder="请选择" @change="changeModel" class="selectLan">
					<el-option v-for="item in problbemDetail.languages" :key="item" :label="item" :value="item">
					</el-option>
				</el-select>
				<p class="title">主题</p>
				<el-select v-model="theme" placeholder="请选择" class="selectTheme">
					<el-option v-for="item in themes" :key="item" :label="item" :value="item"></el-option>
				</el-select>
				<el-button type="primary" round @click="recoverInit">恢复模板</el-button>
			</div>

			<!-- <codemirror class="codeEditor" ref="cmEditor" :value="code" :options="cmOptions" @ready="onCmReady"
				@focus="onCmFocus" @input="onCmCodeChange" /> -->
			<codemirror class="codeEditor" ref="cmEditor" :value="code" :options="cmOptions" @input="onCmCodeChange" />

		</div>
		<!-- 提交固定于页面右下角 -->
		<!-- 提交结果以抽屉展示 从下往上-->
		<el-button class="submit_fixed" type="success" round @click="submission">提交</el-button>
		<!-- <div style="position:absolute;right: 0;bottom: 0;height: 40%;width: 80%;background-color: #008855;z-index: 998;"  ></div> -->
		<!-- -2编译错误 -1答案错误 0正确 1计算超时  2超时 3内存超过 4运行时错误 5传送...  6判题中...  7部分正确 -->
		<el-drawer title="提交结果" :visible.sync="drawer" append-to-body destroy-on-close direction="btt"
			:before-close="handleClose" style="z-index: 9;width:80%;marginLeft:20%" :modal-append-to-body="false"  v-loading="drawer_loading"
			element-loading-text="拼命判题中" element-loading-spinner="el-icon-loading">
			<el-tag style="marginLeft:5%" v-if="submission_status==='编译错误'" type="danger" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='答案错误'" type="danger" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='正确'" type="success" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='计算超时'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='超时'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='内存超过'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='运行时错误'" type="success" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='传送...'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='判题中...'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag style="marginLeft:5%" v-if="submission_status==='服务器响应超时请重试提交...'" type="info" effect="dark">
				{{submission_status}}
			</el-tag>
			<el-tag v-if="submission_status==='部分正确'" type="warning" effect="dark">{{submission_status}}</el-tag>
			<el-tag style="marginLeft:5%" type="success" effect="dark">
				得分：{{score}}
			</el-tag>
			<el-tag style="marginLeft:5%" type="info" effect="dark">
				最后测试用例：{{LastTestcase}}
			</el-tag>
			<el-tag style="marginLeft:5%" type="info" effect="dark">
				正确输出： {{LastDesireOutput}}
			</el-tag>
			<el-tag style="marginLeft:5%" type="info" effect="dark">
				代码输出： {{LastOutput}}
			</el-tag>
			<br />
			<span style="color:red">{{err_info}}</span>
		</el-drawer>
	</div>
</template>

<script>
	import * as echarts from 'echarts'
	// import language js
	import 'codemirror/mode/clike/clike.js'
	import 'codemirror/mode/go/go.js'
	import 'codemirror/mode/python/python.js'
	// import theme style
	import 'codemirror/theme/eclipse.css'
	import 'codemirror/theme/idea.css'
	import 'codemirror/theme/material.css'
	import 'codemirror/theme/monokai.css'
	import 'codemirror/theme/tomorrow-night-bright.css'
	import 'codemirror/theme/base16-dark.css'

	export default {
		name: 'problemDetail',
		mounted() {
			this.getData()
			this.setEchart()
		},
		data() {
			return {
				drawer_loading: false,
				// 初次加载代码标记
				InitLoadCode: false,
				//  重复提交代码计数器
				submissionReply: 0,
				tag_types: ['info', 'warning', 'danger', 'success'],
				drawer: false,
				problbemDetail: {},
				acceptSub: 0,
				totalSub: 0,
				statistic_info: [],
				tags: [],
				msg: '',
				samples: [],
				language: 'C',
				theme: 'monokai',
				themes: [
					'eclipse',
					'idea',
					'material',
					'monokai',
					'tomorrow-night-bright',
					'base16-dark'
				],
				code: '',
				cmOptions: {
					tabSize: 4,
					mode: 'text/x-csrc',
					theme: 'monokai',
					lineNumbers: true,
					line: true,
					smartIndent: true,
					indentUnit: 2
					// more CodeMirror options...
				},
				submission_id: '',
				setInterval_id: '',
				submission_status: '',
				err_info: '',
				score: '',
				LastTestcase: '',
				LastDesireOutput: '',
				LastOutput: '',
				option: {
					tooltip: {
						trigger: 'item'
					},
					legend: {
						top: '5%',
						left: 'center'
					},
					series: [{
						name: '提交分析',
						type: 'pie',
						radius: ['40%', '45%'],
						avoidLabelOverlap: false,
						itemStyle: {
							borderRadius: 10,
							borderColor: '#fff',
							borderWidth: 2
						},
						label: {
							position: 'outside',
							show: true,
							fontSize: '12',
							fontWeight: 'bold',
							formatter: '{b}:{c}'
						},
						labelLine: {
							show: true,
							length: 5
						},
						data: []
					}]
				}
			}
		},
		methods: {
			setEchart() {
				var chartDom = document.getElementById('echart-pass-rate');
				var myChart = echarts.init(chartDom);
				var option;

				option = {
					tooltip: {
						trigger: 'item'
					},
					legend: {
						top: '5%',
						left: 'center'
					},
					color: [],
					series: [{
						name: '提交统计',
						type: 'pie',
						radius: ['40%', '70%'],
						avoidLabelOverlap: false,
						itemStyle: {
							borderRadius: 10,
							borderColor: '#fff',
							borderWidth: 2
						},
						label: {
							show: false,
							position: 'center'
						},
						emphasis: {
							label: {
								show: true,
								fontSize: '40',
								fontWeight: 'bold'
							}
						},
						labelLine: {
							show: false
						},
						data: []
					}],
				};
				this.option = option
				option && myChart.setOption(option);
			},
			handleClose(done) {
				done()
			},
			setStasticPanel(pid) {
				// 获取问题提交情况统计
				this.$problem_axios.GetSubmissionStaticForProblem(pid).then(res => {
					this.selfLog(res)
					this.statistic_info = res.data
					const keyMap = new Map([
						['4', '运行时错误'],
						['3', '内存超过'],
						['2', '超时'],
						['1', '编译错误'],
						['0', '通过'],
						['-1', '答案错误'],
						['-2', '编译错误'],
						['-3', '未完成判题'],
					])
					const colorMap = new Map([
						['4', '#5555ff'],
						['3', '#ff5500'],
						['2', '#ff5500'],
						['1', '#aa0000'],
						['0', '#00b8a9'],
						['-1', '#ff2e63'],
						['-2', '#e84545'],
						['-3', '#550000'],
					])
					for (let key in this.statistic_info) {
						let obj = {
							name: keyMap.get(this.statistic_info[key].result),
							value: this.statistic_info[key].nums
						}
						this.option.color.push(colorMap.get(this.statistic_info[key].result))
						this.option.series[0].data.push(obj)
					}
					var chartDom = document.getElementById('echart-pass-rate');
					var myChart = echarts.init(chartDom);
					this.option && myChart.setOption(this.option);
				})
			},
			async getData() {
				this.setStasticPanel(this.$route.params.id)
				// 用id获取问题的提交次数
				this.$problem_axios.GetAcSubTimes(this.$route.params.id).then(res => {
					this.selfLog(res)
					this.acceptSub = res.data.ac
					this.totalSub = res.data.submissions
				})
				// 用id获取问题详情
				await this.$problem_axios.ProblemDetailById(this.$route.params.id).then(res => {
					this.selfLog("获取问题详情")
					this.selfLog(res)
					if (res.errcode === 204) {
						this.$router.push('/404')
						this.$message({
							message: '该问题不存在~',
							type: 'warning'
						})
					}
					this.problbemDetail = res.data.problem
					this.problbemDetail['templates'] = res.data.templates
					this.InitLoadCode = true
					// 检查是否有前置的代码 有的话就不刷掉当前的代码
					let cookie_code = this.$cookies.get('judge_code' + this.problbemDetail.Pid)
					if (cookie_code === null) {
						this.selfLog('无前置代码')
						this.code = this.problbemDetail.TemplateC
					} else {
						this.selfLog('检测到前置代码')
						// this.selfLog(cookie_code)
						this.code = cookie_code
					}
					// 设置可用的语言选择下拉框
					let lans = []
					for (var key of Object.keys(res.data.templates)) {
						lans.push(key)
					}
					this.problbemDetail['languages'] = lans

					this.$problem_axios.GetProblemTagsById(this.problbemDetail.Pid).then(res => {
						this.selfLog("-=------=-----------")
						this.tags = res.data
					})
					this.setEchart()
					this.score = this.problbemDetail.total_score
					if (this.problbemDetail.submission_number === 0) {
						this.problbemDetail['pass_rate'] = 0
					} else {
						this.problbemDetail['pass_rate'] = parseFloat(
							(
								(this.problbemDetail.AcceptSubmissions /
									this.problbemDetail.TotalSubmissions) *
								100
							).toFixed(2)
						)
					}
					// this.samples = []
					// for (let i = 1; i <= this.problbemDetail.samples.length; i++) {
					// 	let sample = Object.assign({ id: i },
					// 		this.problbemDetail.samples[i - 1]
					// 	)
					// 	this.samples.push(sample)
					// }

				})
			},
			//切换语言换模板
			changeModel(chosen_lan) {
				this.language = chosen_lan
				this.selfLog(chosen_lan)
				// this.selfLog(this.problbemDetail.template)
				this.code = this.problbemDetail.templates[chosen_lan]
			},
			//编辑区域监听
			onCmCodeChange(newCode) {
				this.code = newCode
				this.$cookies.set('judge_code' + this.problbemDetail.Pid, newCode)
				this.selfLog('存储cookie')
				this.selfLog(newCode)
			},
			// 恢复默认模板代码
			recoverInit() {
				this.selfLog(this.problbemDetail)
				this.selfLog('当前语言：')
				this.selfLog(this.language)
				this.code = this.problbemDetail.templates[this.language]
			},
			// 提交代码
			submission() {
				this.err_info = ''
				this.submission_status = '传送...'
				this.drawer = true
				let params = {
					Code: this.code,
					ProblemId: this.problbemDetail.Pid,
					Language: this.language
				}
				this.selfLog(params)
				this.$problem_axios.Submission(params).then(res => {
					this.selfLog(res)
					if (res.errcode === 200) {
						
						this.drawer_loading = true
						// 用id获取问题的提交次数
						this.$problem_axios.GetAcSubTimes(this.$route.params.id).then(res => {
							this.selfLog(res)
							this.acceptSub = res.data.ac
							this.totalSub = res.data.submissions
						})
						this.submission_id = res.data.submissionID
						this.getSubmissionDetail(params.ProblemId, this.submission_id)
					} else if (res.errcode === 204) {
						this.$message({
							message: '提交失败，请先登录',
							type: 'warning'
						})
					} else {
						this.$message({
							message: '提交失败，请稍后重试',
							type: 'warning'
						})
					}
				})
			},
			// 提交代码之后定时查询提交的状态
			getSubmissionDetail(problem_id, submissionid) {
				// 先查有没有对应的提交
				this.$problem_axios.IsSubmissionExsit(submissionid).then(res => {
					// 提交存在
					this.selfLog('查询问题存在与否:')
					this.selfLog(res)
					if (res.errcode === 200) {
						this.selfLog('启动定时任务')
						// 需要设置定时一秒去轮询提交的submission状态 在页面展示之
						clearInterval(this.setInterval_id)
						this.setInterval_id = setInterval(() => {
							this.$problem_axios.GetSubmissionFinal(this.submission_id).then(res => {
								this.selfLog(res)
								if (res.errcode === 200) {
									//  更新提交次数
									this.score = res.data.Score
									this.LastTestcase = res.data.LastTestcase
									this.LastDesireOutput = res.data.LastDesireOutput
									this.LastOutput = res.data.LastOutput
									this.$problem_axios.GetAcSubTimes(problem_id).then(res => {
										this.selfLog(res)
										this.acceptSub = res.data.ac
										this.totalSub = res.data.submissions
									})
									// 更新状态完成就清楚定时任务
									let result_ = res.data.Result
									if (result_ === -3) {
										this.submission_status = '还没做'
									} else if (result_ === -2) {
										this.submission_status = '编译错误'
									} else if (result_ === -1) {
										this.submission_status = '答案错误'
									} else if (result_ === -0) {
										this.submission_status = '正确'
									} else if (result_ === 1) {
										this.submission_status = '计算超时'
									} else if (result_ === 2) {
										this.submission_status = '超时'
									} else if (result_ === 3) {
										this.submission_status = '内存超过'
									} else if (result_ === 4) {
										this.submission_status = '运行时错误'
									} else if (result_ === 5) {
										this.submission_status = '系统错误'
									} else if (result_ === 6) {
										this.submission_status = '传送...'
									} else if (result_ === 7) {
										this.submission_status = '判题中...'
									} else if (result_ === 8) {
										this.submission_status = '部分正确'
									}
									if (!(res.data.Result === 6 || res.data.Result === 7)) {
										this.err_info = res.data.ErrInfo
										this.drawer_loading = false
										clearInterval(this.setInterval_id)
									}
									this.selfLog('查询结束，清除定时任务')
									clearInterval(this.setInterval_id)
								} else if (res.errcode === 202 || res.errcode == 203) {
									this.submission_status = '判题中...'
									this.selfLog(this.submissionReply)
									this.submissionReply++
									// 超过二十次就重新提交一次
									if (this.submissionReply > 20) {
										clearInterval(this.setInterval_id)
										this.$message({
											message: '更新提交状态失败，请重新提交',
											type: 'warning'
										})
										this.submission_status = '服务器响应超时请重试提交...'
										this.drawer_loading = false
										this.submissionReply = 0
									}
								}
								// else if (res.errcode == 203) {
								// 被服务器限流
								else {
									clearInterval(that.setInterval_id)
								}
							})
						}, 3000)
					}
				})
			},
			// 获取提交的submission_id 的状态
			sendMsgRequest() {}
		},
		computed: {
			codemirror() {
				return this.$refs.cmEditor.codemirror
			},
			problbemDetailDifficulty() {
				let ans = ''
				switch (this.problbemDetail.Level) {
					case 1:
						ans = '简单'
						break
					case 2:
						ans = '中等'
						break
					case 3:
						ans = '困难'
						break
				}
				return ans
			}
		},
		watch: {
			//切换主题与语言
			theme: function(newTheme, oldTheme) {
				this.cmOptions.theme = newTheme
			},
			language: function(newLanguage, oldLanguage) {
				const modeMap = new Map([
					['C', 'text/x-csrc'],
					['C++', 'text/x-c++src'],
					['Golang', 'text/x-go'],
					['Java', 'text/x-java'],
					['Python2', 'text/x-python'],
					['Python3', 'text/x-python']
				])
				let cmMode = modeMap.get(newLanguage)
				this.cmOptions.mode = cmMode
			}
		},
		beforeDestroy() {
			clearInterval(this.setInterval_id)
		}
	}
</script>

<style lang='scss' scoped>
	#main {
		display: flex;
		width: 100%;
		justify-content: space-around;

		.q_desc {
			width: 30%;
			background-color: white;
			max-height: 100%;
			overflow-y: scroll;
			text-align: left;

			.header {
				margin: 20px 10px -20px 20px;

				q .title {
					color: black;
					font-weight: bold;
					font-size: 20px;
					margin: auto;
				}

				.low {
					color: #5ab726;
				}

				.mid {
					color: #ffa119;
				}

				.high {
					color: #ef4743;
				}
			}

			.body {
				margin-left: 20px;

				.sampleBox {
					background: #393e46;
					color: white;
					width: 60%;

					.input {
						line-height: 30px;
						margin-left: 10px;
					}

					.output {
						line-height: 30px;
						margin-left: 10px;
					}
				}

				#echart-pass-rate {
					height: 350px;
					width: 350px;
				}
			}
		}

		.ans_editor {
			width: 70%;
			overflow: auto;
			background-color: #393e46;

			.demo-inline {
				display: flex;
				margin-top: 10px;
				margin-bottom: 10px;

				.title {
					color: white;
					margin: 10px;
				}
			}

			.codeEditor {
				margin-left: 10px;
				text-align: left;
				height: 20vh;
			}

			.CodeMirror-scroll {
				height: auto;
				overflow-y: hidden;
				overflow-x: auto;
			}
		}

		.submit_fixed {
			position: absolute;
			z-index: 999;
			width: 200px;
			height: 50px;
			font-size: 20px;
			position: fixed;
			bottom: 21%;
			right: 50px;
		}
	}
</style>
