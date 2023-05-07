<template>
  <div>
    <el-card class="round15 mg_b20">
      <el-row :gutter="20">
        <el-col :span="4" class="setcenter">
          <el-avatar :size="'large'" src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
                     :fit="'cover'"
                     style="margin-right: 10px"></el-avatar>
          <span>admin</span>
        </el-col>
        <el-col :span="4" >

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

</style>