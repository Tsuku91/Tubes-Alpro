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

	// Soal A
	questionBank[0].id = 1
	questionBank[0].soal = "Apa Ibu Kota Indonesia"
	questionBank[0].pilihan[0] = "A. Jakarta"
	questionBank[0].pilihan[1] = "B. Bandung"

	// Soal B
	questionBank[1].id = 2
	questionBank[1].soal = "1 + 1 ="
	questionBank[1].pilihan[0] = "A. 2"
	questionBank[1].pilihan[1] = "B. 1"

	// Soal C
	questionBank[2].id = 3
	questionBank[2].soal = "Kemerdekaan Indonsia Pada Tanggal Berapa?"
	questionBank[2].pilihan[0] = "A. 17  August 1945"
	questionBank[2].pilihan[1] = "A. 01  July 1945"

	// Soal D
	questionBank[3].id = 4
	questionBank[3].soal = "The book is (on / in) the table."
	questionBank[3].pilihan[0] = "A. On"
	questionBank[3].pilihan[1] = "A. In"

	// Soal E
	questionBank[4].id = 5
	questionBank[4].soal = "I am not (feeling / feel) well today."
	questionBank[4].pilihan[0] = "A. Feeling"
	questionBank[4].pilihan[1] = "B. Feel"

	// Soal F
	questionBank[5].id = 6
	questionBank[5].soal = "Pilih kalimat yang menggunakan present perfect tense dengan benar."
	questionBank[5].pilihan[0] = "A. I ate dinner already."
	questionBank[5].pilihan[1] = "B. I have eaten dinner already."

	// Soal G
	questionBank[6].id = 7
	questionBank[6].soal = "Pilih frasa yang tepat untuk meminta tolong."
	questionBank[6].pilihan[0] = "A. Tell me what to do."
	questionBank[6].pilihan[1] = "B. Can you help me, please?"

	// Soal H
	questionBank[7].id = 8
	questionBank[7].soal = "Manakah kalimat yang menggunakan comparative adjective dengan benar?"
	questionBank[7].pilihan[0] = "A. This book is gooder than that one."
	questionBank[7].pilihan[1] = "B. This book is better than that one."

	// Soal I
	questionBank[8].id = 9
	questionBank[8].soal = "It is raining outside, so I (will bring / am bringing) an umbrella."
	questionBank[8].pilihan[0] = "A. will bring"
	questionBank[8].pilihan[1] = "B. am bringing"

	// Soal J
	questionBank[9].id = 10
	questionBank[9].soal = "What is the correct superlative form of 'good'?"
	questionBank[9].pilihan[0] = "A. Gooder"
	questionBank[9].pilihan[1] = "B. Best"

	// Soal K
	questionBank[10].id = 11
	questionBank[10].soal = "Which is the correct past tense of 'eat'?"
	questionBank[10].pilihan[0] = "A. Ate"
	questionBank[10].pilihan[1] = "B. Eaten"

}

// MENGERJAKAN SOAL

// TINGKAT SKOR KESELURUHAN

func main() {
	mainMenu()
}
