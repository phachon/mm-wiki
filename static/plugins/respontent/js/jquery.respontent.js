/*	
 * jQuery respontent v1.2.2
 * @requires jQuery 1.5.0 or later
 *
 * respontent.frebsite.nl
 *	
 * Copyright (c) Fred Heusschen
 * www.frebsite.nl
 *
 * Licensed under the MIT license.
 * http://en.wikipedia.org/wiki/MIT_License
 */


(function( $ ) {

	var _PLUGIN_	= 'respontent',
		_ABBR_		= 'respontent',
		_VERSION_	= '1.2.2';


	//	Plugin already excists
	if ( $[ _PLUGIN_ ] )
	{
		return;
	}


	//	Global variables
	var _c = {},
		_pluginInitiated = false;


	/*
		Class
	*/
	$[ _PLUGIN_ ] = function( $wrapper, opts )
	{
		this.$wrapper	= $wrapper;
		this.opts  		= opts;

		for ( var a = 0; a < $[ _PLUGIN_ ].media.length; a++ )
		{
			if ( typeof this[ 'fit_' + $[ _PLUGIN_ ].media[ a ] ] == 'function' )
			{
				this[ 'fit_' + $[ _PLUGIN_ ].media[ a ] ]();
			}
		}

		return this;
	};

	$[ _PLUGIN_ ].prototype = {

		wrapInParent: function( $t, iClass )
	    {
	        var $p = $t.parent();
	
	        if ( $p.children().length == 1 )
	        {
	            $p.addClass( iClass );
	        }
	        else
	        {
	            $t.wrap( '<div class="' + iClass + '" />' );
	        }
	    }

	};

	$[ _PLUGIN_ ].defaults 	= {};
	$[ _PLUGIN_ ].media 	= [];
	$[ _PLUGIN_ ].version 	= _VERSION_;	


	/*
		jQuery Plugin
	*/
	$.fn[ _PLUGIN_ ] = function( opts )
	{
		//	First time plugin is fired
		if ( !_pluginInitiated )
		{
			_initiatePlugin();
		}

		//	Extend options
		opts = $.extend( true, {}, $[ _PLUGIN_ ].defaults, opts );

		return this.each(
			function()
			{
				var $node = $(this);
				if ( $node.data( _PLUGIN_ ) )
				{
					return;
				}
				$node.data( _PLUGIN_, new $[ _PLUGIN_ ]( $node, opts ) );
			}
		);
	};


	/*
		Private functions
	*/
	function _initiatePlugin()
	{
		_pluginInitiated = true;

		_c = function( c ) { return _ABBR_ + '-' + c; };
		_c.add = function( c )
		{
			c = c.split( ' ' );
			for ( var d in c )
			{
				_c[ c[ d ] ] = _c( c[ d ] );
			}
		};

		$[ _PLUGIN_ ]._c = _c;
	}

})( jQuery );