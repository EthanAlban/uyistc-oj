<el-aside :style="'width: '+screenWidth*0.15+'px;z-index:10'" id="aside">
        <el-menu :default-active="$route.path"  @open="handleOpen" @close="handleClose" background-color="#112d4e"
          text-color="#eeeeee" active-text-color="#3f72af" router class="menuNav">
          >
          <el-menu-item @click="sys_info">
            <div v-if="this.$root.sys_status === 1">
              <div style="color:green">
                <i class="el-icon-cpu"></i>
                <span slot="title">服务器在线</span>
              </div>
            </div>
            <div v-if="this.$root.sys_status === -1">
              <div style="color:red">
                <i class="el-icon-cpu"></i>
                <span slot="title">OJ离线</span>
              </div>
            </div>
            <div v-if="this.$root.sys_status === 0">
              <div style="color:red">
                <i class="el-icon-cpu"></i>
                <span slot="title">服务器离线</span>
              </div>
            </div>
          </el-menu-item>
          <el-menu-item index="/home/nav1">
            <i class="el-icon-s-fold"></i>
            <span slot="title">首页</span>
          </el-menu-item>
          <el-menu-item index="/home/nav2">
            <i class="el-icon-document-copy"></i>
            <span slot="title">题库</span>
          </el-menu-item>
          <el-menu-item index="/home/nav3">
            <i class="el-icon-folder-opened"></i>
            <span slot="title">作业</span>
          </el-menu-item>
          <el-menu-item index="/home/nav4">
            <i class="el-icon-upload"></i>
            <span slot="title">我的提交</span>
          </el-menu-item>
          <el-menu-item index="/home/nav5">
            <i class="el-icon-potato-strips"></i>
            <span slot="title">站内排名</span>
          </el-menu-item>
        </el-menu>
      </el-aside>