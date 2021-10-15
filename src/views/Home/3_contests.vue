<template>
	<div id="main">
		<!-- 竞赛模糊查询 -->
		<el-form :inline="true" :model="formInline" class="demo-form-inline"
			style="margin:10px 0 0 20px;postion:flex;float:left">
			<el-form-item label="竞赛查询">
				<el-input v-model="formInline.keyword" placeholder="输入关键字" @change="search_task"></el-input>
			</el-form-item>
			<el-form-item label>
				<el-select v-model="formInline.status" @change="search_task">
					<el-option label="全部" value></el-option>
					<el-option label="进行中" value="0"></el-option>
					<el-option label="未开始" value="1"></el-option>
					<el-option label="已结束" value="-1"></el-option>
				</el-select>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="search_task">查询</el-button>
			</el-form-item>
		</el-form>
		<!-- 竞赛列表显示 -->
		<el-table :data="tableData" style="width: 100%" max-height="600">
			<el-table-column type="expand">
				<template slot-scope="props">
					<el-form label-position="left" inline class="demo-table-expand">
						<el-form-item label="发布时间">
							<span>{{ props.row.create_time | formatDate}}</span>
						</el-form-item>
						<el-form-item label="开始时间">
							<span>{{ props.row.start_time | formatDate}}</span>
						</el-form-item>
						<el-form-item label="最近修改时间">
							<span>{{ props.row.last_update_time | formatDate}}</span>
						</el-form-item>
						<el-form-item label="截止时间">
							<span>{{ props.row.end_time | formatDate}}</span>
						</el-form-item>
					</el-form>
				</template>
			</el-table-column>
			<el-table-column label="竞赛" prop="ContestName"></el-table-column>
			<el-table-column label="竞赛描述" prop="ContestDesc">
				<template slot-scope="scope">
					<div v-html="scope.row.ContestDesc"></div>
				</template>
			</el-table-column>
			<el-table-column label="创建人" prop="Creater.UserName"></el-table-column>
			<el-table-column label="截止时间" sortable prop="end_time">
				<template slot-scope="scope">
					<i class="el-icon-time"></i>
					<span style="margin-left: 10px">{{ scope.row.EndingTime | formatDate}}</span>
				</template>
			</el-table-column>
			<el-table-column label="竞赛模式" prop="ContestType">
				<template slot-scope="scope">
					<div v-html="scope.row.ContestType"></div>
				</template>
			</el-table-column>
			<el-table-column label="状态">
				<template slot-scope="scope">
					<el-tag v-if="scope.row.status==='1'">未开始</el-tag>
					<el-tag v-if="scope.row.status==='0'" type="success">进行中</el-tag>
					<el-tag v-if="scope.row.status==='-1'" type="danger">已结束</el-tag>
				</template>
			</el-table-column>
			<!-- 只有进行中的竞赛才可参加 -->
			<el-table-column label="操作">
				<template slot-scope="scope">
					<el-tooltip v-if="scope.row.status==='1'" content="竞赛还没开始喔～" placement="top">
						<el-button type="primary" round>参加</el-button>
					</el-tooltip>
					<el-tooltip v-if="scope.row.status==='-1'" content="来晚啦,竞赛已截止～" placement="top">
						<el-button type="primary" round>参加</el-button>
					</el-tooltip>
					<el-tooltip v-if="scope.row.status==='0'" content="点击参加此次竞赛～" placement="top">
						<el-button type="primary" @click="EnterContest(scope.row.Cid)" round>参加</el-button>
					</el-tooltip>
				</template>
			</el-table-column>
		</el-table>
		<!-- 输入密码弹框控件 -->
		<el-dialog title="验证" :visible.sync="dialogFormVisible" append-to-body width="30%" :before-close="handleClose">
			<el-form :model="form">
				<el-form-item label="竞赛密码" label-width="80px">
					<el-input type="password" v-model="form.password" placeholder="请输入竞赛密码" show-password
						prefix-icon="el-icon-lock" clearable></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="cancelYanZhen">取 消</el-button>
				<el-button type="primary" @click="EnterContest_password">确 定</el-button>
			</div>
		</el-dialog>
		<!-- 竞赛问题列表抽屉 -->
		<el-drawer title :visible.sync="table" direction="rtl" size="50%" append-to-body>
			<el-button type="primary" round @click="show_drawer_for_annoce"
				style="postion:absolute;left:20px;top:-20px">查看公告</el-button>
			<el-table :data="problems" style="width: 100%">
				<el-table-column type="expand">
					<template slot-scope="props">
						<el-form label-position="left" inline class="demo-table-expand">
							<el-form-item label="习题描述:">
								<div v-html="props.row.Content"></div>
							</el-form-item>
							<el-form-item label="可使用语言:">
								<el-tag style="margin:10px 0 10px 10px" :type="tag_types[index%5]"
									v-for="(value, index) of props.row.supported_lan" :key="value">{{value}}</el-tag>
							</el-form-item>
						</el-form>
					</template>
				</el-table-column>
				<el-table-column label="ID" prop="Pid"></el-table-column>
				<el-table-column label="标题" prop="Title"></el-table-column>
				<el-table-column label="已完成人数" prop="complete_num"></el-table-column>
				<el-table-column label="操作" width="300" header-align="left">
					<template slot-scope="scope">
						<el-tooltip content="点击去完成习题喔～" placement="top">
							<el-button size="mini" type="success" round @click="doProblem(scope.row.Pid)">
								做它
								<i class="el-icon-thumb"></i>
							</el-button>
						</el-tooltip>
					</template>
				</el-table-column>
			</el-table>
			<!--公告列表抽屉 -->
			<el-drawer title="公告列表" :visible.sync="contest_annon_list_drawer_visble" direction="rtl" size="45%"
				:append-to-body="true">
				<el-table :data="contest_annon_list">
					<el-table-column property="title" label="标题" width="150"></el-table-column>
					<el-table-column property="content" label="内容" width="200"></el-table-column>
					<el-table-column label="创建时间" width="300">
						<template slot-scope="scope">
							<i class="el-icon-time"></i>
							<span style="margin-left: 10px">{{ scope.row.create_time | formatDate}}</span>
						</template>
					</el-table-column>
					<el-table-column property="created_by.username" label="创建人" width="130"></el-table-column>
				</el-table>
			</el-drawer>
		</el-drawer>

		<!-- 分页控件 -->
		<div class="block">
			<el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
				:page-sizes="[10, 20, 40, 80]" :page-size="100" layout="total, sizes, prev, pager, next, jumper"
				:total="total"></el-pagination>
		</div>
	</div>
