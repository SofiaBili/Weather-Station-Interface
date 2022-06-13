<template>
  <div class="main">
    <div class="content-area">
      <div class="container-fluid">
        <h2>Weather Station on Campus</h2>
        <GMapMap
          :center="center"
          :zoom="7"
          map-type-id="terrain"
          style="width: 100%; height: 500px"
        >
          <GMapCluster>
            <GMapMarker
              v-for="marker in markers"
              :key="marker.id"
              :position="marker.position"
              :title="marker.title"
              :icon="marker.icon"
            >
              <GMapInfoWindow>
                <div>
                  <h3>{{ marker.title }}</h3>
                  <p>{{ marker.content }}</p>
                  <b-button>Test</b-button>
                </div>
              </GMapInfoWindow>
            </GMapMarker>
          </GMapCluster>
        </GMapMap>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  setup () {
    return {
      center: {
        lat: 40.7128,
        lng: -74.0060
      },
      markers: [
        {
          position: {
            lat: 40.7128,
            lng: -74.0060
          },
          title: 'New York',
          content: 'New York City'
        }
      ]
    }
  }
}
</script>
