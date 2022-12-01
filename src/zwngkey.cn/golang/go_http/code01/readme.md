https://geektutu.com/post/gee.html

net/http提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。

一些Web开发中简单的需求并不支持，需要手工实现。
- 动态路由：例如hello/:name，hello/*这类的规则。
- 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
- 模板：没有统一简化的HTML机制。
- …

Web 框架的需要的一些核心能力:
- 路由(Routing)：将请求映射到函数，支持动态路由。例如'/hello/:name。
- 模板(Templates)：使用内置模板引擎提供模板渲染机制。
- 工具集(Utilites)：提供对 cookies，headers 等处理机制。
- 插件(Plugin)：Bottle本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。
- …

