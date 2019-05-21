# Pertemuan Minggu-07

## IRIS Go-Framework

## Kubernetes

> Start containers dengan Kubectl

Jalankan perintah `minikube start` untuk download cli kuberctl dan start komponen dari cluster

![minikube start](img/kubectl/001.png)

Jalankan perintah `kubectl get nodes` untuk melihat status dari node yang sudah dijalankan.

![node checking](img/kubectl/002.png)

Untuk menjalankan container berdasarkan docker image, bisa menggunakan perintah `kubectl run <name of deployment> <properties>
`

![Kubernetes run](img/kubectl/003.png)

> Keterangan gambar, kubectl menjalankan sebuah container bernama http berdasarkan docker image `katacoda/docker-http-server:latest` sebanyak 1 replica.

Untuk mendapatkan status dari proses deployment dari kubectl, gunakan perintah  `kubectl get deployments`, perintah ini akan menghasilkan output pada terminal seperti pada gambar dibawah ini:

![Kubernetes get deployments](img/kubectl/004.png)

sedangkan untuk mendapat status deployement spesifik berdasarkan nama:

![Kubernetes get deplyment by container nama](img/kubectl/005.png)

perintah ini akan menghasilkan output yang lebih detail dari container yang di maksud, seperti yang terlihat pada gambar, terdapat informasi berapa replika yang berjalan, label, event yang berhubungan dengan container yang dijalankan.

Setelah deployment selesai dibuat, untuk meng-ekspose service ke sebuah port sehingga bisa di akses gunakan perintah `kubectl expose`.

![Kubernetes expose port](img/kubectl/006.png)

> Keterangan gambar, menggunakan perintah `kubectl expose` untuk meng-ekspose port dari service yaitu 80 ke mesin dengan port 80, sehingga service http bisa di akses melalui port 8080.

Gunakan perintah `curl http://172.17.0.18:8000` untuk test service yang berjalan.

![Kubernetes expose port](img/kubectl/006-curl.png)

### Kubectl menjalankan service dan langsung expose port dalam satu waktu

![Kubernetes expose port](img/kubectl/006-001.png)

Untuk mempersingkat perintah `kubectl` seperti yang ada digambar diatas bisa digunakan untuk menjalankan sebuah service dan langsung meng-ekspos port dari service yang di maksud. Dari gambar diatas bisa di lihat kubectl menjalankan service httpexposed berdasarkan docker image `katacoda/docker-http-server:latest` sebanyak 1 replica dengan mengekspose port 80 ke port 8081.

![Kubernetes expose port](img/kubectl/006-002.png)

Jika ditest menggunakan perintah curl, akan menghasilkan output seperti pada gambar diatas.

![Kubernetes expose port](img/kubectl/006-003.png)

Dengan perintah ini, service tidak akan muncul ketika dilihat dari perintah `kubectl get svc`. Untuk melihat service gunakan perintah: `docker ps | grep httpexposed` sehingga akan tampil seperti pada gambar dibawah ini.

![Kubernetes expose port](img/kubectl/006-004.png)

### Scaling container dengan kubernetes

Setelah service berjalan, dengan kubernetes bisa dengna mudah menambah service ke dalam pods.

![Kubernetes expose port](img/kubectl/007.png)

Perintah diatas akan menambah 3 replika deployment http ke dalam pods. Ketika menjalankan perintah `kubectl get pods` bisa dilihat berapa banyak http deployment yang berjalan.

![Kubernetes expose port](img/kubectl/008.png)

Setelah setiap Pod berjalan, akan otomatis ditambah ke dalam load balancer. Dengan menjalan perintah `kubectl describe svc http` akan muncul deskripsi tentang pods beserta end point dan Pod yang saling terkait dengan service http.

![Kubernetes expose port](img/kubectl/009.png)

Test menggunakan perintah `curl`

![Kubernetes expose port](img/kubectl/010.png)

## Deploy Container dengan File YAML

Salah satu hal yang paling umum dari kubernetes adalah object deploymentnya. Untuk menjalankan sebuah container menggunakan kubernetes, objek-objek yang diinginkan bisa disimpan dalam sebuah file `yaml`. Buat sebuah file dengan nama `deployment.yaml` dan isi dengan:

![deployment file](img/kubectl-yaml/01-001.png)

Kemudian deploye menggunakan perintah kubectl:

![deployment file](img/kubectl-yaml/01-002.png)

Dengan perintah ini kubectl akan mendeploy container webapp1 sesuai dengan yang terdefinisikan di dalam file `deployment.yaml`. untuk memastikan deployemnt berjalan, jalankan perintah seperti pada gambar dibawah ini.

![get deployment](img/kubectl-yaml/01-003.png)

Dan untuk melihat detail infomasi dari deployment yang dibuat:

![deployment file](img/kubectl-yaml/01-004.png)

Dengan perintah diatas akan menampilkan informasi lengkap terkait deployment yang dibuat menggunakan perintah kubectl.

### Kubernetes Service

Kubernetes mempunyai kemampuan yang ampuh untuk mengontrol bagaiman aplikasi saling berkomunikasi. Konfigurasi jaringan ini dalam kubernetes juga bisa di kontrol dalam sebuah file `yaml`. Buat sebuah file `service.yaml`

![service.yaml](img/kubectl-yaml/02-001.png)

Service akan mencari semua service denga label `webapp1` dan akan membuat aplikasi tersedia melalui `NodePort`. Deploy service dengan perintah:

![deploy service](img/kubectl-yaml/02-002.png)

Sama seperti sebelumnya, lihat informasi tentang service dengan perintah `kubectl get svc`

![deploy service](img/kubectl-yaml/02-003.png)

Untuk melihat informasi detail tentang service, jalankan perintah seperti pada gambar dibawah.

![deploy service](img/kubectl-yaml/02-004.png)

Pastikan service bisa di akses dengan perintah curl dari mesin.

![deploy service](img/kubectl-yaml/02-005.png)

### Meningkatkan deployment

Detail tentang file `yaml` bisa berubah setiap saat, tergantung dari kebutuhan development sebuah app. Contoh, untuk menambahkan service yang berjalan, update file `deployment.yaml` dan updata pada bagian `replicas:` menjadi `replicas: 4`. Setelah selesai dengan update file, apply perubahan.

![apply update](img/kubectl-yaml/03-001.png)

Pastikan perubahan sukses di aplikasikan:

![apply update](img/kubectl-yaml/03-002.png)

![apply update](img/kubectl-yaml/03-003.png)

Karena semua pods mempunyai label yang sama, Pods akan di load balancing melalui NodePort. Ketika di test menggunakan `curl` akan menghasilkan output yang berbeda tergantung service mana yang dapat merespond request dari client:

![curl deployed service](img/kubectl-yaml/03-004.png)
