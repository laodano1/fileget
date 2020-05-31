const cd = require('/Users/sophia/Downloads/mapdata/geometryCouties/510100.json');


cd.features.forEach(function (item) {
    console.log(item.properties);
});

