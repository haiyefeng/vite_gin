<template>
  <el-container style="height: 100vh;">
    <el-aside :width="isCollapse ? '64px' : '200px'">
      <router-view></router-view>
      <el-menu
        default-active="1"
        :collapse="isCollapse"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <el-menu-item index="1">
          <i class="el-icon-message"></i>
          <template #title>首页</template>
        </el-menu-item>
        <el-menu-item index="2">
          <i class="el-icon-message"></i>
          <template #title>关于</template>
        </el-menu-item>
        <el-menu-item index="3">
          <i class="el-icon-message"></i>
          <template #title>联系我们</template>
        </el-menu-item>
        <el-menu-item index="4" @click="isCollapse = !isCollapse">
          <i class="el-icon-arrow-left"></i>
          <template #title>折叠菜单</template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-main>
        <h1>欢迎来到你的首页</h1>
        <p>这是你的首页的一些信息。</p>
      </el-main>
      <el-footer>Copyright &copy; 2023</el-footer>
    </el-container>
  </el-container>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { ElContainer, ElAside, ElMain, ElMenu, ElMenuItem, ElFooter } from 'element-plus'
import { getRequest } from '@/utils/request';

export default defineComponent({
  components: {
    ElContainer,
    ElAside,
    ElMain,
    ElMenu,
    ElMenuItem,
    ElFooter
  },
  setup() {
    const isCollapse = ref(false)

    return {
      isCollapse
    }
  },
  mounted() {
    getRequest('/api/hello').then((res: any) => {
      console.log(res)
    })
  }
})
</script>

<style scoped>
.el-header {
  padding: 0 20px;
}
.el-menu-vertical-demo {
  height: 100%;
  border-right: 0;
}
#app {
 font-family: Avenir, Helvetica, Arial, sans-serif;
 -webkit-font-smoothing: antialiased;
 -moz-osx-font-smoothing: grayscale;
 text-align: center;
 color: #2c3e50;
 margin-top: 60px;
}
</style>
