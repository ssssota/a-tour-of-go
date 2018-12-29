package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i1, j1 int = 1, 2

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func hello() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func count() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// LIFO: stack
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

type Vertex struct {
	X, Y float64
}

var (
	v3 = Vertex{1, 2}
	v4 = Vertex{X: 1} // Y:0
	v5 = Vertex{}     // X:0, Y:0
	p3 = &Vertex{1, 2}
)

var pow1 = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

var m2 = map[string]Vertex1{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	x := 1
	y := 0
	return func() int {
		x, y = y, x+y
		return x
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	lastZ, z := x, 1.0

	for math.Abs(z-lastZ) >= 1.0e-6 {
		lastZ, z = z, z-(z*z-x)/(2*z)
	}

	return z, nil
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}

	return n, err
}

func rot13(b byte) byte {
	if b >= 'a' {
		b = (b-'a'+13)%26 + 'a'
	} else if b >= 'A' {
		b = (b-'A'+13)%26 + 'A'
	}
	return b
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sumFunc(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func timeBomb() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("Hello, world")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("Now you have %l problems.", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	var c1, python1, java1 = true, false, "no!"
	fmt.Println(i1, j1, c1, python1, java1)

	var i2, j2 int = 1, 2
	k := 3
	c2, python2, java2 := true, false, "no!"
	fmt.Println(i2, j2, k, c2, python2, java2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i3 int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i3, f, b2, s)

	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f2)
	fmt.Println(x, y, z)

	v1, v2 := "42", 42
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// infinite loop
	/*
		for {
		}
	*/

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	hello()
	count()

	i4, j4 := 42, 2701
	p := &i4
	fmt.Println(*p)
	*p = 21
	fmt.Println(i4)
	p = &j4
	*p = *p / 37
	fmt.Println(j4)

	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	p1 := &v
	(*p1).X = 10
	p1.X = 1e9
	fmt.Println(v)

	fmt.Println(v3, p3, v4, v5)

	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var sl []int = primes[1:4]
	fmt.Println(sl)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a3 := names[0:2]
	b3 := names[1:3]
	fmt.Println(a3, b3)
	b3[0] = "XXX"
	fmt.Println(a3, b3)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s1 := []struct {
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
	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	s2 = s2[1:4]
	fmt.Println(s2)
	s2 = s2[:2]
	fmt.Println(s2)
	s2 = s2[1:]
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)
	s3 = s3[:0]
	printSlice(s3)
	s3 = s3[:4]
	printSlice(s3)
	s3 = s3[2:]
	printSlice(s3)

	var s4 []int
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil {
		fmt.Println("nil!")
	}

	a4 := make([]int, 5)
	printSlice2("a4", a4)
	b4 := make([]int, 0, 5)
	printSlice2("b4", b4)
	c4 := b4[:2]
	printSlice2("c4", c4)
	d4 := c4[2:5]
	printSlice2("d4", d4)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s5 []int
	printSlice(s5)
	s5 = append(s5, 0)
	printSlice(s5)
	s5 = append(s5, 1)
	printSlice(s5)
	s5 = append(s5, 2, 3, 4)
	printSlice(s5)

	for i, v := range pow1 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["Answer"] = 42
	fmt.Println("The value", m3["Answer"])
	m3["Answer"] = 48
	fmt.Println("The value", m3["Answer"])
	delete(m3, "Answer")
	fmt.Println("Ther value", m3["Answer"])

	v6, ok := m3["Answer"]
	fmt.Println("The value", v6, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}

	v7 := Vertex{3, 4}
	fmt.Println(v7.Abs())

	f3 := MyFloat(-math.Sqrt2)
	fmt.Println(f3.Abs())

	v8 := Vertex{3, 4}
	v8.Scale(10)
	Scale(&v8, 0.1)
	fmt.Println(v8)
	fmt.Println(v8.Abs(), Abs(v8))

	p4 := &Vertex{3, 4}
	p4.Scale(10)
	Scale(p4, 0.1)
	fmt.Println(p4)
	fmt.Println(p4.Abs(), Abs(*p4))

	var a5 Abser
	f4 := MyFloat(-math.Sqrt2)
	v9 := Vertex{3, 4}
	a5 = f4
	fmt.Println(a5.Abs())
	//a5 = v9
	a5 = &v9
	fmt.Println(a5.Abs())

	var i5 I
	//describe(i5) // runtime error
	//i5.M()
	var t *T
	i5 = t
	describe(i5)
	i5.M()
	i5 = &T{"hello"}
	describe(i5)
	i5.M()
	i5 = F(math.Pi)
	describe(i5)
	i5.M()

	var i6 interface{}
	describeEmptyInterface(i6)
	i6 = 42
	describeEmptyInterface(i6)
	i7 := i6.(int)
	fmt.Println(i7)
	i6 = "hello"
	describeEmptyInterface(i6)
	s6, ok := i6.(string)
	fmt.Println(s6, ok)
	f5, ok := i6.(float64)
	fmt.Println(f5, ok)
	// f5 = i6.(float64) // panic
	f5, _ = i6.(float64)
	fmt.Println(f)

	do(21)
	do("hello")
	do(true)

	a6 := Person{"Arthur Dent", 42}
	z2 := Person{"Zaphod Beelebrox", 9001}
	fmt.Println(a6, z2)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	r2 := strings.NewReader("Hello, Reader!")
	b5 := make([]byte, 8)
	for {
		n, err := r2.Read(b5)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b5)
		fmt.Printf("b[:n] = %q\n", b5[:n])
		if err == io.EOF {
			break
		}
	}

	s7 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r3 := rot13Reader{s7}
	io.Copy(os.Stdout, &r3)
	fmt.Println()

	m4 := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m4.Bounds())
	fmt.Println(m4.At(0, 0).RGBA())

	go say("world")
	say("hello")

	s8 := []int{7, 2, 8, -9, 4, 0}
	c5 := make(chan int)
	go sumFunc(s8[:len(s8)/2], c5)
	go sumFunc(s8[len(s8)/2:], c5)
	x2, y2 := <-c5, <-c5
	fmt.Println(x2, y2, x2+y2)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 10)
	go fibonacci2(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch3)
		}
		quit <- 0
	}()
	fibonacci3(ch3, quit)

	timeBomb()

	c6 := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c6.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c6.Value("somekey"))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
