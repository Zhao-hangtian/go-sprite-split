# 雪碧图切割工具
使用Go实现的一个简单的PNG图片素材切割工具，基于在Alpha（透明度通道）使用DFS（深度优先搜索）算法做非透明连续区域的切割，得到素材包的雪碧图位置信息。

## 作用
1. 实现素材包雪碧图位置区域的提取（到JSON格式文件），可以精准定位到所需素材。为什么需要雪碧图可以参考[这里](https://www.w3schools.com/css/css_image_sprites.asp#:~:text=An%20image%20sprite%20is%20a,server%20requests%20and%20save%20bandwidth.)或者[这里](https://m.imooc.com/wiki/csssprite-whysprite#:~:text=%E5%88%A9%E7%94%A8%E9%9B%AA%E7%A2%A7%E5%9B%BE%E8%83%BD%E5%A4%9F%E5%BE%88,%E7%9A%84%E5%A4%A7%E5%B0%8F%E8%BF%98%E8%A6%81%E5%B0%8F%E3%80%82)。

2. 实现雪碧图的分割，即从一个大素材图里提取所有的元素。

## 扩展
1. 如果您有其他希望的输出格式，可以自由的修改此代码~

2. 如果您的素材不是带透明通道的PNG图片，可以使用GIMP，Photoshop等工具先进行见到那的编辑（增加透明图层、魔棒/色彩选择去除底部背景）。

## 编译
`go build main.go`

## 例子
原始素材：
![图标集](example/icons.png "Demo Icons Picture")

### 1. 提取雪碧图信息
- 运行
`./main -input example/icons.png -output output  -mode sprite`
- 输出
![img.png](assets/img.png)
  您将在目录`output`下得到分割得到JSON文件，里面包含了雪碧图信息：
![img.png](assets/img_2.png)
坐标以左上方为零点，h是高度，w是宽度，x1，y1是左上坐标，x2，y2是右下坐标。

格式化的素材位置JSON信息可以在您的游戏、应用开发过程中方便的使用。

### 2. 分割大图片
- 运行
`./main --input example/icons.png`
- 输出
![img_1.png](assets/img_1.png)

您将在目录`output`下得到分割得到小图片
![截图1](assets/Screenshot-1.png "输出截图")

