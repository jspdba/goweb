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
    <h4>二维码解析</h4>

    <row>
        <div class="table-responsive">
            <table class="table table-bordered">
                <tr>
                    <td>
                        <textarea autocomplete="off" data-provide="typeahead" name="pic" id="pic" placeholder="图片data" class="form-control"></textarea>
                    </td>
                    <td>
                        <div class="btn-group pull-right">
                            <button id="decode" type="button" class="btn btn-primary">解析</button>
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
<script>
    //解析base64图片数据
    $("#decode").bind("click",function () {
        var picData=$("#pic").val();
        if(picData){
            qrcode.decode(picData)
            qrcode.callback=function (data) {
                console.log(data)
                $("#code").html(data)
            }
        }
    });
</script>
</body>
</html>
