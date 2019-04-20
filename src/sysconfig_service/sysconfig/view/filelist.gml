<html>

<script src="resource/jquery/jquery-1.12.1.js"></script>
<script src="resource/bootstrap3.3.6/js/bootstrap.min.js"></script>

<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/bootstrap-theme.min.css">
<link rel="stylesheet" type="text/css" href="resource/bootstrap3.3.6/css/docs.min.css">

<style>
.div1_container_height{
height:100%;font-size:14px
}

.div2_col_odd{
border:3px solid #DBDBDB;height:100% ; background:#FDFFFF;padding:0
}

.div2_col_eve{
border:3px solid #DBDBDB;height:100% ; padding:0
}

.divx_middle {
   float: none;display: inline-block;vertical-align: middle;
}

.div3_title {
   width:100%;height:50px;align-items: center;display: flex;justify-content: center;
}

.div3_selected {
   width:100%;height:50px;align-items: center;display: flex;border-top:3px solid #DBDBDB;border-bottom:3px solid #DBDBDB
}

.div3_selected_wh {
   width:100%;height:100%
}

.div3_selected_name {
   background:#f5f5f5;height:100%;align-items: center;display: flex;
}

.div3_selected_value {
   background:#FFF;height:100%;
}

.div3_selected_value_wrap {
   height:100%;weight:100%;display: table;
}

.div3_selected_value_text {
   display: table-cell; vertical-align: middle;word-break:break-all;
}

.div3_datas {
   width: 100%; position: absolute; top: 100px ; left: 0 ; bottom: 0;color:gray; overflow-y:scroll
}

.div3_datas_nest {
  padding:10px
}

.list-group-item:hover{
   background:gray;color:white;cursor:hand
}

</style>

<script>

var oldProjectId = null;
function clickProjectId(obj){

   if (oldProjectId!=null) {
      oldProjectId.removeClass("glyphicon-check").addClass("glyphicon-unchecked").removeAttr("style");
   }

   oldProjectId=$(obj).children(".glyphicon-unchecked")
   oldProjectId.removeClass("glyphicon-unchecked").addClass("glyphicon-check").css("color","black");

   $(".obj_projectId_value").text(obj.innerText)

   $.post("filelistjson?level=2&path="+$(obj).attr("path"), function(data) {
     $(".aliasKeyDivList").html(data);
   });
}

var oldAliasKey = null;
function clickAliasKey(obj){

    if (oldAliasKey!=null) {
       oldAliasKey.removeClass("glyphicon-check").addClass("glyphicon-unchecked").removeAttr("style");
    }

    oldAliasKey=$(obj).children(".glyphicon-unchecked")
    oldAliasKey.removeClass("glyphicon-unchecked").addClass("glyphicon-check").css("color","black");

   $(".obj_aliasKey_value").text(obj.innerText)

   $.post("filelistjson?level=3&path="+$(obj).attr("path"), function(data) {
     $(".appNameDivList").html(data);
   });
}

var oldAppName = null;
function clickAppName(obj){
    if (oldAppName!=null) {
       oldAppName.removeClass("glyphicon-check").addClass("glyphicon-unchecked").removeAttr("style");
    }

    oldAppName=$(obj).children(".glyphicon-unchecked")
    oldAppName.removeClass("glyphicon-unchecked").addClass("glyphicon-check").css("color","black");

   $(".obj_appName_value").text(obj.innerText)

   $.post("filelistjson?level=4&path="+$(obj).attr("path"), function(data) {
     $(".filesDivList").html(data);
   });
}


function initDatas(){

  $.post("filelistjson?level=1", function(data) {
    $(".projectDivList").html(data);
  });

}

$(document).ready(function(){

  initDatas();
});

</script>

