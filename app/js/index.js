import { createCharts } from "./data.js";

const mapOptions = {
  center: [38.003, 23.67605],
  zoom: 20,
};

const map = new L.map("map", mapOptions);
const layer = new L.TileLayer(
  "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
);
map.addLayer(layer);

const getStations = async () => {
  const response = await fetch("http://localhost:5000");
  const stations = await response.json();
  return stations;
};

const showModal = (station) => {
  if (station.id != undefined) createCharts(station.id);
  document.getElementById("myModal").style.display = "block";
};

const fillMarkers = async () => {
  const stations = await getStations();
  stations.forEach((station) => {
    const marker = new L.Marker([station.latitude, station.longitude]).on(
      "click",
      () => {
        showModal(station);
      }
    );
    marker.addTo(map);
  });
};

fillMarkers();
