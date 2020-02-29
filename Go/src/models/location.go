package models

import (
	"Connection"
	"fmt"
	"time"
)

type Location struct{
	Id 				int		`gorm:"PRIMARY_KEY"`
	City			string 	`gorm:"VARCHAR(MAX)"`
	Province		string	`gorm:"VARCHAR(MAX)"`
	Country 		string 	`gorm:"VARCHAR(MAX)"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time	`sql : "index"`
}
func init(){
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&Location{})

}
func InsertLocation(city string, province string, country string) *Location{
	db:= Connection.Connect()
	defer db.Close()

	newLoc := &Location{
		City: city,
		Province: province,
		Country: country,
	}

	db.Save(newLoc)

	fmt.Println(newLoc,"Inserted")

	return newLoc
}
func InLoc(){
	InsertLocation("Banda Aceh", "Aceh","Indonesia")
	InsertLocation("Sabang", "Aceh","Indonesia")
	InsertLocation("Denpasar", "Bali","Indonesia")
	InsertLocation("Singaraja", "Bali","Indonesia")
	InsertLocation("Kuta", "Bali","Indonesia")
	InsertLocation("Pangkalpinang", "Bangka Belitung", "Indonesia")
	InsertLocation("Bengkulu", "Bengkulu","Indonesia")
	InsertLocation("Magelang", "Jawa Tengah","Indonesia")
	InsertLocation("Pekalongan", "Jawa Tengah", "Indonesia")
	InsertLocation("Rembang", "Jawa Tengah","Indonesia")
	InsertLocation("Salatiga", "Jawa Tengah","Indonesia")
	InsertLocation("Semarang", "Jawa Tengah","Indonesia")
	InsertLocation("Tegal", "Jawa Tengah","Indonesia")
	InsertLocation("Palu", "Sulawesi Tengah","Indonesia")
	InsertLocation("Banyuwangi", "Jawa Timur","Indonesia")
	InsertLocation("Blitar", "Jawa Timur","Indonesia")
	InsertLocation("Jember", "Jawa Timur","Indonesia")
	InsertLocation("Aceh Barat","Aceh","Indonesia")
	InsertLocation("Aceh Barat Daya","Aceh","Indonesia")
	InsertLocation("Aceh Besar","Aceh","Indonesia")
	InsertLocation("Aceh Jaya","Aceh","Indonesia")
	InsertLocation("Aceh Selatan","Aceh","Indonesia")
	InsertLocation("Aceh Singkil","Aceh","Indonesia")
	InsertLocation("Aceh Tamiang","Aceh","Indonesia")
	InsertLocation("Aceh Tengah","Aceh","Indonesia")
	InsertLocation("Aceh Tenggara","Aceh","Indonesia")
	InsertLocation("Aceh Timur","Aceh","Indonesia")
	InsertLocation("Aceh Utara","Aceh","Indonesia")
	InsertLocation("Administrasi Kepulauan Seribu","Jakarta","Indonesia")
	InsertLocation("Agam","Sumatera Barat","Indonesia")
	InsertLocation("Alor","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Asahan","Sumatera Utara","Indonesia")
	InsertLocation("Asmat","Papua","Indonesia")
	InsertLocation("Badung","Bali","Indonesia")
	InsertLocation("Balangan","Kalimantan Selatan","Indonesia")
	InsertLocation("Bandung","Jawa Barat","Indonesia")
	InsertLocation("Bandung Barat","Jawa Barat","Indonesia")
	InsertLocation("Banggai","Sulawesi Tengah","Indonesia")
	InsertLocation("Banggai Kepulauan","Sulawesi Tengah","Indonesia")
	InsertLocation("Bangka","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Bangka Barat","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Bangka Selatan","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Bangka Tengah","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Bangkalan","Jawa Timur","Indonesia")
	InsertLocation("Bangli","Bali","Indonesia")
	InsertLocation("Banjar","Kalimantan Selatan","Indonesia")
	InsertLocation("Banjarnegara","Jawa Tengah","Indonesia")
	InsertLocation("Bantaeng","Sulawesi Selatan","Indonesia")
	InsertLocation("Bantul","Yogyakarta","Indonesia")
	InsertLocation("Banyuasin","Sumatera Selatan","Indonesia")
	InsertLocation("Banyumas","Jawa Tengah","Indonesia")
	InsertLocation("Banyuwangi","Jawa Timur","Indonesia")
	InsertLocation("Barito Kuala","Kalimantan Selatan","Indonesia")
	InsertLocation("Barito Selatan","Kalimantan Tengah","Indonesia")
	InsertLocation("Barito Timur","Kalimantan Tengah","Indonesia")
	InsertLocation("Barito Utara","Kalimantan Tengah","Indonesia")
	InsertLocation("Barru","Sulawesi Selatan","Indonesia")
	InsertLocation("Batang","Jawa Tengah","Indonesia")
	InsertLocation("Batanghari","Jambi","Indonesia")
	InsertLocation("Batubara","Sumatera Utara","Indonesia")
	InsertLocation("Bekasi","Jawa Barat","Indonesia")
	InsertLocation("Belitung","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Belitung Timur","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Belu","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Bener Meriah","Aceh","Indonesia")
	InsertLocation("Bengkalis","Riau","Indonesia")
	InsertLocation("Bengkayang","Kalimantan Barat","Indonesia")
	InsertLocation("Bengkulu Selatan","Bengkulu","Indonesia")
	InsertLocation("Bengkulu Tengah","Bengkulu","Indonesia")
	InsertLocation("Bengkulu Utara","Bengkulu","Indonesia")
	InsertLocation("Berau","Kalimantan Timur","Indonesia")
	InsertLocation("Biak Numfor","Papua","Indonesia")
	InsertLocation("Bima","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Bintan","Kepulauan Riau","Indonesia")
	InsertLocation("Bireuen","Aceh","Indonesia")
	InsertLocation("Blitar","Jawa Timur","Indonesia")
	InsertLocation("Blora","Jawa Tengah","Indonesia")
	InsertLocation("Boalemo","Gorontalo","Indonesia")
	InsertLocation("Bogor","Jawa Barat","Indonesia")
	InsertLocation("Bojonegoro","Jawa Timur","Indonesia")
	InsertLocation("Bolaang Mongondow","Sulawesi Utara","Indonesia")
	InsertLocation("Bolaang Mongondow Selatan","Sulawesi Utara","Indonesia")
	InsertLocation("Bolaang Mongondow Timur","Sulawesi Utara","Indonesia")
	InsertLocation("Bolaang Mongondow Utara","Sulawesi Utara","Indonesia")
	InsertLocation("Bombana","Sulawesi Tenggara","Indonesia")
	InsertLocation("Bondowoso","Jawa Timur","Indonesia")
	InsertLocation("Bone","Sulawesi Selatan","Indonesia")
	InsertLocation("Bone Bolango","Gorontalo","Indonesia")
	InsertLocation("Boven Digoel","Papua","Indonesia")
	InsertLocation("Boyolali","Jawa Tengah","Indonesia")
	InsertLocation("Brebes","Jawa Tengah","Indonesia")
	InsertLocation("Buleleng","Bali","Indonesia")
	InsertLocation("Bulukumba","Sulawesi Selatan","Indonesia")
	InsertLocation("Bulungan","Kalimantan Utara","Indonesia")
	InsertLocation("Bungo","Jambi","Indonesia")
	InsertLocation("Buol","Sulawesi Tengah","Indonesia")
	InsertLocation("Buru","Maluku","Indonesia")
	InsertLocation("Buru Selatan","Maluku","Indonesia")
	InsertLocation("Buton","Sulawesi Tenggara","Indonesia")
	InsertLocation("Buton Utara","Sulawesi Tenggara","Indonesia")
	InsertLocation("Ciamis","Jawa Barat","Indonesia")
	InsertLocation("Cianjur","Jawa Barat","Indonesia")
	InsertLocation("Cilacap","Jawa Tengah","Indonesia")
	InsertLocation("Cirebon","Jawa Barat","Indonesia")
	InsertLocation("Dairi","Sumatera Utara","Indonesia")
	InsertLocation("Deiyai","Papua","Indonesia")
	InsertLocation("Deli Serdang","Sumatera Utara","Indonesia")
	InsertLocation("Demak","Jawa Tengah","Indonesia")
	InsertLocation("Dharmasraya","Sumatera Barat","Indonesia")
	InsertLocation("Dogiyai","Papua","Indonesia")
	InsertLocation("Dompu","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Donggala","Sulawesi Tengah","Indonesia")
	InsertLocation("Empat Lawang","Sumatera Selatan","Indonesia")
	InsertLocation("Ende","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Enrekang","Sulawesi Selatan","Indonesia")
	InsertLocation("Fakfak","Papua Barat","Indonesia")
	InsertLocation("Flores Timur","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Garut","Jawa Barat","Indonesia")
	InsertLocation("Gayo Lues","Aceh","Indonesia")
	InsertLocation("Gianyar","Bali","Indonesia")
	InsertLocation("Gorontalo","Gorontalo","Indonesia")
	InsertLocation("Gorontalo Utara","Gorontalo","Indonesia")
	InsertLocation("Gowa","Sulawesi Selatan","Indonesia")
	InsertLocation("Gresik","Jawa Timur","Indonesia")
	InsertLocation("Grobogan","Jawa Tengah","Indonesia")
	InsertLocation("Gunung Kidul","Yogyakarta","Indonesia")
	InsertLocation("Gunung Mas","Kalimantan Tengah","Indonesia")
	InsertLocation("Halmahera Barat","Maluku Utara","Indonesia")
	InsertLocation("Halmahera Selatan","Maluku Utara","Indonesia")
	InsertLocation("Halmahera Tengah","Maluku Utara","Indonesia")
	InsertLocation("Halmahera Timur","Maluku Utara","Indonesia")
	InsertLocation("Halmahera Utara","Maluku Utara","Indonesia")
	InsertLocation("Hulu Sungai Selatan","Kalimantan Selatan","Indonesia")
	InsertLocation("Hulu Sungai Tengah","Kalimantan Selatan","Indonesia")
	InsertLocation("Hulu Sungai Utara","Kalimantan Selatan","Indonesia")
	InsertLocation("Humbang Hasundutan","Sumatera Utara","Indonesia")
	InsertLocation("Indragiri Hilir","Riau","Indonesia")
	InsertLocation("Indragiri Hulu","Riau","Indonesia")
	InsertLocation("Indramayu","Jawa Barat","Indonesia")
	InsertLocation("Intan Jaya","Papua","Indonesia")
	InsertLocation("Jayapura","Papua","Indonesia")
	InsertLocation("Jayawijaya","Papua","Indonesia")
	InsertLocation("Jember","Jawa Timur","Indonesia")
	InsertLocation("Jembrana","Bali","Indonesia")
	InsertLocation("Jeneponto","Sulawesi Selatan","Indonesia")
	InsertLocation("Jepara","Jawa Tengah","Indonesia")
	InsertLocation("Jombang","Jawa Timur","Indonesia")
	InsertLocation("Kaimana","Papua Barat","Indonesia")
	InsertLocation("Kampar","Riau","Indonesia")
	InsertLocation("Kapuas","Kalimantan Tengah","Indonesia")
	InsertLocation("Kapuas Hulu","Kalimantan Barat","Indonesia")
	InsertLocation("Karanganyar","Jawa Tengah","Indonesia")
	InsertLocation("Karangasem","Bali","Indonesia")
	InsertLocation("Karawang","Jawa Barat","Indonesia")
	InsertLocation("Karimun","Kepulauan Riau","Indonesia")
	InsertLocation("Karo","Sumatera Utara","Indonesia")
	InsertLocation("Katingan","Kalimantan Tengah","Indonesia")
	InsertLocation("Kaur","Bengkulu","Indonesia")
	InsertLocation("Kayong Utara","Kalimantan Barat","Indonesia")
	InsertLocation("Kebumen","Jawa Tengah","Indonesia")
	InsertLocation("Kediri","Jawa Timur","Indonesia")
	InsertLocation("Keerom","Papua","Indonesia")
	InsertLocation("Kendal","Jawa Tengah","Indonesia")
	InsertLocation("Kepahiang","Bengkulu","Indonesia")
	InsertLocation("Kepulauan Anambas","Kepulauan Riau","Indonesia")
	InsertLocation("Kepulauan Aru","Maluku","Indonesia")
	InsertLocation("Kepulauan Mentawai","Sumatera Barat","Indonesia")
	InsertLocation("Kepulauan Meranti","Riau","Indonesia")
	InsertLocation("Kepulauan Sangihe","Sulawesi Utara","Indonesia")
	InsertLocation("Kepulauan Selayar","Sulawesi Selatan","Indonesia")
	InsertLocation("Kepulauan Siau Tagulandang Biaro","Sulawesi Utara","Indonesia")
	InsertLocation("Kepulauan Sula","Maluku Utara","Indonesia")
	InsertLocation("Kepulauan Talaud","Sulawesi Utara","Indonesia")
	InsertLocation("Kepulauan Yapen","Papua","Indonesia")
	InsertLocation("Kerinci","Jambi","Indonesia")
	InsertLocation("Ketapang","Kalimantan Barat","Indonesia")
	InsertLocation("Klaten","Jawa Tengah","Indonesia")
	InsertLocation("Klungkung","Bali","Indonesia")
	InsertLocation("Kolaka","Sulawesi Tenggara","Indonesia")
	InsertLocation("Kolaka Utara","Sulawesi Tenggara","Indonesia")
	InsertLocation("Konawe","Sulawesi Tenggara","Indonesia")
	InsertLocation("Konawe Selatan","Sulawesi Tenggara","Indonesia")
	InsertLocation("Konawe Utara","Sulawesi Tenggara","Indonesia")
	InsertLocation("Kotabaru","Kalimantan Selatan","Indonesia")
	InsertLocation("Kotawaringin Barat","Kalimantan Tengah","Indonesia")
	InsertLocation("Kotawaringin Timur","Kalimantan Tengah","Indonesia")
	InsertLocation("Kuantan Singingi","Riau","Indonesia")
	InsertLocation("Kubu Raya","Kalimantan Barat","Indonesia")
	InsertLocation("Kudus","Jawa Tengah","Indonesia")
	InsertLocation("Kulon Progo","Yogyakarta","Indonesia")
	InsertLocation("Kuningan","Jawa Barat","Indonesia")
	InsertLocation("Kupang","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Kutai Barat","Kalimantan Timur","Indonesia")
	InsertLocation("Kutai Kartanegara","Kalimantan Timur","Indonesia")
	InsertLocation("Kutai Timur","Kalimantan Timur","Indonesia")
	InsertLocation("Labuhanbatu","Sumatera Utara","Indonesia")
	InsertLocation("Labuhanbatu Selatan","Sumatera Utara","Indonesia")
	InsertLocation("Labuhanbatu Utara","Sumatera Utara","Indonesia")
	InsertLocation("Lahat","Sumatera Selatan","Indonesia")
	InsertLocation("Lamandau","Kalimantan Tengah","Indonesia")
	InsertLocation("Lamongan","Jawa Timur","Indonesia")
	InsertLocation("Lampung Barat","Lampung","Indonesia")
	InsertLocation("Lampung Selatan","Lampung","Indonesia")
	InsertLocation("Lampung Tengah","Lampung","Indonesia")
	InsertLocation("Lampung Timur","Lampung","Indonesia")
	InsertLocation("Lampung Utara","Lampung","Indonesia")
	InsertLocation("Landak","Kalimantan Barat","Indonesia")
	InsertLocation("Langkat","Sumatera Utara","Indonesia")
	InsertLocation("Lanny Jaya","Papua","Indonesia")
	InsertLocation("Lebak","Banten","Indonesia")
	InsertLocation("Lebong","Bengkulu","Indonesia")
	InsertLocation("Lembata","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Lima Puluh Kota","Sumatera Barat","Indonesia")
	InsertLocation("Lingga","Kepulauan Riau","Indonesia")
	InsertLocation("Lombok Barat","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Lombok Tengah","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Lombok Timur","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Lombok Utara","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Lumajang","Jawa Timur","Indonesia")
	InsertLocation("Luwu","Sulawesi Selatan","Indonesia")
	InsertLocation("Luwu Timur","Sulawesi Selatan","Indonesia")
	InsertLocation("Luwu Utara","Sulawesi Selatan","Indonesia")
	InsertLocation("Madiun","Jawa Timur","Indonesia")
	InsertLocation("Magelang","Jawa Tengah","Indonesia")
	InsertLocation("Mahakam Ulu","Kalimantan Timur","Indonesia")
	InsertLocation("Majalengka","Jawa Barat","Indonesia")
	InsertLocation("Majalengka","Jawa Barat","Indonesia")
	InsertLocation("Majene","Sulawesi Barat","Indonesia")
	InsertLocation("Malang","Jawa Timur","Indonesia")
	InsertLocation("Malinau","Kalimantan Utara","Indonesia")
	InsertLocation("Maluku Barat Daya","Maluku","Indonesia")
	InsertLocation("Maluku Tengah","Maluku","Indonesia")
	InsertLocation("Maluku Tenggara","Maluku","Indonesia")
	InsertLocation("Maluku Tenggara Barat","Maluku","Indonesia")
	InsertLocation("Mamasa","Sulawesi Barat","Indonesia")
	InsertLocation("Mamberamo Raya","Papua","Indonesia")
	InsertLocation("Mamberamo Tengah","Papua","Indonesia")
	InsertLocation("Mamuju","Sulawesi Barat","Indonesia")
	InsertLocation("Mamuju Utara","Sulawesi Barat","Indonesia")
	InsertLocation("Mandailing Natal","Sumatera Utara","Indonesia")
	InsertLocation("Manggarai","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Manggarai Barat","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Manggarai Timur","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Manokwari","Papua Barat","Indonesia")
	InsertLocation("Mappi","Papua","Indonesia")
	InsertLocation("Maros","Sulawesi Selatan","Indonesia")
	InsertLocation("Maybrat","Papua Barat","Indonesia")
	InsertLocation("Melawi","Kalimantan Barat","Indonesia")
	InsertLocation("Mempawah","Kalimantan Barat","Indonesia")
	InsertLocation("Merangin","Jambi","Indonesia")
	InsertLocation("Merauke","Papua","Indonesia")
	InsertLocation("Mesuji","Lampung","Indonesia")
	InsertLocation("Mimika","Papua","Indonesia")
	InsertLocation("Minahasa","Sulawesi Utara","Indonesia")
	InsertLocation("Minahasa Selatan","Sulawesi Utara","Indonesia")
	InsertLocation("Minahasa Tenggara","Sulawesi Utara","Indonesia")
	InsertLocation("Minahasa Utara","Sulawesi Utara","Indonesia")
	InsertLocation("Mojokerto","Jawa Timur","Indonesia")
	InsertLocation("Morowali","Sulawesi Tengah","Indonesia")
	InsertLocation("Muara Enim","Sumatera Selatan","Indonesia")
	InsertLocation("Muaro Jambi","Jambi","Indonesia")
	InsertLocation("Mukomuko","Bengkulu","Indonesia")
	InsertLocation("Muna","Sulawesi Tenggara","Indonesia")
	InsertLocation("Murung Raya","Kalimantan Tengah","Indonesia")
	InsertLocation("Musi Banyuasin","Sumatera Selatan","Indonesia")
	InsertLocation("Musi Rawas","Sumatera Selatan","Indonesia")
	InsertLocation("Nabire","Papua","Indonesia")
	InsertLocation("Nagan Raya","Aceh","Indonesia")
	InsertLocation("Nagekeo","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Natuna","Kepulauan Riau","Indonesia")
	InsertLocation("Nduga","Papua","Indonesia")
	InsertLocation("Ngada","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Nganjuk","Jawa Timur","Indonesia")
	InsertLocation("Ngawi","Jawa Timur","Indonesia")
	InsertLocation("Nias","Sumatera Utara","Indonesia")
	InsertLocation("Nias Barat","Sumatera Utara","Indonesia")
	InsertLocation("Nias Selatan","Sumatera Utara","Indonesia")
	InsertLocation("Nias Utara","Sumatera Utara","Indonesia")
	InsertLocation("Nunukan","Kalimantan Utara","Indonesia")
	InsertLocation("Ogan Ilir","Sumatera Selatan","Indonesia")
	InsertLocation("Ogan Komering Ilir","Sumatera Selatan","Indonesia")
	InsertLocation("Ogan Komering Ulu","Sumatera Selatan","Indonesia")
	InsertLocation("Ogan Komering Ulu Selatan","Sumatera Selatan","Indonesia")
	InsertLocation("Ogan Komering Ulu Timur","Sumatera Selatan","Indonesia")
	InsertLocation("Pacitan","Jawa Timur","Indonesia")
	InsertLocation("Padang Lawas","Sumatera Utara","Indonesia")
	InsertLocation("Padang Lawas Utara","Sumatera Utara","Indonesia")
	InsertLocation("Padang Pariaman","Sumatera Barat","Indonesia")
	InsertLocation("Pakpak Bharat","Sumatera Utara","Indonesia")
	InsertLocation("Pamekasan","Jawa Timur","Indonesia")
	InsertLocation("Pandeglang","Banten","Indonesia")
	InsertLocation("Pangkajene dan Kepulauan","Sulawesi Selatan","Indonesia")
	InsertLocation("Paniai","Papua","Indonesia")
	InsertLocation("Parigi Moutong","Sulawesi Tengah","Indonesia")
	InsertLocation("Pasaman","Sumatera Barat","Indonesia")
	InsertLocation("Pasaman Barat","Sumatera Barat","Indonesia")
	InsertLocation("Paser","Kalimantan Timur","Indonesia")
	InsertLocation("Pasuruan","Jawa Timur","Indonesia")
	InsertLocation("Pati","Jawa Tengah","Indonesia")
	InsertLocation("Pegunungan Bintang","Papua","Indonesia")
	InsertLocation("Pekalongan","Jawa Tengah","Indonesia")
	InsertLocation("Pelalawan","Riau","Indonesia")
	InsertLocation("Pemalang","Jawa Tengah","Indonesia")
	InsertLocation("Penajam Paser Utara","Kalimantan Timur","Indonesia")
	InsertLocation("Pesawaran","Lampung","Indonesia")
	InsertLocation("Pesisir Selatan","Sumatera Barat","Indonesia")
	InsertLocation("Pidie","Aceh","Indonesia")
	InsertLocation("Pidie Jaya","Aceh","Indonesia")
	InsertLocation("Pinrang","Sulawesi Selatan","Indonesia")
	InsertLocation("Pohuwato","Gorontalo","Indonesia")
	InsertLocation("Polewali Mandar","Sulawesi Barat","Indonesia")
	InsertLocation("Ponorogo","Jawa Timur","Indonesia")
	InsertLocation("Poso","Sulawesi Tengah","Indonesia")
	InsertLocation("Pringsewu","Lampung","Indonesia")
	InsertLocation("Probolinggo","Jawa Timur","Indonesia")
	InsertLocation("Pulang Pisau","Kalimantan Tengah","Indonesia")
	InsertLocation("Pulau Morotai","Maluku Utara","Indonesia")
	InsertLocation("Puncak","Papua","Indonesia")
	InsertLocation("Puncak Jaya","Papua","Indonesia")
	InsertLocation("Purbalingga","Jawa Tengah","Indonesia")
	InsertLocation("Purwakarta","Jawa Barat","Indonesia")
	InsertLocation("Purworejo","Jawa Tengah","Indonesia")
	InsertLocation("Raja Ampat","Papua Barat","Indonesia")
	InsertLocation("Rejang Lebong","Bengkulu","Indonesia")
	InsertLocation("Rembang","Jawa Tengah","Indonesia")
	InsertLocation("Rokan Hilir","Riau","Indonesia")
	InsertLocation("Rokan Hulu","Riau","Indonesia")
	InsertLocation("Rote Ndao","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sabu Raijua","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sambas","Kalimantan Barat","Indonesia")
	InsertLocation("Samosir","Sumatera Utara","Indonesia")
	InsertLocation("Sampang","Jawa Timur","Indonesia")
	InsertLocation("Sanggau","Kalimantan Barat","Indonesia")
	InsertLocation("Sarmi","Papua","Indonesia")
	InsertLocation("Sarolangun","Jambi","Indonesia")
	InsertLocation("Sekadau","Kalimantan Barat","Indonesia")
	InsertLocation("Seluma","Bengkulu","Indonesia")
	InsertLocation("Semarang","Jawa Tengah","Indonesia")
	InsertLocation("Seram Bagian Barat","Maluku","Indonesia")
	InsertLocation("Seram Bagian Timur","Maluku","Indonesia")
	InsertLocation("Serang","Banten","Indonesia")
	InsertLocation("Serdang Bedagai","Sumatera Utara","Indonesia")
	InsertLocation("Seruyan","Kalimantan Tengah","Indonesia")
	InsertLocation("Siak","Riau","Indonesia")
	InsertLocation("Sidenreng Rappang","Sulawesi Selatan","Indonesia")
	InsertLocation("Sidoarjo","Jawa Timur","Indonesia")
	InsertLocation("Sigi","Sulawesi Tengah","Indonesia")
	InsertLocation("Sijunjung","Sumatera Barat","Indonesia")
	InsertLocation("Sikka","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Simalungun","Sumatera Utara","Indonesia")
	InsertLocation("Simeulue","Aceh","Indonesia")
	InsertLocation("Sinjai","Sulawesi Selatan","Indonesia")
	InsertLocation("Sintang","Kalimantan Barat","Indonesia")
	InsertLocation("Situbondo","Jawa Timur","Indonesia")
	InsertLocation("Sleman","Yogyakarta","Indonesia")
	InsertLocation("Solok","Sumatera Barat","Indonesia")
	InsertLocation("Solok Selatan","Sumatera Barat","Indonesia")
	InsertLocation("Soppeng","Sulawesi Selatan","Indonesia")
	InsertLocation("Sorong","Papua Barat","Indonesia")
	InsertLocation("Sorong Selatan","Papua Barat","Indonesia")
	InsertLocation("Sragen","Jawa Tengah","Indonesia")
	InsertLocation("Subang","Jawa Barat","Indonesia")
	InsertLocation("Sukabumi","Jawa Barat","Indonesia")
	InsertLocation("Sukamara","Kalimantan Tengah","Indonesia")
	InsertLocation("Sukoharjo","Jawa Tengah","Indonesia")
	InsertLocation("Sumba Barat","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sumba Barat Daya","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sumba Tengah","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sumba Timur","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Sumbawa","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Sumbawa Barat","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Sumedang","Jawa Barat","Indonesia")
	InsertLocation("Sumenep","Jawa Timur","Indonesia")
	InsertLocation("Supiori","Papua","Indonesia")
	InsertLocation("Tabalong","Kalimantan Selatan","Indonesia")
	InsertLocation("Tabanan","Bali","Indonesia")
	InsertLocation("Takalar","Sulawesi Selatan","Indonesia")
	InsertLocation("Tambrauw","Papua Barat","Indonesia")
	InsertLocation("Tana Tidung","Kalimantan Utara","Indonesia")
	InsertLocation("Tana Toraja","Sulawesi Selatan","Indonesia")
	InsertLocation("Tanah Bumbu","Kalimantan Selatan","Indonesia")
	InsertLocation("Tanah Datar","Sumatera Barat","Indonesia")
	InsertLocation("Tanah Laut","Kalimantan Selatan","Indonesia")
	InsertLocation("Tangerang","Banten","Indonesia")
	InsertLocation("Tanggamus","Lampung","Indonesia")
	InsertLocation("Tanjung Jabung Barat","Jambi","Indonesia")
	InsertLocation("Tanjung Jabung Timur","Jambi","Indonesia")
	InsertLocation("Tapanuli Selatan","Sumatera Utara","Indonesia")
	InsertLocation("Tapanuli Tengah","Sumatera Utara","Indonesia")
	InsertLocation("Tapanuli Utara","Sumatera Utara","Indonesia")
	InsertLocation("Tapin","Kalimantan Selatan","Indonesia")
	InsertLocation("Tasikmalaya","Jawa Barat","Indonesia")
	InsertLocation("Tebo","Jambi","Indonesia")
	InsertLocation("Tegal","Jawa Tengah","Indonesia")
	InsertLocation("Teluk Bintuni","Papua Barat","Indonesia")
	InsertLocation("Teluk Wondama","Papua Barat","Indonesia")
	InsertLocation("Temanggung","Jawa Tengah","Indonesia")
	InsertLocation("Timor Tengah Selatan","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Timor Tengah Utara","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Toba Samosir","Sumatera Utara","Indonesia")
	InsertLocation("Tojo Una-Una","Sulawesi Tengah","Indonesia")
	InsertLocation("Toli-Toli","Sulawesi Tengah","Indonesia")
	InsertLocation("Tolikara","Papua","Indonesia")
	InsertLocation("Toraja Utara","Sulawesi Selatan","Indonesia")
	InsertLocation("Trenggalek","Jawa Timur","Indonesia")
	InsertLocation("Tuban","Jawa Timur","Indonesia")
	InsertLocation("Tulang Bawang","Lampung","Indonesia")
	InsertLocation("Tulang Bawang Barat","Lampung","Indonesia")
	InsertLocation("Tulungagung","Jawa Timur","Indonesia")
	InsertLocation("Wajo","Sulawesi Selatan","Indonesia")
	InsertLocation("Wakatobi","Sulawesi Tenggara","Indonesia")
	InsertLocation("Waropen","Papua","Indonesia")
	InsertLocation("Way Kanan","Lampung","Indonesia")
	InsertLocation("Wonogiri","Jawa Tengah","Indonesia")
	InsertLocation("Wonosobo","Jawa Tengah","Indonesia")
	InsertLocation("Yahukimo","Papua","Indonesia")
	InsertLocation("Yalimo","Papua","Indonesia")
	InsertLocation("Administrasi Jakarta Barat","Jakarta","Indonesia")
	InsertLocation("Administrasi Jakarta Pusat","Jakarta","Indonesia")
	InsertLocation("Administrasi Jakarta Selatan","Jakarta","Indonesia")
	InsertLocation("Administrasi Jakarta Timur","Jakarta","Indonesia")
	InsertLocation("Administrasi Jakarta Utara","Jakarta","Indonesia")
	InsertLocation("Ambon","Maluku","Indonesia")
	InsertLocation("Balikpapan","Kalimantan Timur","Indonesia")
	InsertLocation("Banda Aceh","Aceh","Indonesia")
	InsertLocation("Bandar Lampung","Lampung","Indonesia")
	InsertLocation("Bandung","Jawa Barat","Indonesia")
	InsertLocation("Banjar","Jawa Barat","Indonesia")
	InsertLocation("Banjarbaru","Kalimantan Selatan","Indonesia")
	InsertLocation("Banjarmasin","Kalimantan Selatan","Indonesia")
	InsertLocation("Batam","Kepulauan Riau","Indonesia")
	InsertLocation("Batu","Jawa Timur","Indonesia")
	InsertLocation("Bau-Bau","Sulawesi Tenggara","Indonesia")
	InsertLocation("Bekasi","Jawa Barat","Indonesia")
	InsertLocation("Bengkulu","Bengkulu","Indonesia")
	InsertLocation("Bima","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Binjai","Sumatera Utara","Indonesia")
	InsertLocation("Bitung","Sulawesi Utara","Indonesia")
	InsertLocation("Blitar","Jawa Timur","Indonesia")
	InsertLocation("Bogor","Jawa Barat","Indonesia")
	InsertLocation("Bontang","Kalimantan Timur","Indonesia")
	InsertLocation("Bukittinggi","Sumatera Barat","Indonesia")
	InsertLocation("Cilegon","Banten","Indonesia")
	InsertLocation("Cimahi","Jawa Barat","Indonesia")
	InsertLocation("Cirebon","Jawa Barat","Indonesia")
	InsertLocation("Denpasar","Bali","Indonesia")
	InsertLocation("Depok","Jawa Barat","Indonesia")
	InsertLocation("Dumai","Riau","Indonesia")
	InsertLocation("Gorontalo","Gorontalo","Indonesia")
	InsertLocation("Gunungsitoli","Sumatera Utara","Indonesia")
	InsertLocation("Jambi","Jambi","Indonesia")
	InsertLocation("Jayapura","Papua","Indonesia")
	InsertLocation("Kediri","Jawa Timur","Indonesia")
	InsertLocation("Kendari","Sulawesi Tenggara","Indonesia")
	InsertLocation("Kotamobagu","Sulawesi Utara","Indonesia")
	InsertLocation("Kupang","Nusa Tenggara Timur","Indonesia")
	InsertLocation("Langsa","Aceh","Indonesia")
	InsertLocation("Lhokseumawe","Aceh","Indonesia")
	InsertLocation("Lubuklinggau","Sumatera Selatan","Indonesia")
	InsertLocation("Madiun","Jawa Timur","Indonesia")
	InsertLocation("Magelang","Jawa Tengah","Indonesia")
	InsertLocation("Makassar","Sulawesi Selatan","Indonesia")
	InsertLocation("Malang","Jawa Timur","Indonesia")
	InsertLocation("Manado","Sulawesi Utara","Indonesia")
	InsertLocation("Mataram","Nusa Tenggara Barat","Indonesia")
	InsertLocation("Medan","Sumatera Utara","Indonesia")
	InsertLocation("Metro","Lampung","Indonesia")
	InsertLocation("Mojokerto","Jawa Timur","Indonesia")
	InsertLocation("Padang","Sumatera Barat","Indonesia")
	InsertLocation("Padangpanjang","Sumatera Barat","Indonesia")
	InsertLocation("Padangsidempuan","Sumatera Utara","Indonesia")
	InsertLocation("Pagar Alam","Sumatera Selatan","Indonesia")
	InsertLocation("Palangka Raya","Kalimantan Tengah","Indonesia")
	InsertLocation("Palembang","Sumatera Selatan","Indonesia")
	InsertLocation("Palopo","Sulawesi Selatan","Indonesia")
	InsertLocation("Palu","Sulawesi Tengah","Indonesia")
	InsertLocation("Pangkal Pinang","Kepulauan Bangka Belitung","Indonesia")
	InsertLocation("Parepare","Sulawesi Selatan","Indonesia")
	InsertLocation("Pariaman","Sumatera Barat","Indonesia")
	InsertLocation("Pasuruan","Jawa Timur","Indonesia")
	InsertLocation("Payakumbuh","Sumatera Barat","Indonesia")
	InsertLocation("Pekalongan","Jawa Tengah","Indonesia")
	InsertLocation("Pekanbaru","Riau","Indonesia")
	InsertLocation("Pematangsiantar","Sumatera Utara","Indonesia")
	InsertLocation("Pontianak","Kalimantan Barat","Indonesia")
	InsertLocation("Prabumulih","Sumatera Selatan","Indonesia")
	InsertLocation("Probolinggo","Jawa Timur","Indonesia")
	InsertLocation("Sabang","Aceh","Indonesia")
	InsertLocation("Salatiga","Jawa Tengah","Indonesia")
	InsertLocation("Samarinda","Kalimantan Timur","Indonesia")
	InsertLocation("Sawahlunto","Sumatera Barat","Indonesia")
	InsertLocation("Semarang","Jawa Tengah","Indonesia")
	InsertLocation("Serang","Banten","Indonesia")
	InsertLocation("Sibolga","Sumatera Utara","Indonesia")
	InsertLocation("Singkawang","Kalimantan Barat","Indonesia")
	InsertLocation("Solok","Sumatera Barat","Indonesia")
	InsertLocation("Sorong","Papua Barat","Indonesia")
	InsertLocation("Subulussalam","Aceh","Indonesia")
	InsertLocation("Sukabumi","Jawa Barat","Indonesia")
	InsertLocation("Sungai Penuh","Jambi","Indonesia")
	InsertLocation("Surabaya","Jawa Timur","Indonesia")
	InsertLocation("Surakarta","Jawa Tengah","Indonesia")
	InsertLocation("Tangerang","Banten","Indonesia")
	InsertLocation("Tangerang Selatan","Banten","Indonesia")
	InsertLocation("Tanjung Pinang","Kepulauan Riau","Indonesia")
	InsertLocation("Tanjungbalai","Sumatera Utara","Indonesia")
	InsertLocation("Tarakan","Kalimantan Utara","Indonesia")
	InsertLocation("Tasikmalaya","Jawa Barat","Indonesia")
	InsertLocation("Tebing Tinggi","Sumatera Utara","Indonesia")
	InsertLocation("Tegal","Jawa Tengah","Indonesia")
	InsertLocation("Ternate","Maluku Utara","Indonesia")
	InsertLocation("Tidore Kepulauan","Maluku Utara","Indonesia")
	InsertLocation("Tomohon","Sulawesi Utara","Indonesia")
	InsertLocation("Tual","Maluku","Indonesia")
	InsertLocation("Kediri", "Jawa Timur","Indonesia")
	InsertLocation("Madiun", "Jawa Timur","Indonesia")
	InsertLocation("Malang", "Jawa Timur","Indonesia")
	InsertLocation("Pasuruan", "Jawa Timur","Indonesia")
	InsertLocation("Probolinggo", "Jawa Timur", "Indonesia")
	InsertLocation("Surabaya", "Jawa Timur","Indonesia")
	InsertLocation("Tuban", "Jawa Timur","Indonesia")
	InsertLocation("Balikpapan", "Kalimantan Timur", "Indonesia")
	InsertLocation("Samarinda", "Kalimantan Timur", "Indonesia")
	InsertLocation("Kupang", "Nusa Tenggara Timur", "Indonesia")
	InsertLocation("Jakarta Timur", "Jakarta","Indonesia")
	InsertLocation("Jakarta Barat", "Jakarta","Indonesia")
	InsertLocation("Jakarta Pusat", "Jakarta","Indonesia")
	InsertLocation("Jakarta Utara", "Jakarta","Indonesia")
	InsertLocation("Jambi","Jambi","Indonesia")
	InsertLocation("Bandar Lampung", "Lampung", "Indonesia")
	InsertLocation("Ambon", "Maluku","Indonesia")
	InsertLocation("Manado", "Sulawesi Utara","Indonesia")
	InsertLocation("Belawan", "Sumatera Utara", "Indonesia")
	InsertLocation("Medan", "Sumatera Utara","Indonesia")
	InsertLocation("Pematangsiantar", "Sumatera Utara", "Indonesia")
	InsertLocation("Sibolga", "Sumatera Utara", "Indonesia")
	InsertLocation("Jayapura", "Papua","Indonesia")
	InsertLocation("Pekanbaru", "Riau","Indonesia")
	InsertLocation("Banjarmasin", "Kalimantan Selatan", "Indonesia")
	InsertLocation("Palangkaraya", "Kalimantan Selatan", "Indonesia")
	InsertLocation("Makassar", "Sulawesi Selatan", "Indonesia")
	InsertLocation("Kendari","Sulawesi Selatan", "Indonesia")
	InsertLocation("Palembang", "Sumatera Selatan", "Indonesia")
	InsertLocation("Pangkalpinang", "Sumatera Selatan", "Indonesia")
	InsertLocation("Kendari", "Sulawesi Tenggara", "Indonesia")
	InsertLocation("Bandung" , "Jawa Barat","Indonesia")
	InsertLocation("Bogor" , "Jawa Barat","Indonesia")
	InsertLocation("Cirebon" , "Jawa Barat","Indonesia")
	InsertLocation("Sukabumi" , "Jawa Barat","Indonesia")
	InsertLocation("Tasikmalaya" , "Jawa Barat", "Indonesia")
	InsertLocation("Pontianak", "Kalimantan Barat", "Indonesia")
	InsertLocation("Mataram", "Nusa Tenggara Barat", "Indonesia")
	InsertLocation("Bukttinggi", "Sumatera Barat", "Indonesia")
	InsertLocation("Padang", "Sumatera Barat","Indonesia")
	InsertLocation("Yogyakarta", "Daerah Istimewa (Yogyakarta","Indonesia")
}
func GetAllLocation()[]Location{
	db:=Connection.Connect()
	defer  db.Close()

	var res []Location
	db.Find(&res)

	return res
}

func GetLocByCity(city string)Location{
	db := Connection.Connect()
	defer db.Close()

	var res Location
	db.Where("City = ?", city).First(&res)

	return res
}

func GetLocByProvince(province string) []Location{
	db := Connection.Connect()
	defer db.Close()

	var res []Location
	db.Where("Province = ?", province).Find(&res)

	return res
}

func getLocByCountry(country string)[]Location{
	db := Connection.Connect()
	defer db.Close()

	var res []Location
	db.Where("Country = ?", country).Find(&res)

	return res
}


