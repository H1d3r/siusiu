package routers

import (
	"log"
	"os"

	"github.com/ShangRui-hash/siusiu/controllers"
	"github.com/ShangRui-hash/siusiu/pkg/exec"
	"github.com/abiosoft/ishell"
)

//Init 初始化路由
func Init(shell *ishell.Shell) error {
	//第三方工具
	shell.AddCmd(&ishell.Cmd{
		Name: "crawlergo",
		Help: "使用chrome headless模式进行URL收集的浏览器爬虫",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/crawlergo", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "nmap",
		Help: "主机发现、端口扫描、服务扫描、版本识别",
		Func: func(c *ishell.Context) {
			exec.Docker("instrumentisto/nmap:7.92", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "sqlmap",
		Help: "SQL注入攻击工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/root/.local/share/sqlmap/output/", "-w", "/root/.local/share/sqlmap/output/", "rickshang/sqlmap:1.6.1"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "dirsearch",
		Help: "目录爆破工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/root/reports", "-w", "/root/reports", "rickshang/dirsearch"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "ffuf",
		Help: "模糊测试工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/ffuf:1.3.1"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "whatweb",
		Help: "web指纹识别",
		Func: func(c *ishell.Context) {
			exec.Docker("bberastegui/whatweb", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "tool-helper",
		Help: "获取工具的帮助文档",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/tool-helper", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "stegseek",
		Help: "爆破隐写术密码",
		Func: func(c *ishell.Context) {
			//获取当前目录
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/steg", "rickdejager/stegseek"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "steghide",
		Help: "隐写术工具",
		Func: func(c *ishell.Context) {
			// docker run -it --rm -v $(pwd):/src -w /src bartimar/steghide info doubletrouble.jpeg
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/src", "-w", "/src", "bartimar/steghide"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "http3-client",
		Help: "支持http3的客户端",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/http3-client", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "gopherus",
		Help: "ssrf漏洞gopher协议payload生成工具",
		Func: func(c *ishell.Context) {
			exec.Docker("rickshang/gopherus", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "firefox-decrypt",
		Help: "firefox浏览器密码提取工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/firefox-decrypt"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "hydra",
		Help: "弱口令爆破工具",
		Func: func(c *ishell.Context) {
			//1.查询本地是否有hydra
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/thc-hydra"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "wfuzz",
		Help: "web应用fuzz工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "ghcr.io/xmendez/wfuzz", "wfuzz"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "cewl",
		Help: "爬去网站关键字以生成字典",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "nocflame/cewl"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "davtest",
		Help: "webdav利用工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/davtest"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "xray",
		Help: "漏洞扫描器",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp/xray", "-w", "/tmp/xray", "rickshang/xray:1.8.4private"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "xray-listen",
		Help: "xray监听工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "-v", currentDir + ":/tmp/xray", "-w", "/tmp/xray", "-p", "7777:7777", "rickshang/xray:1.8.4private", "webscan", "--listen", "0.0.0.0:7777", "--html-output", "xray-result.html"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "ds_store_exp",
		Help: ".DS_Store 文件泄漏利用脚本",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/ds_store_exp"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "svn-exp",
		Help: "svn-exp 文件泄漏利用脚本",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/svn-exp"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "gobuster",
		Help: "目录扫描工具（dirsearch拉跨时备用）",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/gobuster"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "pocsuite3",
		Help: "poc测试框架",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "rickshang/pocsuite3"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "smtp-user-enum",
		Help: "SMTP用户名枚举工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/smtp-user-enum"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "smbmap",
		Help: "smb服务利用工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/smbmap"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "searchsploit",
		Help: "exp/poc搜索工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/searchsploit"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "jsfinder",
		Help: "从js源码提取URL，子域名的工具。",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/tmp", "-w", "/tmp", "rickshang/jsfinder"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "403bypasser",
		Help: "403绕过工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/403bypasser", "-w", "/403bypasser", "rickshang/403bypasser"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "cve-2018-15473-exp",
		Help: "ssh 用户名枚举漏洞利用工具",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/cve-2018-15473-exp", "-w", "/cve-2018-15473-exp", "rickshang/cve-2018-15473-exp"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "waybackurls",
		Help: "查询指定域名的历史页面",
		Func: func(c *ishell.Context) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Println("os.Getwd failed,err:", err)
				return
			}
			params := append([]string{"run", "--rm", "-it", "--network", "host", "-v", currentDir + ":/waybackurls", "-w", "/waybackurls", "rickshang/waybackurls"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "wpscan",
		Help: "wordpress漏洞扫描工具",
		Func: func(c *ishell.Context) {
			params := append([]string{"run", "--rm", "-it", "--network", "host", "wpscanteam/wpscan"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "gau",
		Help: "根据域名进行被动url收集(open threat+wayback machine+common crawl)",
		Func: func(c *ishell.Context) {
			params := append([]string{"run", "--rm", "-i", "--network", "host", "sxcurity/gau:latest"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "amass",
		Help: "信息收集工具",
		Func: func(c *ishell.Context) {
			params := append([]string{"run", "--rm", "-it", "--network", "host", "caffix/amass"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "payloads-all-the-things",
		Help: "payloads大全",
		Func: func(c *ishell.Context) {
			params := append([]string{"run", "--rm", "-it", "--network", "host", "rickshang/payloads-all-the-things"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "linkfinder",
		Help: "用于发现 JavaScript 文件中的端点及其参数",
		Func: func(c *ishell.Context) {
			params := append([]string{"run", "--rm", "-it", "--network", "host", "rickshang/linkfinder"}, c.Args...)
			exec.CmdExec("docker", params...)
		},
	})
	//未找到命令时
	shell.NotFound(controllers.NotFoundHandler)
	return nil
}
