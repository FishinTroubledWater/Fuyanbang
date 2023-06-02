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
					<image class="downArrow" src="@/static/academy-icons/down.png"></image>
				</picker>
				<picker @change="bindPickerChange2" :range="array2" :value="index2" class="selectFormItem">
					<label class="wordSpace">{{array2[index2]}}</label>
					<!-- <label class="downArrow">∨</label> -->
					<image class="downArrow" src="@/static/academy-icons/down.png"></image>
				</picker>
				<picker @change="bindPickerChange3" :range="array3" :value="index3" class="selectFormItem">
					<label class="wordSpace">{{array3[index3]}}</label>
					<!-- <label class="downArrow">∨</label> -->
					<image class="downArrow" src="@/static/academy-icons/down.png"></image>
				</picker>
				<picker @change="bindPickerChange4" :range="array4" :value="index4" class="selectFormItem">
					<label class="wordSpace">{{array4[index4]}}</label>
					<!-- <label class="downArrow">∨</label> -->
					<image class="downArrow" src="@/static/academy-icons/down.png"></image>
				</picker>
			</view>
			<view>
				<view  class="viewMajor" bindtap="click" @click="goProfessional(m.Code)" v-for="m in mes.list">
					<view class="viewText">
						<view class="majorText">
							<text class="majorName">({{m.Code}}){{m.Name}}</text>
							<view class="typeOfDirection">学硕</view>
						</view>
						<view class="typeAndSubject">{{m.SubjectCategory}}-{{m.FirstLevelDiscipline}}</view>
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
				isShow: false,
				
				array1: ['学科门类','理学','工学','农学','哲学','经济学','法学','教育学','文学','历史学','医学','军事学','管理学','艺术学'],	
				array2: ['一级学科','数学','物理','理论经济学','哲学'],
				array3: ['数学类型','数学一','数学二','数学三'],
				array4: ['外语类型','英语一','英语二'],
				index1: 0,
				index2: 0,
				index3: 0,
				index4: 0,
				
				// test: 'aaa bbb ccc',
				active:'',
				mes: [],
				sendMes: [],
				majorName: '',
				isExist: false,
				
				code: 0,
				
				mathematicalType: '数学类型',
				foreignLanguageType: '外语类型',
				firstLevelDiscipline: '一级学科',
				type: '学科门类',
				
				
				// 初始化定时器为空
				time:null,
				// 用户输入的关键词
				keyword:'',
				//搜索数据的数组初始化
				searchList:[],
				//搜索历史初始化
				historyList:[],
				// 初始化搜索发现列表
				foundList:['计算机','软件工程','土木工程','经济与管理学','电子信息','通讯工程'],
				searchContent: '',
			};
		},
		onNavigationBarButtonTap:function(e){
			this.isShow = true;
		 //    console.log(e.text);//提交
			// uni.navigateTo({
			// 	url: "/pages/home/professional/search"
			// })
		},
		onShow(){
			// let pages = getCurrentPages();
			// let currPage = pages[pages.length - 1];
			// if(currPage.searchContent && currPage.searchContent != '') {
			// 	this.majorName = currPage.searchContent;
			// 	uni.$u.http.get('/v1/frontend/major/searchByName/' + this.majorName, {
							
			// 	}).then(res => {
			// 		this.mes = res.data.data;
			// 		console.log(this.mes);
			// 	}).catch(err => {
			// 		this.mes = [];
			// 	})
			// 	console.log(this.majorName)
			// 	currPage.searchContent = '';
			// }
		},
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
			search(){
				this.isShow = false;
				if(this.searchContent != '') {
					this.majorName = this.searchContent;
					uni.$u.http.get('/v1/frontend/major/searchByName/' + this.majorName, {
								
					}).then(res => {
						this.mes = res.data.data;
						console.log(this.mes);
					}).catch(err => {
						this.mes = [];
					})
					console.log(this.majorName)
					this.searchContent = '';
				}
			},
			
			
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg = this.array1[this.index1];
				this.type = this.array1[this.index1];
				console.log(this.mathematicalType);
				console.log(this.foreignLanguageType);
				console.log(this.firstLevelDiscipline);
				console.log(this.type);
				uni.$u.http.post('/v1/frontend/major/searchByRule', {
					subjectCategory: this.type,firstLevelDiscipline: this.firstLevelDiscipline,mathType: this.mathematicalType,foreignType: this.foreignLanguageType,
				}).then(res => {
					console.log("bbbbb")
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					console.log("aaaaa")
				})
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg = this.array2[this.index2];
				this.firstLevelDiscipline = this.array2[this.index2];
				uni.$u.http.post('/v1/frontend/major/searchByRule', {
					subjectCategory: this.type,firstLevelDiscipline: this.firstLevelDiscipline,mathType: this.mathematicalType,foreignType: this.foreignLanguageType,
				}).then(res => {
					console.log("bbbbb")
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					console.log("aaaaa")
				})
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg = this.array3[this.index3];
				this.mathematicalType = this.array3[this.index3];
				uni.$u.http.post('/v1/frontend/major/searchByRule', {
					subjectCategory: this.type,firstLevelDiscipline: this.firstLevelDiscipline,mathType: this.mathematicalType,foreignType: this.foreignLanguageType,
				}).then(res => {
					console.log("bbbbb")
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					console.log("aaaaa")
				})
			},
			bindPickerChange4: function(e) {
				this.index4 = e.target.value;
				this.jg = this.array4[this.index4];
				this.foreignLanguageType = this.array4[this.index4];
				uni.$u.http.post('/v1/frontend/major/searchByRule', {
					subjectCategory: this.type,firstLevelDiscipline: this.firstLevelDiscipline,mathType: this.mathematicalType,foreignType: this.foreignLanguageType,
				}).then(res => {
					console.log("bbbbb")
					this.mes = res.data.data;
					console.log(this.mes)
				}).catch(err => {
					this.mes = [];
					console.log("aaaaa")
				})
			},
			goProfessional(c) {
				var _this = this
				uni.$u.http.get('/v1/frontend/major/detail/' + c, {
				
				}).then(res => {
					console.log("ccccc");
					console.log(res.data.data);
					_this.sendMes = res.data.data[0];
				}).catch(err => {
					
				})
				setTimeout(function() {
					console.log("专业详细信息：");
					console.log(c);
					console.log(_this.sendMes);
					uni.$emit('code2',{ 
						code: _this.sendMes.Code,
						firstLevelDiscipline: _this.sendMes.FirstLevelDiscipline,
						jobOrientation: _this.sendMes.JobDirection,
						JobProspect: _this.sendMes.JobProspect,
						name: _this.sendMes.Name,
						scoreUrl: _this.sendMes.ScoreUrl,
						profile: _this.sendMes.Profile,
						secondLevelDiscipline: _this.sendMes.SecondLevelDiscipline,
						subjectCategory: _this.sendMes.SubjectCategory,
						})
				}, 500)
				uni.navigateTo({
					url: "/pages/home/professional/professional"
				})
			},
		},
		mounted() {
			uni.$u.http.post('/v1/frontend/major/searchByRule', {
				// subjectCategory: '哲学',firstLevelDiscipline: '哲学',mathType: '数学一',foreignType: '英语一',
				subjectCategory: '学科门类',firstLevelDiscipline: '一级学科',mathType: '数学类型',foreignType: '外语类型',
			}).then(res => {
				console.log("bbbbb")
				this.mes = res.data.data;
				console.log(this.mes)
			}).catch(err => {
				this.mes = [];
				console.log("aaaaa")
			})
		},
		onLoad() {
			// uni.getStorage({
			// 	key:'id',   // 储存在本地的变量名
			// 	success:res => {
			// 		// 成功后的回调
			// 		console.log("接收到的数据"); 
			// 		console.log(res.data);   // hello  这里可做赋值的操作
			// 	}
			// })
			
			
			// uni.$on('refreshData1',() => {
			// 	let pages = getCurrentPages();
			// 	let currPage = pages[pages.length - 1];
			// 	if(currPage.searchContent && currPage.searchContent != '') {
			// 		this.majorName = currPage.searchContent;
			// 		uni.$u.http.get('/v1/frontend/major/searchByName/' + this.majorName, {
								
			// 		}).then(res => {
			// 			this.mes = res.data.data;
			// 			console.log(this.mes);
			// 		}).catch(err => {
			// 			this.mes = [];
			// 		})
			// 		console.log(this.majorName)
			// 		currPage.searchContent = '';
			// 	}
			// })
		}
	}
