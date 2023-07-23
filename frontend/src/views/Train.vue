<template>
    <el-container style="height: 100vh;">
        Train
    </el-container>
    <button @click="store.increment()">
    From B: {{ store.count }}
  </button>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue'
  import { ElContainer, ElAside, ElMain, ElMenu, ElMenuItem, ElFooter } from 'element-plus'
  import { store } from '@/store.js'
  
  export default defineComponent({
    components: {
      ElContainer,
      ElAside,
      ElMain,
      ElMenu,
      ElMenuItem,
      ElFooter
    },
    data() {
    return {
      store
    }
  }
  })
  </script>
  
  <style scoped>
  </style>
  