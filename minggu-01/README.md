# Praktikum Pertemuan Minggu ke-01

## Pre Requisite

Sebelum melakukan instalasi `git` pada windows, pastikan terdapat teks editor (notepad++, visual studio code, sublime, vim, etc) sudah terinstall. Keberadaan teks editor ini akan mempengaruhi keberhasilan dari proses instalasi git di komputer nantinya.

## Instalasi Git

1. Download master `git` melalui source link dari [git-scm](https://git-scm.com/download/win)

    ![git-scm download page](img/01-download page.PNG)

2. Setelah selesai proses download file installer, double click pada file untuk menjalankan proses instalasi git pada komputer.

    ![Install git](img/02.file installer.PNG)

    Setelah itu akan muncul halaman informasi licensi dari installer git. Klik **Next** untuk melanjutkan instalasi.

    ![Licensi instalasi](img/03.installler information.PNG)

3. Langkah selanjutnya adalah menentukan letak instalasi git, secara default git akan terinstall di _C:\Program Files\Git_. Ubah lokasi jika ingin meletakkan instalasi git di tempat yang diinginkan.

    ![Select installation destination](img/04.installer select destination.PNG)

    Selain menunjukkan letak instalasi, pada halaman ini juga menampilkan informasi disk space pada partisi disk yang menjadi tujuan. Klik **Next** untuk melanjutkan proses instalasi.

4. Pada langkah ini terdapat pilihan komponen yang akan di sertakan pada instalasi git di komputer, jika ragu abaikan pilihan komponen, biarkan secara default kemudian klik **Next**

    ![Installer Component](img/05.installer select components.PNG)

5. **Select Start Menu Folder** sesuai namanya, halaman ini di gunakan untuk menentukan nama folder pada start menu untuk mengelompokkan nama aplikasi dari instalasi git sehingga bisa di akses melalui start menu. Secara default pada start menu akan muncul kelompok menu dengan nama **Git** ubah sesuai keinginan, atau centang pada checkbox _**Don't create a Start Menu folder**_ jika tidak ingi muncul kelompok menu **Git** pada **Start Menu**. Kemudian klik **Next** untuk melanjutkan.

    ![Select Start Menu Folder](img/06.installer set start menu folder.PNG)

6. **Choosing the default editor used by git** pilih editor default yang akan di gunakan untuk proses editing jika membuka/edit file melalui git gui. Klik pada panah ke bawah dropdown untuk menampilkan daftar text editor yang sudah ada di dalam komputer.

    ![Choose default editor](img/07.installer choose default editor.PNG)

7. Pada saat instalasi, Git menyediakan akses git melalui Bash maupun command prompt. Pilih pilihan kedua supaya bisa menggunakan dari dua antarmuka tersebut. Bash adalah shell di Linux. Dengan menggunakan bash di Windows, pekerjaan di command line Windows bisa dilakukan menggunakan bash - termasuk ekskusi dari Git.

    ![PATH Environment](img/08.installer PATH Environment.PNG)

8. **HTTPS transport backend** Pilih OpenSSL untuk HTTPS. Git menggunakan protocol https untuk akes ke repo GitHub atau repo-repo lain (GitLab, Assembla).

    ![https transport backend](img/09.installer HTTPS transport.PNG)

9. **Configuring the line ending conversion** Pilih pilihan pertama untuk konversi akhir baris (CR-LF).

    ![CR-LF](img/10.installer CR-LF.PNG)

10. **Configuring the terminal emulator to use with Git Bash** Pilihan terminal emulator (Tampilan halaman terminal) yang dipakai oleh Git Bash. Pilih default MinTTY

    ![GitBash terminal emulator](img/11.installer terminal emulator.PNG)

11. **Configuring extra options** Centang fitur tambahan yang akan ikut di sertakan dalam instalasi. Pilih pilihan 1 dan 2 (default)

    ![Extra options](img/12.installer extra options.PNG)

12. **Installing** Proses instalasi berjalan.

    ![Instalasi](img/13.installing.PNG)

13. **Finish instaling git** Klik **Finish** dan proses instalasi selesai. Cek instalasi, dengan membuka start menu, cari group menu **Git**

    ![Finish instalation](img/14.finish.PNG)

    ![Start menu](img/16.start menu.png)

14. **Menjalankan Git** pada start menu, pilih
    1. Git Bash

        ![Git Bash](img/15.git version.PNG)

        Tampilan dari git bash ketika di jalankan, untuk memeriksa instalasi git, ketikkan `git version` pada terminal

    2. Git Gui

        ![Git Gui](img/17.git gui.png)

        Tampilan git gui. Terdapat beberapa menu untuk memulai `create new repository`, `Clone existing repository`, `Open existing repository`
