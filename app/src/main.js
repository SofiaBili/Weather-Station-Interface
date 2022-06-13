import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'

// Bootstrap libraries
import BootstrapVue3 from 'bootstrap-vue-3'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'

// Google Maps API
import VueGoogleMaps from '@fawmi/vue-google-maps'

createApp(App)
  .use(store)
  .use(router)
  .use(BootstrapVue3)
  .use(VueGoogleMaps, {
    load: {
      key: 'AIzaSyBinH7ZodZEFqMAW8U4DeZ2SN86dWu93d0'
    }
  })
  .mount('#app')
