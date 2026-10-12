package main

import (
	"encoding/json"
	"fmt"
)

type QuertionItem struct {
	QuertionID        string `json:"题号"`
	QuertionStatement string `json:"查询语句"`
}

type AnswerRoot struct {
	PlayerName string         `json:"选手"`
	PackageID  string         `json:"参赛包编号"`
	Resp       []QuertionItem `json:"答案"`
}

func main() {
	data := AnswerRoot{
		PlayerName: "刘子琪",
		PackageID:  "pkg-08e82c8d",
		Resp: []QuertionItem{
			{
				QuertionID:        "M001-S009",
				QuertionStatement: "请查询 IP 地址为 20.247.40.92 的资产。",
			},
			{
				QuertionID:        "M002-S004",
				QuertionStatement: "请查询 124.255.254.0/24 网段内的全部资产。",
			},
			{
				QuertionID:        "M003-S002",
				QuertionStatement: "请查询主域是 zendesk.com 的全部资产。",
			},
			{
				QuertionID:        "M004-S003",
				QuertionStatement: "搜索网站名为 www.amazon.com 的资产。",
			},
			{
				QuertionID:        "M005-S004",
				QuertionStatement: "我想找域名里带有 vpn 的所有网站。",
			},
			{
				QuertionID:        "M006-S002",
				QuertionStatement: "搜索位于大不列颠及北爱尔兰联合王国的资产。",
			},
			{
				QuertionID:        "M007-S007",
				QuertionStatement: "搜索韩国江原道的资产。",
			},
			{
				QuertionID:        "M008-S009",
				QuertionStatement: "搜索德国巴伐利亚州的München市的资产。",
			},
			{
				QuertionID:        "M009-S008",
				QuertionStatement: "搜索自治系统号为 4134 的资产。",
			},
			{
				QuertionID:        "M010-S009",
				QuertionStatement: "我想找位于荷兰、使用亚马逊的云计算、对象存储或应用托管服务的资产。",
			},
			{
				QuertionID:        "M011-S005",
				QuertionStatement: "我想查河北地区的华为云资产。",
			},
			{
				QuertionID:        "M012-S009",
				QuertionStatement: "请查询开放 3000 端口的资产。",
			},
			{
				QuertionID:        "M013-S003",
				QuertionStatement: "搜索标题包含 Jenkins 或正文包含 jenkins、国家不是中国，且更新时间在 2026-08-01 之后的资产。",
			},
			{
				QuertionID:        "M014-S002",
				QuertionStatement: "我想找中国的 DNS 服务。",
			},
			{
				QuertionID:        "M015-S001",
				QuertionStatement: "fofa有识别SOCKS的v5版本吗？",
			},
			{
				QuertionID:        "M016-S006",
				QuertionStatement: "搜索 SSH 协议且端口为 22022 的资产。",
			},
			{
				QuertionID:        "M017-S001",
				QuertionStatement: "我想看使用 TCP 协议的数据。",
			},
			{
				QuertionID:        "M018-S006",
				QuertionStatement: "我想找互联网上提供 SNMP 网络管理服务的资产，不限定端口。",
			},
			{
				QuertionID:        "M019-S010",
				QuertionStatement: "我想要搜索Memcached缓存服务协议在FOFA上的数据，地域限定为美国。",
			},
			{
				QuertionID:        "M020-S005",
				QuertionStatement: "请帮我查询真正跑在 TCP 上的 FTP 服务。",
			},
			{
				QuertionID:        "M021-S010",
				QuertionStatement: "请帮我查询真正的 SSH 服务，并且它返回的信息里包含 OpenSSH_10.4。",
			},
			{
				QuertionID:        "M022-S007",
				QuertionStatement: "我想找日本的海康威视摄像头资产。",
			},
			{
				QuertionID:        "M023-S002",
				QuertionStatement: "查 163.com 旗下的域名数据，要求它前面恰好还有两级子域，每一级都只有一个字符，例如 a.z.163.com。",
			},
			{
				QuertionID:        "M024-S005",
				QuertionStatement: "我想找产品版本号为 10.0.0 的资产，不限定具体产品。",
			},
			{
				QuertionID:        "M025-S010",
				QuertionStatement: "我想按产品门类查找连接应用与底层基础服务的中间层软件这一类资产，不限定具体厂商和型号。",
			},
			{
				QuertionID:        "M026-S002",
				QuertionStatement: "搜索网页正文中出现 NoSuchUpload 错误标记的资产。",
			},
			{
				QuertionID:        "M027-S008",
				QuertionStatement: "我想找响应信息中出现 Content-Security-Policy-Report-Only 的资产。",
			},
			{
				QuertionID:        "M028-S011",
				QuertionStatement: "搜索状态码为 308 的网站类资产。",
			},
			{
				QuertionID:        "M029-S010",
				QuertionStatement: "请帮我查询响应信息有 SSH-2.0-OpenSSH_6.1 的资产。",
			},
			{
				QuertionID:        "M030-S003",
				QuertionStatement: "搜索header包含 Set-Cookie:PHPSESSID 的资产。",
			},
			{
				QuertionID:        "M031-S009",
				QuertionStatement: "搜索引用 jquery.min.js 文件的网页。",
			},
			{
				QuertionID:        "M032-S010",
				QuertionStatement: "查荷兰阿姆斯特丹开放 554 端口、响应原文含有 200 OK 的资产。这里的 200 OK 是服务返回的文本。",
			},
			{
				QuertionID:        "M033-S003",
				QuertionStatement: "请帮我查询由 GlobalSign 组织签发、而且证书还没有过期的资产。",
			},
			{
				QuertionID:        "M034-S007",
				QuertionStatement: "我想找证书持有者名称是 apple.com、但域名不是 apple.com 的域名数据。",
			},
			{
				QuertionID:        "M035-S010",
				QuertionStatement: "请查询证书全文中包含 TrustAsia 的资产。",
			},
			{
				QuertionID:        "M036-S007",
				QuertionStatement: "搜索证书持有者组织名称包含 Alibaba 这个关键词的资产。",
			},
			{
				QuertionID:        "M037-S008",
				QuertionStatement: "查 sohu.com 下的资产，要求 Banner 有内容，但其中不包含 apache 关键词。",
			},
			{
				QuertionID:        "M038-S009",
				QuertionStatement: "证书里面只要有adobe.com这个根域名的都给我找出来。",
			},
			{
				QuertionID:        "M039-S001",
				QuertionStatement: "我想搜索证书持有者名称根域是 qq.com 的资产。",
			},
			{
				QuertionID:        "M040-S011",
				QuertionStatement: "请查询证书包含 godaddy，但根域名不是 godaddy.com 的资产。",
			},
			{
				QuertionID:        "M041-S001",
				QuertionStatement: "我想找证书信息里包含 google.com这个域名的资产，蜜罐和欺诈网站都不要。",
			},
			{
				QuertionID:        "M042-S003",
				QuertionStatement: "我想要证书持有者名称包含example.com特征的所有资产或者证书包含example.com这个域名都行，要包含裸域和子域。",
			},
			{
				QuertionID:        "M043-S004",
				QuertionStatement: "找仍在使用 TLS 1.0，且证书尚未过期的资产。",
			},
			{
				QuertionID:        "M044-S007",
				QuertionStatement: "找 JARM 为 21d21d00021d21d00021d21d21d21dce319294781f0e12867fd06896a000ea、网站响应码为 307，且响应头包含 Content-Length: 0 的资产。",
			},
			{
				QuertionID:        "M045-S001",
				QuertionStatement: "我想查操作系统被识别为 linux 的资产。",
			},
			{
				QuertionID:        "M046-S013",
				QuertionStatement: "地域限定为中国大陆，查看没有绑定域名的网站类资产。",
			},
			{
				QuertionID:        "M047-S011",
				QuertionStatement: "我想找证书颁发者与持有者不匹配的资产。",
			},
			{
				QuertionID:        "M048-S001",
				QuertionStatement: "帮我找加拿大的资产，要求证书与域名不匹配。",
			},
			{
				QuertionID:        "M049-S010",
				QuertionStatement: "查找土耳其 TLS JA3S 指纹等于 f4febc55ea12b31ae17cfb7e614afda8 的资产。",
			},
			{
				QuertionID:        "M050-S004",
				QuertionStatement: "帮我找一下网络组织名称里包含Inc 的资产。",
			},
			{
				QuertionID:        "M051-S001",
				QuertionStatement: "我想找美国 OVH SAS 网络里开放 2087 端口、标题完整为 WHM 登录。",
			},
			{
				QuertionID:        "M052-S002",
				QuertionStatement: "请帮我查询 .io 域名下部署的 WordPress 资产。",
			},
			{
				QuertionID:        "M053-S009",
				QuertionStatement: "搜索厦门市 80 端口且网页状态码为 200 的资产。",
			},
			{
				QuertionID:        "M054-S009",
				QuertionStatement: "请帮我查询英国 DigitalOcean 上疑似提供 /api/show 接口的资产。",
			},
			{
				QuertionID:        "M055-S003",
				QuertionStatement: "我想要找中国大陆响应信息包含 Set-Cookie，但是正文有 /admin 的资产。",
			},
			{
				QuertionID:        "M056-S008",
				QuertionStatement: "在 baidu.com 的资产中，正文出现 Login、admin、Dashboard、管理任一项，或响应头出现 Set-Cookie，任意一种命中都要。",
			},
			{
				QuertionID:        "M057-S012",
				QuertionStatement: "在 13.0.0.0/8 网段内查网页：响应码必须是 200，正文包含 Welcome 或 Sign in 或 Dashboard 中任意一项。",
			},
			{
				QuertionID:        "M058-S012",
				QuertionStatement: "查德国与 amazonaws.com 相关的资产，访问地址、主域、CNAME 根域、证书全文、证书域名这五处任一处命中都算。",
			},
			{
				QuertionID:        "M059-S006",
				QuertionStatement: "找 143 端口上证书序列号为 1 的资产，同时证书中有 TAK、ATAK、WinTAK、Squad、CoT 中任意一个词。",
			},
			{
				QuertionID:        "M060-S010",
				QuertionStatement: "我想找 113.247.228.0/24 网段里开放 80、443、8080 或 8443 任意一个端口的资产。",
			},
			{
				QuertionID:        "M061-S005",
				QuertionStatement: "请查询标题包含 Sign in 或 Login、端口为 80 且不是 IPv6 的资产。",
			},
			{
				QuertionID:        "M062-S010",
				QuertionStatement: "请帮我查询新西兰地区放 80 或 443 端口且是使用了云服务商的资产。",
			},
			{
				QuertionID:        "M063-S001",
				QuertionStatement: "搜索 Banner 同时包含 IIS 和 1.，或者响应头包含 IIS 的资产。",
			},
			{
				QuertionID:        "M064-S004",
				QuertionStatement: "找日本已绑定域名的网站：443 端口、响应码 200、证书有效；标题包含 大学、学部、研究 任意一个词，或者正文包含 入学案内、学生 任意一个词就算。",
			},
			{
				QuertionID:        "M065-S004",
				QuertionStatement: "在 baidu.com 主域内找资产：标题有系统、平台、登录、管理或 API，或者子域名中有 api、admin、dev、test、uat、pre，任意一项命中即可。",
			},
			{
				QuertionID:        "M066-S006",
				QuertionStatement: "找被归为邮件系统、标题包含 Mail，但国家不是美国的资产。",
			},
			{
				QuertionID:        "M067-S008",
				QuertionStatement: "找中国大陆开放 8888 端口的资产，香港、澳门、台湾都排除。",
			},
			{
				QuertionID:        "M068-S005",
				QuertionStatement: "找正文包含 Jenkins，但标题不包含同一个关键词的资产。",
			},
			{
				QuertionID:        "M069-S001",
				QuertionStatement: "我想找台湾 AS20940 上 443 端口、使用 TLS 1.3 的资产，但 Let's Encrypt 和 ZeroSSL 组织签发的证书不要。",
			},
			{
				QuertionID:        "M070-S003",
				QuertionStatement: "找日本以外的 HFS 文件服务，状态码为 200，标题完整等于 HFS /，部署在亚马逊云的排除。",
			},
			{
				QuertionID:        "M071-S006",
				QuertionStatement: "请帮我查询不包含港澳台的中国 Grafana 资产。",
			},
			{
				QuertionID:        "M072-S008",
				QuertionStatement: "请帮我查询以 .edu.ar 结尾的高校网站，只看 80 或 443 端口，服务是 Nginx 和 WAF 的不要。",
			},
			{
				QuertionID:        "M073-S003",
				QuertionStatement: "我有 IP 1.1.1.1，想查这台 IP 已经绑定域名的资产记录。",
			},
			{
				QuertionID:        "M074-S008",
				QuertionStatement: "请帮我查询证书里面有根域 aliyun.com、证书当前有效，但资产本身没有绑定域名的结果。",
			},
			{
				QuertionID:        "M075-S006",
				QuertionStatement: "找中国已绑定域名的资产，响应头包含 x-qiniu-zone，但主域不是 qiniudn.com。",
			},
			{
				QuertionID:        "M076-S004",
				QuertionStatement: "找网站标题为空、响应码 200 的资产，fid 和图标哈希都不为空；访问域名包含 admin，且根域满足以下任一种形态：以 .cn 结尾，或 .com 前恰好是四个字符。",
			},
			{
				QuertionID:        "M077-S009",
				QuertionStatement: "我想找网页源码中包含这段原样文本的资产：class=\"container\"。",
			},
			{
				QuertionID:        "M078-S009",
				QuertionStatement: "在网页正文中搜索这段 JSON 异常信息片段，保留双引号和反斜杠：\n\"exception\": \"Symfony\\Component\\ErrorHandler\\Error\\FatalError\"",
			},
			{
				QuertionID:        "M079-S006",
				QuertionStatement: "我想要搜索这个 IP 的整个 B 段：40.88.14.254。",
			},
			{
				QuertionID:        "M080-S006",
				QuertionStatement: "查 47.96.0.1 到 47.96.0.5 这个闭区间里的 IP，要求开放 80 或 443 端口。",
			},
			{
				QuertionID:        "M081-S011",
				QuertionStatement: "找 Grafana 资产：响应头哈希、正文哈希、网站图标哈希、fid 指纹这四项中，至少两项有值。",
			},
			{
				QuertionID:        "M082-S001",
				QuertionStatement: "搜索以下主域资产，都是或的关系，不管这条资产是绑定的任意主域还是他的证书里面有任意主域都算。但是证书要满足被fofa标记为匹配并且有效，最后证书过期时间晚于 2024-01-01 00:00:00 且早于2026-08-30 18:17:49\n主域有：kuaihuoyun.com/wlhyos.cn/wlhyos.com/56ctms.com/xiaokuaikeji.com/zhihuiwuliu.org.cn",
			},
			{
				QuertionID:        "M083-S008",
				QuertionStatement: "请把下面的 Shodan 查询转换成 FOFA：\"Server: nginx\" http.title:\"Welcome\"",
			},
			{
				QuertionID:        "M084-S004",
				QuertionStatement: "把这条 Google Dork 转换为 FOFA：intitle:\"Grafana\" intext:\"JavaScript\"",
			},
			{
				QuertionID:        "M085-S001",
				QuertionStatement: "把这个 urlscan.io 搜索意图转换成 FOFA：HTML 包含 login，并且域名包含 bet、casino 或 kasino。",
			},
			{
				QuertionID:        "M086-S006",
				QuertionStatement: "把这条 Hunter 查询转成 FOFA：domain=\"google\"\u0026\u0026domain.suffix!=\"google.com\"\u0026\u0026web.title=\"\"",
			},
			{
				QuertionID:        "M087-S009",
				QuertionStatement: "将这条 Censys 搜索转换为 FOFA：host.services.endpoints.http.headers: (key: \"server\" and value: \"Tengine\") or host.services.endpoints.http.headers: (key: \"x-cache\" and value: \"MISS\")",
			},
			{
				QuertionID:        "M088-S004",
				QuertionStatement: "修正这条 FOFA 查询，意图是搜索香港地区、Server 精确为 Apache，并排,除 ASN 4134：asn!=”4134″ \u0026\u0026 region=”HK” \u0026\u0026 server==”Apache”",
			},
			{
				QuertionID:        "M089-S010",
				QuertionStatement: "修正这条 FOFA 查询。意图：HFS 的 80 端口都要；8080 端口仅状态码为 200 时要。原查询：app=\"HFS\" \u0026\u0026 port=\"80\" |｜ app=\"HFS” \u0026\u0026 port=\"8080\" \u0026\u0026 status_code=\"200\"",
			},
			{
				QuertionID:        "M090-S008",
				QuertionStatement: "HTTP/2.0 200 OK\nContent-Length: 1171\nAccept-Ranges: bytes\nCache-Control: max-age=900\nCache-Control: public, max-age=900, stale-while-revalidate=60\nContent-Type: text/html; charset=utf-8\nDate: Mon, 14 Sep 2026 07:38:38 GMT\nEtag: \"6aa522a3-493\"\nLast-Modified: Sat, 12 Sep 2026 10:00:03 GMT\nServer: nginx\nVary: Accept-Encoding\n\n依据上面的响应片段生成搜索语句。Etag、Content-Length 保留材料中的值，其余响应头只要求字段名存在。本题不限定状态码。",
			},
			{
				QuertionID:        "M091-S009",
				QuertionStatement: "搜索开放 1000000 端口的资产。",
			},
			{
				QuertionID:        "M092-S007",
				QuertionStatement: "搜索地址或网段为 1.2.3.4.5 的资产。",
			},
			{
				QuertionID:        "M093-S004",
				QuertionStatement: "查同一条资产记录，国家代码必须等于 US，同时国家代码又必须不等于 US。",
			},
			{
				QuertionID:        "M094-S004",
				QuertionStatement: "帮我查最安全的网站。",
			},
			{
				QuertionID:        "M095-S010",
				QuertionStatement: "找出当前正在运行的容器内存使用率超过 80% 的服务器，没有容器监控数据或运行时访问权限。",
			},
			{
				QuertionID:        "M096-S001",
				QuertionStatement: "阅读这份官方文档的 TCP Agent Listener Port 一节，提取示例诊断响应中以产品名称或 Remoting 开头的全部字段名。把它们组合成FOFA查询：网页正文或服务响应原文出现任意一个字段名都要，不限定端口，不保留示例中的版本和会话值。https://www.jenkins.io/doc/book/security/services/",
			},
			{
				QuertionID:        "M097-S001",
				QuertionStatement: "帮我找和https://octodex.github.com/ 网站图标相同的其他资产。",
			},
			{
				QuertionID:        "M098-S001",
				QuertionStatement: "找和msp.gov.cn.poliec.com  一样仿冒特征的网站，总结搜索语句",
			},
			{
				QuertionID:        "M099-S001",
				QuertionStatement: "CVE-2021-41773对应的产品在FOFA上怎么搜？",
			},
			{
				QuertionID:        "M100-S001",
				QuertionStatement: "查找使用以下 TLS 证书的资产：证书颁发者通用名称为 ZeroSSL ECC DV SSL CA 2，证书详情中的 Serial Number 显示为 63:BD:83:71:62:14:83:6E:F8:00:75:35:61:61:EB:06。 转换成fofa能查的特征。",
			},
		},
	}
	jsonBytes, err := json.MarshalIndent(data, "", "")
	if err != nil {
		fmt.Printf("序列化失败：%v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
