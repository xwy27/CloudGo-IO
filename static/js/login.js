$('#login-form').form({
  fields: {
    email: 'empty',
    password: 'empty'
  }
});

$('.submit').on('click', () => {
  $('#login-form').submit();
});