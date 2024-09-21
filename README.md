# Program Konversi Suhu
Muhammad Nafal Zakin Rustanto - 24/5352555/TK/59364 - Teknologi Informasi

## Deskripsi Program

Program Go ini meminta pengguna untuk memasukkan suhu dalam Celsius dan kemudian menawarkan pilihan untuk mengkonversinya ke Fahrenheit, Kelvin, atau Reamur. Program ini menggunakan mekanisme conditional statement `swicth - case` untuk menentukan konversi suhunya

## Alur Program

1. Program dimulai dengan mengimpor package yang diperlukan (`fmt` untuk operasi input/output).

2. Dalam fungsi `main()`:
   - Program mendeklarasikan variabel `suhu` bertipe `float32` untuk menyimpan suhu input.
   - Pengguna diminta untuk memasukkan suhu dalam Celsius.
   - Program membaca input menggunakan `fmt.Scanln()`.

3. Program kemudian menampilkan pilihan konversi kepada pengguna:
   - Case 1: Fahrenheit
   - Case 2: Kelvin
   - Case 3: Reamur

5. Pengguna diminta untuk memilih opsi dengan memasukkan angka (1, 2, atau 3).

6. Menggunakan pernyataan `switch`, program melakukan konversi berdasarkan pilihan pengguna:
   - Case 1: Mengkonversi ke Fahrenheit menggunakan rumus `(C * 9/5) + 32`
   - Case 2: Mengkonversi ke Kelvin dengan menambahkan 273 ke suhu Celsius
   - Case 3: Mengkonversi ke Reamur menggunakan rumus `C * 4/5`
   - Default: Menampilkan pesan error untuk input yang tidak valid

7. Hasil konversi kemudian dicetak ke terminal.

## Cara Menjalankan

Untuk menjalankan program ini, pastikan Go sudah terinstall dan ikuti langkah-langkah berikut:

1. Simpan program dalam file dengan ekstensi `.go` (misalnya, `ConvertSuhu.go`).
2. Buka terminal atau command prompt.
3. Navigasikan ke direktori yang berisi file Go tersebut.
4. Jalankan perintah: `go run ConvertSuhu.go`
5. Ikuti petunjuk di layar untuk memasukkan suhu dan memilih opsi konversi.

## Catatan

Program dibuat sepenuhnya oleh saya tanpa bantuan AI, dan dibuat untuk penugasan Pelatihan Web Development KMTETI
