import { defineStore } from "pinia";
import { getStation } from "@/services/station";
import { ref } from "vue";

export const useDataStore = defineStore("data", () => {
  const station = ref(undefined);

  const getStationData = async (id) => {
    station.value = await getStation(id);
  };

  return { station, getStationData };
});
