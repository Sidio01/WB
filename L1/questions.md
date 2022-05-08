Вопрос 1. Какой самый эффективный способ конкатенации строк?

```golang
package main

import "strings"

func main() {
	var str strings.Builder

	str.Grow(32)

	str.WriteString("hello")
	str.WriteRune(' ')
	str.WriteString("world")
	str.WriteRune(' ')
	str.WriteString("!!!")

	println(str.String())
}
```

Вопрос 2. Что такое интерфейсы, как они применяются в Go?

Интерфейс - это сущность, определяющая поведение структуры и описывающая методы, которые должны быть реализованы для других структур, которые будут удовлетворять этому интерфейсу. Интерфейсы дают возможность изменения поведения разных структур, обладающих одними и теми же методами.

Вопрос 3. Чем отличаются RWMutex от Mutex?

Mutex блокирует участок памяти на чтение и запись. RWMutex делает доступным чтение для нескольких потоков.

Вопрос 4. Чем отличаются буферизированные и не буферизированные каналы?

Небуферизированный канал принимает одно значение и затем блокируется до момента чтения. Буферизированный может принять несколько значений до блокировки.

Вопрос 5. Какой размер у структуры struct{}{}?
0

Вопрос 6. Есть ли в Go перегрузка методов или операторов?

Нет. Ответ из документации:
Why does Go not support overloading of methods and operators?

Method dispatch is simplified if it doesn't need to do type matching as well. Experience with other languages told us that having a variety of methods with the same name but different signatures was occasionally useful but that it could also be confusing and fragile in practice. Matching only by name and requiring consistency in the types was a major simplifying decision in Go's type system.

Regarding operator overloading, it seems more a convenience than an absolute requirement. Again, things are simpler without it. 

В общих словах, в целях упрощения синтаксиса языка.

Вопрос 7. В какой последовательности будут выведены элементы map[int]int?

Пример:
```golang
m[0]=1
m[1]=124
m[2]=281
```

В порядке возрастания значений ключей.

Вопрос 8. В чем разница make и new?
make - создает новый слайс, канал, мапу, а new создает указатель на данные типы.

Вопрос 9. Сколько существует способов задать переменную типа slice или map?

```golang
var x map[int]int
var x map[int]int{0:1, 1:0}
var x = make(map[int]int)
```

Вопрос 10. Что выведет данная программа и почему?
```golang
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
1, так как изменение значение переменной происходит локально внутри другой функции.
a - 0xc000132000 1
p before update - 0xc000132000 1
b - 0xc000132010 2
p from update - 0xc000132010 2
p after update - 0xc000132000 1

Вопрос 11. Что выведет данная программа и почему?

```golang
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
0,1,2,3,4 в любом порядке и затем deadlock, так как WaitGroup передается не по указателю.

Вопрос 12. Что выведет данная программа и почему?

```golang
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
0, так как переменная n объявлена в локальном окружении. Чтобы ответ был 2, необходимо удалить повторное объявление переменной и просто присвоить ей значение 1.

Вопрос 13. Что выведет данная программа и почему?

```golang
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
Слайс характеризуется дополнительными свойствами как length и capacity. При инициализации слайса со значениями, capacity равно количеству элементов. Если при append количество элементов в слайсе будет больше capacity, то создасться новый слайс, поэтому с будет изменен только первый элемент, а в остальном слайс останется прежним.

Вопрос 14. Что выведет данная программа и почему?

```golang
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
[b b a][a a]. Слайсы всегда передаются по указателю, но так как сначала в функции происходит append, то в дальнейшем работа идет уже с другим слайсом (из-за увеличения capacity), поэтому первоначальный слайс останется неизменным.