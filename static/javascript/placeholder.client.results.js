var placeholder = placeholder || {};
placeholder.client = placeholder.client || {};

placeholder.client.results = (function(){

    var marker_opts = {
	radius: 8,
	fillColor: "#ff7800",
	color: "#000",
	weight: 1,
	opacity: 1,
	fillOpacity: 0.8
    };

    var focus;
    
    var self = {

	'drawResultsMap': function(map, rows){

	    var count = rows.length;

	    if (! count){
		return false;
	    }

	    var minx = 0.0;
	    var miny = 0.0;
	    var maxx = 0.0;
	    var maxy = 0.0;
	    
	    for (var i=0; i < count; i++){
		
		var row = rows[i];
		var bbox = row.getAttribute("data-bounding-box");
		
		if (! bbox){
		    continue;
		}
		
		bbox = bbox.split(",");
		
		if (bbox.length != 4){
		    continue;
		}
		
		var row_minx = bbox[0];
		var row_miny = bbox[1];
		var row_maxx = bbox[2];
		var row_maxy = bbox[3];
		
		minx = Math.min(minx, row_minx);
		miny = Math.min(miny, row_miny);
		maxx = Math.max(maxx, row_maxx);
		maxy = Math.max(maxy, row_maxy);
	    }
	    
	    var sw = [ miny, minx ];
	    var ne = [ maxy, maxx ];
	    
	    var bounds = [ sw, ne ];
	    map.fitBounds(bounds);

	    // maybe?
	    // L.Nextzen.hash({map: map});
	    
	    return true;
	},

	'drawResultsFeatures': function(map, rows){

	    var feature_collection = self.resultsAsFeatureCollection(rows);

	    var layer = L.geoJSON(feature_collection, {
		pointToLayer: function (feature, latlng) {
		    return L.circleMarker(latlng, marker_opts);
		},
		onEachFeature: function(feature, layer){
		    
		    if (feature.properties){
			
			var wofid = feature.properties["wof:id"];
			var name = feature.properties["wof:name"];
			var placetype = feature.properties["wof:placetype"];

			var popup_text = name + " (" + wofid + ") is a " + placetype;
			layer.bindPopup(popup_text);
		    }
		}
	    });
	    
	    layer.addTo(map);
	    return layer;
	},

	'assignHoverEvents': function(rows, map, features_layer){

	    var wof_layer;
	    
	    var panTo = function(e){
		
		var el = e.target;
		
		var wof_id = el.getAttribute("data-whosonfirst-id");
		var wof_el;
		
		if (wof_id){
		    wof_el = el;
		}
		
		else {
		    
		    var parent_el = el.parentNode;
		    wof_id = parent_el.getAttribute("data-whosonfirst-id")
		    
		    if (wof_id){
			wof_el = parent;
		    }
		}
		
		if (! wof_el){
		    console.log("Unable to determine WOF element");
		    return;
		}
		
		var latitude_id = "result-" + wof_id + "-latitude";
		var latitude_el = document.getElementById(latitude_id);
		
		if (! latitude_el){
		    console.log("Missing latitude element", latitude_id);
		    return;
		}
		
		var longitude_id = "result-" + wof_id + "-longitude";
		var longitude_el = document.getElementById(longitude_id);
		
		if (! longitude_el){
		    console.log("Missing longitude element", longitude_id);
		    return;
		}
		
		var lat = parseFloat(latitude_el.innerText);
		var lon = parseFloat(longitude_el.innerText);		
		
		if ((! lat) || (lat == NaN)){
		    console.log("Invalid latitude", latitude_el.innerText);
		    return;
		}
		
		if ((! lon) || (lon == NaN)){
		    console.log("Invalid longitude", longitude_el.innerText);
		    return;
		}

		if (wof_layer){
		    map.removeLayer(wof_layer);
		}
		
		map.flyTo([lat, lon], 8);
		focus = wof_id;
		
		var wof_url = whosonfirst.uri.id2abspath(wof_id);

		var on_success = function(f){
		    
		    var f_props = f["properties"];
		    var f_wofid = f_props["wof:id"];

		    if (f_wofid != focus){
			return;
		    }

		    var f_geom = f["geometry"];

		    if (f_geom["type"] == "Point"){
			return;
		    }
		    
		    var bbox = whosonfirst.geojson.derive_bbox(f);

		    var minx = bbox[0];
		    var miny = bbox[1];
		    var maxx = bbox[2];
		    var maxy = bbox[3];		    

		    var sw = [ miny, minx ];
		    var ne = [ maxy, maxx ];
		    var bounds = [ sw, ne ];
		    
		    wof_layer = L.geoJSON(f);
		    wof_layer.addTo(map);
		    
		    features_layer.bringToFront();

		    map.fitBounds(bounds, {
			"padding": [ 50, 50 ]
		    });
		};

		var on_fail = function(rsp){
		    console.log("FAIL", rsp);
		};
		
		whosonfirst.net.fetch(wof_url, on_success, on_fail);
	    };

	    var count = rows.length;
	    
	    for (var i=0; i < count; i++){
		
		var row = rows[i];
		var wof_id = row.getAttribute("data-whosonfirst-id");
		
		if (! wof_id){
		    continue;
		}
		
		var wofid_id = "result-" + wof_id + "-id";
		var wofid_el = document.getElementById(wofid_id);
		
		if (wofid_el){
		    wofid_el.onclick = panTo;
		}
		
	    }
	},
	
	'resultsAsFeatureCollection': function(rows){

	    var features = [];
	    
	    var count = rows.length;
	    
	    for (var i=0; i < count; i++){

		var row = rows[i];
		var wof_id = row.getAttribute("data-whosonfirst-id");

		if (! wof_id){
		    console.log("Row is missing data-whosonfirst-id attribute");
		    continue;
		}

		var name_id = "result-" + wof_id + "-name";
		var name_el = document.getElementById(name_id);

		if (! name_el){
		    console.log("Missing name element", name_id);
		    continue;
		}
		
		var placetype_id = "result-" + wof_id + "-placetype";
		var placetype_el = document.getElementById(placetype_id);

		if (! placetype_el){
		    console.log("Missing placetype element", placetype_id);
		    continue;
		}
		    
		var latitude_id = "result-" + wof_id + "-latitude";
		var latitude_el = document.getElementById(latitude_id);

		if (! latitude_el){
		    console.log("Missing latitude element", latitude_id);
		    continue;
		}
		
		var longitude_id = "result-" + wof_id + "-longitude";
		var longitude_el = document.getElementById(longitude_id);

		if (! longitude_el){
		    console.log("Missing longitude element", longitude_id);
		    continue;
		}

		var lat = parseFloat(latitude_el.innerText);
		var lon = parseFloat(longitude_el.innerText);		

		if ((! lat) || (lat == NaN)){
		    console.log("Invalid latitude", latitude_el.innerText);
		    continue;
		}

		if ((! lon) || (lon == NaN)){
		    console.log("Invalid longitude", longitude_el.innerText);
		    continue;
		}

		if (! lat){
		    continue;
		}
		
		if (! lon){
		    continue;
		}

		var name = name_el.innerText;
		var placetype = placetype_el.innerText;		
		
		var coords = [lon, lat];
		
		var geom = {
		    "type": "Point",
		    "coordinates": coords,
		};

		var props = {
		    "wof:id": wof_id,
		    "wof:placetype": placetype,
		    "wof:name": name,
		};

		var feature = {
		    "type": "Feature",
		    "geometry": geom,
		    "properties": props,
		};

		features.push(feature);
	    }

	    var collection = {
		"type": "FeatureCollection",
		"features": features,
	    };

	    return collection;
	}
	
    };

    return self;
    
})();
