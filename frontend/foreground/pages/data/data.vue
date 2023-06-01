<template>
  <view>

    <view class="tabs-box">
      <view class="one-nav" style="padding: 20rpx;font-size: 33rpx;"
        :class="currentSwiperIndex === 0 ? 'nav-actived' : '' " @tap="swiperChange(0)">加油站
      </view>
      <view class="one-nav" style="padding: 20rpx;font-size: 33rpx;"
        :class="currentSwiperIndex === 1 ? 'nav-actived' : '' " @tap="swiperChange(1)">求解答
      </view>
    </view>
    <swiper class="swiper-box" style="height: 2000upx" :current="currentSwiperIndex">
      <!-- 加油站 -->
      <swiper-item class="swiper-item sns-que">
        <view class="topic" @click="gotoPage('/pages/data/activityDetail/activityDetail')">
          <text class="contentTopWord">攒图计划招募令</text>
          <text>\n</text>
          <text style="color:#efefef;margin-left: 20rpx;font-size: 24rpx;">每日一画，不画就出局\n</text>
         <!-- <navigator class="contentTopWordDetails" url="/pages/data/activityDetail/activityDetail">了解详情</navigator> -->
        </view>
    
        <u-list>
          <u-list-item v-for="(item, index) in indexList" :key="index">
            <uni-card @click="clicknews(item.postId)" :title="item.name" sub-title="帖子信息" :extra="item.time"
              :thumbnail="item.icon" class="trends-box-item">
              <u-text :lines="3" :text="item.summary"></u-text>
              <!-- <image class="newsimage" :src="indexList[index].img[0]"></image> -->
              <!-- <view class="u-content">
                <u-parse :content="indexList[index].summary"></u-parse>
              </view> -->
            </uni-card>
          </u-list-item>
        </u-list>
      </swiper-item>
    
      <!-- 求解答 -->
      <swiper-item class="swiper-item sns-oil">
        <view class="topic" @click="gotoPage('/pages/data/activityDetail/activityDetail')">
          <text class="contentTopWord">21天攒图计划招募令</text>
          <text>\n</text>
          <text style="color:#FFFFFF;margin-left: 20rpx;">每日一画，不画就出局\n</text>
      <!--    <navigator class="contentTopWordDetails" url="/pages/data/activityDetail/activityDetail">了解详情</navigator> -->
        </view>
        <u-list>
          <u-list-item v-for="(item, index) in questionsList" :key="index">
            <uni-card @click="clickquestions(item.queId)" :title="item.title" :sub-title="item.time" :extra="item.name"
              :thumbnail="item.icon" class="trends-box-item">
              <u-text :lines="3" :text="item.summary"></u-text>
              <u-row customstyle="margin-bottom: 10px">
                <u-col span="6">
                  <text style="">悬赏学币：</text>
                </u-col>
                <u-col span="6" offset="-4">
                  <u-text lines="1" :text="item.Reward"></u-text>
                </u-col>
              </u-row>
            </uni-card>
          </u-list-item>
        </u-list>
      </swiper-item>
    </swiper>

  </view>
</template>

