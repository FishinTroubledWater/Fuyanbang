<template>
  <el-drawer
      :title="title"
      :visible.sync="drawer"
      :direction="direction"
      :before-close="handleClose"
      size="40%">
    <template #default="scope">
      <slot name="details"
            :index="scope.$index"
            :row="scope.row"></slot>
    </template>
  </el-drawer>
</template>

<script>
export default {
  name: "Drawer",
  props: {
    drawer: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
    }
  },
  data() {
    return {
      direction: 'rtl',
    }
  },
  methods: {
    handleClose() {
      this.$emit('closed',false)
    }
  }
}
</script>

<style scoped>

</style>