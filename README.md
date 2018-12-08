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

# 指令集和解释器
每一个类或者接口都会被 java 编译器编译成一个 class 文件, 类或接口的方法信息就放在 class 文件的 method_info 结构中.
如果方法不是抽象的, 也不是本地方法, 方法的 java 代码就会被编译器编译成字节码
(即使方法是空的, 编译器也会生成一条 return 语句) 存放在 method_info 结构的 Code属性中

字节码中存放编码后的 java 虚拟机指令. 每条指令都以一个单自己的操作码(opcode) 开头, 由一字节表示
java 虚拟机最多只能支持 256(2 ^ 8)条指令
如果指令想象成函数的话, 操作数就是它的参数, 为了让字节码更加紧凑
很多操作码本身就隐含了操作数
比如把常数 0 推入操作数栈的指令是 iconst_0

java 虚拟机规范把已经定义的 205 条指令按用途分成了 11 类
分别是
- 常量(constant)
- 加载(loads)
- 存储(store)
- 操作数栈(stack)
- 数学(math)
- 转换(conversions)
- 比较(comparisons)
- 控制(control)
- 引用(references)
- 扩展(extended)
- 保留(reserved)


# 类和对象
todo

# 方法调用和返回
## 方法的分类
* 静态方法
* 实例方法

静态方法在编译时就绑定了
动态方法可能要推迟到运行期才能知道

java 7 之前虚拟机规范一共提供了 4 条方法调用指令

* invokestatic
    用于调用静态方法
* invokespecial
	用来调用无需动态绑定的实例方法(构造函数, 私有方法)
* super
	调用超类(父类)方法
* 动态绑定
	* invokeinterface
		接口类型的引用调用方法
	* invokevirtual
		剩余情况

invokedynamic 是 java7 新增的一条指令, 用于更好的支持动态语言, 详情查看 vm 规范

方法方回执的指令
* return 用于没有返回值的情况
* areturn	返回引用
* ireturn  返回 int
* lreturn 返回 long
* freturn	返回 float
* dreturn 返回 double

A 方法调用 B 方法, B 方法结束后从栈帧的操作数栈取数返回值推入 A 方法帧的操作数栈顶

6条返回指令和4条调用指令加上字节码解析栈帧之间传递变量, 实现了 java 的方法调用




