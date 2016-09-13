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

$("#subject_success").click(function () {
    def_post({status:11});
});

$("#subject_fail").click(function () {
    def_post({status:12});
});

$("#task_success").click(function () {
    def_post({status:21});
});

$("#task_fail").click(function () {
    def_post({status:22});
});

$("#primary_success").click(function () {
    def_post({status:31});
});

$("#primary_fail").click(function () {
    def_post({status:32});
});

$("#interim_success").click(function () {
    def_post({status:41});
});

$("#interim_fail").click(function () {
    def_post({status:42});
});

$("#paper_success").click(function () {
    def_post({status:51});
});

$("#paper_fail").click(function () {
    def_post({status:52});
});

$("#translate_success").click(function () {
    def_post({status:61});
});

$("#translate_fail").click(function () {
    def_post({status:62});
});

</script>


