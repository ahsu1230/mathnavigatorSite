console.log("My Javascript works");

var database = JSON.parse(data);

$(document).ready(function() {
	console.log("jQuery ready!");
	$(".search-form button").on("click", function() {
		var searchText = $(".search-form input").val();
		console.log("search " + searchText);
		
		$("ul").empty();
		
		var numResults = 0;
		
		for (var i=0; i<database.length; i++) {
			var vehicleId = database[i].id;
			var vehicleName = database[i].name.english;
			var vehicleImg = getThumbnailFromId(vehicleId);
			var vehicleInfo = database[i].info.moreinfo;
			if (matches(vehicleName, searchText)) {
				createRow(vehicleId, vehicleName, vehicleImg, vehicleInfo);
				numResults = 1 + numResults;
			}
		}
		numOfVehicle(numResults);
	});
});

function numOfVehicle (numResults) {
	if (numResults == 0) {
		$("h3").text("No Vehicles Found");
	}
	else {
		$("h3").text(numResults + " Vehicle(s) found");
	}
}

function createRow (vehicleId, vehicleName, vehicleImg, vehicleInfo) {
	
	var row = $(document.createElement("li"));
	var span1 = $(document.createElement("span"));
	span1.text(vehicleId);
	row.append(span1);
	var pokeImg = $(document.createElement("img"));
	pokeImg.attr("src", vehicleImg);
	row.append(pokeImg);
	var span2 = $(document.createElement("span"));
	span2.text(vehicleName);
	row.append(span2);
	var span3 = $(document.createElement("span"));
	span3.text(vehicleInfo);
	row.append(span3);
	$("ul").append(row);
}

function getThumbnailFromId(id) {
	var idDigit = ('000' + id).substr(-3);
	return "../assets/thumbnails/" + idDigit + ".jpg";
}

function matches(text1, text2) {
	if (!text1 && !text2) {
		return false;
	}
	var textLC1 = text1.toLowerCase();
	var textLC2 = text2.toLowerCase();
	return textLC1.indexOf(textLC2) >= 0;
}