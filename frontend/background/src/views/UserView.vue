<template>
  <el-table
      :data="tableData"
      style="width: 100%">
    <el-table-column
        label="联系方式"
        width="180">
      <template slot-scope="scope">
        <i class="el-icon-user"></i>
        <span style="margin-left: 10px">{{ scope.row.phoneNumber }}</span>
      </template>
    </el-table-column>
    <el-table-column
        label="姓名"
        width="180">
      <template slot-scope="scope">
        <el-popover trigger="hover" placement="top">
          <p>姓名: {{ scope.row.name }}</p>
          <p>住址: {{ scope.row.address }}</p>
          <div slot="reference" class="name-wrapper">
            <el-tag size="medium">{{ scope.row.name }}</el-tag>
          </div>
        </el-popover>
      </template>
    </el-table-column>
    <el-table-column
        label="用户院校"
        width="180">
      <template slot-scope="scope">
        <span style="margin-left: 10px">{{ scope.row.college }}</span>
      </template>
    </el-table-column>
    <el-table-column
        label="用户类别"
        width="180">
      <template slot-scope="scope">
        <span style="margin-left: 10px">{{ scope.row.role }}</span>
      </template>
    </el-table-column>

    <el-table-column label="操作">
      <template slot-scope="scope">
        <el-button
            size="mini"
            @click="handleEdit(scope.$index, scope.row)">编辑
        </el-button>
        <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)">删除
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "UserView",
  data() {
    const item = {
      date: '2016-05-02',
      name: '王小虎',
      phoneNumber: '110',
      role: '普通用户',
      college: '福州大学',
    };
    return {
      tableData: Array(10).fill(item)
    }
  },
  methods: {
    handleEdit(index, row) {
      console.log(index, row);
    },
    handleDelete(index, row) {
      console.log(index, row);
    }
  }


}
</script>

<style scoped>

</style>