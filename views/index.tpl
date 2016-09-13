<div class="row">
    <!-- blue-->
    <div class="col-md-4">
        <div class="panel panel-primary">
            <div class="panel-heading">
                <div class="row">
                    <div class="col-xs-3">
                        <div class="middle">今日日期</div>
                    </div>
                    <div class="col-xs-9 text-right">
                        <div class="huge">{{dateformat .nowDate "2006-01-02"}}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- green-->
    <div class="col-md-4">
        <div class="panel panel-green">
            <div class="panel-heading">
                <div class="row">
                    <div class="col-xs-3">
                        <div class="middle">距离结束</div>
                        <!-- <i class="fa fa-tasks fa-5x"></i>  -->
                    </div>
                    <div class="col-xs-9 text-right">
                        <div class="huge"> {{.endDay}}天</div>
                    </div>
                </div>
            </div>

        </div>
    </div>


</div>
