package module

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gido/2D_WebSocket_Game/db"
	"github.com/gorilla/mux"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/index.html", "/client"))
}
func serveWorld(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/world.html", "/client"))
}
func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/login.html", "/client"))
}
func serveRegisterPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/register.html", "/client"))
}

// LoginHandler takes care of authenticating Form data and sending player info to client
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("Could not get login information from client")
	}
	uName := r.FormValue("username")
	uPassword := r.FormValue("password")

	if uName == "" || uPassword == "" {
		http.Redirect(w, r, fmt.Sprintf("./%s/login.html", "/client"), 302)
	}

	authorized, loginData := db.LoginPlayer(db.Database, uName, uPassword)
	if !authorized {
		http.Redirect(w, r, fmt.Sprintf("/login"), 302)
	} else {
		// Redirect to world after succesfull login and send data about player via url params
		log.Println("User successfully loged in")

		http.Redirect(w, r, fmt.Sprintf("/world?ID=%s&PosX=%f&PosY=%f&Class=%s", loginData.ID, loginData.PosX, loginData.PosY, loginData.Class), 302)
	}
}

//RegisterHandler takes care of registering players to database
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not get login information from client")
		http.Redirect(w, r, "/register", 302)
	}
	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	uClass := r.FormValue("check")

	if uName == "" || uPassword == "" || uClass == "" {
		http.Redirect(w, r, fmt.Sprintf("/register"), 302)

	} else {
		player := db.PlayerInfo{ID: GetToken(10), PosX: 250, PosY: 250, Class: uClass}

		// Create Inventory for new Player
		err = db.CreateInventoryTable(db.Database, player.ID)
		if err != nil {
			log.Println("Can not create InventoryTable: ", err)
		}

		// Register new Player
		err = db.RegisterPlayer(db.Database, uName, uPassword, player)
		if err != nil {
			log.Println("Register user failed: ", err)
			http.Redirect(w, r, "/register", 302)
		} else {
			log.Println("User registered")
			http.Redirect(w, r, "/login", 302)
		}
	}

}

// Handle all static files
func handleStaticFiles() {
	s := http.StripPrefix("/client/", http.FileServer(http.Dir("./client/")))
	router.PathPrefix("/client/").Handler(s)

	router.Handle("/css", http.FileServer(http.Dir("css")))
	router.Handle("/js", http.FileServer(http.Dir("js")))

}

var router = mux.NewRouter()

// StartApi start handle functions and server
func StartAPI() {
	router.HandleFunc("/login", serveLoginPage)
	router.HandleFunc("/register", serveRegisterPage)

	handleStaticFiles()
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/world", serveWorld)
	router.HandleFunc("/ws", ServeWs)

	router.HandleFunc("/registerHandler", RegisterHandler).Methods("POST")
	router.HandleFunc("/loginHandler", LoginHandler).Methods("POST")

	log.Println("Server running on port 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)

	}

}
