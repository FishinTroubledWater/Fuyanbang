<template>
	<view class="content">
		<u-popup :show="showGet" mode="center" @close="close">
			<view class="pop">
				<text class="title">当前获取方式</text>
				<hr class='divider' />
				<text class="ans">1.每日签到可获得1学币</text>
			</view>
		</u-popup>
		<view class="box1">
			<text class="coin">
				{{user.balance}}
			</text>
			<text class="balen">
				学币余额
			</text>
			<button class="getBtn" @click="show1 = true">充值</button>
			<view>
				<button class="getBtn" @click="show2 = true">兑换</button>
			</view>
		</view>

		<hr class='divider' />
		<view class="box2">
			<text class="title">什么是学币</text>
			<text class="ans">学币是仅限于在福研帮中使用的虚拟货币</text>
			<text class="title" id="t1">如何使用学币</text>
			<text class="ans">1.学币是仅限于在福研帮中使用的虚拟货币</text>
			<text class="ans">2.学币是仅限于在福研帮中使用的虚拟货币</text>
		</view>
		<view>
			<u-popup :show="show1" @close="close2" @open="open1">


				<view class="recharge-buttons">
					<button class="recharge-button" @click="recharge(10)">10学币（{{ calculateAmount(10) }}元）</button>
					<button class="recharge-button" @click="recharge(20)">20学币（{{ calculateAmount(20) }}元）</button>
					<button class="recharge-button" @click="recharge(50)">50学币（{{ calculateAmount(50) }}元）</button>
					<!-- 可按需添加其他充值金额的按钮 -->

				</view>
				<view style="text-align: center; margin-bottom: 50rpx;">
					<text style="text-align: center; margin-bottom: 20rpx;">1块钱=10学币哦！</text>
				</view>
			</u-popup>
			<u-popup :show="show2" @close="close3" @open="open2">
				<view class="exchange-content">
					<view class="exchange-options">
						<radio-group v-model="selectedOption">
							<label class="option-label">
								<radio class="option-radio" value="option1"></radio>
								<text class="option-text">1个月bilibili大会员(200学币)</text>
							</label>
							<label class="option-label">
								<radio class="option-radio" value="option2"></radio>
								<text class="option-text">1个月网易云音乐会员(149学币)</text>
							</label>
							<label class="option-label">
								<radio class="option-radio" value="option3"></radio>
								<text class="option-text">1个月百度网盘会员(149学币)</text>
							</label>
						</radio-group>
					</view>

					<view class="input-container">
						<input class="input-field" v-model="inputValue" placeholder="请输入充值账号/手机号"></input>
					</view>

					<button class="exchange-button" @click="submitExchange">提交！！</button>
				</view>
			</u-popup>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				show1: false,
				show2: false,
				showGet: false,
				selectedOption: '',
				inputValue: '',
				id: '',
				user: {
					balance: '100',
				},
				selectedAmount: 0, // 新增的属性，保存当前选择的充值金额
			}
		},

		onShow() {
			this.refresh();
		},
		mounted() {
			this.refresh();
		},
		methods: {
			submitExchange() {
				// 提交兑换操作，可以在这里处理相关的逻辑
				console.log('选项:', this.selectedOption);
				console.log('输入值:', this.inputValue);
			},
			calculateAmount(coin) {
				return coin / 10; // 假设每十学币对应1人民币
			},
			open1() {
				// console.log('open');
			},
			open2() {
				// console.log('open');
			},
			close2() {
				// console.log("搞到这了吗？");
				this.show1 = false
				// console.log('close');
			},
			close3() {
				// console.log("搞到这了吗？");
				this.show2 = false
				// console.log('close');
			},
			recharge(amount) {
				const actualAmount = this.calculateAmount(amount);
				console.log(`充值学币: ${amount}（${actualAmount}元）`);
				this.selectedAmount = amount;
				// 其他充值操作的逻辑代码...

				// 关闭充值弹窗
				this.showGet = false;
			},
			refresh() {
				uni.getStorage({
					key: 'userId', // 储存在本地的变量名
					success: res => {
						// 成功后的回调
						// console.log(res.data);   // hello  这里可做赋值的操作
						this.id = res.data;
						console.log(this.id)
					}
				})
				uni.$u.http.get('v1/frontend/user/basicUserInfo?id=' + this.id, {

				}).then(res => {
					console.log("获取学币成功！");
					console.log(res.data.data);
					this.user.balance = res.data.data.user.Balance;

				}).catch(err => {
					console.log("获取学币失败！！！");
				})
			},
			getCoin() {
				this.showGet = true;
			},
			close() {
				this.showGet = false;
			},

		}
	}
</script>

<style>
	.exchange-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		/* 居中对齐 */
		height: 100%;
	}

	.exchange-title {
		font-size: 20px;
		margin-bottom: 20px;
	}

	.exchange-options {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		margin-top: 100rpx;
		margin-bottom: 20px;
	}

	.option-label {
		display: flex;
		align-items: center;
		margin-bottom: 10px;
	}

	.option-radio {
		margin-right: 10px;
	}

	.option-text {
		font-size: 16px;
	}

	.input-container {
		margin-bottom: 20px;
	}

	.input-field {
		width: 200px;
		height: 30px;
		padding: 5px;
		border: 1px solid #ccc;
		border-radius: 5px;
	}

	.submit-button {
		padding: 10px 20px;
		background-color: #4caf50;
		color: #fff;
		border: none;
		border-radius: 5px;
		cursor: pointer;
	}

	.submit-button:hover {
		background-color: #45a049;
	}

	.recharge-buttons {
		height: 250rpx;
		display: flex;
		justify-content: space-around;
		margin-top: 20px;
	}

	.recharge-button {
		width: 200rpx;
		height: 100rpx;
		border-radius: 20px;
		background-color: #f4ce69;
		color: #ffffff;
	}
	
	.exchange-button {
		width: 250rpx;
		height: 100rpx;
		border-radius: 20px;
		background-color: #f4ce69;
		color: #ffffff;
	}

	.content {

		display: flex;
		flex-direction: column;
	}

	.box1 {
		height: 600rpx;
		display: flex;
		flex-direction: column;
		text-align: center;

	}

	.box2 {
		height: 500px;
		display: flex;
		flex-direction: column;
	}

	.pop {
		width: 300px;
		height: 150px;
		display: flex;
		flex-direction: column;
	}

	.title {
		margin-top: 10%;
		margin-left: 8%;
		font-size: 20px;
		font-weight: bold;
	}

	.ans {
		margin-top: 10px;
		margin-left: 8%;
	}

	.coin {
		font-size: 100upx;
		margin-top: 50px;
		color: #f9bd10;
	}

	.balen {
		font-size: 30upx;
		margin-top: 10px;
		color: #d1c2d3;
	}

	.getBtn {
		margin-top: 20px;
		width: 40%;
		border-radius: 50px;
		background-color: #f4ce69;
	}

	.divider {
		margin: 12px;
		align-self: center;
		width: 90%;
		color: #987cb9;
		size: 10;
	}
</style>