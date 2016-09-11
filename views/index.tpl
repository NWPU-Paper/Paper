<div id="page-wrapper">
    <div class="container-fluid">
        <!-- Page Heading -->
        <div class="row">
            <div class="col-lg-12">
                <h1 class="page-header">
                    高校论文管理平台
                </h1>
            </div>
        </div>
        <!-- /.row -->
        <!-- /.row -->
        <div class="row">
            <!-- blue-->
            <div class="col-lg-3 col-md-6">
                <div class="panel panel-primary">
                    <div class="panel-heading">
                        <div class="row">
                            <div class="col-xs-3">
                                <div class="middle">今日日期</div>
                            </div>
                            <div class="col-xs-9 text-right">
                                <div class="huge"> 456</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <!-- green-->
            <div class="col-lg-3 col-md-6">
                <div class="panel panel-green">
                    <div class="panel-heading">
                        <div class="row">
                            <div class="col-xs-3">
                                <div class="middle">结束时间</div>
                                <!-- <i class="fa fa-tasks fa-5x"></i>  -->
                            </div>
                            <div class="col-xs-9 text-right">
                                <div class="huge"> 123</div>

                            </div>
                        </div>
                    </div>

                </div>
            </div>

            {{if eq .CurrentUser.Type 2}}

            <!-- yellow-->
            <div class="col-lg-3 col-md-6">
                <div class="panel panel-yellow">
                    <div class="panel-heading">
                        <div class="row">
                            <div class="col-xs-3">
                                <div class="middle">指导人数</div>
                                <!-- <i class="fa fa-shopping-cart fa-5x"></i> -->
                            </div>
                            <div class="col-xs-9 text-right">
                                <div class="huge">4人</div>

                            </div>
                        </div>
                    </div>

                </div>
            </div>


            <!-- red-->
            <div class="col-lg-3 col-md-6">
                <div class="panel panel-red">
                    <div class="panel-heading">
                        <div class="row">
                            <div class="col-xs-3">
                                <div class="middle">已选人数</div>
                                <!--<i class="fa fa-support fa-5x"></i> -->
                            </div>
                            <div class="col-xs-9 text-right">
                                <div class="huge">4人</div>

                            </div>
                        </div>
                    </div>

                </div>
            </div>

            {{end}}

        </div>
    </div>
</div>