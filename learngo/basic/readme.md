## 基础语法

### 变量定义

### 内建变量类型

- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune(类似于其他语言的 char，但是 rune 是 4 字节，char 是 1 字节)
- float32, float64, complex64, comple128


### 强制类型转换

- 类型转换是强制的, 不能隐式转换
- `var a,b int = 3,4`
- `var c int = math.Sqrt(a*a + b*b)` ❌
- `var c int = int(math.Sqrt(float64(a*a + b*b)))` ✅

### 常量的定义

- `const filename = "abc.txt"`
- `const` 数值可作为各种类型使用
- `const a,b = 3,4`
- `var c int = int(math.Sqrt(a*a + b*b))`

#### 枚举类型

- iota

#### 变量定义要点回顾

- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有 char, 只有 rune 32位