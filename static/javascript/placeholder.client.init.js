window.addEventListener("load", function load(event){

    var api_key = document.body.getAttribute("data-nextzen-api-key");

    if (! api_key){
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	return;
    }
    
    var map = placeholder.client.maps.getMap(map_el, api_key);

    if (! map){
	return;
    }

    var rows = document.getElementsByClassName("result");
    
    if (! placeholder.client.results.drawResults(map)){
	return;
    }
    
    map_el.style.display = "block";
    
}, false);
