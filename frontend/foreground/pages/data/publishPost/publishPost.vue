<template>
	<view>
		<view>
			<input class="uniInput" placeholder-style="color:#c7c7c7" focus placeholder="标题 (1-30字之间) "
				v-model="title" />
		</view>
		<view class="container">

			<editor id="editor" class="qlContainer" :placeholder="placeholder" @input="getText" @ready="onEditorReady"
				:adjust-position="false"></editor>

		</view>
		<view class="below">
			<view class="selectForm">
				<picker @change="bindPickerChange1" :range="array1" :value="index1" class="selectFormItem">
					<label class="">{{array1[index1]}}</label>
					<!-- <label class="downArrow"> ∨ </label> -->
					<label class="down">
						<image class="downArrow" src="@/static/academy-icons/down.png"></image>
					</label>
				</picker>
				<view class="inputNum" v-show="isShow">
					<label>请输入悬赏的金额：</label>
					<input type='number' v-model="num" class="intext" />
				</view>
			</view>
			<view class="operator">
				<button @click="chooseImage1()" class="uploadingBtn">上传图片</button>
				<button @click="publish()" class="publishBtn">发布</button>
			</view>
		</view>
	</view>

</template>

<script>
	import Axios from 'axios'
	Axios.defaults.baseURL = '/'
	// eslint-disable-next-line no-unused-vars
	const axios = require('axios')

	import {
		pathToBase64,
		base64ToPath
	} from '@/js/image-tools/index.js'

	export default {
		data() {
			return {
				id: 0,
				placeholder: '开始输入文章内容...',
				title: '',
				context: '',
				synopsis: '',
				array1: ['请选择文章类型', '加油站', '求解答'],
				index1: 0,
				type: 0,
				baseImageList: [],
				img: '@/static/academy-icons/sight.png',
				reward: 0,
				num: 0,
				isShow: false,
				coin: 0,
			}
		},
		methods: {
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg = this.array1[this.index1];
				// this.type = this.array1[this.index1];
				this.type = this.index1;
				if (this.index1 == 2) this.isShow = true
				else this.isShow = false
			},
			onEditorReady() {
				uni.createSelectorQuery().select('#editor').context((res) => {
					this.editorCtx = res.context
				}).exec()
			},
			undo() {
				this.editorCtx.undo()
			},
			getText(e) {
				// console.log(e.detail.html);//带标签内容
				this.context = e.detail.html;
				this.synopsis = e.detail.text;
				console.log(this.context); //带标签内容
				console.log(this.synopsis);
			},

			publish() {
				if (this.title == "") {
					console.log("请先输入文章标题");
					uni.showToast({
						title: '请先输入文章标题',
						icon: 'none', //如果要纯文本，不要icon，将值设为'none'
						duration: 2000 //持续时间为 2秒
					})
				} else if (this.title.length <= 0 || this.title.length >= 31) {
					console.log("标题的字数请控制在1-30之间");
					uni.showToast({
						title: '标题的字数请控制在1-30之间',
						icon: 'none', //如果要纯文本，不要icon，将值设为'none'
						duration: 2000 //持续时间为 2秒
					})
				} else if (this.context == "") {
					console.log("请输入文章内容");
					uni.showToast({
						title: '请输入文章内容',
						icon: 'none', //如果要纯文本，不要icon，将值设为'none'
						duration: 2000 //持续时间为 2秒
					})
				} else if (this.index1 == 0) {
					console.log("请先选择文章类型");
					uni.showToast({
						title: '请先选择文章类型',
						icon: 'none', //如果要纯文本，不要icon，将值设为'none'
						duration: 2000 //持续时间为 2秒
					})
				} else {
					this.num = parseInt(this.num);
					if (this.index1 == 2 && this.num < 0) {
						console.log("悬赏的金额不能小于0");
						uni.showToast({
							title: '悬赏的金额不能小于0',
							icon: 'none', //如果要纯文本，不要icon，将值设为'none'
							duration: 2000 //持续时间为 2秒
						})
					} else if (this.coin < this.num) {
						console.log("您的学币不足");
						uni.showToast({
							title: '您的学币不足',
							icon: 'none', //如果要纯文本，不要icon，将值设为'none'
							duration: 2000 //持续时间为 2秒
						})
					} else {
						this.reward = parseInt(this.num);
						console.log("作者id：" + this.id);
						console.log("文章类型：" + this.type);
						console.log("标题：" + this.title);
						console.log("内容：" + this.context);
						console.log("图片地址：" + this.img)
						console.log("简介：" + this.synopsis)
						console.log("悬赏数额：" + this.reward)
						uni.$u.http.post('/v1/frontend/circle/uploadPost', {
							userId: this.id,
							title: this.title,
							content: this.context,
							type: this.type,
							summary: this.synopsis,
							img: this.img,
							reward: this.reward,
						}).then(res => {
							console.log("发布成功")
							uni.showToast({
								title: '发布成功',

								icon: 'none', //如果要纯文本，不要icon，将值设为'none'
								duration: 2000 //持续时间为 2秒
							})
							uni.navigateBack({
								// delta: 1
							})
						}).catch(err => {
							console.log("发布失败")
							uni.showToast({
								title: '发布失败',
								icon: 'none', //如果要纯文本，不要icon，将值设为'none'
								duration: 2000 //持续时间为 2秒
							})
						})
					}

				}
			},
			async chooseImage1() {
				var _this = this;
				// uni.chooseImage({
				//     count: 1, //默认9
				//     sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
				//     sourceType: ['album'], //从相册选择
				//     success: function (res) {
				//         // console.log(JSON.stringify("图片地址：" + res.tempFilePaths[0]));
				// 		var data = JSON.stringify(res.tempFilePaths[0]);
				// 		data = data.substr(6,data.length);
				// 		console.log("图片地址：" + data);
				// 		_this.editorCtx.insertImage({
				// 			width: '100%', //设置宽度为100%防止宽度溢出手机屏幕
				// 			height: 'auto',
				// 			// src: 'https://pic35.photophoto.cn/20150511/0034034892281415_b.jpg', //服务端返回的url
				// 			src: "/static/my-assets/taiku.png",
				// 			alt: '图像',
				// 		})
				//     }
				// });

				// const res = await uni.chooseImage({
				// 	count: 1,
				// 	success: (res) => {
				// 		_this.editorCtx.insertImage({
				// 			width: '100%', //设置宽度为100%防止宽度溢出手机屏幕
				// 			height: 'auto',
				// 			src: res.tempFilePaths[0],
				// 			alt: '图像',
				// 		})
				// 	}
				// });

				// const imagePath = res.tempFilePaths[0];

				// const fileInfo = await uni.getFileInfo({
				//     filePath: imagePath, // 图片文件路径
				// 	success: (res) => {
				// 		// console.log("kkkkkkkk");
				// 	}
				// });
				// // console.log("kkkkkkkk");
				//   // 将图片文件转换为File对象
				// const file = new File([imagePath], fileInfo.fileName, {
				// 	type: fileInfo.type, // 设置文件类型
				// });
				// // console.log("kkkkkkkk");
				// console.log("图片的file文件：" + file);

				try {
					// 调用uni.chooseImage方法选择本地图片文件

					const res = await new Promise((resolve, reject) => {
						uni.chooseImage({
							count: 1, // 选择图片的数量，这里选择1张
							sourceType: ['album'], // 选择图片的来源，这里选择相册
							success: resolve,
							fail: reject,
						});
					});

					// 从返回结果中获取选中的图片文件路径
					const imagePath = res.tempFilePaths[0];



					const res1 = pathToBase64(imagePath)
						.then(base64 => {
							// console.log(base64)
							this.baseImageList = this.baseImageList.concat(base64);
							var s = base64.substr(base64.indexOf(',') + 1, base64.length);
							console.log(s)
							uni.$u.http.post('https://api.superbed.cn/upload', {
								token: '1000766339bd4c248a8ad625a87f687d',
								b64_data: s,
							}).then(res => {
								console.log("图片上传成功");
								console.log(res.data.url);
								_this.img = res.data.url;
								_this.editorCtx.insertImage({
									width: '100%', //设置宽度为100%防止宽度溢出手机屏幕
									height: 'auto',
									src: res.data.url,
									alt: '图像',
								})
							}).catch(error => {
								console.log(error);
								console.log("失败")
							})
						})
						.catch(error => {
							console.error(error)
						})


					// const uploadRes = await new Promise((resolve, reject) => {
					//       uni.uploadFile({
					//         url: 'https://sm.ms/api/v2/upload', // 图床服务器的URL地址
					//         filePath: "imagePath", // 图片文件路径
					//         name: 'smfile', // 上传文件的字段名
					//         headers: {
					//             'Authorization': 'giKtHZm1SL0Paj1l7Ye0lrQ0Ilq6VuxX', // 添加Authorization字段，替换your_token_here为实际的token值
					//         },
					//         success: resolve,
					//         fail: reject,
					//       });
					//     });

					//     // 上传成功后，可以在uploadRes中获取服务器返回的响应信息
					//     console.log('上传成功:', uploadRes.data);


					//     //调用uni.getFileInfo方法获取图片文件的信息
					//     const fileInfo = await new Promise((resolve, reject) => {
					//       uni.getFileInfo({
					//         filePath: imagePath, // 图片文件路径
					//         // success: resolve,
					// 		success: res => {
					// 			console.log("file exist!!!!",res.size)
					// 		},
					//         fail: reject,
					//       });
					//     });


					// 将图片文件转换为File对象
					// const file = new File([imagePath], fileInfo.fileName, {
					//   type: fileInfo.type, // 设置文件类型
					// });

					// var file1 = new File(imagePath);

					// console.log("图片的file文件：", file);



					// const formData = new FormData();
					// formData.append('smfile', file);

					// await Axios.post('https://sm.ms/api/v2/upload', formData, {
					// headers: {
					//     'Content-Type': 'multipart/form-data',
					//     'Authorization': 'giKtHZm1SL0Paj1l7Ye0lrQ0Ilq6VuxX',
					//   },
					// }).then((res) => {
					//   console.log("返回信息：" + res);
					// }).catch(err => {
					// 	console.log("图片转换接口请求失败");
					// });


				} catch (error) {
					console.error('上传失败:', error);
				}
			},

			// 将本地图片文件转换为Base64编码的字符串
		},
		mounted() {
			uni.getStorage({
				key: 'userId', // 储存在本地的变量名
				success: res => {
					// 成功后的回调
					this.id = res.data;
					console.log(this.id)
					uni.$u.http.get('v1/frontend/user/basicUserInfo?id=' + this.id, {

					}).then(res => {
						console.log("获取学币成功！");
						this.coin = res.data.data.user.Balance;
						console.log(this.coin);

					}).catch(err => {
						console.log("获取学币失败！！！");
					})
				}
			})
		}
	}
