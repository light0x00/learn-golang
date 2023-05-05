package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/tour/tree"
)

func TestForLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func swap(i, j int) {
	if t := i; t > 0 {

	}
}

func Sqrt(x float64) float64 {
	var r float64 = 1 << 64
	for z := x; ; {
		if math.Abs(x-z*z) < math.Abs(x-r*r) {
			r = z
		} else {
			return r
		}
		z -= (z*z - x) / (2 * z)
	}
}

func TestSqrt(t *testing.T) {
	fmt.Println(math.Sqrt(2))
	// fmt.Println(Sqrt(2))
}

/*
===================== Pointer =====================
*/

func TestPointer(t *testing.T) {
	//声明一个int指针
	var p *int
	var i = 1
	//&i 生成一个i的指针,赋给p
	p = &i
	p2 := &i
	fmt.Println(p, p2)

	i = 2
	//直接访问 i, 通过指针p访问i (前者是立即寻址,数据随指令到达寄存器; 后者是直接寻址,数据的地址随指令进入寄存器,cpu再通过地址访问数据)
	fmt.Println(i, *p)

	//使用指针更改 i
	*p = 3
	fmt.Println(i, *p)
}

/*
===================== Struct =====================
*/

func Test_Struct(t *testing.T) {
	var v = Vertex{1, 2}
	withStructParam(v)
	withStructPointerParam(&v)
	vp := &v
	fmt.Println(&vp)
}

type Vertex struct {
	X int
	Y int
}

/*
值传递,传递的是拷贝副本
*/
func withStructParam(v Vertex) {
	v.X = 3
}

/*
 */
func withStructPointerParam(v *Vertex) {
	//通过指针更新结构变量
	(*v).Y = 11
	//简写形式
	v.Y = 11
}

func TestStruct2(t *testing.T) {
	//内联结构
	a := struct {
		i int
		b bool
	}{
		i: 1,
		b: true,
	}
	fmt.Println(a)
}

/*
===================== Slice =====================
*/

func TestSlice(t *testing.T) {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	var names2 = names[2:3]

	names2[0] = "Light"

	fmt.Println(names[2]) //slice 是编译器层面实现的虚拟数组,其并没有真的分配内存给新数组,而是共享原数组

	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Println(len(names2))
	fmt.Println(cap(names2))

	names3 := names2[:1]
	fmt.Println(names3)
}

func TestArray(t *testing.T) {
	a1 := [5]int{}
	a2 := make([]int, 5)
	fmt.Println(a1, a2)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func TestSliceLiteral(t *testing.T) {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	//结构数组
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func TestSliceResize(t *testing.T) {
	var s []int
	fmt.Println(&s)
	s = append(s, 2, 3, 4)
	fmt.Println(s)
	fmt.Println(&s)

}

func TestForRange(t *testing.T) {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("index: %d ,value: %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("value: %d\n", v)
	}

	for i := range pow {
		fmt.Printf("index: %d\n", i)
	}

}

func create2DimArray1(x, y int) [][]uint8 {
	r := [][]uint8{}
	for i := 0; i < x; i++ {
		r = append(r, make([]uint8, y))
	}
	return r
}
func create2DimArray2(x, y int) [][]uint8 {
	r := make([][]uint8, x)
	for x := range r {
		r[x] = make([]uint8, y)
	}
	return r
}

func create2DimArray3(x, y int) [][]uint8 {
	arr := make([]uint8, x*y)
	var r [][]uint8
	for i := 0; i < len(arr); i += y {
		r = append(r, arr[i:i+y])
	}
	return r
}

func TestPic2(t *testing.T) {
	create2DimArray3(10, 10)
	// for i := 0; i < min(dx, dy); i++ {
	// 	r[i][i] = 2
	// }
	// return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
===================== Map =====================
*/

func TestMap(t *testing.T) {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])
}

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	for _, w := range strings.Fields(s) {
		val, exist := m[w]
		if !exist {
			m[w] = 1
		} else {
			m[w] = val + 1
		}
	}
	return m
}

func TestMapExcersice(t *testing.T) {
	fmt.Println(WordCount("to be or not to be"))
}

