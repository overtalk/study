# Python3 notes

## session1 : basic attention
- .py文件最后需要空出一行，否则PyCharm会有提示，很烦人
- 单行注释（#），多行注释（'''/"""）
- 同一个代码块的语句必须包含相同的缩进空格数，例如 if 之后需要执行多个语句，则他们需要相同的缩进
- 多行语句可以使用反斜杠(\)来实现
    - eg：total = item_one + \
    -             item_two + \
    -             item_three
- python 中数字有四种类型：整数（int）、布尔型(bool)、浮点数(float)和复数(complex)
- python 中字符串
    - 单引号也可以双引号（单行）
    - 三引号（多行）
    - 反斜杠可以用来转义，使用r可以让反斜杠不发生转义。。 如 r"this is a line with \n" 则\n会显示，并不是换行。
    - Python 中的字符串有两种索引方式，从左往右以 0 开始，从右往左以 -1 开始
- print 默认输出是换行的，如果要实现不换行需要在变量末尾加上 end=""：
    - print( x, end=" " )
- 导包
    - import some-module
    - from some-module import some-func/*
    
 ## session2 : basic data structure
- Python 中的变量不需要声明。每个变量在使用前都必须赋值，变量赋值以后该变量才会被创建。
- Python 中，变量就是变量，它没有类型，我们所说的"类型"是变量所指的内存中对象的类型。
    - counter = 100
    - name = "python rookie"
- python 允许同时为多个变量赋值
    - a = b = c = 1   
    - a, b, c = 1, 2, "python rookie"
- python3 中有六个基本的数据类型
    - 不可变
        - Number（数字）
            - int, float, bool, complex
            - 在混合计算时，Python会把整型转换成为浮点数
                - 2 / 4 （除法，得到一个浮点数 0.5）
                - 2 // 4 (除法，得到一个整数 0)
                - 17 % 3 (取余，2)
                - 2 ** 5 （乘方 32）
        - String（字符串）
            - Python 字符串不能被改变。向一个索引位置赋值，比如word[0] = 'm'会导致错误
            - 字符串的截取的语法格式如下：变量[头下标:尾下标]
            - 索引值以 0 为开始值，-1 为从末尾的开始位置
        - Tuple（元组）
            - 元组（tuple）与列表类似，不同之处在于元组的元素不能修改。元组写在小括号 () 里，元素之间用逗号隔开
                - tup1 = ()    # 空元组
                - tup2 = (20,) # 一个元素，需要在元素后添加逗号
    - 可变
        - List（列表）
            - Python 中使用最频繁的数据类型
            - 列表是写在方括号 [] 之间、用逗号分隔开的元素列表
            - 和字符串一样，列表同样可以被索引和截取，列表被截取后返回一个包含所需元素的新列表
                - 列表截取的语法格式如下：变量[头下标:尾下标]
                    - 索引值以 0 为开始值，-1 为从末尾的开始位置
                    - Python 列表截取可以接收第三个参数，参数作用是截取的步长
            - 与Python字符串不一样的是，列表中的元素是可以改变的
                - a[0] = 9
            - 列表中元素的类型可以不相同，它支持数字，字符串甚至可以包含列表（所谓嵌套）
            - List 内置了有很多方法，例如 append()、pop() 等等
        - Set（集合）
            - 集合（set）是由一个或数个形态各异的大小整体组成的，构成集合的事物或对象称作元素或是成员
            - 基本功能是进行成员关系(集合的交集、差集、并集...)测试和删除重复元素
            - 可以使用大括号 { } 或者 set() 函数创建集合，注意：创建一个空集合必须用 set() 而不是 { }，因为 { } 是用来创建一个空字典
                - parame = {value1, value2, value3}
                - set(value)
        - Dictionary（字典）
            - 字典（dictionary）是Python中另一个非常有用的内置数据类型
            - 列表(List)是有序的对象集合，字典(Dictionary)是无序的对象集合
                - 两者之间的区别在于：字典当中的元素是通过键来存取的，而不是通过偏移存取
            - 字典是一种映射类型，字典用 { } 标识，它是一个无序的 键(key) : 值(value) 的集合
            - 键(key)必须使用不可变类型, 在同一个字典中，键(key)必须是唯一的
            - 构造函数 dict() 可以直接从键值对序列中构建字典:
                - dict([('Runoob', 1), ('Google', 2), ('Taobao', 3)])
- del 语句用于删除单个或者多个对象


## session3 : details of each data structure
- 数字Number
    - int(x) : 将x转换为一个整数
    - float(x) : 将x转换到一个浮点数
    - ceil(x) : 返回数字的上入整数，如math.ceil(4.1) 返回 5
    - random() : 随机生成下一个实数，它在[0,1)范围内  ( from random import random )
    
    
- 字符串String
    - 字符串运算
        - "H" in a : 字符串a中是否含有"H"
        - "H" not in a 
    - 字符串格式化
        - print ("我叫 %s 今年 %d 岁!" % ('小明', 10))
        - %c : 格式化字符及其ASCII码
        - %s : 格式化字符串
        - %d : 格式化整数
        - %u : 格式化无符号整型
        - %o : 格式化无符号八进制数
        - %x : 格式化无符号十六进制数
        - %X : 格式化无符号十六进制数（大写）
        - %f : 格式化浮点数字，可指定小数点后的精度
        - %e : 用科学计数法格式化浮点数
        - %E : 作用同%e，用科学计数法格式化浮点数
        - %g : %f和%e的简写
        - %G : %f 和 %E 的简写
        - %p : 用十六进制数格式化变量的地址
    - 内建方法
        - capitalize() : 将字符串的第一个字符转换为大写

- 列表List
    - for x in [1, 2, 3]: print(x, end=" ")

- 元组Tuple
    - Python 的元组与列表类似，不同之处在于元组的元素不能修改
    - 创建空元组 : tup1 = ()
    - 创建只有一个元素的元组 : tup1 = (50,)

- 字典Dictionary
    - 键值是不可变的类型
    - dict = {'Alice': '2341', 'Beth': '9102', 'Cecil': '3258'}

- 集合Set
    - 可以使用大括号 { } 或者 set() 函数创建集合，注意：创建一个空集合必须用 set() 而不是 { }，因为 { } 是用来创建一个空字典

## 循环语句
- Python中的循环语句有 for 和 while
    - while 判断条件：
        - 语句
    - for <variable> in <sequence>:
        - <statements>
    - for i in range(5):
    - for i in range(5, 9):
    
## 迭代器 : 
- 迭代器对象从集合的第一个元素开始访问，直到所有的元素被访问完结束。迭代器只能往前不会后退。
- 迭代器有两个基本的方法：iter() 和 next()。
```bash
list=[1,2,3,4]
it = iter(list)    # 创建迭代器对象
for x in it:
    print (x, end=" ")
```

## 函数入参，分成两类
- 不可变类型：值传递 （Number，string。。。）
- 可变类型：指针传递 （List，。。。）
- 函数入参数
    - def printinfo( name, age = 35 ): 参数默认值
    - 加了星号 * 的参数会以元组(tuple)的形式导入，存放所有未命名的变量参数。
        ```bash
        # 可写函数说明
        def printinfo( arg1, *vartuple ):
           "打印任何传入的参数"
           print ("输出: ")
           print (arg1)
           for var in vartuple:
              print (var)
           return
        ```
    - 加了两个星号 ** 的参数会以字典的形式导入
        ```bash
        # 可写函数说明
        def printinfo( arg1, **vardict ):
           "打印任何传入的参数"
           print ("输出: ")
           print (arg1)
           print (vardict)
         
        # 调用printinfo 函数
        printinfo(1, a=2,b=3)
        ```
    - 声明函数时，参数中星号 * 可以单独出现
        - 如果单独出现星号 * 后的参数必须用关键字传入
        ```bash
        def f(a,b,*,c):
            return a+b+c
        
        f(1,2,3)   # 报错
        f(1,2,c=3) # 正常
        ```
- 匿名函数
    - python 使用 lambda 来创建匿名函数
    ```bash
    # 可写函数说明
    sum = lambda arg1, arg2: arg1 + arg2
     
    # 调用sum函数
    print ("相加后的值为 : ", sum( 10, 20 ))
    print ("相加后的值为 : ", sum( 20, 20 ))
    ```
- python 也可以返回多个数，但是返回的结果就是一个Tuple    


## 类
- 两个下划线开头，声明该属性为私有
- 两个下划线开头，声明该方法为私有方法，只能在类的内部调用 ，不能在类的外部调用。

- demo
    ```python
    class Human:
        age = 1
        name = ""
        __private = "隐藏的属性"
    
        def __init__(self, name="no name",  age=0,):
            self.age = age
            self.name = name
        def say(self):
            self.__privatefunc()
            print("my name is %s, age = %d" % (self.name, self.age))
    
        def __privatefunc(self):
            print(self.__private)
    
    
    Human("tom", 12).say()
    Human("jack", 22).say()
    ```
- 单继承    
```python
# 类定义
class people:
    # 定义基本属性
    name = ''
    age = 0
    # 定义私有属性,私有属性在类外部无法直接进行访问
    __weight = 0

    # 定义构造方法
    def __init__(self, n, a, w):
        self.name = n
        self.age = a
        self.__weight = w

    def speak(self):
        print("%s 说: 我 %d 岁。" % (self.name, self.age))


# 单继承示例
class student(people):
    grade = ''

    def __init__(self, n, a, w, g):
        # 调用父类的构函
        people.__init__(self, n, a, w)
        self.grade = g

    # 覆写父类的方法
    def speak(self):
        print("%s 说: 我 %d 岁了，我在读 %d 年级" % (self.name, self.age, self.grade))


s = student('ken', 10, 60, 3)
s.speak()
```


- 多继承
- class DerivedClassName(Base1, Base2, Base3):
- 需要注意圆括号中父类的顺序，若是父类中有相同的方法名，而在子类使用时未指定，python从左至右搜索 即方法在子类中未找到时，从左到右查找父类中是否包含方法。
```python
# 类定义
class people:
    # 定义基本属性
    name = ''
    age = 0
    # 定义私有属性,私有属性在类外部无法直接进行访问
    __weight = 0

    # 定义构造方法
    def __init__(self, n, a, w):
        self.name = n
        self.age = a
        self.__weight = w

    def speak(self):
        print("%s 说: 我 %d 岁。" % (self.name, self.age))


# 单继承示例
class student(people):
    grade = ''

    def __init__(self, n, a, w, g):
        # 调用父类的构函
        people.__init__(self, n, a, w)
        self.grade = g

    # 覆写父类的方法
    def speak(self):
        print("%s 说: 我 %d 岁了，我在读 %d 年级" % (self.name, self.age, self.grade))


# 另一个类，多重继承之前的准备
class speaker():
    topic = ''
    name = ''

    def __init__(self, n, t):
        self.name = n
        self.topic = t

    def speak(self):
        print("我叫 %s，我是一个演说家，我演讲的主题是 %s" % (self.name, self.topic))


# 多重继承
class sample(speaker, student):
    a = ''

    def __init__(self, n, a, w, g, t):
        student.__init__(self, n, a, w, g)
        speaker.__init__(self, n, t)


test = sample("Tim", 25, 80, 4, "Python")
test.speak()  # 方法名同，默认调用的是在括号中排前地父类的方法
```
