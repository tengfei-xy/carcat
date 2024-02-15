$('#submit').on('click', function() {
    let from = {}
    from.action = $.act.Login_user()
    from.loginid = $("#loginid").text()
    window.location.href = $.page.index() + "?id=" + from.loginid
    return
    $.post_send.ajax(form, function(res) {
        if (res.status != 0) {
            $.toast.text(res.explain)
            return
        } else {
            window.location.href = $.page.index() + "?id=" + from.loginid
        }
    })
})