<body>
<div class="container-fluid div1_container_height" >
<div class="row">
  <div class="col-xs-3 div2_col_odd" >
    <div class="div3_title bg-primary" style="background:#6b6b6d">
      项目名称(projectId)
    </div>
    <div class="div3_selected bg-info" >
      <div class='div3_selected_wh'>
        <div class="col-xs-3 div3_selected_name" style='padding:5px'>
        筛选项：
        </div>
        <div class="col-xs-9 div3_selected_value" >
        <div class="div3_selected_value_wrap">
            <div class="div3_selected_value_text obj_projectId_value">
                <div style="color:BBBBBB">请选择...</div>
            </div>
        </div>
        </div>
      </div>
    </div>
    <div class="div3_datas">
      <div class="div3_datas_nest">
          <ul class="list-group projectDivList">
              <!--  1
              <li class="list-group-item" onclick='clickProjectId(this)' path='' filename=''>
                  <span class="glyphicon glyphicon-unchecked" aria-hidden="true"></span>
                  <span style="padding-left:5px">project002</span>
              </li>
            -->
          </ul>
      </div>
    </div>
  </div>
  <div class="col-xs-3 div2_col_eve" >
    <div class="div3_title bg-primary" style="background:#6b6b6d">
      配置项(aliasKey)
    </div>
    <div class="div3_selected bg-info" >
      <div class='div3_selected_wh'>
        <div class="col-xs-3 div3_selected_name" style='padding:5px'>
        筛选项：
        </div>
        <div class="col-xs-9 div3_selected_value" >
        <div class="div3_selected_value_wrap">
            <div class="div3_selected_value_text obj_aliasKey_value">
                <div style="color:BBBBBB">请选择...</div>
            </div>
        </div>
        </div>
      </div>
    </div>
    <div class="div3_datas">
    <div class="div3_datas_nest">
        <ul class="list-group  aliasKeyDivList">
          <!--  2
            <li class="list-group-item" onclick='clickAliasKey(this)' path='' filename=''>
               <span class="glyphicon glyphicon-unchecked" aria-hidden="true"></span>
               <span style="padding-left:5px">advs_remote</span>
            </li>
          -->
        </ul>
    </div>
    </div>
  </div>
  <div class="col-xs-3 div2_col_odd" >
    <div class="div3_title bg-primary" style="background:#6b6b6d">
      应用服务(appName)
    </div>
    <div class="div3_selected bg-info" >
      <div class='div3_selected_wh'>
        <div class="col-xs-3 div3_selected_name" style='padding:5px'>
        筛选项：
        </div>
        <div class="col-xs-9 div3_selected_value" >
        <div class="div3_selected_value_wrap">
            <div class="div3_selected_value_text obj_appName_value">
                <div style="color:BBBBBB">请选择...</div>
            </div>
        </div>
        </div>
      </div>
    </div>
    <div class="div3_datas">
      <div class="div3_datas_nest">
          <ul class="list-group appNameDivList">
            <!-- 3
              <li class="list-group-item" onclick='clickAppName(this)'>
                 <span class="glyphicon glyphicon-unchecked" aria-hidden="true"></span>
                 <span style="padding-left:5px">app_d2</span>
              </li>
            -->

          </ul>
      </div>
    </div>
  </div>
  <div class="col-xs-3 div2_col_eve" >
    <div class="div3_title bg-primary" style="background:#6b6b6d">
      打包的配置文件列表
    </div>
    <div class="div3_selected " style='background:#eceff1'>

    </div>
    <div class="div3_datas">
      <div class="div3_datas_nest">
      <ul class="list-group filesDivList">
         <!-- 4
          <li class="list-group-item" style='color:#555555'>
             <span class="glyphicon glyphicon-file" aria-hidden="true"></span>
             <span style="padding-left:5px">web.toml</span>
          </li>
          <li class="list-group-item" style='color:#999999'>
             <span class="glyphicon glyphicon-folder-open" aria-hidden="true"></span>
             <span style="padding-left:5px">simple</span>
          </li>
          <li class="list-group-item" style='color:#555555;margin-left:20px'>
             <span class="glyphicon glyphicon-file" aria-hidden="true"></span>
             <span style="padding-left:5px">simple-dist.toml</span>
          </li>
          <li class="list-group-item" style='color:#555555;margin-left:20px'>
             <span class="glyphicon glyphicon-file" aria-hidden="true"></span>
             <span style="padding-left:5px">simple-dubbo.toml</span>
          </li>
          <li class="list-group-item" style='color:#555555'>
             <span class="glyphicon glyphicon-file" aria-hidden="true"></span>
             <span style="padding-left:5px">db.toml</span>
          </li>
        -->
      </ul>
      </div>
    </div>
  </div>
</div>

</div>
</body>
</html>
