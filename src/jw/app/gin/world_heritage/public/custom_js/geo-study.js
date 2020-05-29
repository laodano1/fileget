var mymap = L.map('map').setView([30.66243, 104.0625], 9);  // 天府广场

// var mymap = L.map('map').setView([37.8, -96], 11);  //

L.tileLayer.chinaProvider('GaoDe.Normal.Map',{maxZoom:18,minZoom:5}).addTo(mymap);  // add china 高德 map layer
//
// L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
//     // L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png', {
//     attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="http://mapbox.com">Mapbox</a>',
//     // maxZoom: 18,
//     // tileSize: 512,
//     // zoomOffset: -1,
//     id: 'mapbox.streets',
//     accessToken: 'pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpejY4NXVycTA2emYycXBndHRqcmZ3N3gifQ.rJcFIG214AriISLbB6B5aw'
// }).addTo(mymap);   // 'addTo': Adds the control to the given map.


bounds = L.latLngBounds([[ 30.3, 104.1625], [ 30.54, 104.3623]]);

// var marker = L.marker([ 32, -130]).addTo(mymap);
// marker.bindPopup('first').openPopup();
//
// var marker2 = L.marker([ 13, -100]).addTo(mymap);
// marker2.bindPopup('second').openPopup();

L.rectangle(bounds).addTo(mymap);

// mymap.fitBounds(bounds);

options = {opacity: 0.8};
// var imageUrl = ['www.baidu.com/img/bd_logo1.png'];
// var overlay = L.imageOverlay( imageUrl, bounds, options ).addTo(mymap);

var videoUrls = [
    'http://localhost:3001/pic/patricia_nasa.webm'

];
var videoOverlay = L.videoOverlay( videoUrls, bounds, options ).addTo(mymap);

// mymap.addLayer(videoOverlay);

videoOverlay.on('load', function () {
    var MyPauseControl = L.Control.extend({
        onAdd: function() {
            var button = L.DomUtil.create('button');
            button.innerHTML = '⏸';
            L.DomEvent.on(button, 'click', function () {
                videoOverlay.getElement().pause();
            });
            return button;
        }
    });
    var MyPlayControl = L.Control.extend({
        onAdd: function() {
            var button = L.DomUtil.create('button');
            button.innerHTML = '▶️';
            L.DomEvent.on(button, 'click', function () {
                videoOverlay.getElement().play();
            });
            return button;
        }
    });

    var pauseControl = (new MyPauseControl()).addTo(mymap);
    var playControl = (new MyPlayControl()).addTo(mymap);
});

//
// $.get('/cdcs', function (data, status) {
//     // var myGeoStyle = {
//     //     "color": "blue",
//     //     "weight": 3,
//     //     "opacity": 0.35
//     // };
//
//     if(status === 'success') {
//         L.geoJSON(data, {
//             style: style,
//             onEachFeature: onEachFeature
//         }).addTo(mymap);
//     }
// });


var style = function (feature) {
    // console.log('in style');
    // console.log(feature);
    return {
        fillColor: "#e3b926",
        weight: 3,
        opacity: 0.3,
        color: 'blue',
        fillOpacity: 0.3
    };
};


var onEachFeature = function (feature, layer) {
    // console.log('in onEachFeature');
    // console.log(feature);
    // console.log(layer);
    layer.on({
       mouseover: onMouseover,
       mouseout: onMouseout,
       click: onClick
    });

};

var onClick = function (e) {
    var countryName        = e.target.feature.properties.name;

    // if (countryName !== marker.getContent())
    // if ( marker !== '') { marker.remove(); }
    var marker = L.marker(e.latlng).addTo(mymap);

    marker.bindPopup(countryName).openPopup();

    // console.log(e.latlng);
    //
    setTimeout(function () {
        marker.remove();
    }, 1000);

    // L.popup().setLatLng(e.latlng).setContent(countryName).openOn(mymap);
};

var onMouseout = function (e) {
    // console.log('mouse out');

};

var onMouseover = function (e) {
    // console.log('mouse over');
    // console.log(e);
    var countryName        = e.target.feature.properties.name;
    var countryCoordinates = e.target.feature.geometry.coordinates[0][0];  // 里面的坐标的经纬度需要互换



    // console.log(countryCoordinates);

    var coors = [[30.67453427248429, 104.1167449951172],
        [30.67453427248429,104.01992797851562],
        [30.60423095372707,104.03709411621095],
        [30.60423095372707,104.12910461425781]];
    // console.log(coors);
    var polygon = L.polygon(coors).addTo(mymap);
    mymap.fitBounds(polygon.getBounds());
};













