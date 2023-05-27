<template>
	<view>
		<view>
			<input class="uniInput" placeholder-style="color:#c7c7c7" focus placeholder="标题 (6-30字之间) " v-model="title" />
		</view>
		<view class="container">
			<editor id="editor" class="qlContainer" :placeholder="placeholder" @input="getText" @ready="onEditorReady"></editor>
		</view>
		<view class="selectForm">
			<picker @change="bindPickerChange1" :range="array1" :value="index1" class="selectFormItem">
				<label class="">{{array1[index1]}}</label>
				<label class="downArrow"> ∨ </label>
			</picker>
		</view>
		<view class="operator">
			<button @click="chooseImage1()" class="uploadingBtn">上传图片</button>
			<button @click="test()" class="publishBtn">发布</button>
		</view>
	</view>
	
</template>

<script>
	import Axios from 'axios'
	Axios.defaults.baseURL = '/'
	// eslint-disable-next-line no-unused-vars
	const axios = require('axios')
	
	
	 export default {
	    data() {
	        return {
	            placeholder: '开始输入文章内容...',
				title: '',
				context: '',
				synopsis: '',
				array1: ['请选择文章类型','加油站','求解答','学长学姐说'],
				index1: 0,
				type: '',
	        }
	    },
	    methods: {
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg = this.array1[this.index1];
				this.type = this.array1[this.index1];
			},
	        onEditorReady() {
	            uni.createSelectorQuery().select('#editor').context((res) => {
	                this.editorCtx = res.context
	            }).exec()
	        },
	        undo() {
	            this.editorCtx.undo()
	        },
			getText(e){
			    // console.log(e.detail.html);//带标签内容
				this.context = e.detail.html;
				this.synopsis = e.detail.text;
				console.log(this.context);//带标签内容
				console.log(this.synopsis);
			},
			test(){
				if(this.index1 != 0 || this.title != "") {
					console.log("文章类型：" + this.type);
					console.log("标题：" + this.title);
					console.log("内容：" + this.context);
				}
				else if(this.title == ""){
					console.log("请先输入文章标题");
					uni.showToast({
						title: '请先输入文章标题',
						icon: 'none',    //如果要纯文本，不要icon，将值设为'none'
						duration: 2000    //持续时间为 2秒
					})  
				}
				else{
					console.log("请先选择文章类型");
					uni.showToast({
						title: '请先选择文章类型',
						icon: 'none',    //如果要纯文本，不要icon，将值设为'none'
						duration: 2000    //持续时间为 2秒
					})  

				}
				
			},
			async chooseImage1(){
				let _this = this;
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
				
					_this.editorCtx.insertImage({
						width: '100%', //设置宽度为100%防止宽度溢出手机屏幕
						height: 'auto',
						src: res.tempFilePaths[0],
						alt: '图像',
					})
					
					// const uploadRes = await new Promise((resolve, reject) => {
					//       uni.uploadFile({
					//         url: 'https://sm.ms/api/v2/upload', // 图床服务器的URL地址
					//         filePath: imagePath, // 图片文件路径
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
					
				    //调用uni.getFileInfo方法获取图片文件的信息
				    const fileInfo = await new Promise((resolve, reject) => {
				      uni.getFileInfo({
				        filePath: imagePath, // 图片文件路径
				        success: resolve,
				        fail: reject,
				      });
				    });
				
				    // 将图片文件转换为File对象
				    const file = new File([imagePath], fileInfo.fileName, {
				      type: fileInfo.type, // 设置文件类型
				    });
				
				//     console.log("图片的file文件：", file);
				//   } catch (error) {
				//     console.error("选择图片出错：", error);
				//   }
				  
					
				const formData = new FormData();
				formData.append('smfile', file);
						
				axios.post('https://sm.ms/api/v2/upload', formData, {
				headers: {
				    'Content-Type': 'multipart/form-data',
				    'Authorization': 'giKtHZm1SL0Paj1l7Ye0lrQ0Ilq6VuxX',
				  },
				}).then((res) => {
				  console.log("返回信息：" + res);
				}).catch(err => {
					console.log("图片转换接口请求失败");
				});
				
				
			}catch (error) {
				console.error('上传失败:', error);
			}
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
	height: calc(100vh - 460rpx);
	line-height: 160%;
	font-size: 34rpx;
	overflow-y: auto;
}
.operator{
	display: flex;
	margin-top: 10rpx;
}
.selectForm{
	display: flex;
	align-items: center;
	border-width:2px;
	border-style:solid;
	border-color:#f0f0f0;
	height: 80rpx;
}
.selectFormItem{
	display: flex;
	justify-content: center;
	align-items: center;
	margin-left: 10rpx;

}
.publishBtn{
	display: flex;
	justify-content: center;
	align-items: center;
	background-color: #3a8afb;
	color: #ffffff;
	width: 320rpx;
	height: 100rpx;
	margin-right: 25rpx;
}
.uploadingBtn{
	display: flex;
	justify-content: center;
	align-items: center;
	width: 320rpx;
	height: 100rpx;
	color: #3a8afb;
	border-style:solid;
	border-width:1px;
	border-color:#3a8afb;
}
.downArrow{
	margin-left: 20rpx;
}
</style>