</script>

<style>
	.uniInput {
		height: 120rpx;
		font-size: 40rpx;
		border-width: 2px;
		border-bottom-style: solid;
		border-color: #f0f0f0;
		margin-left: 20rpx;
		margin-right: 20rpx;
	}

	.container {
		padding: 10px;
	}

	#editor {}

	.downArrow {
		margin-left: 5rpx;
		// padding-top: 10rpx;
		height: 20rpx;
		width: 20rpx;
	}

	.qlContainer {
		height: calc(100vh - 460rpx);
		/* height: 800rpx; */
		line-height: 160%;
		font-size: 34rpx;
		/* 	border-style:solid;
	border-width:2px;
	border-color:#f0f0f0; */

		/* 	overflow-y: auto; */
	}

	.operator {
		display: flex;
		margin-top: 10rpx;
	}

	.selectForm {
		display: flex;
		align-items: center;
		border-width: 2px;
		border-style: solid;
		border-color: #f0f0f0;
		height: 80rpx;
	}

	.selectFormItem {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-left: 10rpx;

	}

	.publishBtn {
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: #3a8afb;
		color: #ffffff;
		width: 320rpx;
		height: 100rpx;
		margin-right: 25rpx;
	}

	.uploadingBtn {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 320rpx;
		height: 100rpx;
		color: #3a8afb;
		border-style: solid;
		border-width: 1px;
		border-color: #3a8afb;
	}

	.downArrow {
		margin-left: 20rpx;
	}

	.inputNum {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-left: 100rpx;
	}

	.intext {
		text-align: center;
		width: 100rpx;
		background-color: #f1f1f1;
		border-radius: 10rpx;
		border: 1rpx solid #E0E3DA;
	}

	.below {}
</style>