package main

import (
  "bytes"
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
    "log"
)
type Profile struct {
  Name    string
  Hobbies []string
}
type Parkir struct {
  id int64 `json:"id"`
  tipe string `json:"tipe"`
  plat string `json:"plat"`
}
type Error struct{
  Status int64
  Message string
}
func main()  {
  http.HandleFunc("/parkir_masuk",home)
  http.HandleFunc("/parkir_keluar",out)
  fmt.Println("Client running...")
  http.ListenAndServe(":8000",nil)
}
func home(w http.ResponseWriter, r *http.Request)  {
  var client = &http.Client{}
  w.Header().Set("Content-Type", "application/json")
  if r.Method == "POST" {
    w.WriteHeader(200)
    request, _ := http.NewRequest("POST","http://localhost:8001/get_id",nil)
    request.Header.Set("Content-Type", "application/json")
    resp,_ := client.Do(request)
    defer resp.Body.Close()
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    w.Write([]byte(string(responseData)))
  }else {
    error := Error{400,"Method Not Allowed"}
    js, err := json.Marshal(error)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.WriteHeader(400)
    w.Write(js)
  }
}
func out(w http.ResponseWriter, r *http.Request)  {
  var client = &http.Client{}
  w.Header().Set("Content-Type", "application/json")
  if r.Method == "POST" {
    w.WriteHeader(200)
    resp2,_ := ioutil.ReadAll(r.Body)
    request, _ := http.NewRequest("POST","http://localhost:8001/get_total",bytes.NewBuffer(resp2))
    request.Header.Set("Content-Type", "application/json")
    resp,_ := client.Do(request)
    defer resp.Body.Close()
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    w.Write([]byte(string(responseData)))
  }else {
    error := Error{400,"Method Not Allowed"}
    js, err := json.Marshal(error)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.WriteHeader(400)
    w.Write(js)
  }
}
