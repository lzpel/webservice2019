function getNowDate(){
  var dt = new Date();
  var y = dt.getFullYear();
  var m = ("00" + (dt.getMonth()+1)).slice(-2);
  var d = ("00" + dt.getDate()).slice(-2);
  var result = y + "/" + m + "/" + d;
  return result;
}
function nowTime(){
  var node = document.getElementById("name");
  node.textContent = "あいうえお";
}
nowTime();