$(function () {
    $('.spinner input').bind('keyup', function () {
        this.value = this.value.replace(/[^\d]/g,'');
    });
    $('.spinner input').bind('afterpaste', function () {
        this.value = this.value.replace(/[^\d]/g,'')
    });
    $('.spinner .btn:first-of-type').on('click', function() {
        var val = $('.spinner input').val();
        if(val == "") {
            $('.spinner input').val(0);
        }
        $('.spinner input').val( parseInt($('.spinner input').val(), 10) + 1);
    });
    $('.spinner .btn:last-of-type').on('click', function() {
        var val = $('.spinner input').val();
        if(val == "") {
            $('.spinner input').val(0);
        }
        var n = parseInt($('.spinner input').val(), 10) - 1;
        if (n <= 0) {
            $('.spinner input').val(0);
        }else {
            $('.spinner input').val(n);
        }
    });
});