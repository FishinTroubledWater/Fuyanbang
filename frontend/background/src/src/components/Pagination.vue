<template>
<div style="padding: 10px 0">
  <el-pagination
      background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page.sync="queryInfo.pageNum"
    :page-sizes="[5,10,15,20]"
    :page-size="queryInfo.pageSize"
    layout="total,sizes,prev,pager,next,jumper"
    :total="total" class="mg whiteback pd">
  </el-pagination>
</div>
</template>

<script>
export default {
  name: "Pagination",
  methods: {
    handleSizeChange(newSize) {
      this.$emit('page-size-change', newSize);
    },
    handleCurrentChange(newPage) {
      this.$emit('page-change', newPage);
    },
  },
  props:{
    total: Number,
    queryInfo: Object
  },
  data() {
    return {
    };
  }
}
</script>

<style scoped>

</style>