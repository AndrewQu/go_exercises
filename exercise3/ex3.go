package main

import (
   "fmt"
   "strings"
)

func main() {
   fmt.Println("***** Exercise 3 *****")
   fmt.Println("*** Try map")
   try_map()
   count_string := "The quick fox jumped over the lazy dog lying on the ground."
   wc := try_countWord(count_string)
   fmt.Printf("Word counts in \"%v\"\n", count_string)
   fmt.Println(wc)
}

type Point2d struct {
   X, Y float64
}

func try_countWord(s string) map [string] int {
   str_map := make(map [string] int)
   words := strings.Fields(s)
   for _, w := range words {
      w = strings.Trim(w, ",.:;!?")
      n := str_map[w]
      str_map[w] = 1 + n
   }
   return str_map
}

func try_map(){
   city_locations := make(map[string]Point2d)
   city_locations["London"] = Point2d { 200.0, 102.0 }
   city_locations["New York"] = Point2d { 201.0, 222.0 }
   city_locations["Paris"] = Point2d { 198.0, 99.0 }
   fmt.Println("Position of London = ", city_locations["London"])
   for k, v := range city_locations {
      fmt.Printf("City: %-10s  Location: %v\n", k, v)
   }
   
   more_cities := map[string]Point2d {
      "Berlin": { 200.0, 133.3 },
      "Rio de Jeniro": { 300.0, 345.0 },
      "Tokyo": { 121.0, 54.0 },
   }
   fmt.Println("\nMore cities     Location")
   for k, v := range more_cities {
      fmt.Printf("%-15s %v\n", k, v)
   }
   
   loc_deli, mapped := city_locations["Deli"]
   fmt.Println("Location of Deli", loc_deli, "  Mapped ?", mapped )
   
   more_cities["Deli"] = Point2d { 102.11, 34.55 }
   loc_deli, mapped = more_cities["Deli"]
   fmt.Println("Location of Deli", loc_deli, "  Mapped in more_cities ?", mapped )
   
   delete(more_cities, "Tokyo");
   loc, mapped := more_cities["Tokyo"]
   fmt.Println("Deleted city Tokyo ?", !mapped, loc)
}
