# siusiu （suitesuite）
一款基于docker的渗透测试工具箱，致力于做到渗透工具随身携带、开箱即用、按需下载。

## Features

siusiu 将常用的渗透测试工具都封装为了docker镜像，并推送到了dockerhub中。并且为用户提供了一个shell控制台，通过该控制台，可以：

- 查看第三方安全工具列表
- 按需下载第三方安全工具
- 运行第三方安全工具
- 查看第三方安全工具的说明文档与使用样例（通过demos命令）

同时siusiu也支持非交互模式，便于siusiu被其他程序调用,例如:siusiu exec help

## Usage：
```
siusiu:/ > help

Commands:
  403bypasser                  403绕过工具
  amass                        信息收集工具
  arjun                        传参发现工具
  cewl                         爬取网站关键字以生成字典
  clear                        clear the screen
  cloudfail                    用于寻找cloudflare背后的真实IP
  crawlergo                    使用chrome headless模式进行URL收集的浏览器爬虫
  cve-2018-15473-exp           ssh 用户名枚举漏洞利用工具
  davtest                      webdav利用工具
  dirsearch                    目录爆破工具
  ds_store_exp                 .DS_Store 文件泄漏利用脚本
  exit                         exit the program
  fetcher                      用于将指定目录制作为字典
  ffuf                         模糊测试工具
  firefox-decrypt              firefox浏览器密码提取工具
  gau                          根据域名进行被动url收集(open threat+wayback machine+common crawl)
  githack:bugscanteam          git 泄漏利用工具(会下载.git文件夹，方便检索历史版本)
  githack:lijiejie             git 泄漏利用工具(只会下载当前版本)
  gobuster                     目录扫描工具（dirsearch拉跨时备用）
  gopherus                     ssrf漏洞gopher协议payload生成工具
  help                         display help
  http3-client                 支持http3的客户端
  hydra                        弱口令爆破工具
  input-scanner                用于提取js文件中的URL
  jsfinder                     从js源码提取URL，子域名的工具。
  ksubdomain                   子域名爆破工具
  linkfinder                   用于发现 JavaScript 文件中的端点及其参数
  nmap                         主机发现、端口扫描、服务扫描、版本识别
  pacu                         aws 利用框架
  paramspider                  参数挖掘工具
  payloads-all-the-things      payloads大全
  php_mt_seed                  PHP伪随机数种子破解器
  pocsuite3                    poc测试框架
  rip-hg.pl                    .hg 文件泄漏利用脚本(可以下载.gh文件夹，方便检查历史版本)
  rip-svn.pl                   .svn 文件泄漏利用脚本(可以下载.svn文件夹，方便检查历史版本)
  searchsploit                 exp/poc搜索工具
  smbmap                       smb服务利用工具
  smtp-user-enum               SMTP用户名枚举工具
  sqlmap                       SQL注入攻击工具
  sqlmapapi                    sqlmap api
  steghide                     隐写术工具
  stegseek                     爆破隐写术密码
  subfinder                    子域名查询工具
  svn-exp                      svn-exp 文件泄漏利用脚本
  tool-helper                  获取工具的帮助文档
  waybackurls                  查询指定域名的历史页面
  wfuzz                        web应用fuzz工具
  whatweb                      web指纹识别
  wpscan                       wordpress漏洞扫描工具
  xray                         漏洞扫描器
  xray-listen                  xray监听工具

```

## Tested On  

新版本基于docker构建，只要是安装了docker的主机的，都可以正常运行。老版本基于shell脚本构建，只能在linux和mac环境下运行。

## Installation

### 下载二进制文件

- step1. 点击 Releases ，下载对应的版本
- step2. 修改可执行文件名为 siusiu
- step3. 赋予该文件可执行权限

### git 安装

```shell
git clone --depth 1 https://github.com/ShangRui-hash/siusiu.git
cd siusiu
go env -w  GOPROXY=https://goproxy.cn,https://goproxy.io,direct
go build -o siusiu  
```

### go安装

```shell
go get github.com/ShangRui-hash/siusiu@latest
```

或者

```shell
go install github.com/ShangRui-hash/siusiu@latest 
```

## Screenshots

如果用户未安装pocsuite3，则自动下载 pocsuite3,然后自动运行 
![avatar](https://img-blog.csdnimg.cn/20211006160456729.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)

在siusiu控制台中运行sqlmap和dirsearch
![avatar](https://img-blog.csdnimg.cn/20211006160557298.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)

## 使用技巧

### 使用方式
siusiu 有两种使用方式，一种是交互式控制台，直接输入siusiu即可启动控制台， tab 键可以自动补全命令，help 可以查看所有命令。另一种是非交互方式，通过`siusiu exec <tool-name>` 来运行指定工具


### 搜索指定工具
例如：搜索和web相关的工具
```
siusiu exec help  | grep web
  davtest                      webdav利用工具
  wfuzz                        web应用fuzz工具
  whatweb                      web指纹识别
```
```
siusiu exec help  | grep 扫描    
  gobuster                     目录扫描工具（dirsearch拉跨时备用）
  nmap                         主机发现、端口扫描、服务扫描、版本识别
  wpscan                       wordpress漏洞扫描工具
  xray                         漏洞扫描器
```
### 构建工作流
使用linkfinder获取js文件中的url,然后用wc -l 统计数目
```shell
siusiu exec linkfinder -i target.js -o cli | wc -l
``` 
使用fetcher 将当前目录制作为目录字典
```
siusiu exec fetcher ./ > dict.txt 
``` 

### 关于searchsploit的一些使用说明
由于 searchsploit 所有的exp和poc都封装在容器中，所以如果想要将exp/poc拷贝到本机，可以通过下面的命令：
step1. 查看exp/poc的具体位置
```
searchsploit -p <id>
```
step2. 拷贝exp/poc到当前目录
```
searchsploit cp <exp/poc的绝对路径>
```
### 关于tool-helper
tool-helper 是一个封装了工具使用样例的镜像，通过tool-helper <工具名>可以在不离开shell的情况下查看工具的使用样例
```
tool-helper
Usage: ./help <filename>
        crawlergo
        dirsearch
        ffuf
        gdb
        gobuster
        hydra
        masscan
        nmap
        searchsploit
        smbmap
        spawn_shell
        sqlmap
        steghide
        stegseek
        url-collector
        wafw00f
        xray
```
## Contact me  
### QQ 交流群

<img src="https://store.heytapimage.com/cdo-portal/feedback/202301/04/ed1d5ac9f0c48af0a154037fb892024f.png" height="250px" width="250px" alt="图片.png" title="图片.png" referrerPolicy="no-referrer" />

### 微信公众号
<img src="https://i.imgtg.com/2023/03/10/YwJ3S.jpg" height="250px" width="250px" alt="图片.png" title="图片.png" referrerPolicy="no-referrer" >

## QA

问：我喜欢作者怎么办？  
答：⁄(⁄ ⁄•⁄ω⁄•⁄ ⁄)⁄  

问：能不能一键日卫星？  
答：~~不能，至少目前现阶段不可以。抱歉。么么哒~~ 可以一键日卫星，我的小可爱。  

问：How to build and install siusiu on raspberry pi ?  
答：首先，请不要说英文，请用普通话。然后， I can not replay your question you just mentioned in mandarin excuse me. If you want to run it on embed platforms like raspberry pi 3 model b, you must need to know that these platforms have a ARM core, for example, pi 3b's cpu is Contex-A53 with a ARM v8 architecture, so, set GOOS=linux and GOARCH=aarch64 then run go build main.go in your shell, enjoyed it !

