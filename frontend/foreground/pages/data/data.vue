<template>
	<view>
		<text>圈子</text>
		<view class="tabs-box">
			<view class="one-nav" :class="currentSwiperIndex === 0 ? 'nav-actived' : '' " @tap="swiperChange(0)">求解答
			</view>
			<view class="one-nav" :class="currentSwiperIndex === 1 ? 'nav-actived' : '' " @tap="swiperChange(1)">加油站
			</view>
		</view>
		<swiper class="swiper-box" style="height: 1000upx" :current="currentSwiperIndex">
			<swiper-item class="swiper-item sns-que">
				<div>
					<image class="appContentTop" src="../../static/background/bg1.png"></image>
					<text class="contentTopWord">21天攒图计划招募令</text>
					<navigator class="contentTopWord" url="">了解详情</navigator>
				</div>
				<div>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
					<uni-card title="linjinkun" sub-title="教育部" extra="2023年03月10日"
						thumbnail="../../static/background/bg1.png" class="trends-box-item">
						<u--text :lines="3" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
					</uni-card>
				</div>

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
						summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
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
					}


				],
			};
		},
		methods: {
			//求解答、加油站、关注切换方法
			swiperChange(index) {
				this.currentSwiperIndex = index
			},
			
		}
	}
</script>

<style lang="scss">
	.tabs-box {
		flex-flow: row;
		justify-content: space-around;
		display: flex;
	}

	.nav-actived {
		color: #0000ff;
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
	}

	.contentPic {
		object-fit: cover;
		height: 100%;
		width: 100%;
	}
</style>