</template>

<script>
	export default {
		beforeMount() {
			this.GetAllContests()
			this.$user_axios.user_profile().then(res => {
				this.selfLog(res)
				this.user_profile = res['data']
				this.selfLog(this.user_profile)
			})
		},
		data() {
			return {
				tag_types: ['info', 'warning', 'danger', '', 'success'],
				user_profile: null,
				// 分页查看
				page: 1,
				offset: 0,
				limit: 10,
				total: 0,
				contest_id: 0,
				tableData: [],
				dialogFormVisible: false,
				form: { password: '' },
				formInline: {
					keyword: '',
					status: ''
				},
				problems: {},
				complete_num:[],
				table: false,
				contest_annon_list_drawer_visble: false,
				contest_annon_list: []
			}
		},
		methods: {
			handleClose(done) {
				this.form.password = ''
				done()
			},
			cancelYanZhen() {
				this.form.password = ''
				this.dialogFormVisible = false
			},
			GetAllContests() {
				this.$contests_axios.GetAllContests(this.offset, this.limit).then(res => {
					this.tableData = res.data
					for (let contest of this.tableData) {
						// 转化类型码
						if (contest['ContestType'] === 0){
							contest['ContestType'] = 'ACM'
						}else if (contest['ContestType'] === 1){
							contest['ContestType'] = 'OI'
						}else if(contest['ContestType'] === 2){
							contest['ContestType'] = 'IOI'
						}
						var curDate = new Date(),
							beginDate = new Date(contest.StartTime),
							endDate = new Date(contest.EndingTime)
							// 正在进行 0  未开始1   已结束-1
						if (curDate >= beginDate && curDate <= endDate) {
							contest['status'] = '0'
						}else if(curDate>endDate){
							contest['status'] = '-1'
						}else{
							contest['status'] = '1'
						}
						this.selfLog(this.tableData)
					}
				})
				this.$contests_axios.GetContestQuantity().then(res => {
					this.total = res.data
				})
			},
			// // 查询公告列表
			get_annonce_for_contest() {
				this.$axios.get_annonce_for_contest_stu(this.contest_id).then(res => {
					this.selfLog(res)
					this.contest_annon_list = res.data
					this.contest_annon_list_drawer_visble = true
				})
			},
			// 展示公告抽屉
			async show_drawer_for_annoce() {
				await this.get_annonce_for_contest(this.contest_id)
				this.contest_annon_list_drawer_visble = true
			},
			// 监听分页控件
			handleSizeChange(val) {
				this.limit = val
				this.offset = (this.page - 1) * this.limit
				this.Get_Contests_List()
				this.selfLog(`每页 ${val} 条`)
			},
			handleCurrentChange(val) {
				this.page = val
				this.offset = (this.page - 1) * this.limit
				this.Get_Contests_List()
				this.selfLog(`当前页: ${val}`)
			},
			// 参加按钮监听
			EnterContest(contest_id) {
				if (this.user_profile === null) {
					this.$message({
						message: '查看习题请先登录',
						type: 'warning'
					})
					return false
				}
				this.selfLog(contest_id)
				this.contest_id = contest_id
				this.$contests_axios.ContestAccess(contest_id).then(res => {
					this.selfLog(res)
					// 没有权限就要输密码
					if (res.errcode===205) {
						this.$message({
							message: '暂无权限，请输入密码',
							type: 'warning'
						})
						this.dialogFormVisible = true
					} else if(res.errcode===204){
						this.$message({
							message: '查看习题请先登录',
							type: 'warning'
						})
					}else if(res.errcode===200){
						this.$message({
							message: '拥有权限，可进入',
							type: 'success'
						})
						// 查询竞赛的问题列表
						this.$contests_axios.ContestProblemList(this.contest_id).then(res => {
							this.selfLog(res)
							if (res.data === 'Contest has not started yet.') {
								this.$message({
									message: '竞赛尚未到达开始时间',
									type: 'warning'
								})
							} else {
								this.problems = res.data.problems
								this.complete_num = res.data.complete_num
								this.selfLog( this.complete_num)
								for(let problem of this.problems){
									this.$problem_axios.GetProblemSupportLanguage(problem.Pid).then(res=>{
										problem['supported_lan'] = res.data
										problem['complete_num'] = this.complete_num[problem['Pid']]
									})
								}
								this.dialogFormVisible = false
								this.table = true
							}
						})
					}
				})
				// this.dialogFormVisible = false
			},
			EnterContest_password(id) {
				if (this.user_profile === null) {
					this.$message({
						message: '查看习题请先登录',
						type: 'warning'
					})
					return false
				}
				if (this.form.password === '') {
					this.$message({
						message: '密码不能为空',
						type: 'warning'
					})
					return false
				}
				this.$contests_axios.ContestPasswordCheck(this.contest_id,this.form.password).then(res => {
					this.selfLog(res)
					if (res.errcode === 204) {
						this.$message({
							message: '密码错误或者密码已过期',
							type: 'error'
						})
					} else if (res.errcode === 200) {
						this.$message({
							message: '密码正确~',
							type: 'success'
						})
						// 查询竞赛的问题列表
						this.$contests_axios.ContestProblemList(this.contest_id).then(res => {
							this.selfLog(res)
							if (res.data === 'Contest has not started yet.') {
								this.$message({
									message: '竞赛尚未到达开始时间',
									type: 'warning'
								})
							} else {
								this.problems = res.data.problems
								this.complete_num = res.data.complete_num
								this.selfLog( this.complete_num)
								for(let problem of this.problems){
									problem['complete_num'] = this.complete_num[problem['Pid']]
									this.$problem_axios.GetProblemSupportLanguage(problem.Pid).then(res=>{
										problem['supported_lan'] = res.data
									})
								}
								this.dialogFormVisible = false
								this.table = true
							}
						})
					}
				})
			},
			doProblem(id) {
				this.selfLog('前端传入路径参数问题id：'+id)
				this.$router.push(`/problemDetail_Contest/${id}/${this.contest_id}`)
			},
			search_task() {
				this.$axios
					.search_assign(this.offset, this.limit, this.formInline)
					.then(res => {
						this.total = res.data.total
						this.tableData = res.data.results
						this.selfLog(res)
					})
			}
		},
		computed: {}
	}
</script>

<style lang='scss' scoped>
	.demo-table-expand {
		font-size: 0;
	}

	.demo-table-expand label {
		width: 90px;
		color: #99a9bf;
	}

	.demo-table-expand .el-form-item {
		margin-right: 0;
		margin-bottom: 0;
		width: 50%;
	}
</style>
