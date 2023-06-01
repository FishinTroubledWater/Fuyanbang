<template>
	<view>
		<view v-for="m in mes">
			<view class="title">{{m.Title}}</view>
			<view class="user">
				<image class="headPortrait" src="@/static/academy-icons/photo.jpg"></image>
				<view class="mes">
					<view class="name">{{m.Author}}</view>
					<view class="time">{{m.PublishTime}}</view>
				</view>
			</view>
			<view class="content">
				<u-parse :content="m.Content"></u-parse>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				id: 0,
				title: "从跨专业到2021清华建筑学硕状元，初试402分，快题理论双130+，我是怎样...",
				avatarUrl: "@/static/academy-icons/photo.jpg",
				author: "咸鱼学长",
				publishTime: "2023-06-01",
				pageUrl: "@/static/academy-icons/sight.png",
				content: "写下这封长信只因为我知道作为一个跨考者，孤独是很难避免的，这种感觉真实，但它没什么坏处——无处也无需回避。如果你恰是一名“三跨”（跨专业、跨学校、跨地区）考生，那么基本不会有概率遇到同行者，考研之路必然会要你去独自走完,这其实是件挺酷的事。",
				mes: [],
			}
		},
		onLoad:function(option){
			var _this = this
			this.id = option.id;
			uni.$u.http.get('/v1/frontend/recipe/detail/' + this.id, {
				
			}).then(res => {
				this.mes = res.data.data;
				console.log("成功")
				console.log(this.mes);
			}).catch(err => {
				console.log("失败")
			})
		},
	}
</script>

<style>
.title{
	margin: 20rpx 30rpx;
	color: #000;
	font-weight: 600;
	font-size: 34rpx;
/* 	font-family: "黑体"; */
}
.user{
	display: flex;
	margin-left: 30rpx;
	margin-top: 40rpx;
/* 	justify-content: center; */
	/* align-items: center; */
}
.headPortrait{
	height: 110rpx;
	width: 110rpx;
	border-radius: 50%;
}
.mes{
	margin-left: 24rpx;
}
.name{
	font-weight: 700;
	font-size: 30rpx;
	font-family: "思源黑体";
}
.time{
	margin-top: 15rpx;
	font-size: 30rpx;
	color: #CECECE;
}
.content{
	margin: 30rpx;
	font-size: 30rpx;
	font-family: "思源黑体";
}
</style>