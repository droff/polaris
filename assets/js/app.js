$(document).ready(function() {
  $('#snapshot').hide();

  $(document).on("submit", "form[data-remote=true]", function(e){
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
    $.ajax({
      url: '/getimage',
      method: 'POST',
      dataType:"image/jpeg",
    }).always(function(data) {
      var rawImage = data.responseText;
      $("#image").attr("src", "data:image/jpeg;base64," + rawImage);
    });
  });
});
