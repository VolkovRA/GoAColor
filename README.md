# GoAColor (Golang ANSI Colors)

Описание
------------------------------

Очень простенький модуль для возможности раскраски текста в консоли. Использует управляющие ANSI символы, подробнее: [тут](https://ru.wikipedia.org/wiki/%D0%A3%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D1%8F%D1%8E%D1%89%D0%B8%D0%B5_%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8_ANSI "Управляющие последовательности ANSI").
Может не работать в Windows из-за отсутствия поддержки терминалом управляющих ANSI символов. 

Не имеет зависимостей

Пример использования
------------------------------
```
var msg =   acolor.Apply(acolor.Bold, acolor.Green) + "[OK]" + acolor.Clear() + 
            acolor.Apply(acolor.Italic) + " Now this message is beautiful!" + acolor.Clear()

fmt.Println(msg)
```
![Пример вывода](https://github.com/VolkovRA/GoAColor/blob/master/example.png)
