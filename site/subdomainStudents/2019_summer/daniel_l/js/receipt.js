console.log("My Javascript works");

var toppings = [];

var imgMap = {
  brownsugar: "../assets/caramel.jpg",
  jasmine: "../assets/cappuccino.jpg",
  oolong: "../assets/black_coffee.jpg",
  thai: "../assets/latte.jpg",
  taro: "../assets/cold_brew.jpg"
};

var flavorStringMap = {
  brownsugar: "Caramel Coffee",
  jasmine: "Cappuccino",
  oolong: "Black Coffee",
  thai: "Latte",
  taro: "Cold Brew"
};

var iceStringMap = {
  full: "More ice",
  half: "Normal ice",
  none: "No Ice"
};

var sugarStringMap = {
  hundred: "Two packets of sugar",
  fifty: "One packet of sugar",
  twentyfive: "No sugar"
};

var toppingStringMap = {
  bubbles: "Foam",
  grassjelly: "Expresso",
  pudding: "Cream"
}

var priceMap = {
  brownsugar: 4,
  jasmine: 2.5,
  oolong: 2.5,
  thai: 3,
  taro: 3
};

function init() {
  console.log("jQuery ready!");
  var map = convertSearch(window.location.search);
  var name1 = map["name"];
  var name2 = name1.replace("%20"," ");
  var totalPrice = priceMap[map["flavor"]]+.5*toppings.length;
  console.log(JSON.stringify(map));
  console.log(JSON.stringify(toppings));
  $("img").attr("src",imgMap[map["flavor"]]);
  $(".greetings").text("Hi " + name2 + ", thanks for ordering at Starbucks! This is your receipt:");
  $(".flavor").text(flavorStringMap[map["flavor"]]);
  $(".ice").text("Ice Amount: " + iceStringMap[map["ice"]]);
  $(".sugar").text("Sugar Amount: " + sugarStringMap[map["sugar"]]);
  $(".price").text("Price: " + totalPrice + " dollars");
  if (toppings.length == 0){
    $(".toppings").text("Toppings: No Toppings");
  }
  if (toppings.length == 1){
	$(".toppings").text("Toppings: "+toppingStringMap[toppings[0]]);
  }
  if (toppings.length == 2){
	$(".toppings").text("Toppings: "+toppingStringMap[toppings[0]]+", "+toppingStringMap[toppings[1]]);
  }
  if (toppings.length == 3){
	$(".toppings").text("Toppings: "+toppingStringMap[toppings[0]]+", "+toppingStringMap[toppings[1]]+", "+toppingStringMap[toppings[2]]);
  }
}

$(document).ready(init);

function convertSearch(search) {
  var map = {};
  search = search || "";
  search = search.substring(1);
  var pairList = search.split("&");
  for (var i=0; i < pairList.length; i++) {
    var pair = pairList[i];
    var split = pair.split("=");
    var key = split[0];
    var value = split[1];
	if(key == "toppings") {
		toppings.push(value);
	} else {
		map[key] = value;
	}
  }
  return map;
}