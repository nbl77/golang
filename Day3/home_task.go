package main

import (
  "fmt"
  "strconv"
  "time"
  "math/rand"
)
var timePark = make([]int64, 0)
var park = make([]string, 0)

func main()  {
  var key string

  for key != "4"{
    makeMenu(
      "Parkir Masuk",
      "Parkir Keluar",
      "Lihat id parkir kendaraan",
      "Keluar",
    )
    fmt.Scanln(&key)
    switch key {
    case "1":
      fmt.Println("ID Parkir anda",servParkIn())
    case "2":
      var tipe string
      var plat string
      var id_parkir string
      fmt.Println("Masukan Tipe kendaraan anda : ")
      fmt.Scanln(&tipe)
      fmt.Println("Masukan plat kendaraan anda : ")
      fmt.Scanln(&plat)
      fmt.Println("Masukan id parkir anda : ")
      fmt.Scanln(&id_parkir)
      fmt.Println(servParkOut(tipe, plat, id_parkir))
    case "3":
      fmt.Println(park)
    case "4":
      fmt.Println("Terimakasih telah menggunakan layanan parkir kami")
    default:
      fmt.Println("Opsi yang anda masukan salah")
    }
  }
}
func makeMenu(menu ...string)  {
  fmt.Println("===== Menu =====")
  for key,val := range menu{
    a:= strconv.Itoa(key + 1)
    fmt.Println(a + ". " + val)
  }
  fmt.Println("Silahkan Pilih :")
}

func servParkIn() string {
  lenBef := len(park)
  timePark = append(timePark, time.Now().Unix())
  park = append(park, strconv.Itoa(rand.Intn(1000000)))
  return park[lenBef]
}
func servParkOut(tipe, plat, id_parkir string) string {
  parkir := false
  var timeParkir int64

  for key, val := range park{
    if id_parkir == val{
      parkir = true
      timeParkir = timePark[key]
      RemoveIndexInt64(&timePark, key)
      RemoveIndexStr(&park, key)
      break
    }
  }
  if !parkir {
    return "ID Parkir tidak ditemukan"
  }
  total := 0
  currentTime := time.Now().Unix() - timeParkir
  switch tipe {
  case "mobil":
    for i := 0; i < int(currentTime); i++ {
        if i < 1 {
          total += 5000
          continue
        }
        total += 3000
    }
  case "motor":
    for i := 0; i < int(currentTime); i++ {
        if i < 1 {
          total += 3000
          continue
        }
        total += 2000
    }
  default:
    return "Tipe yang anda masukan tidak tersedia"
  }
  return "Waktu parkir anda " + strconv.Itoa(int(currentTime)) + " detik, Jumlah uang yang harus anda bayar adalah " + strconv.Itoa(total)
}

func RemoveIndexInt64(arr *[]int64, index int) {
    newArr := *arr
    *arr = append(newArr[:index], newArr[index+1:]...)
}
func RemoveIndexStr(arr *[]string, index int) {
  newArr := *arr
  *arr = append(newArr[:index], newArr[index+1:]...)
}
