
<!-- /.row -->
<!-- Flot Charts -->
<div class="row">
    <div class="col-lg-12">
        <h2 class="page-header">用户资料</h2>
    </div>
</div>
<!-- /.row -->
<div class="row">
    <div class="col-lg-12">
        <div class="form-group">
            <p class="form-control-static">姓名：{{.CurrentUser.Name}}</p>
            <p class="form-control-static">邮箱：{{.CurrentUser.Email}}</p>
            <p class="form-control-static">电话：{{.CurrentUser.Phone}}</p>
            {{if .CurrentUser.Major}}
            <p class="form-control-static">研究方向：{{.CurrentUser.Major.MajorName}}</p>
            {{end}}
        </div>
    </div>
</div>
<!-- /.row -->

<!-- Flot Charts -->
<div class="row">
    <div class="col-lg-12">
        <h2 class="page-header"></h2>

    </div>
</div>
<!-- /.row -->
<div class="row">
    <div class="col-md-2">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#edit_modal">
            编辑
        </button>
    </div>
</div>


<!-- 编辑Modal -->
<div class="modal fade" id="edit_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <form role="form" action="" method="post">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span
                                aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">编辑个人信息</h4>
                </div>
                <div class="modal-body">
                    <div class="form-group">
                        <label>邮箱：</label>
                        <input type="email" class="form-control" name="email" placeholder="请输入邮箱" value="{{.CurrentUser.Email}}">
                        <br>
                        <label>电话：</label>
                        <input type="text" class="form-control" name="phone" placeholder="请输入电话" value="{{.CurrentUser.Phone}}">
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                    <button type="submit" class="btn btn-primary">保存</button>
                </div>
            </div>
        </div>
    </form>
</div>
