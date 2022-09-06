window.addEventListener("load", function load(event){

    var placeholder_el = document.getElementById("placeholder");
    var search_button = document.getElementById("search-button");
    
    var ready_check = placeholder_el.getAttribute("data-enable-ready-check");
    var ready_check_url = placeholder_el.getAttribute("data-ready-check-url");    

    if (ready_check == "true"){

	var placeholder_ready_cb = function(e){
	    
	    var rsp = e.target;
	    var status = rsp.status;
	    
	    // console.log("STATUS", status);
	    
	    if (status == 200){
		console.log("Placeholder is running and accepting connections");
		search_button.innerText = "Search";
		search_button.removeAttribute("disabled");
		return;
	    }
	    
	    if (status == 503){
		setTimeout(placeholder_ready, 2500);
		return;
	    }
	    
	    console.log("Unable to determine Placeholder status", rsp);
	    return false;
	};
	
	var placeholder_ready = function(){
	    var req = new XMLHttpRequest();
	    req.addEventListener("load", placeholder_ready_cb);
	    req.open("GET", ready_check_url, true);
	    req.send();
	};
	
	placeholder_ready();
    }
    
    var rows = document.getElementsByClassName("result");
    var count = rows.length;

    if (! count){
	return;
    }
        
    var api_key = document.body.getAttribute("data-nextzen-api-key");
    var style_url = document.body.getAttribute("data-nextzen-style-url");
    var tile_url = document.body.getAttribute("data-nextzen-tile-url");    
    
    if (! api_key){
	console.log("Missing API key");
	return;
    }

    if (! style_url){
	console.log("Missing style URL");
	return;
    }

    if (! tile_url){
	console.log("Missing tile URL");
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	console.log("Missing map element");	
	return;
    }

    var map_args = {
	"api_key": api_key,
	"style_url": style_url,
	"tile_url": tile_url,
    };

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
