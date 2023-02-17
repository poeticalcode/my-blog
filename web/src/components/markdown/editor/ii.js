const initPasteDragImg = (Editor, uploadImg) => {
  let doc = document.getElementById(Editor.id)
  doc.addEventListener('paste', function (event) {
      let items = (event.clipboardData || window.clipboardData).items;
      let file = null;
      if (items && items.length) {
          // 搜索剪切板items
          for (let i = 0; i < items.length; i++) {
              if (items[i].type.indexOf('image') !== -1) {
                  file = items[i].getAsFile();
                  break;
              }
          }
      } else {
          console.log("当前浏览器不支持");
          return;
      }
      if (!file) {
          console.log("粘贴内容非图片");
          return;
      }
      uploadImg(file, Editor);
  });
  let dashboard = document.getElementById(Editor.id)
  dashboard.addEventListener("dragover", function (e) {
      e.preventDefault()
      e.stopPropagation()
  })
  dashboard.addEventListener("dragenter", function (e) {
      e.preventDefault()
      e.stopPropagation()
  })
  dashboard.addEventListener("drop", function (e) {
      e.preventDefault()
      e.stopPropagation()
      let files = this.files || e.dataTransfer.files;
      uploadImg(files[0], Editor);
  })
}


function uploadImg(file, Editor) {
  let formData = new FormData();
  let fileName = new Date().getTime() + "." + file.name.split(".").pop();
  formData.append('editormd-image-file', file, fileName);
  $.ajax({
      url: Editor.settings.imageUploadURL,
      type: 'post',
      data: formData,
      processData: false,
      contentType: false,
      dataType: 'json',
      success: function (msg) {
          let success = msg ['success'];
          if (success === 1) {
              let url = msg["url"];
              if (/\.(png|jpg|jpeg|gif|bmp|ico)$/.test(url)) {
                  Editor.insertValue("![图片alt](" + msg["url"] + " '图片title')");
              } else {
                  Editor.insertValue("[下载附件](" + msg["url"] + ")");
              }
          } else {
              console.log(msg);
              alert("上传失败");
          }
      }
  });
}