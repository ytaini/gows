[TOC]

# context 包

https://segmentfault.com/a/1190000040917752


```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
- Deadlne方法：当Context自动取消或者到了取消时间被取消后返回
- Done方法：当Context被取消或者到了deadline返回一个被关闭的channel
- Err方法：当Context被取消或者关闭后，返回context取消的原因
- Value方法：获取设置的key对应的值

这个接口主要被3个类继承实现，分别是`emptyCtx`、`ValueCtx`、`cancelCtx`采用匿名接口的写法，这样可以对任意实现了该接口的类型进行重写。


## 创建根Context

在我们调用`context.Background、context.TODO`时创建的对象就是`empty`：
```go
var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)

func Background() Context {
    return background
}

func TODO() Context {
    return todo
}
```

```
Background和TODO还是一模一样的,官方说：
    background它通常由主函数、初始化和测试使用，并作为传入请求的顶级上下文；
    TODO是当不清楚要使用哪个 Context 或尚不可用时，代码应使用 context.TODO，后续在在进行替换掉，归根结底就是语义不同而已。
```

### emptyCtx类
emptyCtx主要是给我们创建根Context时使用的，其实现方法也是一个空结构，实际源代码长这样：
```go
type emptyCtx int
func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key any) any {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}
```


## WithValue的实现

`withValue`内部主要就是调用`valueCtx`类：
```go
func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}
```


### valueCtx类
valueCtx目的就是为Context携带键值对，因为它采用匿名接口的继承实现方式，他会继承父Context，也就相当于嵌入Context当中了

```go
type valueCtx struct {
	Context
	key, val any
}
```

实现了String方法输出Context和携带的键值对信息：

```go
func (c *valueCtx) String() string {
    return contextName(c.Context) + ".WithValue(type " +
        reflectlite.TypeOf(c.key).String() +
        ", val " + stringify(c.val) + ")"
}
```

实现Value方法来存储键值对：
```go
func (c *valueCtx) Value(key any) any {
    if c.key == key {
        return c.val
    }
    return value(c.Context, key)
}

func value(c Context, key any) any {
    for {
        switch ctx := c.(type) {
        case *valueCtx:
            if key == ctx.key {
                return ctx.val
            }
            c = ctx.Context
        case *cancelCtx:
            if key == &cancelCtxKey {
                return c
            }
            c = ctx.Context
        case *timerCtx:
            if key == &cancelCtxKey {
                return &ctx.cancelCtx
            }
            c = ctx.Context
        case *emptyCtx:
            return nil
        default:
            return c.Value(key)
        }
    }
}
```

我们在调用`Context`中的`Value`方法时`会层层向上调用直到最终的根节点`，中间要是找到了key就会返回，否会就会找到最终的emptyCtx返回nil。


## WithCancel的实现

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    }
    c := newCancelCtx(parent)
    propagateCancel(parent, &c)
    return &c, func() { c.cancel(true, Canceled) }
}
```

这个函数执行步骤如下：
- 创建一个`cancelCtx`对象，作为`子context`
- 然后调用`propagateCancel构建父子context之间的关联关系`，这样当父context被取消时，子context也会被取消。
- 返回`子context对象和子树取消函数`


### cancelCtx类
`cancelCtx`继承了`Context`，也实现了接口`canceler`:
```go
type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}

type cancelCtx struct {
    Context

    mu       sync.Mutex            // protects following fields
    done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
    children map[canceler]struct{} // set to nil by the first cancel call
    err      error                 // set to non-nil by the first cancel call
}
```

字段解释：
- mu：就是一个互斥锁，保证并发安全的，所以context是并发安全的
- done：用来做context的取消通知信号，之前的版本使用的是chan struct{}类型，现在用atomic.Value做锁优化
- children：key是接口类型canceler，目的就是存储实现当前canceler接口的子节点，当根节点发生取消时，遍历子节点发送取消信号
- error：当context取消时存储取消信息

这里实现了`Done`方法，返回的是一个只读的`channel`，目的`就是我们在外部可以通过这个阻塞的channel等待通知信号。`

<br>

`propagateCancel`方法
```go
func propagateCancel(parent Context, child canceler) {
  // 如果返回nil，说明当前父`context`从来不会被取消，是一个空节点，直接返回即可。
    done := parent.Done()
    if done == nil {
        return // parent is never canceled
    }

  // 提前判断一个父context是否被取消，如果取消了也不需要构建关联了，
  // 把当前子节点取消掉并返回
    select {
    case <-done:
        // parent is already canceled
        child.cancel(false, parent.Err())
        return
    default:
    }

  // 这里目的就是找到可以“挂”、“取消”的context
    if p, ok := parentCancelCtx(parent); ok {
        p.mu.Lock()
    // 找到了可以“挂”、“取消”的context，但是已经被取消了，那么这个子节点也不需要
    // 继续挂靠了，取消即可
        if p.err != nil {
            child.cancel(false, p.err)
        } else {
      // 将当前节点挂到父节点的childrn map中，外面调用cancel时可以层层取消
            if p.children == nil {
        // 这里因为childer节点也会变成父节点，所以需要初始化map结构
                p.children = make(map[canceler]struct{})
            }
            p.children[child] = struct{}{}
        }
        p.mu.Unlock()
    } else {
    // 没有找到可“挂”，“取消”的父节点挂载，那么就开一个goroutine
        atomic.AddInt32(&goroutines, +1)
        go func() {
            select {
            case <-parent.Done():
                child.cancel(false, parent.Err())
            case <-child.Done():
            }
        }()
    }
}
```

