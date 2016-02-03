$(document).ready(function() {
  var session_key = null;

  $('#snapshot').hide();
  $('#loading').hide();

  $(".fancybox").fancybox();

  $(document).on('submit', 'form[data-remote=true]', function(e) {
    e.preventDefault();

    $.ajax({
      url: '/auth',
      method: 'POST',
      dataType: 'json',
      data: '{ "code": "' + $('#code').val() + '" }',
      error: function() {
        console.log('error');
      },
      complete: function(data) {
        session_key = JSON.parse(data.responseText).session_key;
        $('#snapshot').show();
      }
    });
  });

  $('#snapshot').click(function() {
    $('#loading').toggle();

    $.ajax({
      method: 'post',
      url: '/getimage',
      dataType: 'text',
      data: '{ "session_key": "' + session_key + '" }',
      error: function(xhr, textStatus, errorThrown) {
        $('#loading').toggle();
        console.log(xhr, textStatus, errorThrown + 'error');
      },
      complete: function(data) {
        $('#loading').toggle();
        $('.fancybox').attr('href', 'data:image/jpeg;base64,' + data.responseText);
        $('#image').attr('src', 'data:image/jpeg;base64,' + data.responseText);
      }
    });
  });
});
