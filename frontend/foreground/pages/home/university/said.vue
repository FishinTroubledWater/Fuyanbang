<template>
	<view>
		<view v-for = "m in mes1">
			<view class="viewSaid" @click="goPost(m.ID)">
				<view class="viewUser">
					<image class="headPortrait" src="@/static/academy-icons/photo.jpg"></image>
					<view class="userMes">
						<text class="userName">{{m.Author}}</text>
					</view>
				</view>
				<view class="saidContent">
					<u--text :lines="3" :text="m.Content" size="24rpx" lineHeight="34rpx" margin="20rpx 6rpx 0rpx 0rpx" padding="6px 4px"
						height="94px"></u--text>
					<image class="sights" :src="m.PageUrl"></image>
				</view>
			</view>
		</view>
			
<!-- 			<view class="viewSaid">
				<view class="viewUser">
					<image class="headPortrait" src="@/static/academy-icons/photo.jpg"></image>
					<view class="userMes">
						<text class="userName">玲玲学姐</text>
					</view>
				</view>
				<view class="saidContent">

					<view class="textContent">{{this.content}}</view>
					<image class="sights" src="@/static/academy-icons/sight.png"></image>

				</view>
			</view> -->
	</view>
</template>

<script>
	export default{
		data() {
			return {
				mes1: [],
				code: 1,
				content: "致所有考研的学弟学妹，见字如面：写下这封长信只因为我知道作为一个跨考者，孤独是很难避免的，"
			}
		},
		props:['title'],
		methods: {
			goPost(postId) {
				uni.navigateTo({
					url: "/pages/home/university/post?id=" + postId
				})
			},
		},
		onShow() {
			// console.log('eee')
		},
		created() {	
			var _this = this;
			const on = uni.$on('code1', function(data) {
				_this.mes1 = data.saidMes;
				console.log("said内部的Mes是：");
				console.log(_this.mes1)
				
				// uni.$u.http.get('/v1/frontend/recipe/list/1', {
					
				// }).then(res => {
				// 	this.mes = res.data.data;
				// 	console.log("成功")
				// 	console.log(this.mes);
				// }).catch(err => {
				// 	console.log("失败")
				// })
			})
			
		},
		mounted() {
			// uni.$u.http.get('/v1/frontend/recipe/list/1', {
				
			// }).then(res => {
			// 	this.mes = res.data;
			// 	console.log("成功")
			// 	console.log(this.mes);
			// }).catch(err => {
			// 	console.log("失败")
			// })
		},
	}
</script>

<style lang="scss">
.viewUser{
	display: flex;
	flex-direction: row;
	align-items: center;
}
.viewSaid {
	height: auto;
	  
	/* 圆角 */
	border-radius: 18rpx;
	
	/* 边 */
	border: 1rpx solid #E0E3DA;
	/* 阴影 */
	box-shadow:2rpx 7rpx 0rpx #ebebeb;
	
	background-color: #ffffff;
	margin-left:30rpx;
	margin-right:30rpx;
	margin-top: 25rpx;
	
	/* padding使得文字和图片不至于贴着边框 */
	padding: 25rpx;
	
}
.headPortrait{
	height: 110rpx;
	width: 110rpx;
	border-radius: 50%;
}
.userMes {
	margin-left: 30rpx;
	display: flex;
	flex-direction: column;
}
.userName{
	font-size: 30rpx;
	font-family: "黑体";
}
.typeOfAdmission{
	margin-top: 10rpx;
	width: auto;
	display: flex;
	justify-content: center;
	align-items: center;
	border-radius: 18rpx;
	background-color: #96C5F1;
	font-size: 24rpx;
	color: #ffffff;
	font-family: "黑体";
}
.textOfAdmission{
	margin: 10rpx;
}
.lacation{
	margin-left: 30rpx;
	width: 200rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	font-family: "黑体";
	font-size: 36rpx;
}
.saidContent{
	display: flex;
	flex-direction: row;
	align-items: center;
}
.textContent{
	width: 450rpx;
	margin-top: 30rpx;
	margin-right: 20rpx;
	font-size: 20rpx;
}
.sights{
	float: right;
	width: 190rpx;
	height: 125rpx;
	border-radius: 18rpx;
}
.readingQuantity{
	margin-top: 30rpx;
	font-size: 26rpx;
	color: #CECECE;
}
</style>