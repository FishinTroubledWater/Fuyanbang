<template>
	<view>
		<uni-card :title="indexList.name" sub-title="问题详情" :extra="indexList.time" :thumbnail="indexList.icon"
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
			<textarea class="textarea" placeholder="说说你的看法吧,在此处输入回答." placeholder-style="font-size:28rpx"
				maxlength="200" @input="descInput" v-model="desc" />
			<view class="num">{{ desc.length }}/200</view>
			<button @click="clickSent">发送回答</button>
		</view>


		<text style="font-size: 40rpx; font-weight: 800;">回答:</text>
		<uni-card v-for="(item, index) in answer" :title="item.name" :sub-title="item.time"
			:thumbnail="item.icon" class="trends-box-item">
			<u--text :text="item.answer"></u--text>
		</uni-card>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				queID:'',
				whetherLike:'false',
				desc: '',
				myanswer: 'null',

				txt: "txt",
				academyName: '福州大学',
				indexList: {
					name: 'zhang',
					time: '2022-12-21',
					icon: '../../../static/background/activityDetails.png',
					queId: '123456',
					summary: `<p>露从今夜白</p>
					<img src="../../static/background/activityDetails.png" />`,
					// isImage: true,
					// img: ['../../../static/background/activityDetails.png',
					// 	'../../../static/background/bg1.png',
					// 	'../../../static/background/bg2.png'
					// ],
				},
				answer: [],

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
				this.myanswer = e.detail.value
			},
			clickSent() {
				console.log(this.myanswer)
				//post请求

			},

		},
		mounted() {
			uni.$u.http.get('/v1/frontend/circle/queAnswer/3' , {
		
				}).then(res => {
					console.log(res.data.data);
					this.answer=res.data.data;
				}).catch(err => {
		
				})
			// uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.queID, {
			
			// 	}).then(res => {
			// 		console.log(res.data.data);
			// 		this.indexList=res.data.data[0];
			// 	}).catch(err => {
			
			// 	})
				
				// uni.$u.http.post('/v1/frontend/academy/searchByRule', {
				// 	region: '福建',
				// 	level: '985',
				// 	type: '法学',
				// }).then(res => {
				// 	console.log(res.data)
				// }).catch(err => {
		
				// })
		},

		onLoad: function(option) {
			console.log(option.id)
			this.queID=option.id
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
</style>