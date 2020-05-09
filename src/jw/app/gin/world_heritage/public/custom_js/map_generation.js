var mymap = L.map('mapid').setView([30.66243, 104.0625], 10);  // 天府广场

// API: http://leafletjs.com/reference-1.2.0.html#tilelayer
// L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
//     // L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png', {
//     attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="http://mapbox.com">Mapbox</a>',
//     // maxZoom: 18,
//     // tileSize: 512,
//     // zoomOffset: -1,
//     id: 'mapbox.streets',
//     accessToken: 'pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpejY4NXVycTA2emYycXBndHRqcmZ3N3gifQ.rJcFIG214AriISLbB6B5aw'
// }).addTo(mymap);   // 'addTo': Adds the control to the given map.

L.tileLayer.chinaProvider('GaoDe.Normal.Map',{maxZoom:18,minZoom:5}).addTo(mymap);  // 高德地图插件

var circle;
var polygon;

var month = '201805';

$.get('/lpjson?month=' + month, function (data, status) {
    var loupanInfo = data;

    if (loupanInfo.length < 1) {return;}

    if (status !== 'success') {return;}  // exception handle

// collected markers
    var TFNA = [];
    var longquan = [];
    var notTFNA = [];
    allAreas = {};  //全部区域
    var oneArea = {};  //tmp区域
    Areas = {};  //tmp区域

    loadAllLoupan(loupanInfo);

// console.log("sametimeStamp: " + TFNA.length);
    var allItemLayer = L.layerGroup();
// console.log(Areas);

    var allMakrkers = [];
    var baseLayers = {};

    Object.keys(Areas).forEach(function (oneArea) {
        // console.log(oneArea);
        baseLayers[oneArea] = L.layerGroup(allAreas[oneArea]);

        allAreas[oneArea].forEach(function (oneMarker) {
            allMakrkers.push(oneMarker);
        });
    });

    var allMarkersLayer = L.layerGroup(allMakrkers).addTo(mymap);
// var notTFNALayer = L.layerGroup(notTFNA).addTo(mymap);

// console.log(allMakrkers);
    baseLayers["全城"] = allMarkersLayer;
//
    L.control.layers(baseLayers, null).addTo(mymap);

});



// 默认加载的楼盘的标记
function loadAllLoupan(loupanInfo) {
    // console.log("loupanInfo.loupans.length: " + loupanInfo.loupans.length);

    // var redMarker = L.AwesomeMarkers.icon({
    //     icon: 'star',
    //     markerColor: 'blue'
    // });

    loupanInfo.forEach(function (data) {
        var lpName = "<b>楼盘: </b> " + data.name + "<br>";
        if(data.hasOwnProperty('wenzurl')) { // 如果有微信文章链接
            lpName = "<b>楼盘: </b><a href=' " + data.wenzurl + "'> " + data.name + "</a><br>";
        }
        else  {  //如果没有微信页面链接
            lpName = "<b>楼盘: </b><a href='http://47.100.37.81/prof?lp=" + (data.id-1) + "'> " + data.name + "</a><br>";
        }

        var lpBlock = "<b>区域: </b>" + data.block + "<br>";
        var lpType = "<b>类型: </b> " + data.type + "<br>";
        var lpPrice = "<b>价格: </b> " + data.price + "<br>";
        var lpHuxing = "<b>户型: </b> " + data.huxing + "<br>";

        lpBlock = "<p style=\"color: darkgray;margin-bottom: 0px;margin-top: 0px;\">" + lpBlock + "</p>";
        var popCnt = lpName + lpBlock + lpType + lpPrice + lpHuxing;

        if(data.taoshu !== '') {
            var taoshu = "<b>套数: </b> " + data.taoshu + "<br>";
            popCnt += taoshu;
        }


        if (data.hasOwnProperty('subway') && data.subway.length != 0) {
            var subway = "<b>地铁: </b> ";
            data.subway.forEach(function (value) {
                subway += value + "<br>";
            });
            popCnt += subway;
        }
        if (data.hasOwnProperty('xuequ') && data.xuequ.length != 0) {
            var xuequ = "<b>学区: </b> ";
            data.xuequ.forEach(function (value) {
                xuequ += value + "<br>";
            });
            popCnt += xuequ;
        }

        // var marker1 = L.marker(data.coordinate, {opacity: 0.5, icon: redMarker})
        var marker1 = L.marker(data.coordinate, {opacity: 0.5})
                        .bindPopup(popCnt)
                        .openPopup();
        var circle = L.circle(data.coordinate, {
            color: 'red',
            fillColor: null,
            fillOpacity: 0.5,
            radius: 10
        });

        if (!allAreas.hasOwnProperty(data.block)) {
        //     allAreas.add(data.block, [].push(marker1));
            allAreas[data.block] = [marker1, circle];
        }
        else {
            allAreas[data.block].push(marker1);
            allAreas[data.block].push(circle);
        }

        // console.log(data.block);
        if (!Areas.hasOwnProperty(data.block)) {
            Areas[data.block] = 0;
        }
    });
}

