<template>
  <view>
    <uni-card :title="indexList.name" sub-title="帖子详情" :extra="indexList.time" :thumbnail="indexList.icon"
      class="trends-box-item" style="border-radius: 24rpx;margin-bottom: 8rpx;box-shadow: 12rpx 12rpx 10rpx #bfbfbf;">
      <view class="u-content">
        <u-parse :content="indexList.summary"></u-parse>
      </view>

      <u-row customstyle="margin: 5rpx 0 5rpx 2rpx;">
        <u-col span="6">
          <u-icon v-if="whetherCollect==='false'" style="padding-left: 0rpx;padding-top: 15rpx;" label="收藏"
            color="#808A87" size="20" name="star" @click="clickCollect"></u-icon>
          <u-icon v-if="whetherCollect==='true'" style="padding-left: 0rpx;padding-top: 15rpx;" label="收藏"
            color="#fed71a" size="20" name="star-fill" @click="clickCollect"></u-icon>
        </u-col>
        <u-col span="6" offset="-4">
          <u-icon v-if="whetherLike==='false'" style="padding-left: 20rpx;padding-top: 15rpx;" :label="this.likeNum"
            color="#808A87" size="20" name="heart" @click="clickLike"></u-icon>
          <u-icon v-if="whetherLike==='true'" style="padding-left: 20rpx;padding-top: 15rpx;" :label="this.likeNum"
            color="#FF0000" size="20" name="heart-fill" @click="clickLike"></u-icon>
        </u-col>
      </u-row>

    </uni-card>
    <!-- <view class="content">
			<textarea class="uni-title uni-common-pl" v-model="txt"></textarea>
		</view> -->

    <view class="textarea_box">
      <textarea class="textarea" placeholder="说说你的看法吧，在此处留下你的足迹。👣👣👣" placeholder-style="font-size:28rpx" maxlength="200"
        @input="descInput" v-model="desc" />
      <view class="num">{{ desc.length }}/200</view>
    </view>

    <button class="sendButton" @click="clickSent">发送评论</button>

    <text style="margin-left: 40rpx; font-size: 40rpx; font-weight: 800;">评论:</text>
    <uni-card v-for="(item, index) in comment" :title="item.name" :sub-title="item.time" :thumbnail="item.icon"
      class="trends-box-item" style="border-radius: 24rpx;margin-bottom: 8rpx;box-shadow: 12rpx 12rpx 10rpx #bfbfbf;">
      <u--text :text="item.content"></u--text>
    </uni-card>

    <u-gap height="20" bgColor="#fff"></u-gap>
  </view>
</template>

<script>
  import uCard from '../../../uni_modules/uni-card/uni-card.vue'
  import uSteps from '../../../uni_modules/uni-steps/uni-steps.vue'
  import uniIcons from '../../../uni_modules/uni-icons/uni-icons.vue'
  import uSection from '../../../uni_modules/uni-section/uni-section.vue'
