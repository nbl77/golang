package main
import (
    "fmt"
    "math"
    "strconv"
)
func main() {
  var inp string
  for inp != "99"{
    fmt.Println("Pilih :")
    fmt.Println("1.Perkalian")
    fmt.Println("2.Bagi")
    fmt.Println("3.Tambah")
    fmt.Println("4.Kurang")
    fmt.Println("5.Akar")
    fmt.Println("6.Pangkat")
    fmt.Println("7.Luas Persegi")
    fmt.Println("8.Luas Lingkaran")
    fmt.Println("9.Volume tabung")
    fmt.Println("10.Volume balok")
    fmt.Println("11.Volume prisma")
    fmt.Println("====================")
    fmt.Scanln(&inp)
    switch inp {
    case "1":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil Perkalian :",perkalian(newA, newB))
    case "2":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil pembagian :",pembagian(newA, newB))
    case "3":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil penjumlahan :",penjumlahan(newA, newB))
    case "4":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil pengurangan :",pengurangan(newA, newB))
    case "5":
      fmt.Println("Masukan Nilai :")
      var a float64
      fmt.Scanln(&a)
      fmt.Println("Hasil akar :",akar(a))
    case "6":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil pangkat :",pangkat(newA, newB))
    case "7":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil luas persegi :",luas_persegi(newA, newB))
    case "8":
      fmt.Println("Masukan Nilai :")
      var a string
      fmt.Scanln(&a)
      newA,_ := strconv.Atoi(a)
      fmt.Println("Hasil luas lingkaran :",luas_lingkaran(newA))
    case "9":
      fmt.Println("Masukan Nilai :")
      var a, b string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      fmt.Println("Hasil Volume Tabung :",volume_tabung(newA, newB))
    case "10":
      fmt.Println("Masukan Nilai :")
      var a, b, c string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      fmt.Scanln(&c)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      newC,_ := strconv.Atoi(c)
      fmt.Println("Hasil Volume balok :",volume_balok(newA, newB, newC))
    case "11":
      fmt.Println("Masukan Nilai :")
      var a, b, c string
      fmt.Scanln(&a)
      fmt.Scanln(&b)
      fmt.Scanln(&c)
      newA,_ := strconv.Atoi(a)
      newB,_ := strconv.Atoi(b)
      newC,_ := strconv.Atoi(c)
      fmt.Println("Hasil Volume prisma :",volume_prisma(newA, newB, newC))
    }
  }
}
func perkalian(a ...int) int {
  res := 1
  for _,val := range a{
    res = res * val
  }
  return res
}
func pembagian(a,b int) int {
  return a/b
}
func penjumlahan(a ...int) int {
  res := 0
  for _,val := range a{
    res += val
  }
  return res
}
func pengurangan(a ...int) int {
  res := a[0]
  for key,val :=range a{
    if key > 0 {
        res -= val
      }
    }
  return res
}
func akar(a float64) float64 {
  return math.Sqrt(a)
}
func pangkat(a,b int) int {
  res := 1
  for i := 0; i < b; i++ {
    res = res * a;
  }
  return res
}
func luas_persegi(a,b int) int {
  return a*b
}
func luas_lingkaran(a int) int {
  return (22/7)*(a*a)
}

func volume_tabung(a,b int) int {
  return (22/7)*a*a*b
}
func volume_balok(a,b,c int) int {
  return a*b*c
}
func volume_prisma(a,b,c int) int {
  return a*b*c
}
