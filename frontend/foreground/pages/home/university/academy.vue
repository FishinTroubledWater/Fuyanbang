<template>
	<view class="container">
		<view class="search" v-show="isShow">
			<view class="searchContainer">
				<uni-search-bar @input="input" :radius="20" placeholder="请输入搜索内容" bgColor="#F7F7F7" cancelButton="none" class="searchBox" v-model="searchContent"></uni-search-bar>
				<button class="searchButton" @click="search()">搜索</button>
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
		
		<view v-show="!isShow">
			<view class="selectForm">
				<picker @change="bindPickerChange1" :range="array1" :value="index1" class="selectFormItem">
					<label class="wordSpace">{{array1[index1]}}</label>
					<label class="down">
						<image class="downArrow" src="@/static/academy-icons/down.png"></image>
					</label>
				</picker>
				<picker @change="bindPickerChange2" :range="array2" :value="index2" class="selectFormItem">
					<label class="wordSpace">{{array2[index2]}}</label>
					<!-- <label class="downArrow">∨</label> -->
					<label class="down">
						<image class="downArrow" src="@/static/academy-icons/down.png"></image>
					</label>
				</picker>
				<picker @change="bindPickerChange3" :range="array3" :value="index3" class="selectFormItem">
					<label class="wordSpace">{{array3[index3]}}</label>
					<!-- <label class="downArrow">∨</label> -->
					<label class="down">
						<image class="downArrow" src="@/static/academy-icons/down.png"></image>
					</label>
				</picker>
			</view>
			<view class="searchAcademy">
				<text class="searchText" v-if="isExist == true">
					共搜索到 <text class="searchNum">{{mes.num}}</text> 所院校
				</text>
				<text class="searchText" v-if="isExist == false">
					共搜索到 <text class="searchNum">0</text> 所院校
				</text>
			</view>
			<view>
				<view  class="viewAcademy" v-for="m in mes.list" @click="goUniverity(m.Code)" :style="active">
					<image class="academyLogo" :src="m.Logo"></image>
					<view class="viewText">
						<text class="academyName">{{m.Name}}</text>
						<view class="academyType">
							<view class="typeOfScoreLine">{{m.LineType}}</view>
							<text class="typeOfInstitution">{{m.Level}}</text>
						</view>
					</view>
					<text class="lacation">{{m.Region}}</text>
				</view>
			</view>
		</view>
		
	</view>
</template>

<script>
import Axios from 'axios'
Axios.defaults.baseURL = '/'
// eslint-disable-next-line no-unused-vars
const axios = require('axios')

