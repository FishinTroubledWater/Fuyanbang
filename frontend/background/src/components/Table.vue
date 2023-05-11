<template>
  <el-table :data="tableData" style="width: 100%" border stripe>
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
    <el-table-column label="状态" v-if="showState" width="100" formatter="dateFormat">
      <template #default="scope">
        <slot name="state"
              :index="scope.$index"
              :row="scope.row"></slot>
      </template>
    </el-table-column>
    <el-table-column label="操作">
      <template #default="scope">
        <slot name="operation"
              :index="scope.$index"
              :row="scope.row">
        </slot>
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
  },
  methods:{

  }

}
</script>

<style scoped>

</style>