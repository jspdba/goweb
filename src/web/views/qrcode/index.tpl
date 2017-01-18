<!DOCTYPE html>
<html lang="zh-cn">
<style>
    body {
        padding-top: 70px;
        padding-bottom: 50px;
    }
</style>
{{template "common/header_flat.tpl"}}
<title>二维码</title>
<body>
{{template "common/navibar.tpl"}}

<div class="container">
    <h4>二维码处理</h4>
    <div class="row">
        <div class="box">
            <div class="col-lg-12 text-center" role="form">
                <div class="form-group">
                    <label for="url">网址</label>
                    <textarea autocomplete="off" data-provide="typeahead" name="url" id="url" placeholder="输入网址" class="form-control"></textarea>
                </div>
                <input id="make" type="button" class="btn btn-primary" value="生成">
            </div>
            <div id="code">
            </div>
        </div>
    </div>
</div>

{{template "common/script.tpl"}}
<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>
<script src="/static/js/jquery.qrcode.min.js"></script>
<script>
    /**
     * 转换中文字符,生成二维码
     * @param str
     * @returns {string}
     */
    function toUtf8(str) {
        var out, i, len, c;
        out = "";
        len = str.length;
        for(i = 0; i < len; i++) {
            c = str.charCodeAt(i);
            if ((c >= 0x0001) && (c <= 0x007F)) {
                out += str.charAt(i);
            } else if (c > 0x07FF) {
                out += String.fromCharCode(0xE0 | ((c >> 12) & 0x0F));
                out += String.fromCharCode(0x80 | ((c >>  6) & 0x3F));
                out += String.fromCharCode(0x80 | ((c >>  0) & 0x3F));
            } else {
                out += String.fromCharCode(0xC0 | ((c >>  6) & 0x1F));
                out += String.fromCharCode(0x80 | ((c >>  0) & 0x3F));
            }
        }
        return out;
    }

    $("#make").bind("click",function () {
        var url=$("#url").val();
        if(url){
            //开始生成二维码
            var str = toUtf8(url);
//        $('#code').qrcode(str);
            $("#code").html("");
            $("#code").qrcode({
//            render: "table", //table方式
                correctLevel:0,
                width: 200, //宽度
                height:200, //高度
                text: str //任意内容
            });
        }
    })
</script>
</body>
</html>
