<template>
	<view class="content">
		<u-popup :show="showGet" mode="center" @close="close" >
			<view class="pop">
				<text class="title">当前获取方式</text>
				<hr class='divider' />
				<text class="ans">1.每日签到可获得1学币</text>
			</view>
		</u-popup>
		<view class="box1">
			<text class="coin">
				{{balance}}
			</text>
			<text class="balen">
				学币余额
			</text>
			<button class="getBtn" @click="getCoin()">获得</button>
		</view>
		<hr class='divider' />
		<view class="box2">
			<text class="title">什么是学币</text>
			<text class="ans">学币是仅限于在福研帮中使用的虚拟货币</text>
			<text class="title" id="t1">如何使用学币</text>
			<text class="ans">1.学币是仅限于在福研帮中使用的虚拟货币</text>
			<text class="ans">2.学币是仅限于在福研帮中使用的虚拟货币</text>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showGet:false,
				balance: '100'
			}
		},
		methods: {
			getCoin(){
				this.showGet=true;
			},
			close() {
				this.showGet=false;
			},

		}
	}
</script>

<style>
	.content {

		display: flex;
		flex-direction: column;
	}

	.box1 {
		height: 250px;
		display: flex;
		flex-direction: column;
		text-align: center;

	}

	.box2 {
		height: 500px;
		display: flex;
		flex-direction: column;
	}
	.pop{
		width: 300px;
		height: 150px;
		display: flex;
		flex-direction: column;
	}
	.title{
		margin-top: 10%;
		margin-left: 8%;
		font-size: 20px;
		font-weight:bold;
	}
	.ans{
		margin-top: 10px;
		margin-left: 8%;
	}

	.coin {
		font-size: 100upx;
		margin-top: 50px;
		color: #f9bd10;
	}

	.balen {
		font-size: 30upx;
		margin-top: 10px;
		color: #d1c2d3;
	}
	.getBtn{
		margin-top: 20px;
		width: 40%;
		border-radius: 50px;
		background-color:#f4ce69;
	}
	 .divider {
		 margin: 12px;
	        align-self: center;
	        width: 90%;
	        color: #987cb9;
	        size: 10;
	    }
</style>