<template>
  <view class="content">

    <!-- 倒计时 -->
    <uni-section class="header-box" title="" sub-title="" type="">

      <u-row customStyle="margin: 2px 14px">
        <u-col span="4">
          <u--text :text="RemainingTime" size="90rpx" color="#ffffff" class="countDown"></u--text>
        </u-col>
      </u-row>

      <u-row customStyle="margin: 2px">
        <u-col span="3">
          <u--text></u--text>
        </u-col>
        <u-col span="4">
          <u--text></u--text>
        </u-col>
        <u-col span="4">
          <u--text :text="ExamYear" size="34rpx" color="#ffffff" class="yearDown"></u--text>
        </u-col>
      </u-row>

      <!--      <u-row style="margin: 16rpx 8rpx; padding: 0rpx 12rpx;">
        <u--text :line="2" text="满怀激情地投入到对事理的探究中来，至于建筑变成了自然而然的结果" size="30rpx" lineHeight="30rpx" color="#f1f1f1"></u--text>
      </u-row> -->
      <view class="remark">
        满怀激情地投入到对事理的探究中来，至于建筑变成了自然而然的结果。
      </view>

      <!-- 时间轴 -->
      <view class="timeline">
        <uni-steps :options="TimeList" active-icon="flag" :active="active" />
      </view>

    </uni-section>

    <!-- 标签按钮 -->
    <my-card :is-shadow="true" style="opacity: 0.6;width: 620rpx;">
      <view class="tabs-box">
        <view class="tabs-box-item" style="margin-right: 22px;">
          <u-icon label="院校" labelPos="bottom" labelSize="34rpx" labelColor="#000" margin="2" size="40"
            name="/static/home-images/academy.png" @click="gotoPage('/pages/home/university/academy')"></u-icon>
        </view>

        <view class="tabs-box-item" style="margin: 0px 22px;">
          <u-icon label="专业" labelPos="bottom" labelSize="34rpx" labelColor="#000" margin="2" size="40"
            name="/static/home-images/major.png" @click="gotoPage('/pages/home/professional/major')"></u-icon>
        </view>

        <view class="tabs-box-item" style="margin-left: 22px;">
          <u-icon label="资讯" labelPos="bottom" labelSize="34rpx" labelColor="#000" margin="2" size="40"
            name="/static/home-images/news.png" @click="gotoPage('/pages/home/information/news')"></u-icon>
        </view>
      </view>
    </my-card>

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
    background-color: #DCE2F1;
  }

  .text-area {
    display: flex;
    justify-content: center;
  }

  /* 倒计时样式 */
  .header-box {
    margin: 10rpx 0rpx;
    border-radius: 10rpx;
    width: 690rpx;
    opacity: 0.8;
    /* 背景图片 */
    background: url('../../static/FZU_building.jpg') no-repeat;
    background-size: 100%;
    background-attachment: fixed;
  }

  /* 时间轴 */
  .timeline {
    margin: 34rpx 0rpx;
  }

  /* 标签栏样式 */
  .tabs-box {
    margin: 8rpx 8rpx;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  /* 标签栏子元素 */
  .tabs-box-item {
    margin: 0 20rpx;
  }

  /* 最新动态样式 */
  .trends-box {
    border-radius: 15rpx;
    margin: 6rpx 0rpx;
  }

  .trends-box-title {
    /*    margin: 6rpx 16rpx; */
    margin-left: 40rpx;
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

  .countDown {
    width: 300rpx;
    font-family: "思源黑体";
    font-weight: 700;
  }

  .yearDown {
    width: 300rpx;
    font-family: "思源黑体";
    font-size: 32rpx;
  }

  .remark {
    margin-top: 50rpx;
    margin-left: 25rpx;
    width: 500rpx;
    color: #f5f5f5;
    font-size: 28rpx;
    font-family: "思源黑体";
  }




  .viewUser {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 15rpx;
  }

  .viewSaid {
    height: auto;

    /* 圆角 */
    border-radius: 18rpx;

    /* 边 */
    border: 1rpx solid #E0E3DA;
    /* 阴影 */
    box-shadow: 2rpx 7rpx 0rpx #ebebeb;

    background-color: #ffffff;
    margin-left: 30rpx;
    margin-right: 30rpx;
    margin-top: 25rpx;

    /* padding使得文字和图片不至于贴着边框 */
    padding: 25rpx;

  }

  .headPortrait {
    height: 100rpx;
    width: 100rpx;
    border-radius: 50%;
  }

  .userMes {
    margin-left: 30rpx;
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
    margin-left: 40rpx;
  }

  .title {
    font-size: 36rpx;
    font-weight: 600;
    margin-left: 10rpx;
  }
</style>