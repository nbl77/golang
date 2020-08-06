package main

import (
    "fmt"
    "grpc/model"
)

func main() {
  var user = &model.User{
    Id:"123",
    Nama:"Nabil",
    Alamat:"Lampung",
  }
  fmt.Println(user)
}