import { onLoad } from 'uview-ui/libs/mixin/mixin';
	export default {
		data() {
			return {
				isShow: false,
				
				array1: ['院校地区','北京','福建','天津','上海','重庆',
				'内蒙古','广西','西藏','宁夏','新疆','山西','辽宁','吉林',
				'黑龙江','江苏','浙江','安徽','江西','山东','河北','河南','湖北',
				'湖南','广东','海南','四川','贵州','云南','陕西','甘肃','青海','台湾'],		
				array2: ['院校层次','985','211','本科','大专','高职'],
				array3: ['院校类型','综合','理工','师范','农林','政法','医药','财经','民族','语言','艺术','体育','军事','旅游'],
				index1: 0,
				index2: 0,
				index3: 0,
				searchnum: 111,
				active:'',
				
				mes: [],
				sendMes: [],
				saidMes: [],
				isExist: false,
				
				region: '院校地区',
				level: '院校层次',
				type: '院校类型',
				academyName: '福州大学',
				
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
				searchContent: '',
			};
		},
		onNavigationBarButtonTap:function(e){
			this.isShow = true;
			// uni.navigateTo({
			// 	url: "/pages/home/university/search"
			// })
		},
		async onLoad() {
			const oMeta = document.createElement('meta');
			oMeta.name = "referrer";
			oMeta.content = "no-referrer"
			document.getElementsByTagName('head')[0].appendChild(oMeta);
			// const result = uni.$u.http.post('/PairProject/players?gender=F&seed=true').then(res => {
			// 	console.log("aaa");
			// }).catch(err => {
			// 	console.log("失败");
			// 	console.log(err);
			// })
			
			// const result = uni.$u.http.post('/PairProject/players?gender=F&seed=true').then(res => {
			// 	console.log("aaa");
			// }).catch(err => {

			// })
			
			// let mes = await this.$u.api.getInfo();
			// this.mes = postMenu({ custom: { auth: true }}).then(() => {
				
			// }).catch(() =>{
				
			// })
			// console.log(this.mes)
			
			var _this= this;
			// 详见官网：https://uniapp.dcloud.io/api/request/request
			// uni.request({
			// 	url:'/PairProject/players?gender=F&seed=true',
			// 	method: 'POST',
			// 	data: {
					
			// 	},
			// 	success: res => {
			// 		_this.mes = res.data;
			// 		console.log(_this.mes)					
			// 	},						
			// });
			
			// uni.request({
			// 	// url:'http://124.222.141.238:8088/v1/frontend/academy/searchByRule',
			// 	url:'/v1/frontend/user/passwordLogin',
			// 	method: 'GET',
			// 	data: {
			// 		region: "福州",
			// 		level: '985',
			// 		type: '法学',
			// 	},
			// 	header: {
			// 	 //    'Connection': 'close',
			// 		// 'Content-Length': '793',
			// 		// 'Content-Type': 'application/json;charset=utf-8',
			// 		// 'Data': '',
			// 	},
			// 	success: res => {
			// 		_this.mes = res.data;
			// 		console.log(_this.mes)					
			// 	},						
			// });
			
			// const result = await Axios.post('http://124.222.141.238:8088/v1/frontend/user/passwordLogin', {
			//         account: 'test123',
			//         password: 'test123'
			//         }).then(res =>{
			//           console.log("成功");
			// 		  console.log(res.data)
			//         }).catch(error =>{
			//           console.log(error);
			// 		  console.log("失败")
			//         })
		
		
		
		
			// const result = await Axios.post('http://124.222.141.238:8088/v1/frontend/academy/searchByRule', {
			// 		  region: '福州',
			// 		  level: '985',
			// 		  type: '法学',
			// 		}).then(res =>{
			// 			console.log("成功");
			// 			console.log(res.data.data)
			// 		}).catch(error =>{
			// 		  console.log(error);
			// 		  console.log("失败")
			// 		})
			
			
			// uni.$on('refreshData',() => {
			// 	let pages = getCurrentPages();
			// 	let currPage = pages[pages.length - 1];
			// 	if(currPage.searchContent && currPage.searchContent != '') {
			// 		this.academyName = currPage.searchContent;
			// 		uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
								
			// 		}).then(res => {
			// 			this.isExist = true;
			// 			this.mes = res.data.data;
			// 			console.log(this.mes);
			// 		}).catch(err => {
			// 			this.mes = [];
			// 			this.isExist = false;
			// 		})
			// 		currPage.searchContent = '';
			// 	}
			// })
		
		},
		
		onShow(){  
			// let pages = getCurrentPages();
			// let currPage = pages[pages.length - 1];
			// if(currPage.searchContent && currPage.searchContent != '') {
			// 	this.academyName = currPage.searchContent;
			// 	uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
							
			// 	}).then(res => {
			// 		this.isExist = true;
			// 		this.mes = res.data.data;
			// 		console.log(this.mes);
			// 	}).catch(err => {
			// 		this.mes = [];
			// 		this.isExist = false;
			// 	})
			// 	currPage.searchContent = '';
			// }
		},

		// onShow() {
		//     // uni.$off('searchContent')//建议先销毁一次监听，再进行新的一次监听，否则会出现重复监听的现象
		// 	uni.$once('searchContent',function(data){
		// 		console.log("bbbb");
		// 		console.log(data);
		// 		this.academyName = data;
		// 		uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
							
		// 		}).then(res => {
		// 		  this.mes = res.data.data;
		// 		  this.mes.num = 1;
		// 		  console.log(this.mes.num);
		// 		  console.log(this.mes);
		// 		}).catch(err => {
							
		// 		})
		// 		//这的data就是B页面传递过来的数据
		// 	})
		// 	this.$forceUpdate();
		// },
		methods: {
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
			search() { 
				this.isShow = false;
				if(this.searchContent != '') {
					this.academyName = this.searchContent;
					uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
								
					}).then(res => {
						this.isExist = true;
						this.mes = res.data.data;
						console.log(this.mes);
					}).catch(err => {
						this.mes = [];
						this.isExist = false;
					})
					this.searchContent = '';
				}
			},
			
			
			reload() {
				this.$nextTick(() => {
					let pages = getCurrentPages();
					let currPage = pages[pages.length - 1];
					if(currPage.searchContent && currPage.searchContent != '') {
						this.academyName = currPage.searchContent;
						uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
									
						}).then(res => {
							this.isExist = true;
							this.mes = res.data.data;
							console.log(this.mes);
						}).catch(err => {
							this.mes = [];
							this.isExist = false;
						})
						currPage.searchContent = '';
					}
				})
			},
			touchStart(){
				this.active="background-color:#e6eff9"
			},
			touchEnd(){
				this.active=""
			},
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg = this.array1[this.index1];
				this.region = this.array1[this.index1];
				console.log(this.region);
				console.log(this.level);
				console.log(this.type);
				uni.$u.http.post('/v1/frontend/academy/searchByRule', {
					region: this.region,level: this.level,type: this.type,
					// region: '院校地区',level: '院校层次',type: '院校类型',
				}).then(res => {
					console.log("bbbbb")
					this.isExist = true;
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					this.isExist = false;
					console.log("aaaaa")
				})
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg = this.array2[this.index2];
				this.level = this.array2[this.index2];
				console.log(this.region);
				console.log(this.level);
				console.log(this.type);
				uni.$u.http.post('/v1/frontend/academy/searchByRule', {
					region: this.region,level: this.level,type: this.type,
					// region: '院校地区',level: '院校层次',type: '院校类型',
				}).then(res => {
					console.log("bbbbb")
					this.isExist = true;
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					this.isExist = false;
					console.log("aaaaa")
				})
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg = this.array3[this.index3];
				this.type = this.array3[this.index3];
				console.log(this.region);
				console.log(this.level);
				console.log(this.type);
				uni.$u.http.post('/v1/frontend/academy/searchByRule', {
					region: this.region,level: this.level,type: this.type,
					// region: '院校地区',level: '院校层次',type: '院校类型',
				}).then(res => {
					console.log("bbbbb")
					this.isExist = true;
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					this.isExist = false;
					console.log("aaaaa")
				})
			},
			goUniverity(c) {
				var _this = this
				uni.$u.http.get('/v1/frontend/academy/detail/' + c, {
				
				}).then(res => {
				    console.log(res.data.data);
					_this.sendMes = res.data.data[0];
				}).catch(err => {
					
				})
				uni.$u.http.get('/v1/frontend/recipe/list/' + c, {
					
				}).then(res => {
					_this.saidMes = res.data.data;
					console.log("成功")
					console.log(_this.saidMes);
				}).catch(err => {
					console.log("失败")
				})
				setTimeout(function() {
					uni.$emit('code1',{ 
						codeID: c ,
						belong: _this.sendMes.Belong,
						type: _this.sendMes.Type,
						code: _this.sendMes.Code,
						guide: _this.sendMes.Guide,
						level: _this.sendMes.Level,
						lineType: _this.sendMes.LineType,
						logo: _this.sendMes.Logo,
						name: _this.sendMes.Name,
						profile: _this.sendMes.Profile,
						region: _this.sendMes.Region,
						saidMes: _this.saidMes,
						})
				}, 500)
				// uni.setStorage({
				//     key: 'id',   // 存储值的名称
				//     data: 'hello',   //   将要存储的数据
				//     success:res => {
				//     	// 成功后的回调
				// 		console.log("成功传出去的数据");
				//         console.log(res);
				//     }
				// });
				uni.navigateTo({
					url: "/pages/home/university/university?code=" + c
					// url: "/pages/home/major"
					// url: "/pages/home/professional/professional"
				})
			},
		},
		mounted() {
			// 基本用法，注意：get请求的参数以及配置项都在第二个参数中
			//      uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
			
			//      }).then(res => {
			//        console.log(res.data.data);
			//        this.mes = res.data.data;
			//      }).catch(err => {
			
			//      })
			
			uni.$u.http.post('/v1/frontend/academy/searchByRule', {
				// region: '福州',level: '985',type: '法学',
				region: '院校地区',level: '院校层次',type: '院校类型',
			}).then(res => {
				console.log("bbbbb")
				this.isExist = true;
				this.mes = res.data.data;
				console.log(this.mes)
			}).catch(err => {
				this.mes = [];
				this.isExist = false;
				console.log("aaaaa")
			})
		}
	}
