# myt-androidApi-demo

1.启动参数说明，-help可查看参数

    -command string                                                                                                           
    操作指令, snap--获取屏幕截图, devRandom--随机设备信息, devCustom--自定义设备信息, hideApp--隐藏应用 (default "snap")

    -host string                                                                                                              
    安卓host 

    -p string
    安卓容器接口的端口，例9082，10005 

    -app string                                                                                                               
    隐藏app时的可选参数，可填多个用逗号分割，例com.ss.android.ugc.aweme,com.mth_player.oaid

    -l int
    获取屏幕截图时的可选参数，可选值有1,2,3 (default 3)

2.安卓容器接口的端口说明

    当Android容器为桥接网络模式时 ip为安卓容器实例IP port为9082

    当Android共享ip时 ip为当前主机ip port为动态计算得出 计算方法为: 
    第一个Android容器实例 10005 第二个 10008 第三个 10011 .... 以此类推 公式为 (索引下标-1) *3 + 10005

3.获取屏幕截图
    
    .\demo.exe -command=snap -host='192.168.100.10' -p=10005 -l=3

4.随机设备信息

    .\demo.exe -command=devRandom -host='192.168.100.10' -p=10005

5.自定义设备信息

    .\demo.exe -command=devCustom -host='192.168.100.10' -p=10005

    自定义设备信息是设置的dev.json中的数据，可以自行修改自定义

6.隐藏应用
    
    .\demo.exe -command=hideApp -host='192.168.100.10' -p=10005 -app=com.ss.android.ugc.aweme,com.mth_player.oaid

    app参数里是应用的包名

