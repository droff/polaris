$(document).ready(function() {
  $('#snapshot').hide();

  $(document).on('submit', 'form[data-remote=true]', function(e){
    e.preventDefault();
    $.ajax({
      url: '/auth',
      method: 'POST',
      dataType: 'json',
      data: '{ "code": "' + $('#code').val() + '" }'
    }).error(function() {
      console.log('error');
    }).done(function(data) {
      $('#snapshot').show();
    });
  });

  $('#snapshot').click(function() {
    $('#image').attr('src', '/assets/images/loading.gif');
    $.ajax({
      method: 'post',
      url: '/getimage',
      dataType: 'text',
      complete: function(data) {
        var rawImage = data.responseText;
        $('#image').attr('width', '1280');
        $('#image').attr('height', '720');
        $('#image').attr('src', 'data:image/jpeg;base64,' + rawImage);
      },
      error: function(xhr, textStatus, errorThrown) {
        console.log(xhr, textStatus, errorThrown + 'error');
      }
    });
  });
});
