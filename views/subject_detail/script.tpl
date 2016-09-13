<script>//设置默认action
var operate = 1;

function def_post(v) {
    v.operate = operate;
    $.post("./{{.Subject.Id}}",v,function ($data) {
//            if($data.status==200)
//                alert("操作成功");
//            else
//                alert("操作失败<br>" + data.message);
        window.location.reload();
    })
}

$("#task_success").click(function () {
    def_post({status:21});
});

$("#task_fail").click(function () {
    def_post({status:22});
});

</script>


