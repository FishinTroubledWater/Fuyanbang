<template>
	<view>
		<uni-card :title="indexList.name" sub-title="帖子详情" :extra="indexList.time" :thumbnail="indexList.icon"
			class="trends-box-item">
			<u--text :text="indexList.summary"></u--text>
			<image v-for="(item, index) in indexList.img" :src="indexList.img[index]"></image>

		</uni-card>
		<!-- <view class="content">
			<textarea class="uni-title uni-common-pl" v-model="txt"></textarea>
		</view> -->

		<view class="textarea_box">
			<textarea class="textarea" placeholder="说说你的看法吧,在此处输入评论." placeholder-style="font-size:28rpx"
				maxlength="200" @input="descInput" v-model="desc" />
			<view class="num">{{ desc.length }}/200</view>
			<button @click="clickSent">发送评论</button>
		</view>


		<text style="font-size: 40rpx; font-weight: 800;">评论:</text>
		<uni-card v-for="(item, index) in comment" :title="comment[index].name" :sub-title="comment[index].time"
			:thumbnail="comment[index].icon" class="trends-box-item">
			<u--text :text="comment[index].content"></u--text>
		</uni-card>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				desc: '',
				myComment: 'null',

				txt: "txt",
				academyName: '福州大学',
				indexList: {
					name: 'zhang',
					time: '2022-12-21',
					icon: '../../../static/background/activityDetails.png',
					postId: '123456',
					summary: '近日，福州大学爆出，今年福大考研学生取得前所未有的成绩，尤其计算机与大数据学院，计算机与大数据学院的考研学生们表现优秀。',
					isImage: true,
					img: ['../../../static/background/activityDetails.png',
						'../../../static/background/bg1.png',
						'../../../static/background/bg2.png'
					],
				},
				comment: [{
					name: '吴彦祖',
					icon: '../../../static/background/bg2.png',
					content: '评论内容',
					time: '2022-12-21'
				}, {
					name: '吴彦祖',
					icon: '../../../static/background/bg1.png',
					content: '评论内容',
					time: '2022-12-21'
				}],

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
			descInput(e) {
				console.log(e.detail.value.length, '输入的字数')
				this.myComment = e.detail.value
			},
			clickSent() {
				console.log(this.myComment)
				//post请求

			},

		},

		onLoad: function(option) {
			console.log(option.id)
		},

		mounted() {
			uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {

				}).then(res => {
					console.log(res.data);
				}).catch(err => {

				}),
				uni.$u.http.post('/v1/frontend/academy/searchByRule', {
					region: '福州',
					level: '985',
					type: '法学',
				}).then(res => {
					console.log(res.data)
				}).catch(err => {

				})
		}
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