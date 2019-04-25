package module

import (
	"fmt"
	"net/http"

	"github.com/gido/2D_WebSocket_Game/server/db"
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

func (WsClient WsClient) LoginHandler(w http.ResponseWriter, r *http.Request) {
	player := WsClient.Player
	fmt.Println(player)
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Could not get login information from client")
	}
	uName := r.FormValue("username")
	uPassword := r.FormValue("password")

	if uName == "" || uPassword == "" {
		http.Redirect(w, r, fmt.Sprintf("./%s/login.html", "/client"), 302)
	}

	authorized, loginData := db.LoginPlayer(db.Database, uName, uPassword)
	if !authorized {
		fmt.Println("Something went wrong, can not log in")

	} else {
		player.ID = loginData.ID
		player.Position.X = loginData.PosX
		player.Position.Y = loginData.PosY

		http.Redirect(w, r, "/world", 302)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Could not get login information from client")
	}
	uName := r.FormValue("username")
	uPassword := r.FormValue("password")

	if uName == "" || uPassword == "" {
		http.Redirect(w, r, fmt.Sprintf("./%s/register.html", "/client"), 302)
	}

	player := Player{ID: GetToken(10), Position: Position{X: 250, Y: 250},
		Velocity: Velocity{X: 3, Y: 3},
		Control:  Control{Right: false, Left: false, Up: false, Down: false}}

	err = db.RegisterPlayer(db.Database, uName, uPassword, player.ID)
	if err != nil {
		fmt.Println("Register player failed")
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func handleStaticFiles() {
	s := http.StripPrefix("/client/", http.FileServer(http.Dir("./client/")))
	router.PathPrefix("/client/").Handler(s)

	router.Handle("/css", http.FileServer(http.Dir("css")))
	router.Handle("/js", http.FileServer(http.Dir("js")))

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r)
	})

}

var router = mux.NewRouter()

func StartApi() {

	handleStaticFiles()
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/world", serveWorld)

	router.HandleFunc("/login", serveLoginPage)
	router.HandleFunc("/register", serveRegisterPage)

	router.HandleFunc("/registerHandler", RegisterHandler).Methods("POST")

	fmt.Println("Server running on port 3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)

	}

}
