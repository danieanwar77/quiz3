package auth

import (
	"quiz3/database"
	repository "quiz3/repositories"
	"quiz3/structs"
)

func Auth(username string, password string) bool {
	var (
		user structs.Users
	)

	user.UserName = username
	user.Password = password

	users, err := repository.AuthUser(database.DbConnection, user)

	// fmt.Println(http.StatusOK, users)

	if err != nil {
		return false
	} else {
		if users.ID != 0 {
			return true
		} else {
			return false
		}
	}

}

// func Check() {
// 	// konfigurasi server
// 	server := &http.Server{
// 		Addr: ":8000",
// 	}

// 	// routing
// 	http.Handle("/", Auth(http.HandlerFunc(CheckAuth)))

// 	// jalankan server
// 	fmt.Println("server running at http://localhost:800")
// 	err := server.ListenAndServe()
// 	fmt.Println("error server", err)
// }

// // Fungi Log yang berguna sebagai middleware
// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		uname, pwd, ok := r.BasicAuth()
// 		if !ok {
// 			w.Write([]byte("Username atau Password tidak boleh kosong"))
// 			return
// 		}

// 		var auth = controllers.AuthUser(uname, pwd)

// 		if auth == true {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		w.Write([]byte("Username atau Password tidak sesuai"))
// 		return
// 	})
// }

// // Fungsi GetMovie untuk mengampilkan text string di browser
// func CheckAuth(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		w.Write([]byte("<h1>Anda Berhasil Login </h1>"))
// 	}
// }
