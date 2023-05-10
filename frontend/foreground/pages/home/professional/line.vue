<template>
	<view>
<!-- 		<view class="imgList" v-for="mes in lineImgSrc">
			<image class="lineImg" src="@/static/major-icons/line.png"></image>
		</view> -->
		<view class="imgList">
			<image class="lineImg" :src="this.scoreUrl" @click="imgClick(scoreUrl)"></image>
<!-- 			<view v-for="m in mes.list">
				<image class="lineImg" :src="m"></image>
			</view> -->
		</view>
	</view>
</template>

<script>
	export default{
		data() {
			return {
				lineImgSrc: ["https://tse3-mm.cn.bing.net/th/id/OIP-C.azDHmgED6aBCU06PuGkgXwHaEo?pid=ImgDet&rs=1",
				"https://tse4-mm.cn.bing.net/th/id/OIP-C.UZ63fE0dy5PHLMJVMqVihwHaEo?pid=ImgDet&rs=1"],
				scoreUrl: '',
				mes: [],
			}
		},
		onLoad() {
		},
		mounted() {
			var _this= this;
			const on = uni.$on('code2',function(data) {
				_this.scoreUrl = data.scoreUrl;
				console.log(_this.scoreUrl)
			})
			const oMeta = document.createElement('meta');
			oMeta.name = "referrer";
			oMeta.content = "no-referrer"
			document.getElementsByTagName('head')[0].appendChild(oMeta);
		},
		methods: {
			imgClick(url){
				console.log(url)
				uni.previewImage({
					urls: [url]
				});
			},
		},
		created() {
			uni.$u.http.post('/v1/frontend/academy/score', {
				type: '学科门类',firstLevelDiscipline: '一级学科',secondLevelDiscipline: '二级学科',
			}).then(res => {
				this.mes = res.data.data;
				console.log(this.mes)
			}).catch(err => {
				this.mes = [];
			})
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