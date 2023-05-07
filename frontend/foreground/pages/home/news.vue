<template>
  <view>
    <!-- 标签 -->
    <view class="tabs-box">
      <u-tag text="考研常识" type="success" plain shape="circle" plain size="large" :plainFill="true"></u-tag>
      <u-tag text="考研政策" plain shape="circle" plain size="large" :plainFill="true"></u-tag>
      <u-tag text="选择院校" plain shape="circle" plain size="large" :plainFill="true"></u-tag>
      <u-tag text="备考指南" plain shape="circle" plain size="large" :plainFill="true"></u-tag>
    </view>
    <!-- 分割线 -->
    <u-line color="#262626"></u-line>
    <!-- 内容 -->
    <u-list @scrolltolower="scrolltolower">
      <u-list-item v-for="(item, index) in indexList" :key="index">
        <uni-card :is-shadow="true" padding="10px">
          <u--text :line="1" :text="item.title" size="22" color="#000" lineHeight="24px" margin="10px 2px"></u--text>

          <u-row customStyle="margin: 4px" justify="space-between" gutter="2">
            <u-col span="8">
              <u--text :line="3" :text=item.overview size="16" color="#000" lineHeight="20px" margin="6px 2px"
                height="94px"></u--text>
            </u-col>
            <u-col span="4">
              <u--image :src="item.img" shape="square" mode="aspectFill" width="150px" height="94px"></u--image>
            </u-col>
          </u-row>

          <u--text :line="1" :text="item.readNum" size="12" color="#000" lineHeight="12px" margin="2px 2px"></u--text>
        </uni-card>
      </u-list-item>
    </u-list>
  </view>
</template>

<script>
  import uCard from '../../components/uni-card/uni-card.vue'
  import uIcons from '../../components/uni-icons/uni-icons.vue'

  export default {
    components: {
      uCard,
      uIcons
    },
    data() {
      return {
        indexList: [{
          id: "1",
          title: "考研考本校和外校的区别",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/1.jpg",
          readNum: "65,210阅读"
        }, {
          id: "2",
          title: "研究生就读途径有哪些",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/2.jpg",
          readNum: "76,453阅读"
        }, {
          id: "3",
          title: "研究生就读途径有哪些",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/3.jpg",
          readNum: "47,099阅读"
        }, {
          id: "4",
          title: "考研考本校和外校的区别",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/4.jpg",
          readNum: "16,120阅读"
        }, {
          id: "5",
          title: "考研考本校和外校的区别",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/5.jpg",
          readNum: "14,111阅读"
        }, {
          id: "6",
          title: "考研考本校和外校的区别",
          overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          img: "https://cdn.uviewui.com/uview/album/6.jpg",
          readNum: "55,211阅读"
        }]
      };
    },
    methods: {
      // 滚动到底部触发事件
      scrolltolower() {
        console.log("页面到底了")
      },
      // 点击标签
      clickTag() {

      }
    }
  }
</script>

<style lang="scss" scoped>
  // 标签列表
  .tabs-box {
    margin: 12px 6px;
    padding: 6px 12px;
    flex-flow: row;
    justify-content: space-around;
    display: flex;
  }
</style>