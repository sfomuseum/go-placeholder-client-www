window.addEventListener("load", function load(event){

    var rows = document.getElementsByClassName("result");
    var count = rows.length;

    if (! count){
	return;
    }
        
    var api_key = document.body.getAttribute("data-nextzen-api-key");
    var url_prefix = document.body.getAttribute("data-url-prefix");
    var is_apigw = document.body.getAttribute("data-is-api-gateway");    
    
    if (! api_key){
	console.log("Missing API key");
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	console.log("Missing map element");	
	return;
    }

    var map_args = {
	"api_key": api_key,
    };

    if (url_prefix != ""){
	map_args["url_prefix"] = url_prefix;
    }

    if (is_apigw != ""){
	map_args["is_api_gateway"] = 1;
    }
    
    // we need to do this _before_ Tangram starts trying to draw things
    map_el.style.display = "block";
    
    var map = placeholder.client.maps.getMap(map_el, map_args);

    if (! map){
	console.log("Unable to instantiate map");
	return;
    }
    
    if (! placeholder.client.results.drawResultsMap(rows, map)){
	map_el.style.display = "none";
	console.log("Failed to draw results maps");
	return;
    }

    var features_layer = placeholder.client.results.drawResultsFeatures(rows, map);

    if (! features_layer){
	console.log("Failed to render features layer");
	return;
    }

    placeholder.client.results.assignHoverEvents(rows, map, features_layer);
    
}, false);
