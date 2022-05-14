# Решение задач

## 1
Дано трехзначное число. Найдите сумму его цифр. 

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16
```GO
    var nb int
    fmt.Scan(&nb)
    snb := strconv.Itoa(nb)
    n1, err := strconv.Atoi(snb[0:1])
    n2, err := strconv.Atoi(snb[1:2])
    n3, err := strconv.Atoi(snb[2:3])
    
    if err != nil {
        panic(err)
    }
    fmt.Print(n1 + n2 + n3)
```

## 

```GO
```