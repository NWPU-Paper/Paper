<div class="page-header">
    <h1><a  href="{{urlfor "TeacherController.Subject"}}">< 课题列表</a> - 添加新课题</h1>
</div>
<div class="row">
    <div  style="margin:20px">

        <form role="form" action="{{urlfor "TeacherController.ExecAddSubject"}}" method="post">
            <div class="input-group input-group-lg" style="padding: 10px 0px 10px;">
                <span class="input-group-addon">课题名</span>
                <input type="text" name="title"class="form-control" required>
            </div>
            <div class="form-group">
                <label for="name">课题简介</label>
                <textarea class="form-control" rows="5" name="introduce" required></textarea>
            </div>
            <div style="width:50%;margin-left: auto;margin-right: auto">
                <button type="submit"  class="btn btn-default" style="width:100%;">保存课题</button>
            </div>
        </form>
    </div>

</div>