let mapOptions = {
    center: [38.00303, 23.67605],
    zoom: 20,
}

let map = new L.map('map', mapOptions);
let layer = new L.TileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png');
map.addLayer(layer);

let marker1 = new L.Marker([38.002275, 23.67632]).on('click', function(e) {
  showData("marker1")
});

marker1.addTo(map);

let marker2 = new L.Marker([38.001582, 23.675108]).on('click', function(e) {
  showData("marker2")
});
marker2.addTo(map);
  

let marker3 = new L.Marker([38.001941, 23.674325]).on('click', function(e) {
  showData("marker3")
});
marker3.addTo(map);
  

let marker4 = new L.Marker([38.004025, 23.675966]).on('click', function(e) {
  showData("marker4")
});
marker4.addTo(map);

function showData(markerName){

  var newContent = markerName;

  var contentHolder = document.getElementById('modalTitle');
  contentHolder.innerHTML = newContent;

  document.getElementById('myModal').style.display = "block";
}