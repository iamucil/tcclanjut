# Pertemuan Minggu-04

## Docker - Menghubungkan antar jaringan

> Referensi : [https://katacoda.com/courses/docker/5](https://katacoda.com/courses/docker/5)

Docker memiliki dua pendekatan untuk jaringan. Yang pertama mendefinisikan hubungan antara dua kontainer. `links` memperbarui file `/etc/hosts` dan **environment variabel** untuk memungkinkan kontainer saling berkomunikasi.

Pendekatan alternatif yang kedua adalah membuat `docker network` untuk menghubungukan kontainer. Jaringan memiliki atribut yang mirip dengan jaringan fisik, memungkinkan kontainer untuk saling terhubung lebih bebas daripada saat menggunakan `links`.

### Step 1 - Start Redis

Skenario yang paling umum dalam implementasi `links` adalah menghubungkan container (aplikasi) dengan data-store container, dalam praktikum kali ini menggunakan docker image _redis data store_. Secara default ketika menjalankan sebuah container, docker akan membuat nama untuk container yang dijalankan. Untuk kemudahan dalam penggunaan `links` definisikan `name` container yang dijalankan.

```markdown
docker run -d --name redis-server redis
```

### Step 2 - Create Link

Untuk menghubungkan dengan container asal gunakan opsi `--link <container-name|id>:<alias>` ketika menjalankan container yang baru.

> Bagaimana Link bekerja?

Pertama, docker akan mengeset beberapa environment variable berdasarkan container sumber. Untuk menampilkan environment variable gunakan perintah: `docker run --link redis-server:redis alpine env`

![show variables](img/01.communicating-between-containers/Selection_002.jpg)

Selanjutnya, Docker akan melakukan update terhadap file `HOSTS` pada container dengan 3 variable yaitu: _the original, the alias and the hash-id_

jalankan perintah `docker run --link redis-server:redis alpine cat /etc/hosts`

![cat variable](img/01.communicating-between-containers/Selection_003.jpg)

Ketika link sudah terbuat lakukan ping terhadap container dengan perintah `docker run --link redis-server:redis alpine ping -c 1 redis`

![ping src container](img/01.communicating-between-containers/Selection_004.jpg)

## Docker - Networks

> Referensi : [https://katacoda.com/courses/docker/networking-intro](https://katacoda.com/courses/docker/networking-intro)
