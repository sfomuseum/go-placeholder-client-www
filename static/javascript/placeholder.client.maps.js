var placeholder = placeholder || {};
placeholder.client = placeholder.client || {};

placeholder.client.maps = (function(){

    var refill_style = "/tangram/refill-style.zip";

    var tiles_template = "https://{s}.tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt";

    var attribution = '<a href="https://github.com/tangrams" target="_blank">Tangram</a> | <a href="http://www.openstreetmap.org/copyright" target="_blank">&copy; OpenStreetMap contributors</a> | <a href="https://www.nextzen.org/" target="_blank">Nextzen</a>';
   
    var maps = {};

    var self = {

	'getMap': function(map_el, args){

	    if (! args){
		args = {};
	    }

	    if (! args["api_key"]){
		return null;
	    }

	    var api_key = args["api_key"];
	    
	    var map_id = map_el.getAttribute("id");

	    if (! map_id){
		return;
	    }
	    
	    if (maps[map_id]){
		return maps[map_id];
	    }

	    var attribution = self.getAttribution();
	    var tangram_opts = self.getTangramOptions(args);
	    
	    var nextzen_opts = {
		apiKey: api_key,
		attribution: attribution,
		tangramOptions: tangram_opts,
	    };
	    
	    var map = L.Nextzen.map('map', nextzen_opts);
	    maps[map_id] = map;
	    
	    return map;
	},

	'getTangramOptions': function(args){

	    if (! args){
		args = {};
	    }

	    if (! args["api_key"]){
		return null;
	    }

	    var api_key = args["api_key"];

	    var style_url = refill_style;
	    
	    if (args["url_prefix"]){
		style_url = args["url_prefix"] + style_url;
	    }

	    var tangram_opts = {
		scene: {
		    import: [
			style_url,
		    ],
		    sources: {
			mapzen: {
			    url: tiles_template,
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
