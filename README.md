## 一个跨平台局域网共享剪贴板文字的工具

### 介绍

**背景**

Mac 与iPad之间可以共享剪切板，用起来非常方便，但是这个功能仅限在苹果设备之间，在苹果、Windows和Linux之间缺少类似的工具，于是便就有了这个工具。

**功能**

在局域网共享不同电脑之间的剪贴版文字，内容传输加密，支持局域网多个用户同时使用，支持同一用户不同设备之间共享。

### 安装

点击 https://github.com/FanXingGuo/IAO/releases/tag/v1.0 （或者本页面 右栏 Releases）

备用地址：链接: https://pan.baidu.com/s/1w_mXzFfHSTmuFkIzpAj9bQ  密码: glt9

选择对应的平台版本，下载、解压：

##### Windows平台（amd64）：

双击运行，如果提示请求网络访问，点击允许

##### **Linux(amd64、arm、arm64)、Mac平台（amd64）：**

双击运行或者使用命令行，如果共享时没有反应，请添加管理员 权限：

linux: requires 'xclip' or 'xsel' command to be installed


```bash
sudo ./IAO
```

运行时截图

![20220126](http://cdn.51dream.top/blog/20220126.png)

### 使用

工具会自动检测剪贴板变化，当发生变化时自动同步到其他开启本程序且相同文件名（不包括后缀）的电脑剪切板中

不同设备同步的前提条件（同时满足）：

1. 在同一个网络下（连接的是同一个路由器）
2. 本程序打开，有网络访问权限
3. 相同文件名（若有后缀名，后缀名之前的文件名）

**关于相同文件名：**

区分局域网不同用户使用，默认相同文件名的传递有效，Windows 平台不包括后缀名，如程序 IAO.exe 剪贴板内容会传递给Linux或Mac下的 “IAO”（Mac、Linux 可执行文件 无后缀名），Linux：程序名为zhangsan会共享给 Windows平台下zhangsan.exe，lisi.exe、lisi 则不会共享 

**关于异常：panic: listen udp 0.0.0.0:9999: bind: address already in use**

报错显示端口已被占有。IAO 使用的是9999端口，如果其他程序占用该端口，请先关闭占用的程序。或者有可能，IAO打开了两次，第二次打开报错端口被占用


