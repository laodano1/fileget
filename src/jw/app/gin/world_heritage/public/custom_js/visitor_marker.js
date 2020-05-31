
loadMap();

function loadMap() {
    var default_location = [30.66243, 104.0625];
    var mymap = L.map('mapid').setView(default_location, 11);  //

    L.tileLayer.chinaProvider('GaoDe.Normal.Map',{maxZoom:18,minZoom:5}).addTo(mymap);  // 高德地图插件

    var marker1 = L.marker(default_location, {opacity: 0.5}).addTo(mymap);

}


$('li.litem').mouseover(function () {
    $(this).css('background-color', '#eeeeee');
    // $('this.span.stext').css('color', 'green');
}).mouseout(function () {
    $(this).css('background-color', 'white').css('color', 'black');
});

// $(document).ready(handleVistorInfo);
$('#bt').click(function () {
    // $('#16').scrollIntoView();
    var record = document.getElementById("l6");
    record.scrollIntoView();
    $('#l6').css('background-color', 'yellow');
    // var timer = setInterval(
    //     $('#l6').css('background-color', 'yellow')
    //     , 3000);

    // $('#l6').css('background-color', 'white');
});

$('#l1')
    .mouseover(function () {
        $(this).css('background-color', 'gray');
    })
    // .mouseout(function () {
    //     $(this).css('background-color', 'white');
    // });

handleVistorInfo();

function handleVistorInfo() {
    var ip = '194.182.64.181';
    var ip2 = '1.64.139.29';
    var ipArr = [ip, ip2];
    var i = 0;
    var arrTHeadItems = ['city', 'country', 'countryCode', 'isp', 'regionName', 'timezone', 'org', 'latitude', 'longitude'];
    $('#visitorInfo').empty();
    $('#visitorInfo').append(
        $('<table></table>').attr('id','visitors').append(
            $('<tr></tr>').attr('id','th')
        )
    );
    arrTHeadItems.forEach(function (value) {
        $('#th').append(
            $('<th></th>').text(value)
        );
    });

    $.get('/visitors', function (visitorInfo, status) {
        visitorInfo.forEach(function (item) {
            i++;
            if(i >= 2) {return;}
            getVistorDetails(item.ip, i);
        });
    });

    // ipArr.forEach(
function getVistorDetails(ip, i) {
        $.get('http://ip-api.com/json/' + ip, function (data, status) {
            console.log(status);

            if (data.status !== 'success') {return;}
            // if (data.status === 'success') {
                // console.log(data);
                var city = data.city;
                var country = data.country;
                var countryCode = data.countryCode;  // country initials
                var isp = data.isp;
                var regionName = data.regionName;   // province
                var timezone = data.timezone;   // province
                var org = data.org;
                var lat = data.lat;
                var lon = data.lon;


                var arrRepItems = [city, country, countryCode, isp, regionName, timezone, org, lat, lon];
                var tHead;

                $('#visitors').append(
                    $('<tr></tr>').attr('id', 'l' + i)
                );
                arrRepItems.forEach(function (value) {
                    $('#l' + i).append(
                        $('<td></td>').text(value)
                    );
                });

            // }
            // else {
            //     // $('#visitorInfo').css('text-align', 'center').text('Can not get the IP information!');
            //     console.log('get location of ' + ip + ' fail');
            //     $('#visitorInfo').css({'text-align': 'center', 'font-weight': 'bold', 'color': 'red'}).text(data.message);
            // }
        });
};
    // );

}

// http://ip-api.com/json/47.100.37.81



