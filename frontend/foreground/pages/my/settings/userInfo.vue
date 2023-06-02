<template>
	<view class="backGround">
		<view class="infoCard">
			<view class="content">
				<u-cell-group>
					<u-cell>
						<text slot="title">修改头像</text>
						<u-image @click="changeHead()" width='100rpx' height='100rpx' slot="value" :src="user.avatarUrl"
							shape="circle"></u-image>
					</u-cell>
					<u-cell>
						<text slot="title">修改昵称</text>
						<input class="right" slot="value" v-model="user.nickName"
							placeholder="请输入昵称">{{user.nickName}}</input>
					</u-cell>

					<picker :range="sex" @confirm="bindSexChange($event)" @change="bindSexChange($event)">
						<u-cell>
							<text slot="title">设置性别</text>
							<text slot="value">{{user.sex}}</text>
						</u-cell>
					</picker>


					<u-cell>
						<text slot="title">所在地区</text>
						<input class="right" slot="value" v-model="user.area" placeholder="请输入地区">{{user.area}}</input>
					</u-cell>


					<picker :range="colleges" @confirm="bindCollegeChange($event)" @change="bindCollegeChange($event)">
						<u-cell>
							<text slot="title">本科院校</text>
							<text slot="value">{{user.college}}</text>
						</u-cell>

					</picker>

					<picker :range="majors" @confirm="bindMajorChange($event)" @change="bindMajorChange($event)">
						<u-cell>
							<text slot="title">本科专业</text>
							<text slot="value">{{user.major}}</text>
						</u-cell>
					</picker>

					<u-cell>
						<text slot="title">考研年份</text>
						<input class="right" slot="value" v-model="user.year"
							placeholder="请输入考研年份">{{user.year}}</input>
					</u-cell>

					<picker :range="colleges" @confirm="bindTargetCollegeChange($event)"
						@change="bindTargetCollegeChange($event)">
						<u-cell>
							<text slot="title">报考院校</text>
							<text slot="value">{{user.targetCollege}}</text>
						</u-cell>
					</picker>

					<view class="box">
						<text class="slogan">个性签名</text>
						<u--textarea v-model="user.slogan" placeholder="请在此处编辑您的个性签名"
							count>{{user.slogan}}</u--textarea>
					</view>
					<button class="upButton" @click="upInfo()">保存</button>
				</u-cell-group>
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
				id: "",
				user: {
					avatarUrl: '/static/my-assets/泰裤辣.png',
					nickName: "",
					sex: "不明",
					area: "",
					college: "",
					major: "",
					year: "",
					targetCollege: "",
					slogan: ""
				},

				//选择器数据
				sex: ['男', '女'],
				colleges: ['福州大学', '清华大学', '贵州大学', '上海大学', '北京大学', '北京大学', '清华大学', '清华大学', '清华大学', '清华大学', '清华大学', '清华大学',
					'南京大学',
				],
				majors: ['金融', '应用统计', '税务', '国际商务', '保险', '资产评估', '审计', '法律', '社会工作', '警务', '教育', '体育', '应用心理', ],
			}
		},
		onShow() {
			console.log("执行onShow()函数")
			this.refresh();
		},
		onPullDownRefresh() {
			setTimeout(() => {
				this.refresh();
				uni.stopPullDownRefresh();
			}, 1000)
		},
		mounted() {
			console.log("执行onLoad（）1");
			this.refresh();
			console.log("执行onLoad（）end")
		},
		methods: {
			refresh() {
				uni.getStorage({
					key: 'userId', // 储存在本地的变量名
					success: res => {
						// 成功后的回调
						this.id = res.data;
					}
				})
				uni.$u.http.get('v1/frontend/user/basicUserInfo?id=' + this.id, {}).then(res => {
					console.log("获取数据成功！");
					console.log(res.data.data);
					this.user.avatarUrl = res.data.data.user.AvatarUrl;
					this.user.nickName = res.data.data.user.NickName;
					this.user.sex = res.data.data.user.Sex;
					this.user.area = res.data.data.user.Area;
					this.user.slogan = res.data.data.user.Slogan;
					this.user.useageDays = res.data.data.userDay;
					this.user.college = res.data.data.user.College;
					this.user.major = res.data.data.user.Major;
					this.user.year = res.data.data.user.Year;
					this.user.targetCollege = res.data.data.user.TargetCollege;
				}).catch(err => {
					console.log(this.id);
					console.log("获取数据失败！");
				})

			},
			async changeHead() {
				var _this = this;

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
						var s = base64.substr(base64.indexOf(',') + 1, base64.length);
						console.log(s)
						uni.$u.http.post('https://api.superbed.cn/upload', {
							token: '1000766339bd4c248a8ad625a87f687d',
							b64_data: s,
						}).then(res => {
							console.log("图片上传成功");
							console.log(res.data.url);
							_this.user.avatarUrl = res.data.url
						}).catch(error => {
							console.log(error);
							console.log("失败")
						})
					}).catch(error => {
						console.error(error)
					})


			},
			bindSexChange(e) {
				this.user.sex = this.sex[e.target.value]
			},
			bindCollegeChange(e) {
				this.user.college = this.colleges[e.target.value]
			},
			bindMajorChange(e) {
				this.user.major = this.majors[e.target.value]
			},
			bindTargetCollegeChange(e) {
				this.user.targetCollege = this.colleges[e.target.value]
			},
			//上传用户信息的方法
			upInfo() {
				uni.$u.http.post('/v1/frontend/user/settings', {
					id: this.id,
					avatarUrl: this.user.avatarUrl,
					nickName: this.user.nickName,
					sex: this.user.sex,
					area: this.user.area,
					college: this.user.college,
					major: this.user.major,
					year: this.user.year,
					targetCollege: this.user.targetCollege,
					slogan: this.user.slogan
				}).then(res => {
					console.log(res);
					uni.showToast({
						title: '修改成功',
						//将值设置为 success 或者直接不用写icon这个参数
						icon: 'success',
						//显示持续时间为 1.5秒
						duration: 1500
					})
					this.timer = setInterval(() => {
						//TODO 
						uni.navigateBack({
							delta: 1, //返回层数，2则上上页
						})
					}, 1500);
				}).catch(err => {
					console.log("失败了。。。");
					console.log(err);
				})
			},
			onUnload: function() {
				if (this.timer) { //在页面卸载时清除定时器有时会清除不了，可在页面跳转时清除
					clearInterval(this.timer);
					this.timer = null;
				}
			}

		}
	}
</script>

<style>
	.backGround{
		width: 100%;
		height: 1600rpx;
		background-color:#fafafa
	}
	.right {
		text-align: right;
	}

	.box {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		padding-top: 8px;
		padding-left: 20px;
		padding-right: 20px;

	}

	.slogan {
		width: 25%;
		height: 30px;
		margin-left: 30rpx;
		font-size: 15px;
	}

	.upButton {
		width: 80%;
		height: 50px;
		background-image: linear-gradient(112deg, #08507880, #85d8ce);
		border-radius: 20px;
		margin-top: 100rpx;
		/* margin-bottom: 50rpx; */
		color: #fff;
		border: none;
		box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
	}

	.u-cell {
		padding-left: 20px;
		padding-right: 20px;
	}

	.infoCard {
		width: auto;
		height: 1000rpx;
		margin: 25rpx;
		margin-top: 60rpx;
		margin-bottom: 60rpx;
		background-color: #fff;
		box-shadow: 12rpx 12rpx 10rpx #bfbfbf;
		border-radius: 20rpx;
	}
</style>