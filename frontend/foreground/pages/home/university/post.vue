<template>
	<view>
		<view>学长学姐说</view>
		<view>{{this.id}}</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				id: 0,
			}
		},
		onLoad:function(option){
			var _this = this
			this.id = option.id;
		},
	}
</script>

<style>
</style>