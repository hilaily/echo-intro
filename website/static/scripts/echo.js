(function() {
  // Add table class
  var tables = document.getElementsByTagName('table');
  for (var i = 0; i < tables.length; i++) {
    var t = tables[i];
    t.classList.add('table');
    t.classList.add('is-narrow');
  }

  // Prism
  Prism.languages.sh = Prism.languages.bash;
})();
