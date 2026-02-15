package main

import "fmt"

func main() {
	// 1. Cara standar
	var nama string = "Stark"
	var umur int = 20

	//2. The Go Way
	tinggi := 174.5

	// 3. Konsep "Zero Value"
	var isStudent bool
	var saldo int

	fmt.Println("Nama : ", nama)
	fmt.Println("Umur : ", umur)
	fmt.Println("Tinggi (Tb) : ", tinggi)
	fmt.Println("Status Mahasiswa : ", isStudent)
	fmt.Println("Sisa Uang : ", saldo)

	memori()
	looping()
}

func memori() {
	// 1. ARRAY
	var loker [3]string
	loker[0] = "Buku"
	loker[1] = "Tas"

	// 2.Slice
	var playlist []string //ini kosong

	//tambahkan dengan append()
	playlist = append(playlist, "Mixed songs")
	playlist = append(playlist, "Instrumental")
	playlist = append(playlist, "Lagu Anime")

	fmt.Println("Isi Array Loker: ", loker)
	fmt.Println("Isi Slice Playlist: ", playlist)

	//Hitung banyak / panjang data
	fmt.Println("Jumlah Loker: ", len(loker))
	fmt.Println("Jumlah Playlist: ", len(playlist))
}

func looping() {
	// kita punya playlist musik (Slice)
	playlist := []string{"Mixed songs", "Instrumental", "Lagu Anime"}

	fmt.Println("--- PLAYLIST HARI")

	// 1. Cara Classic Loop
	for i := 0; i < len(playlist); i++ {
		fmt.Println("Putar (klasik):", playlist[i])
	}

	// 2. Range (The Go way)
	//range mengembalikan 2 hal : Index(urutan) dan Value(isinya)
	for index, lagu := range playlist {
		fmt.Printf("Lagu ke -%d: %s\n", index+1, lagu)
	}

	// 3. Mengabaikan index (pakai Underscore _)
	for _, lagu := range playlist {
		fmt.Println("sedang memutar: ", lagu)
	}

}
