<template>
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
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { ElAside, ElMenu, ElMenuItem, ElIcon } from 'element-plus';
import { getRequest } from '@/utils/request';

export default defineComponent({
  components: {
    ElAside,
    ElMenu,
    ElMenuItem,
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
.el-menu-vertical:not(.el-menu--collapse) {
  width: 200px;
}
.el-menu-vertical {
  height: 93vh;
}


.layout-left {
  /* 左侧不设置宽度 */
  width: auto;
  height: 95vh;
  /* position: relative; */
  background-color: #545c64;
}
</style>
