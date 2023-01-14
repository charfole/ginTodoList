<template>
  <el-container>
    <el-header>ginTodoList</el-header>
    <el-main>
      <el-row type="flex" justify="center">
        <el-col :xs="20" :span="12" >
          <div class="grid-content">
            <el-divider>
              <h1>待办小清单</h1>
            </el-divider>
            <TodoList></TodoList>
          </div>
        </el-col>
      </el-row>
    </el-main>
    <el-footer>
      <div>
        {{ nowDate }}{{ nowWeek }}
      </div>
    </el-footer>
  </el-container>
</template>

<script>
// @ is an alias to /src
import TodoList from "@/components/TodoList.vue";

export default {
  name: "Index",
  components: {
    TodoList
  },
  data(){
    return {
      formatDateTime: "",
      nowDate: "",    // 当前日期
      nowTime: "",    // 当前时间
      nowWeek: ""     // 当前星期
    }
  },
  methods:{
    dealWithTime(data) { // 获取当前时间
      let Y = data.getFullYear();
      let M = data.getMonth() + 1;
      let D = data.getDate();
      let H = data.getHours();
      let Min = data.getMinutes();
      let S = data.getSeconds();
      let W = data.getDay();
      H = H < 10 ? "0" + H : H;
      Min = Min < 10 ? "0" + Min : Min;
      S = S < 10 ? "0" + S : S;
      switch (W) {
        case 0:
          W = "日";
          break;
        case 1:
          W = "一";
          break;
        case 2:
          W = "二";
          break;
        case 3:
          W = "三";
          break;
        case 4:
          W = "四";
          break;
        case 5:
          W = "五";
          break;
        case 6:
          W = "六";
          break;
        default:
          break;
      }
      this.nowDate = Y + "年" + M + "月" + D + "日 ";
      this.nowWeek = "周" + W ; 
      this.nowTime = H + ":" + Min + ":" + S;
    },
  },
  mounted() { 
    // 页面加载完后显示当前时间
    this.dealWithTime(new Date())
    // 定时刷新时间
    this.timer = setInterval(()=> {
       this.dealWithTime(new Date()) // 修改数据date
    }, 500)
  }, 
  destroyed() {
    if (this.timer) {  // 注意在vue实例销毁前，清除我们的定时器
      clearInterval(this.timer);
    }
  }
}
</script>



<style>
.el-header,
.el-footer {
  background-color: #409eff;
  color: #fff;
  text-align: center;
  line-height: 60px;
}
.el-footer {
  background-color: #909399;
  display: block;
  width: 100%;
  position: fixed;
  bottom: 0;
}
</style>