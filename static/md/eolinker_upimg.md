* eolinker默认没有开启图片的上传,用的3.5开源免费版,不知道是否收费版支持。eolinker是一款比较强大的接口管理工具，官方网址:www.eolinker.com
- 今天接到需求，需要eolinker写详细说明文档支持图片上传,eolinker使用两款编辑器实现,一款是wangEditor，我用的开源版使用的是wangEditor2。文档地址: https://www.kancloud.cn/wangfupeng/wangeditor2
- 官网最新的版本已经到了wangEditor3。文档不同,这点要注意。另一款就是editorMd
- 开始修改实现图片上传,这里选取修改wangEditor。
- 修改scripts/app-1ee7ea79f2.js， 文件名稍微不同,都是app-开头的，只有这一个文件。搜索 `config.menus`，找到周围有， 菜单配置增加`"img"`，在该配置前面增加上传配置 `l.config.uploadImgUrl="/upload.php"`。其中 l是 new wangEditor的参数名，如果不同对应修改即可


<img src="/static/md/imgs/1.png" style="width:100%;" />
<!-- ![](/static/md/imgs/1.png){:width="1491"} -->

- eoliker根目录简单增加增加上传文件,保存到upload目录
  
  demo

```php
    <?php
    $file = array_pop($_FILES);
    if(empty($file)){

        die("上传参数错误:使用multipart上传");
    }
    $uploadDir = "upload/" . date("Ymd");
    if(!file_exists($uploadDir)){
        @mkdir($uploadDir);
    }
    //生成随即文件名
    $newName = uniqid();
    //补气后缀
    //var_dump(pathinfo($file["name"])["extension"]);die();
    $newName .=  pathinfo($file["name"])["extension"] ? "." . pathinfo($file["name"])["extension"] : '';
    //保存文件
    if(move_uploaded_file($file["tmp_name"], $uploadDir . "/" . $newName)){

        echo "/" . $uploadDir . "/" . $newName; 
    }else{
        echo "move error";
    }
```
