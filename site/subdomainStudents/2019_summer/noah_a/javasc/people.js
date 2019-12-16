console.log("My Javascript works");

$(document).ready(function() {
  console.log("jQuery ready!");

  $("h3").on("click", function(event) {
	  var target = event.target;
	  var siblings = $(target).siblings("ul");
	  console.log(target);
	  console.log(siblings);
	  if ($(siblings).hasClass("show")) {
		$(siblings).removeClass("show")
		} 
	else {
		$(siblings).addClass("show")
		}
  });	
});