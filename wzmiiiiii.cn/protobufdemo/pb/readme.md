# ProtoBuf

protocol buffers （ProtoBuf）是一种语言无关、平台无关、可扩展的序列化结构数据的方法，它可用于（数据）通信协议、数据存储等。

Protocol Buffers 是一种灵活，高效，自动化机制的结构数据序列化方法－可类比 XML，但是比 XML 更小（3 ~ 10倍）、更快（20 ~ 100倍）、更为简单。

json\xml都是基于文本格式，protobuf是二进制格式。

可以通过 ProtoBuf 定义数据结构，然后通过 ProtoBuf 工具生成各种语言版本的数据结构类库，用于操作 ProtoBuf 协议数据



> 编译 

对`.proto`文件进行编译,生成对应的`.go`文件
- 命令 `protoc --proto_path=IMPORT_PATH --go_out=DST_DIR $SRC_DIR/helloworld.proto`
- 命令 `protoc --go_out=. *.proto`