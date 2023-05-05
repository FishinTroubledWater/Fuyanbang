<template>
  <view>
    <uni-card :is-shadow="false">
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
          <p>露从今夜白，月是故乡明。</p>
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