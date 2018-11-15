use go write jvm from book of jvm

# classpath
通过 classpath 搜索 class 文件, 并解析为 byte[]

# classfile
通过 jvm 规范解析 byte[] 为 classfile 文件
- 通过 vm 规范解析版本魔数等
- 解析 class 常量池
- 解析属性表
