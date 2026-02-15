package main

import (
	"fmt"
)

/*
   ESTRUCTURAS

*/

// Usuario representa a una persona registrada en la plataforma
type Usuario struct {
	Correo   string
	Nombre   string
	Password string
}

// Video representa una película o contenido disponible
type Video struct {
	ID     int
	Titulo string
}

// Plataforma representa el sistema de streaming
type Plataforma struct {
	Usuarios []Usuario
	Videos   []Video
	Logueado *Usuario
}

/*
==========================
   MÉTODOS DEL SISTEMA
==========================
*/

// Registrar nuevo usuario
func (p *Plataforma) Registrar(correo, nombre, password string) {

	for _, u := range p.Usuarios {
		if u.Correo == correo {
			fmt.Println(" El correo ya está registrado.")
			return
		}
	}

	p.Usuarios = append(p.Usuarios, Usuario{
		Correo:   correo,
		Nombre:   nombre,
		Password: password,
	})

	fmt.Println(" Usuario registrado correctamente.")
}

// Iniciar sesión
func (p *Plataforma) Login(correo, password string) {

	for i := range p.Usuarios {
		if p.Usuarios[i].Correo == correo && p.Usuarios[i].Password == password {
			p.Logueado = &p.Usuarios[i]
			fmt.Println(" Bienvenido,", p.Logueado.Nombre)
			return
		}
	}

	fmt.Println(" Credenciales incorrectas.")
}

// Ver catálogo
func (p *Plataforma) VerCatalogo() {

	if p.Logueado == nil {
		fmt.Println(" Debes iniciar sesión primero.")
		return
	}

	fmt.Println("\n Catálogo disponible:")
	for _, v := range p.Videos {
		fmt.Printf("ID: %d | %s\n", v.ID, v.Titulo)
	}
}

// Reproducir película
func (p *Plataforma) Reproducir(id int) {

	if p.Logueado == nil {
		fmt.Println(" Debes iniciar sesión primero.")
		return
	}

	for _, v := range p.Videos {
		if v.ID == id {
			fmt.Println(" Reproduciendo:", v.Titulo)
			return
		}
	}

	fmt.Println(" Película no encontrada.")
}

// Cerrar sesión
func (p *Plataforma) Logout() {
	if p.Logueado != nil {
		fmt.Println(" Sesión cerrada.")
		p.Logueado = nil
	} else {
		fmt.Println(" No hay sesión activa.")
	}
}

/*
==========================
   FUNCIÓN PRINCIPAL
==========================
*/

func main() {

	plataforma := Plataforma{}

	// Catálogo inicial
	plataforma.Videos = append(plataforma.Videos,
		Video{ID: 1, Titulo: "Interestelar"},
		Video{ID: 2, Titulo: "Matrix"},
		Video{ID: 3, Titulo: "El Padrino"},
	)

	var opcion int

	for {
		fmt.Println("\n===== DMPLAYS - Streaming =====")
		fmt.Println("1. Registrar")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Ver catálogo")
		fmt.Println("4. Reproducir película")
		fmt.Println("5. Cerrar sesión")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione opción: ")

		fmt.Scanln(&opcion)

		switch opcion {

		case 1:
			var correo, nombre, pass string

			fmt.Print("Correo: ")
			fmt.Scanln(&correo)
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Contraseña: ")
			fmt.Scanln(&pass)

			plataforma.Registrar(correo, nombre, pass)

		case 2:
			var correo, pass string

			fmt.Print("Correo: ")
			fmt.Scanln(&correo)
			fmt.Print("Contraseña: ")
			fmt.Scanln(&pass)

			plataforma.Login(correo, pass)

		case 3:
			plataforma.VerCatalogo()

		case 4:
			var id int
			fmt.Print("ID de la película: ")
			fmt.Scanln(&id)
			plataforma.Reproducir(id)

		case 5:
			plataforma.Logout()

		case 6:
			fmt.Println(" Cerrando aplicación...")
			return

		default:
			fmt.Println(" Opción inválida.")
		}
	}
}

