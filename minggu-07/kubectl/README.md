# Start Containers Menggunakan `kubectl`

## kubectl

> Start containers menggunakan Kubectl

Jalankan perintah `minikube start` untuk download cli kuberctl dan start komponen dari cluster

![minikube start](img/001.png)

Jalankan perintah `kubectl get nodes` untuk melihat status dari node yang sudah dijalankan.

![node checking](img/002.png)

Untuk menjalankan container berdasarkan docker image, bisa menggunakan perintah `kubectl run <name of deployment> <properties>
`

![Kubernetes run](img/003.png)

> Keterangan gambar, kubectl menjalankan sebuah container bernama http berdasarkan docker image `katacoda/docker-http-server:latest` sebanyak 1 replica.

Untuk mendapatkan status dari proses deployment dari kubectl, gunakan perintah  `kubectl get deployments`, perintah ini akan menghasilkan output pada terminal seperti pada gambar dibawah ini:

![Kubernetes get deployments](img/004.png)

sedangkan untuk mendapat status deployement spesifik berdasarkan nama:

![Kubernetes get deplyment by container nama](img/005.png)

perintah ini akan menghasilkan output yang lebih detail dari container yang di maksud, seperti yang terlihat pada gambar, terdapat informasi berapa replika yang berjalan, label, event yang berhubungan dengan container yang dijalankan.

Setelah deployment selesai dibuat, untuk meng-ekspose service ke sebuah port sehingga bisa di akses gunakan perintah `kubectl expose`.

![Kubernetes expose port](img/006.png)

> Keterangan gambar, menggunakan perintah `kubectl expose` untuk meng-ekspose port dari service yaitu 80 ke mesin dengan port 80, sehingga service http bisa di akses melalui port 8080.

Gunakan perintah `curl http://172.17.0.18:8000` untuk test service yang berjalan.

![Kubernetes expose port](img/006-curl.png)

### Kubectl menjalankan service dan langsung expose port dalam satu waktu

![Kubernetes expose port](img/006-001.png)

Untuk mempersingkat perintah `kubectl` seperti yang ada digambar diatas bisa digunakan untuk menjalankan sebuah service dan langsung meng-ekspos port dari service yang di maksud. Dari gambar diatas bisa di lihat kubectl menjalankan service httpexposed berdasarkan docker image `katacoda/docker-http-server:latest` sebanyak 1 replica dengan mengekspose port 80 ke port 8081.

![Kubernetes expose port](img/006-002.png)

Jika ditest menggunakan perintah curl, akan menghasilkan output seperti pada gambar diatas.

![Kubernetes expose port](img/006-003.png)

Dengan perintah ini, service tidak akan muncul ketika dilihat dari perintah `kubectl get svc`. Untuk melihat service gunakan perintah: `docker ps | grep httpexposed` sehingga akan tampil seperti pada gambar dibawah ini.

![Kubernetes expose port](img/006-004.png)

### Scaling container dengan kubernetes

Setelah service berjalan, dengan kubernetes bisa dengna mudah menambah service ke dalam pods.

![Kubernetes expose port](img/007.png)

Perintah diatas akan menambah 3 replika deployment http ke dalam pods. Ketika menjalankan perintah `kubectl get pods` bisa dilihat berapa banyak http deployment yang berjalan.

![Kubernetes expose port](img/008.png)

Setelah setiap Pod berjalan, akan otomatis ditambah ke dalam load balancer. Dengan menjalan perintah `kubectl describe svc http` akan muncul deskripsi tentang pods beserta end point dan Pod yang saling terkait dengan service http.

![Kubernetes expose port](img/009.png)

Test menggunakan perintah `curl`

![Kubernetes expose port](img/010.png)
