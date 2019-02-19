定义一个描述错误的数据格式, 用来实现golang在运行时的错误传递, 目标:

- 设计数据结构的前提假设:
  - 模块的用户, 90% 只判断有没有err, 不判断err是什么.
  - 每个模块定义自己本层的错误, 但提供接口允许用户在需要时拿到最底层的错误(例如errno等).
  - 避免直接向上传递错误(上层模块拿到底层错误用处不大, 例如创建block时只得到一个permission deny的errno, 无法有效的定位到到底是哪个文件/目录出现了问题).
  - error的设计目标是为了简化查错, 提供完整的日志, 方便统计分析, 出错时的整条调用链的检查.

- 可以用于API层的错误描述, 一般API处理时收到错误并需要将错误返回给客户端, 要求错误有:

  - 具体唯一确定的error code, 便于客户端判断
  - 人类可读的error message.
  - 发生错误的相关的东西是什么.

- Message通过Code和Resource生成出来. 只提供一个Message接口用来输出message.

- 希望这个模块可以作为go的errors包的无修改替代品, 和 https://github.com/pkg/errors 的无修改替代品.

- 提供一个方便的接口让用户直接取得最底层的错误.

- 记录引发错误的底层错误是什么. 类似在存储服务中, 一个底层的IO错误导致API失败, 如果能在日志中记录引发错误的错误, 可以方便定位问题.

  有一类上层错误由几个下层错误引起, 例如多数派写中, nw = 5,3, 这时写失败3个后端会引起整个API调用失败, 这时需要记录多个下层错误.

- error 结构里可选的带有stacktrace 信息,方便打印日志(参考了 https://github.com/pkg/errors )

```go
type Error struct {
  Code       string  // error code
  Cause    []error   // 0 or seveal error that cause this error.
  Resource   string  // optional
  *stack             // optional traceback info of this error
}


// 实现系统error interface
func (e *Error) Error() string // 直接返回Code的string

// 兼容 https://github.com/pkg/errors 的接口:
func (e *Error) Cause() string // 返回第一个cause
// 其他接口没差别不列出了


// 扩展的接口
func (e *Error) AllCause() string // 返回所有cause的slice
func (e *Error) RootCause() string // 返回最初的cause
func (e *Error) AllRootCause() string // 返回所有的最底层的cause; cause的树的叶子节点.
func (e *Error) Message() string  // 给人看的, 通过Code和Resource拼装起来.

```

## 例子🌰: s2, group not found的一个可能的错误信息

以s2中场景为例, 描述下如何表示一个具体的错误,
假设一个错误是group没读取到.
而引发这个错误的是通过dbproxy读取group信息失败引发的.
而读dbproxy时重试了2次都失败了, 一次是mysql被置位readonly, 一次是socket读取超时:

```yaml
Code: GroupNotFound
Resource: "group_id://123"
Cause: 
    - Code: DBProxyReadError
      Resource: "dbproxy://127.0.0.1:3303"
      Cause:
        - Code: MysqlReadonly
          Resource: "mysql://192.168.2.2:3306"
        - Code: InvalidResponse
          Resource: "mysql://192.168.9.9:3306"
          Cause:
              - Code: SocketReadTimeout
                Resource: "tcp4://192.168.9.9:3306"
```

## 例子🌰: ec, block build error的一个可能的错误信息

```yaml
Code: BlockBuildError
Resource: "block://aabbccddeeff"
Cause:
    - Code: NeedleWriteError
      Resource: "needle://3cp/foo/bar.jpg"
      Cause:
          - Code: FSIsReadonly
            Resource: "file:///ecdrives/bbb" # schema of local file url in browser
            Cause:
                - <a native fs error> # may be an error with errno.
```
