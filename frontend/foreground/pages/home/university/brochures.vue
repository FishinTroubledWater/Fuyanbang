<template>
	<view class="brochuresMes">
<!-- 		<button @click="test">aaaaa</button> -->
		<view class="introduct">招生简章链接：</view>
		<view class="brochuresTitile" @click="goWeb()"> {{this.name}} 2023 年硕士研究生招生简章（校本部）</view>
<!-- 		<view class="brochuresTitile"> {{this.guide}}</view> -->

<!-- 		<text class="mesText" space="ensp" >{{mes}}</text> -->
	</view>
</template>

<script>
	export default{
		data() {
			return {
				code: 0,
				guide: '',
				name: '',
				mes: [],
				webUrl: 'https://admission.pku.edu.cn/docs/20220919154422268077.pdf?CSRFT=NM1Z-V0SB-3KTR-ZOJM-B4B0-ZPUE-Q2IB-GCN7'
			}
		},
		props:['title'],
		mounted() {
			// console.log("上方");
			// var _this= this;
			// uni.$on('code1',function(data) {
			// 	_this.guide = data.guide;
			// 	_this.name = data.name;
			// 	console.log("内部的guide是：");
			// 	console.log(_this.guide)
			// })
			// console.log("下方");
		},
		methods: {
			test(){
				console.log("HHHHHH");
				console.log(this.guide)
			},
			goWeb() {
				uni.navigateTo({
					url: '/pages/home/university/web?url=' + this.guide
				})
			},
		},
		created() {
			console.log("上方");
			var _this= this;
			uni.$on('code1',function(data) {
				_this.guide = data.guide;
				_this.name = data.name;
				console.log("内部的guide是：");
				console.log(_this.guide)
			})
			console.log("下方");
			// uni.$on('code',res=>{
			// 	this.code = res
			// 	uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {
				
			// 	}).then(res => {
			// 		this.mes = res.data.data;
			// 		console.log("成功1")
			// 		console.log(this.mes);
			// 	}).catch(err => {
			// 		console.log("失败")
			// 	})
			// })
			// var _this= this;
			// const on = uni.$on('code1',function(data) {
			// 	_this.guide = data.guide;
			// 	_this.name = data.name;
			// 	console.log("内部的guide是：");
			// 	console.log(_this.guide)
			// })
			// console.log("外部的guide是：");
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