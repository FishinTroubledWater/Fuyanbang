<template>
	<view>
		<uni-card :title="indexList.name" sub-title="帖子详情" :extra="indexList.time" :thumbnail="indexList.icon"
			class="trends-box-item">
			<view class="u-content">
				<u-parse :content="indexList.summary"></u-parse>
			</view>

		</uni-card>
		<!-- <view class="content">
			<textarea class="uni-title uni-common-pl" v-model="txt"></textarea>
		</view> -->
		<u-icon v-if="whetherLike==='false'" style="padding-left: 50rpx;" label="收藏" color="#2979ff" size="20" name="star" @click="clickLike"></u-icon>
		<u-icon v-if="whetherLike==='true'" style="padding-left: 50rpx;" label="收藏" color="#2979ff" size="20" name="star-fill" @click="clickLike"></u-icon>
		<view class="textarea_box">
			<textarea class="textarea" placeholder="说说你的看法吧,在此处输入评论." placeholder-style="font-size:28rpx"
				maxlength="200" @input="descInput" v-model="desc" />
			<view class="num">{{ desc.length }}/200</view>
			<button @click="clickSent">发送评论</button>
		</view>


		<text style="font-size: 40rpx; font-weight: 800;">评论:</text>
		<uni-card v-for="(item, index) in comment" :title="item.name" :sub-title="item.time"
			:thumbnail="item.icon" class="trends-box-item">
			<u--text :text="item.content"></u--text>
		</uni-card>
	</view>
</template>

<script>
	import uCard from '../../../uni_modules/uni-card/uni-card.vue'
	import uSteps from '../../../uni_modules/uni-steps/uni-steps.vue'
	import uIcons from '../../../uni_modules/uni-icons/uni-icons.vue'
	import uSection from '../../../uni_modules/uni-section/uni-section.vue'
	export default {
		components: {
		  uCard,
		  uSteps,
		  uIcons,
		  uSection
		},
		data() {
			return {
				id:'',
				postId:'0',
				whetherLike:'false',
				desc: '',
				myComment: 'null',

				txt: "txt",
				academyName: '福州大学',
				indexList: {
					// name: 'zhang',
					// time: '2022-12-21',
					// icon: '../../../static/background/activityDetails.png',
					// postId: '123456',
					// summary: `<p>露从今夜白</p>
					// <img src="../../static/background/activityDetails.png" />`,
					// // isImage: true,
					// // img: ['../../../static/background/activityDetails.png',
					// // 	'../../../static/background/bg1.png',
					// // 	'../../../static/background/bg2.png'
					// // ],
				},
				comment: [],

			}
		},
		watch: {
			txt(txt) {
				if (txt.indexOf('\n') != -1) { //敲了回车键了
					uni.hideKeyboard() //收起软键盘
				}
			}
		},
		methods: {
			clickLike(){
				if(this.whetherLike==='true'){
					this.whetherLike='false'
				}else{
					this.whetherLike='true'
				}
				
			},
			descInput(e) {
				console.log(e.detail.value.length, '输入的字数')
				this.myComment = e.detail.value
			},
			clickSent() {
				console.log(this.myComment)
				uni.getStorage({
					key:'userId',   // 储存在本地的变量名
					success:res => {
						// 成功后的回调
						// console.log(res.data);   // hello  这里可做赋值的操作
						this.id=res.data;
						console.log(this.id)
					}
				})
				uni.$u.http.post('/v1/frontend/circle/postComment', {
					postId: this.postId,
					userId: this.id,
					comment: this.myComment,
				}).then(res => {
					console.log(res.data)
					if(res.data.code==200){
						uni.showToast({
							title:"评论成功",
							duration:1000,
						})
					}
				}).catch(err => {
				
				})

			},

		},

		onLoad: function(option) {
			console.log(option.id)
			this.postId=option.id
		},

		mounted() {
			uni.$u.http.get('/v1/frontend/circle/newinfoComment/' + this.postId, {

				}).then(res => {
					console.log(res.data.data);
					this.comment=res.data.data;
				}).catch(err => {

				}),
			uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.postId, {
			
				}).then(res => {
					console.log(res.data.data);
					this.indexList=res.data.data[0];
				}).catch(err => {
			
				})
				
				// uni.$u.http.post('/v1/frontend/academy/searchByRule', {
				// 	region: '福建',
				// 	level: '985',
				// 	type: '法学',
				// }).then(res => {
				// 	console.log(res.data)
				// }).catch(err => {

				// })
		},
	}
</script>

<style>
	.textarea_box {
		padding: 20rpx;
		background-color: #F2F2F2;

		/deep/ .uni-textarea-textarea {
			font-size: 28rpx;
			line-height: 45rpx;
		}

		.num {
			text-align: right;
			color: gray
		}
	}
	.textarea{
		width: 100%;
		height: 100rpx;
	}
</style>