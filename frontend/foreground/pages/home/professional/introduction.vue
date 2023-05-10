<template>
	<view>
		<view class="introduction">
			<view>
				<text class="introduct">{{this.name}}</text>
			</view>
		</view>
<!-- 		<view class="introductionMes" v-for="mes in introductionMes">
			<text class="mesText" space="ensp" >{{mes}}</text>
		</view> -->
		<view class="introductionMes">
			<text class="mesText" space="ensp" >{{this.profile}}</text>
		</view>
	</view>
</template>

<script>
	export default{
		data() {
			return {
				introductionMes: [
					  '    哲学是人文科学领域内的基础学科，是一门非常古老的学问，它的本质是认识和思考我们所处的世界。在希腊文中，哲学是爱智慧的意思。学哲学，就是学习智慧。对此，吉林大学哲学系教授孙正聿曾提出较深刻的见解："学习哲学，最重要的就是锻炼和培养一种善于发现问题和提出问题并进行 前提批判 的理论思维能力。哲学的爱智，无论是对自然的惊讶，还是认识人自己，都不仅仅是一种对知识的追求，更重要的是一种对生活意义的关切，对生活境界的陶冶。哲学，是使人崇高起来的学问。为学与做人，其道一也。哲学的爱智，还是一种反思的、批判的思想活动，它要追究各种知识的根据，思考历史进步的尺度，询问真善美的标准，探索生活信念的前提。"哲学一级学科，下设8个主流二级学科马克思主义哲学、中国哲学、外国哲学、逻辑学、伦理学、美学、宗教、科学技术哲学，二级学科每个学校不一定都会设置硕士点；研究方向的划分，不同院系也是根据自身的科研条件和师资力量来确立 ' ,
					 ],
				profile: '',
				name: '',
			}
		},
		mounted() {
			var _this= this;
			const on = uni.$on('code2',function(data) {
				_this.name = data.name;
				_this.profile = data.profile;
			})
		},
		methods: {
			
		}
	}
</script>

<style>
	.introduction{
		margin-top: 20rpx;
		height:70rpx;
	}
	.introduct{
		font-size: 38rpx;
		font-weight: 600;
		margin-left: 20rpx;
	}
	.introductionMes{
		margin: 0rpx 35rpx;
	}
	.mesText{
		color: #434343;
		white-space: pre-wrap;
		word-break: break-word;
	}
</style>