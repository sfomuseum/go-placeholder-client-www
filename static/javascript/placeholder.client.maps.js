var placeholder = placeholder || {};
placeholder.client = placeholder.client || {};

placeholder.client.maps = (function(){

    var attribution = '<a href="https://github.com/tangrams" target="_blank">Tangram</a> | <a href="http://www.openstreetmap.org/copyright" target="_blank">&copy; OpenStreetMap contributors</a> | <a href="https://www.nextzen.org/" target="_blank">Nextzen</a>';
    
    var maps = {};

    var self = {

	'getMap': function(map_el, api_key){

	    var map_id = map_el.getAttribute("id");

	    if (! map_id){
		return;
	    }
	    
	    if (maps[map_id]){
		return maps[map_id];
	    }

	    var attribution = self.getAttribution();
	    var tangram_opts = self.getTangramOptions(api_key);
	    
	    var nextzen_opts = {
		apiKey: api_key,
		attribution: attribution,
		tangramOptions: tangram_opts,
	    };
	    
	    var map = L.Nextzen.map('map', nextzen_opts);
	    maps[map_id] = map;
	    
	    return map;
	},

	'getTangramOptions': function(api_key){

	    var tangram_opts = {
		scene: {
		    import: [
			'/tangram/refill-style.zip',	// something something something prefixes
		    ],
		    sources: {
			mapzen: {
			    url: 'https://{s}.tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt',
			    url_subdomains: ['a', 'b', 'c', 'd'],
			    url_params: {api_key: api_key},
			    tile_size: 512,
			    max_zoom: 16
			}
		    }
		}
	    };

	    return tangram_opts;
	},

	'getAttribution': function(){
	    return attribution;
	},
    };

    return self;
    
})();
