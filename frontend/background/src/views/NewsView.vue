<template>
  <div>
    <!--    面包屑-->
    <Breadcrumb class="round15 pd whiteback mg" :title="title"></Breadcrumb>
    <!--    资讯查询-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold;"> 资讯查询</div>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-input placeholder="请输入作者"  v-model="queryInfo.query">
            <el-button slot="append" icon="el-icon-search" @click="getNewsList"></el-button>
          </el-input>
        </el-col>
        <el-col :span=2>
          <el-button @click="addDialogVisible = true" type="primary"
                     round icon="el-icon-plus">添加资讯
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <!--    资讯列表-->
    <el-card class="round15 mg">
      <div style="font-size: 20px;font-weight: bold"> 资讯列表</div>
      <Table :table-data="newsList" :columns="columns" >
        <!--        操作区-->
        <template #operation="scope">
          <el-button size="mini" type="success" icon="el-icon-view" round
                     @click="showDetails(scope.row)">详情
          </el-button>
          <el-button size="mini" type="primary" icon="el-icon-edit" round
                     @click="showEditDialog(scope.row.ID)">编辑
          </el-button>
          <el-button size="mini" type="danger" icon="el-icon-delete" round
                     @click="removeNewsById(scope.row.ID)">删除
          </el-button>
        </template>
      </Table>
    </el-card>
    <!--    分页器-->
    <Pagination :total="total" :query-info="queryInfo"
                @page-size-change="handlePageSizeChange"
                @page-change="handlePageChange"></Pagination>
    <!--    添加资讯的对话框-->
    <el-dialog title="添加资讯" :visible.sync="addDialogVisible" width="50%"
               @close="addDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="addFormRef" :model="addForm" label-width="80px"
               :rules="addFormRules">
        <el-form-item label="作者" prop="author">
          <el-input v-model="addForm.author"></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="addForm.title"></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="part">
          <el-select v-model="addForm.type" placeholder="请选择">
            <el-option
                v-for="item in types"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="内容" prop="summary">
          <quill-editor v-model="addForm.content"></quill-editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
         <el-button @click="addDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="addNews">确 定</el-button>
      </span>
    </el-dialog>
    <!--    修改资讯对话框-->
    <el-dialog title="修改资讯" :visible.sync="editDialogVisible" width="50%"
               @close="editDialogClosed">
      <!--      内容主体区域-->
      <el-form ref="editFormRef" :model="editForm" label-width="80px"
               >
        <el-form-item label="作者" prop="author">
          <el-input v-model="editForm.Author"></el-input>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="editForm.Title"></el-input>
        </el-form-item>
        <el-form-item label="分类" prop="part">
          <el-select v-model="editForm.Type" placeholder="请选择">
            <el-option
                v-for="item in types"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="内容" prop="summary">
          <quill-editor v-model="editForm.Content"></quill-editor>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="editDialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="editNewsInfo">确 定</el-button>
  </span>
    </el-dialog>
    <!--    抽屉-->
    <Drawer :drawer="detailsDrawer" :title="drawerTitle" @closed="drawerClosed">
      <template #details="scope">
        <el-descriptions direction="vertical" :column="2" border class="mg">
          <el-descriptions-item label="作者"> {{ editForm.Author }}</el-descriptions-item>
          <el-descriptions-item label=类别"><el-tag>{{ editForm.type }}</el-tag> </el-descriptions-item>
          <el-descriptions-item label="标题" :span="2"> {{ editForm.Title }}</el-descriptions-item>
          <el-descriptions-item label="发布时间"> {{ formattedPublishTime }}</el-descriptions-item>
        </el-descriptions>
        <div class="mg">内容</div>
        <quill-editor v-model="editForm.Content" @focus="focus($event)" class="mg"></quill-editor>
      </template>
    </Drawer>
  </div>
</template>

