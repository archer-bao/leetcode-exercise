# Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
```
L   C   I   R
E T O E S I I G
E   D   H   N
```
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：
```
string convert(string s, int numRows);
```

示例 1:
```
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```

示例 2:
```
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```


## 解答
这道题通过找规律可以发现 Z 字形是有周期的，也就是每 2 * numRows - 2 的周期变换。

为什么是 2 * numRows - 2？

因为除了第一行和最后一行只出现1个字符之外，其余行都是出现2个字符。所以是 numRows + numRows - 2

知道周期之后就好办了，总体以第一行作为参考，逐行扫描，第一行的字符索引一定是周期，比如上面的索引0、6、12。

然后第二行的E和O是第一个周期内的字符，其偏移位置为 周期0 + 行数（第二行为1）和  周期1 - 行数（第二行为1）

然后第三行的E和C是第一个周期内的字符，其偏移位置为 周期0 + 行数（第三行为2）和  周期1 - 行数（第三行为2）

最后一行和第一行类似，为周期数+行数
