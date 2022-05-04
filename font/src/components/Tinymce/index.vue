<template>
  <div :class="{fullscreen:fullscreen}" class="tinymce-container" :style="{width:containerWidth}">
    <textarea :id="tinymceId" class="tinymce-textarea" />
  </div>
</template>

<script lang="ts">
/**
 * docs:
 * https://panjiachen.github.io/vue-element-admin-site/feature/component/rich-editor.html#tinymce
 */
import plugins from './plugins'
import toolbar from './toolbar'
import load from './dynamicLoadScript'
// why use this cdn, detail see https://github.com/PanJiaChen/tinymce-all-in-one
const tinymceCDN = '/tinymce-need-file/cdn.js'//'https://cdn.jsdelivr.net/npm/tinymce-all-in-one@4.9.3/tinymce.min.js';
// const tinymceCDN = './tinymce.min.ts';
export default {
  name: 'Tinymce',
  props: {
    id: {
      type: String,
      default: function() {
        return 'vue-tinymce-' + +new Date() + ((Math.random() * 1000).toFixed(0) + '')
      }
    },
    maxLength:{
      type:Number,
      default:12000,
    },
    modelValue: {
      type: String,
      default: ''
    },
    toolbar: {
      type: Array,
      required: false,
      default() {
        return []
      }
    },
    menubar: {
      type: String,
      default: 'file edit insert view format table'
    },
    height: {
      type: [Number, String],
      required: false,
      default: 150
    },
    width: {
      type: [Number, String],
      required: false,
      default: 'auto'
    },
    wordcount:{
      type:Number,
      default:0,
    },
    contentBgColor:{
      type:String,
      default:"#000000",
    },
    textColorMap:{
      type:Array,
      default:[],
    }

  },
  emits:["update:modelValue", 'update:wordcount', 'input'],
  data() {
    return {
      hasChange: false,
      hasInit: false,
      tinymceId: this.id,
      fullscreen: false,
      count:0,
    }
  },
  computed: {
    containerWidth() {
      const width = this.width
      if (/^[\d]+(\.[\d]+)?$/.test(width)) { // matches `100`, `'100'`
        return `${width}px`
      }
      return width
    }
  },
  watch: {
    modelValue(val) {
      var _this =this
      if(this.hasInit){
         _this.compWordCount(val)
      }
      if (!this.hasChange && this.hasInit) {
        this.$nextTick(() =>{
           this.setContent(val || '');
        })
      }
    }
  },
  mounted() {
    this.init(this.modelValue)
  },
  activated() {
    if (window.tinymce) {
      this.initTinymce()
    }
  },
  deactivated() {
    this.destroyTinymce()
  },
  destroyed() {
    this.destroyTinymce()
  },
  methods: {
    init(val) {
      // dynamic load tinymce from cdn
      load(tinymceCDN, (err) => {
        if (err) {
          this.$message.error(err.message)
          return
        }
        this.initTinymce(val)
      })
    },
    // 计算输入的字符数
    compWordCount(value){
        var html = value || '';
        var re1 = new RegExp("<.+?>","g");
        var txt = html.replace(re1,'');
        txt = txt.replace(/\n/g,'');
        txt = txt.replace(/&nbsp;/g,' ');
        this.$emit('update:wordcount',txt.length)
        let domId = `wordcount_${this.id}`
        if( window.tinymce.get(this.id) && window.tinymce.get(this.id).editorContainer){
          this.count=txt.length;
          let spanDom = window.tinymce.get(this.id).editorContainer.querySelector(`#${domId}`);
          // 设置颜色
          if(spanDom){
              let color = "black";
            if (this.count > this.maxLength){
              color  = "red"
            }
            spanDom.style.color=color;
            spanDom.innerHTML = this.count+"字"
          }
          
        }
    },

    // min_height: 150,
    // max_height: 150,
    initTinymce(val,textColorMap) {
      const _this = this
      if(!(window.colorData && window.colorData.length >4)){
        textColorMap = [
          "#000000", "Black","#993300", "Burnt orange","#333300", "Dark olive", "#003300", "Dark green", "#003366", "Dark azure",  "#000080", "Navy Blue",  "#333399", "Indigo", 
          "#333333", "Very dark gray", "#800000", "Maroon", "#FF6600", "Orange", "#808000", "Olive", "#008000", "Green", "#008080", "Teal",  "#0000FF", "Blue", "#666699", "Grayish blue", 
          "#808080", "Gray", "#FF0000", "Red", "#FF9900", "Amber", "#99CC00", "Yellow green", "#339966", "Sea green", "#33CCCC", "Turquoise", "#3366FF", "Royal blue", 
          "#800080", "Purple", "#999999", "Medium gray", "#FF00FF", "Magenta", "#FFCC00", "Gold", "#FFFF00", "Yellow", "#00FF00", "Lime", "#00FFFF", "Aqua", "#00CCFF", 
          "Sky blue", "#993366", "Red violet", "#FFFFFF", "White", "#FF99CC", "Pink", "#FFCC99", "Peach", "#FFFF99", "Light yellow", "#CCFFCC", "Pale green", "#CCFFFF", 
          "Pale cyan", "#99CCFF", "Light sky blue", "#CC99FF", "Plum"]
      } else {
        textColorMap = window.colorData
      }
     
      window.tinymce.init({
        textcolor_map: textColorMap,
        branding: false,
        selector: `#${this.tinymceId}`,
        language: 'zh_CN',
        height: this.height,
        body_class: "panel-body",
        object_resizing: false,
        toolbar: this.toolbar.length > 0 ? this.toolbar : toolbar,
        menubar: false,
        plugins: plugins,
        end_container_on_empty_block: true,
        powerpaste_word_import: 'clean',
        code_dialog_height: 250,
        code_dialog_width: 1000,
        advlist_bullet_styles: 'square',
        advlist_number_styles: 'default',
        imagetools_cors_hosts: ['www.tinymce.com', 'codepen.io'],
        default_link_target: '_blank',
        link_title: false,
        nonbreaking_force_tab: true, // inserting nonbreaking space &nbsp; need Nonbreaking Space Plugin
        init_instance_callback: editor => {
          if (_this.value) {
            editor.setContent(_this.value)
          }
          // 挂载字数统计
          var childNode = document.createElement('span');
          childNode.setAttribute("id", `wordcount_${this.id}`);   // 设置属性
          childNode.className ="mce-wordcount mce-widget mce-label mce-flow-layout-item";   // 设置属性
          childNode.innerText = "0 字";   // 设置text值
          let parentNode = window.tinymce.activeEditor.editorContainer.querySelector(".mce-resizehandle").parentElement 
          parentNode.appendChild(childNode)
          
          // 模式切换 
          var bgnode = document.createElement('span');
          bgnode.className ="mce-wordcount mce-widget mce-label  custom-switch"; 
          bgnode.innerText = " ■  背景切换"
          let bg_color = this.contentBgColor
          bgnode.addEventListener("click",function(){
            let content_body =window.tinymce.activeEditor.editorContainer.querySelector("iframe").contentWindow.document.querySelector(".mce-content-body")
            if (content_body.classList.contains("bg_back")){
              this.style.color = "#1890FF"
                content_body.classList.remove("bg_back")
                content_body.style.backgroundColor = "inherit"
            }else{
              this.style.color = "black"
              content_body.classList.add("bg_back")
              content_body.style.backgroundColor = bg_color
            } 
          })
          parentNode.appendChild(bgnode)
          // 显示一下
          if(_this.count<=0){
            _this.compWordCount(this.value)
          }
          _this.hasInit = true
          editor.on('NodeChange Change KeyUp SetContent', () => {
            this.hasChange = true
            this.$emit('update:modelValue', editor.getContent())
          })
        },
        setup(editor) {
          editor.on('FullscreenStateChanged', (e) => {
            _this.fullscreen = e.state
          })
        },
        // without images_upload_url set, Upload tab won't show up
        images_upload_url: window.base_url + 'admin/fileupload/uploads',
        // we override default upload handler to simulate successful upload
        images_upload_handler: function (blobInfo, success, failure) {
          let xhr, formData
          xhr = new XMLHttpRequest()
          xhr.withCredentials = false
          //xhr.open('POST', window.base_url+'admin/fileupload/uploads');
          xhr.open(
            'POST',
            'http://csis.cm.com/script/imgupload?domain=van.oa.com&referer=http%3A%2F%2Fvan.oa.com%2Ffaq%2Fupload&internet=true'
          )

          xhr.onload = function () {
            let json
            if (xhr.status != 200) {
              failure('HTTP Error: ' + xhr.status)
              return
            }

            /*
            json = JSON.parse(xhr.responseText);
            if (!json || !json.data || typeof json.data.url != 'string') {
                failure('Invalid JSON: ' + xhr.responseText);
                return;
            }
            success(json.data.url);
            */

            //通过csis接口上传图片，但是返回的结果是完整的html代码，需要通过下面的正则匹配到正确的图片地址
            //匹配图片地址，如：http://file.service.qq.com/user-files/uploads/201704/fb039e7820afc8eacdb9f807824a2333.png
            const picurl = xhr.responseText.match(
              /(http\:\/\/file\.service\.qq\.com.+?)\'/
            )
            if (picurl && picurl[1]) {
              success(picurl[1].replace('http://', 'https://'))
              console.log('=success=', xhr)
              //this.$message("上传成功，请点击“确定”按钮保存");
            } else {
              alert('上传失败,请务必重新上传此图片')
              success('/uploadErr.png')
              //this.$message.error("上传失败");
              console.log('=====blobInfo err ====', blobInfo)
            }
            // if(picurl && picurl[1]){
            // 	console.log("=====upload ok ====",picurl[1]);
            //     success(picurl[1]);
            //     this.$message("上传成功，请点击“确定”按钮保存");
            // }else{
            // 	console.log("=====upload err ====",picurl[1]);
            //     this.$message.error("上传失败");
            // }
          }
          formData = new FormData()
          formData.append(
            'data[Upload][file]',
            blobInfo.blob(),
            blobInfo.filename()
          )
          xhr.send(formData)
        },
        images_upload_credentials: true,
        convert_urls: false
      })

      if (val) {
        this.setContent(val)
        setTimeout(function(){
          _this.compWordCount(val)
        },200)
      }

      // setTimeout(() => {
      //   document.getElementById('mceu_48').remove()
      // },100)
    },
    destroyTinymce() {
      const tinymce = window.tinymce.get(this.tinymceId)
      if (this.fullscreen) {
        tinymce.execCommand('mceFullScreen')
      }

      if (tinymce) {
        tinymce.destroy()
      }
    },
    setContent(value) {
      window.tinymce.get(this.tinymceId).setContent(value)
      this.compWordCount(value)
    },
    getContent() {
      window.tinymce.get(this.tinymceId).getContent()
    },
    setContentempty() {
      window.tinymce.get(this.tinymceId).setContent()
    },
    imageSuccessCBK(arr) {
      arr.forEach(v => window.tinymce.get(this.tinymceId).insertContent(`<img class="wscnph" src="${v.url}" >`))
    }
  }
}
</script>

<style lang="scss" >

.panel-body{
  
}

.tinymce-container {
  position: relative;
  line-height: normal;
}

.tinymce-container {
  ::v-deep {
    .mce-fullscreen {
      z-index: 1000;
    }
  }
}

.tinymce-textarea {
  visibility: hidden;
  z-index: -1;
}

.editor-custom-btn-container {
  position: absolute;
  right: 4px;
  top: 4px;
  /*z-index: 2005;*/
}

.fullscreen .editor-custom-btn-container {
  z-index: 100;
  position: fixed;
}

.editor-upload-btn {
  display: inline-block;
}
</style>
