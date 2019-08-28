window.addEventListener("load", function load(event){

    var rows = document.getElementsByClassName("result");
    var count = rows.length;

    if (! count){
	return;
    }
        
    var api_key = document.body.getAttribute("data-nextzen-api-key");

    if (! api_key){
	console.log("Missing API key");
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	console.log("Missing map element");	
	return;
    }

    // we need to do this _before_ Tangram starts trying to draw things
    map_el.style.display = "block";
    
    var map = placeholder.client.maps.getMap(map_el, api_key);

    if (! map){
	console.log("Unable to instantiate map");
	return;
    }
    
    if (! placeholder.client.results.drawResultsMap(map, rows)){
	map_el.style.display = "none";
	console.log("Failed to draw results maps");
	return;
    }

    var features_layer = placeholder.client.results.drawResultsFeatures(map, rows);

    if (! features_layer){
	console.log("Failed to render features layer");
	return;
    }

    placeholder.client.results.assignHoverEvents(rows, map, features_layer);
    
}, false);
