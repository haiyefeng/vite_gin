<template>
    <el-container style="height: 100vh;">
        Home
        From A: {{ store.count }}
    </el-container>
  </template>
  
<script lang="ts">
import { defineComponent } from 'vue'
import { ElContainer, ElAside, ElMain, ElMenu, ElMenuItem, ElFooter } from 'element-plus'
import { store } from '@/store.js'

export default defineComponent({
  data() {
    return {
      store
    }
  },
  components: {
    ElContainer,
    ElAside,
    ElMain,
    ElMenu,
    ElMenuItem,
    ElFooter
  }
})
</script>

<style scoped>
</style>
