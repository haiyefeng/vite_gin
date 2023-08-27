import { createApp } from 'vue'
import './style.css'
import router from '@/router'
import App from './App.vue'
import * as ElIcons from '@element-plus/icons-vue'

const app = createApp(App)
for (const name in ElIcons){
	app.component(name,(ElIcons as any)[name])
}

app.use(router).mount('#app')
