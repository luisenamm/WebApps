package main

import (
  "net/http"
  "encoding/json"


)

var( data = make([]human, 0))

type human struct{
  Name string `json: name`
  Hobby string  `json: hobby`
  Music string `json: music`
}

func MostrarHTML(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "vue_abril2021.html")
}

func GetHumans(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{	
		json.NewEncoder(w).Encode(data)
	  }else{
		w.WriteHeader(http.StatusBadRequest)
	  }
 }
func AddHuman(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{	
		r.ParseForm() 
		name := r.Form.Get("name")
    	hobby := r.Form.Get("hobby")
		music := r.Form.Get("music")
		exist := false 
		for _, element := range data {
			if name == element.Name{
				exist = true
				break
			}
		}
		if exist{
			w.WriteHeader(http.StatusConflict)
		}else{
			data=append(data, human{name,hobby,music})
			w.WriteHeader(http.StatusCreated)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateHuman(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT"{	
		r.ParseForm() 
		name := r.Form.Get("name")
    	hobby := r.Form.Get("hobby")
		music := r.Form.Get("music")
		exist := false 
		for index, element := range data {
			if name == element.Name{
				data[index]=human{name, hobby, music}
				exist = true
				break
			}
		}
		if exist{
			w.WriteHeader(http.StatusNoContent)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
}
func PatchHobby(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH"{	
		r.ParseForm() 
		name := r.Form.Get("name")
    	hobby := r.Form.Get("hobby")
		exist := false 
		for index, element := range data {
			if name == element.Name{
				data[index].Hobby=hobby
				exist = true
				break
			}
		}
		if exist{
			w.WriteHeader(http.StatusNoContent)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
}
func PatchMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH"{	
		r.ParseForm() 
		name := r.Form.Get("name")
    	music := r.Form.Get("music")
		exist := false 
		for index, element := range data {
			if name == element.Name{
				data[index].Music=music
				exist = true
				break
			}
		}
		if exist{
			w.WriteHeader(http.StatusNoContent)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
}
func RemoveIndex(s []human, index int) []human {
    return append(s[:index], s[index+1:]...)
}
func DeleteHuman(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE"{	
		r.ParseForm() 
		name := r.Form.Get("name")
		exist := false 
		for index, element := range data {
			if name == element.Name{
				data=RemoveIndex(data,index)
				exist = true
				break
			}
		}
		if exist{
			w.WriteHeader(http.StatusOK)
		}else{
			w.WriteHeader(http.StatusNotFound)
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
}





func main() {

	human1:=human{"John Doe", "Photography", "Classic"}
	human2:=human{"Lorena Sumip", "Painting", "Soul"}
	human3:=human{"Bob Racso", "Music", "Blues"}

	data=append(data, human1)
	data=append(data, human2)
	data=append(data, human3)

	

  http.HandleFunc("/",MostrarHTML)
  http.HandleFunc("/GetHumans",GetHumans)
  http.HandleFunc("/AddHuman",AddHuman)
  http.HandleFunc("/UpdateHuman",UpdateHuman)
  http.HandleFunc("/PatchHobby",PatchHobby)
  http.HandleFunc("/PatchMusic",PatchMusic)
  http.HandleFunc("/DeleteHuman",DeleteHuman)

  
  err := http.ListenAndServe("localhost"+":"+"8080", nil)
  if err != nil {
    return
  }
}