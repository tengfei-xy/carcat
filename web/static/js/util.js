pathstr = window.location.pathname
pth = pathstr.substring(1, pathstr.indexOf("/", 2))
pd = pth.substring(0, 4)
const baseurl = window.location.origin + "/" + pth + "/"
const url = baseurl + "carcat"

$.page = {
    index: function() { return baseurl + "index.html" },
    get: function() { return baseurl + "get.html" },
    db: function() { return baseurl + "db.html" },
}
$.act = {
    Login_user: function() { return "login-user-" }
}
$.post_send = {
    // ajax封装,https://www.jianshu.com/p/26348205b871

    ajax: function(data, success, error) {

        //console.log("post_data", data)
        const type = 'post'; //请求类型
        const async = true; //异步请求
        const alone = false; //独立提交（一次有效的提交）
        var success = success || function(data) {
            console.log('请求成功');
            setTimeout(function() {
                $.toast.text(data.msg); //通过layer插件来进行提示信息
            }, 500);
            if (data.status) { //服务器处理成功
                setTimeout(function() {
                    if (data.url) {
                        location.replace(data.url);
                    } else {
                        location.reload(true);
                    }
                }, 1500);
            }
            // else{//服务器处理失败
            //     if(alone){//改变ajax提交状态
            //         ajaxStatus = true;
            //     }
            // }
        };
        var error = error || function(data) {
            console.error('请求成功失败', data);
            /*data.status;//错误状态吗*/
            setTimeout(function() {
                if (data.status == 404) {
                    $.toast.text('请求失败，请求未找到');
                } else if (data.status == 503) {
                    $.toast.text('请求失败，服务器内部错误');
                } else {
                    $.toast.text('请求失败,网络连接超时');
                }
                //ajaxStatus = true;
            }, 500);
        };
        //ajaxStatus = false;//禁用ajax请求
        /*正常情况下1秒后可以再次多个异步请求，为true时只可以有一次有效请求（例如添加数据）*/
        // if(!alone){
        //     setTimeout(function () {
        //         ajaxStatus = true;
        //     },1000);
        // }
        $.ajax({
            'url': url,
            'data': JSON.stringify(data),
            'type': type,
            "contentType": "application/json;charset=utf-8",
            'dataType': "json",
            'success': success,
            'error': error
        });
    }
}
$.toast = {
    text: function(data) {
        $("#weui-toast__content").text(data)
        var $textToast = $('#textToast');
        if ($textToast.css('display') != 'none') return;

        $textToast.fadeIn(100);
        setTimeout(function() {
            $textToast.fadeOut(100);
        }, 2000);

    }
}
$.append = {
    table_cell: function(data) {
        let div = $('<div class="weui-flex__item weui-tabbar__item table_plus"></div>')
        let div_in_p = $('<p class="table_text_plus">' + data + '</p>')
        div.append(div_in_p)

        return div
    },
    table_small_cell: function(data) {
        return $('<div class="weui-flex__item table_small_space"></div>').text(data)
    },
    table: function(tableid, a, b, c) {
        let $tableid = $(tableid)

        let div = $('<div class="weui-flex"></div>')
        div.append(a, b)
        if (typeof(c) != "undefined") {
            div.append(c)
        }
        $tableid.append(div)

    }
}
$.time = {
    current_time_date: function() {
        return current_date() + " " + current_time()
    },
    current_date: function() {
        let d = new Date()
        let year = d.getFullYear()
        let month = d.getMonth() + 1
        let day = d.getDate()
        month = ('' + month).length === 1 ? '0' + month : '' + month
        day = ('' + day).length === 1 ? '0' + day : '' + day
        return year + "-" + month + "-" + day
    },
    current_time: function() {
        let d = new Date()
        let h = d.getHours()
        let m = d.getMinutes()
        h = ('' + h).length === 1 ? '0' + h : '' + h
        m = ('' + m).length === 1 ? '0' + m : '' + m

        return h + ":" + m
    },
    transTimestamp: function(date, time) {
        let d = date + " " + time + ":00.0"
        d = d.substring(0, 19);
        d = d.replace(/-/g, '/');
        return String(new Date(d).getTime())
    },
    wxpickerdate: function(id, name) {
        weui.datePicker({
            start: 2020,
            end: new Date().getFullYear(),
            onConfirm: function(result) {
                $(id).text($.time.wxtransdate(result))
            },
            //onChange: function (result) {},
            title: name
        });
    },
    wxpickertime: function(id, name) {
        let hours = new Array(),
            minites = new Array()
        for (var i = 0; i < 24; i++) {
            var hours_item = {};
            hours_item.label = ('' + i).length === 1 ? '0' + i : '' + i;
            hours_item.value = j;
            hours.push(hours_item);
        }
        for (var j = 0; j < 60; j++) {
            var minites_item = {};
            minites_item.label = ('' + j).length === 1 ? '0' + j : '' + j;
            minites_item.value = j;
            minites.push(minites_item);
        }

        weui.picker(hours, minites, {
            defaultValue: [new Date().getHours(), new Date().getMinutes()],
            onConfirm: function(result) {
                $(id).text($.time.wxtranstime(result))
            },
            title: name
        });
    },
    wxtransdate: function(e) {
        y = e[0].value
        m = e[1].value
        d = e[2].value
        m = ('' + m).length === 1 ? '0' + m : '' + m;
        d = ('' + d).length === 1 ? '0' + d : '' + d;
        return y + "-" + m + "-" + d
    },
    wxtranstime: function(e) {
        return e[0].value + ":" + e[1].value
    },
    wxcenturytime: function(id, name) {
        weui.datePicker({
            start: 1900,
            end: new Date().getFullYear(),
            onConfirm: function(result) {
                $(id).text($.time.wxtransdate(result))
            },
            //onChange: function (result) {},
            title: name
        });
    },
    keepnew: function() {
        return "?_=1480405186452"
    },
    keepnewt: function() {
        return "&_=1480405186452"
    }
}
$.auto = {
    choose: function(id, data) {
        switch (data) {
            case "男":
                $(id).eq(0).attr("checked", true)
                break;
            case "女":
                $(id).eq(1).attr("checked", true)
                break;
        }
        if (data == "") { return }
        if (typeof(data) === "undefined") {
            $.toast.text("数据获取失败")
        }
        let e = data.replace("diy", "0").split(',')
        for (let i = 0; i <= e.length; i++) {

            let seq = parseInt(e[i]) - 1
            $(id).eq(seq).attr("checked", true)
        }
    },
    choose_diy: function(id, data, diy) {
        if (data == "" || diy == "") { return }

        if (typeof(data) === "undefined" || typeof(data) === "diy") {
            $.toast.text("数据获取失败")
        }
        if (data.indexOf("diy") != -1) {
            $(id).val(diy)
        }
    },
    array: function(id) {
        let r = ""
        $(id).each(function() {
            r += $(this).val() + ","
        })
        if (r != "") {
            r = r.substr(0, r.length - 1)
        }
        return r
    }
}
$.check = {
    isString: function(str) {
        if (str = "") { return true }
        if (typeof(str) == "undefined" || str.length == 0) { return false }

        for (var i in str) {
            var asc = str.charCodeAt(i);
            if ((asc >= 65 && asc <= 90 || asc >= 97 && asc <= 122)) {
                return true;
            }
        }
        return false;
    },
    isNumber: function(str) {
        let n = Number(str);
        if (!isNaN(n)) {
            return true;
        }
        return false
    },
}
$.download = {
    filename: function(a, b, c) {
        return a + "__" + b + "__" + c + ".xlsx"
    }
}
$.te = {
    born: function() { if (pd == "c213" || pd == "24gf" || pd == "423n") { return true } else { return false } }
}
$('html').css("user-select", "none")
$('input').attr("autocomplete", "off")