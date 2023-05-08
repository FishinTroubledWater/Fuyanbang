<template>
	<view>
		<view class="academyMes">
			<image class="academyLogo" src="@/static/academy-icons/北大.png"></image>
			<view class="viewText">
				<text class="academyName">{{mes.Name}}</text>
				<view class="academyType">
					<view class="typeOfScoreLine">{{mes.LineType}}</view>
					<text class="typeOfInstitution">{{mes.Level}}</text>
				</view>
				<text class="academyCode">{{mes.Region}}|院校代码：{{mes.Code}}</text>
			</view>
		</view>
		<view class="academyBar">
			<view class="barItem" @click="show(1)" :class="index===1? 'active':''">院校简介</view>
			<view class="barItem" @click="show(2)" :class="index===2? 'active':''">招生简章</view>
			<view class="barItem" @click="show(3)" :class="index===3? 'active':''">考研分数线</view>
			<view class="barItem" @click="show(4)" :class="index===4? 'active':''">学长学姐说</view>
		</view>
		<view>
			<!-- <keep-alive>
				<component :is="comp" v-show="isShow"></component>
			</keep-alive> -->
			<one v-show="isShow1" :title='mes'></one>
			<Two v-show="isShow2" :title='mes'></Two>
			<Three v-show="isShow3" :title='mes'></Three>
			<four v-show="isShow4" :title='mes'></four>
		</view>
	</view>
</template>

<script>
	import one from '@/pages/home/university/introduction.vue'
	import Two from '@/pages/home/university/brochures.vue'
	import Three from '@/pages/home/university/line.vue'
	import four from '@/pages/home/university/said.vue'
	export default {
		components: {
		    one:one,
		    Two:Two,
		    Three:Three,
		    four:four,
		  },
		data() {
			return {
				// 控制切换按钮后按钮改变样式
				index: 1,
				// 控制子组件显示
				comp: 'one',
				// 控制点击按钮后子组件显示，再次点击隐藏
				isShow: true,
				isShow1: true,
				isShow2: false,
				isShow3: false,
				isShow4: false,
				active: '',
				code: 0,
				mes: [],
				guide: '',
				name: '',
			}
		},
		onLoad:function(option){
			var _this = this
			this.code = option.code;
			uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {
			
			}).then(res => {
			    console.log(res.data.data);
				this.mes = res.data.data;
			}).catch(err => {
			
			})
			// uni.$on('code1',function(data) {
			// 	_this.guide = data.guide;
			// 	_this.name = data.name;
			// 	console.log("内部的guide是：");
			// 	console.log(_this.guide)
			// })
		},
		methods: {
		    show (value) {
				if (this.index === value) {
					this.index = 0
				} else {
					this.index = value
					this.isShow = true
				}
				// if (value === 1) this.comp = 'one'
				// if (value === 2) this.comp = 'two'
				// if (value === 3) this.comp = 'three'
				// if (value === 4) this.comp = 'four'
				if (value === 1){
					this.isShow1 = true
					this.isShow2 = false
					this.isShow3 = false
					this.isShow4 = false
				}
				if (value === 2){
					this.isShow1 = false
					this.isShow2 = true
					this.isShow3 = false
					this.isShow4 = false
				}
				if (value === 3){
					this.isShow1 = false
					this.isShow2 = false
					this.isShow3 = true
					this.isShow4 = false
				}
				if (value === 4){
					this.isShow1 = false
					this.isShow2 = false
					this.isShow3 = false
					this.isShow4 = true
				}
				this.$route.meta.index = 0
		    }
			
		}
	}
</script>

<style>
	.academyMes{
		display: flex;
		flex-direction: row;
		justify-content: center;
		width: 100%;
		height: 180rpx;
		background-color: #F8F8F8;
	}
	.viewText {
		display: flex;
		flex-direction: column;
	}
	.academyLogo{
		margin-right: 40rpx;
		height: 150rpx;
		width: 150rpx;
	}
	.academyName{
		font-size: 40rpx;
		font-family: "黑体";
	}
	.academyType{
		display: flex;
		margin-top: 5rpx;
	}
	.typeOfScoreLine{
		width: 150rpx;
		height: 50rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 18rpx;
		background-color: #96C5F1;
		font-size: 26rpx;
		color: #ffffff;
		font-family: "黑体";
	}
	.typeOfInstitution{
		width: 100rpx;
		height: 50rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-left: 20rpx;
		border-radius: 18rpx;
		background-color: #96C5F1;
		font-size: 26rpx;
		color: #ffffff;
		font-family: "黑体";
	}
	.academyCode{
		margin-top: 5rpx;
		font-size: 30rpx;
	}
	.academyBar{
		display: flex;
		justify-content: center;
	}
	.barItem:nth-child(-n+2){
		width: 180rpx;
		height: 80rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 30rpx;
	}
	.barItem:nth-child(n+3){
		width: 200rpx;
		height: 80rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 30rpx;
	}
	.active {
		color: #ffffff;
		background: #8fc9ff;
		border-bottom: transparent;
/* 		font-weight: 700;
		border-bottom:4rpx solid #BD3124; */
	}
</style>
