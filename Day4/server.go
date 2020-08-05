package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "strconv"
  "time"
  "encoding/json"
  "math/rand"
)
var timePark = make([]int64, 0)
var park = make([]string, 0)
type Send struct {
  Status int64
  Id string
  Message string
}
type Customers struct {
  Id int64 `json:"id"`
  Tipe string `json:"tipe"`
  Plat string `json:"plat"`
}
func main()  {
  http.HandleFunc("/get_id",sendID)
  http.HandleFunc("/get_total",getTotal)
  fmt.Println("Server is running...")
  http.ListenAndServe(":8001",nil)
}
func sendID(w http.ResponseWriter, r *http.Request)  {
  send_id := Send{200,servParkIn(),"Masukan ID parkir untuk keluar"}
  w.WriteHeader(200)
  w.Header().Set("Content-Type", "application/json")
  js, _ := json.Marshal(send_id)
  w.Write(js)
}
func getTotal(w http.ResponseWriter, r *http.Request)  {
  resp,_ := ioutil.ReadAll(r.Body)
  var customer Customers
  err := json.Unmarshal(resp, &customer)
  if err != nil {
        fmt.Println(err.Error())
        return
    }
  res := servParkOut(customer.Tipe,customer.Plat, strconv.Itoa(int(customer.Id)))
  var result = Send{200, strconv.Itoa(int(customer.Id)), res}
  js, _ := json.Marshal(result)
  w.Write(js)
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
