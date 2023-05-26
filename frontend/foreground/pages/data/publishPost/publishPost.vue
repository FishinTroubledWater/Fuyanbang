<template>
	<view>
		<view>
			<input class="uniInput" placeholder-style="color:#c7c7c7" focus placeholder="标题 (6-30字之间) " v-model="title" />
		</view>
		<view class="container">
			<editor id="editor" class="qlContainer" :placeholder="placeholder" @ready="onEditorReady" v-model="context"></editor>
		</view>
		<view>
			<button @click="chooseImage1()">上传图片</button>
			<button @click="test()">test</button>
		</view>
	</view>
	
</template>

<script>
	
	 export default {
	    data() {
	        return {
	            placeholder: '开始输入...',
				title: '',
				context: '',
	        }
	    },
	    methods: {
	        onEditorReady() {
	            uni.createSelectorQuery().select('#editor').context((res) => {
	                this.editorCtx = res.context
	            }).exec()
	        },
	        undo() {
	            this.editorCtx.undo()
	        },
			
			test(){
				console.log("bbbbbb" + this.title);
				console.log("aaaaaa" + this.context);
			},
			
			chooseImage1(){
				let _this = this;
				uni.chooseImage({
				    count: 6, //默认9
				    sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
				    sourceType: ['album'], //从相册选择
				    success: function (res) {
				        console.log(JSON.stringify(res.tempFilePaths));
						var data = JSON.stringify(res.tempFilePaths)
						_this.editorCtx.insertImage({
							width: '100%', //设置宽度为100%防止宽度溢出手机屏幕
							height: 'auto',
							src: 'https://pic35.photophoto.cn/20150511/0034034892281415_b.jpg', //服务端返回的url
							alt: '图像',
						})
				    }
				});
			},
			
			async chooseImage(type) {
				let _this = this;
				uni.chooseImage({
						count: 1, //最多可以选择的图片张数，默认9
						sizeType: ['original', 'compressed'], //original 原图，compressed 压缩图，默认二者都有
						sourceType: ['album', 'camera'], //album 从相册选图，camera 使用相机，默认二者都有
						success: (respone) => {
							_this.$request.urlRequest(
								'/gate/oss/token', {},
								'GET',
								(res) => {
									if (res.code == 200) {
									let data = res.result;
									let env = {
										uploadImageUrl: 'https://58d.oss-cn-hangzhou.aliyuncs.com/', // 默认存在根目录，可根据需求改
										AccessKeySecret: data.AccessKeySecret, // AccessKeySecret 去你的阿里云上控制台上找
										OSSAccessKeyId: data.AccessKeyId, // AccessKeyId 去你的阿里云上控制台上找
										stsToken: data.SecurityToken,
										timeout: 87600 //这个是上传文件时Policy的失效时间
									}
									let dir = 'images/';
									let filePath = respone.tempFilePaths[0];
									const aliyunFileKey = dir + new Date().getTime() + Math.floor(Math.random() * 150) + '.png';
									const aliyunServerURL = env.uploadImageUrl; //OSS地址，需要https
									const accessid = env.OSSAccessKeyId;
									const policyBase64 = _this.getPolicyBase64(env);
									const signature = _this.getSignature(policyBase64, env); //获取签名
									const stsToken = env.stsToken;
					
									let param = {
										'key': aliyunFileKey,
										'policy': policyBase64,
										'OSSAccessKeyId': accessid,
										'signature': signature,
										'success_action_status': '200',
										'x-oss-security-token': stsToken,
										'stsToken': stsToken,
									};
					
									uni.uploadFile({
										url: "https://58d.oss-cn-hangzhou.aliyuncs.com/", //开发者服务器 url
										filePath: filePath, //要上传文件资源的路径
										name: 'file', //必须填file
										formData: param,
										success: (res) => {
											if(type){
												_this.detail.imageUrl = aliyunServerURL + aliyunFileKey;
												return;
											}
											_this.editorCtx.insertImage({
												src: aliyunServerURL + aliyunFileKey, // 此处需要将图片地址切换成服务器返回的真实图片地址
												alt: '图片',
												width:"320upx",// 此处可以对图片进行尺寸的限制，否则页面中的图片会展示不全
												mode:'widthFix',
												success: function(e) {}
											});
										},
										fail: (err) => {
											// _this.msg.push(JSON.stringify(err));
											// err.wxaddinfo = aliyunServerURL;
											// failc(err);
										},
									})
								}
							}
						)
					}
				});
			}
		}
	}
</script>

<style>
.uniInput{
	height: 120rpx;
	font-size: 40rpx;
	border-width:2px;
	border-bottom-style:solid;
	border-color:#f0f0f0;
	margin-left: 20rpx;
	margin-right: 20rpx;
}
.container{
	padding: 10px;
}
#editor {
	
}
.qlContainer{
	height: calc(100vh - 380rpx);
	line-height: 160%;
	font-size: 34rpx;
	overflow-y: auto;
}
</style>
