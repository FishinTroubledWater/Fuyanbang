<template>
  <view>
    <uni-card :is-shadow="true">
      <u--text :text="title" size="24" lineHeight="24" margin="4px"></u--text>
      <span style="margin: 16px 16px;font-size: 14px;line-height: 24px;">{{author}}</span>
      <span style="margin-right: 16px;font-size: 14px;float: right;">{{time}}</span>
      <view class="u-content">
        <u-parse :content="content"></u-parse>
      </view>
    </uni-card>
  </view>
</template>

<script>
  import uCard from '../../components/uni-card/uni-card.vue'

  export default {
    components: {
      uCard
    },
    data() {
      return {
        title: "考研考本校和外校的区别，你一定得知道！",
        author: "中国青年报",
        time: "2022-06-22",
        content: `
					<p>考研到底是要考本校还是外校呢？这是每一个考研学子都应该要知道的。</p>
					<img src="https://cdn.uviewui.com/uview/swiper/2.jpg" />
          <p>首先，本校是你已经生活过四年的地方，当然有的同学是五年，在这为数不短的日子里相信你对校园一定了如指掌。哪里有好玩的，哪家的饭菜最好吃，哪里最适合约会等等。这些你都比任何人清楚，所以在你考上之后，基本上你会感觉和本科不会有太多的不同。</p>
          <p>而考取外校的话你可以进入新的环境，认识新的同学，可以把以前的黑历史都抹杀掉，不会有人认识你。但是这也意味着你可能需要一段时间来熟悉校园，甚至上课都需要拿着校园地图来问师兄师姐们。而且哪家店最实惠你也需要多吃几次亏才能明白。没有对比就没有伤害，你会经常拿现在的环境和以前的学校对比，从景色到食堂美味，这些都和以前不同。</p>
          <p>考研的话毫无疑问是考本校比较简单。首先获取资源那是相当的容易，你可以直接去找师兄师姐们，询问他们考试重点，或者考试注意事项。你的初试成绩达到了本校的复试线，那你几乎已经确定考上研究生了，因为本校的保护政策，一般学校会优先录取本校的学生。而且老师也比较熟悉你，更加愿意教自己的学生。甚至你可以直接去找学校的老师，一般来说学校的老师都很“护犊子”，老师对于本校的学生都很关爱，一定会帮你划重点的哟!</p>
				`
      }
    },
    methods: {

    }
  }
</script>

<style scoped>
  /* 资讯内容 */
  .u-content {
    padding: 4rpx;
    font-size: 24rpx;
    color: #000;
    line-height: 1.6;
  }
</style>