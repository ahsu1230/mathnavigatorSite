console.log("My Javascript works");

$(document).ready(function() {
	console.log("jQuery ready!");
	
	$(".overlay-wrapper").hide();

	$("li").on("click", function(e) {
		var imgSrc = e.target.src;
		var fullImgSrc = imgSrc.replace("thumb", "full")
		console.log(imgSrc); // <- Check this! What does this give us
		
		/* Your code goes here */
		
		$(".overlay-wrapper").show();
		$(".overlay-content img").attr("src", fullImgSrc);
		
	});

	$(".dismiss").on("click", function() {
		/* Your code goes here */
		
		$(".overlay-wrapper").hide();
		
	});
});