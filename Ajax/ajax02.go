package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type book struct {
	Titulo string `json: titulo`
	Autor  string `json: autor`
}

type UserData struct {
	User    string
	Mail    string
	Lastcon string
}

type SendUser struct {
	User    bool
	Mail    bool
	Lastcon bool
}

var data = [4]UserData{
	UserData{"Nick", "nicks@anymail.com", "02-28-2021 09:38:22"},
	UserData{"Luis", "luis.eds@anymail.com", "02-27-2021 19:12:02"},
	UserData{"Chris", "chris@anymail.com", "02-25-2021 12:08:12"},
	UserData{"Oscar", "the_odb@anymail.com", "02-20-2021 17:58:09"},
}

var changed = [4]SendUser{
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
	SendUser{false, false, false},
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "ajax02.html")
}

func darMensaje(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	x := r.Form.Get("y")

	fmt.Printf("%s", x)

	libro := book{
		Titulo: "La Casa",
		Autor:  "Paco Roca",
	}

	//w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(libro)

}

func mostrarHTML2(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "reto1.html")
}

func mostrarHTML3(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "reto2.html")
}

func mostrarHTML4(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "reto3.html")
}

func reto1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	a := [5]int{1, 2, 3, 4, 5}

	switch r.Method {
	case "GET":
		for i := 0; i < 5; i++ {
			fmt.Fprintf(w, strconv.Itoa(a[i]))
		}
	case "POST":
		position := r.Form.Get("y")
		value := r.Form.Get("n")
		i, _ := strconv.Atoi(value)
		j, _ := strconv.Atoi(position)
		a[j%5] = i
		fmt.Fprintf(w, "Ok")
	}
}

func reto2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	position := r.Form.Get("y")

	a := [6]string{"{\"Persona\":\"Nick\",\"Libro\":\"Orgullo y Prejuicio\",\"mascota\":\"Manchas\"}",
		"{\"Persona\":\"Luis\",\"Libro\":\"Harry Potter\",\"mascota\":\"Sissi\"}",
		"{\"Persona\":\"Chris\",\"Libro\":\"Quiubole con\",\"mascota\":\"Nina\"}",
		"{\"Persona\":\"Oscar\",\"Libro\":\"Sapiens\",\"mascota\":\"Pirata\"}"}

	i, _ := strconv.Atoi(position)
	fmt.Fprint(w, a[i%4])

}

func reto3(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		toSend := [4]UserData{} // Informacion que va a enviar realmente
		for i, v := range changed {
			var n, m, d string

			// Guardará los datos que no esten bloqueados en las variables "n, m, d"
			if !v.User {
				n = data[i].User
				v.User = true
			} else if v.User {
				n = ""
			}

			if !v.Mail {
				m = data[i].Mail
				v.Mail = true
			} else if v.Mail {
				m = ""
			}

			if !v.Lastcon {
				d = data[i].Lastcon
				v.Lastcon = true
			} else if v.Lastcon {
				d = ""
			}

			// Actualiza los permisos modificados a la variable de permisos
			changed[i] = v
			// Crea el objeto a enviar con la informacion permitida
			toSend[i] = UserData{n, m, d}

		}
		json.NewEncoder(w).Encode(toSend)
	} else if r.Method == "POST" {

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()

		id, _ := strconv.Atoi(r.Form.Get("id"))
		name := r.Form.Get("name")
		mail := r.Form.Get("mail")
		date, _ := strconv.ParseBool(r.Form.Get("date"))

		fmt.Println("Updating data...")

		// Actualiza la informacion de "data" solo cuando hay algo

		if name != "" {
			fmt.Println("name: " + data[id].User + "->" + name)
			data[id].User = name
			changed[id].User = false
		}
		if mail != "" {
			fmt.Println("mail: " + data[id].Mail + "->" + mail)
			data[id].Mail = mail
			changed[id].Mail = false
		}
		if date {
			newDate := time.Now().Format("02-01-2006 15:04:05")
			fmt.Println("date: " + data[id].Lastcon + "->" + newDate)
			data[id].Lastcon = newDate
			changed[id].Lastcon = false
		}
	}

}

func main() {

	http.HandleFunc("/dato", darMensaje)
	http.HandleFunc("/", mostrarHTML)
	http.HandleFunc("/reto1", reto1)
	http.HandleFunc("/reto2", reto2)
	http.HandleFunc("/reto3", reto3)
	http.HandleFunc("/ejemplo", mostrarHTML2)
	http.HandleFunc("/ejemplo2", mostrarHTML3)
	http.HandleFunc("/ejemplo3", mostrarHTML4)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
