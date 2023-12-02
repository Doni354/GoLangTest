package say_hello

import "fmt"

func SayHello(name *string) {
    fmt.Println("Hello World!")
    fmt.Println("What's your name?")

    fmt.Scan(name)

    fmt.Println("My Name is", *name)
    fmt.Println("Nice to see you", *name)
}

func BasicMath(a, b, c *int) {
    fmt.Println("BASIC CALCULATOR")
    fmt.Println("Input your first number : ")
    fmt.Scan(a)
    fmt.Println("Input your second number : ")
    fmt.Scan(b)
    fmt.Println(": OPTIONS :")
    fmt.Println("1.Tambah")
    fmt.Println("2.Kurang")
    fmt.Println("3.Kali")
    fmt.Println("4.Bagi")
    fmt.Println("Choose Option")
    fmt.Scan(c)

    switch *c {
    case 1:
        *c = *a + *b
        fmt.Println("Hasil Penjumlahan", *c)
    case 2:
        *c = *a - *b
        fmt.Println("Hasil Pengurangan", *c)
    case 3:
        *c = *a * *b
        fmt.Println("Hasil Perkalian", *c)
    case 4:
        if *b == 0 {
            fmt.Println("Error: Tidak dapat dibagi 0")
        } else {
            *c = *a / *b
            fmt.Println("Hasil Pembagian", *c)
        }
    default:
        fmt.Println("Opsi yang kamu pilih salah")
    }
}
