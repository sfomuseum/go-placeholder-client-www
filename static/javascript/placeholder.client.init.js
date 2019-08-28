window.addEventListener("load", function load(event){

    var api_key = document.body.getAttribute("data-nextzen-api-key");

    if (! api_key){
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	return;
    }

    map_el.style.display = "block";
    
    var map = placeholder.client.maps.getMap(map_el, api_key);

    if (! map){
	return;
    }

    var zoom = 14;
    var lat = 37.6185;
    var lon = -122.3829;

    map.setView([ lat, lon ], zoom);
    
    var rows = document.getElementsByClassName("result");
    
    if (! placeholder.client.results.drawResultsMap(map, rows)){
	return;
    }

    placeholder.client.results.drawResultsFeatures(map, rows)
    
}, false);
