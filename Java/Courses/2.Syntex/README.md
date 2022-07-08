# 2. Синтаксис языка Java

## Переменные
Переменная это именнованная ячейка памяти которое моежт быть присвоенное значение
```Java
int age  = 27;
String userName = "Saydum";
```

## Условный оператор if else
```Java
if (age >= 18) {
    // OK
} else {
    // НЕ OK :)
}
```
`&&` - and
`||` - or

Для сравнения двух строк используется `equels` \
```Java
if (userName.equels(user)) {
    true;
} else {
    false;
}
```
### switch case
```Java
switch (age) {
    case: 18
        // RUN
        break;
    case: 0;
        // NOT RUN
        break;
    default:
        // default code
}

## Типы данных
В Java 8 примитивных и простых типов

```Java
/** Целые числа */ 
byte   1 байт     от -128    до 127
short  2 байт     от -32 768 до 32 767
int    4 байт     от         до
long   8 байт     от         до

/** Дробные числа (Примитивные типы )*/ 
float  4 байт     3.14f
double 8 байт     4.14

```

### Приведение типов
В Java нет автомотического преобразование типов от малого к большему его можно указать вручную.
```Java
int a = 10;

Здесь int присваивается к byte
byte b = a; //!!!ОШИБКА!!!\\

// Приведении типов в ручную
byte b = (byte)a;
```

Если `byte` присвоить к `int` тогда будет автоматическа преобразование типов. \
То есть тип с большим байтом нельзя без ручного преобразования присвоить к малому байту. \
Малый байт преобразуется к большому байту автоматически
```Java
byte a = 10;

int b = a; // ВСЕ OK!
```

## Циклы
- wjile
- for
- do while

```Java
while (true) {
    // RUN
}

for (int i = 0; i <= 10; i++) {
    // RUN
}

do {
   // RUN 
}
while (true);
```

## Массивы

Создание массива \
```Java
int[] myArray = new int[10];
  ↑       ↑              ↑
 Тип     Имя        кол-во ячеек
```
Цыкл для массива `В данном цыкле нельзя изменять массив, т.е присваивать новые значение его элемента.`
```Java
for (int i: myArray) {
    System.out.Println(i);
}
```
Создание массива без `new`. \
`int nums = {1, 3, 5, 7, 8};`


# ООП
- Класс - это шаблон с характеристиками
- Объект - структура данных, содержашийся описание свойст внешнего объекта.

## Констрктор
```java
class Box {
    // Констрктор
    public Box() {
        //...
    }
}
```
## [Перегрузка методов](https://metanit.com/java/tutorial/2.18.php)
Это когда мы можем использовать методы с одним и тем же именем, но с разными типами и/или кол-вом параметров.
```java
class Box {
    public static void main(String[] args) {
        System.out.println(sum(2, 3));          // 5
        System.out.println(sum(4.5, 3.2));      // 7.7
        System.out.println(sum(4, 3, 7));       // 14
    }

    static int sum(int x, int y){         
        return x + y;
    }

    static double sum(double x, double y){
             
        return x + y;
    }

    static int sum(int x, int y, int z){         
        return x + y + z;
    }
}
```