</script>

<style lang="scss">
.container{
	// background-image: linear-gradient(179deg,#83a4d488,#b6fbff88);
	background-color: #fafafa;
}
.selectForm{
	display: flex;
	justify-content: center;
	border-bottom-style:solid;
	border-width: 1rpx;
	border-color:#c1c1c1;
}
.wordSpace{
	letter-spacing: 2rpx;
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
.viewMajor {
	height: 100rpx;
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
.viewText {
	display: flex;
	flex-direction: column;
}
.majorName{
	font-size: 34rpx;
	font-family: "思源黑体";
	
}
.majorText{
	display: flex;

}
.typeOfDirection{
	width: 90rpx;
	height: 35rpx;
	display: flex;
	justify-content: center;
	align-items: center;
	background-color: #F06977;
	font-size: 26rpx;
	color: #ffffff;
	font-family: "思源黑体";
	margin-left: 20rpx;
	margin-top: 5rpx;
	border-style:solid;
	border-width:3rpx;
	border-color:#c3c3c3;
	border-radius: 10rpx;
}
.typeAndSubject{
	margin-top: 10rpx;
	margin-left: 20rpx;
	font-size: 28rpx;
	font-family: "思源黑体";
	color: #A8A8A8;
}
.downArrow{
	margin-left: 5rpx;
	height: 20rpx;
	width: 20rpx;
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
