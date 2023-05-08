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
		<view>
			<view class="lineTable">
				<view class="tableTitle">北京大学建筑学硕士考研近3年复试分数线</view>
				<table class="line">
					<tr>
						<td>年份</td>
						<td>数学</td>
						<td>政治</td>
						<td>外语</td>
						<td>专业</td>
						<td>总分</td>
						<td>复试</td>
						<td>录取</td>
					</tr>
					<tr>
						<td>2018</td>
						<td>90</td>
						<td>60</td>
						<td>60</td>
						<td>90</td>
						<td>380</td>
						<td>11</td>
						<td>7</td>
					</tr>
					<tr>
						<td>2019</td>
						<td>90</td>
						<td>60</td>
						<td>60</td>
						<td>90</td>
						<td>380</td>
						<td>4</td>
						<td>2</td>
					</tr>
				</table>
			</view>
			<view class="lineTable">
				<view class="tableTitle">北京大学建筑学硕士考研近3年复试分数线</view>
				<table class="line">
					<tr>
						<td>年份</td>
						<td>数学</td>
						<td>政治</td>
						<td>外语</td>
						<td>专业</td>
						<td>总分</td>
						<td>复试</td>
						<td>录取</td>
					</tr>
					<tr>
						<td>2018</td>
						<td>90</td>
						<td>60</td>
						<td>60</td>
						<td>90</td>
						<td>380</td>
						<td>11</td>
						<td>7</td>
					</tr>
				</table>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				mes: [],
				code: 0,
				
				array1: ['学科门类','全部','理学','工学','农学','哲学','经济学','法学','教育学','文学','历史学','医学','军事学','管理学','艺术学'],
				array2: ['一级学科','数学','物理','化学'],
				array3: ['二级学科','基础数学','计算数学','概率论与数理统计'],
				index1: 0,
				index2: 0,
				index3: 0,
				type: '学科门类',
				firstLevelDiscipline: '一级学科',
				secondLevelDiscipline: '二级学科'
			};
		},
		onShow() {
			// console.log('ddd')
		},
		created() {
			uni.$on('code',res=>{
				this.code = res
				uni.$u.http.get('/v1/frontend/academy/detail/' + this.code, {
				
				}).then(res => {
					this.mes = res.data.data;
					console.log("成功3")
					console.log(this.mes);
				}).catch(err => {
					console.log("失败")
				})
			})
		},
		methods: {
			bindPickerChange1: function(e) {
				this.index1 = e.target.value;
				this.jg = this.array1[this.index1];
				this.type = this.array1[this.index1];
				console.log(this.type);
				console.log(this.firstLevelDiscipline);
				console.log(this.secondLevelDiscipline);
			},
			bindPickerChange2: function(e) {
				this.index2 = e.target.value;
				this.jg = this.array2[this.index2];
				this.firstLevelDiscipline = this.array2[this.index2];
				console.log(this.type);
				console.log(this.firstLevelDiscipline);
				console.log(this.secondLevelDiscipline);
			},
			bindPickerChange3: function(e) {
				this.index3 = e.target.value;
				this.jg = this.array3[this.index3];
				this.secondLevelDiscipline = this.array3[this.index3];
				console.log(this.type);
				console.log(this.firstLevelDiscipline);
				console.log(this.secondLevelDiscipline);
			},
		}
	}
</script>

<style>
	.selectForm{
		display: flex;
		justify-content: center;
	}
	.selectFormItem{
		width: 200rpx;
		height: 80rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 32rpx;
	}
	.downArrow{
		margin-left: 5rpx;
	}
	.lineTable{
		height: auto;
		border: 2rpx solid #BBBBBB;
		margin: 20rpx;
		border-radius: 18rpx;
		margin-bottom: 40rpx;
	}
	.tableTitle{
		display: flex;
		justify-content: center;
		align-items: center;
		font-weight: 700;
	}
	.line{
		text-align: center;
		margin-top: 15rpx;
	}
	table td {
	  width: 80rpx;
	  height: 60rpx;
	}
</style>