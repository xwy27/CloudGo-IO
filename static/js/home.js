// set author name and contact info from server
$(document).ready(() => {
  $.ajax({
    url: '/api/info',
    data: '',
    type: 'GET',
  }).then((res) => {
    $('.author').append(res.author);
    $('.contact').append(res.contact);
  });
});

// login button
$('#login').on('click', () => {
  window.location.href = '/login';
});