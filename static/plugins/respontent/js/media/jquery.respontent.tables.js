/*	
 * jQuery respontent tables media
 */


(function( $ ) {

	var _PLUGIN_	= 'respontent',
		_MEDIA_		= 'tables';

	var _c = {},
		_mediaInitiated = false;


	$[ _PLUGIN_ ].prototype[ 'fit_' + _MEDIA_ ] = function( $table, type )
	{
		var that = this;

		if ( !_mediaInitiated )
		{
			_initMedia();
		}


		//	Single table
		if ( $table )
		{
			//	Find type
			if ( typeof type == 'undefined' )
			{
				type = this.opts[ _MEDIA_ ];
			}
			if ( typeof type == 'function' )
			{
				type = type.call( $table[ 0 ] );
			}
			if ( typeof type == 'undefined' )
			{
				type = getType.call( $table[ 0 ] );
			}

            if ( !$table.find( 'tbody' ).length )
            {
                $table.wrapInner( '<tbody />' );
            }

			//	Apply fix
			if ( type )
			{
				switch( type )
				{
					case 'stack':
						$table.addClass( _c[ 'table-stack' ] );
						break;
					
					case 'scroll':
						this.wrapInParent( $table, _c[ 'table-scroll' ] );
						break;
					
					case 'list':
						$table.addClass( _c[ 'table-list' ] );
	
						var $tr = $('tr', $table),
			                $td = $table.children( 'tbody' ).find( 'td' ),
			                $th = $table.children( 'thead' ).find( 'th, td' ),
			                txt = '';
			
			            $th.each(
			                function( i )
			                {
			                	txt = $.trim( $(this).text() );
			                	var $cur = $td.filter( ':nth-child(' + ( i + 1 ) + ')' );
		
			                	if ( txt.length )
			                	{
				                    $cur.attr( 'data-respontent-title', txt );
			                	}
			                	else
			                	{
				                	$cur.addClass( _c[ 'table-list-title' ] );
			                	}
			                }
			            );
						break;
				}
			}
		}

		//	Find all tables
		else
		{	
		    this.$wrapper
		    	.find( 'table' )
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
//	$[ _PLUGIN_ ].defaults[ _MEDIA_ ] = false / 'scroll' / 'stack' / 'list' / function;


	//	Add to plugin
	$[ _PLUGIN_ ].media.push( _MEDIA_ );


	//	Private functions
	function getType()
	{
		if (
			$(this).find( 'th' ).length || 
			$(this).find( 'thead' ).length || 
			$(this).find( 'tfoot' ).length
		) {
			return 'scroll';
		}
		else
		{
			return 'stack';
		}
	}
	function _initMedia()
	{
		_mediaInitiated = true;
		_c = $[ _PLUGIN_ ]._c;
		_c.add( 'table-stack table-scroll table-list table-list-title' );
	}

})( jQuery );