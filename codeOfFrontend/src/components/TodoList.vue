<template>
  <el-card class="box-card">
    <el-row :gutter="20">
      <el-col :span="16" :offset="2">
        <el-input v-model="newTitle" size placeholder="请输入待办事项..." />
      </el-col>
      <el-col :span="6">
        <el-button type="primary" icon="el-icon-plus" @click="handleAdd" circle></el-button>
      </el-col>
    </el-row>
    <el-divider></el-divider>
    <el-table :data="tableData" style="width: 100%" :row-class-name="tableRowClassName">
      <el-table-column type="index" width="50"></el-table-column>
      <el-table-column align="center" label="待办事项" prop="title"></el-table-column>
      <el-table-column align="center" label="创建日期" prop="date"></el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
          <el-button
            type="success"
            icon="el-icon-check"
            v-show="!scope.row.status"
            @click="handleEdit(scope.$index, scope.row)"
            circle
          ></el-button>
          <el-button
            type="warning"
            icon="el-icon-refresh-left"
            v-show="scope.row.status"
            @click="handleEdit(scope.$index, scope.row)"
            circle
          ></el-button>
          <el-button
            type="danger"
            icon="el-icon-close"
            @click="handleDelete(scope.$index, scope.row.id)"
            circle
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
export default {
  name: "TodoList",
  data() {
    return {
      tableData: [],
      newTitle: "",
      nowDate: "",
    };
  },
  mounted() {
    // 页面加载完后显示当前时间
    this.dealWithTime(new Date())
    // 定时刷新时间
    this.timer = setInterval(()=> {
       this.dealWithTime(new Date()) // 修改数据date
    }, 500)
    this.axios
      .get("/v1/todo")
      .then(response => (this.tableData = response.data));
  },
  destroyed() {
    if (this.timer) {  // 注意在vue实例销毁前，清除我们的定时器
      clearInterval(this.timer);
    }
  },
  methods: {
    tableRowClassName({ row }) {
      if (row.status) {
        return "success-row";
      } else {
        return "";
      }
    },
    getTodoList() {
      this.axios
        .get("/v1/todo")
        .then(response => (this.tableData = response.data));
    },
    handleEdit(index, row) {
      let messageSuffix = row.status ? " 置为未完成" : " 置为已完成";
      this.axios
        .put("/v1/todo/" + row.id, {
          status: !row.status
        })
        .then(() => {
          this.tableData[index].status = !row.status;
          this.$message({
            showClose: true,
            duration: 1500,
            message: `<${row.title}> ${messageSuffix}`,
            type: "success"
          });
        });
    },
    handleDelete(index, id) {
      this.axios.delete("/v1/todo/" + id).then(() => {
        this.tableData.splice(index, 1);
        this.$message({
          showClose: true,
          duration: 1500,
          message: "删除待办事项成功",
          type: "success"
        });
      });
    },
    handleAdd() {
      if (this.newTitle != "") {
        this.axios
          .post("/v1/todo", {
            title: this.newTitle,
            date: this.nowDate
          })
          .then(() => {
            this.getTodoList();
            this.$message({
              showClose: true,
              duration: 1500,
              message: "添加待办事项成功",
              type: "success"
            });
          });
        this.newTitle = "";
      } else {
        this.$message({
          showClose: true,
          duration: 1500,
          message: "title不能为空哦",
          type: "warning"
        });
      }
    },
    dealWithTime(data) { // 获取当前时间
      let Y = data.getFullYear();
      let M = data.getMonth() + 1;
      let D = data.getDate();
      this.nowDate = Y + "-" + M + "-" + D;
    }
  }
};

</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  text-decoration: line-through;
}
</style>