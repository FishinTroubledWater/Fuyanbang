<template>
  <view>
    <uni-card :title="indexList.name" sub-title="问题详情" :extra="indexList.time" :thumbnail="indexList.icon"
      class="trends-box-item" style="border-radius: 24rpx;margin-bottom: 8rpx;box-shadow: 12rpx 12rpx 10rpx #bfbfbf;">
      <view class="u-content">
        <u-parse :content="indexList.content"></u-parse>
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

    <view class="textarea_box">
      <textarea class="textarea" placeholder="写下你的回答吧，救救孩子吧。😭😭😭" placeholder-style="font-size:28rpx" maxlength="200"
        @input="descInput" v-model="desc" />
      <view class="num">{{ desc.length }}/200</view>
    </view>

    <button class="sendButton" @click="clickSent">发送回答</button>

    <text style="margin-left: 40rpx; font-size: 40rpx; font-weight: 800;">回答:</text>
    <uni-card v-for="(item, index) in answer" :title="item.name" :sub-title="item.time" :thumbnail="item.icon"
      class="trends-box-item" @click="clickanswer(item.answerId,item.isAccepted,queID)"
      style="border-radius: 24rpx;margin-bottom: 8rpx;box-shadow: 12rpx 12rpx 10rpx #bfbfbf;">
      <u--text lines="3" :text="item.answer"></u--text>
      <u-row customstyle="margin-bottom: 10px">
        <u-col span="3">
          <text>回答状态：</text>
        </u-col>
        <u-col span="3">
          <u--text lines="1" :text="item.isAccepted"></u--text>
        </u-col>
      </u-row>
    </uni-card>

    <u-gap height="20" bgColor="#fff"></u-gap>

  </view>
</template>

