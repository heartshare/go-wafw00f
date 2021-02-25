# GO-WAFW00F

## 介绍

- WAFW00F是一款优秀的web应用防火墙识别开源工具：https://github.com/EnableSecurity/wafw00f
- 使用Golang重写的原因：Python环境配置不便利，Golang打包生成可执行文件直接运行
- 目前还在开发阶段，规则解析存在小问题，考虑加入协程提高执行效率

## 快速开始

```shell
./go-wafw00f.exe -u http://www.xxx.com/
```

![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/waf/waf.png)

## 简单原理

- 使用官方的py文件，但不以python执行，因为太过于麻烦，使用正则解析匹配规则
- 将解析后的规则集保存为json文件，每次执行首先检测json文件，提高效率
- 如果官方库有更新，直接替换`./waf/lib`目录即可同步更新
