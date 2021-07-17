# Po-Struct

使用 Go 实现 Python 中的基础数据结构及其特性，包括但不限于：

- [ ] float
- [ ] int
- [ ] bytes
- [ ] str
- [ ] list
- [ ] dict

# 结构特性

## list

参考 Python List 设计，实现了自动扩缩容

- [x] `Srting()`序列化
- [x] `Size()`获取长度
- [x] `Cap()`获取容量
- [x] `GetItem()`获取元素，含反向索引 `l.GetItem(-1)`
- [x] `SetItem()`设置元素，含反向索引
- [x] `Append()`尾部添加，可添加多个元素
- [x] `Insert()`中间插入，含反向索引
- [x] `GetSlice()`中间切片，含反向索引
- [x] `Extend()`拼接
- [x] `Reverse()`倒转
- [x] `Pop()`尾部出栈，删除指定位置元素
- [x] `Clear()`清空
- [x] `Range()` 接受一个方法，遍历
- [ ] `Sort()`排序

# set

- [x] `Srting()`序列化
- [x] `Size()`获取长度
- [x] `Copy()`拷贝
- [x] `IsSubSet()`判断子集
- [x] `IsSuperSet()`判断超集
- [x] `Equal()`判断相等
- [x] `Union()`返回并集
- [x] `Difference()`返回差集
- [x] `Intersection()`返回交集
- [x] `Range()` 接受一个方法，遍历


## float

- [ ] 对象缓冲池

## int

- [ ] 小数静态对象池
- [ ] 不溢出的大数运算

