##### 终端录屏程序
- 这是一个**小巧的终端录屏程序**
- `gors is the meaning of Go Record Screen`

##### 编译二进制文件
`go build -o gors main.go`

##### 录屏
- **指定文件的方式** `gors record --filename panqd.json`
- **不指定文件的方式** `gors record` **录屏文件生成在/tmp目录下, 以gors文件名开头**

##### 回放
`gors play --filename panqd.json`

##### 帮助
- `gors --help`
- `gors record --help`
- `gors play --help`
