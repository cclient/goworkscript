go的一些脚本工具

个人实际把go当一个简版的c,轻量的java,无回调地狱的nodejs,性能更高的python来用

这个项目不是什么大工程，都很轻，更像是类似python的业余脚本

####src/client/pdf

下载极邦客网站pdf

索引来源(无意中搜到一个，发现可以搜到好多，就写这个脚本都拿下来)

http://cn.bing.com/search?q=site:ppt.geekbang.org+ppt

google和bing搜索的结果基本一致

建表语句在
src/client/pdf/db.sql


用网站源pdf id 作排重

如下例

'1','Bing','580984e1c891a','如何打造大规模互联网企业的 监控告警平台','http://ppt.geekbang.org/slide/download/467/580984e1c891a.pdf','如何打造大规模互联网企业的 监控告警平台-- 以携程hickwall为例 author rhtang@ctrip.com'

网站id 为580984e1c891a


google,!@#$%&,不是每个人都,!@#$%&,为用适用更多人(也为省点代理流量),用bing作搜索引擎

最大的瓶颈在下载文件的网络io,并行无甚意义，所以解析和下载完全串行化

能被搜索引擎爬到,说明极邦客允许无登录访问。不知网站是有意为知,还是权限部分设计缺陷所致


爬虫拿到的部分pdf文件信息示例


##src/client/spider 

后台并行网页抓取(只保留了基本的get请求,需要加代理,加header头,多节点分发的可以自已补充)

通过redis解耦

客户端提交url信息至redis

后台批量并行请求url(默认并行请求数20条)

流程说明

1共提交1000条至redis(往list前入)

2每隔10秒从redis取100条(从list后取)

2因并行数是20，100条共分为5组,每组之间FIFO

3取完后移100条出list(从list后移出)

4完成后继续从redis取100条(从list前取)

go http 底层会复用tcp连接,请求效率很高

并发比较粗糙，实际可以封装的通用一些。


##src/client/wangpiao

网票网抓取内容解析，网站改版，部分已失效

启动方式

修改src/app.go

执行 go run src/app.go

也可以打docker启用

*————————————————————————————*


实际以上具备了爬虫及解析最基本的功能。

https://github.com/cclient/gowebframework +  + src/client/spider 

可以组合成简版的爬虫服务后台，提交需要请求url，然后再从redis直接取网站内容解析。