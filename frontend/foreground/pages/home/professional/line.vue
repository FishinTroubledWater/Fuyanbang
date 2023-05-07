<template>
	<view>
		<view class="imgList" v-for="mes in lineImgSrc">
<!-- 			<image class="lineImg" :src="mes"></image> -->
			<image class="lineImg" src="@/static/major-icons/line.png"></image>
		</view>
	</view>
</template>

<script>
	export default{
		data() {
			return {
				lineImgSrc: ["https://tse3-mm.cn.bing.net/th/id/OIP-C.azDHmgED6aBCU06PuGkgXwHaEo?pid=ImgDet&rs=1",
				"https://tse4-mm.cn.bing.net/th/id/OIP-C.UZ63fE0dy5PHLMJVMqVihwHaEo?pid=ImgDet&rs=1"],
			}
		},
	}
</script>

<style>
	.lineImg{
		height: 650rpx;
		width: 650rpx;
		margin-left: 55rpx;
		margin-top: 30rpx;
	}
</style>