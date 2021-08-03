<template>
	<view>
		<u-avatar :src="form.avatar"></u-avatar>
		<u-avatar :text="text"></u-avatar>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				src: 'http://pic2.sc.chinaz.com/Files/pic/pic9/202002/hpic2119_s.jpg',
				text: '无头像',
				form:{
					avatar:""
				}
			}
		},
		mounted() {
			this.$user_axios.GetUserByUsernamePwd().then(res=>{
				this.selfLog(res)
				this.form.avatar = res.data.Avatar
			})
		}
	}
</script>

<style>
</style>
