<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    帖子查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 帖子查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入用户名" v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getPostList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加帖子
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    帖子列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 帖子列表</div>
      <Table :table-data="postList" :columns="columns" :show-state="true">
        <!--        状态区-->
        <template #state="scope">
          <el-tag v-if="scope.row.State === 0" type="warning">未审核</el-tag>
          <el-tag v-else type="success">已审核</el-tag>

        </template>
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="warning" icon="el-icon-finished" round
                     @click="checkPostById(scope.row.ID)"
                     :disabled="scope.row.State=== 1">审核
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.ID)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removePostById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    添加帖子的对话框-->
    <el-dialog title="添加帖子" :visible.sync="addDialogVisible" width="50%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="用户名" prop="account">
          <el-input v-model="addForm.account"></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="addForm.title"></el-input>
        </el-form-item>
        <el-form-item label="板块" prop="partID">
          <el-select v-model="addForm.partID" placeholder="请选择">
            <el-option
                v-for="item in parts"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="state">
          <el-radio v-model="addForm.state" label='0'>未审核</el-radio>
          <el-radio v-model="addForm.state" label='1'>已审核</el-radio>
        </el-form-item>
        <el-form-item label="概要" prop="summary">
          <el-input type="textarea" v-model="addForm.summary" :rows="5" resize="none"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <!--          <quill-editor v-model="addForm.content"></quill-editor>-->
          <Editor :value.sync="addForm.content"></Editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addPost">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改帖子对话框-->
    <el-dialog title="修改帖子" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               :rules="editFormRules">
        <el-form-item label="用户名" prop="Account">
          <el-input v-model="editForm.Author.Account" disabled></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="Title">
          <el-input v-model="editForm.Title"></el-input>
        </el-form-item>
        <el-form-item label="板块" prop="PartId">
          <el-select v-model="editForm.Part.PartName" placeholder="请选择">
            <el-option v-for="item in parts" :key="item.value"
                       :label="item.label" :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="State">
          <el-radio v-model="editForm.State" :label='0'>未审核</el-radio>
          <el-radio v-model="editForm.State" :label='1'>已审核</el-radio>
        </el-form-item>
        <el-form-item label="概要" prop="Summary">
          <el-input type="textarea" v-model="editForm.Summary" :rows="5" resize="none"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="Content">
          <!--          <quill-editor v-model="editForm.Content"></quill-editor>-->
          <Editor :value.sync="editForm.Content"></Editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editPostInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <div class="mg setcenter"><el-image :src="editForm.CoverUrl" style="width: 70%"></el-image></div>
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="用户名"> {{ editForm.Author.Account }}</el-descriptions-item>
          <el-descriptions-item label="收藏数"> {{ editForm.Favorite }}</el-descriptions-item>
          <el-descriptions-item label="板块"> {{ editForm.Part.PartName }}</el-descriptions-item>
          <el-descriptions-item label="发布时间"> {{ formattedPublishTime }}</el-descriptions-item>
          <el-descriptions-item label="悬赏"> {{ editForm.Reward }}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag v-if="editForm.State === 0" type="warning">未审核</el-tag>
            <el-tag v-else type="success">已审核</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <div class="mg">内容</div>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg"></quill-editor>
      </template>
    </Drawer>

  </div>

</template>

<script>


