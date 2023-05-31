<template>
  <view class="content">

    <uni-section class="header-box" title="" sub-title="" type="">

      <!-- 倒计时 -->
      <u-row customStyle="margin: 0rpx 24rpx">
        <u-col span="8">
          <u--text :text="RemainingTime" size="72rpx" color="#ffffff" class="countDown"></u--text>
        </u-col>
      </u-row>

      <u-row customStyle="margin: 2px">
        <u-col span="3">
          <u--text></u--text>
        </u-col>
        <u-col span="4">
          <u--text></u--text>
        </u-col>
        <u-col span="5">
          <u--text :text="ExamYear" size="30rpx" color="#ffffff" class="yearDown"></u--text>
        </u-col>
      </u-row>

      <!--      <u-row style="margin: 16rpx 8rpx; padding: 0rpx 12rpx;">
        <u--text :line="2" text="满怀激情地投入到对事理的探究中来，至于建筑变成了自然而然的结果" size="30rpx" lineHeight="30rpx" color="#f1f1f1"></u--text>
      </u-row> -->
      <!-- <view class="remark">
        满怀激情地投入到对事理的探究中来，至于建筑变成了自然而然的结果。
      </view> -->

      <!-- 时间轴 -->
      <view class="timeline">
        <uni-steps deactive-color="#ffffff" active-color="#00aa00" :options="TimeList" active-icon="flag"
          :active="active" />
      </view>

    </uni-section>

    <!-- 标签按钮 -->
    <u-row class="tabs-box">
      <!-- 院校 -->
      <u-col span="6" style="margin: 0 -40rpx 0 24rpx;height: 256rpx;">
        <my-card :is-shadow="true" class="tabs-box-item" @click="gotoPage('/pages/home/university/academy')"
          style="height: 600rpx;">
          <u-row style="margin: 2rpx 2rpx;">
            <u-col span="7">
              <u--text text="院 校" size="20" color="#ffffff"></u--text>
            </u-col>
            <u-col span="5" offset="1">
              <u-avatar src="/static/home-images/academy.png" shape="square" size="36"></u-avatar>
            </u-col>
          </u-row>
          <u-row>
            <u--text text="选院校的确是一件十分不容易的事情，专业的选择为我们敲定了努力的方向，学校的选择才是给我们确定了准确的目标。" :lines="2" size="12"
              color="#e3e3e3"></u--text>
          </u-row>
        </my-card>
      </u-col>

      <!-- 专业和资讯 -->
      <u-col span="6" style="margin: 0 4rpx 0 -4rpx;height: 256rpx;">
        <my-card :is-shadow="true" class="tabs-box-item" @click="gotoPage('/pages/home/professional/major')" style="margin-bottom: -16rpx;
          background-image: linear-gradient(270deg, #614385ee, #516395ee);">
          <u-row style="margin-top: -4rpx;">
            <u-col span="7" offset="1">
              <u--text text="专 业" size="20" color="#ffffff"></u--text>
            </u-col>
            <u-col span="5" offset="-1">
              <u-avatar src="/static/home-images/major.png" shape="square" size="36"></u-avatar>
            </u-col>
          </u-row>
        </my-card>

        <my-card :is-shadow="true" class="tabs-box-item" @click="gotoPage('/pages/home/information/news')" style="
          background-image: linear-gradient(270deg, #614385ee, #516395ee);">
          <u-row style="margin-top: -6rpx;">
            <u-col span="7" offset="1">
              <u--text text="资 讯" size="20" color="#ffffff"></u--text>
            </u-col>
            <u-col span="5" offset="-1">
              <u-avatar src="/static/home-images/news.png" shape="square" size="36"></u-avatar>
            </u-col>
          </u-row>
        </my-card>
      </u-col>

    </u-row>

    <!-- 动态 -->
    <view class="trends-box">
      <view class="trends-box-title">最新动态</view>
      <u-line color="#000000"></u-line>
      <!--      <u-list @scrolltolower="scrolltolower">
        <u-list-item v-for="(item, index) in indexList" :key="index">
          <uni-card :is-shadow="true" :title="item.title" :sub-title="item.subTitle" :extra="item.publishTime"
            :thumbnail="item.pageImage" class="trends-box-item" @click="gotoPage('/pages/home/detail', item.id)">
            <u--text :lines="1" :text="item.content"></u--text>
          </uni-card>
        </u-list-item>
      </u-list> -->

      <u-skeleton v-if="isArticleLoading" class="downloading" rows="5" title
        :rowsWidth="['640rpx', '640rpx', '640rpx', '640rpx', '640rpx']"
        :rowsHeight="['20px', '20px', '20px', '20px', '20px']" loading></u-skeleton>

      <u-list @scrolltolower="scrolltolower">
        <u-list-item v-for="(item, index) in indexList" :key="index">
          <view class="viewSaid" @click="gotoDetail('/pages/home/information/detail', item.id)">
            <!-- 					<view class="title" v-if="item.title.length >= 16">{{item.title.substr(0,17)}}...</view>
					<view class="title" v-else>{{item.title}}</view> -->
            <u--text :lines="1" :text="item.title" bold="" size="36rpx" margin="0 10rpx"></u--text>
            <view class="viewUser">
              <image class="headPortrait" src="@/static/academy-icons/photo.jpg"></image>
              <view class="userMes">
                <!-- <text class="userName">{{item.subTitle.substr(0,10)}}</text> -->
                <u--text :lines="1" :text="item.subTitle" size="26rpx" margin="0 10rpx"></u--text>
              </view>
              <view class="publishTime">{{item.publishTime}}</view>
            </view>
            <view class="saidContent">
              <!-- <view class="textContent">{{item.content.substr(0,30)}}...</view> -->
              <u--text :lines="3" :text="item.content" size="20rpx" margin="0 10rpx"></u--text>
              <image class="sights" src="@/static/academy-icons/sight.png"></image>
            </view>
          </view>
        </u-list-item>
      </u-list>

    </view>

  </view>
</template>

<script>
  import Axios from 'axios'
  Axios.defaults.baseURL = '/'
  // eslint-disable-next-line no-unused-vars
  const axios = require('axios')

  import MyCard from '../../components/my-card/my-card.vue'
  import uSteps from '../../uni_modules/uni-steps/uni-steps.vue'
  import UniIcons from '../../components/uni-icons/uni-icons.vue'
  import uSection from '../../uni_modules/uni-section/uni-section.vue'

  export default {
    components: {
      MyCard,
      uSteps,
      UniIcons,
      uSection
    },
    data() {
      return {
        // 时间轴
        active: 3,

        //文章是否显示
        isArticleLoading: true,

        // 最新动态
        indexList: [
          //   {
          //   id: "666",
          //   title: "快讯!2023年考研国家线发布",
          //   subTitle: "教育部",
          //   publishTime: "2023年03月10日",
          //   pageImage: "/static/building.png",
          //   content: "近日，教育部部署2023年全国硕士研究生招生复试录取工作..."
          // },
        ],

        cover: 'https://web-assets.dcloud.net.cn/unidoc/zh/shuijiao.jpg',
        avatar: 'https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png',
        extraIcon: {
          color: '#4cd964',
          size: '22',
          type: 'gear-filled'
        },
        urls: [
          'https://cdn.uviewui.com/uview/album/1.jpg',
          'https://cdn.uviewui.com/uview/album/2.jpg',
        ]
      }
    },
    computed: {
      // 考研倒计时
      RemainingTime() {
        //date1当前时间
        let date1 = new Date();
        //date2结束时间
        let date2 = new Date(date1.getFullYear(), "11", "25", date1.getHours(), date1.getMinutes(), date1.getSeconds(),
          date1.getMilliseconds());
        const diff = date2.getTime() - date1.getTime(); //目标时间减去当前时间
        const diffDate = diff / (24 * 60 * 60 * 1000); //计算当前时间与结束时间之间相差天数
        if (diffDate < 0) diffDate + 365;
        return diffDate + "天";
      },
      // 考研年份
      ExamYear() {
        let date = new Date();
        return date.getFullYear() + "考研倒计时"
      },
      // 今天星期几
      TimeList() {
        let date = new Date();
        var weekday = new Array(7);
        weekday[0] = "SUN";
        weekday[1] = "MON";
        weekday[2] = "TUE";
        weekday[3] = "WED";
        weekday[4] = "THU";
        weekday[5] = "FRI";
        weekday[6] = "STA";
        let list = [{
          title: 'SUN'
        }, {
          title: 'MON'
        }, {
          title: 'TUE'
        }, {
          title: 'WED'
        }, {
          title: 'THU'
        }, {
          title: 'FRI'
        }, {
          title: 'STA'
        }];
        for (let i = 0; i < list.length; i++) {
          let j = date.getDay() - 3 + i
          if (j < 0) j += 7;
          if (j > 6) j -= 7;
          list[i].title = weekday[j];
        }
        return list;
      },
    },
    onLoad() {},
    async mounted() {

      // const result = await Axios.get('http://localhost:8088/v1/frontend/news/detail', {
      //         }).then(res =>{
      //           console.log(res.data.data)
      //           for (var i = 0; i < res.data.data.newses.length; i++) {
      //             let tmp = res.data.data.newses[i];
      //             this.indexList.push({
      //               id: tmp.ID,
      //               title: tmp.Title,
      //               subTitle: tmp.Author,
      //               publishTime: uni.$u.timeFormat(tmp.PublishTime, 'yyyy年mm月dd日'),
      //               pageImage: "/static/building.png",
      //               content: tmp.Content
      //             })
      //           }
      //         }).catch(error =>{
      //           console.log(error);
      // 		  console.log("失败")
      //         })

      // 基本用法，注意：get请求的参数以及配置项都在第二个参数中
      uni.$u.http.get('/v1/frontend/news/detail', {

      }).then(res => {
        this.isArticleLoading = false;
        console.log(res.data.data)
        for (var i = 0; i < res.data.data.newses.length; i++) {
          let tmp = res.data.data.newses[i];
          this.indexList.push({
            id: tmp.ID,
            title: tmp.Title,
            subTitle: tmp.Author,
            publishTime: uni.$u.timeFormat(tmp.PublishTime, 'yyyy年mm月dd日'),
            pageImage: "/static/building.png",
            content: tmp.Content
          })
        }
      }).catch(err => {
        console.log("出错了...")
      })
    },
    async onReachBottom() {
      console.log("整个页面到底了！");
    },
    methods: {
      // 页面跳转
      gotoPage(url) {
        console.log("跳转的页面是" + url);
        uni.navigateTo({
          url
        })
      },
      // 详情跳转
      gotoDetail(u, id) {
        console.log("跳转的详情页面是" + u);
        uni.navigateTo({
          url: u + "?id=" + id,
        })
      },
      // 滚动到底部触发事件
      scrolltolower() {
        console.log("页面到底了")
      },
      onClick(e) {
        console.log(e)
      },
      actionsClick(text) {
        uni.showToast({
          title: text,
          icon: 'none'
        })
      }
    }
  }
</script>

<style>
  /* 整体内容样式 */
  .content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background-image: linear-gradient(178deg, #bac8ff, #c3fae855);
  }

  .text-area {
    display: flex;
    justify-content: center;
  }

  /* 倒计时样式 */
  .header-box {
    margin: 20rpx 0rpx 10rpx 0rpx;
    border-radius: 20rpx;
    width: 672rpx;
    opacity: 0.8;
    box-shadow: 5rpx 5rpx 5rpx #646464;
    /* 背景图片 */
    /* background: url('../../static/FZU_building.jpg') no-repeat; */
    /* background-size: 100%; */
    /* background-attachment: fixed; */
    background-image: linear-gradient(135deg, #614385, #516395);
  }

  /* 时间轴 */
  .timeline {
    margin: 30rpx 0rpx;
    color: #ffffff;
  }

  /* 标签栏样式 */
  .tabs-box {
    margin: 2rpx 2rpx;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  /* 卡片按钮 */
  .tabs-box-item {
    opacity: 0.7;
    border-radius: 32rpx;
    background-image: linear-gradient(90deg, #614385ee, #516395ee);
  }

  /* 最新动态样式 */
  .trends-box {
    border-radius: 20rpx;
    margin: 6rpx 0rpx;
    width: 646rpx;
  }

  .trends-box-title {
    /*    margin: 6rpx 16rpx; */
    margin-left: 16rpx;
    margin-bottom: 20rpx;
    font-size: 34rpx;
    color: #3F3F3F;
    font-weight: 700;
    /* 	font-family: "思源黑体"; */
  }

  .trends-box-item {
    opacity: 0.8;
    width: 680rpx;
  }

  .downloading {
    width: 100%;
    margin-top: 16px;
  }

  .countDown {
    width: 420rpx;
    font-family: "思源黑体";
    font-weight: 600;
  }

  .yearDown {
    width: 320rpx;
    font-family: "思源黑体";
    font-size: 32rpx;
    padding-left: 24rpx;
  }

  /*  .remark {
    margin-top: 50rpx;
    margin-left: 25rpx;
    width: 500rpx;
    color: #f5f5f5;
    font-size: 28rpx;
    font-family: "思源黑体";
  } */

  .viewUser {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 15rpx;
  }

  .viewSaid {
    height: auto;
    width: 600rpx;

    /* 圆角 */
    border-radius: 20rpx;

    /* 边 */
    border: 1rpx solid #E0E3DA;
    /* 阴影 */
    box-shadow: 6rpx 6rpx 8rpx #ebebeb;

    background-color: #ffffff;
    margin-left: 4rpx;
    margin-top: 24rpx;

    /* padding使得文字和图片不至于贴着边框 */
    padding: 16rpx;

  }

  .headPortrait {
    height: 100rpx;
    width: 100rpx;
    border-radius: 50%;
  }

  .userMes {
    margin-left: 16rpx;
    display: flex;
    flex-direction: column;
    width: 270rpx;
  }

  .userName {
    font-size: 30rpx;
    font-family: "黑体";
  }

  .saidContent {
    display: flex;
    flex-direction: row;
    align-items: center;
  }

  .textContent {
    width: 450rpx;
    margin-top: 0rpx;
    margin-right: 20rpx;
    font-size: 20rpx;
  }

  .sights {
    float: right;
    width: 190rpx;
    height: 125rpx;
    border-radius: 18rpx;
  }

  .publishTime {
    font-size: 20rpx;
    color: #9A9A9A;
    margin-left: 32rpx;
  }

  .title {
    font-size: 36rpx;
    font-weight: 600;
    margin-left: 10rpx;
  }
</style>