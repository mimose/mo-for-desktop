<template>
    <v-app id="inspire">
        <v-overlay z-index="999" :value="appOverlay">
          <v-progress-circular
            indeterminate 
            size="64"
          ></v-progress-circular> 
        </v-overlay>

        <v-system-bar app 
          height="36"
          fixed
          color="accent darken-4"  
        >
          <v-banner single-line @click:icon="goHomePage">
            <!-- <v-icon
              slot="icon"
              color="warning"
              size="25"
            >mdi-wifi-strength-alert-outline</v-icon> -->
            <span style="color: white">built by mimose / mimosecy@gmail.com</span>
          </v-banner>
        </v-system-bar>

        <v-navigation-drawer
            v-model="drawer"
            app
            width="20%"
            fixed
            permanent
        >
            <moToolBar></moToolBar>

            <!-- <v-sheet
                color="grey lighten-5"
                height="5%"
                width="100%"
            ></v-sheet> -->
            
            <record-list :appOverlaySync.sync='appOverlay'></record-list>

        </v-navigation-drawer>

    </v-app>
</template>

<script>
import "vuetify/dist/vuetify.min.css";
import "./assets/css/main.css";
import "material-design-icons-iconfont";

import MoToolBar from './components/MoToolBar.vue';
import RecordList from './components/RecordList.vue';

export default {
  name: "mo-app",
  data: () => ({
    drawer: null,
    appOverlay: false
  }),
  components: {
    moToolBar: MoToolBar,
    recordList: RecordList
  },
  computed: {
  },
  methods: {
    goHomePage() {
      window.open('https://www.baidu.com','_blank')
    },
  }
};

</script>
