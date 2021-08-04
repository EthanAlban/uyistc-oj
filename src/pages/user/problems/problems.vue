<template>
	<view>
		<view class="cards" v-for="problem in problems">
			<u-card :title="problem.Title" sub-title="" :thumb="problem.Uid.Avatar">
				<view class="" slot="body">
					<view class="u-body-item u-flex u-border-bottom u-col-between u-p-t-0">
						<view class="u-body-item-title u-line-2">{{problem.Content}}</view>
						<image
							src="https://img11.360buyimg.com/n7/jfs/t1/94448/29/2734/524808/5dd4cc16E990dfb6b/59c256f85a8c3757.jpg"
							mode="aspectFill" style="left:10px;"></image>
					</view>
				</view>
				<view class="" slot="foot">
					<u-icon name="chat-fill" size="34" color="" :label="problem.TotalSubmissions+'次提交'"></u-icon>
				</view>
			</u-card>
		</view>
	</view>

</template>

<script>
	export default {
		data() {
			return {
				problems:[]
			}
		},
		mounted() {
			this.$problem_axios.GetPagesProblems({
				offset: 0,
				limit: 10
			}).then(res => {
				this.problems = res.data
				this.selfLog(res)
			})
		}
	}
</script>

<style scoped lang="scss">
	.u-card-wrap {
		background-color: $u-bg-color;
		padding: 1px;
	}

	.u-body-item {
		font-size: 32rpx;
		color: #333;
		padding: 20rpx 10rpx;
	}

	.u-body-item image {
		width: 120rpx;
		flex: 0 0 120rpx;
		height: 120rpx;
		border-radius: 8rpx;
		margin-left: 12rpx;
	}
</style>
