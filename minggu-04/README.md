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

Docker Network merupakan sebuah opsi menu yang memungkinkan kita untuk melakukan segala hal yang berhubungan dengan manajemen administrasi jaringan, seperti membuat jaringan, menghubungkan, melihat informasi jaringan.

Dalam praktikum ini melalui course katacoda akan mencoba membuat docker network yang memungkinkan setiap docker container bisa saling berkomunikasi.

### Step 1 -  Membuat Jaringan

Jalankan perintah `docker network create backend-network` perintah ini digunakan untuk membuat docker network dengan nama `backend-network`. [Info detail](https://docs.docker.com/network/)

![docker network create predefined-name](img/02.networking-intro/Selection_001.jpg)

Setelah docker network terbuat, sertakan opsi `--net` ketika menjalankan docker container. Contoh menggunakan docker image redis

![docker run with net attribute](img/02.networking-intro/Selection_002.jpg)

Perintah diatas akan menjalankan container redis dan menjalankan menggunakan docker network `backend-network`.

### Step 2 - Network Communication

Tidak seperti links, `docker network` berperilaku seperti jaringan tradisional, dimana setiap node bisa di pasang/lepas.

Lihat environment variable dari sebuah container dengan menjalankan perintah `docker run --net=backend-network alpine env`, perintah ini akan menampilkan semua environment variable dari sebuah container.

![docker run network env](img/02.networking-intro/Selection_003.jpg)

![docker run network cat hosts](img/02.networking-intro/Selection_004.jpg)

Semua dns pada docker network akan di assign melalui IP `127.0.0.11` yang sudah di set pada `/etc/resolv.conf`

![docker run network cat resolv.conv](img/02.networking-intro/Selection_005.jpg)

Dan ketika 1 container melakukan ping menggunakan hostname dari container tujuan, ping akan mengembalikan output berupa IP Address dari hostname container tujuan. Contoh ping redis container yang sudah jalan dari container alpine.

![docker run network ping redis](img/02.networking-intro/Selection_006.jpg)

### Step 3 - Menghubungkan 2 Container

Docker support multiple nerwork dan container bisa di attach ke 1 jaringan secara bersamaan.

Buat docker network baru dengan nama `frontend-network`

![docker network frontend-network](img/02.networking-intro/Selection_007.jpg)

Untuk menghubungkan container yang berjalan dengan suatu docker network gunakan perintah `docker connect`

![docker attach redis to fontend-network](img/02.networking-intro/Selection_008.jpg)

Perintah ini akan me-attach redis ke dalam `frontend-network` yang baru saja dibuat. Jalankan web server dan assign network ke dalam `frontend-network`, sama dengan `redis` sehingga webserver bisa berkomunikasi dengan redis melalui jaringan yang sama.

![docker run web server](img/02.networking-intro/Selection_009.jpg)

Ketika ditest menggunakan curl, bisa di akses pada 1 hostname yaitu `docker`

![docker curl](img/02.networking-intro/Selection_010.jpg)

### Step 4 - Buat Alias

Alias pada docker akan mempermudah ketika akan mengakses kontainer. Dengan alias ini kita bisa mendefinisikan nama (dns) dari container yang akan kita akses. Dalam course ini, redis menggunakan alias `db` yang lebih bisa merepresentasikan bahwa redis yang dijalankan disini akan digunakan sebagai database dari webserver.

Buat network baru dengan nama `frontend-network2` dengan perintah `docker network create frontend-network2`

Kemudian jalankan redis pada network yang baru dibuat dan beri alias db:

![docker network alias](img/02.networking-intro/Selection_011.jpg)

Dengan begitu redis bisa di akses dengan nama db.

![docker ping alias](img/02.networking-intro/Selection_012.jpg)

### Step 5 - Disconnect Containers

![docker network ls](img/02.networking-intro/Selection_013.jpg)

`docker network ls` adalah perintah docker untuk menampilkan `docker network` yang sudah ada di komputer. Dari informasi jaringan ini kita bisa menampilkan informasi detail dari docker network termasuk docker container mana yang menggunakan docker network tertentu.

![docker network inspect](img/02.networking-intro/Selection_014.jpg)

Untuk melepaskan container dari docker network tertentu, gunakan perintah `docker network disconnect frontend-network nama-container`.

![docker network disconnect](img/02.networking-intro/Selection_015.jpg)
