<template>
	<view class="brochuresMes">
		<view class="introduct">招生简章链接：</view>
		<view class="brochuresTitile" @click="goWeb()"> 北京大学 2023 年硕士研究生招生简章（校本部）</view>

<!-- 		<text class="mesText" space="ensp" >{{mes}}</text> -->
	</view>
</template>

<script>
	export default{
		data() {
			return {
				code: 0,
				mes: [],
				webUrl: 'https://admission.pku.edu.cn/docs/20220919154422268077.pdf?CSRFT=NM1Z-V0SB-3KTR-ZOJM-B4B0-ZPUE-Q2IB-GCN7'
			}
		},
		methods: {
			goWeb() {
				uni.navigateTo({
					url: '/pages/home/university/web?url=' + this.webUrl
				})
			},
		},
		created() {
			uni.$on('code',res=>{
				this.code = res
				uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {
				
				}).then(res => {
					this.mes = res.data.data;
					console.log("成功1")
					console.log(this.mes);
				}).catch(err => {
					console.log("失败")
				})
			})
		},
		mounted() {
		},
	}
</script>

<style>
	.brochuresMes{
		margin: 0rpx 30rpx;
	}
	.mesText{
		color: #434343;
		white-space: pre-wrap;
		word-break: break-word;
	}
	.brochuresTitile{
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 30rpx;
		text-decoration: underline
	}
	.introduct{
		font-size: 38rpx;
		font-weight: 600;
		margin-top: 20rpx;
		height: 80rpx;
	}
</style>