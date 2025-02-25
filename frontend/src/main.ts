import './assets/style/base.less'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'
import { lightTheme, darkTheme } from './plugins/theme'

import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()
const vuetify = createVuetify({
    theme: {
        defaultTheme: 'lightTheme',
        themes: {
            lightTheme,
            darkTheme
        }
    }
})

app.use(pinia)
app.use(router)
app.use(vuetify)

app.mount('#app')

export { vuetify }
