use go write jvm from book of jvm

# classpath
通过 classpath 搜索 class 文件, 并解析为 byte[]

# classfile
通过 jvm 规范解析 byte[] 为 classfile 文件
- 通过 vm 规范解析版本魔数等
- 解析 class 常量池
- 解析属性表

# 线程栈帧(线程栈帧)
- 线程共享数据区
    需要在 vm 启动时创建好, 在 vm 退出时销毁
    共享区主要存放类数据和对象
    对象数据存放在堆(Heap)中
    类数据存放在方法区(Method Area)
    堆由垃圾收集器定期清理
    类数据包括字段和方法信息, 方法的字节码, 运行时常量池
- 线程私有数据区
    在创建线程时创建, 销毁线程时销毁, 生命周期和线程绑定
    线程私有的运行时数据区用于辅助执行 java 字节码
    每个线程都有自己的 pc 寄存器(Program Counter) 和 java 虚拟机栈(JVM Stack)
    Stack Frame(栈帧) 保存方法的执行状态
    包括局部变量表(Local Variable) 和操作数栈(Operand Stack)
    线程执行某个方法时, 该方法叫做当前类
    如果当前方法是 java 方法则 pc 寄存器中存放当前正在执行的 java 虚拟机指令的地址
    否则当前方法是本地方法, pc 寄存器中的值没有明确定义

