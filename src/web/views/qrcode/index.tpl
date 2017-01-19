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

    <row>
        <div class="table-responsive">
            <table class="table table-bordered">
                <tr>
                    <td id="index1">http://dingying.m.womai.com</td>
                    <td>
                        <div class="btn-group pull-right">
                            <button type="button" class="use btn btn-primary" for="index1">使用</button>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td>
                        <textarea autocomplete="off" data-provide="typeahead" name="url" id="url" placeholder="输入网址" class="form-control"></textarea>
                    </td>
                    <td>
                        <div class="btn-group pull-right">
                            <button id="make" type="button" class="btn btn-primary">生成</button>
                            <button type="button" class="append btn btn-primary" data-prefix="http://">http</button>
                            <button type="button" class="append btn btn-primary" data-www="www">www</button>
                            <button type="button" class="append btn btn-primary" data-suffix=".com">.com</button>
                        </div>
                    </td>
                </tr>
            </table>
        </div>
    </row>
    <div class="row">
        <div id="code">
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
        url && makeQrcode(url)
    });

    function makeQrcode(url) {
        if(url){
            var str = toUtf8(url);
            $("#code").html("");
            $("#code").qrcode({
                correctLevel:0,
                width: 200, //宽度
                height:200, //高度
                text: str //任意内容
            });
        }
    }

    $(".use").bind("click",function () {
        makeQrcode($("#"+$(this).attr("for")).html())
    })
    $(".append").bind("click",function () {
        var prefix=$(this).data("prefix");
        var suffix=$(this).data("suffix");
        var www=$(this).data("www");
        var txt=$("#url").val();
        if(prefix){
            if(txt && txt.indexOf(prefix)>-1){
                return
            }
            txt=prefix+txt
            $("#url").val(txt);
        }
        if(suffix){
            if(txt && txt.lastIndexOf(suffix)>-1){
                return
            }
            txt=txt+suffix
            $("#url").val(txt);
        }

        if(www){
            if(txt && txt.indexOf(www)>-1){
                return
            }
            if(txt && txt.indexOf("//")>-1){
                arr=txt.split("//")
                arr[1]="www."+arr[1]
                txt = arr.join("//")
            }else{
                txt="www."+txt
            }
            $("#url").val(txt);
        }
    })
</script>
</body>
</html>
