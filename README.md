## Circular Printer

Ada sebuah lingkaran huruf dari A sampai Z. Pointer awal selalu dimulai dari A. Setiap perpindahan dari posisi awal ke huruf selanjutnya dihitung 1 detik dan pilih waktu terkecil, misal:
- jika A ke B maka 1 detik
- jika A ke C maka 2 detik
- jika A ke Z maka 1 detik
- jika Z ke B maka 2 detik
- jika Z ke A maka 1 detik

# contoh
- input: AZB maka 0 + 1 + 2 = 3 detik
- input: ZNMD maka 1 + 12 + 1 + 9 = 23 detik