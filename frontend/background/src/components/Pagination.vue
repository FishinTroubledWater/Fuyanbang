<template>
<div style="padding: 10px 0">
  <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage1"
    :page-sizes="[5,10,15,20]"
    :page-size="10"
    layout="total,sizes,prev,pager,next,jumper"
    :total="20" class="mg whiteback pd">
  </el-pagination>
</div>
</template>

<script>
export default {
  name: "Pagination",
  methods: {
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    }
  },
  // props:{
  //   total: Number,
  //   queryInfo: Object
  // },
  data() {
    return {
      currentPage1: 5,
      currentPage2: 5,
      currentPage3: 5,
      currentPage4: 4
    };
  }
}
</script>

<style scoped>

</style>