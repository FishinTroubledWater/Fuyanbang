<template>
    <el-table :data="tableData"  style="width: 100%" border stripe>
    <!--        索引-->
    <el-table-column type="index"></el-table-column>
    <el-table-column
        v-for="column in columns"
        :key="column.prop"
        :prop="column.prop"
        :label="column.label"
        :width="column.width"
        :formatter="column.formatter"
        :sortable="column.sortable"
        :sort-by="column.index"
        :show-overflow-tooltip="column.showOverflowTooltip">
    </el-table-column>
      <el-table-column v-if="showState" label="状态" width="120px">
        <template>
<!--          这里该怎么写-->
          <slot name="state"></slot>
        </template>
      </el-table-column>
      <el-table-column  label="操作" >
        <template #>
         <slot name="op"></slot>
        </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "Table",
  props: {
    tableData: {
      type: Array,
      default: () => []
    },
    columns: {
      type: Array,
      default: () => []
    },
    showState: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>

</style>