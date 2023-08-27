<template>
  <div class="common-layout">
    <el-container height="100%">
      <el-header>Header</el-header>
      <el-container height="100%">
        <LeftAside></LeftAside>
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
import LeftAside from '@/components/LeftAside.vue'

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
    const istest = ref(false)
    return {
      istest,
    }
  },
  mounted() {
    getRequest('/api/hello').then((res: any) => {
      this.istest = res
    })
  },
  methods: {
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
</style>
