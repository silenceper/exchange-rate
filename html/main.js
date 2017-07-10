$(function () {
    var query_path="/exchange";
    $("#exchange-btn").click(function () {
        var from_val = $("#from_val").val();
        if (from_val == "") {
            alert("请先输入要转换的金额");
            return;
        }
        var rate_from = $("#rate_from").val();
        var rate_to = $("#rate_to").val();

        $.ajax({
　　　　　　url: query_path,
　　　　　　type: 'get',
　　　　　　dataType: 'json',
　　　　　　data: {"rate_from":rate_from,"rate_to":rate_to,"from_val":from_val},
　　　　}).done(function(data) {
　　　　　　if(data.retCode!=0){
            alert(data.retMsg);
          }else{
            $("#to_val").val(data.data);
          }
　　　　}).fail(function() {
　　　　　　console.log("request error");
　　　　})
    });
});