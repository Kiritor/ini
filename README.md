## ini
Golang实现的读取ini配置文件(sections-key-value)
## 安装方法
<ceode>
	go get github.com/kiritor/ini
</ceode>
## 使用方法
INI文件格式是某些平台或软件上的配置文件的非正式标准，以节(section)和键(key)构成,如下形式:
<code>
	[database]
	username = root
	password = root
</code>

#### 初始化
<code>
	iniC:=ini.NewIni("config.ini")
</code>
ini.NewIni(filePath),根据文件路径生成ini配置对象.

#### 获取所有配置信息
<code>
	iniC.DictList()
</code>
DictList()返回的是[]Dict类型,底层是[]map[string]map[string]string数据类型(sections-key-value数组)

#### 获取单个配置信息
<code>
	iniC.GetValue("database","username")
</code>
GetValue(section,key):获取某个section(节)下key的value值

# 删除单个配置信息
<code>
	iniC.DeleteValue("database","username")
</code>
DeleteValue(section,key):删除某个section(节)下key-value,删除成功返回true,反之不存在section或者key则删除失败,返回false

# 设置或添加单个配置信息
<code>
    iniC.SetValue("database","username","root")
</code>
SetValue(section,key,value):设置某个section下某个key的value,1、如果不存在section,则直接添加section-key-value,2、如果存在section,但是不存在key则在该section下添加key-value,3、如果存在section和key则修改value.