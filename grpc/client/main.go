package main

import (
    "context"
    "log"
    // "net"
    "fmt"
    "grpc/common/config"
    "grpc/common/model"
    "encoding/json"
    "github.com/golang/protobuf/ptypes/empty"
    "google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
   port := config.SERVICE_GARAGE_PORT
   conn, err := grpc.Dial(port, grpc.WithInsecure())
   if err != nil {
       log.Fatal("could not connect to", port, err)
   }

   return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
    port := config.SERVICE_USER_PORT
    conn, err := grpc.Dial(port, grpc.WithInsecure())
    if err != nil {
        log.Fatal("could not connect to", port, err)
    }

    return model.NewUsersClient(conn)
}

func main() {
  user1 := model.User{
        Id:       "n002",
        Name:     "Ahmad Nabil",
        Password: "Wtpmjgda",
        Gender:   model.UserGender(model.UserGender_value["MALE"]),
    }

    // garage1 := model.Garage{
    //     Id:   "q001",
    //     Name: "Quel'thalas",
    //     Coordinate: &model.GarageCoordinate{
    //         Latitude:  45.123123123,
    //         Longitude: 54.1231313123,
    //     },
    // }

    // serviceUser := serviceUser()
    garageService := serviceGarage()
    // addUser(serviceUser, user1)
    // showUser(serviceUser)
    // addGarage(garageService, user1, garage1)
    showGarage(garageService, user1)

}
func addUser(serviceUser model.UsersClient, user model.User)  {
  // register user1
  serviceUser.Register(context.Background(), &user)
  fmt.Println("\n", "===========> user test")
}
//
func showUser(serviceUser model.UsersClient)  {
  // show all registered users
  res1, err := serviceUser.List(context.Background(), new(empty.Empty))
  if err != nil {
      log.Fatal(err.Error())
  }
  res1String, _ := json.Marshal(res1.List)
  log.Println(string(res1String))
}

func addGarage(garageService model.GaragesClient, user model.User, garage model.Garage)  {
  fmt.Println("\n", "===========> garage test A")

  garageService.Add(context.Background(), &model.GarageAndUserId{
     UserId: user.Id,
     Garage: &garage,
  })
}

func showGarage(garageService model.GaragesClient, user model.User)  {
  // show all garages of user1
  res2, err := garageService.List(context.Background(), &model.GarageUserId{UserId: user.Id})
  if err != nil {
     log.Fatal(err.Error())
  }
  res2String, _ := json.Marshal(res2.List)
  log.Println(string(res2String))
}
