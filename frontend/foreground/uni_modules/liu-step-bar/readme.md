### liu-step-bar适用于uni-app项目的步骤条组件
### 本组件目前兼容微信小程序、H5
### 本组件支持自定义步骤内容、步骤回显、点击
### 如使用过程中有问题或有一些好的建议，欢迎qq联系：2364518038

### 使用方式（uni_modules版本）
``` html
<liu-step-bar :step="step" :stepList="stepList" @clickStep="clickStep"></liu-step-bar>
```
``` javascript
export default {
	data() {
		return {
			step: 1, //当前步骤
			stepList: [{
				name: '第一步',
				id: 1
			}, {
				name: '第二步',
				id: 2
			}, {
				name: '第三步',
				id: 3
			}], //步骤列表
		};
	},
	methods: {
		//点击步骤
		clickStep(e){
			this.step = e.id
			console.log('============:', e)
		}
	}
}
```

### 使用方式2（uni_modules版本）
``` html
<liu-step-bar :step="step" :stepList="stepList" @clickStep="clickStep"></liu-step-bar>
```
``` javascript
export default {
	data() {
		return {
			step: 1, //当前步骤
			stepListTab: [{
					name: '选择房屋',
					id: 1,
					checkedImg:'../../static/yhRegist/step1.png',
					unCheckedImg:'../../static/yhRegist/step-one.png',
					checkedColor:'red',
					unCheckedColor:'blue',
					checkedLine:'yellow',
					unCheckedLine:'green'
				}, {
					name: '完善信息',
					id: 2,
					checjedImg:'../../static/yhRegist/step2.png',
					unCheckedImg:'../../static/yhRegist/step-two.png',
				}, {
					name: '提交审核',
					id: 3,
					checkedImg:'../../static/yhRegist/step3.png',
					unCheckedImg:'../../static/yhRegist/step-three.png',
				}] //支持字体颜色，线条等变量
		};
	},
	methods: {
		//点击步骤
		clickStep(e){
			this.step = e.id
			console.log('============:', e)
		}
	}
}
```


### 属性说明
| 名称                         | 类型            | 默认值                | 描述             |
| ----------------------------|--------------- | ---------------------- | ---------------|
| step                        | Number         | 1             | 当前步骤
| stepList                    | Array          | []            | 步骤信息
| checkedImg                  | String         |      checkedImg      | 已完成的图片路径
| unCheckedImg                | String         |   unCheckedImg         | 未完成的图片路径
| checkedColor                | String         | #287BF8             | 已完成的字体颜色
| unCheckedColor              | String          | #333333          | 未完成的字体颜色
| checkedLine                 | String         |  #287BF8                  | 已完成的线条颜色
| unCheckedLine               | String          | #bebebe            | 未完成的线条颜色