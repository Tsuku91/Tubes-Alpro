package main

import "fmt"

// STRUCT ADMIN
type admin struct {
	username, password string
}

// STRUCT USER
type user struct {
	username, kelas, level string
}

// STRUCT SOAL
type Question struct {
	id      int
	soal    string
	pilihan [2]string
}

var questionBank []Question

var peserta []user

func mainMenu() {
	var pilihan int
	bankSoal()
	for {
		fmt.Println("-Selamat Datang di Aplikasi Quiz Language Mastery-")
		fmt.Println("Silahkan Login atau Register terlebih dahulu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println("Login Sebagai :")
			fmt.Println("1. Admin")
			// fmt.Println("2. Peserta")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				loginAdmin()
			} else {
				// loginPeserta()
			}
		} else if pilihan == 2 {
			register()
		} else {
			fmt.Println("Terima Kasih Telah Berkunjung")
			break
		}
	}
}

// ADMIN
func loginAdmin() {
	var a admin
	fmt.Println("Masukkan Username :")
	fmt.Scan(&a.username)
	fmt.Println("Masukkan Password :")
	fmt.Scan(&a.password)
	if a.username == "admin" && a.password == "admin123" {
		fmt.Println("Login Sebagai Admin Berhasil")
		menuAdmin()
	}
}
func menuAdmin() {
	var pilihan int
	fmt.Println("Apa yang Ingin di Lakukan?")
	fmt.Println("1. Edit Soal")
	fmt.Println("2. Tambah Soal")
	fmt.Println("3. Hapus Soal")
	fmt.Println("4. Melihat Bank Soal")
	fmt.Println("5. Exit To Main Menu")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		editQuestion()
	} else if pilihan == 2 {
		addQuestion()
	} else if pilihan == 3 {
		deleteQuestion()
	} else if pilihan == 4 {
		bankSoal()
		melihatBankSoal()
	} else {
		mainMenu()
	}
}

func addQuestion() {
	var z Question
	var keputusan int
	for {
		fmt.Println("Masukan Nomor Soal :")
		fmt.Scan(&z.id)
		fmt.Println("Masukan Soal :")
		fmt.Scan(&z.soal)
		fmt.Println("Masukkan Opsi Jawaban A :")
		fmt.Scan(&z.pilihan[0])
		fmt.Println("Masukkan Opsi Jawaban B :")
		fmt.Scan(&z.pilihan[1])
		storeQuestion(z)
		fmt.Println("Ingin Menambahkan Soal Lagi ?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&keputusan)
		if keputusan == 1 {
			addQuestion()
		} else {
			break
		}
	}
}
func deleteQuestion() {

}
func editQuestion() {

}

func melihatBankSoal() {
	for _, question := range questionBank {
		fmt.Println("ID:", question.id)
		fmt.Println("Soal:", question.soal)
		for i, pilihan := range question.pilihan {
			fmt.Println("Pilihan", i+1, ":", pilihan)
		}
		fmt.Println()
	}
}

// PESERTA
func register() {
	var regist user

	fmt.Println("Silahkan masukkan username yang Ingin di gunakan :")
	fmt.Scan(&regist.username)

	fmt.Println("Silahkan masukkan kelas yang sedang kamu duduki :")
	fmt.Scan(&regist.kelas)

	fmt.Println("Silahkan masukkan level : (Beginner, Intermediate, Advanced)")
	fmt.Scan(&regist.level)

	peserta = append(peserta, regist) // Append new user to the slice
	menuPeserta()
	// loginPeserta()
}

// func loginPeserta(){
// 	fmt.Print("Masukan")

// 	fmt.Print("Login Berhasil")
// 	menuPeserta()
// }
func menuPeserta() {
	var pilihan int
	fmt.Println("Apa yang ingin dilakukan :")
	fmt.Println("1. Kerjakan Kuis")
	fmt.Println("2. Lihat Skor Keselurahan Peserta")
	fmt.Println("3. Exit To Main Menu")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		// kerjakanQuiz()
	} else if pilihan == 2 {
		// leaderboard()
	} else {
		mainMenu()
	}
}

// BANK SOAL
func storeQuestion(q Question) {
	questionBank = append(questionBank, q)
}

func bankSoal() {
	var a, b, c, d, e, f, g, h, i, j, k Question

	// Soal A
	a.id = 1
	a.soal = "Apa Ibu Kota Indonesia"
	a.pilihan[0] = "A. Jakarta"
	a.pilihan[1] = "B. Bandung"

	// Soal B
	b.id = 2
	b.soal = "1 + 1 ="
	b.pilihan[0] = "A. 2"
	b.pilihan[1] = "B. 1"

	// Soal C
	c.id = 3
	c.soal = "Kemerdekaan Indonsia Pada Tanggal Berapa?"
	c.pilihan[0] = "A. 17  August 1945"
	c.pilihan[1] = "A. 01  July 1945"

	// Soal D
	d.id = 4
	d.soal = "The book is (on / in) the table."
	d.pilihan[0] = "A. On"
	d.pilihan[1] = "A. In"

	// Soal E
	e.id = 5
	e.soal = "I am not (feeling / feel) well today."
	e.pilihan[0] = "A. Feeling"
	e.pilihan[1] = "B. Feel"

	// Soal F
	f.id = 6
	f.soal = "Pilih kalimat yang menggunakan present perfect tense dengan benar."
	f.pilihan[0] = "A. I ate dinner already."
	f.pilihan[1] = "B. I have eaten dinner already."

	// Soal G
	g.id = 7
	g.soal = "Pilih frasa yang tepat untuk meminta tolong."
	g.pilihan[0] = "A. Tell me what to do."
	g.pilihan[1] = "B. Can you help me, please?"

	// Soal H
	h.id = 8
	h.soal = "Manakah kalimat yang menggunakan comparative adjective dengan benar?"
	h.pilihan[0] = "A. This book is gooder than that one."
	h.pilihan[1] = "B. This book is better than that one."

	// Soal I
	i.id = 9
	i.soal = "It is raining outside, so I (will bring / am bringing) an umbrella."
	i.pilihan[0] = "A. will bring"
	i.pilihan[1] = "B. am bringing"

	// Soal J
	j.id = 10
	j.soal = "What is the correct superlative form of 'good'?"
	j.pilihan[0] = "A. Gooder"
	j.pilihan[1] = "B. Best"

	// Soal K
	k.id = 11
	k.soal = "Which is the correct past tense of 'eat'?"
	k.pilihan[0] = "A. Ate"
	k.pilihan[1] = "B. Eaten"

	questionBank = append(questionBank, a, b, c, d, e, f, g, h, i, j, k)
}

// MENGERJAKAN SOAL

// TINGKAT SKOR KESELURUHAN

func main() {
	mainMenu()
}
