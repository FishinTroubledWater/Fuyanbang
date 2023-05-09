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
		
		<uni-list>
			<view  class="viewMajor" bindtap="click" @click="goProfessional">
				<view class="viewText">
					<view class="majorText">
						<text class="majorName">(010101)哲学</text>
						<view class="typeOfDirection">学硕</view>
					</view>
					<view class="typeAndSubject">哲学-哲学</view>
				</view>
			</view>
			<view  class="viewMajor" bindtap="click" @click="goProfessional">
				<view class="viewText">
					<view class="majorText">
						<text class="majorName">(010101)马克思主义哲学</text>
						<view class="typeOfDirection">学硕</view>
					</view>
					<view class="typeAndSubject">哲学-哲学</view>
				</view>
			</view>
		</uni-list>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				array1: ['学科门类','全部','理学','工学','农学','哲学','经济学','法学','教育学','文学','历史学','医学','军事学','管理学','艺术学'],	
				array2: ['一级学科','数学','物理','化学'],
				array3: ['数学类型','数学一','数学二','数学三'],
				array4: ['外语类型','英语一','英语二'],
				index1: 0,
				index2: 0,
				index3: 0,
				index4: 0,
				
				active:'',
				mes: [],
				majorName: '',
				
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
				// uni.$u.http.get('/v1/frontend/academy/searchByName/' + this.academyName, {
							
				// }).then(res => {
				// 	this.isExist = true;
				// 	this.mes = res.data.data;
				// 	console.log(this.mes);
				// }).catch(err => {
				// 	this.mes = [];
				// 	this.isExist = false;
				// })
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
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg = this.array2[this.index2];
				this.firstLevelDiscipline = this.array2[this.index2];
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg = this.array3[this.index3];
				this.mathematicalType = this.array3[this.index3];
			},
			bindPickerChange4: function(e) {
				this.index4 = e.target.value;
				this.jg = this.array4[this.index4];
				this.foreignLanguageType = this.array4[this.index4];
			},
			goProfessional() {
				uni.navigateTo({
					url: "/pages/home/professional/professional"
				})
			},
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
