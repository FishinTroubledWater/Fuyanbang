<template>
	<view>
		<text class="quanzi">圈子</text>
		<view class="tabs-box">
			<view class="one-nav" :class="currentSwiperIndex === 0 ? 'nav-actived' : '' " @tap="swiperChange(0)">加油站
			</view>
			<view class="one-nav" :class="currentSwiperIndex === 1 ? 'nav-actived' : '' " @tap="swiperChange(1)">求解答
			</view>
		</view>
		<swiper class="swiper-box" style="height: 1000upx" :current="currentSwiperIndex">
			<swiper-item class="swiper-item sns-que">
				<view class="topic">
					<text class="contentTopWord">21天攒图计划招募令</text>
					<text>\n</text>
					<text style="color:blanchedalmond;">每日一画，不画就出局
							来吧，一起坚持21天，做行动派！</text>
					<navigator class="contentTopWord" url="/pages/data/activityDetail/activityDetail">了解详情</navigator>
									
				</view>
					
					<u-list>
						<u-list-item v-for="(item, index) in indexList" :key="index">
							<uni-card @click="clicknews(index)" title="linjinkun" sub-title="教育部" extra="2023年03月10日"
								thumbnail="../../static/background/bg1.png" class="trends-box-item" >
								<u--text :lines="3" :text="indexList[index].summary"></u--text>
								<image class="newsimage" src="../../static/background/activityDetails.png"></image>
							</uni-card>
						</u-list-item>
					</u-list>
				

			</swiper-item>
			<swiper-item class="swiper-item sns-oil">
				<u-list>
					<u-list-item v-for="(item, index) in indexList" :key="index">
						<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
							thumbnail="../../static/background/bg1.png" class="trends-box-item">
							<u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
						</uni-card>
					</u-list-item>
				</u-list>
			</swiper-item>

		</swiper>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				currentSwiperIndex: 0,
				list1: [{
					name: '求解答',
				}, {
					name: '加油站',
				}],
				
				indexList: [{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					}, {
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123457',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					},{
						name: '1',
						time:'2022-12-21',
						icon:'',
						newId:'123456',
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
						isImage:true,
						img:['www.baidu.com',
						'www.baidu.com',
						'www.baidu.com']
					}


				],
				commentList:[{
					name:'吴彦祖',
					icon:'',
					content:'回复内容',
					time:'2022-12-21'
				},{
					name:'吴彦祖',
					icon:'',
					content:'回复内容',
					time:'2022-12-21'
				}
					
				]
			};
		},
		methods: {
			//求解答、加油站、关注切换方法
			swiperChange(index) {
				this.currentSwiperIndex = index
			},
			clicknews(index){
				// console.log('1'),
				uni.navigateTo({
					url:'/pages/data/newsDetails/newsDetails?id='+index,
				})
			}
			
		}
	}
</script>

<style lang="scss">
	.newsimage{
		max-height: 200rpx;
		max-width: 200px;
	}
	.quanzi{
		font-size: 50rpx;
		font-weight: 800;
	}
	.tabs-box {
		flex-flow: row;
		justify-content: space-around;
		display: flex;
	}

	.nav-actived {
		color:blanchedalmond;
		font-weight: 700;
	}

	.appContentTop {
		z-index: -1;
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 200rpx;
	}

	.contentTopWord {
		color: #ffffff;
		font-size: 40rpx;
	}

	.contentPic {
		object-fit: cover;
		height: 100%;
		width: 100%;
	}
	.topic{
		background:url("../../static/background/bg2.png") center center;
	}
</style>