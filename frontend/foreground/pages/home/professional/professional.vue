<template>
	<view>
		<view class="professionalMes">
			<view class="viewText">
				<text class="professionalName">(010100)哲学</text>
				<view class="professionalType">
					<text class="type">门类：哲学</text>
					<text class="firstLevelDiscipline">一级学科：哲学</text>
				</view>
			</view>
		</view>
		<view class="academyBar">
			<view class="barItem" @click="show(1)" :class="index===1? 'active':''">专业介绍</view>
			<view class="barItem" @click="show(2)" :class="index===2? 'active':''">分数线</view>
			<view class="barItem" @click="show(3)" :class="index===3? 'active':''">就业前景</view>
		</view>
		<view>
			<!-- <keep-alive>
				<component :is="comp" v-show="isShow"></component>
			</keep-alive> -->
			<one v-show="isShow1" :title='mes'></one>
			<two v-show="isShow2" :title='mes'></two>
			<three v-show="isShow3" :title='mes'></three>
		</view>
	</view>
</template>

<script>
	import one from '@/pages/home/professional/introduction.vue'
	import two from '@/pages/home/professional/line.vue'
	import three from '@/pages/home/professional/prospect.vue'
	
	export default {
		components: {
		    one:one,
		    two:two,
		    three:three,
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
				active: '',
				mes: [],
				code: 0,
			}
		},
		mounted() {
			console.log("profession外部的code是：");
			console.log(this.code)
			var _this= this;
			const on = uni.$on('code1',function(data) {
				_this.code = data.codeID;
				console.log("profession内部的code是：");
				console.log(_this.code)
			})
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
				if (value === 1){
					this.isShow1 = true
					this.isShow2 = false
					this.isShow3 = false
				}
				if (value === 2){
					this.isShow1 = false
					this.isShow2 = true
					this.isShow3 = false
				}
				if (value === 3){
					this.isShow1 = false
					this.isShow2 = false
					this.isShow3 = true
				}
				this.$route.meta.index = 0
			}
		}
	}
</script>

<style>
	.professionalMes{
		display: flex;
		flex-direction: row;
		align-items: center;
		width: 100%;
		height: 180rpx;
		background-color: #E7ECFF;
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
	.professionalName{
		font-size: 40rpx;
		font-family: "思源黑体";
		font-weight: 600;
	}
	.professionalType{
		display: flex;
		margin-top: 20rpx;
	}
	.type{
		margin-top: 5rpx;
		font-size: 30rpx;
	}
	.firstLevelDiscipline{
		margin-left: 30rpx;
	}

	.academyBar{
			display: flex;
			justify-content: center;
		}
		.barItem{
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
