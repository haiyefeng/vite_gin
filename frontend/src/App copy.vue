<template>
  <div class="common-layout">
    <el-container height="100%">
      <el-header>Header</el-header>
      <el-container height="100%">
        <el-aside class="layout-left">
            <el-menu
            default-active="2"
            active-text-color="#ffd04b"
            background-color="#545c64"
            class="el-menu-vertical"
            text-color="#fff"
            :collapse="isCollapse"
            @open="handleOpen"
            @close="handleClose"
          >
            <el-sub-menu index="1">
              <template #title>
                <el-icon><location /></el-icon>
                <span>Navigator One</span>
              </template>
              <el-menu-item-group>
                <template #title><span>Group One</span></template>
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
              </el-menu-item-group>
              <el-menu-item-group title="Group Two">
                <el-menu-item index="1-3">item three</el-menu-item>
              </el-menu-item-group>
              <el-sub-menu index="1-4">
                <template #title><span>item four</span></template>
                <el-menu-item index="1-4-1">item one</el-menu-item>
              </el-sub-menu>
            </el-sub-menu>
            <el-menu-item index="2">
              <el-icon><icon-menu /></el-icon>
              <template #title>Navigator Two</template>
            </el-menu-item>
            <el-menu-item index="3" disabled>
              <el-icon><document /></el-icon>
              <template #title>Navigator Three</template>
            </el-menu-item>
            <el-menu-item index="4">
              <el-icon><setting /></el-icon>
              <template #title>Navigator Four</template>
            </el-menu-item>
          </el-menu>
          <el-icon size="20px" text-color="red" background-color="red" @click="toggleCollapse"><DArrowLeft v-if="!isCollapse"/><DArrowRight v-if="isCollapse"/></el-icon>
        </el-aside>
        <el-container>
          <el-main class="main">Main</el-main>
          <el-footer class="footer">Footer</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue';
import { ElContainer, ElAside, ElMain, ElMenu, ElMenuItem, ElFooter, ElIcon } from 'element-plus';
import { getRequest } from '@/utils/request';

export default defineComponent({
  components: {
    ElContainer,
    ElAside,
    ElMain,
    ElMenu,
    ElMenuItem,
    ElFooter,
    ElIcon
  },
  setup() {
    const isCollapse = ref(false)
    const expendIcon = ref('<DArrowRight />')
    // <DArrowRight />
    return {
      isCollapse,
      expendIcon
    }
  },
  mounted() {
    getRequest('/api/hello').then((res: any) => {
      console.log(res)
    })
  },
  methods: {
    handleOpen (key: string, keyPath: string[]) {
      console.log(key, keyPath)
    },
    handleClose (key: string, keyPath: string[]) {
      console.log(key, keyPath)
    },
    toggleCollapse(){
      this.isCollapse = !this.isCollapse;
    }
  }
})
</script>

<style scoped>
.common-layout {
  height: 100%;
  width: 100%;
}
.el-header {
  padding: 0 auto;
  margin: 0 auto;
  width: 100%;
  height: 5vh;
  background-color: aqua;
}
.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
}
.el-menu-vertical {
  height: 93vh;
}
.main {
  height: 90vh;
  width: 100%;
  background-color: green;
  padding: 0 auto;
  margin: 0 auto;
}

.footer {
  height: 5vh;
  width: 100%;
  background-color: blueviolet;
  padding: 0 auto;
  margin: 0 auto;
}

.layout-left {
  /* 左侧不设置宽度 */
  width: auto;
  height: 95vh;
  /* position: relative; */
  background-color: #545c64;
}
/* #app {
  height:100vh;
  margin: 60px 20px;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
} */
</style>