</script>

<style lang="scss">
.container{
	// background-image: linear-gradient(179deg,#83a4d488,#b6fbff88);
	// background-image: linear-gradient(178deg, #bac8ff, #c3fae855);
	background-color: #fafafa;
}
.wordSpace{
	letter-spacing: 4rpx;
}
.selectForm{
	display: flex;
	justify-content: center;
	margin-left: 10rpx;
}
.selectFormItem{
	width: 220rpx;
	height: 100rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	font-size: 30rpx;
}
.downArrow{
	margin-left: 5rpx;
	// padding-top: 10rpx;
	height: 20rpx;
	width: 20rpx;
}
.searchAcademy{
	background-color: #efefef55;
}
.searchText{
	display: flex;
	align-items: center;
	height: 80rpx;
	margin-left: 30rpx;
	font-size: 30rpx;
}
.searchNum{
	color: #bd3124
}


.viewAcademy {
	height: 120rpx;
	display: flex;
	flex-direction: row;
	align-items: center;
	  
	/* 圆角 */
	border-radius: 18rpx;
	
	/* 边 */
	border: 1rpx solid #E0E3DA;
	/* 阴影 */
	box-shadow:2rpx 7rpx 0rpx #d4d4d4;
	
	background-color: #ffffff88;
	margin-left:30rpx;
	margin-right:30rpx;
	margin-top: 25rpx;
	
	/* padding使得文字和图片不至于贴着边框 */
	padding: 25rpx;
	
}
.academyLogo{
	height: 120rpx;
	width: 120rpx;
}
.viewText {
	margin-left: 30rpx;
	display: flex;
	flex-direction: column;
}
.academyName{
	font-size: 36rpx;
	font-family: "黑体";
}
.academyType{
	display: flex;
	margin-top: 5rpx;
}
.typeOfScoreLine{
	width: 150rpx;
	height: 50rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	border-radius: 18rpx;
	background-color: #91d5ff;
	font-size: 26rpx;
	color: #ffffff;
	font-family: "黑体";
}
.typeOfInstitution{
	width: 100rpx;
	height: 50rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	margin-left: 20rpx;
	border-radius: 18rpx;
	background-color: #8fc9ff;
	font-size: 26rpx;
	color: #ffffff;
	font-family: "黑体";
}
.lacation{
	margin-left: 30rpx;
	width: 200rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	font-family: "黑体";
	font-size: 36rpx;
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
