package main

import (
  "context"
  "log"
  "fmt"
  "Day5/model"
  "strconv"
  "time"
  "github.com/golang/protobuf/ptypes/empty"
  "google.golang.org/grpc"
)

func QueueService() model.ParkirServiceClient {
  port := ":9000"
  conn, err := grpc.Dial(port, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect to ",port, err)
  }
  return model.NewParkirServiceClient(conn)
}

func main()  {
  queueServ := QueueService()
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
      result := GetID(queueServ)
      log.Println("ID Parkir Anda :",result)
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
      parkir := &model.Parkir{
        Id: id_parkir,
        Tipe: tipe,
        Plat: plat,
      }
      result := ParkingExit(queueServ, parkir)
      log.Println(result)
    case "3":
      result := GetList(queueServ)
      log.Println(result)
    case "4":
      fmt.Println("Terimakasih telah menggunakan layanan parkir kami")
    default:
      fmt.Println("Opsi yang anda masukan salah")
    }
  }
}

func GetID(queueServ model.ParkirServiceClient) string {
  resp,err := queueServ.GetID(context.Background(), new(empty.Empty))
  if err != nil {
    log.Fatalf(err.Error())
  }
  return resp.Id
}

func GetList(queueServ model.ParkirServiceClient) string {
  teks := "ID Parkir Yang Tersedia :"
  resp,err := queueServ.GetList(context.Background(), new(empty.Empty))
  if err != nil {
    log.Fatalf(err.Error())
  }
  if len(resp.QList) == 0 {
    return "Tidak ada kendaraan yang sedang parkir!"
  }
  for _,val := range resp.QList{
    wkt := (int(time.Now().Unix()) - int(val.Time))
    teks += "\nID Parkir :" + val.Id +"\n"
    teks += "Waktu Berlalu :" + strconv.Itoa(wkt) + "\n"
    teks += "============================\n"
  }
  return teks
}

func ParkingExit(queueServ model.ParkirServiceClient, parkir *model.Parkir) string {
  resp,err := queueServ.Exit(context.Background(), parkir)
  if err != nil {
    log.Fatalf(err.Error())
  }
  if resp.Status == 404 {
    return resp.Message
  }
  return resp.Message + " " + resp.Second + ", Total Yang harus Dibayar :" + resp.Total
}
func makeMenu(menu ...string)  {
  fmt.Println("===== Menu =====")
  for key,val := range menu{
    a:= strconv.Itoa(key + 1)
    fmt.Println(a + ". " + val)
  }
  fmt.Println("Silahkan Pilih :")
}
