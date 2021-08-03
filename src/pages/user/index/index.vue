<template>
	<view class="">
		<view>
			<u-toast ref="uToast" />
		</view>	
		<view class="top_background">
			<u-count-down :timestamp="count_down_secs" separator="zh" :show-seconds="false" :show-border="true" font-size="40" color="#606266" border-color="#909399"></u-count-down>
		</view>
		<view class="grids">
			<swiper class="swiper" @change="change">
				<swiper-item>
					<u-grid :col="3" @click="click" hover-class="hover-class">
						<u-grid-item v-for="(item, index) in list" :index="index" :key="index">
							<u-icon :name="item" :size="46"></u-icon>
							<text class="grid-text">{{ '宫格' + (index + 1) }}</text>
						</u-grid-item>
					</u-grid>
				</swiper-item>
				<swiper-item>
					<u-grid :col="3" @click="click">
						<u-grid-item v-for="(item, index) in list" :index="index + 9" :key="index">
							<u-icon :name="item" :size="46"></u-icon>
							<text class="grid-text">{{ '宫格' + (index + 1) }}</text>
						</u-grid-item>
					</u-grid>
				</swiper-item>
				<swiper-item>
					<u-grid :col="3" @click="click">
						<u-grid-item v-for="(item, index) in list" :index="index + 18" :key="index">
							<u-icon :name="item" :size="46"></u-icon>
							<text class="grid-text">{{ '宫格' + (index + 1) }}</text>
						</u-grid-item>
					</u-grid>
				</swiper-item>
			</swiper>
			<view class="indicator-dots" v-if="isSwiper">
				<view class="indicator-dots-item" :class="[current == 0 ? 'indicator-dots-active' : '']">
				</view>
				<view class="indicator-dots-item" :class="[current == 1 ? 'indicator-dots-active' : '']">
				</view>
				<view class="indicator-dots-item" :class="[current == 2 ? 'indicator-dots-active' : '']">
				</view>
			</view>
		</view>
		<view class="cards" v-for="annoncement in annoncements">
			<u-card :title="annoncement.Title" :sub-title="annoncement.Time" :thumb="annoncement.Uid.Avatar">
				<view class="" slot="body">
					<view class="u-body-item u-flex u-border-bottom u-col-between u-p-t-0">
						<view class="u-body-item-title u-line-2">{{annoncement.Content}}</view>
						<image src="https://img11.360buyimg.com/n7/jfs/t1/94448/29/2734/524808/5dd4cc16E990dfb6b/59c256f85a8c3757.jpg" mode="aspectFill"></image>
					</view>
				</view>
				<view class="" slot="foot"><u-icon name="chat-fill" size="34" color="" label="30评论"></u-icon></view>
			</u-card>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isSwiper:true,
				current: 0,
				list: ['integral', 'kefu-ermai', 'coupon', 'gift', 'scan', 'pause-circle', 'gift', 'scan', 'pause-circle'],
				count_down_secs:0,
				title: '素胚',
				subTitle: '2020-05-15',
				thumb: 'https://avatar.csdnimg.cn/4/B/D/3_qq_45959912_1610702505.jpg',
				annoncements:[]
			}
		},
		mounted() {
			this.count_down_secs = this.get_secs_two_dates()*24*60*60
			// 获取公告信息
			this.$annonce_axios.GetAnnoncements({offset:0,limit:10}).then(res=>{
				this.annoncements = res.data
			})
		},
		onPullDownRefresh(){
			this.$annonce_axios.GetAnnoncements({offset:0,limit:10}).then(res=>{
				this.annoncements = res.data
				uni.stopPullDownRefresh();
				this.$refs.uToast.show({
					title: '刷新成功',
					type: 'primary ',
					position:'top',
					duration:800
				})
			})
		},
		methods: { 
			change(e) {
				this.current = e.detail.current
			},
			// 计算两天之间的距离
			get_secs_two_dates(){
				var time1 = Date.parse(new Date().toLocaleDateString())
				var time2 = Date.parse('2021/12/24')
				var nDays = Math.abs(parseInt((time2 - time1) / 1000 / 3600 / 24))
				return nDays
			}
		}
	}
</script>

<style scoped lang="scss">
	/* 下方这些scss变量为uView内置变量，详见开发  组件-指南-内置样式 */

	.grid-text {
		font-size: 28rpx;
		margin-top: 4rpx;
		color: $u-type-info;
	}
	
	.swiper {
		height: 480rpx;
	}
	
	.indicator-dots {
		margin-top: 40rpx;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	
	.indicator-dots-item {
		background-color: $u-tips-color;
		height: 6px;
		width: 6px;
		border-radius: 10px;
		margin: 0 3px;
	}
	
	.indicator-dots-active {
		background-color: $u-type-primary;
	}
	.top_background{
		background-color: #1296db;
		width: 100%;
		height: 80px;
		top: 0;
		// border: 1px black solid;
	}
	.grids{
		// position: absolute;
		top:40px;
		left: 0px;
		width: 100%;
		float: left;
	}
	.cards{
		height: 100%;
		width: 90%;
		align-items: center;
		margin-top: 20px;
		margin-left: 5%;
		float: left;
	}
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