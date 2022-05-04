Вопрос 1. Какой самый эффективный способ конкатенации строк?


Вопрос 2. Что такое интерфейсы, как они применяются в Go?


Вопрос 3. Чем отличаются RWMutex от Mutex?


Вопрос 4. Чем отличаются буферизированные и не буферизированные каналы?


Вопрос 5. Какой размер у структуры struct{}{}?


Вопрос 6. Есть ли в Go перегрузка методов или операторов?


Вопрос 7. В какой последовательности будут выведены элементы map[int]int?

Пример:
```golang
m[0]=1
m[1]=124
m[2]=281
```


Вопрос 8. В чем разница make и new?


Вопрос 9. Сколько существует способов задать переменную типа slice или map?


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