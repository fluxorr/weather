package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
    Location struct{
      Name string `json:"name"`
      Region string `json:"region"`
      Country string `json:"country"`
    } `json:"location"`
    Current struct{
      TempC  float64 `json:"temp_c"`
      TempF  float64 `json:"temp_f"`
      Condition struct{
        Text string `json:"text"`
      } `json:"condition"`
      FeelsLike float64 `json:"feelslike_c"`
    } `json:"current"`
}


func main () {
    res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=0335f5cdfbfc40318ca42339253003&q=Pune")
    if err != nil {
      panic(err)
    }

    defer res.Body.Close()

    if res.StatusCode != 200 {
      panic("The api is currently not available")
    }

    body, err := io.ReadAll(res.Body)
    if err != nil{
      panic(err)
    }
    // fmt.Println(string(body))

    var weather Weather
    err = json.Unmarshal(body,  &weather)
    if err != nil {
      panic(err)
    }

    City , State , Country := weather.Location.Name, weather.Location.Region, weather.Location.Country
    fmt.Println(" City: ",City ,"\n", "State: ", State,"\n","Country: ", Country)

    Tempc, TempF, Condition, FeelsLikeC := weather.Current.TempC  , weather.Current.TempF, weather.Current.Condition.Text , weather.Current.FeelsLike

    fmt.Printf(" %s Temp: %.2f°C\n Temp(F): %.2f°F\n Feels Like: %.2f°C\n",Condition, Tempc, TempF, FeelsLikeC)

}