addBROptionsPanel();
function addBROptionsPanel() {
    var brPanel = L.control({position: 'bottomright'});
    brPanel.onAdd = function (map) {
        var div = L.DomUtil.create('div', 'brPanel');
        var labels = [];

        // labels.push('<button style="background: #0ea432"></button>');
        // div.innerHTML = labels.join('<br>');
        var element = '<div id="buttonl">选项</div>';
        div.innerHTML = element;
        return div;

    };

    brPanel.addTo(mymap);
}

$('#buttonl').on('click', function (e) {

    console.log('button clicked!');
    if($('.pboxBR').length === 0) {
        addPopupBoxBR();
    }
    else {
        $('.pboxBR').remove();
    }

});
//
// $(document).on('click', function (e) {
//     console.log('document clicked!');
//     if($('#areaBR').is(':visible')) {
//         console.log('in if ');
//         $('.pboxBR').remove();
//     }
// });

function addPopupBoxBR() {
    var brPanel = L.control({position: 'bottomright'});
    brPanel.onAdd = function (map) {
        var div = L.DomUtil.create('div', 'pboxBR');

        // console.log(d1);
        var element = '<div id="areaBR">'
            + '<div class="innerIcon" id="d1">精装</div>'
            + '<div class="innerIcon" id="d2">清水</div>'
            + '<div class="innerIcon" id="d3">叠拼</div>'
        '</div>';
        div.innerHTML = element;
        return div;
    };

    brPanel.addTo(mymap);
}


// addBLOptionsPanel();
function addBLOptionsPanel() {
    var brPanel = L.control({position: 'bottomleft'});
    brPanel.onAdd = function (map) {
            var div = L.DomUtil.create('div', 'blPanel');
        var labels = [];

        // labels.push('<button style="background: #0ea432"></button>');
        // div.innerHTML = labels.join('<br>');
        var element = '<select id="select">' +
            '<option value="school">学校</option>' +
            '<option value="loupan">楼盘</option>'
            + '</select>';
        div.innerHTML = element;
        // div.innerHTML = '<div id="blBut" >更多.左</div>';

        // $('#blPanel').on('click', function (e) {
        //     console.log('blPanel clicked');
        //     var win =  L.control.window(map, { title:'Hello world!',maxWidth:400, modal: true})
        //         .content('Info: ')
        //         .prompt({callback:function(){alert('This is called after OK click!')}})
        //         .show();
        // });

        return div;

    };

    brPanel.addTo(mymap);
}


// $('#blPanel').on('click', function (e) {
//     console.log('blPanel clicked');
//     var win =  L.control.window(mymap, { title:'Hello world!',maxWidth:400, modal: true})
//         // .content('Info:' + Feature.properties.Comments)
//         .content('Info: ')
//         .prompt({callback:function(){alert('This is called after OK click!')}})
//         .show();
// });

// var CanvasLayer = L.GridLayer.extend({
//     createTile: function(coords, done){
//         var error;
//         // create a <canvas> element for drawing
//         var tile = L.DomUtil.create('canvas', 'leaflet-tile');
//         // setup tile width and height according to the options
//         var size = this.getTileSize();
//         tile.width = size.x;
//         tile.height = size.y;
//         // draw something asynchronously and pass the tile to the done() callback
//         setTimeout(function() {
//             done(error, tile);
//         }, 1000);
//         return tile;
//     }
// });
//
// var greenIcon = L.icon({
//     iconUrl: 'pic/icon-normal-big.png'
//     ,
//     shadowUrl: 'leaf-shadow.png',
//
//     iconSize:     [38, 95], // size of the icon
//     shadowSize:   [50, 64], // size of the shadow
//     iconAnchor:   [22, 94], // point of the icon which will correspond to marker's location
//     shadowAnchor: [4, 62],  // the same for the shadow
//     popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
// });
