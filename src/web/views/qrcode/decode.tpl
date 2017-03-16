<!DOCTYPE html>
<html lang="zh-cn">
<style>
    body {
        padding-top: 70px;
        padding-bottom: 50px;
    }
</style>
{{template "common/header_flat.tpl"}}
<title>二维码反向解析</title>
<body>
{{template "common/navibar.tpl"}}

<div class="container">
    <h4>二维码解析</h4>
    <row>
        <div class="table-responsive">
            <table class="table table-bordered">
                <caption>解析截图或网址</caption>
                <tr>
                    <td>
                        <textarea autocomplete="off" data-provide="typeahead" name="pic" id="pic" placeholder="截图或网址" class="form-control"></textarea>
                    </td>
                    <td>
                        <div class="btn-group pull-right">
                            <button id="decode" type="button" class="btn btn-primary">解析二维码</button>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <img id="theImg" src="">
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <div id="code"></div>
                    </td>
                </tr>
                <!--<tr>
                    <td>
                        <textarea autocomplete="off" data-provide="typeahead" name="url" id="url" placeholder="图片url" class="form-control">https://static.zhihu.com/static/revved/img/index/qr-code.d6565408.png</textarea>
                    </td>
                    <td>
                        <div class="btn-group pull-right">
                            <button id="url2Cavas" type="button" class="btn btn-primary">图片转Canvas</button>
                        </div>
                    </td>
                </tr>-->

            </table>
        </div>
    </row>
</div>
{{template "common/script.tpl"}}
<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>

<script type="text/javascript" src="/static/jsqrcode/grid.js"></script>
<script type="text/javascript" src="/static/jsqrcode/version.js"></script>
<script type="text/javascript" src="/static/jsqrcode/detector.js"></script>
<script type="text/javascript" src="/static/jsqrcode/formatinf.js"></script>
<script type="text/javascript" src="/static/jsqrcode/errorlevel.js"></script>
<script type="text/javascript" src="/static/jsqrcode/bitmat.js"></script>
<script type="text/javascript" src="/static/jsqrcode/datablock.js"></script>
<script type="text/javascript" src="/static/jsqrcode/bmparser.js"></script>
<script type="text/javascript" src="/static/jsqrcode/datamask.js"></script>
<script type="text/javascript" src="/static/jsqrcode/rsdecoder.js"></script>
<script type="text/javascript" src="/static/jsqrcode/gf256poly.js"></script>
<script type="text/javascript" src="/static/jsqrcode/gf256.js"></script>
<script type="text/javascript" src="/static/jsqrcode/decoder.js"></script>
<script type="text/javascript" src="/static/jsqrcode/qrcode.js"></script>
<script type="text/javascript" src="/static/jsqrcode/findpat.js"></script>
<script type="text/javascript" src="/static/jsqrcode/alignpat.js"></script>
<script type="text/javascript" src="/static/jsqrcode/databr.js"></script>
<script type="text/javascript" src="/static/paste/paste.js"></script>
<script>

    $("#decode").bind("click",function () {
        var picData=$("#pic").val().trim();
        if(picData){
            qrcode.decode(picData)
            qrcode.callback=function (data) {
                $("#code").html(data)
            }
        }
    });

    $("#url2Cavas").bind("click",function () {
        var v=$("#url").val()
        if(v){
            draw("url")
        }
    });

    /*function draw(id) {
        // Get the canvas element and set the dimensions.
        var canvas = document.getElementById('canvas');
        canvas.height = window.innerHeight;
        canvas.width = window.innerWidth;

        // Get a 2D context.
        var ctx = canvas.getContext('2d');

        // create new image object to use as pattern
        var img = new Image();
//        img.crossOrigin = "anonymous";
        img.src = document.getElementById(id).value;
        img.onload = function(){
            // Create pattern and don't repeat!
            var ptrn = ctx.createPattern(img,'no-repeat');
             ctx.fillStyle = ptrn;
             ctx.fillRect(0,0,canvas.width,canvas.height);
        };
    }*/

    //canvas to dataUrl
    /*$("#dataUrl").bind("click",function () {
        var canvas = document.getElementById('canvas');
        Obj.dateUrl=canvas.toDataURL("image/png")
        $("#pic").val(Obj.dateUrl)
    });*/

    $(function () {
        function setImg(data) {
            var v=jQuery.trim(data);
            if(v && (v.indexOf("data")>-1 || v.indexOf("http")>-1)){
                $("#theImg").attr("src",v)
            }
        }

        //粘贴图片
        $('#pic').pastableTextarea().on('pasteImage', function(ev, data){
//            var blobUrl = URL.createObjectURL(data.blob);
//            console.log(blobUrl)
            var s = data.dataURL
            $("#pic").val(s)
            setImg(s)
//            console.log(data.width, data.height,data.dataURL);
        }).on('pasteImageError', function(ev, data){
            console.log(data.message);
            if(data.url){
                console.log('But we got its url anyway:' + data.url)
            }
        }).on('pasteText', function(ev, data){
//            console.log(data.text)
            setImg(data.text)
        });
    })
</script>
</body>
</html>
