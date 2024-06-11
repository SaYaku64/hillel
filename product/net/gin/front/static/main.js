$("document").ready(() => {
    var sessionCookie = Cookies.get('sessionID');
    if (sessionCookie == undefined || sessionCookie == "") {
        $.get( "/api/v1/getSessionID", function( data ) {
            Cookies.set('sessionID', data, { expires: 365});
            document.getElementById("navSessionID").innerHTML = data;
        });
    } else {
        document.getElementById("navSessionID").innerHTML = sessionCookie;
    }

    $("#btnName").click(() => {
        $.get( "/api/v1/alert", function( data ) {
            alert(data)
        });
    });
});
