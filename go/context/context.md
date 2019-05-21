# Context

- context的主要数据结构是一种嵌套的结构或者说是单向的继承关系的结构，比如最初的context是一个小盒子，里面装了一些数据，之后从这个context继承下来的children就像在原本的context中又套上了一个盒子，然后里面装着一些自己的数据。或者说context是一种分层的结构，根据使用场景的不同，每一层context都具备有一些不同的特性，这种层级式的组织也使得context易于扩展，职责清晰。

## context 包的核心是 struct Context，声明如下：
type Context interface {
    <p>&emsp;  Deadline() (deadline time.Time, ok bool)</p>
    <p>&emsp;  Done() <-chan struct{}</p>
    <p>&emsp;  Err() error</p>
    <p>&emsp;  Value(key interface{}) interface{}</p>
}

- 复制代码可以看到Context是一个interface，在golang里面，interface是一个使用非常广泛的结构，它可以接纳任何类型。Context定义很简单，一共4个方法，我们需要能够很好的理解这几个方法

- Deadline方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

- Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。之后，Err 方法会返回一个错误，告知为什么 Context 被取消。

- Err方法返回取消的错误原因，因为什么Context被取消。

- Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

## Context 的继承
- 有了如上的根Context，那么是如何衍生更多的子Context的呢？这就要靠context包为我们提供的With系列的函数了。

    - func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

    - func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

    - func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

    - func WithValue(parent Context, key, val interface{}) Context

- 通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。
- WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。
- WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
- WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
- WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，这是我们实际用经常要用到的技巧，一般我们想要通过上下文来传递数据时，可以通过这个方法，如我们需要tarce追踪系统调用栈的时候。
