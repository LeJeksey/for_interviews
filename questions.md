### ООП
Могут расспрашивать про ООП в го. Хорошо могут поймать на вопросе "как организовано наследование в go".
<details>
Тут важно понимать, что привычного наследования нет, но есть композиция, которая может перекрыть какие-то кейсы.

Можно пожевать пример

    type Transport struct{}
    func (t *Transport) Move() {
        fmt.Println("I'm moving")
    }  

    type Car struct {
        Transport
        Wheels int
    }
    type Plane struct {
        Transport
        Wings int
    }

То что мы сделали - это композиция. Теперь мы можем вызывать методы Transport через Car и Plane. Но если мы переопределим метод Move в Car, то он перекроет метод Move в Transport. То есть, если мы вызовем Move у Car, то вызовется метод Move из Car, а не из Transport.
</details>

### Конкурентность, параллельность
- могут душить по поводу разницы между конкурентностью и параллельностью
- 

### Прочее
- как проверить слайс на пустоту?
- 