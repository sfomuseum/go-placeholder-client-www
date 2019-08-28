var placeholder = placeholder || {};
placeholder.client = placeholder.client || {};

placeholder.client.results = (function(){

    var self = {

	'drawResults': function(map, rows){

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
	    console.log("BOUNDS", bounds);
	    
	    console.log("MAP", map);
	    
	    map.fitBounds(bounds);
	    L.Nextzen.hash({map: map});

	    return true;
	}
    };

    return self;
    
})();
