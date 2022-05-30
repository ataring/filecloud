
$(function () {

    


    // 上传
    $("#cb-upload-input").on("change", function (event) {
        var inst = new mdui.Dialog("#cb-loading-view", {modal: true,});
        inst.open();

        var files = $('#cb-upload-input').prop('files');
        var data = new FormData();
        data.append("basePath", location.pathname);
        for (var i = 0; i < files.length; i++) {
            data.append("files", files[i]);
        }

        $.ajax({
            url: "/cloudupload",
            type: "POST",
            data: data,
            cache: false,
            processData: false,
            contentType: false,
            xhr: function () {
                myXhr = $.ajaxSettings.xhr();
                if (myXhr.upload) {
                    myXhr.upload.addEventListener('progress', progressHandlingFunction, false);
                }
                return myXhr;
            },
            success: function (msg) {
                if (msg = 200) {
                    toast("上传成功！");
                    setTimeout(function () {
                        location.reload();
                    }, 1000);
                } else {
                    toast("上传失败！");
                    
                }
            },
            error: function (request, textStatus, errorThrown) {
                toast(textStatus + ":" + errorThrown.toString())
            },
            complete: function () {
                inst.close();
            }
        });
    });



    //转移文件
    $(".cb-chage-btn").on("click", function () {
        let inst = new mdui.Dialog("#cb-change-view");
        inst.open();
    });
    // 删除
    let filePath = $(this).attr("attr-filePath");
    $(".cb-delete-btn").on("click", function () {
        filePath = $(this).attr("attr-filePath");
        let inst = new mdui.Dialog("#cb-remove-tips");
        inst.open();
    });

    // 确认删除
    $(".cb-delete-confirm").on("click", function () {
        let query = "path=" + encodeURIComponent(filePath);
        $.get("/remove?" + query, function (data) {
            if (data.code === -1000) {
                toast(data.message);
            } else {
                location.reload();
            }
        })
    })
});



function toast(message) {
    mdui.snackbar(message, {
        position: "right-top",
    });
}

function progressHandlingFunction(e) {
    if (e.lengthComputable) {
        var percent = Math.floor(e.loaded / e.total * 100);
        $("#cb-loading-view .mdui-progress-determinate").css("width", percent + "%");
    }
}