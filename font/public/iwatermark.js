/**
 * ----------------------------------------------------------------------------------
 * 生成文字水印
 * 页面水印：
 * 1、生成透明的div平铺
 * 参数说明：
 * mask_txt： 水印文字 （必填）
 * load_mark: 0|1 是否页面加载后就加载水印
 * color: 水印文字颜色
 * font_size：水印文字大小，默认为16号字体
 * front_font：字体，默认微软雅黑
 * front_x：第一个水印位置X坐标
 * front_y：第一个水印位置Y坐标
 * front_x_space: 调横向密度
 * front_y_space: 调纵向密度
 * width：水印宽度
 * height：水印高度
 * angle：水印角度
 * front_txt_alpha：前景水印文字透明度
 * front_rows：显示几行，留空为横向平铺
 * front_cols：显示几列,留空为纵向平铺
 * 调用示例1(整个页面打水印)：
 * <script type="text/javascript" src="iwatermark.js?mask_txt=linklin"></script>
 * 调用示例2(页面某个区域打水印)：
 * <script type="text/javascript" src="iwatermark.js?mask_txt=linklin&load_mark=0"></script>
 * <script type="text/javascript">
 * WMJS.watermark($("#task_info")[0]);
 * 如果引入js不加mask_txt参数时，可以这样调用WMJS.watermark($("#task_info")[0],'linklin');
 * </script>
 * 注意：iwatermark.js脚本要在body后引入
 *  ----------------------------------------------------------------------------------
 */

