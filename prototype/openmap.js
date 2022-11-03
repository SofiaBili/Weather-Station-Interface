let mapOptions = {

    center: [38.00303, 23.67605],
    zoom: 17,
}

let map = new L.map('map', mapOptions);
let layer = new L.TileLayer('http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png');
map.addLayer(layer);

let marker = new L.Marker([38.00303, 23.67605]);