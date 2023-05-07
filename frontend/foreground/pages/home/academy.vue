<template>
	<view>
		<view class="selectForm">
			<picker @change="bindPickerChange1" :range="array1" :value="index1" class="selectFormItem">
				<label class="">{{array1[index1]}}</label>
				<label class="downArrow">∨</label>
			</picker>
			<picker @change="bindPickerChange2" :range="array2" :value="index2" class="selectFormItem">
				<label class="">{{array2[index2]}}</label>
				<label class="downArrow">∨</label>
			</picker>
			<picker @change="bindPickerChange3" :range="array3" :value="index3" class="selectFormItem">
				<label class="">{{array3[index3]}}</label>
				<label class="downArrow">∨</label>
			</picker>
		</view>
		<view class="searchAcademy">
			<text class="searchText">
				共搜索到 <text class="searchNum">{{mes.num}}</text> 所院校
			</text>
		</view>
		<view>
			<view  class="viewAcademy" v-for="m in mes.list" @click="goUniverity(m.Code)" @touchstart="touchStart" @touchend="touchEnd" :style="active">
				<image class="academyLogo" src="@/static/academy-icons/北大.png"></image>
				<view class="viewText">
					<text class="academyName">{{m.Name}}</text>
					<view class="academyType">
						<view class="typeOfScoreLine">{{m.LineTtpe}}</view>
						<text class="typeOfInstitution">{{m.Level}}</text>
					</view>
				</view>
				<text class="lacation">{{m.Region}}</text>
			</view>
			
<!-- 			<view class="viewAcademy" v-for="m in mes.list" bindtap="click" @click="goUniverity">
				<image class="academyLogo" src="@/static/academy-icons/清华大学.png"></image>
				<view class="viewText">
					<text class="academyName">清华大学</text>
					<view class="academyType">
						<view class="typeOfScoreLine">自划线</view>
						<text class="typeOfInstitution">985</text>
					</view>
				</view>
				<text class="lacation">北京</text>
			</view> -->
			
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
				
				region: '院校地区',
				level: '院校层次',
				type: '院校类型',
				academyName: '福州大学'
			};
		},
		onNavigationBarButtonTap:function(e){
			uni.navigateTo({
				url: "/pages/home/university/search"
			})
		},
		async onLoad() {
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
		
		
		},
		
		onShow(){  
			let pages = getCurrentPages();
			let currPage = pages[pages.length - 1];
			if(currPage.searchContent) {
				this.academyName = currPage.searchContent;
				uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
							
				}).then(res => {
				  this.mes = res.data.data;
				  console.log(this.mes);
				}).catch(err => {
							
				})
			}
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
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg = this.array2[this.index2];
				this.level = this.array2[this.index2];
				console.log(this.region);
				console.log(this.level);
				console.log(this.type);
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg = this.array3[this.index3];
				this.type = this.array3[this.index3];
				console.log(this.region);
				console.log(this.level);
				console.log(this.type);
			},
			goUniverity(c) {
				uni.navigateTo({
					url: "/pages/home/university/university?code=" + c
				})
				uni.$emit('code1',c)
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
				region: '福州',level: '985',type: '法学',
			}).then(res => {
				this.mes = res.data.data;
				console.log(this.mes)
			}).catch(err => {
			
			})
		}
	}
</script>

<style lang="scss">
.selectForm{
	display: flex;
	justify-content: center;
}
.selectFormItem{
	width: 200rpx;
	height: 100rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	font-size: 30rpx;
}
.downArrow{
	margin-left: 5rpx;
}
.searchAcademy{
	background-color: #efefef;
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
	box-shadow:2rpx 7rpx 0rpx #ebebeb;
	
	background-color: #ffffff;
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
	background-color: #96C5F1;
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
	background-color: #96C5F1;
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

</style>
