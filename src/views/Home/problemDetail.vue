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
					<el-tag style="margin:10px 0 10px 10px" :type="tag_types[idx%5]"
						v-for="(tag,idx) in problbemDetail.tags" :key="tag">{{tag}}</el-tag>
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
					<span>通过次数:{{ problbemDetail.AcceptSubmissions }}</span>
					<el-divider direction="vertical"></el-divider>
					<span>提交次数:{{ problbemDetail.TotalSubmissions }}</span>
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
			</div>

			<!-- <codemirror class="codeEditor" ref="cmEditor" :value="code" :options="cmOptions" @ready="onCmReady"
				@focus="onCmFocus" @input="onCmCodeChange" /> -->
			<codemirror class="codeEditor" ref="cmEditor" :value="code" :options="cmOptions" @input="onCmCodeChange" />

		</div>
		<!-- 提交固定于页面右下角 -->
		<!-- 提交结果以抽屉展示 从下往上-->
		<el-button class="submit_fixed" type="success" round @click="submission">提交</el-button>

		<!-- -2编译错误 -1答案错误 0正确 1计算超时  2超时 3内存超过 4运行时错误 5传送...  6判题中...  7部分正确 -->
		<el-drawer title="提交结果" :visible.sync="drawer" append-to-body destroy-on-close direction="btt"
			:before-close="handleClose" style="width:60%;marginLeft:39.8%">
			<el-tag v-if="submission_status==='编译错误'" type="danger" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='答案错误'" type="danger" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='正确'" type="success" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='计算超时'" type="info" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='超时'" type="info" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='内存超过'" type="info" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='运行时错误'" type="success" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='传送...'" type="info" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='判题中...'" type="info" effect="dark">{{submission_status}}</el-tag>
			<el-tag v-if="submission_status==='部分正确'" type="warning" effect="dark">{{submission_status}}</el-tag>
			<br />
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
		},
		data() {
			return {
				tag_types: ['info', 'warning', 'danger', '', 'success'],
				drawer: false,
				problbemDetail: {},
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
					smartIndent:true
					// more CodeMirror options...
				},
				submission_id: '',
				setInterval_id: '',
				submission_status: '',
				err_info: '',
				score: '',
				option: {
					tooltip: { trigger: 'item' },
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
				//配置做题分析数据
				this.option.series[0].data = []
				const keyMap = new Map([
					['0', '通过'],
					['-1', '超时'],
					['-2', '编译错误']
				])
				const colorMap = new Map([
					['0', '#00b8a9'],
					['-1', '#ff2e63'],
					['-2', '#e84545']
				])
				for (let key in this.problbemDetail.statistic_info) {
					let obj = {
						name: keyMap.get(key),
						value: this.problbemDetail.statistic_info[key],
						itemStyle: { color: colorMap.get(key) }
					}
					this.option.series[0].data.push(obj)
				}
				let chartDom = document.getElementById('echart-pass-rate')
				let myChart = echarts.init(chartDom)
				myChart.setOption(this.option, true)
			},
			handleClose(done) {
				done()
			},
			async getData() {
				// 用id获取问题详情
				await this.$problem_axios.ProblemDetailById(this.$route.params.id).then(res => {
					this.selfLog(res)
					if (res.errcode === 204) {
						this.$router.push('/404')
						this.$message({
							message: '该问题不存在~',
							type: 'warning'
						})
					}
					this.problbemDetail = res.data
					this.code = res.data.Template
					this.$problem_axios.GetProblemTagsById(this.problbemDetail.Pid).then(res => {
						this.problbemDetail['tags'] = res.data
						this.selfLog(this.problbemDetail)
					})
					this.setEchart()
					// 检查是否有前置的代码 有的话就不刷掉当前的代码
					if (this.code === '') {
						this.selfLog('无前置代码')
						this.code = this.problbemDetail.template[this.language]
					} else {
						this.selfLog('检测到前置代码')
					}
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
					this.samples = []
					for (let i = 1; i <= this.problbemDetail.samples.length; i++) {
						let sample = Object.assign({ id: i },
							this.problbemDetail.samples[i - 1]
						)
						this.samples.push(sample)
					}
					if (this.code === '') {
						this.code = this.problbemDetail.template.C
					}

				})
			},
			//切换语言换模板
			changeModel(chosen_lan) {
				this.selfLog(chosen_lan)
				this.selfLog(this.problbemDetail.template)
				this.code = this.problbemDetail.template[chosen_lan]
			},
			// onCmReady(cm) {
			// 	this.selfLog('the editor is readied!', cm)
			// },
			// onCmFocus(cm) {
			// 	this.selfLog('the editor is focused!', cm)
			// },
			onCmCodeChange(newCode) {
				this.selfLog(newCode)
				this.code = newCode
			},
			// 提交代码
			submission() {
				this.drawer = true
				let params = {
					Code: this.code,
					ProblemId: this.problbemDetail.Pid,
					Language: this.language
				}
				this.selfLog(params)
				this.$problem_axios.Submission(params).then(res => {
					this.selfLog(res)
					if (res.error === null) {
						//提交之后刷新页面
						this.getData()
						// 把用户代码再写进去避免每次提交都被刷新掉
						//   this.problbemDetail.template[this.language] = this.code;
						this.submission_id = res.data.submission_id
						this.getSubmissionDetail(params.problem_id, this.submission_id)
					} else if (res.error === 'permission-denied') {
						this.$message({
							message: '提交失败，请先登录',
							type: 'warning'
						})
					} else if (res.data === 'code: This field is required.') {
						this.$message({
							message: '提交失败，代码区域为空',
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
			getSubmissionDetail(problem_id) {
				const that = this
				// 先查有没有对应的提交
				this.$axios.submission_exists(problem_id).then(res => {
					// 提交存在
					this.selfLog(res)
					if (res.data) {
						this.selfLog('启动定时任务')
						// 需要设置定时一秒去轮询提交的submission状态 在页面展示之
						that.setInterval_id = setInterval(() => {
							this.$axios.submission_id(that.submission_id).then(res => {
								this.selfLog(res)
								if (res.error === null) {
									// 更新状态完成就清楚定时任务
									let result_ = res.data.result
									if (result_ === -2) {
										that.submission_status = '编译错误'
									} else if (result_ === -1) {
										that.submission_status = '答案错误'
									} else if (result_ === -0) {
										that.submission_status = '正确'
									} else if (result_ === 1) {
										that.submission_status = '计算超时'
									} else if (result_ === 2) {
										that.submission_status = '超时'
									} else if (result_ === 3) {
										that.submission_status = '内存超过'
									} else if (result_ === 4) {
										that.submission_status = '运行时错误'
									} else if (result_ === 5) {
										that.submission_status = '系统错误'
									} else if (result_ === 6) {
										that.submission_status = '传送...'
									} else if (result_ === 7) {
										that.submission_status = '判题中...'
									} else if (result_ === 8) {
										that.submission_status = '部分正确'
									}
									if (!(res.data.result === 6 || res.data.result === 7)) {
										that.err_info = res.data.statistic_info.err_info
										that.score = res.data.statistic_info.score
										clearInterval(that.setInterval_id)
									}
								} else {
									this.$message({
										message: '更新提交状态失败',
										type: 'warning'
									})
								}
							})
						}, 1000)
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
			max-height: 60%;
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
			bottom: 150px;
			right: 50px;
		}
	}
</style>
