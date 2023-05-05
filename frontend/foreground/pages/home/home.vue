<template>
  <view class="content">

    <!-- 倒计时 -->
    <uni-section class="header-box" title="" sub-title="" type="">

      <u-row customStyle="margin: 2px 8px">
        <u-col span="4">
          <u--text text="260天" size="40px" color="#ffffff"></u--text>
        </u-col>
      </u-row>

      <u-row customStyle="margin: 2px">
        <u-col span="4">
          <u--text></u--text>
        </u-col>
        <u-col span="4">
          <u--text></u--text>
        </u-col>
        <u-col span="4">
          <u--text text="2024考研倒计时" size="16px"></u--text>
        </u-col>
      </u-row>

      <u-row style="margin: 16rpx 8rpx; padding: 0rpx 8rpx;">
        <u--text :line="2" text="满怀激情地投入到对事理的探究中来，至于建筑变成了自然而然的结果" size="14px"></u--text>
      </u-row>

      <!-- 时间轴 -->
      <view class="timeline">
        <uni-steps :options="list" active-icon="flag" :active="active" />
      </view>

    </uni-section>

    <!-- 标签按钮 -->
    <uni-card :is-shadow="true" style="opacity: 0.6;width: 86%;">
      <view class="tabs-box">
        <view class="tabs-box-item" style="margin-right: 22px;">
          <u-icon label="院校" labelPos="bottom" labelSize="20px" labelColor="#000" margin="2" size="40"
            name="/static/home-images/academy.png" @click="gotoPage('/pages/home/academy')"></u-icon>
        </view>

        <view class="tabs-box-item" style="margin: 0px 22px;">
          <u-icon label="专业" labelPos="bottom" labelSize="20px" labelColor="#000" margin="2" size="40"
            name="/static/home-images/major.png" @click="gotoPage('/pages/home/major')"></u-icon>
        </view>

        <view class="tabs-box-item" style="margin-left: 22px;">
          <u-icon label="资讯" labelPos="bottom" labelSize="20px" labelColor="#000" margin="2" size="40"
            name="/static/home-images/news.png" @click="gotoPage('/pages/home/news')"></u-icon>
        </view>
      </view>
    </uni-card>

    <!-- 动态 -->
    <view class="trends-box">
      <view class="trends-box-title">最新动态
      </view>
      <u-line color="#000000"></u-line>
      <u-list @scrolltolower="scrolltolower">
        <u-list-item>
          <uni-card :is-shadow="true" title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日"
            thumbnail="/static/building.png" class="trends-box-item" @click="gotoPage('/pages/home/detail')">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>

        <u-list-item>
          <uni-card :is-shadow="true" title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日"
            thumbnail="/static/building.png" class="trends-box-item">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>

        <u-list-item>
          <uni-card :is-shadow="true" title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日"
            thumbnail="/static/building.png" class="trends-box-item">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>

        <u-list-item>
          <uni-card title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日" thumbnail="/static/building.png"
            class="trends-box-item">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>

        <u-list-item>
          <uni-card title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日" thumbnail="/static/building.png"
            class="trends-box-item">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>

        <u-list-item>
          <uni-card title="快讯!2023年考研国家线发布" sub-title="教育部" extra="2023年03月10日" thumbnail="/static/building.png"
            class="trends-box-item">
            <u--text :lines="1" text="近日，教育部部署2023年全国硕士研究生招生复试录取工作..."></u--text>
          </uni-card>
        </u-list-item>
      </u-list>

    </view>

  </view>
</template>

<script>
  import uCard from '../../components/uni-card/uni-card.vue'
  import uSteps from '../../components/uni-steps/uni-steps.vue'
  import uIcons from '../../components/uni-icons/uni-icons.vue'
  import uSection from '../../components/uni-section/uni-section.vue'

  export default {
    components: {
      uCard,
      uSteps,
      uIcons,
      uSection
    },
    data() {
      return {
        // 时间轴
        active: 3,
        list: [{
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
        }],

        // 最新动态
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
    onLoad() {},
    methods: {
      // 页面跳转
      gotoPage(url) {
        uni.navigateTo({
          url
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
    opacity: 0.8;
    /* 背景图片 */
    background: url('../../static/FZU_building.jpg') no-repeat;
    background-size: 100%;
    background-attachment: fixed;
  }

  /* 时间轴 */
  .timeline {
    margin: 16rpx 0rpx;
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
    margin: 6rpx 16rpx;
    font-size: 24rpx;
    color: #454545;
  }

  .trends-box-item {
    opacity: 0.8;
  }
</style>