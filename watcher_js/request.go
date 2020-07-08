package watcher_js
const RequestJs = `
document.addEventListener("DOMContentLoaded", function () {
    hljs.initHighlightingOnLoad();
});

function openTab(tabName) {
  var i;
  var x = document.getElementsByClassName("view-tab");
  for (i = 0; i < x.length; i++) {
    x[i].style.display = "none";
  }
  document.getElementById(tabName).style.display = "block";
}
`
