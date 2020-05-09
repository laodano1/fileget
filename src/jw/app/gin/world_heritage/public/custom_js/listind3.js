var width = $(window).width();
var height = $(window).height() * 0.08;

var cntWidth = 400;
var cntHeight = $(window).height() * 0.1;

var danjiaWay = '123';
var zongjiaWay = '123';

var loupans = [];

var bt1  =  $('<div></div>').attr('id', 'danjia').text('单价排序');
// var dvd1 =  $('<div></div>').attr('id', 'divider1').text('');
var bt2  =  $('<div></div>').attr('id', 'zongjia').text('总价排序');
var toTop  =  $('<button></button>').attr('id', 'back2Top')
            .attr('type', 'button')
            .html('分享')
            .on('click', function (e) {
                console.log('back2Top clicked');
                // $("html, body").animate({scrollTop: 0}, 1000);  // back to top
                _handleClick();
             });

var desc = $('<div></div>').attr('id', 'desc').text('欢迎关注微信公众号 ---- "成都好房指南"');
var img = $('<img></img>').attr('id', 'shareQR').attr('src', 'pic/qrcode-cdhfzn.jpg');
var back2List = $('<div></div>').attr('id', 'back2List').text('返回列表');

// add two buttons
$('.sort').append(bt1).append(bt2); //.append(toTop);
$('.firstLine').append(toTop);

function _handleClick(obj, item, index) {
//     $('#back').show();
    $('.sort').fadeOut();
    $('.list').fadeOut();
    $('body').append(img).append(desc).append(back2List);

    $('#back2List').on('click', function (e) {
        location.reload();
    })
}

$.get('/lpjson?month=201805', function (list, status) {
   if(status === 'success') {
       // loupans = list;
       generateList(list);

       $('#danjia').on('click', function () {
           // console.log('单价 clicked!');
           // console.log(danjiaWay);
           var title1 = '单价排序';
           // var partlist1 = list.slice(3, 7);

           var danjia = _clickHandle(danjiaWay, list, title1);

           $('.list').empty();
           generateList(danjia);
       });

       $('#zongjia').on('click', function () {
           var title2 = '总价排序';
           // var partlist2 = list.slice(3, 7);
           // var zongjia = _clickHandle(zongjiaWay, partlist2, title2);
           var zongjia = _clickHandle(zongjiaWay, list, title2);
           // console.log(zongjia.length);
           $('.list').empty();
           generateList(zongjia);
       });
   }
   else {
       $('.list').text('获取数据失败！');
   }
});

function _handleSVGClick(item, index) {
    if (item.hasOwnProperty('wenzurl')) {
        window.location = item.wenzurl;
    }
    else {
        window.location = 'http://47.100.37.81/prof?lp=' + (item.id-1);
    }
}


function generateList(list) {
    var listItem = d3.select('body').select('.list')
        .selectAll('svg')
        .data(list)
        .enter()
        .append('svg')
            .attr('width', function () {
                return width;
            })
            .attr('height', function () {
                return cntHeight * 2;
            })
            .attr('position', 'fixed')
            .attr('id', function (d, i) {
                return 'svg' + i;
            })
            .on('click', function (d, i) {
            // console.log('svg clicked' + i);
            // $("html, body").animate({scrollTop: 0}, 1000);  // back to top
            _handleSVGClick(d, i);
        });

// for locating picture
    listItem.append("rect")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.1 + ") rotate(0)")
        .attr('fill', 'white')
        .attr('stroke', 'green')
        .attr('width', '100px')
        .attr('height', '100px')
        .attr('stroke-dasharray', 2)
        .attr('x', 0)
        .attr('y', cntHeight * 0.003);

    // add picture
    listItem.append("svg:image")
        // .attr('x', width * 0.0)
        .attr('x', width * 0.03)
        .attr('y', cntHeight * 0.1)
        // .attr('y', cntHeight * 0.003)
        .attr('width', '95px')
        .attr('height', '100px')
        // .attr("xlink:href", "pic/cdhfc.jpeg");
        .attr("xlink:href", "pic/timg.jpeg");

    // for first title
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.3 + ", " + cntHeight * 0.32 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .attr('fill', 'green')
        .style('text-anchor', 'start')
        .attr("font-size", 16)
        .attr("text-anchor", "middle")
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            // return "楼盘标题1\n";
            return d.name + ' [' + d.block + ']';
        });
    // for second title
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.3 + ", " + cntHeight * 0.99 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('fill', '#96c8da')
        .style('text-anchor', 'start')
        // .attr('style', 'gray')
        .attr("font-size", 12)
        .attr("text-anchor", "middle")
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d.huxing + '/' + d.type ;
        });

    // for second title
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.3 + ", " + cntHeight * 1.39 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('fill', 'red')
        .style('text-anchor', 'start')
        // .attr('style', 'gray')
        .attr("font-size", 13)
        .attr("text-anchor", "middle")
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            if (d.price.match(/元/)) {
                return '' + d.price;
            }
            else {
                return '' + d.price + '元/㎡';
            }

        });
    //
}

// sort by 单价 or 总价
function sortList(data, key, way, title) {
    return data.sort(function (a, b) {
        var x,y;
        if (title === '单价排序') {
            // 通过运算，自动转换成数字，不然是按照字母顺序排序
            x = handlePirce($.trim(a.price)) * 1;
            y = handlePirce($.trim(b.price)) * 1;
        }

        else if(title === '总价排序') {
            // a.unit
            var p1 = handlePirce($.trim(a.price));
            var p2 = handlePirce($.trim(b.price));

            var u1 = handlePirce($.trim(a.huxing));
            var u2 = handlePirce($.trim(b.huxing));

            x = p1 * u1;
            y = p2 * u2;
        }
        else {
            console.log('排序没有指定');
            return -1;
        }

        if (way === '123')  {
            return ((x > y) ? 1 : -1);
        }
        if (way === '321') {

            return ((x < y) ? 1 : -1);
        }
    })
}

//
function _clickHandle(sortWay, data, title) {
    // if (touched != 0) {
    var lps = [];
    var way = sortWay;
        if (way === '123') {
            // console.log(data);
            lps = sortList(data, 'price', '321', title);

            if (title === '单价排序') {
                // console.log('单价');
                $('#danjia').text(title + ' ↓');
                $('#zongjia').text('总价排序');
            }
            if (title === '总价排序') {
                // console.log('总价');
                $('#zongjia').text(title + ' ↓');
                $('#danjia').text('单价排序');
            }

            way = '321';
        }
        else {
            lps = sortList(data, 'price', '123', title);

            if (title === '单价排序') {
                $('#danjia').text(title + ' ↑');
                $('#zongjia').text('总价排序');
            }
            if (title === '总价排序') {
                $('#zongjia').text(title + ' ↑');
                $('#danjia').text('单价排序');
            }
            way = '123';
        }

    // return [lps, touched, sortWay];
    if(title === '单价排序') {danjiaWay = way;}
    if(title === '总价排序') {zongjiaWay = way;}

    return lps;
}

function handlePirce(price) {
    // console.log(price);
    if (price === '待定') {
        price = 0;
    }
    else {
        // console.log(x);
        if (/-/.test(price)) {
            price = (price.split(/-/)[0]).match(/\d+/);
        }
        else {
            price = price.match(/\d+/);
        }
    }
    // console.log(price);
    return price;
}

