<template>
	<view>
		<view class="directionText">就业方向</view>
		<view class="directionItem">
			<view v-for="direction in jobOrientation.substr(0,jobOrientation.length).split(' ')">
				<view class="Item">{{direction}}</view>
			</view>
		</view>
		
		<view class="prospectText">就业前景</view>
<!-- 		<view class="prospectMes">
			<text class="mesText" space="nbsp" >  论从地域方面，还是领域方面来看，哲学系毕业生分布都比较广泛，其主要就业方向有：
	  ①公务员。毕业生可在国家、省、市等行政管理部门从事管理或文字性工作，哲学专业的毕业生分析问题的能力较强，从事这方面的工作具有很大的优势。
	  ②文教事业或新闻出版部门。除公务员外，该专业毕业生还可到学校、科研单位或新闻出版等部门从事研究性、采编类工作，但这些单位对毕业生的学历等条件要求较高。
	  ③各类企业等。还有一部分毕业生可到企业的党办、文秘、人事管理、财务管理等部门从事各类实际工作</text>
		</view> -->
		<view class="prospectMes">
			<text class="mesText" space="nbsp" >{{this.JobProspect}}</text>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				JobProspect: '',
				jobOrientation: '',
			}
		},
		mounted() {
			var _this= this;
			const on = uni.$on('code2',function(data) {
				_this.JobProspect = data.JobProspect;
				_this.jobOrientation = data.jobOrientation;
			})
		},
		methods: {
			
		}
	}
</script>

<style>
	.directionText{
		font-size: 38rpx;
		font-weight: 600;
		margin-top: 20rpx;
		height: 80rpx;
		margin-left: 20rpx;
	}
	.directionItem{
		display: flex;
		align-items: center;
	}
	.Item{
		width: auto;
		height: 50rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: #EDF3FF;
		font-size: 26rpx;
		color: #1579ec;
		font-family: "思源黑体";
		padding: 4rpx 18rpx;
		margin-left: 25rpx;
		font-size: 30rpx;
		border-style:solid;
		border-width:2rpx;
		border-color: #BBBBBB;
	}
	.prospectText{
		font-size: 38rpx;
		font-weight: 600;
		margin-top: 30rpx;
		height: 80rpx;
		margin-left: 20rpx;
	}
	.prospectMes{
		margin: 0rpx 30rpx;
	}
	.mesText{
		margin-left: 20rpx;
		color: #434343;
		white-space: pre-wrap;
		word-break: break-word;
	}
</style>