<script>
export default {
  name: "TrendView",
  computed:{
    formattedPublishTime() {
      const publishTime = this.editForm.PublishTime;
      return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data(){
    return{
      title: '资讯管理',
      drawerTitle: "资讯详情",
      //添加对话框
      addDialogVisible: false,
      //控制详情抽屉可见否
      detailsDrawer: false,
      //控制修改对话框的显示与隐藏
      editDialogVisible: false,
      //数据总数
      total: 0,
      //获取资讯列表的参数对象
      queryInfo: {
        query: '',
        pageNum: 1,
        pageSize: 10
      },
      //资讯类型数据集合
      types:[
        {value:'考研常识',label:'考研常识'},
        {value:'考研政策',label:'考研政策'},
        {value:'选择院校',label:'选择院校'},
        {value:'备考指南',label:'备考指南'},
      ],
      //查询到的资讯信息
      editForm: {

      },
      //添加的表单数据
      addForm: {
        author: '',
        title: '',
        content: '',
        type: '1',
      },
      // 添加规则
      addFormRules: {
        author: [
          {required: true, message: '请输入作者名', trigger: 'blur'},
          {min: 2, max: 10, message: '作者名的长度在2~10个字符间', trigger: 'blur'}
        ],
        title:[
          {required: true, message: '标题不能为空', trigger: 'blur'},
        ],
      },
      //资讯数据集合
      newsList:[],
      columns: [
        {prop: 'Author', label: '作者', width: '150px',sortable: true},
        {prop: 'PublishTime', label: '发布时间', width: '200px', sortable: true,
          formatter: (row, column) => {
            const publishTime = row[column.property];
            return this.$moment(publishTime).format('YYYY-MM-DD HH:mm:ss');
          }},
        {prop: 'Title', label: '标题', width: '180px', showOverflowTooltip: true},
        {prop: 'Type', label: '分类', width: '100px', sortable: true}
      ],
    }
  },
  created() {
    this.getNewsList()
  },
  methods: {
    // 查询方法
    async getNewsList() {
      const {data: res} = await this.axios.post('news/list', {
        query: this.queryInfo.query,
        pageSize: this.queryInfo.pageSize,
        pageNum: this.queryInfo.pageNum
      }, {
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      });
      console.log(res)
      if (res.code !== 200) return this.$message.error('获取资讯列表失败！')
      this.newsList = res.data.newses
      this.total = res.data.total
      console.log(this.newsList);
    },
    //处理每页显示数量变化
    handlePageSizeChange(newSize) {
      this.queryInfo.pageSize = newSize;
      this.getNewsList()
    },
    //处理页码变化
    handlePageChange(newPage) {
      this.queryInfo.pageNum = newPage;
      this.getNewsList()
    },
    //删除资讯
    async removeNewsById(id) {
      const result = await this.$confirm('确定要删除该资讯？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).catch(err => err)
      console.log(result)
      if (result !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      console.log(id);
      const {data: res} = await this.axios.delete('news/delete', {
        params: {'id': id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) return this.$message.error('删除资讯失败！')
      this.$message.success('删除资讯成功！')
      await this.getNewsList()
    },
    // 显示详情
    showDetails(row) {
      console.log(row);
      this.editForm=row;
      this.detailsDrawer = true;
    },
    //关闭详情
    drawerClosed() {
      this.detailsDrawer = false;
    },
    //添加资讯对话框关闭
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    //点击按钮添加资讯
    addNews() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        console.log(this.addForm)
        try{
          // 发起添加资讯网络请求
          const {data: res} = await this.axios.post('news/add', this.addForm, {
            headers: {
              'Authorization': window.sessionStorage.getItem("token")
            }
          })
          if (res.code !== 200) {
            this.$message.error('添加资讯失败！')
          } else this.$message.success('添加资讯成功！')
          //隐藏对话框
          this.addDialogVisible = false
          //刷新用户列表
          await this.getNewsList()
        }catch (err){
          this.$message.error('添加资讯失败')
        }
      })
    },
    //展示编辑的对话框
    async showEditDialog(id) {
      console.log(id);
      const {data: res} = await this.axios.get('news/searchById', {
        params: {'id':id},
        headers: {
          'Authorization': window.sessionStorage.getItem("token")
        }
      })
      if (res.code !== 200) {
        return this.$message.error('查询用户信息失败！')
      }
      this.editForm = res.data
      this.editDialogVisible = true
    },
    // 修改对话框关闭
    editDialogClosed() {
      this.$refs.editFormRef.resetFields()
    },
    // 修改信息并提交
    editNewsInfo() {
      this.$refs.editFormRef.validate(async valid => {
        console.log(valid)
        if (!valid) return
        console.log(this.editForm)
        // 发起修改资讯信息的数据请求
        const {data: res} = await this.axios.patch('news/update', {
          'id': this.editForm.ID,
          'author': this.editForm.Author,
          'type': this.editForm.Type,
          'title': this.editForm.Title,
          'Content': this.editForm.Content
        },{
          headers: {
            'Authorization': window.sessionStorage.getItem("token")
          }
        })
        if (res.code !== 200) {
          return this.$message.error('更新资讯失败！')
        }
        // 关闭对话框
        this.editDialogVisible = false
        // 刷新数据列表
        await this.getNewsList()
        this.$message.success('更新资讯成功！')
      })
    },
    //获取焦点事件
    focus(event){
      event.enable(false);
    }
  }

}
</script>

<style scoped>

</style>