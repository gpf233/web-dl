# web-dl

#### usage:

```sh
./web-dl [-h] [-p] [-t] <-i hlsUrl.json> [-o outputDir]
```

#### 查看帮助：

```
./web-dl -h
```

#### hlsUrl获取方法：

1. 进入课程录播页面，按F12打开开发者工具，选中Network标签栏
2. 按F5刷新页面，在filter过滤器中输入`getViewUrlHls`

3. 右键点击查询得到的请求，Copy - Copy response
4. 使用[在线工具](https://www.urldecoder.io/)解码URL，将URL中的info键对应的json值复制到input.json中

