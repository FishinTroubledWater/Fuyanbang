<template>
	<view class="content">
		<u-popup :show="showCon" mode="center" @close="close">
			<view class="contact">
				<text class="ttext">
					通过下方邮箱联系我们
				</text>
				<text class="ttext">
					1xxxxxxxxxxxx@qq.com
				</text>
			</view>
		</u-popup>
		<u-popup :show="showCopr" mode="center" @close="close">
			<view class="contact">
				<text class="ttext">
					合作请联系下方邮箱
				</text>
				<text class="ttext">
					1xxxxxxxxxxxx@qq.com
				</text>
			</view>
		</u-popup>
		<u-popup :show="showRep" mode="center" @close="close">
			<view class="contact">
				<text class="ttext">
					举报电话
				</text>
				<text class="ttext">
					16546516xxxxx
				</text>
			</view>
		</u-popup>
		<u-popup :show="showVer" mode="center" @close="close">
			<view class="contact">
				<text class="ttext">
					当前版本
				</text>
				<text class="ttext">
					{{version}}
				</text>
			</view>
		</u-popup>
		<view class="about_60">
			<view class="about_62">
				<image src="/static/my-assets//about/Logo.jpg" mode="scaleToFill" border="0" class="about_22"></image>
				<text decode="true" class="about_35">福研帮</text>
				<text decode="true" class="about_36">{{version}}</text>
			</view>
		</view>
		<view class="box">
			<u-cell-group>
				<u-cell isLink @click='showVersion()'>
					<text slot="title" class="title">检查更新</text>
				</u-cell>
				<u-cell isLink @click='showCooperate()'>
					<text slot="title" class="title">合作商洽</text>
				</u-cell>
				<u-cell isLink @click='showReport()'>
					<text slot="title" class="title">申诉举报</text>
				</u-cell>
				<u-cell isLink @click='showContact()'>
					<text slot="title" class="title">联系我们</text>
				</u-cell>
			</u-cell-group>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				version:'V1.0.0',
				showVer: false,
				showCopr: false,
				showCon: false,
				showRep: false,
			}
		},
		methods: {
			showContact() {
				this.showCon = true;
			},
			showCooperate() {
				this.showCopr = true;
			},
			showReport() {
				this.showRep = true;
			},
			showVersion() {
				this.showVer = true;
			},
			close() {
				this.showVer = false;
				this.showCon = false;
				this.showCopr = false;
				this.showRep = false;

			}
		},
	}
</script>

<style>
	.box {
		width: 600upx;
		height: 490upx;
		border-radius: 10px;
		background-color: #eeeeee;
		position: relative;
		top: 240px;
		margin: 0 auto;
	}

	.contact {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		border-radius: 10px;
	}

	.ttext {
		margin: 20px;
	}

	.title {
		margin: 10px;
	}

	.about_60 {
		white-space: normal;
		width: 749upx;
		position: relative;
		height: 100%;
		padding: 0upx;
		clear: both;
		float: left;
		background-color: #ffffff;
		text-align: left;
		border-radius: 0upx;
		font-size: 8upx;
	}

	.about_60 .about_62 {
		white-space: normal;
		width: 609upx;
		height: 300upx;
		padding: 0upx;
		margin-top: 127upx;
		margin-left: 84upx;
		float: left;
		position: absolute;
		top: 0upx;
		left: 0upx;
		text-align: left;
		border-radius: 0upx;
		font-size: 8upx;
	}

	.about_60 .about_62 .about_22 {
		white-space: normal;
		width: 180upx;
		height: 180upx;
		padding: 0upx;
		clear: both;
		margin-top: 10upx;
		margin-left: 200upx;
		float: left;
		text-align: left;
		border-radius: 0upx;
		font-size: 8upx;
		line-height: 127upx;
	}

	.about_60 .about_62 .about_35 {
		white-space: normal;
		width: 97upx;
		height: 36upx;
		padding: 0upx;
		clear: both;
		margin-top: 29upx;
		margin-left: 241upx;
		float: left;
		text-align: center;
		justify-content: center;
		border-radius: 0upx;
		color: #000000;
		font-size: 31upx;
		line-height: 36upx;
	}

	.about_60 .about_62 .about_36 {
		white-space: normal;
		width: 473upx;
		height: 26upx;
		padding: 0upx;
		clear: both;
		margin-top: 18upx;
		margin-left: 52upx;
		float: left;
		text-align: center;
		justify-content: center;
		border-radius: 0upx;
		color: #5c5f66;
		font-size: 20upx;
		line-height: 26upx;
	}
</style>