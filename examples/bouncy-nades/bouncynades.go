package main

import (
	"fmt"
	"log"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang"
	events "github.com/markus-wa/demoinfocs-golang/events"
	ex "github.com/markus-wa/demoinfocs-golang/examples"
	st "github.com/markus-wa/demoinfocs-golang/sendtables"
)

// Run like this: go run bouncynades.go -demo /path/to/demo.dem
func main() {
	f, err := os.Open(ex.DemoPathFromArgs())
	checkError(err)
	defer f.Close()

	p := dem.NewParser(f)

	_, err = p.ParseHeader()
	checkError(err)

	stp := p.SendTableParser()

	p.RegisterEventHandler(func(events.DataTablesParsedEvent) {
		stp.FindServerClassByName("CBaseCSGrenadeProjectile").RegisterEntityCreatedHandler(func(e st.EntityCreatedEvent) {
			prop := e.Entity.FindProperty("m_nBounces")
			prop.RegisterPropertyUpdateHandler(func(st.PropValue) {
				fmt.Println(e.Entity.ID, e.Entity.Position())
			})
		})
	})

	err = p.ParseToEnd()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
