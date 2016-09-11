<div class="page-header">
    <h1><a  href="{{urlfor "TeacherController.Subject"}}">< 课题列表</a> - 添加新课题</h1>
</div>
<div class="row">
    <div class="col-lg-12">
        <div class="input-group input-group-lg" style="padding: 10px 0px 10px;">
            <span class="input-group-addon">课题名</span>
            <input type="text" class="form-control">
        </div><br>

        <div class="input-group input-group-lg">
            <span class="input-group-addon">教师姓名</span>
            <input type="text" class="form-control">
        </div><br>
        <form role="form">
            <div class="form-group">
                <label for="name">课题简介</label>
                <textarea class="form-control" rows="3"></textarea>
            </div>
        </form>
    </div>
    <div style="width:50%;margin-left: auto;margin-right: auto">
        <button type="button" class="btn btn-default" style="width:100%;">保存课题</button>
    </div>
</div>