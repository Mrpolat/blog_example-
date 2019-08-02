package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt" 
    "net/http"
    "text/template"
    "strings"
    "log"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	    for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Polat!") // write data to response

}
func Login(w http.ResponseWriter, r *http.Request) {
	
    
	fmt.Println("method:", r.Method) //get request method
	     if r.Method == "GET" {
		    t, _ := template.ParseFiles("login.html")
		    t.Execute(w, nil)
	} else {
		    r.ParseForm()
		    t, _ := template.ParseFiles("login.html")
		    // we logged into the port
		    if r.Form["username"][0] == "jack" && r.Form["password"][0] == "daniel" {
			//"if" checked
			blogYazisi := getProduct("")
			fmt.Println("Deneme: " + blogYazisi.Baslik + " - " + blogYazisi.Icerik)
			t, _ = template.ParseFiles("sql.html")
			t.Execute(w, map[string]string{
				"a": blogYazisi.Baslik,
				"b": blogYazisi.Icerik,
			})
                
		} else {
			t, _ = template.ParseFiles("login.html")
			t.Execute(w, nil)
			fmt.Println("giriş yanlış")
		}
	}
}

type Product struct {
	Baslik , Icerik string
}

var db *sql.DB
var err error

func getProduct(productCode string) (Product) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog")

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()
	
	if err != nil {
		fmt.Println(err.Error())
	}
    var p Product
        i := 2
	    err = db.QueryRow("select baslik , icerik from blogexample WHERE id = ?", i).Scan(&p.Baslik, &p.Icerik)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("baslik: %s\n icerik: %s\n ", p.Baslik, p.Icerik)
	
    return p
    
}

func main(){
    
    fmt.Println("Go MySQL Tutorial")

        db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog")

   
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
        insert, err := db.Query("INSERT INTO `blogexample` (`baslik`, `icerik`) VALUES ('selam', 'merhaba');")

    if err != nil {
	    panic(err.Error())
	
	}
    defer insert.Close()

        fmt.Println("Successfully inserted into user tables")

        getProduct("g43")

    fmt.Println("Access URL: localhost:9090/login")
	http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", Login)
	err = http.ListenAndServe(":9090", nil) // setting listening porterr 
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}