/*
===================== Funciton =====================
*/

func asynRead(fn func(a ...any) (n int, err error)) {
	fn("Hello World")
}

func TestFunction(t *testing.T) {
	asynRead(fmt.Println)
	asynRead(func(a ...any) (n int, err error) {
		return fmt.Println(a...)
	})
}

func fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		c := a + b
		a = b
		b = c
		return a
	}

}

func TestFibonacci(t *testing.T) {

	gen := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}

/*
===================== Methods =====================
*/
// [示例1]
// 定义一个结构类型
type Fibonacci struct {
	a int
	b int
}

func NewFibonacci() Fibonacci {
	f := Fibonacci{0, 1}
	fmt.Printf("%p\n", &f)
	return f
}

func NewFibonacci2() *Fibonacci {
	f := Fibonacci{0, 1}
	fmt.Printf("%p\n", &f)
	return &f
}

func TestStructMethod(t *testing.T) {
	f := NewFibonacci()
	fmt.Printf("%p\n", &f)

	for i := 0; i < 5; i++ {
		fmt.Println(f.gen())
	}

	f2 := NewFibonacci2()
	fmt.Printf("%p\n", f2)
	for i := 0; i < 5; i++ {
		fmt.Println((*f2).gen())
	}
}

// 为类型定义一个方法
func (f *Fibonacci) gen() int {
	c := f.a + f.b
	f.a = f.b
	f.b = c
	return f.a
}

// [示例2]
// 1.定义一个类型
type MyInt int

func (i *MyInt) Abs() int {
	if *i < 0 {
		return int(-*i)
	} else {
		return int(*i)
	}
}

// 2.给类型加方法
func (i MyInt) Abs2() int {
	if i < 0 {
		return int(-i)
	} else {
		return int(i)
	}
}

func TestTypeMethod(t *testing.T) {
	i := MyInt(-2)
	fmt.Println(i.Abs())
	fmt.Println((&i).Abs())
}

//[总结] 本质上,结构、类型方法, 是编译器的语法糖, 等价于方法接受参数

/*
===================== Interfaces =====================
*/

// [示例1]
// 1.定义接口
type AST interface {
	eval()
}

// 2-1 定义结构体,并实现接口方法
type UnaryExpr struct {
	operator string
	epxr     AST
}

func (u *UnaryExpr) eval() {
	u.epxr.eval()
	fmt.Printf("eval unary expr %v\n", u.operator)
}

// 2-2 定义类型,并实现接口方法
type Lexeme string

func (l *Lexeme) eval() {
	fmt.Printf("eval lexeme:%v\n", l)
}

func TestInterfaces(t *testing.T) {
	var ast AST

	l := Lexeme("123")
	u := UnaryExpr{"+", &l}
	u.eval()

	ast = &u
	ast = &l
	ast = &l
	ast = &u

	recognizeType(&l)
	recognizeType(&u)
	fmt.Printf("%v,%T\n", ast, ast)
}

//[总结]
// 实现方法的receiver决定了“实现类”是谁
// func (l *Lexeme) eval() 实现类是指针类型 *Lexeme
// 而 func (l Lexeme) eval() 实现类是 Lexeme

//==> Type assertions 判断接口实现的类型

func recognizeType(ast AST) (string, error) {
	val, isLexeme := ast.(*Lexeme)
	if isLexeme {
		return fmt.Sprintf("it's a lexeme,%v", *val), nil
	}
	expr, isUnaryExpr := ast.(*UnaryExpr)
	if isUnaryExpr {
		return fmt.Sprintf("it's a UnaryExpr,%v", (*expr).operator), nil
	}
	return "", fmt.Errorf("Unknown ast type %T", ast)
}

func TestTypeAssertions(t *testing.T) {
	l := Lexeme("Hello")
	u := UnaryExpr{operator: "-", epxr: &l}

	fmt.Println(recognizeType(&l))
	fmt.Println(recognizeType(&u))
}

//==>Type switches

