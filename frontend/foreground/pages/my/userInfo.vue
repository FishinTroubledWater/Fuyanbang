<template>
	<!-- TODO: 所在地区部分的多列选择器存在一定问题 -->
	<view class="content">
		<u-cell-group>
			<u-cell>
				<text slot="title">修改头像</text>
				<u-image @click="changeHead()" width='100rpx' height='100rpx' slot="value" :src="user.avatarUrl"
					shape="circle"></u-image>
			</u-cell>
			<u-cell>
				<text slot="title">修改昵称</text>
				<input class="right" slot="value" v-model="user.nickName" placeholder="请输入内容">{{user.nickName}}</input>
			</u-cell>


			<picker :range="sex" @confirm="bindSexChange($event)" @change="bindSexChange($event)">
				<u-cell>
					<text slot="title">设置性别</text>
					<text slot="value">{{user.sex}}</text>
				</u-cell>
			</picker>

			<picker @confirm="bindAreaConfirm($event)" @change="bindAreaChange($event)"
				@columnchange="bindAreaChange1($event)" mode='multiSelector' :range="areas" range-key="text">
				<u-cell>
					<text slot="title">所在地区</text>
					<text slot="value">{{user.area}}</text>
				</u-cell>
			</picker>

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
				<input class="right" slot="value" v-model="user.year" placeholder="请输入考研年份">{{user.year}}</input>
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
				<u--textarea v-model="user.slogan" placeholder="请在此处编辑您的个性签名" count>{{user.slogan}}</u--textarea>
			</view>
			<button class="upButton" @click="upInfo()">保存</button>
		</u-cell-group>
	</view>
</template>

<script>
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
				areas: [
					[{
						text: '贵州省'
					}, {
						text: '福建省'
					}],

					[{
						key: 1,
						text: '贵阳市'
					}, {
						key: 2,
						text: '清镇市'
					}, {
						key: 3,
						text: '遵义市'
					}],

					[{
						text: '云岩区'
					}, {
						text: '南明区'
					}, {
						text: '花溪区'
					}, {
						text: '金阳新区'
					}]
				],
				pos1: {},
				pos2: {},
				pos3: {}
			}
		},
		methods: {
			changeHead() {
				uni.chooseImage({
					count: 1,
					success: (res) => {
						this.headImg = res.tempFilePaths[0]
					}
				});
			},
			bindSexChange(e) {
				this.user.sex = this.sex[e.target.value]
			},
			bindAreaChange1(e) {
				let array = [
					[{
						text: '云岩区'
					}, {
						text: '南明区'
					}, {
						text: '花溪区'
					}, {
						text: '金阳新区'
					}],


					[{
						text: '清镇一区'
					}, {
						text: '清镇二区'
					}, {
						text: '清镇三区'
					}, {
						text: '清镇四区'
					}],


					[{
						text: '遵义一区'
					}, {
						text: '遵义二区'
					}, {
						text: '遵义三区'
					}, {
						text: '遵义四区'
					}]
				];
				if (e.detail.column == 1) {
					this.$set(this.area, 2, array[e.target.value]);
				} else {
					return;
				}
			},
			bindAreaChange(e) {
				let v = e.target.value;
				let index1 = v[0];
				let index2 = v[1];
				let index3 = v[2];
				this.pos1 = this.area[0][index1];
				this.pos2 = this.area[1][index2];
				this.pos3 = this.area[2][index3];
			},
			bindAreaConfirm(e) {
				let v = e.target.value;
				let index1 = v[0];
				let index2 = v[1];
				let index3 = v[2];
				this.pos1 = this.area[0][index1];
				this.pos2 = this.area[1][index2];
				this.pos3 = this.area[2][index3];
				this.user.area = this.pos1.text + this.pos2.text + this.pos3.text;
				console.log(this.pos1.text + this.pos2.text + this.pos3.text)
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
				uni.showToast({
					title: '修改成功',
					//将值设置为 success 或者直接不用写icon这个参数
					icon: 'success',
					//显示持续时间为 2秒
					duration: 1500
				}) 
				// 点击上传信息按钮触发的方法
				// var that = this
				// uni.request({
				// 	//api地址
				// 	url: 'http://localhost:3000/web/api/rest/user/',
				// 	header: {
				// 		'content-type': 'application/x-www-form-urlencoded'
				// 	},
				// 	method: 'POST',
				// 	data: {
				// 		// 将json数据转化成字符串格式进行上传
				// 		information: JSON.stringify(that.user)
				// 	},
				// 	success: (res) => {
				// 		console.log(res)
				// 	},
				// 	error(err) {
				// 		console.log(err)
				// 	}
				// })
				this.timer = setInterval(() => {
				    //TODO 
					uni.navigateBack({
							delta:1,//返回层数，2则上上页
						})
				}, 1500);
			},
			onUnload:function(){  
			    if(this.timer) {  //在页面卸载时清除定时器有时会清除不了，可在页面跳转时清除
			        clearInterval(this.timer);  
			        this.timer = null;  
			    }  
			}

		}
	}
</script>

<style>
	.right {
		text-align: right;
	}

	.box {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		padding-top: 8px;
		padding-left: 15px;
		padding-right: 15px;

	}

	.slogan {
		width: 20%;
		height: 30px;
		font-size: 15px;
	}

	.upButton {
		width: 80%;
		height: 50px;
		background-color: bisque;
		border-radius: 20px;
		margin-top: 50px;
	}
</style>