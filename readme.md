<h1 align="center">WIKI SCRAPER</h1>

## Project Overview
This project is aimed at fulfilling the requirements for Small Task 2 of Algorithm Strategy, which involves implementing the Brute Force and Divide And Conquer algorithms in creating Bezier Curves. The implementation of Bezier Curves also includes the creation of a GUI that can display visual results of the executed algorithms.

This project aimed at fulfilling the requirements for Big Task 2 of Algorithm Strategy, which involves implementing the BFS and IDS for To create a scraper to scrap a wikipedia to get from a certain url to a target url. The implementaion includes the creation of a Web that can be used as user input

## Implementation
### BFS
Algoritma BFS yang kami implementasikan menggunakan sebuah priority queue dalam melakukan perhitungannya hal ini kami lakukan agar sekiranya dapat menemukan node yang coock terlebih dahulu antara url saat ini dengan target sehingga perhitungan menjadil lebih sedikit. Kami menggunakan string matching sebagai prioritasnya dimana semakin mirip target dengan current url makan semakin tinggi prioritasnya
### IDS 
Algoritma IDS yang kami implementasikan akan melakukan iterasi sebanyak depth_limit yang kami tentukan sendiri. Kami membatasi depth_limit kami 6 karena kami mengasumsikan depth terjauh adalah 6. Untuk tiap iterasi tersebut kami akan melakukan scraping untuk nodes hingga depth saat itu hingga target url dapat ditemukan. Jika blm ditemukan iterasi akan dilanjutkan dengan depth ditambah 1.
## Setup Project

### Requirements
1. Download Node js:
<br>Link : 
    ```
    https://nodejs.org/en/download
    ```
2. Download Golang: <br>
    Windows <br>
    ```
    https://go.dev/doc/install
    ```
    Linux <br>
    1. Install manual
    ```
    https://go.dev/doc/install
    ```
    2. Install from terminal
    ```
    sudo apt install golang-go
    ```
3. Clone the Repo :
    ```
    git clone https://github.com/Loxenary/Tubes2_Wiki_Scrapper.git
    ```
4. Navigate to wikiscrapper
    ```
    cd src/wikiscrapper
    ```
5. Install Frontend dependencies
    ```
    npm install
    ```
6. Navigate to backend
    ```
    cd backend
    ```
7. Run go mod tidy
    ```
    go mod tidy
    ```
8. Install neeeded dependencies
    ```
    go get github.com/PuerkitoBio/goquery
    ```

### Running

1. Proceeed to src (From backend directory)
    ```
    cd ../..
    ```
2. Launch The Apps <br>
    Windows : 
    ```
    ./run
    ```

    Linux : 
    1. Manual Launch
        1. Create 2 terminal instances
        2. Then run following command
            ```
            Backend
            cd wikiscrapper/backend
            go run main.go prioqueue.go bfs.go links.go ids.go safemap.go

            Frontend
            cd wikiscrapper
            npm run dev
            ```

    2.  Using Shell 
        - Linux
        1. Create the shell executable 
            ```
            chmod +x run.sh
            ```
        2. Run the file
            ```
            ./run.sh
            ```
        - WSL
        1. Install dos2unix
            ```
            sudo apt install dos2unix
            ``` 
        2. Run the following command
            ```
            dos2unix run.sh
            ```
        3. Create the shell executable
            ```
            chmod +x run.sh
            ```
        4. Run the shell 
            ```
            ./run.sh
            ```
The project is properly setup

## Authors

<b>Special thanks to our contributors:</b>
1. Auralea Alvinia Syaikha (13522148) 
2. Muhammad Davis Adhipramana (13522157)
3. Pradipta Rafa Mahesa (13522162)
