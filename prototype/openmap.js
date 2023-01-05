let mapOptions = {

    center: [38.00303, 23.67605],
    zoom: 20,
}

let map = new L.map('map', mapOptions);
let layer = new L.TileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png');
map.addLayer(layer);

let marker1 = new L.Marker([38.002275, 23.67632]).on('click', function(e) {
  showmodal();
});
marker1.addTo(map);
//.bindPopup('Εδώ έχει λιακάδα!');

let marker2 = new L.Marker([38.001582, 23.675108]).on('click', function(e) {
  showmodal();
});
marker2.addTo(map);
  

let marker3 = new L.Marker([38.001941, 23.674325]).on('click', function(e) {
  showmodal();
});
marker3.addTo(map);
  

let marker4 = new L.Marker([38.004025, 23.675966]).on('click', function(e) {
  showmodal();
});
marker4.addTo(map);


//functions

//when you click the marker marker1 a modal appears
function showmodal() {
  document.getElementById('myModal').style.display = "block";
}