package main

import (
	"bufio"
	"excavation_calculator/hotspots"
	"excavation_calculator/materials"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func main(){
	for i, material := range materials.AllMaterials {
		// Request the HTML page.
		res, err := http.Get(material.Url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".infobox-quantity").Each(func(j int, s *goquery.Selection) {
			// For each item found, get the band and title
			text , _ := s.Attr("data-val-each")
			materials.AllMaterials[i].Value, _ = strconv.Atoi(text)
		})
	}

	for i, z := range hotspots.AllHotspots {
		for _, x := range z.MaterialsWithDecimal{
			for _ , y := range materials.AllMaterials {
				if x.Material.Id == y.Id{
					hotspots.AllHotspots[i].Total = hotspots.AllHotspots[i].Total + (float64(y.Value) * x.Decimal)
				}
			}
		}
	}

	sort.Slice(hotspots.AllHotspots[:], func(i, j int) bool {
		return hotspots.AllHotspots[i].Total > hotspots.AllHotspots[j].Total
	})

	fmt.Println("Name : Average Value of material : Min Level Required")
	for _, h := range hotspots.AllHotspots{
		fmt.Printf("%s : %v : %v\n", h.Name, h.Total/10, h.Level)
	}


	fmt.Println("-----------------------------------------------------")
	fmt.Println("Helper made by ClueHunterX add me in game if you find any bug :)")
	fmt.Print("Press 'Enter' to Exit...\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')


}
