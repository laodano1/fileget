// $(document).ready( function () {
//     $('#myTable').DataTable({
//         data: tableData,
//         columns: [
//             {data: 'ID'},
//             {data: '楼盘名称'},
//             {data: '户型'},
//             {data: '单价'}
//         ]
//     });
// } );
//
// tableData = [
//     {"ID" : ""},
//     {"楼盘名称" : "中海国际"},
//     {"户型" : "70-100"},
//     {"单价" : "6666"}
// ];

$('#submit').click(function () {
    // console.log('button clicked');
    flag = 'ok';
    $('#notification').empty();

    var name = $('#name').val();   // mandatory

    var type = $('#type').val();    // mandatory
    var huxing = $('#huxing').val();   //
    var price = $('#price').val();    //
    var block = $('#sl option:selected').val();   // ----------

    var lattitude = $('#lat').val();
    var longitude = $('#lon').val();


    var sb1 = $('#sb1').is(":checked") ? $('#sb1').val() : '';
    var sb2 = $('#sb2').is(":checked") ? $('#sb2').val() : '';
    var sb3 = $('#sb3').is(":checked") ? $('#sb3').val() : '';
    var sb4 = $('#sb4').is(":checked") ? $('#sb4').val() : '';
    var sb5 = $('#sb5').is(":checked") ? $('#sb5').val() : '';
    var sb6 = $('#sb6').is(":checked") ? $('#sb6').val() : '';
    var sb7 = $('#sb7').is(":checked") ? $('#sb7').val() : '';
    var sb8 = $('#sb8').is(":checked") ? $('#sb8').val() : '';
    var sb9 = $('#sb9').is(":checked") ? $('#sb9').val() : '';
    var sb10 = $('#sb10').is(":checked") ? $('#sb10').val() : '';
    var sb11 = $('#sb11').is(":checked") ? $('#sb11').val() : '';

    var ditie = [sb1, sb2, sb3, sb4, sb5, sb6, sb7, sb8, sb10];
    ditie = removeEmptyElement(ditie);

    var xuequ = $('#xuequ').val();
    var schools = [];
    if (!xuequ.match(/,/)) {
        if (xuequ.match(/\n/)) {
            schools = xuequ.split(/\n/);
        }
    }
    else {
        setWarning('学区里的学校名称只能通过行隔开')
    }

    var taoshu = $('#taoshu').val();

    isValEmpty(name, '楼盘名称不能为空');
    isValEmpty(huxing, '户型不能为空');
    isValEmpty(block, '区域不能为空');
    isValEmpty(lattitude, '纬度不能为空');
    isValEmpty(longitude, '经度不能为空');

    var latLon = [lattitude, longitude];

    schools = removeEmptyElement(schools);
    const currentMonth = $('#mon option:selected').val();
    // var i = 0;
    $.get('/lpjson?month=' + currentMonth, function (data, status) {
        if(status === 'success') {
            var i = data.length;
            var itemJson = {
                id: i+1,
                name: name,
                type: type,
                huxing: huxing,
                price: price,
                block: block,
                coordinate: latLon,
                subway: ditie,
                xuequ: schools,
                taoshu: taoshu,
                wenzurl: ''
            };
            console.log(itemJson);
            console.log(flag);

            if (flag === 'ok') {
                $.post('/record', itemJson, function (data, status) {
                    $('#notification').css('color','green').text('Upload : ' + status);
                    valueClear();
                })
                    .fail(function (err) {
                        // console.log(err);
                        // console.log(obj);
                        $('#notification').css('color','red').text(err.statusText + ' : ' + err.responseText);
                        valueClear();
                    });
            }
            else {
                console.log('input wrong!');
                valueClear();
            }
        }
        else {
            console.log('get lou pan data failed!')
        }
    });
    // console.log(ditie);




});

function valueClear() {
    $('#name').val('');
    $('#lat').val('');
    $('#lon').val('');
    $('#type').val('');
    $('#huxing').val('');
    $('#price').val('');
    $('#xuequ').val('');
}

function isValEmpty(val, msg) {
    if (val === '') {
        setWarning(msg);
        flag = 'wrong';
        // return ;
    }
}

function removeEmptyElement(array) {
    var tmp = [];
    for(var i = 0; i < array.length; i++) {
        if (array[i] !== '') {
            tmp.push(array[i]);
        }
    }
    return tmp;
}


function setWarning(msg) {
    $('#notification').text(msg);
}

function isEmpty(val) {
    setWarning('');
}