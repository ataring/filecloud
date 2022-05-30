var $$ = mdui.JQ;
    //监听鼠标右击事件 / 移动端长按事件
    $$(document).on('contextmenu', function (e) {
        // console.log(e);

        //0：移动端长按（iOS 测试未通过）
        //2：电脑端右键
        if(e.button === 2 || e.button === 0) {
            e.preventDefault();//阻止冒泡，阻止默认的浏览器菜单

            //（修改了这里）鼠标点击位置，相对于页面
            var _x = e.pageX,
                _y = e.pageY;

            let $div = $$("<div></div>").css({
                position: 'absolute',
                top: _y+'px',
                left: _x+'px',
            });
            $$('body').append($div);//创建临时DOM

            // anchorSelector 表示触发菜单的元素的 CSS 选择器或 DOM 元素
            // menuSelector 表示菜单的 CSS 选择器或 DOM 元素
            // options 表示组件的配置参数，见下面的参数列表
            // 完整文档列表：https://doc.nowtime.cc/mdui/menu.html
            var inst = new mdui.Menu($div, '#menu');
            inst.open();//打开菜单栏
            $div.remove();//销毁创建的临时DOM
        }
    });
