package main
import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    //"container/list"
    )

func main() {
   //queue := list.New()
   command := make(chan string)
   output := make(chan []byte)
   http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
	//queue.PushBack(r.FormValue("run"))
	//fmt.Println(queue.Front())
	command <- r.FormValue("run");
	fmt.Fprintf(w, string(<- output)+"\n")
   });

   http.HandleFunc("/poll", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, <-command);
   });

   http.HandleFunc("/output", func(w http.ResponseWriter, r*http.Request) { 
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println(string(reqBody))
	output <- reqBody
   });

   fmt.Println("Server Started.....")
   log.Print(http.ListenAndServe(":8081", nil))
   }
