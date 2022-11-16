let mapOptions = {

    center: [38.00303, 23.67605],
    zoom: 17,
}

let map = new L.map('map', mapOptions);
let layer = new L.TileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png');
map.addLayer(layer);

let marker1 = new L.Marker([38.002275, 23.67632]);

marker1
  .addTo(map)
  .bindPopup('Εδώ γενικά έχει φουλ άνεμο προσέχετε!');

  let marker2 = new L.Marker([38.001582, 23.675108]);

marker2
  .addTo(map)
  .bindPopup('Εδώ έχει λιακάδα!');

  let marker3 = new L.Marker([38.001941, 23.674325]);

marker3
  .addTo(map)
  .bindPopup('Εδώ χιονίζει!');

  let marker4 = new L.Marker([38.004025, 23.675966]);

marker4
  .addTo(map)
  .bindPopup('Εδώ ο αισθητήρας ανιχνεύει υγρασία! Προσοχή σε αυτούς που έχουν ρευματικά!');