func recognizeType2(o interface{}) (string, error) {
	switch v := o.(type) {
	case int:
		return fmt.Sprintf("int,%v", v), nil
	case string:
		return fmt.Sprintf("string,%v", v), nil
	case Lexeme:
		v.eval()
		return fmt.Sprintf("Lexeme,%v", v), nil
	default:
		return "", fmt.Errorf("Unknown type: %T,%v", o, v)
	}
}

func TestTypeSwitch(t *testing.T) {
	fmt.Println(recognizeType2("sophisticated"))
	fmt.Println(recognizeType2(11))
	fmt.Println(recognizeType2(Lexeme("Hello")))
	fmt.Println(recognizeType2(false))
}

// ==>Stringers 重写 toString
// 1.定义一个类型
type IPAddr [4]byte

// 2.实现 Stringer interface
func (addr IPAddr) String() string { // 这里 receiver 较小且无状态, 故不用指针, 使其被分配在栈上
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func TestStringers(t *testing.T) {
	a := IPAddr{127, 0, 0, 1}
	b := IPAddr{192, 168, 0, 1}
	fmt.Println(a, b)
}

//==>Errors 自定义异常

// 1.定义一个结构体表达错误
type MyError struct {
	msg string
}

// 2.实现 error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func run() error {
	return &MyError{"it didn't work"}
}

func TestError(t *testing.T) {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/*
===================== Generics =====================
*/
//[示例1] 实现一个泛型链表
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

func newList[T any]() *List[T] {
	head := *new(Node[T])
	tail := *new(Node[T])
	head.next = &tail
	tail.prev = &head

	l := List[T]{&head, &tail}
	return &l
}

func (list *List[T]) add(ele T) {
	n := Node[T]{val: ele}
	last := list.tail.prev

	last.next = &n
	n.prev = last

	n.next = list.tail
	list.tail.prev = &n
}

func (list *List[T]) get(i int) (T, error) {
	idx := 0
	for node := list.head.next; node != list.tail; node = node.next {
		if idx == i {
			return node.val, nil
		}
		idx++
	}
	return *new(T), nil // *new(T) 创建  zero value
}

func TestList(t *testing.T) {
	list := newList[int]()

	list.add(8)
	list.add(2)
	list.add(4)

	fmt.Println(list.get(2))
}

/*
===================== Goroutines =====================
*/
//===> Goroutine
func httpGet(url string) {
	fmt.Printf("http get %v\n", url)
}
func TestGoroutines(t *testing.T) {
	go httpGet("http://github.com")
	fmt.Println("finished")
}

// ===> Channel

// [示例1] 生产者消费者模型
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //关闭 channel,使得 consumer 那边退出阻塞
}

// 通过channel的返回方式循环接收
func consumer(ch chan int) {
	for d, more := <-ch; more; d, more = <-ch {
		fmt.Printf("received %v\n", d)
	}
}

// range 方式循环接收
func consumer2(ch chan int) {
	for d := range ch {
		fmt.Printf("received %v\n", d)
	}
}

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int, 2) //默认 chan 的容量为 1
	go producer(ch)
	consumer(ch)
}

// [示例2]

func TestCountDownLatch(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		ch <- 1
		println("1")
		ch <- 2
		println("2")
		close(ch)
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//当 channel 关闭以后, 继续读取会得到 zero value
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// [练习] 等价二叉树

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walk0(t, ch)
	close(ch)
}

func Walk0(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk0(t.Left, ch)
	}
	if t.Right != nil {
		Walk0(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 1)
	go Walk(t1, ch1)
	ch2 := make(chan int, 1)
	go Walk(t2, ch2)

	for {
		val1, hasMore1 := <-ch1
		val2, hasMore2 := <-ch2
		if hasMore1 != hasMore2 {
			return false
		} else if val1 != val2 {
			return false
		} else if hasMore1 == false {
			break
		}
	}
	return true
}
func TestEquvialentBinTree(t *testing.T) {
	tree := tree.New(1)
	fmt.Println(Same(tree, tree))
	// fmt.Println(Same(tree.New(1), tree.New(1)))

	// ch1 := make(chan int, 1)
	// go Walk(tree.New(1), ch1)
	// for v := range ch1 {
	// 	fmt.Println(v)
	// }

}

/*
===================== Mutex =====================
*/

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func TestMutext(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
