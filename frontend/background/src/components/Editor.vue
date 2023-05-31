<template>
  <div>
    <quill-editor  @change="onEditorChange"
        id="content" ref="quill"  v-model="value" :options="editorOption">
    </quill-editor>
  </div>
</template>

<script>
import { quillEditor, Quill } from 'vue-quill-editor'
import { container, ImageExtend, QuillWatch } from 'quill-image-extend-module'
Quill.register('modules/ImageExtend', ImageExtend)

export default {
  name: "Editor",
  props: {
    value: {
      type: String,
    }
  },
  data(){
    return {
      editorOption: {
        modules: {
          ImageExtend: {
            loading: true,
            name: 'img',//图片参数名
            size: 5, // 可选参数 图片大小，单位为M，1M = 1024kb
            action: 'https://api.superbed.cn/upload',//上传的服务器地址，如果action为空，则采用base64插入图片
            // response 为一个函数，用来获取服务器返回的具体图片地址
            response: res => {
              console.log(res)
              return res.url
            },
            success: () => {
              console.log('ImageExtend中的success：上传成功')
            }, // 可选参数  上传成功触发的事件
            change: (xhr, formData) => {
              // xhr.setRequestHeader('myHeader','myValue')
              formData.append('token', 'f387bdf137f04651bfa69c18910bc2d6')
            } // 可选参数 每次选择图片触发，也可用来设置头部，但比headers多了一个参数，可设置formData
          },
          // 如果不上传图片到服务器，此处不必配置
          toolbar: {
            // container为工具栏，此次引入了全部工具栏，也可自行配置
            container: [
              ["bold", "italic", "underline", "strike"],
              ["blockquote", "code-block"],
              [{ header: 1 }, { header: 2 }],
              [{ list: "ordered" }, { list: "bullet" }],
              [{ script: "sub" }, { script: "super" }],
              [{ indent: "-1" }, { indent: "+1" }],
              [{ direction: "rtl" }],
              [{ size: ["small", false, "large", "huge"] }],
              [{ header: [1, 2, 3, 4, 5, 6, false] }],
              [{ color: [] }, { background: [] }],
              [{ font: [] }],
              [{ align: [] }],
              ["image"]
            ],
            // 上传成功，回显图片（会进入如上面ImageExtend的各过程，返回<img src="http://xx.xx.xx.xx:xxxx/file/xxx.jpg">）
            handlers: {
              image: function() {
                // 劫持原来的图片点击按钮事件
                QuillWatch.emit(this.quill.id)
              }
            }
          }
        }
      },

    }
  },
}
</script>

<style scoped>

</style>