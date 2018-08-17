package main

import (
	"fmt"
	"net/http"
  "html"
  "text/template"
	"errors"
	"encoding/csv"
	"os"
	"bufio"
	"log"
	"io"
	"strconv"
)

type Record struct {
  Id int
  Firstname string
  Lastname string
  Phonenumber string
}

type Tpldata struct {
	Record Record
	Nb_item int
}

type error interface {
    Error() string
}

/**
 * aim of this function is to check if requires get parameters are available
 */
func checkParameter(parms [3]string,  r *http.Request) (map[string]string, error){
	result := make(map[string]string)
	fmt.Printf("==> ")
  for i := 0; i < len(parms); i++ {
    value, ok := r.URL.Query()[parms[i]]
    if !ok || len(value) < 1 {

			return result, errors.New(fmt.Sprintf("==> Error %s not filled", parms[i]))
    }
    fmt.Printf("%s=%s ", parms[i], value[0])
    result[parms[i]] = value[0]
  }
	fmt.Printf("\n")
  return result, nil

}


func getRecord(number string, records []Record) ( Record, error){
	for i := 0 ; i < len(records); i++ {
		if records[i].Phonenumber == number {
			return records[i], nil
		}
	}
	// we return first item... but it's wrong.... to be corrected when
	// I known better this langage
	return records[0], errors.New(fmt.Sprintf("number '%s' not found", number))
}

/**
 * List of phonenumber in the csv file
 */
func readCsv(filename string)([]Record){
	csvFile, _ := os.Open(filename)
  reader := csv.NewReader(bufio.NewReader(csvFile))
	//skip first line
	reader.Read()
	var people []Record
	for {
			 line, error := reader.Read()
			 if error == io.EOF {
					 break
			 } else if error != nil {
					 log.Fatal(error)
			 }
			 id, err := strconv.Atoi(line[0])
			 if err == nil {
				 people = append(people, Record {
						 Id: id,
						 Firstname: line[1],
						 Lastname:  line[2],
						 Phonenumber: line[3]})
			 }else{
				 log.Fatal(err)
			 }

	 }
	 return people
}

// command args :
// phonebook CONFIG_FILE LISTEN_PORT
func main() {

	const tpl = `<?xml version="1.0" encoding="UTF-8" ?>
	<list response="get_list" type="pb" reqid="22AB2" total="{{.Nb_item}}" first="1" last="{{.Nb_item}}">
	  <entry id="{{.Record.Id}}">
	    <ln>{{.Record.Firstname}}</ln>
	    <fn>{{.Record.Lastname}}</fn>
	    <hm>{{.Record.Phonenumber}}</hm>
	  </entry>
	</list>
`

	phonebook := readCsv(os.Args[1])

	// phonebook service
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

    fmt.Printf(html.EscapeString(r.URL.Path))

    a := [3]string{"command", "type", "hm"}

    params, err:= checkParameter(a, r)
    if err == nil {
      if params["command"] == "get_list" && params["type"] == "pb" {
				record,err := getRecord(params["hm"], phonebook)
				if err == nil {
					data := Tpldata{ Record: record, Nb_item: len(phonebook) }
					t := template.New("tpl")
					t, err := t.Parse(tpl)
					if err == nil{
						 err := t.Execute(w, data)
						 if err != nil{
							 log.Fatal("==> can't execute template", err)
						 }
					}else{
						log.Fatal("==> can't create template", err)
					}
				}else{
        	log.Fatal("==> number not found")
				}
      }else{
        log.Fatal("==> params not handled")
      }
    }
	})

	// launch deamon
	err := http.ListenAndServe(":"+os.Args[2], nil)
  if err != nil {
        fmt.Printf("ListenAndServe:", err,"\n")
  }
}
