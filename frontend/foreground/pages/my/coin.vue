<template>
	<view class="content">
		<view class="box1">
			<text class="coin">
				{{balance}}
			</text>
			<text class="balen">
				学币余额
			</text>
		</view>
		<view class="box2"></view>
	</view>
</template>

<script>
	export default{
		data(){
			return{
				balance:'100'
			}
		},
		methods:{
			
		}
	}
</script>

<style>
	.content{

		display: flex;
		flex-direction: column;
	}
	.box1{
		height: 250px;
		background-color: aqua;
		
	}
	.box2{
		height: 500px;
		background-color: bisque;
		
		
	}
	.coin{
		
	}
	.balen{
		
	}
</style>