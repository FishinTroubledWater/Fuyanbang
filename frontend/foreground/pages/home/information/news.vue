<template>
  <view>
    <!-- 标签 -->
    <view class="tabs-box">
      <u-tabs :current="current" :list="tagList" lineWidth="32" lineHeight="2" lineColor="#f56c6c" :scrollable="false"
        :activeStyle="{
                 color: '#303133',
                 fontWeight: 'bold',
                 transform: 'scale(1.05)'
             }" :inactiveStyle="{
                 color: '#606266',
                 transform: 'scale(1)'
             }" itemStyle="padding-left: 16px; padding-right: 16px; height: 32px;" @change="change">
      </u-tabs>
    </view>
    <!-- 分割线 -->
    <u-line color="#262626"></u-line>

    <!-- 内容 -->
    <u-list @scrolltolower="scrolltolower">
      <u-list-item v-for="(item, index) in indexList" :key="index">
        <my-card :is-shadow="true" padding="10px" @click="gotoPage('/pages/home/information/detail', item.id)">
          <u--text :lines="1" :text="item.title" bold="" size="22" color="#000" lineHeight="24px"
            margin="10px 2px"></u--text>

          <!--          <u-row customStyle="margin: 4px" justify="space-between" gutter="4">
            <u-col span="8">
              <u--text :lines="3" :text="item.overview" size="14" color="#000" lineHeight="20px"
                margin="0px 4px 0px 0px" padding="0px 4px" height="94px"></u--text>
            </u-col>
            <u-col span="4">
              <u--image class="UImg" :src="item.img" radius="16rpx" shape="square" mode="aspectFill" width="110px"
                height="90px" margin="0px -16px 0px 8px"></u--image>
            </u-col> -->
          </u-row>

          <view class="saidContent">
            <u--text :lines="3" :text="item.overview" size="14" color="#000" lineHeight="20px" class="textContent"
              margin="0px" padding="0px 4px" height="94px"></u--text>
            <image class="sights" :src="item.img"></image>
          </view>

          <u--text :lines="1" :text="item.readNum" size="12" color="#000" lineHeight="12px" margin="2px 2px"></u--text>
        </my-card>
      </u-list-item>
    </u-list>
  </view>
</template>

<script>
  import MyCard from '../../../components/my-card/my-card.vue'
  import UniIcons from '../../../components/uni-icons/uni-icons.vue'

  export default {
    components: {
      MyCard,
      UniIcons
    },
    data() {
      return {
        // 标签列表
        tagList: [{
          name: '考研常识',
        }, {
          name: '考研政策',
        }, {
          name: '选择院校'
        }, {
          name: '备考指南'
        }],
        //选中的tag标签             
        current: 0,
        // 资讯列表
        indexList: [
          //   {
          //   id: "666",
          //   title: "考研考本校和外校的区别",
          //   overview: "考研到底要考本校还是考取别的院校呢？考研考取本校和考外校的区别在哪里。这些都是每一位",
          //   img: "https://cdn.uviewui.com/uview/album/1.jpg",
          //   readNum: (Math.floor(Math.random() * 90) + 10) + "," + (Math.floor(Math.random() * 900) + 100) + " 阅读"
          // },
        ],
        List1: [],
        List2: [],
        List3: [],
        List4: []
      };
    },
    methods: {
      // 页面跳转
      gotoPage(u, id) {
        uni.navigateTo({
          url: u + "?id=" + id,
        })
      },
      // 滚动到底部触发事件
      scrolltolower() {
        console.log("页面到底了")
      },
      // 点击标签
      change(index) {
        this.current = index.index;
        console.log(this.current);

        switch (this.current) {
          case 0:
            this.indexList = this.List1;
            break;
          case 1:
            this.indexList = this.List2;
            break;
          case 2:
            this.indexList = this.List3;
            break;
          case 3:
            this.indexList = this.List4;
            break;
        }

        // const pages = getCurrentPages();
        // pages.mounted();
      }
    },
    mounted() {
      // 基本用法，注意：get请求的参数以及配置项都在第二个参数中
      uni.$u.http.get('/v1/frontend/news/detail', {

      }).then(res => {
        console.log(res.data.data.newses)

        for (var i = 1; i < res.data.data.newses.length; i++) {
          let tmp = res.data.data.newses[i];
          console.log(tmp.Type);

          if (tmp.Type == "考研常识") {
            this.List1.push({
              id: tmp.ID,
              title: tmp.Title,
              overview: tmp.Content,
              img: "https://cdn.uviewui.com/uview/album/1.jpg",
              readNum: (Math.floor(Math.random() * 90) + 10) + "," + (Math.floor(Math.random() * 900) + 100) +
                " 阅读",
              contentType: tmp.Type
            })
          }
          if (tmp.Type == "考研政策") {
            this.List2.push({
              id: tmp.ID,
              title: tmp.Title,
              overview: tmp.Content,
              img: "https://cdn.uviewui.com/uview/album/1.jpg",
              readNum: (Math.floor(Math.random() * 90) + 10) + "," + (Math.floor(Math.random() * 900) + 100) +
                " 阅读",
              contentType: tmp.Type
            })
          }
          if (tmp.Type == "选择院校") {
            this.List3.push({
              id: tmp.ID,
              title: tmp.Title,
              overview: tmp.Content,
              img: "https://cdn.uviewui.com/uview/album/1.jpg",
              readNum: (Math.floor(Math.random() * 90) + 10) + "," + (Math.floor(Math.random() * 900) + 100) +
                " 阅读",
              contentType: tmp.Type
            })
          }
          if (tmp.Type == "备考指南") {
            this.List4.push({
              id: tmp.ID,
              title: tmp.Title,
              overview: tmp.Content,
              img: "https://cdn.uviewui.com/uview/album/1.jpg",
              readNum: (Math.floor(Math.random() * 90) + 10) + "," + (Math.floor(Math.random() * 900) + 100) +
                " 阅读",
              contentType: tmp.Type
            })
          }
        }
        
        this.indexList = this.List1;
        
      }).catch(err => {
        console.log("出错了..." + err)
      })
    }
  }
</script>

<style lang="scss" scoped>
  // 标签列表
  .tabs-box {
    //   margin: 12px 6px;
    //   padding: 6px 12px;
    //   flex-flow: row;
    //   justify-content: space-around;
    //   display: flex;
  }

  .saidContent {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  .sights {
    float: right;
    width: 190rpx;
    height: 125rpx;
    border-radius: 18rpx;
  }

  .textContent {
    width: 450rpx;
    margin-top: 0rpx;
    margin-right: 20rpx;
    font-size: 20rpx;
  }
</style>