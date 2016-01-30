var form = $("#SignupForm");
var un = $("#inputUsername");
var p1 = $("#inputPassword");
var p2 = $("#inputPasswordCfm");

un.on('input', (function() {
	$("#helpBlock1").html('')
	console.log(un.val())
	$.ajax({
	    url:"/api/checkusername",
	    data: un.val(),
	    type: "POST",
	    success: function(data){
	        console.log(data)
	        if (data == "true"){
	        	un.parent().removeClass("has-success").addClass("has-error")
	        	$("#helpBlock1").html('Username taken.');
	      	} else {
	            un.parent().removeClass("has-error").addClass("has-success")

   	        }
	    }

	});
}))

form.submit(function(e){
	var ok = validatePasswords();
	console.log(ok)
	if ( !ok ) {
		p1.val("");
        p2.val("");
        e.preventDefault();
    }
	// e.preventDefault();
});

function validatePasswords() {
        $("#helpBlock2").text("");
        if (p1.val() == '' || p2.val() == '') {
            $("#helpBlock2").text('Enter a password.');
            return false;
        }
        if ( p1.val().length < 1) {
        	$("#helpBlock2").html('Password cannot be blank');
            return false;
        }
        if (p1.val() !== p2.val()) {
            $("#helpBlock2").html('Your passwords did not match.<br/>Please re-enter your passwords.');
            return false;
        }
        if ( (un.parent().hasClass("has-error") === true) ) {
        	$("#helpBlock2").html('Please use a validate username');
        	return false;
        }
        return true;
    };
