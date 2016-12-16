<style>
    .active,a:active {
        color: #0000FF;
        background-color: #2aabd2;
    } /* 被选择的链接 */
</style>
<ol>
    {{range $Path:=.}}
    <li class="item"><a href="{{$Path}}" target="showframe">{{$Path}}</a></li>{{end}}
</ol>
<script src="../js/jquery-1.8.3.min.js"></script>
<script>
    $(".item").bind("click",function () {
        $(".active").removeClass("active");
        $(this).addClass("active");
    })
</script>