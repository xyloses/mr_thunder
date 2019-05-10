## 软件说明
	这个只是配置文件，我们众所周知的排版语言TeX，但是貌似被html取代了. 是的QML很多节目层的排版标准都已经是html. 我非常大胆的使用html做数据层的描述语言。 对于文章这种数据结构，xml确实是一个非常好的表述。
	thunder 类似于hexo这样的静态网站生成器，但是thunder更适合中国，除了可以生成网站和app，还可以生成公众号，等内容平台文案，以及小程序。


## Markdown语法

标题
这是最为常用的格式，在平时常用的的文本编辑器中大多是这样实现的：输入文本、选中文本、设置标题格式。
而在 Markdown 中，你只需要在文本前面加上 # 即可，同理、你还可以增加二级标题、三级标题、四级标题、五级标题和六级标题，总共六级，只需要增加 #  即可，标题字号相应降低。例如：

# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

```
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题
```

## 列表

列表格式也很常用，在 Markdown 中，你只需要在文字前面加上 - 就可以了，例如：

- 文本1
- 文本2
- 文本3
如果你希望有序列表，
也可以在文字前面加上 1. 2. 3. 就可以了，例如：

1. 文本1
2. 文本2
3. 文本3

## 链接
在 Markdown 中，插入链接不需要其他按钮，你只需要使用 [显示文本](链接地址) 这样的语法即可，例如：
[hweijie](hweijie.com)

```
[hweijie](hweijie.com)
```

## 图片
插入图片的语法和链接的语法很像，只是前面多了一个 !
![hweijie](hweijie.com/logo)
```
![hweijie](hweijie.com/logo)
```

## 引用

在我们写作的时候经常需要引用他人的文字，这个时候引用这个格式就很有必要了，在 Markdown 中，你只需要在你希望引用的文字前面加上 > 就好了，例如：
> bababab
```
> bababab
```

## 粗体和斜体

Markdown 的粗体和斜体也非常简单，用两个 * 包含一段文本就是粗体的语法，用一个 * 包含一段文本就是斜体的语法。例如：

*何嘉源的博客*
**始于颜值、陷于才华、忠于人品**

```
*何嘉源的博客*
**始于颜值、陷于才华、忠于人品**
```


## 代办清单To-do List

你可以在Markdown中编写代办清单
相关代码

- [x] 已完成项目1
  - [x] 已完成事项
  - [ ] 代办事项
- [ ] 代办项目2
- [ ] 代办项目3

```
- [x] 已完成项目1
  - [x] 已完成事项
  - [ ] 代办事项
- [ ] 代办项目2
- [ ] 代办项目3
```

## 表格
使用 | 来分隔不同的单元格，使用 - 来分隔表头和其他行：
为了美观，可以使用空格对齐不同行的单元格，并在左右两侧都使用 | 来标记单元格边界：
在表头下方的分隔线标记中加入 :，即可标记下方单元格内容的对齐方式：
:--- 代表左对齐
:--: 代表居中对齐
---: 代表右对齐


| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

```
| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |
```


## 高效绘制 流程图、序列图、甘特图、流程图
```

graph TD
    A[Christmas] -->B(Go Shopping)
    B --> C{Let me think}
    C -->|one| D[Laptop]
    C -->|two| E[iPhone]
    C -->|three| F[Car]
    
```
graph TD
    A[Christmas] -->B(Go Shopping)
    B --> C{Let me think}
    C -->|one| D[Laptop]
    C -->|two| E[iPhone]
    C -->|three| F[Car] 
    
``` 
sequenceDiagram
    loop every day
        Alice->John: Hello John, how are you?
        John->Alice: Great!
    end
```
sequenceDiagram
    loop every day
        Alice->John: Hello John, how are you?
        John->Alice: Great!
    end
    
### 甘特图    
```
gantt
dateFormat YYYY-MM-DD
title 产品计划表
section 初期阶段
明确需求: 2016-03-01, 10d
section 中期阶段
跟进开发: 2016-03-11, 15d
section 后期阶段
走查测试: 2016-03-20, 9d
```

gantt
dateFormat YYYY-MM-DD
title 产品计划表
section 初期阶段
明确需求: 2016-03-01, 10d
section 中期阶段
跟进开发: 2016-03-11, 15d
section 后期阶段
走查测试: 2016-03-20, 9d


```
    gantt
    dateFormat YYYY-MM-DD
    title 产品计划表
    section 初期阶段
    明确需求: 2016-03-01, 10d
    section 中期阶段
    跟进开发: 2016-03-11, 15d
    section 后期阶段
    走查测试: 2016-03-20, 9d
```

## 书写数学公式


inline math: `$\dfrac{
\tfrac{1}{2}[1-(\tfrac{1}{2})^n] }{
1-\tfrac{1}{2} } = s_n$`.

```
inline math: `$\dfrac{
\tfrac{1}{2}[1-(\tfrac{1}{2})^n] }{
1-\tfrac{1}{2} } = s_n$`.
```



以后在更新咯！