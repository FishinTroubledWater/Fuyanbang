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
			<picker @change="bindPickerChange4" :range="array4" :value="index4" class="selectFormItem">
				<label class="">{{array4[index4]}}</label>
				<label class="downArrow">∨</label>
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
<!-- 			<view  class="viewMajor" bindtap="click" @click="goProfessional">
				<view class="viewText">
					<view class="majorText">
						<text class="majorName">(010101)马克思主义哲学</text>
						<view class="typeOfDirection">学硕</view>
					</view>
					<view class="typeAndSubject">哲学-哲学</view>
				</view>
			</view> -->
		</view>
<!-- 		<view v-for="t in test.substr(0,test.length).split(' ')">
			{{t}}
		</view> -->
	</view>
</template>

<script>
	export default {
		data() {
			return {
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
				type: '学科门类'
			};
		},
		onNavigationBarButtonTap:function(e){
		    console.log(e.text);//提交
			uni.navigateTo({
				url: "/pages/home/professional/search"
			})
		},
		onShow(){
			let pages = getCurrentPages();
			let currPage = pages[pages.length - 1];
			if(currPage.searchContent && currPage.searchContent != '') {
				this.majorName = currPage.searchContent;
				uni.$u.http.get('/v1/frontend/major/searchByName/' + this.majorName, {
							
				}).then(res => {
					this.mes = res.data.data;
					console.log(this.mes);
				}).catch(err => {
					this.mes = [];
				})
				console.log(this.majorName)
				currPage.searchContent = '';
			}
		},
		methods: {
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
						jobOrientation: _this.sendMes.JobOrientation,
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
.viewMajor {
	height: 80rpx;
	display: flex;
	flex-direction: row;
	// align-items: center;
	
	/* 边 */
	border: 1rpx solid #E0E3DA;
	/* 阴影 */
	box-shadow:2rpx 4rpx 0rpx #ebebeb;
	
	background-color: #ffffff;

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
}
.typeAndSubject{
	margin-top: 10rpx;
	margin-left: 20rpx;
	font-size: 28rpx;
	font-family: "思源黑体";
	color: #A8A8A8;
}
</style>
