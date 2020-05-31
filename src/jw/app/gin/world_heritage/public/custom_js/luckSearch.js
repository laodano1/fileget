var width = $(window).width();
var height = $(window).height() * 0.08;

var cntWidth = 400;
var cntHeight = $(window).height() * 0.1;

var danjiaWay = '123';
var zongjiaWay = '123';

var loupans = [];

// var bt1  =  $('<div></div>').attr('id', 'danjia').text('单价排序');
// // var dvd1 =  $('<div></div>').attr('id', 'divider1').text('');
// // var bt2  =  $('<div></div>').attr('id', 'zongjia').text('总价排序');
// var toTop  =  $('<button></button>').attr('id', 'back2Top')
//             .attr('type', 'button')
//             .html('分享')
//             .on('click', function (e) {
//                 console.log('back2Top clicked');
//                 // $("html, body").animate({scrollTop: 0}, 1000);  // back to top
//                 _handleClick();
//              });

// var desc = $('<div></div>').attr('id', 'desc').text('欢迎关注微信公众号 ---- "成都好房指南"');
// var img = $('<img></img>').attr('id', 'shareQR').attr('src', 'pic/qrcode-cdhfzn.jpg');
// var back2List = $('<div></div>').attr('id', 'back2List').text('返回列表');


// var bt1  =  $('<input></input>').attr('id', 'searchbox').text('单价排序');

// add two buttons
// $('.sort').append(bt1).append(bt2); //.append(toTop);
// $('.firstLine').append(toTop);

// function _handleClick(obj, item, index) {
// //     $('#back').show();
//     $('.sort').fadeOut();
//     $('.list').fadeOut();
//     $('body').append(img).append(desc).append(back2List);
//
//     $('#back2List').on('click', function (e) {
//         location.reload();
//     })
// }


$.get('/yaohaojson?type=short', function (list, status) {
   if(status === 'success') {
       // loupans = list;
       console.log('get yaohao result');
       // var partlist1 = list.slice(3, 7);
       generateList(list);

       $('#sbutton').on('click', function () {
           // console.log('单价 clicked!');
           // console.log(danjiaWay);

           var lpName = $('select option:selected').val();
           var subStr = $('#sb2').val();
           $('#sb2').val('');
           // var partlist1 = list.slice(3, 7);

           var retVal = _clickHandle(lpName, subStr);


       });

       // $('#zongjia').on('click', function () {
       //     var title2 = '总价排序';
       //     // var partlist2 = list.slice(3, 7);
       //     // var zongjia = _clickHandle(zongjiaWay, partlist2, title2);
       //     var zongjia = _clickHandle(zongjiaWay, list, title2);
       //     // console.log(zongjia.length);
       //     $('.list').empty();
       //     generateList(zongjia);
       // });
   }
   else {
       $('.list').text('获取摇号数据失败！');
   }
});


//
function _clickHandle(lpName, subStr) {
    var tmpStr = lpName + '+' + subStr;
    if (subStr.trim() === '') { return;}
    $.get('/yaohaojson?sc=' + tmpStr, function (data, status) {
        if (status === 'success') {
            $('.list').empty();
            generateList(data);
        }
        else {
            $('.list').text('查询信息未成功！');
        }
    });


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
            });


    generateDescription(listItem);
    generateContents(listItem);
// // for locating picture
//     listItem.append("rect")
//         .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.1 + ") rotate(0)")
//         .attr('fill', 'white')
//         .attr('stroke', 'green')
//         .attr('width', '100px')
//         .attr('height', '100px')
//         .attr('stroke-dasharray', 2)
//         .attr('x', 0)
//         .attr('y', cntHeight * 0.003);

    // add picture
    // listItem.append("svg:image")
    //     // .attr('x', width * 0.0)
    //     .attr('x', width * 0.03)
    //     .attr('y', cntHeight * 0.1)
    //     // .attr('y', cntHeight * 0.003)
    //     .attr('width', '95px')
    //     .attr('height', '100px')
    //     // .attr("xlink:href", "pic/cdhfc.jpeg");
    //     .attr("xlink:href", "pic/timg.jpeg");

    // for descriptions

    //
    // // for first title
    // listItem.append("text")
    //     .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.32 + ") rotate(0)")
    //     .attr("font-family", "sans-serif")
    //     .attr('fill', 'green')
    //     .style('text-anchor', 'start')
    //     .attr("font-size", 14)
    //     .attr("text-anchor", "middle")
    //     .attr('fill', 'black')
    //     .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
    //         // return "楼盘标题1\n";
    //         // console.log('in text svg');
    //         return '公证摇号编号: ' + d['公证摇号编号'];
    //     });
    //
    //
    // // for second title
    // listItem.append("text")
    //     .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.59 + ") rotate(0)")
    //     .attr("font-family", "sans-serif")
    //     .style('fill', 'black')
    //     .style('text-anchor', 'start')
    //     // .attr('style', 'gray')
    //     .attr("font-size", 14)
    //     .attr("text-anchor", "middle")
    //     .attr('fill', 'black')
    //     .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
    //         return '身份证照号码： ' + d['身份证照号码'];
    //     });
    //
    // // for second title
    // listItem.append("text")
    //     .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.86 + ") rotate(0)")
    //     .attr("font-family", "sans-serif")
    //     .style('fill', 'black')
    //     .style('text-anchor', 'start')
    //     // .attr('style', 'gray')
    //     .attr("font-size", 14)
    //     .attr("text-anchor", "middle")
    //     .attr('fill', 'black')
    //     .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
    //             return '姓名: ' + d['姓名'] ;
    //     });
    //
}


function generateDescription(listItem) {
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.32 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .attr('fill', 'green')
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .attr("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            // return "楼盘标题1\n";
            // console.log('in text svg');
            return '公证摇号编号: ';
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.62 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .attr("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return '身份证照号码: ';
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 0.92 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .style("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return '姓名: ';
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 1.22 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .style("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return '购房登记号: ';
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 1.52 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .style("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return '选房顺序号: ';
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.02 + ", " + cntHeight * 1.82 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        // .style("text-anchor", "end")
        .attr('fill', 'gray')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return '备注: ';
        });
}

function generateContents(listItem) {
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 0.32 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            // return "楼盘标题1\n";
            // console.log('in text svg');
            return d['公证摇号编号'];
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 0.62 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d['身份证照号码'];
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 0.92 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d['姓名'];
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 1.22 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d['购房登记号'];
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 1.52 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d['选房顺序号'];
        });
    listItem.append("text")
        .attr("transform", "translate(" + cntWidth * 0.23 + ", " + cntHeight * 1.82 + ") rotate(0)")
        .attr("font-family", "sans-serif")
        .style('text-anchor', 'start')
        .attr("font-size", 12)
        .attr('fill', 'black')
        .text(function (d) {    // 根据绑定的data进行赋值 [2009, 2010]
            return d['备注'];
        });
}



