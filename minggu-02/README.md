# Pertemuan Minggu-02

## Berkolaborasi di git repository

Selain untuk mengelola aset digital milik diri sendiri, kita bisa menggunakan Git untuk berkolaborasi dalam suatu repo di GitHub yang bisa diakses bersama. Dalam kasus seperti ini, berarti ada 2 peran:

1. Pemilik repo, sering disebut sebagai *upstream author*.
2. Kontributor, yaitu orang-orang yang akan berkontribusi memberikan konten.

Untuk situasi seperti ini, diasumsikan:

1. *Upstream author* telah membuat repo git di GitHub
2. Kontributor telah mengetahui adanya repo tersebut, tertarik untuk berkontribusi, sudah mengetahui apa yang akan diberikan ke proyek (repo GitHub *upstream author*) tersebut.
3. Pembahasan selanjutnya adalah tentang bagaimana kontributor bisa mengirimkan kontribusi ke repo GitHub milik *upstream author*.

Skenario:

1. *Upstream author* adalah *hasan354313*.
2. Kontributor adalah *iamucil*
3. Repo dari *upstream author* adalah **playground** yang bisa diakses di [https://github.com/hasan354313/tcclanjut](https://github.com/hasan354313/tcclanjut)

## Forking

Fork adalah membuat clone dari suatu repo di GitHub milik *upstream author*, diletakkan ke milik kontributor. Fork hanya dilakukan sekali saja. Kontributor harus mem-*fork* repo *upstream author* sehingga di repo kontributor muncul repo tersebut. Proses *forking* ini dijelaskan dengan langkah-langkah berikut:

1. Login ke Github
2. Akses repository *Upstream Author* di [https://github.com/hasan354313/tcclanjut](https://github.com/hasan354313/tcclanjut)

3. Pada sisi kanan atas, pilih fork.

    ![Fork Repository](img/01_fork_repository.png)

4. Pilih repository akan ditaruh di akun yang mana

    ![Choose account](img/02_fork_repository_pilih_lokasi_fork.png)

    dalam kasus ini, hasil forking akan ditaruh pada akun *iamucil*

    ![Process forking](img/03_proses_forking.png)

5. Setelah selesai proses forking, clone repository ke local untuk mulai berkontribusi. Dengan cara, pilih tombol *Clone or Download* yang ada di sisi kanan atas, pilih *https* kemudian copy url dari repository. `tcclanjut-1` adalah nama repository dari hasil fork karena sudah ada nama repository `tcclanjut` untuk akun `iamucil`

    ![clone url](img/04_copy_repo_url.png)

    Clone repostory ke local dengan menjalanlan perintah:

    ```git clone repository_url```

    ![Clone local](img/05_clone_repository.png)

    Keterangan: Clone repository [https://github.com/iamucil/tcclanjut-1](https://github.com/iamucil/tcclanjut-1) ke folder local `~/git/hasan`

6. Buat perubahan. Dalam kasus ini buat folder baru.

    ![Make some change](img/06_make_some_change.png)

    Keterangan: `cd tcclanjut-1` hasil cloning repository, kemudian buat folder baru dengan nama `pull-request-175410208` dengan perintah `mkdir pull-request-175410208` buat file kosong dengan `.empty` dengan perintah `touch .empty` hal ini karena git akan mengabaikan direktori kosong ketika proses push ke remote repository.

7. Lihat perubahan dengan perintah `git status`

    ![Git Status](img/07_git_status.png)

    Keterangan: `git status` adalah perintah `git` untuk melihat perubahan apa saja yang ada di local, dalam hal ini terlihat ada 1 buath file `.empty` di dalam direktori `pull-request-175410208` yang belum ada di dalam `remote repository`.

8. Tambahkan file/folder baru ke remote repository dengan perintah `git add .`

    ![Git add](img/08_git_add.png)

    ![Git status added](img/09_git_status_new_file.png)

    Keterangan: `git add .` perintah ini digunakan untuk menambahkan `untracked files` yang ada di `local` ke `remote`. Jalankan perintah `git status` untuk memastikan file sudah daftarkan untuk di tambah ke remote repository ketika di lakukan push, ditandai dengan status `new file: .empty` dengan warna hijau.

9. `git commit`: Setelah perubahan yang diinginkan sudah selesai di lakukan, file/folder sudah ditata dan sudah ditambahkan, langkah selanjutnya rekam perubahan (commit).

    ![Commit](img/10_git_comit.png)

    Keterangan: `git commit -m 'tcclanjut: Make some pr'` perintah ini digunakan untuk merekam perubahan yang ada di local dan memberikan pesan perubahan apa yang dilakukan atau memeberi pesan kepada author/kontributor lain tentang pekerjaan yang sudah dilakukan.

10. Push/kirim perubahan yang ada di local ke repository git, setelah perubahan yang ada dilocal di commit.

    ![Git push](img/11_git_push.png)

    Keterangan: `git push origin master` > Push (kirim) perubahan ke remote:origin yaitu repository milik sendiri, `master` adalah nama branch yang digunakan untuk menyimpan file/asset yang sudah di ubah.

    ![login](img/12_login_form.png)

    Muncul halaman login, karena informasi username dan password dari git tidak di simpan di local, karena alasan keamanan. Mengingat komputer yang dipakai adalah komputer untuk umum.

    ![Push succed](img/13_git_push_success.png)

11. Akses repository git untuk melihat perubahan yang dilakukan sudah berhasil di kirim.

    ![New Change](img/14_new_change.png)

    Terdapat pesan `this branch is 1 commit ahead of hasan354313:master` menandakan perubahan sudah berhasil dilakukan dan sudah terkirim ke repository yang ada di github, dan siap untuk di *PR*

12. Pull request ke `upstream` repo, dengan cara klik pada tombol `New pull request`

    ![Compare changes](img/15_compare_repo.png)

    Muncul halaman baru untuk melakukan pull request, halaman ini digunakan untuk melihat dan membandingkan perubahan apa saja yang di lakukan di head repository dengan base repository. Halaman ini juga digunakan untuk memastikan bahwa perubahan sudah sesuai dengan keinginan dan tidak ada `conflict` ketika akan di lakukan *PR*.

    Jika sudah yakin dengan perubahan yang dilakukan, klik tombol hijau : `Create pull request` untuk melakukan Pull Request.

13. *Open Pull Request*

    ![Open Pull Request](img/16_open_pr.png)

    Jika file yang diubah tidak ada conflict akan ada notifikasi `Able to merge` ini menandakan tidak ada masalah dengan perubahan yang kita lakukan dengan base repository, dan pull request bisa di merge. Form ini digunakan untuk mencatat pesan PR ke pemilik base repository. Klik tombol `Create pull request` untuk melakukan `pull request`.

    ![PR Opened](img/17_pr_opened.png)

    Keterangan: Tampilan PR sudah dibuka, dan menunggu proses `aprove` dari pemilik base/upstream repository.

14. Pada repo `upstream author`, muncul angka 1 (artinya jumlahnya 1) pada Pull requests di bagian atas.

15. *Upstream author* bisa menyetujui setelah melakukan review: klik pada ```Pull requests```, akan muncul PR dengan message seperti yang ditulis oleh kontributor (*Add: contributor*). Klik pada PR tersebut, review kemudian klik ```Merge pull request``` diikuti dengan ```Confirm merge```. Setelah itu, status akan berubah menjadi ```Merged```.

    ![PR Merged](img/18_pr_merged.png)