<br>

`cancel方法`
```go
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
  // 取消时传入的error信息不能为nil, context定义了默认error:var Canceled = errors.New("context canceled")
    if err == nil {
        panic("context: internal error: missing cancel error")
    }
  // 已经有错误信息了，说明当前节点已经被取消过了
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return // already canceled
    }
  
    c.err = err
  // 用来关闭channel，通知其他协程
    d, _ := c.done.Load().(chan struct{})
    if d == nil {
        c.done.Store(closedchan)
    } else {
        close(d)
    }
  // 当前节点向下取消，遍历它的所有子节点，然后取消
    for child := range c.children {
        // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err)
    }
  // 节点置空
    c.children = nil
    c.mu.Unlock()
  // 把当前节点从父节点中移除，只有在外部父节点调用时才会传true
  // 其他都是传false，内部调用都会因为c.children = nil被剔除出去
    if removeFromParent {
        removeChild(c.Context, c)
    }
}
```


## withDeadline、WithTimeout的实现

先看`WithTimeout`方法，它内部就是调用的`WithDeadline`方法：
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}
```
所以我们重点来看`withDeadline`是如何实现的：

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
  // 不能为空`context`创建衍生context
    if parent == nil {
        panic("cannot create context from nil parent")
    }
  
  // 当父context的结束时间早于要设置的时间，则不需要再去单独处理子节点的定时器了
    if cur, ok := parent.Deadline(); ok && cur.Before(d) {
        // The current deadline is already sooner than the new one.
        return WithCancel(parent)
    }
  // 创建一个timerCtx对象
    c := &timerCtx{
        cancelCtx: newCancelCtx(parent),
        deadline:  d,
    }
  // 将当前节点挂到父节点上
    propagateCancel(parent, c)
  
  // 获取过期时间
    dur := time.Until(d)
  // 当前时间已经过期了则直接取消
    if dur <= 0 {
        c.cancel(true, DeadlineExceeded) // deadline has already passed
        return c, func() { c.cancel(false, Canceled) }
    }
    c.mu.Lock()
    defer c.mu.Unlock()
  // 如果没被取消，则直接添加一个定时器，定时去取消
    if c.err == nil {
        c.timer = time.AfterFunc(dur, func() {
            c.cancel(true, DeadlineExceeded)
        })
    }
    return c, func() { c.cancel(true, Canceled) }
}
```

`withDeadline`相较于`withCancel`方法`也就多了一个定时器去定时调用cancel方法`，这个`cancel`方法在`timerCtx`类中进行了重写，我们先来看一下`timerCtx`类，他是基于`cancelCtx`的，多了两个字段：


### timerCtx类
```go
type timerCtx struct {
    cancelCtx
    timer *time.Timer // Under cancelCtx.mu.

    deadline time.Time
}
```

timerCtx实现的cancel方法，内部也是调用了cancelCtx的cancel方法取消：
```go
func (c *timerCtx) cancel(removeFromParent bool, err error) {
  // 调用cancelCtx的cancel方法取消掉子节点context
    c.cancelCtx.cancel(false, err)
  // 从父context移除放到了这里来做
    if removeFromParent {
        // Remove this timerCtx from its parent cancelCtx's children.
        removeChild(c.cancelCtx.Context, c)
    }
  // 停掉定时器，释放资源
    c.mu.Lock()
    if c.timer != nil {
        c.timer.Stop()
        c.timer = nil
    }
    c.mu.Unlock()
}
```
## context使用原则
- context.Background 只应用在最高等级，作为所有派生 context 的根。
- context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
- 不要把Context放在结构体中，要以参数的方式传递。
- 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
- 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
- Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递。context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
- context是线程安全的，可以放心的在多个goroutine中传递。同一个Context可以传给使用其的多个goroutine，且Context可被多个goroutine同时安全访问。
- Context 结构没有取消方法，因为只有派生 context 的函数才应该取消 context。


Go 语言中的 context.Context 的主要作用还是`在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用`，虽然它也有传值的功能，但是这个功能我们还是很少用到。在真正使用传值的功能时我们也应该非常谨慎，使用 `context.Context` 进行传递参数请求的所有参数一种非常差的设计，比较常见的使用场景是`传递请求对应用户的认证令牌以及用于进行分布式追踪的请求 ID`。
## context优缺点


