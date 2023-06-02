<template>
	<view>
		<!-- <view class="mes">
			<view class="academyMes">
				<view class="type">院校类型</view>
				<view class="typeName">{{title.Type}}</view>
			</view>
			<view class="academyMes">
				<view class="type">院校隶属</view>
				<view class="typeName">{{title.Belong}}</view>
			</view>
		</view>
		<view class="introduction">
			<view>
				<text class="introduct">院校简介</text>
				<text class="code">院校代码：{{title.Code}} {{title.Region}}</text>
			</view>
		</view> -->
		<!-- 		<view class="introductionMes" v-for="Mes in introductionMes">
			<text class="mesText" space="ensp" >{{Mes}}</text>
		</view> -->
		<!-- 		<view class="introductionMes">
			<text class="mesText" space="ensp" >{{title.Profile}}</text>
		</view> -->

		<view>
			<u-row >
				<u-col span="1">
					
				</u-col>
				<u-col span="5">
					<view class="academyMes">
						<view class="type">院校类型</view>
						<view class="typeName">{{this.type}}</view>
					</view>
				</u-col>
				<u-col span="5">
					<view class="academyMes">
						<view class="type">院校隶属</view>
						<view class="typeName">{{this.belong}}</view>
					</view>
				</u-col>
				<u-col span="1">
					
				</u-col>
<!-- 				<view class="mes">
					<view class="academyMes">
						<view class="type">院校类型</view>
						<view class="typeName">{{this.type}}</view>
					</view> -->
					
				

			</u-row>
			<view class="introduction">
				<view>
					<text class="introduct">院校简介</text>
					<text class="code">院校代码：{{this.code}} {{this.region}}</text>
				</view>
			</view>
			<!-- 		<view class="introductionMes" v-for="Mes in introductionMes">
						<text class="mesText" space="ensp" >{{Mes}}</text>
					</view> -->
			<view class="introductionMes">
				<text class="mesText" space="ensp">{{this.profile}}</text>
			</view>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				introductionMes: [
					'     北京大学（英语：Peking University，缩写：PKU），简称北大，创建于1898年，初名京师大学堂，成立之初为中国最高学府，同时也是国家最高教育行政机关，行使教育部职能，统管全国教育。中华民国建立后，校名改为北京大学校，后又改名为国立北京大学。1916年—1927年，蔡元培任北京大学校长时期，“循思想自由原则、取兼容并包之义”，推行改革，把北大办成以文、理两科为重点的综合性大学，使北京大学成为新文化运动的中心、五四运动的策源地。1937年抗日战争爆发后，北大与清华、南开合并组建国立西南联合大学。日本战败后，北京大学回到北平复校。1952年院系调整后，北京大学聚集了原北大、清华、燕大三校的自然科学、人文学者，奠定了北大文理学科在中国高校中长期领先的地位。2000年，北京医科大学回归北京大学，成为北京大学医学部。',
					'     北京大学是中华人民共和国最顶尖的高校之一，现时属于“QS世界百强大学”、“泰晤士高等教育世界百强大学”及“美国新闻与世界报道世界百强大学”，是“双一流A类”和原“985工程”、原“211工程”重点建设大学。如今，北京大学是一所兼具自然科学、人文学科、社会科学、医学以及新型工科的综合性大学，也是大中华地区和亚太地区综合实力最强的高等教育机构之一。北京大学现有六个校区。北京大学校址原本在沙滩红楼、公主府等地，1952年院系调整后迁入位于北京西北郊的原燕京大学校园（燕园），燕园兼收中国古典韵味和西方校园规划理念，主要景观有未名湖、博雅塔等。',
					'     2022年6月9日，在2023QS世界大学排名中，北京大学在中国(大陆)排名最高，排名上升6位，至全球第12位'
				],
				mes: [],
				belong: '',
				code: 0,
				type: '',
				region: '',
				profile: '',


			}
		},
		props: ['title'],
		// onShow() {
		// 	// console.log('bbb')
		// 	// uni.$on('code',res=>{
		// 	// 	console.log('aaa')
		// 	// 	console.log(this.code)
		// 	// })
		// },
		methods: {

		},
		mounted() {
			// const on = uni.$once('code1',res=>{
			// 	_this.code = res
			// 	console.log("内部的code是：");
			// 	console.log(_this.code)
			// 	// uni.$u.http.get('/v1/frontend/academy/detail/2010', {

			// 	// }).then(res => {
			// 	// 	this.mes = res.data.data;
			// 	// 	console.log("成功2")		
			// 	// 	console.log(this.mes)
			// 	// 	console.log("MMMM")	
			// 	// }).catch(err => {
			// 	// 	console.log("失败")
			// 	// })
			// })

			// console.log("KKKHH");
			// console.log(this.code)
			var _this = this;

			const on = uni.$on('code1', function(data) {
				_this.code = data.codeID;
				_this.belong = data.belong;
				_this.region = data.region;
				_this.type = data.type;
				_this.profile = data.profile;
				console.log("introduction内部的code是：");
				console.log(_this.code)
			})

			// setTimeout(function() {
			// 	this.code = on.code
			// 	console.log("外部的code是");
			// 	console.log(this.code)
			// 	uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {

			// 	}).then(res => {
			// 		this.mes = res.data.data;
			// 		console.log("成功2")		
			// 		console.log(this.mes)
			// 		console.log("MMMM")	
			// 	}).catch(err => {
			// 		console.log("失败")
			// 	})
			// }, 200)

			// this.code = on.code

			// console.log("外部的code是");
			// console.log(this.code)

			// uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {

			// }).then(res => {
			// 	this.mes = res.data.data;
			// 	console.log("成功2")		
			// 	console.log(this.mes)
			// 	console.log("MMMM")	
			// }).catch(err => {
			// 	console.log("失败")
			// })


		},
		onUnload() {
			// uni.$off('code1')
		},
	}
</script>

<style>
	.mes {
/* 		display: flex;
		justify-content: center; */
		height: 220rpx;
	}

	.academyMes {
		/* display: flex; */
		/* justify-content: center; */
		width: 250rpx;
		height: 170rpx;
		border: 2rpx solid #BBBBBB;
		margin: 20rpx;
		border-radius: 18rpx;
	}

	.type {
		display: block;
		text-align: center;
		width: 200rpx;
		margin-top: 30rpx;
		font-size: 30rpx;
		color: #767676;
	}

	.typeName {
		display: block;
		text-align: center;
		margin-top: 5rpx;
		font-size: 38rpx;
		font-weight: 600;
	}

	.introduction {
		height: 70rpx;
	}

	.introduct {
		font-size: 38rpx;
		font-weight: 600;
		margin-left: 20rpx;
	}

	.code {
		font-size: 30rpx;
		float: right;
		margin-right: 20rpx;
		color: #767676;
	}

	.introductionMes {
		margin: 0rpx 20rpx;
	}

	.mesText {
		color: #434343;
		white-space: pre-wrap;
		word-break: break-word;
	}
</style>