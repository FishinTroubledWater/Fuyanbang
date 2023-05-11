<template>
  <div>
    <el-card class="round15 mg_b20">
      <el-row :gutter="20">
        <el-col :span="2" class="setcenter">
          <el-avatar :size="60" src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
                     :fit="'cover'"
                     style=""></el-avatar>
<!--          <span>预留</span>-->
        </el-col>
        <el-col :span="5">
          <div class="adminMes">
            <div class="account">admin</div>
            <div class="phoneNumber">电话：110</div>
            <div class="state">账号状态：<span class="startUsing">启用</span></div>
          </div>
        </el-col>
        <el-col :span="7" class="addNumMes">
          <div class="addNum">
            <div class="addUserNum">昨日新增用户<span class="numColor">9</span>位</div>
            <div class="addPostNum">昨日新增文章<span class="numColor">5</span>篇</div>
            <div class="reviewedNum">当前待审核文章<span class="numColor">7</span>篇</div>
          </div>
        </el-col>
        <el-col :span="3" class="numShow" >
          <el-avatar :size="'large'" :src="require('../assets/images/文章.svg')"
                     :fit="'cover'"
                     style=""></el-avatar>
          <div class="showNum">文章数量</div>
          <div class="showNumColor">333</div>
        </el-col>
        <el-col :span="3" class="numShow">
          <el-avatar :size="'large'" :src="require('../assets/images/用户.svg')"
                     :fit="'cover'"
                     style=""></el-avatar>
          <div class="showNum">用户数量</div>
          <div class="showNumColor">333</div>
        </el-col>
        <el-col :span="3" class="numShow">
          <el-avatar :size="'large'" :src="require('../assets/images/反馈.svg')"
                     :fit="'cover'"
                     style=""></el-avatar>
          <div class="showNum">反馈数量</div>
          <div class="showNumColor">333</div>
        </el-col>

      </el-row>


    </el-card>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-card class="round15">
          <span id="chart1" style="width: 500px;height: 400px;"></span>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card class="round15">
          <span id="chart2" style="width: 500px;height: 400px;"></span>
        </el-card>
      </el-col>

    </el-row>

  </div>
</template>

<script>
import * as echarts from 'echarts';

export default {
  name: "DashBoard",
  data(){
      return {}
  },
  created() {},
  //等地页面上元素渲染完毕
  mounted() {
    //文章发布统计分析
    var chart1 = echarts.init(document.getElementById('chart1'));
    var option1 = {
      title: {
        text: '文章发布统计分析',
      },
      color:[
        '#4394c5'
      ],
          tooltip: {},
      xAxis: {
        data: ['1月', '2月', '3月', '4月', '5月', '6月']
      },
      yAxis: {},
      series: [
        {
          name: '数据',
          type: 'bar',
          data: [5, 20, 36, 10, 10, 20],
        }
      ]
    }

    //全天文章发布分析
    var chart2 = echarts.init(document.getElementById('chart2'));
    var option2 = {
      title: {
        text: '全天文章发布分析',
      },
      color:[
        '#4394c5'
      ],
      tooltip: {},
      xAxis: {
        data: ['0.00', '5.00', '10.00', '15.00', '20.00', '24.00']
      },
      yAxis: {},
      series: [
        {
          name: '数据',
          type: 'line',
          smooth:'true',
          data: [1500, 500, 1000, 1500, 2000, 2500],
        }
      ]
    }
    // 绘制图表
    chart1.setOption(option1);
    chart2.setOption(option2);
  },


  methods:{}
}
</script>

<style scoped>
.setcenter{
  height: 70px;
}
.adminMes{
  /*margin-left: 5px;*/
}
.account{
  font-weight: 700;
  font-size: 16px;
  margin-bottom: 2px;
}
.phoneNumber{
  font-size: 8px;
  margin-bottom: 2px;
}
.state{
  font-size: 8px;
  margin-bottom: 2px;
}
.startUsing{
  /*border-radius: 10px;*/
  /*background-color: #F0F2F7;*/
  color: #3952FD;
}
.addUserNum {
  font-size: 6px;
  margin-top: 6px;
  font-family: 黑体,sans-serif;
}
.addPostNum{
  font-size: 6px;
  margin-top: 6px;
  font-family: 黑体,sans-serif;
}
.reviewedNum{
  font-size: 6px;
  margin-top: 6px;
  font-family: 黑体,sans-serif;
}
.numColor{
  margin-left: 2px;
  margin-right: 2px;
  color: #F57D2D;
}
.addNumMes{
  border-right:1px dashed #BBBBBB;
}
.numShow{
  text-align: center;
}
.showNum{
  font-size: 6px;
}
.showNumColor{
  color: #3952FD;
}
</style>
