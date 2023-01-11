import { defineStore } from "pinia";
import { getStations } from "@/services/station";
import { ref } from "vue";

export const useMapStore = defineStore("map", () => {
  const markers = ref(undefined);

  const getMarkers = async () => {
    const stations = await getStations();
    markers.value = stations.map((station) => [
      station.latitude,
      station.longitude,
    ]);
  };

  return { markers, getMarkers };
});
