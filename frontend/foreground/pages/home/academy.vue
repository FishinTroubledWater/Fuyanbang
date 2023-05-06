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
				共搜索到 <text class="searchNum">{{searchnum}}</text> 所院校
			</text>
		</view>
		<uni-list>
			<view  class="viewAcademy" bindtap="click" @click="goUniverity" @touchstart="touchStart" @touchend="touchEnd" :style="active">
				<image class="academyLogo" src="../../static/academy-icons/北大.png"></image>
				<view class="viewText">
					<text class="academyName">北京大学</text>
					<view class="academyType">
						<view class="typeOfScoreLine">自划线</view>
						<text class="typeOfInstitution">985</text>
					</view>
				</view>
				<text class="lacation">北京</text>
			</view>
			<view  class="viewAcademy" bindtap="click" @click="goUniverity">
				<image class="academyLogo" src="../../static/academy-icons/清华大学.png"></image>
				<view class="viewText">
					<text class="academyName">清华大学</text>
					<view class="academyType">
						<view class="typeOfScoreLine">自划线</view>
						<text class="typeOfInstitution">985</text>
					</view>
				</view>
				<text class="lacation">北京</text>
			</view>
		</uni-list>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				array1: ['院校地区','全部','北京','福建'],
				array2: ['院校层次','985','211'],
				array3: ['院校类型','综合','理工','师范'],
				index1: 0,
				index2: 0,
				index3: 0,
				searchnum: 111,
				active:''
			};
		},
		methods: {
			touchStart(){
				this.active="background-color:#e6eff9"
			},
			touchEnd(){
				this.active=""
			},
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg=this.array1[this.index1]
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg=this.array2[this.index2]
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg=this.array3[this.index3]
			},
			goUniverity() {
				uni.navigateTo({
					url: "/pages/home/university/university"
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
	box-shadow:2rpx 6rpx 0rpx #B8B8B8;
	
	background-color: #ffffff;
	margin: 30rpx;
	
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