export default {
  name: "post",
  computed: {
    formattedPublishTime() {
      const publishTime = this.editForm.PublishTime;
      return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
    },
    mappedPartID() {
      const part = this.parts.find((item) => item.value === this.editForm.Part.PartName);
      return part ? part.value : '';
    },
  },
  data() {
    return {
      title: '帖子管理',
      drawerTitle: "帖子详情",
      //添加对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改对话框的显示与隐藏
      editDialogVisible: false,
      //数据总数
      total: 0,
      //获取列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //板块数据集合
      parts: [
        {value: '1', label: '帖子'},
        {value: '2', label: '问题'},
        {value: '3', label: '官方'},
      ],
      //查询到的信息
      editForm: {
        ID: '',
        Author: {},
        Part: {},
        State: '',
      },
      //帖子数据集合
      postList: [],
      //添加的表单数据
      addForm: {
        account: '',
        partID: '1',
        summary: '',
        title: '',
        state: '0',
        content: ''
      },
      // 添加规则
      addFormRules: {
        account: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {min: 3, max: 10, message: '用户名的长度在3~10个字符间', trigger: 'blur'}
        ],
        title: {required: true, message: '标题不能为空', trigger: 'blur'},
        summary: {required: true, message: '概要不能为空', trigger: 'blur'},
        content: {required: true, message: '内容不能为空', trigger: 'blur'},
      },
      // 修改规则
      editFormRules: {
        Title: {required: true, message: '标题不能为空', trigger: 'blur'},
        Summary: {required: true, message: '概要不能为空', trigger: 'blur'},
        Content: {required: true, message: '内容不能为空', trigger: 'blur'},
      },
      //表格配置
      columns: [
        {prop: 'Author.Account', label: '用户名', width: '150px'},
        {
          prop: 'PublishTime', label: '发布时间', width: '200px', sortable: true,
          formatter: (row, column) => {
            const publishTime = row[column.property];
            return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
          }
        },
        {prop: 'Part.PartName', label: '板块', width: '100px', sortable: true},
        {prop: 'Title', label: '标题', width: '200px', showOverflowTooltip: true},
      ],
    }
  },
  created() {
    this.getPostList()
  },
  methods: {
    // 查询方法
    async getPostList() {
      const {data: res} = await this.axios.post('post/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取帖子列表失败！')
      this.postList = res.data.posts
      this.total = res.data.total
      console.log(this.postList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getPostList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getPostList()
    },
    //删除
    async removePostById(id) {
      const result = await this.$confirm('确定要删除该帖子？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('post/delete', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除帖子失败！')
      this.$message.success('删除帖子成功！')
      await this.getPostList()
    },
    //审核
    async checkPostById(id) {
      const result = await this.$confirm('确定要通过该帖子？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消通过')
      }
      console.log(id);
      // 发起审核帖子信息的数据请求
      this.editForm.State = '1'
      console.log(id);
      const {data: res} = await this.axios.patch('post/update', {
        'id': id,
        'state': this.editForm.State
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('审核帖子失败！')
      }
      this.$message.success('审核成功')
      await this.getPostList()
    },

    // 显示详情
    showDetails(row) {
      console.log(row);
      this.editForm = row;
      this.detailsDrawer = true;
    },
    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },
    //添加对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    //点击按钮添加
    addPost() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        // 发起添加帖子网络请求
        const {data: res} = await this.axios.post('post/add', this.addForm, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          this.$message.error('添加帖子失败！')
        } else this.$message.success('添加帖子成功！')
        //隐藏对话框
        this.addDialogVisible = false
        //刷新用户列表
        await this.getPostList()
      })
    },
    //展示编辑的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('post/searchById', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询帖子信息失败！')
      }
      this.editForm = res.data
      this.editDialogVisible = true
    },
    // 修改对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },
    // 修改信息并提交
    editPostInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        const partID = this.mappedPartID || this.editForm.PartID; // 设置默认值或采取其他处理方式
        console.log(this.editForm)
        // 发起修改帖子信息的数据请求
        const {data: res} = await this.axios.patch('post/update', {
          'id': this.editForm.ID,
          // 'partID': this.editForm.PartID,
          'partID': partID,
          'summary': this.editForm.Summary,
          'content': this.editForm.Content,
          'state': this.editForm.State
        }, {
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新帖子失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getPostList()
        this.$message.success('更新帖子成功！')
      })
    },
    //获取焦点事件
    focus(event) {
      event.enable(false);
    }
  }

}
</script>

<style scoped>
</style>