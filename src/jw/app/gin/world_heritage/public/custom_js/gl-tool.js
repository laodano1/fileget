$('#vm').click(function () {
    $('#display').empty();
    $.get('/gl?page=vmarker', function (data, status) {
        if(status === 'success') {
            $('#display').append(data);
        }
        else {
            $('#display').append('page request failed!');
        }
    });
});


$('#dt').click(function () {
    $('#display').empty();
    $.get('/gl?page=record', function (data, status) {
       if(status === 'success') {
           $('#display').append(data);
       }
       else {
           $('#display').append('page request failed!');
       }
    });
});