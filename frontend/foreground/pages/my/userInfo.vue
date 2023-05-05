<template>
	<view class="content">
		<u-cell-group>
			<u-cell>
				<text slot="title">修改头像</text>
				<u-image @click="changeHead()" width='100rpx' height='100rpx' slot="value" :src="avatarUrl"
					shape="circle"></u-image>
			</u-cell>
			<u-cell>
				<text slot="title">修改昵称</text>
				<input class="right" slot="value" v-model="user.nickName" placeholder="请输入内容"></input>
				<text>{{user.nickName}}</text>
			</u-cell>


			<picker :range="sex" @confirm="bindSexChange($event)" @change="bindSexChange($event)">
				<u-cell>
					<text slot="title">设置性别</text>
					<text slot="value">{{user.sex}}</text>
				</u-cell>
			</picker>

			<picker @confirm="bindAreaConfirm($event)" @change="bindAreaChange($event)" @columnchange="bindAreaChange1($event)" mode='multiSelector' :range="area" range-key="text">
				<u-cell>
					<text slot="title">所在地区</text>
					<text slot="value">{{user.area}}</text>
				</u-cell>
			</picker>
		</u-cell-group>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				account: "",
				avatarUrl: '/static/my-assets/泰裤辣.png',
				user: {
					nickName: "",
					sex: "性别",
					area: "",
					college: "",
					major: "",
					year: "",
					targetCollege: "",
					slogan: ""
				},
				sex: ['男', '女'],
				area: [
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
				this.user.area = this.pos1.text+this.pos2.text+this.pos3.text;
				console.log(this.pos1.text+this.pos2.text+this.pos3.text)
			}
		}
	}
</script>

<style>
	.right {
		text-align: right;
	}
</style>