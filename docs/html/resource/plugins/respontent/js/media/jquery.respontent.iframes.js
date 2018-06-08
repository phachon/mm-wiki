/*	
 * jQuery respontent iframs media
 */


(function( $ ) {

	var _PLUGIN_	= 'respontent',
		_MEDIA_		= 'iframes';

	var _c = {},
		_mediaInitiated = false;


	$[ _PLUGIN_ ].prototype[ 'fit_' + _MEDIA_ ] = function( $iframe, type )
	{
		var that = this;

		if ( !_mediaInitiated )
		{
			_initMedia();
		}


		//	Single iframe
		if ( $iframe )
		{
			//	Find type
			if ( typeof type == 'undefined' )
			{
				type = this.opts[ _MEDIA_ ];
			}
			if ( typeof type == 'function' )
			{
				type = type.call( $iframe[ 0 ] );
			}
			if ( typeof type == 'undefined' )
			{
				type = getType.call( $iframe[ 0 ] );
			}

			//	Apply fix
			if ( type )
			{
				switch( type )
				{
					case 'square':
					case 'youtube':
						this.wrapInParent( $iframe, _c[ 'iframe-' + type ] );
						break;
				}
			}
		}

		//	Find all iframes
		else
		{	
		    this.$wrapper
		    	.find( 'iframe' )
		    	.each(
			        function()
			        {
			        	that[ 'fit_' + _MEDIA_ ]( $(this) );
			        }
			    );
		}

		return this;
	};


	//	Options
//	$[ _PLUGIN_ ].defaults[ _MEDIA_ ] = false / 'square' / 'yuotube' / 'maps' / function;


	//	Add to plugin
	$[ _PLUGIN_ ].media.push( _MEDIA_ );


	//	Private functions
	function getType()
	{
		if ( 
			$(this).is( '[src*="youtube.com"]' ) ||
			$(this).is( '[src*="youtube-nocookie.com"]' )
		) {
			return 'youtube';
		}
		else
		{
			return 'square';
		}
	}
	function _initMedia()
	{
		_mediaInitiated = true;
		_c = $[ _PLUGIN_ ]._c;
		_c.add( 'iframe-square iframe-youtube' );
	}

})( jQuery );