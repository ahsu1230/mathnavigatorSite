console.log("My Javascript works");

function init() {
  console.log("jQuery ready!");
  var flavor = window.localStorage.getItem("flavor");
  $("select.milk-tea").val(flavor);
  var ice = window.localStorage.getItem("ice");
  $("select.ice-amount").val(ice);
  var sugar = window.localStorage.getItem("sugar");
  $("select.sugar-amount").val(sugar);  
  var bubbles = window.localStorage.getItem("bubbles");
  var isChecked = bubbles =="true";
  $("input.bubbles").prop("checked", isChecked);
  var grassJelly = window.localStorage.getItem("grass-jelly");
  var isChecked2 = grassJelly =="true";
  $("input.grass-jelly").prop("checked", isChecked);
  var pudding = window.localStorage.getItem("pudding");
  var isChecked3 = pudding =="true";
  $("input.pudding").prop("checked", isChecked);

  // Setup on-click handler
  $("button").on("click", handleClick);
  $(".milk-tea").on("change", handleSelectChange);
  $(".ice-amount").on("change", handleSelectChange);
  $(".sugar-amount").on("change", handleSelectChange);
  $(".bubbles").on("change", handleCheckboxChange);
  $(".grass-jelly").on("change", handleCheckboxChange);
  $(".pudding").on("change", handleCheckboxChange);
}


function handleSelectChange(e) {
  var targetVal = e.target.value;
  var targetClass = e.target.className;
  if (targetClass == "milk-tea"){
	window.localStorage.setItem("flavor", targetVal);
  }
  if (targetClass == "ice-amount"){
	window.localStorage.setItem("ice", targetVal);
  }
  if (targetClass == "sugar-amount"){
	window.localStorage.setItem("sugar", targetVal);
  }
}

function handleCheckboxChange(e) {
  var targetVal = e.target.checked;
  var targetClass = e.target.className;
  if (targetClass == "bubbles"){
    window.localStorage.setItem("bubbles", targetVal);
  }
  if (targetClass == "grass-jelly"){
    window.localStorage.setItem("grass-jelly", targetVal);
  }
  if (targetClass == "pudding"){
    window.localStorage.setItem("pudding", targetVal);
  }
}

function handleClick() {
  var milkTeaSelected = $(".milk-tea").val();
  var iceAmount = $(".ice-amount").val();
  var sugarAmount = $(".sugar-amount").val();
  var name = $(".enterName").val();
  var url = "html/receipt.html?flavor="+milkTeaSelected+"&ice="+iceAmount+"&sugar="+sugarAmount+"&name="+name;
  if ($(".bubbles").prop("checked") == true) {
	  url += "&toppings=bubbles";
  }
  if ($(".grass-jelly").prop("checked") == true) {
	  url += "&toppings=grassjelly";
  }
  if ($(".pudding").prop("checked") == true) {
	  url += "&toppings=pudding";
  }
  console.log(url);
  window.location = url;
}


$(document).ready(init);