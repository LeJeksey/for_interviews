## Варианты задач

Задачи читать по одному пункту, не читать целиком, если хочешь выполнять
### "создай словарь" 
<details><summary></summary>
Прям писать + приложил тестик для проверки myDict_test.go
<details>
<summary>1)</summary> Наш словарь должен уметь:
   - Get(key string) (interface{}, bool)
   - Set(key string, value interface{})</details>
<details><summary>2)</summary>  Проверь, продумал ли ты то, как будут создавать твой словарь (минус карма, если нет, но сейчас самое время продумать)</details>
<details><summary>3)</summary> Потом сделай ее потокобезопасной (на это задание еще и пытаются вопросами сами навести, типа "а что тут не так")</details>
</details>

### "append"
<details><summary></summary>
Решать предлагают без возможности запустить код
<details><summary>1)</summary> 
Что выведет код:

    x := make([]int, 0, 4)

	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)

	fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
	fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
	fmt.Printf("%v len=%d cap=%d\n", z, len(z), cap(z))
</details>
<details><summary>2)</summary>
Что выведет код? 

    x := make([]int, 0, 3)

    x = append(x, 1, 2, 3)
    y := append(x, 4)
    z := append(x, 5)

    fmt.Printf("%v len=%d cap=%d\n", x, len(x), cap(x))
    fmt.Printf("%v len=%d cap=%d\n", y, len(y), cap(y))
    fmt.Printf("%v len=%d cap=%d\n", z, len(z), cap(z))
</details>
<details><summary>3)</summary>
В чем отличие? Почему выводы отличаются?
</details>
</details>

### "Конкурентность внутри цикла" 
<details><summary></summary>
Запускать также не дают
<details><summary>1)</summary>
Что выведет код и почему?
    
        func main() {
            var wg sync.WaitGroup
            for i := 0; i < 10; i++ {
                wg.Add(1)
                go func() {
                    fmt.Println(i)
                    wg.Done()
                }()
            }
            wg.Wait()
        }
</details>
<details><summary>2)</summary>
Как исправить, чтобы вывести числа от 0 до 9? (на порядок пофигу)
</details>
</details>

### Конкурентная сумма квадратов
<details><summary></summary>
Посчитать сумму квадратов чисел от 0 до 1000. Но сделать это в несколько потоков. (решение писать лень)
</details>

