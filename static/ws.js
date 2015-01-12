/*jslint browser: true */
(function () {
  'use strict';

  var hide = function (elem) {
    elem.classList.add('hidden');
  };
  var show = function (elem) {
    elem.classList.remove('hidden');
  };

  var btnConnect = document.getElementById('ButtonConnect'),
      txtUsername = document.getElementById('InputName'),
      pSignup = document.getElementById('Signup'),
      pCoffee = document.getElementById('Coffee');
  hide(pCoffee);

  var connectSocket = function (username, lat, lng) {
    var ws = new WebSocket('ws://localhost:8080/ws');
    var connected = false;
    ws.addEventListener('open', function () {
      ws.send(JSON.stringify({Name: username, Lat: lat, Lng: lng}));
    });
    ws.addEventListener('message', function (evt) {
      var res = JSON.parse(evt.data);
      console.log(res);
      if (!connected) {
        if (res.Status === 'Success') {
          hide(pSignup);
          show(pCoffee);
        }
      }
    });
  };
  
  btnConnect.addEventListener('click', function (evt) {
    evt.preventDefault();
    if (txtUsername.value !== '') {
      navigator.geolocation.getCurrentPosition(function (pos) {
        connectSocket(txtUsername.value, pos.coords.latitude, pos.coords.longitude);
      });
    }
  });

}());