import { onMounted } from "vue"
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
        postId: '0',
        whetherLike: 'false',
        whetherCollect: 'false',
        likeNum: '0',
        desc: '',
        myComment: 'null',

        txt: "txt",
        academyName: '福州大学',
        indexList: {
          // name: 'zhang',
          // time: '2022-12-21',
          // icon: '../../../static/background/activityDetails.png',
          // postId: '123456',
          // summary: `<p>露从今夜白</p>
          // <img src="../../static/background/activityDetails.png" />`,
          // // isImage: true,
          // // img: ['../../../static/background/activityDetails.png',
          // // 	'../../../static/background/bg1.png',
          // // 	'../../../static/background/bg2.png'
          // // ],
        },
        comment: [],

      }
    },
    watch: {
      txt(txt) {
        if (txt.indexOf('\n') != -1) { //敲了回车键了
          uni.hideKeyboard() //收起软键盘
        }
      }
    },
    methods: {
      clickCollect() {
        if (this.whetherCollect === 'true') {
          this.whetherCollect = 'false'
        } else {
          this.whetherCollect = 'true'
        }


        this.id = this.id + '',
          console.log(this.id),
          uni.$u.http.post('/v1/frontend/circle/postCollected', {
            articleID: this.postId,
            userID: this.id,
            isCollected: this.whetherCollect,
          }).then(res => {
            console.log(res.data)
            if (res.data.code == 200) {
              uni.showToast({
                title: "操作成功",
                duration: 1000,
              })
            } else {
              uni.showToast({
                title: "操作不成功",
                duration: 1000,
              })
              if (this.whetherCollect === 'true') {
                this.whetherCollect = 'false'
              } else {
                this.whetherCollect = 'true'
              }
            }

          }).catch(err => {

          })
      },
      clickLike() {
        if (this.whetherLike === 'true') {
          this.whetherLike = 'false'
          this.likeNum = this.likeNum - 1
        } else {
          this.whetherLike = 'true'
          this.likeNum = this.likeNum + 1
        }

        this.id = this.id + '',
          console.log(this.id),
          uni.$u.http.post('/v1/frontend/circle/postLiked', {
            postId: this.postId,
            userId: this.id,
            isLiked: this.whetherLike,
          }).then(res => {
            console.log(res.data)
            if (res.data.code == 200) {
              uni.showToast({
                title: "操作成功",
                duration: 1000,
              })
            } else {
              uni.showToast({
                title: "操作不成功",
                duration: 1000,
              })
              if (this.whetherLike === 'true') {
                this.whetherLike = 'false'
                this.likeNum = this.likeNum - 1
              } else {
                this.whetherLike = 'true'
                this.likeNum = this.likeNum + 1
              }
            }
            // setTimeout(() => {
            // 	this.$router.go(0)
            // }, 500)

          }).catch(err => {

          })

      },
      descInput(e) {
        console.log(e.detail.value.length, '输入的字数')
        this.myComment = e.detail.value
      },
      refresh(){
        uni.getStorage({
          key: 'userId', // 储存在本地的变量名
          success: res => {
            // 成功后的回调
            // console.log(res.data);   // hello  这里可做赋值的操作
            this.id = res.data;
            console.log(this.id)
            uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.postId + '/' + this.id, {
        
            }).then(res => {
              console.log(res.data.data);
              this.indexList = res.data.data[0];
              this.whetherLike = this.indexList.isLiked;
              this.whetherCollect = this.indexList.isCollected;
              this.likeNum = this.indexList.likeNum;
            }).catch(err => {
        
            })
          }
        })
        uni.$u.http.get('/v1/frontend/circle/newinfoComment/' + this.postId, {
        
          }).then(res => {
            console.log(res.data.data);
            this.comment = res.data.data;
          }).catch(err => {
        
          }),
          uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.postId + '/' + this.id, {
        
          }).then(res => {
            console.log(res.data.data);
            this.indexList = res.data.data[0];
            this.whetherLike = this.indexList.isLiked;
            this.whetherCollect = this.indexList.isCollected;
            this.likeNum = this.indexList.likeNum;
          }).catch(err => {
        
          })
      },
      clickSent() {
        var _this=this;
        console.log(this.myComment)
        uni.getStorage({
          key: 'userId', // 储存在本地的变量名
          success: res => {
            // 成功后的回调
            // console.log(res.data);   // hello  这里可做赋值的操作
            this.id = res.data;
            console.log(this.id)
          }
        })
        uni.$u.http.post('/v1/frontend/circle/postComment', {
          postId: this.postId,
          userId: this.id,
          comment: this.myComment,
        }).then(res => {
          console.log(res.data)
          if (res.data.code == 200) {
            uni.showToast({
              title: "评论成功",
              duration: 1000,
            })
          }
          setTimeout(() => {
            _this.refresh();
            _this.desc='';
          }, 500)

        }).catch(err => {

        })

      },

    },
onShow(){
  mounted();
},
    onLoad: function(option) {
      console.log(option.id)
      this.postId = option.id
    },

    mounted() {
      uni.getStorage({
        key: 'userId', // 储存在本地的变量名
        success: res => {
          // 成功后的回调
          // console.log(res.data);   // hello  这里可做赋值的操作
          this.id = res.data;
          console.log(this.id)
          uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.postId + '/' + this.id, {

          }).then(res => {
            console.log(res.data.data);
            this.indexList = res.data.data[0];
            this.whetherLike = this.indexList.isLiked;
            this.whetherCollect = this.indexList.isCollected;
            this.likeNum = this.indexList.likeNum;
          }).catch(err => {

          })
        }
      })
      uni.$u.http.get('/v1/frontend/circle/newinfoComment/' + this.postId, {

        }).then(res => {
          console.log(res.data.data);
          this.comment = res.data.data;
        }).catch(err => {

        }),
        uni.$u.http.get('/v1/frontend/circle/newinfoDetails/' + this.postId + '/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.indexList = res.data.data[0];
          this.whetherLike = this.indexList.isLiked;
          this.whetherCollect = this.indexList.isCollected;
          this.likeNum = this.indexList.likeNum;
        }).catch(err => {

        })

      // uni.$u.http.post('/v1/frontend/academy/searchByRule', {
      // 	region: '福建',
      // 	level: '985',
      // 	type: '法学',
      // }).then(res => {
      // 	console.log(res.data)
      // }).catch(err => {

      // })
    },
  }
</script>

<style>
  .textarea_box {
    margin: 48rpx 40rpx 32rpx 40rpx;
    padding: 20rpx;
    background-color: #F2F2F2;
    border-radius: 20rpx;
    box-shadow: -5rpx -8rpx 5rpx #bfbfbf;
    height: 250rpx;

    /deep/ .uni-textarea-textarea {
      font-size: 28rpx;
      line-height: 42rpx;
    }

    .num {
      text-align: right;
      color: gray
    }
  }

  .sendButton {
    width: 45%;
    height: 40px;
    background-image: linear-gradient(112deg, #08507880, #85d8ce);
    border-radius: 20px;
    margin-bottom: 50rpx;
    color: #fff;
    font-size: 32rpx;
    border: none;
    box-shadow: 5rpx 10rpx 5rpx #bfbfbf;
  }

  .textarea {
    width: 100%;
    height: 80%;
  }
</style>