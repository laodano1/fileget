// var $j = jQuery.noConflict();
getLoupans();

var unittouched = '0';
var pricetouched = '0';
// var touched = '0';
var unitsortWay = '0';
var pricesortWay = '0';
// var data = $('ul');
//
// popup();
//
// function popup() {
//     $('#s4').on('click', '.open-about', function (e) {
//         e.prventEvent()
//         console.log('s4 clicked');
//         $.popup('.popup-about');
//     });
//     //
//     // $('#s4').on('click','.open-services', function () {
//     //     $.popup('.popup-services');
//     // });
// }


function getLoupans() {
    $.getJSON('/lpjson', function (data) {

        $('#unitprize').on('click', function () {
            // console.log('unitprize button touched');
            // unittouched = 1;
            // unitsortWay = '123';
            // var lps;
            var title = '单价排序';
            $.toast(title);
            // console.log('unit: ' + unittouched + ' ' + unitsortWay);
            var values = _clickHandle(unittouched, unitsortWay, data, title);

            // console.log(values[0]);
            unittouched = values[1];
            unitsortWay = values[2];
            generateSortList(values[0]);

        });

        $('#wholeprize').on('click', function () {
            var title = '总价排序';
            $.toast(title);
            // console.log('price: ' + pricetouched + ' ' + pricesortWay);
            // console.log(lps);
            var values = _clickHandle(pricetouched, pricesortWay, data, title);
            pricetouched = values[1];
            pricesortWay = values[2];
            generateSortList(values[0]);
        });

    })
}


function _clickHandle(touched, sortWay, data, title) {
    if (touched != 0) {
        if (sortWay === '123') {
            lps = sortList(data, 'price', '321', title);
            // lps.forEach(function (object) {
            //     console.log(object.price);
            // });
            if (title === '单价排序') {
                $('#unitprize').text(title + '↓');
            }
            if (title === '总价排序') {
                $('#wholeprize').text(title + '↓');
            }
            // $('#unitprize').text(title + '↓');
            sortWay = '321';
        }
        else {
            lps = sortList(data, 'price', '123', title);
            // lps.forEach(function (object) {
            //     console.log(object.price);
            // });
            // $('#unitprize').text(title + '↑');
            if (title === '单价排序') {
                $('#unitprize').text(title + '↑');
            }
            if (title === '总价排序') {
                $('#wholeprize').text(title + '↑');
            }
            sortWay = '123';
        }

    }
    else {
        lps = sortList(data, 'price', '123', title);
        if (title === '单价排序') {
            $('#unitprize').text(title + '↑');
        }
        if (title === '总价排序') {
            $('#wholeprize').text(title + '↑');
        }
        // $('#unitprize').text(title + '↑');
        sortWay = '123';
        touched = 1;
    }

    return [lps, touched, sortWay];
}

//
function generateSortList(lps) {
    $('ul').empty();
    // console.log($('ul').children().get(0));
    lps.forEach(function (lpItem) {
        var tmpli = $('<li></li>').attr('id', 2)
            .append(
                $('<a></a>').addClass('item-link item-content').attr('href', '#')
                    .append(
                        $('<div></div>').addClass('item-media').append($('<img></img>').attr('src', 'pic/cdhfc.jpeg'))
                    )
                    .append(
                        $('<div></div>').addClass('item-inner')
                            .append(
                                $('<div></div>').addClass('item-title-row')
                                    .append($('<div></div>').addClass('item-title').text(lpItem.name))
                                    .append($('<div></div>').addClass('item-after'))
                            )
                            .append(
                                $('<div></div>').addClass('item-subtitle').text(lpItem.price)
                            )
                            .append(
                                $('<div></div>').addClass('item-text').text(lpItem.unit)
                            )
                    )

            );
        // console.log(tmpli);
        $('ul').append(tmpli);

    });

}

// 排序
function sortList(data, key, way, title) {
    // $j('ul').children()

    return data.sort(function (a, b) {
        // console.log(a.toString().properties);
        // $('<li>')

        // var str = a.toString().get('.item-text');
        // console.log( str);

        // var texta = a.find('.item-title').text();
        // var textb = b.find('.item-title').text();
        // // console.log(texta);
        var x,y;
        if (title === '单价排序') {
            // x = $.trim(a[key]);
            // y = $.trim(b[key]);

            x = handlePirce($.trim(a[key]));
            y = handlePirce($.trim(b[key]));
        }

        else if(title === '总价排序') {
            // a.unit
            p1 = handlePirce($.trim(a.price));
            p2 = handlePirce($.trim(b.price));

            u1 = handlePirce($.trim(b.unit));
            u2 = handlePirce($.trim(b.unit));

            // x = $.trim(a[key]);
            // y = $.trim(b[key]);
            x = p1 * u1;
            y = p2 * u2;
        }
        else {
            console.log('排序没有指定');
            return -1;
        }

        //
        // var x = $.trim(a[key]);
        // var y = $.trim(b[key]);
        // var firstPrize = 0

        // if (y === 8400) {
        //     console.log('x: ' + x + ' y :' + y);
        // }

        if(way === '123')  {
            return ((x > y) ? 1 : -1);
        }
        if (way === '321') {

            return ((x < y) ? 1 : -1);
        }

    })
}
//
// function handleUnit(unit) {
//     if (unit == "") {
//         unit = 0;
//     }
//     else {
//         if (/-/)
//     }
// }
//

function handlePirce(price) {
    if (price === '待定') {
        price = 0;
    }
    else {
        // console.log(x);
        if (/-/.test(price)) {
            price = price.split(/-/)[0];
        }
    }

    return price;
}


