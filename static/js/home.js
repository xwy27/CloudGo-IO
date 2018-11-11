$(document).ready(() => {
  $.ajax({
    url: '/api/id',
    data: '',
    type: 'GET',
  }).then((res) => {
    $('.id').append(res.id);
  });
});