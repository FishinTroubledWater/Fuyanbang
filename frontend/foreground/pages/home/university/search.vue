<template>
	<view class="search">
		<view class="searchContainer">
			<uni-search-bar @input="input" :radius="20" placeholder="请输入搜索内容" bgColor="#F7F7F7" cancelButton="none" class="searchBox" v-model="searchContent"></uni-search-bar>
			<button class="searchButton" @click="goBack()">搜索</button>
		</view>
		<!-- 搜索列表 -->
		<view v-if="searchList.length!=0">
			<view class="searchList" v-for="(searchItem,searchIndex) in searchList" :key="searchIndex" >
				<uni-icons type="search" size="18" color="#C0C0C0"/>
				<view class="searchItem" >
					{{searchItem.word}}
				</view>
			</view>
		</view>
		<view v-else>
			<!-- 搜索历史 -->
<!-- 			<view class="history" v-model="his">
				<view class="history-title">
					<text>历史搜索</text>
					<uni-icons type="trash" size="20" color="#C0C0C0" @click="clearHistory"></uni-icons>
				</view>
				<view class="history-content" v-if="historyList.length!=0">
					<view class="history-item" v-for="(historyItem,historyIndex) in historyList" :key="historyIndex">
						{{historyItem}}
					</view>
				</view>
				<view class="history-none" v-else>
					<text>无搜索历史</text>
				</view>
			</view> -->

			<view class="found">
				<view class="foundTitle">
					<text>热门搜索</text>
					<uni-icons type="" size="20" color="#C0C0C0"></uni-icons>
				</view>
				<view class="foundContent">
					<view class="foundItem" v-for="(foundItem,foundIndex) in foundList" :key="foundIndex" @click="selectItem(foundItem)">
						{{foundItem}}
					</view>
				</view>
			</view>
		</view>
		
	</view>	
</template>
 
<script>
	export default {
		data() {
			return {
				// 初始化定时器为空
				time:null,
				// 用户输入的关键词
				keyword:'',
				//搜索数据的数组初始化
				searchList:[],
				//搜索历史初始化
				historyList:[],
				// 初始化搜索发现列表
				foundList:['福州大学','北京大学','清华大学','厦门大学','电子科技大学','浙江大学','复旦大学','南京大学'],
				searchContent: ''
			};
		},
		methods:{
			goBack(){
				// var data = this.searchContent;
				// uni.$emit('searchContent', data);
				// uni.navigateBack({
				// 	delta: 1
				// })
				// 在提交成功后，返回首页需要刷新（带上参数）
				let pages = getCurrentPages();
				let prevPage = pages[pages.length - 2];//上一个页面
				//直接调用上一个页面的setData()方法，把数据存到上一个页面中去
				// prevPage.setData({
				// 	mydata: { isRefresh: 1 }
				// })
				prevPage.$vm.searchContent = this.searchContent
				// 关闭当前页面，返回上一页面

				setTimeout(() => {
					uni.$emit('refreshData');
					uni.navigateBack({
						delta: 1
					})
				}, 250)

			},
			selectItem(index){
				this.searchContent = index;
			},
			
			//用户输入时可以获取用户输入的内容
			input:function(e){
				//每次使用先清空定时器，优化
				clearTimeout(this.timer);
				this.timer=setTimeout(()=>{
					this.keyword=e
					//获取搜索数据
					this.getSearchContent()
				},500)
				console.log(e)
			},
			//获取搜索列表的方法
			async getSearchContent(){
				// if(this.keyword.length==0){
				// 	this.searchList=[];
				// 	return
				// }else{
				// 	const res=await uni.$http.get('/search')
				// 	const {data,status}=res.data;
				// 	if(status!=200){
				// 		alert('获取数据失败')
				// 	}else{
				// 		this.searchList=data
				// 		//把搜索的关键字保存到historyList中
				// 		this.saveHistory()
				// 	}
				// }
			},
			// 保存历史记录
			saveHistory(){
				// if(this.historyList.indexOf(this.keyword)==-1){
				// 	this.historyList.unshift(this.keyword)
				// 	// 把用户输入的内容保存到历史记录当中
				// 	uni.setStorageSync('kw',JSON.stringify(this.historyList||'[]'))
				// }				
			},
			// 清空历史记录
			clearHistory(){
				// this.historyList=[]
				// uni.setStorageSync('kw','[]')
				// if(his.length==0){
				// 	this.his=!this.his
				// }
			}
		},
		onLoad() {
			// 从缓存中读取历史记录
			// this.historyList=JSON.parse(uni.getStorageSync('kw'))
		}
	}
</script>
 
<style lang="scss">
*{
	margin: 0;
	padding: 0;
}
.search{
	width: 100%;
	height: 100vh;
	background-color: #FFF;
	.searchContainer{
		width: 100%;
		display: flex;
		.searchBox{
			width: 550rpx;
		}
		.searchButton{
			display: flex;
			justify-content: center;
			align-items: center;
			width: 120rpx;
			height: 65rpx;
			margin-top: 20rpx;
			border-radius: 40rpx;
			background-color: #4ca9e7;
			color: #ffffff;
			font-size: 30rpx;
		}
	}
	.searchList{
		width: 100%;
		height: 80rpx;
		line-height: 80rpx;
		display: flex;
		border-bottom: 1px solid #eee;
		uni-icons{
			margin:0 20rpx;
		}
	}
	.history{
		.history-title{
			width: 90%;
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin: 0 auto;
			.text{
				font-weight: bold;
				font-size: 34rpx;
			}
		}
		.history-content{
			width: 90%;
			margin: 10rpx auto;
			display: flex;
			flex-wrap: wrap;
			.history-item{
				height: 50rpx;
				line-height: 50rpx;
				background-color: #F8F8F8;
				margin-top: 10rpx;
				margin-right: 20rpx;
				padding:0 20rpx;
				border-radius: 20rpx;
			}
		}
		.history-none{
			width: 100%;
			height: 100rpx;
			text-align: center;
			line-height: 100rpx;
		}
	}
	.found{
		margin-top: 50rpx;
		.foundTitle{
			width: 90%;
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin: 0 auto;
			.text{
				font-weight: bold;
				font-size: 34rpx;
			}
		}
		.foundContent{
			width: 90%;
			margin: 10rpx auto;
			display: flex;
			flex-wrap: wrap;
			.foundItem{
				display: flex;
				justify-content: center;
				align-items: center;
				height: 60rpx;
				line-height: 50rpx;
				background-color: #ffffff;
				margin-top: 20rpx;
				margin-right: 20rpx;
				padding:0 30rpx;
				border-radius: 20rpx;
				border-style:solid;
				border-width:1px;
				border-color: #CECECE;
				color: #CECECE;
				font-size: 26rpx;
			}
		}
	}
}
 
</style>