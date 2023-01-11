<template>
  <div class="main">
    <div class="content-area">
      <div class="container-fluid">
        <h2>Weather Station on Campus</h2>
        <l-map style="height: 350px; z-index: 0" :zoom="zoom" :center="center">
          <l-tile-layer :url="url" :attribution="attribution"></l-tile-layer>
          <l-marker
            v-for="(marker, index) in markers"
            :key="index"
            :lat-lng="marker"
          ></l-marker>
        </l-map>
      </div>
    </div>
  </div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer, LMarker } from "@vue-leaflet/vue-leaflet";
import { ref } from "vue";

export default {
  name: "Map",
  components: {
    LMap,
    LTileLayer,
    LMarker,
  },
  setup() {
    const zoom = ref(20);
    const center = ref([47.31322, -1.319482]);
    const url = ref("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png");
    const attribution = ref(
      '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors'
    );
    const markers = ref([[38.002275, 23.67632]]);

    return {
      zoom,
      center,
      url,
      attribution,
      markers,
    };
  },
};
</script>
