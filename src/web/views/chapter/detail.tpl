<!DOCTYPE html>
<html lang="zh-cn">
<style>
    body {
        padding-top: 70px;
        padding-bottom: 50px;
    }
</style>
{{template "common/header_flat.tpl"}}
<title>{{.entity.Title}}</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <h4>{{.entity.Title}}</h4>
    <p>
        {{str2html .entity.Content}}
    </p>
    <div class="navbar-fixed-bottom">
        <div class="row">
            <div class="col-md-12">
                <form action="{{urlfor "ChapterController.FindByTitle" ":id" .entity.Book.Id}}" class="form-inline" role="form" method="post">
                    <div class="form-group">
                        <div class="input-group">
                            <input name="Title" type="text" class="form-control" placeholder="搜索">
                            <span class="input-group-btn">
                                <button class="btn btn-default" type="submit">Go!</button>
                                <a id="pre" type="button" class="btn btn-primary btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .pre.Id}}">上一页</a>
                                <a id="next" type="button" class="btn btn-info btn-sm" href="{{urlfor "ChapterController.Detail" ":id" .next.Id}}">下一页</a>
                                <button id="er" type="button" class="btn btn-info btn-sm">二维码</button>
                                <a id="update" type="button" class="btn btn-info btn-sm" href="{{urlfor "ChapterController.Update" ":id" .entity.Id}}">更新</a>
                                <button id="addToFav" type="button" class="btn btn-info btn-sm" link="{{urlfor "LinkController.PostLink"}}">收藏</button>
                            </span>
                        </div>
                    </div>
                </form>
            </div><!-- /.col-md-12 -->
        </div><!-- /.row -->
    </div>
</div><!-- /.container -->

<!-- Modal -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title" id="myModalLabel">添加 Tag</h4>
            </div>
            <div class="modal-body">
                <input id="tag" type="text" value="tag1">
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button id="useTag" type="button" class="btn btn-primary">使用</button>
            </div>
        </div>
    </div>
</div>

<!-- 二维码 -->
<div class="modal fade" id="qrCodeModal" tabindex="-1" role="dialog" aria-labelledby="qrCodeModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title" id="qrCodeModalLabel">二维码</h4>
            </div>
            <div class="modal-body">
                <div class="row">
                    <div class="col-sm-12">
                        <div id="code">
                        </div>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">确定</button>
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
    var tag = Cookies.get("__TAG")
    if(!tag){
        $('#myModal').modal('show')
    }
    //使用tag
    $("#useTag").bind("click",function (data) {
        var value = $("#tag").val();
        if(value){
            if($(this).hasClass("sure")){
                Cookies.set('__TAG', value, { expires: 365 });
                $('#myModal').modal('hide')
            }else{
                var url="/log/tag/"+value;
                $.getJSON(url,function (data) {
                    if(data.code==0){
                        if(data.result==0){
                            Cookies.set('__TAG', value, { expires: 365 });
                            $('#myModal').modal('hide')
                        }else{
                            toastr.warning(value+" 已存在！")
                            $("#useTag").html("确定使用")
                            if(!$("#useTag").hasClass("sure")){
                                $("#useTag").addClass("sure")
                            }
                        }
                    }
                })
            }
        }
    });

    var url="{{urlfor "ChapterLogController.Add" ":tag" "tag1" ":bookId" .entity.Book.Id ":index" .entity.Index}}";
    if(url){
      $.getJSON(url,function (data) {
          if(data.code==0){
              console.log("添加日志成功")
          }
      })
    }

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
    //生产二维码
    $("#er").bind("click",function () {
        var str = toUtf8(window.location.href);
        $("#code").html("");
        $("#code").qrcode({
            correctLevel:0,
            width: 200, //宽度
            height:200, //高度
            text: str //任意内容
        });
        $('#qrCodeModal').modal('show')
    })
    //添加收藏
    $("#addToFav").bind("click",function () {
        var url=$(this).attr("link");
        $.post(url,{
            "Url":window.location.href,
            "Tag.name":"",
            "Title":$("title").html(),
            "Tags.Name":"",

        },function (data) {
            if(data.code==0){
                toastr.info("收藏成功")
            }else{
                toastr.warning("收藏失败")
            }
        },"json")
        })
</script>
</body>
</html>
