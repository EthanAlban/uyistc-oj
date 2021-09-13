<template>
	<div>
		<el-button type="primary" style="float:left;margin:10px 0 10px 10px" round @click="GetSubmissions">刷新提交记录
		</el-button>
		<!-- 将页面高度存进vuex，这边动态绑定最大高度 -->
			<el-table :data="submissons_list" style="width: 100%" max-height="600">
				<el-table-column type="expand">
					<template slot-scope="props">
						<el-form label-position="left" inline class="demo-table-expand">
							<el-form-item label="提交时间">
								<span>{{ props.row.create_time | formatDate}}</span>
							</el-form-item>
							<el-form-item label>
								<el-button type="primary" @click="getCode(props.row.Code)" plain>查看代码</el-button>
							</el-form-item>
							<el-form-item label="错误提示">
								<span style="color: red">
									{{props.row.ErrInfo}}
								</span>
							</el-form-item>
						</el-form>
					</template>
				</el-table-column>

				<el-table-column label="问题编号" prop="problem">
					<template slot-scope="scope">
						<el-button @click="problemDetail(scope.row.ProblemId.Pid)" type="text" size="small">
							{{scope.row.ProblemId.Pid}}</el-button>
					</template>
				</el-table-column>
				<el-table-column label="提交时间" sortable prop="create_time">
					<template slot-scope="scope">
						<i class="el-icon-time"></i>
						<span style="margin-left: 10px">{{ scope.row.CreateTime | formatDate}}</span>
					</template>
				</el-table-column>
				<el-table-column label="使用语言" prop="Language.Language_"></el-table-column>
				<el-table-column label="结果" prop="result">
					<template slot-scope="scope">
						<div v-if="scope.row.result!='正确'">
							<span style="color:red">{{scope.row.result}}</span>
						</div>
						<div v-if="scope.row.result==='正确'" title="正确">
							<span style="color:green">{{scope.row.result}}</span>
						</div>
					</template>
				</el-table-column>
				<el-table-column label="得分" prop="Score"></el-table-column>
				<el-table-column label="用户名" prop="UserName"></el-table-column>
			</el-table>
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
		mounted() {
			this.GetSubmissions(this.page, this.limit, this.offset);
			if (this.user_profile === null) {
				const h = this.$createElement;
				this.$notify({
					title: "未登录",
					message: h("i", {
						style: "color: teal"
					}, "未登录无法查看提交记录")
				});
			}
		},
		data() {
			return {
				submissons_list: [],
				limit: 10,
				page: 1,
				total: 0,
				offset: 0
			};
		},

		methods: {
			tableRowClassName({
				row,
				rowIndex
			}) {
				if (rowIndex === 1) {
					return "warning-row";
				} else if (rowIndex === 3) {
					return "success-row";
				}
				return "";
			},
			// 获取问题列表
			GetSubmissions(page, limit, offset) {
				this.selfLog("更新提交列表");
				this.offset = (this.page - 1) * this.limit;
				this.$submission_axios.GetUserSubmissions(this.limit, this.offset).then(res => {
					if (res.errcode === 204) {
						const h = this.$createElement;
						this.$notify({
							title: "未登录",
							message: h("i", {
								style: "color: teal"
							}, "未登录无法查看提交记录")
						});
						return
					}
					this.selfLog(res)
					// 将编译结果和编码对应
					this.submissons_list = res.data.submissions;
					for (let i = 0; i < this.submissons_list.length; i++) {
						let result_ = this.submissons_list[i].Result;
						if (result_ === -2) {
							this.submissons_list[i].result = "编译错误";
						} else if (result_ === -1) {
							this.submissons_list[i].result = "答案错误";
						} else if (result_ === 0) {
							this.submissons_list[i].result = "正确";
						} else if (result_ === 1) {
							this.submissons_list[i].result = "计算超时";
						} else if (result_ === 2) {
							this.submissons_list[i].result = "超时";
						} else if (result_ === 3) {
							this.submissons_list[i].result = "内存超过";
						} else if (result_ === 4) {
							this.submissons_list[i].result = "运行时错误";
						} else if (result_ === 5) {
							this.submissons_list[i].result = "系统错误";
						} else if (result_ === 6) {
							this.submissons_list[i].result = "传送...";
						} else if (result_ === 7) {
							this.submissons_list[i].result = "判题中...";
						} else if (result_ === 8) {
							this.submissons_list[i].result = "部分正确";
						}
					}
					this.selfLog(this.submissons_list)
					this.total = res.data.total;
				});
			},
			// 监听分页控件
			handleSizeChange(val) {
				this.limit = val;
				this.offset = (this.page - 1) * this.limit;
				this.GetSubmissions(this.page, this.limit, this.offset);
				this.selfLog(`每页 ${val} 条`);
			},
			handleCurrentChange(val) {
				this.page = val;
				this.offset = (this.page - 1) * this.limit;
				this.GetSubmissions(this.page, this.limit, this.offset);
				this.selfLog(`当前页: ${val}`);
			},
			// 进入问题详情页
			problemDetail(id) {
				this.$router.push(`/problemDetail/${id}`);
			},
			// 提交列表查看提交的代码
			getCode(Code) {
				this.$confirm(Code, "提交代码", {
						distinguishCancelAndClose: true,
						confirmButtonText: "保存到剪贴板",
						cancelButtonText: "取消"
					})
					.then(() => {
						this.$copyText(Code);
						this.$message({
							type: "info",
							message: "已复制"
						});
					})
					.catch(action => {
						this.selfLog(action)
						this.$message({
							type: "info",
							message: action === "cancel" ? "放弃保存并离开页面" : "未保存"
						});
					});
			}
		}
	};
</script>

<style lang='scss' scoped>
	#main {
		background-color: #3d84a8;
		width: 100%;
		line-height: 600px;
	}

	.el-table .warning-row {
		background: oldlace;
	}

	.el-table .success-row {
		background: #f0f9eb;
	}
</style>