<script>
  import uCard from '../../uni_modules/uni-card/uni-card.vue'
  import uSteps from '../../uni_modules/uni-steps/uni-steps.vue'
  import uniIcons from '../../uni_modules/uni-icons/uni-icons.vue'
  import uSection from '../../uni_modules/uni-section/uni-section.vue'
  
  export default {
    components: {
      uCard,
      uSteps,
      uniIcons: uniIcons,
      uSection
    },
    data() {
      return {
        id: '',
        content: `
					<p>露从今夜白，月是故乡明</p>
					<img src="../../static/background/activityDetails.png" />
				`,
        currentSwiperIndex: 0,
        list1: [{
          name: '求解答',
        }, {
          name: '加油站',
        }],

        indexList: [],

        // indexList:  [{
        // 		name: '1',
        // 		PublishTime:'2022-12-21',
        // 		icon:'',
        // 		PartID:'123456',
        // 		Summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
        // 		isImage:true,
        // 		img:'www.baidu.com',
        // 	},{
        // 		name: '1',
        // 		PublishTime:'2022-12-21',
        // 		icon:'',
        // 		PartID:'123456',
        // 		Summary:'近日，教育部部署2023年全国硕士研究生招生复试录取工作...',
        // 		isImage:true,
        // 		img:'www.baidu.com',
        // }],
        questionsList: []
        // questionsList: [{
        // 		name: '1',
        // 		time: '2022-12-21',
        // 		icon: '../../static/background/activityDetails.png',
        // 		queId: '123456',
        // 		summary: '第一条',
        // 		isImage: true,
        // 		img: ['../../static/background/activityDetails.png',
        // 			'../../static/background/activityDetails.png',
        // 			'../../static/background/activityDetails.png'
        // 		]
        // 	},
        // 	{
        // 		name: '2',
        // 		time: '2022-12-21',
        // 		icon: '../../static/background/activityDetails.png',
        // 		queId: '123',
        // 		summary: '第二条',
        // 		isImage: true,
        // 		img: ['../../static/background/activityDetails.png',
        // 			'../../static/background/activityDetails.png',
        // 			'../../static/background/activityDetails.png'
        // 		]
        // 	}
        // ]
      };
    },
    methods: {
      gotoPage(url){
        uni.navigateTo({
          url
        })
      },
      //求解答、加油站、关注切换方法
      swiperChange(index) {
        this.currentSwiperIndex = index
      },
      clicknews(index) {
        // console.log('1'),
        uni.navigateTo({
          url: '/pages/data/newsDetails/newsDetails?id=' + index,
        })
      },
      clickquestions(index) {
        uni.navigateTo({
          url: '/pages/data/questionsDetails/questionsDetails?id=' + index,
        })
      }

    },
    onShow() {
      uni.getStorage({
        key: 'userId', // 储存在本地的变量名
        success: res => {
          // 成功后的回调
          // console.log(res.data);   // hello  这里可做赋值的操作
          this.id = res.data;
          console.log(this.id)
        }
      })
      uni.$u.http.get('/v1/frontend/circle/newinfo/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.indexList = res.data.data
        }).catch(err => {

        }),
        uni.$u.http.get('/v1/frontend/circle/newque/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.questionsList = res.data.data
        }).catch(err => {

        })
    },
    mounted() {
      uni.getStorage({
        key: 'userId', // 储存在本地的变量名
        success: res => {
          // 成功后的回调
          // console.log(res.data);   // hello  这里可做赋值的操作
          this.id = res.data;
          console.log(this.id)
          uni.$u.http.get('/v1/frontend/circle/newinfo/' + this.id, {

            }).then(res => {
              console.log(res.data.data);
              this.indexList = res.data.data
            }).catch(err => {

            }),
            uni.$u.http.get('/v1/frontend/circle/newque/' + this.id, {

            }).then(res => {
              console.log(res.data.data);
              this.questionsList = res.data.data
            }).catch(err => {

            })
        }
      })
      uni.$u.http.get('/v1/frontend/circle/newinfo/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.indexList = res.data.data
        }).catch(err => {

        }),
        uni.$u.http.get('/v1/frontend/circle/newque/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.questionsList = res.data.data
        }).catch(err => {

        })
    },

  }
</script>

<style lang="scss">
  .trends-box-item {

    /* 圆角 */
    border-radius: 18rpx;

    /* 边 */
    border: 1rpx solid #E0E3DA;
    /* 阴影 */
    box-shadow: 2rpx 7rpx 0rpx #ebebeb;

  }

  .u-content {
    padding: 24rpx;
  }

  .newsimage {
    max-height: 200rpx;
    max-width: 200px;
  }

  .quanzi {
    font-size: 50rpx;
    font-weight: 800;
  }

  .tabs-box {
    margin-top: 60rpx;
    flex-flow: row;
    justify-content: space-around;
    display: flex;
  }

  .one-nav {
    color: #606266;
  }

  .nav-actived {
    color: #0b7285;
    font-weight: 700;
    //background-color: #74759b;

  }

  .appContentTop {
    z-index: -1;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 200rpx;
  }

  .contentTopWord {
    color: #b77b57;
    font-size: 50rpx;
    margin-left: 10rpx;
    margin-top: 24rpx;
    font-family: Arial;
    font-weight: bold;
  }

  .contentTopWordDetails {
    color: #74759b;
    font-size: 40rpx;
    margin-left: 20rpx;
    height: 80rpx;
  }

  .contentPic {
    object-fit: cover;
    height: 100%;
    width: 100%;
  }

  .topic {
    background: url("../../static/background/bg2.png") top center;
    background-size: cover;
    height: 250rpx;
  }
</style>