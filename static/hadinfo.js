let agreementCheckbox = $("#agreement");
let addFormSubmitButton = $("#addFormSubmitButton");
window.onload = function () {
    agreementCheckbox.change(function () {
        $("#addFormSubmitButton")[0].className = "ui button " + (agreementCheckbox[0].checked ? "" : "disabled");
    });
    addFormSubmitButton.on("click", function () {
        $("#addFormSubmitButton")[0].className = "ui button disabled"
        $("#content").val(window.btoa($("#content_input").val())); // base64 encoding
        let fail=function (err) {
            console.log("error"+err)
        }
        $.ajax({
            url: "api/add",
            method: "post",
            data: $("#addForm").serialize(),
            success:function(result){
                console.log(result)
                if(result["message"]==="done") {
                    window.location.href="/v/"+result["pasteId"];
                }
                else {
                    fail(result)
                }
            },
            error: fail
        })
    })
}