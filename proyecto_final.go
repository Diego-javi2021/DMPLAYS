package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Usuario struct {
	Correo   string `json:"correo"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
}

type Video struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
}

type Plataforma struct {
	Usuarios []Usuario
	Videos   []Video
	Logueado *Usuario
}

var plataforma = Plataforma{}

// ///////////////////////////////////
// HABILITAR CORS
// ///////////////////////////////////
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

// ///////////////////////////////////
// REGISTRAR USUARIO
// ///////////////////////////////////
func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevo Usuario
	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range plataforma.Usuarios {
		if u.Correo == nuevo.Correo {
			http.Error(w, "Correo ya registrado", http.StatusBadRequest)
			return
		}
	}

	plataforma.Usuarios = append(plataforma.Usuarios, nuevo)
	json.NewEncoder(w).Encode(nuevo)
}

// ///////////////////////////////////
// LOGIN
// ///////////////////////////////////
func login(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	var datos Usuario
	json.NewDecoder(r.Body).Decode(&datos)

	for i := range plataforma.Usuarios {
		if plataforma.Usuarios[i].Correo == datos.Correo &&
			plataforma.Usuarios[i].Password == datos.Password {

			plataforma.Logueado = &plataforma.Usuarios[i]
			json.NewEncoder(w).Encode("Login exitoso")
			return
		}
	}

	http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
}

// ///////////////////////////////////
// LOGOUT
// ///////////////////////////////////
func logout(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	plataforma.Logueado = nil
	json.NewEncoder(w).Encode("Sesión cerrada")
}

// ///////////////////////////////////
// LISTAR VIDEOS
// ///////////////////////////////////
func listarVideos(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	json.NewEncoder(w).Encode(plataforma.Videos)
}

// ///////////////////////////////////
// OBTENER VIDEO POR ID
// ///////////////////////////////////
func obtenerVideo(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, v := range plataforma.Videos {
		if v.ID == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	http.Error(w, "Video no encontrado", http.StatusNotFound)
}

// ///////////////////////////////////
// CREAR VIDEO
// ///////////////////////////////////
func crearVideo(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	var nuevo Video
	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plataforma.Videos = append(plataforma.Videos, nuevo)
	json.NewEncoder(w).Encode(nuevo)
}

// ///////////////////////////////////
// ACTUALIZAR VIDEO
// ///////////////////////////////////
func actualizarVideo(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var actualizado Video
	json.NewDecoder(r.Body).Decode(&actualizado)

	for i := range plataforma.Videos {
		if plataforma.Videos[i].ID == id {
			plataforma.Videos[i].Titulo = actualizado.Titulo
			json.NewEncoder(w).Encode(plataforma.Videos[i])
			return
		}
	}

	http.Error(w, "Video no encontrado", http.StatusNotFound)
}

// ///////////////////////////////////
// ELIMINAR VIDEO
// ///////////////////////////////////
func eliminarVideo(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, v := range plataforma.Videos {
		if v.ID == id {
			plataforma.Videos = append(plataforma.Videos[:i], plataforma.Videos[i+1:]...)
			json.NewEncoder(w).Encode("Video eliminado")
			return
		}
	}

	http.Error(w, "Video no encontrado", http.StatusNotFound)
}

// ///////////////////////////////////
// MAIN
// ///////////////////////////////////
func main() {

	plataforma.Videos = append(plataforma.Videos,
		Video{ID: 1, Titulo: "Interestelar"},
		Video{ID: 2, Titulo: "Matrix"},
		Video{ID: 3, Titulo: "El Padrino"},
	)

	http.HandleFunc("/usuarios", registrarUsuario)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/videos", listarVideos)
	http.HandleFunc("/video", obtenerVideo)
	http.HandleFunc("/crear-video", crearVideo)
	http.HandleFunc("/actualizar-video", actualizarVideo)
	http.HandleFunc("/eliminar-video", eliminarVideo)
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Servidor ejecutándose en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
