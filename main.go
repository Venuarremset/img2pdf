package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/jung-kurt/gofpdf"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

//func setupRoutes() {
//	server := http.Server{
//	Addr: "127.0.0.1:8480",
//	}
//	server.ListenAndServe()
//}
func handler(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "index.html", nil)
}

var myTemplate *template.Template

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Hello World")
	myTemplate = template.Must(template.ParseGlob("index.html"))
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+port, nil)

	//setupRoutes()

	f, err := os.Open("image.png") //opening a file with os package
	if err != nil {
		log.Fatal(err) //to display no such file or directory
	}

	img, err := png.Decode(f) //decoding the file f using image/png package
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", img)
	pdf := gofpdf.New("P", "mm", "A4", "") //Create a PDF using gofpdf package (potrait mode in millimiters terms of size A4)
	pdf.AddPage()                          //adds a page
	pdf.SetFont("Arial", "B", 15)          //sets font name, type and size
	pdf.MultiCell(160, 5, "", "0", "0", false)
	pdf.Image("image.png", 50, 100, 0, 100, true, "", 0, "")

	//pdf.ImageOptions(string(content), -10, 10, 30, 0, false, gofpdf.ImageOptions{
	//		ReadDpi: true,
	//	}, 0, "")

	_ = pdf.OutputFileAndClose("hello.pdf")
	//	fmt.Println(err)

}