<script>
  export default {
    data() {
      return {
        id: '',
        queID: '',
        whetherLike: 'false',
        whetherCollect: 'false',
        isSolved: 'false',
        likeNum: '0',
        desc: '',
        myanswer: 'null',
        isMine: '',
        txt: "txt",
        academyName: '福州大学',
        indexList: {},
        answer: [],

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
      clickanswer(index, isAccepted, queID) {
        var _this = this
        if (this.isMine == true && isAccepted == "未采纳" && _this.isSolved == "false") {
          uni.showModal({
            title: '',
            confirmText: '采纳',
            cancelText: '取消',
            content: '是否采纳该回答？',
            success: function(res) {
              if (res.confirm) {
                console.log(this.queID)
                index = index + '',
                  uni.$u.http.post('/v1/frontend/circle/postAnswerStatus', {
                    answerId: index,
                    queId: queID,
                  }).then(res => {
                    console.log(res.data)
                    if (res.data.code == 200) {
                      uni.showToast({
                        title: "操作成功",
                        duration: 1000,
                      })
                      setTimeout(() => {
                       _this.refresh();
                      }, 500)
                    } else {
                      uni.showToast({
                        title: "操作不成功",
                        duration: 1000,
                      })

                    }

                  }).catch(err => {

                  })

              } else if (res.cancel) {
                console.log('用户点击取消');
              }

            }
          });
        } else if (_this.isSolved == "true") {
          // uni.showModal({
          //   title: '',
          //   confirmText: '确定',
          //   content: '该问题已被采纳',
          // })
        }

      },
      refresh() {
        uni.getStorage({
          key: 'userId', // 储存在本地的变量名
          success: res => {
            // 成功后的回调
            // console.log(res.data);   // hello  这里可做赋值的操作
            this.id = res.data;
            console.log(this.id)
            uni.$u.http.get('/v1/frontend/circle/queDetails/' + this.queID + '/' + this.id, {

            }).then(res => {
              console.log(res.data.data);
              this.indexList = res.data.data[0];
              this.whetherLike = this.indexList.isLiked;
              this.whetherCollect = this.indexList.isCollected;
              this.likeNum = this.indexList.likeNum;
              this.isSolved = this.indexList.isSolved;
            }).catch(err => {

            })
            uni.$u.http.get('/v1/frontend/circle/queAnswer/' + this.queID, {

            }).then(res => {
              console.log(res.data.data);
              this.answer = res.data.data;
            }).catch(err => {

            })
          }
        })
        uni.$u.http.get('/v1/frontend/circle/queAnswer/' + this.queID, {

        }).then(res => {
          console.log(res.data.data);
          this.answer = res.data.data;
        }).catch(err => {

        })
        uni.$u.http.get('/v1/frontend/circle/queDetails/' + this.queID + '/' + this.id, {

        }).then(res => {
          console.log(res.data.data);
          this.indexList = res.data.data[0];
          this.whetherLike = this.indexList.isLiked;
          this.whetherCollect = this.indexList.isCollected;
          this.likeNum = this.indexList.likeNum;
          this.isMine = this.indexList.isMine;
          this.isSolved = this.indexList.isSolved;
        }).catch(err => {

        })
      },
      clickCollect() {
        if (this.whetherCollect === 'true') {
          this.whetherCollect = 'false'
        } else {
          this.whetherCollect = 'true'
        }
        this.id = this.id + '',
          console.log(this.id),
          uni.$u.http.post('/v1/frontend/circle/postCollected', {
            articleID: this.queID,
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
            postId: this.queID,
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
        this.myanswer = e.detail.value
      },
      clickSent() {
        var _this = this;
        console.log(this.myanswer)
        //post请求
        uni.getStorage({
          key: 'userId', // 储存在本地的变量名
          success: res => {
            // 成功后的回调
            // console.log(res.data);   // hello  这里可做赋值的操作
            this.id = res.data;
            console.log(this.id)
          }
        })
        uni.$u.http.post('/v1/frontend/circle/postAnswer', {
          queId: this.queID,
          userId: this.id,
          answer: this.myanswer,
        }).then(res => {
          console.log(res.data)
          if (res.data.code == 200) {
            uni.showToast({
              title: "回答成功",
              duration: 1000,
            })
            setTimeout(() => {
              _this.refresh();
              _this.desc = '';
            }, 500)

          }
        }).catch(err => {

        })
      },
    },

    mounted() {
      uni.getStorage({
        key: 'userId', // 储存在本地的变量名
        success: res => {
          // 成功后的回调
          // console.log(res.data);   // hello  这里可做赋值的操作
          this.id = res.data;
          console.log(this.id)
          uni.$u.http.get('/v1/frontend/circle/queDetails/' + this.queID + '/' + this.id, {

          }).then(res => {
            console.log(res.data.data);
            this.indexList = res.data.data[0];
            this.whetherLike = this.indexList.isLiked;
            this.whetherCollect = this.indexList.isCollected;
            this.likeNum = this.indexList.likeNum;
            this.isSolved = this.indexList.isSolved;
          }).catch(err => {

          })
          uni.$u.http.get('/v1/frontend/circle/queAnswer/' + this.queID, {

          }).then(res => {
            console.log(res.data.data);
            this.answer = res.data.data;
          }).catch(err => {

          })
        }
      })
      uni.$u.http.get('/v1/frontend/circle/queAnswer/' + this.queID, {

      }).then(res => {
        console.log(res.data.data);
        this.answer = res.data.data;
      }).catch(err => {

      })
      uni.$u.http.get('/v1/frontend/circle/queDetails/' + this.queID + '/' + this.id, {

      }).then(res => {
        console.log(res.data.data);
        this.indexList = res.data.data[0];
        this.whetherLike = this.indexList.isLiked;
        this.whetherCollect = this.indexList.isCollected;
        this.likeNum = this.indexList.likeNum;
        this.isMine = this.indexList.isMine;
        this.isSolved = this.indexList.isSolved;
      }).catch(err => {

      })
    },

    onLoad: function(option) {
      console.log(option.id)
      this.queID = option.id
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
    margin-bottom: 40rpx;
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