if (!this.WMJS) {
    WMJS = function() {
        this.options = {
            width: 16,
            height: 16,
            font_size: 16,
            mask_txt: '',
            show_obj: document.getElementById('app'),
            load_mark: 1,
            pageHeight: 0,
            reloadTimer: '',
            is_load: false,
            color: '#c9c9c9', //#33e3e3
            angle: 35,
            front_x: 60,
            front_y: 60,
            front_font: 'msyhbd',
            front_x_space: 180,
            front_y_space: 150,
            front_txt_alpha: '0.5',
            front_rows: 0,
            front_cols: 0,
            original_front_rows: 0,
            original_front_cols: 0
        }
    };

    /**
     * body或者div（页面部分区域）显示水印
     * @param object load_mark_obj 页面html对象，如document.body 或者$('body')[0] 或者 $('#div')[0]
     * @param string my_mask_txt   要显示的水印字
     */
    WMJS.prototype.watermark = function(load_mark_obj, my_mask_txt) { //load_mark_obj是元素对象(调用示例：WMJS.watermark($("#task_info")[0]);)
        if (this.options.is_load) {
            return false;
        }
        this.options.is_load = true;
        if (!load_mark_obj && this.options.load_mark == 0) {
            console.log("error: no load_mark_obj and load_mark is false!");
            return false;
        }
        if (my_mask_txt) {
            this.options.mask_txt = my_mask_txt;
            this.options.width = 16 * this.options.mask_txt.length;
        }
        this.remove_old_watermask();
        // console.log(load_mark_obj, 'load_mark_objload_mark_objload_mark_objload_mark_obj')
        // console.log(document.getElementById('main-container'), 'load_mark_objload_mark_objload_mark_objload_mark_obj')
        var wmTarget = document.getElementById('app') || load_mark_obj || document.body;
        if (!wmTarget) {
            console.log("error: no load_mark_obj!");
            return false;
        }
        this.options.show_obj = wmTarget;
        var oTemp = document.createDocumentFragment(),
            wmTargetWidth = wmTarget.scrollWidth,
            wmTargetHeight = wmTarget.scrollHeight,
            max_width = wmTargetWidth,
            max_height = wmTargetHeight;
        this.options.pageHeight = wmTarget.offsetHeight;
        if (load_mark_obj && parseInt(wmTarget.style.width) && wmTarget.style.width.indexOf('%') == -1) {
            max_width = parseInt(wmTarget.style.width) - 20,
                max_height = parseInt(wmTarget.style.height) - 20;
        }

        this.options.front_cols = Math.ceil((max_width - this.options.front_x) / (this.options.width + this.options.front_x_space));
        this.options.front_rows = Math.ceil((max_height - this.options.front_y) / (this.options.height + this.options.front_y_space));

        var mask_elem = document.createElement('div');
        var M = this.getRotation(-this.options.angle);
        mask_elem.id = 'mask_elem00';
        mask_elem.appendChild(document.createTextNode(this.options.mask_txt));
        mask_elem.style.webkitTransform = "rotate(-" + this.options.angle + "deg)";
        mask_elem.style.MozTransform = "rotate(-" + this.options.angle + "deg)";
        mask_elem.style.msTransform = "rotate(-" + this.options.angle + "deg)";
        mask_elem.style.OTransform = "rotate(-" + this.options.angle + "deg)";
        mask_elem.style.transform = "rotate(-" + this.options.angle + "deg)";
        mask_elem.style.visibility = "";
        mask_elem.style.position = "absolute";
        mask_elem.style.left = this.options.front_x + 'px';
        mask_elem.style.top = this.options.front_y + 'px';
        mask_elem.style.overflow = "hidden";
        //mask_elem.style.border="solid #eee 1px";
        mask_elem.style.opacity = this.options.front_txt_alpha;
        if (isIE9) {
            mask_elem.style.filter = "progid:DXImageTransform.Microsoft.Alpha(opacity=" + this.options.front_txt_alpha * 100 + ")";
        } else {
            mask_elem.style.filter = "progid:DXImageTransform.Microsoft.Alpha(opacity=" + this.options.front_txt_alpha * 100 + ") progid:DXImageTransform.Microsoft.Matrix(sizingMethod='auto expand', M11=" + M[0] + ", M12=" + M[1] + ", M21=" + M[2] + ", M22=" + M[3] + ")";
        }
        mask_elem.style.fontSize = this.options.font_size + "px";
        mask_elem.style.fontFamily = this.options.front_font;
        mask_elem.style.color = this.options.color;
        mask_elem.style.textAlign = "center";
        //mask_elem.style.width = this.options.width+'px';
        //mask_elem.style.height = this.options.height+'px';
        mask_elem.style.display = "block";
        mask_elem.style.pointerEvents = "none";
        mask_elem.style.zIndex = 99999;
        oTemp.appendChild(mask_elem);
        var x, y;
        for (var i = 0; i < this.options.front_rows; i++) {
            y = this.options.front_y + (this.options.front_y_space + this.options.height) * i;
            for (var j = 0; j < this.options.front_cols; j++) {
                x = this.options.front_x + (this.options.width + this.options.front_x_space) * j;
                if (i != 0 || j != 0) {
                    //clone
                    var new_elem = mask_elem.cloneNode(true);
                    new_elem.id = 'mask_elem' + i + j;
                    new_elem.style.left = x + 'px';
                    new_elem.style.top = y + 'px';
                    oTemp.appendChild(new_elem);
                }
            };
        };
        wmTarget.appendChild(oTemp);
        this.options.is_load = false;
        console.log('watermark ok');
        if (!WMJS.options.reloadTimer) {
            //添加事件监听
            window.onresize = function(e) {
                WMJS.watermark(WMJS.options.show_obj, WMJS.options.mask_txt);
            };
            //这里监控高度，定时两秒判断高度是否变化，如果变化则重新生成水印
            WMJS.options.reloadTimer = setInterval(function() {
                // console.log('reloadTimerreloadTimerreloadTimerreloadTimer')
                var tmpHeight = WMJS.options.show_obj.offsetHeight;
                // console.log(WMJS.options.pageHeight)
                if (tmpHeight != WMJS.options.pageHeight) {
                    WMJS.options.pageHeight = tmpHeight;
                    WMJS.watermark(WMJS.options.show_obj, WMJS.options.mask_txt);
                }
            }, 2000);
        }
    };

    WMJS.prototype.init = function(mask_txt) {
            isIE9 = document.all && document.addEventListener && !window.atob,
                int_key_data = {
                    'width': 1,
                    'font_size': 1,
                    'height': 1,
                    'front_x': 1,
                    'front_y': 1,
                    'original_front_rows': 1,
                    'original_front_cols': 1,
                    'front_x_space': 1,
                    'front_y_space': 1,
                    'angle': 1
                };
            this.options.mask_txt = mask_txt;
            this.options.width = 16 * this.options.mask_txt.length;
            if (this.options.color && this.options.color.substring(0, 1) != "#") {
                this.options.color = "#" + this.options.color;
            }
            if (this.options.color.length > 7) {
                this.options.color = '#c9c9c9';
            }
            for (var item in this.options) {
                if (int_key_data[item]) {
                    this.options[item] = parseInt(this.options[item]); //int型处理
                } else if (isNaN(this.options[item])) {
                    // this.options[item] = this.options[item].replace("\"", '').replace("'", '').replace("<", '').replace(">", ''); //过虑html代码
                }
            }
            //添加事件
            if (this.options.mask_txt) {
                window.onload = function() {
                    WMJS.watermark(WMJS.options.show_obj, WMJS.options.mask_txt); //页面初始化
                };
                console.log(WMJS.options.show_obj, 'WMJS.options.show_obj,WMJS.options.show_obj,WMJS.options.show_obj,init')
                WMJS.watermark(WMJS.options.show_obj, WMJS.options.mask_txt);
            }
        }
        /**
         * 获取脚本src参数
         * @returns {Array}
         */
    WMJS.prototype.get_src_params = function() {
            var params_src = [];
            var scripts = document.getElementsByTagName('script');
            var currentScript = scripts[scripts.length - 1];
            var jssrc = currentScript.src;
            var src = jssrc.substring(jssrc.lastIndexOf("?") + 1, (jssrc.length + 1));
            var params_tmp = src.split("&");
            for (var i = 0; i < params_tmp.length; i++) {
                var myparams = params_tmp[i].split("=");
                params_src[myparams[0]] = myparams[1];
            }
            return params_src;
        }
        /**
         * 删除水印
         */
    WMJS.prototype.remove_old_watermask = function() {
        for (var i = 0; i <= this.options.front_rows; i++) {
            for (var j = 0; j <= this.options.front_cols; j++) {
                var rmElm = document.getElementById('mask_elem' + i + j);
                if (rmElm) {
                    rmElm.parentNode.removeChild(rmElm);
                }
            }
        }
    }
    WMJS.prototype.getRotation = function(deg) {
            var deg2rad = Math.PI * 2 / 360;
            rad = deg * deg2rad;
            costheta = Math.cos(rad);
            sintheta = Math.sin(rad);
            M11 = costheta;
            M12 = -sintheta;
            M21 = sintheta;
            M22 = costheta;
            return [M11, M12, M21, M22];
        }
        //new WMJS对象，页面调用
    WMJS = new WMJS();
}
