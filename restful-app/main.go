package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"sync"
)

type Album struct {
	ID     int
	Name   string
	Artist string
	Year   int
}

var (
	albums      []Album
	albumsMutex sync.Mutex
	csvFile     = "albums.csv"
	templates   = template.Must(template.ParseGlob("templates/*.html"))
)

func main() {
	// Load albums from CSV
	loadAlbums()

	// Define routes
	http.HandleFunc("/", listAlbumsHandler)
	http.HandleFunc("/new", newAlbumHandler)
	http.HandleFunc("/edit", editAlbumHandler)
	http.HandleFunc("/delete", deleteAlbumHandler)
	http.HandleFunc("/save", saveAlbumHandler)

	// Start the server
	port := ":8080"
	log.Printf("Starting server on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadAlbums() {
	log.Println("Loading albums from CSV")
	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	file, err := os.Open(csvFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No CSV file found, starting with an empty album list")
			return // No CSV file yet
		}
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	albums = nil
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		year, _ := strconv.Atoi(record[3])
		album := Album{
			ID:     id,
			Name:   record[1],
			Artist: record[2],
			Year:   year,
		}
		albums = append(albums, album)
		log.Printf("Loaded album: %+v", album)
	}
	log.Println("Finished loading albums from CSV")
}

func saveAlbums() {
	log.Println("Saving albums to CSV")
	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	// Create the CSV file
	file, err := os.Create(csvFile)
	if err != nil {
		// Unlock the mutex before logging a fatal error
		albumsMutex.Unlock()
		log.Fatalf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write each album to the CSV file
	for _, album := range albums {
		record := []string{
			strconv.Itoa(album.ID),
			album.Name,
			album.Artist,
			strconv.Itoa(album.Year),
		}
		if err := writer.Write(record); err != nil {
			log.Printf("Error writing record to CSV: %v", err)
		} else {
			log.Printf("Saved album to CSV: %+v", album)
		}
	}
	log.Println("Finished saving albums to CSV")
}

func listAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering listAlbumsHandler")
	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	if err := templates.ExecuteTemplate(w, "list.html", albums); err != nil {
		log.Printf("Error rendering template in listAlbumsHandler: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	log.Println("Exiting listAlbumsHandler")
}

func newAlbumHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering newAlbumHandler")
	if err := templates.ExecuteTemplate(w, "form.html", nil); err != nil {
		log.Printf("Error rendering template in newAlbumHandler: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	log.Println("Exiting newAlbumHandler")
}

func editAlbumHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering editAlbumHandler")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	log.Printf("Editing album with ID: %d", id)

	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	for _, album := range albums {
		if album.ID == id {
			log.Printf("Album found: %+v", album)
			if err := templates.ExecuteTemplate(w, "form.html", album); err != nil {
				log.Printf("Error rendering template in editAlbumHandler: %v", err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			log.Println("Exiting editAlbumHandler")
			return
		}
	}

	log.Printf("Album with ID %d not found", id)
	http.NotFound(w, r)
	log.Println("Exiting editAlbumHandler")
}

func deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering deleteAlbumHandler")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	log.Printf("Deleting album with ID: %d", id)

	albumsMutex.Lock()
	// Find and remove the album
	for i, album := range albums {
		if album.ID == id {
			log.Printf("Album found and deleted: %+v", album)
			albums = slices.Delete(albums, i, i+1)
			albumsMutex.Unlock() // Unlock the mutex before calling saveAlbums
			saveAlbums()
			log.Println("Exiting deleteAlbumHandler")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	albumsMutex.Unlock() // Unlock the mutex if the album is not found

	log.Printf("Album with ID %d not found", id)
	http.NotFound(w, r)
	log.Println("Exiting deleteAlbumHandler")
}

func saveAlbumHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering saveAlbumHandler")
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	artist := r.FormValue("artist")
	year, _ := strconv.Atoi(r.FormValue("year"))

	log.Printf("Received form data - ID: %d, Name: %s, Artist: %s, Year: %d", id, name, artist, year)

	// Lock the mutex only for modifying the albums slice
	albumsMutex.Lock()
	if id == 0 {
		// New album
		id = len(albums) + 1
		newAlbum := Album{ID: id, Name: name, Artist: artist, Year: year}
		albums = append(albums, newAlbum)
		log.Printf("New album added: %+v", newAlbum)
	} else {
		// Update existing album
		for i, album := range albums {
			if album.ID == id {
				updatedAlbum := Album{ID: id, Name: name, Artist: artist, Year: year}
				albums[i] = updatedAlbum
				log.Printf("Album updated: %+v", updatedAlbum)
				break
			}
		}
	}
	albumsMutex.Unlock() // Unlock the mutex after modifying the albums slice

	// Save albums to CSV (this function already locks the mutex)
	saveAlbums()

	log.Println("Exiting saveAlbumHandler")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
