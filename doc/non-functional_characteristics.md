# Nem-funkcionális jellemzők vizsgálata

## Benchmark

A benchmarkokat a Go-ban írt algoritmusokon futtattuk,
mivel a Go testing könyvtárát használtuk, ezen belül is a Benchmark
funkciókat: https://pkg.go.dev/testing#hdr-Benchmarks

Egészen sok dolgot lehet csinálni ezzel a funkcióval, és időbe került működésének elsajátítása.

Létre kell hozni egy ugyanolyan nevű fájlt `_test` végződéssel, mint az amelyiknek függvényeit akarjuk tesztelni.
Ezután a megfelelő szintaxist használva meghívhatóak a függvények. Ezután futtatni kell a tesztet a megfelelő
kapcsolókkal és a testing könyvtár automatikusan kiértékeli a benchmarkot.

A teszt futtatása elsőre nem volt triviális, de végül rájöttünk, hogy hogyan kell elindítani.
A futtatás során meg kell adni a futtatandó teszt fájlt és azt a fájlt, amiben a tesztelendő függvény található.

```bash
go test -bench=. bogo_sort_test.go bogo_sort.go
```

Ez kimenetként az alábbit adja:

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkBogoSort/input_size_6-12               27335032                41.66 ns/op
BenchmarkBogoSort/input_size_12-12              20608626                52.43 ns/op
BenchmarkBogoSort/input_size_13-12                     1        21105931474 ns/op
PASS
ok      command-line-arguments  23.431s
```

Rájöttünk, hogy ilyenkor csak egyszer futtatja a benchmarkot, viszont egy kapcsolóval megadható, hogy hányszor futtassa:

```bash
go test -bench=. -count 5 bogo_sort_test.go bogo_sort.go
```

Ennek kimenete:

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkBogoSort/input_size_6-12               25353583                42.29 ns/op
BenchmarkBogoSort/input_size_6-12               27953912                40.76 ns/op
BenchmarkBogoSort/input_size_6-12               28296075                41.82 ns/op
BenchmarkBogoSort/input_size_6-12               27125326                40.88 ns/op
BenchmarkBogoSort/input_size_6-12               28280616                40.45 ns/op
BenchmarkBogoSort/input_size_12-12              19808017                52.30 ns/op
BenchmarkBogoSort/input_size_12-12              21285954                53.48 ns/op
BenchmarkBogoSort/input_size_12-12              21606840                53.72 ns/op
BenchmarkBogoSort/input_size_12-12              21639661                52.82 ns/op
BenchmarkBogoSort/input_size_12-12              20797857                54.11 ns/op
BenchmarkBogoSort/input_size_13-12                     1        21279557446 ns/op
BenchmarkBogoSort/input_size_13-12              20139840                57.38 ns/op
BenchmarkBogoSort/input_size_13-12              20895702                58.11 ns/op
BenchmarkBogoSort/input_size_13-12              20782876                56.04 ns/op
BenchmarkBogoSort/input_size_13-12              19720352                55.92 ns/op
PASS
ok      command-line-arguments  37.939s
```

Valamint egy benchstat nevű go program segítségével kiértékelhetőek a különböző statisztikák, valamint össze is lehet
hasonlítani különböző méréseket.

```
name                       time/op
BogoSort/input_size_6-12   41.2ns ± 3%
BogoSort/input_size_12-12  53.3ns ± 2%
```

A benchmarkokat különböző bemenetekre lefuttattuk és mindegyiknél megadtunk egy bizonyos minimum időt, ameddig kell
futnia, hogy a mérés minél pontosabb és valósabb legyen.

A mért adatokat lementettük txt fájlokba, majd a benchstat segítségével összehasonlítottuk azokat:

```
name \ time/op               comb.txt     counting.txt  bubble.txt   gnome.txt    heap.txt
Sort/input_size=10;-12       37.2ns ± 2%   52.0ns ± 3%   6.5ns ± 0%   8.1ns ± 2%   243.7ns ± 3%
Sort/input_size=19;-12       91.4ns ±10%   79.9ns ± 3%  12.8ns ± 8%  15.9ns ± 4%   422.7ns ± 4%
Sort/input_size=30;-12        180ns ± 2%    228ns ± 5%    18ns ± 1%    24ns ± 2%     595ns ± 1%
Sort/input_size=1000;-12     16.2µs ± 6%    3.3µs ± 2%   0.6µs ± 2%   0.8µs ± 0%    54.3µs ± 2%
Sort/input_size=10000;-12     236µs ± 2%     38µs ±17%     6µs ± 1%     8µs ± 1%     708µs ± 3%
Sort/input_size=100000;-12   3.14ms ± 5%   0.41ms ± 5%  0.06ms ± 1%  0.08ms ± 2%    8.27ms ± 2%
Sort/input_size=250000;-12   8.30ms ± 4%   1.29ms ±15%  0.15ms ± 3%  0.20ms ± 4%   21.84ms ± 2%
Sort/input_size=1000000;-12                4.41ms ± 2%                            103.12ms ± 3%
```

A kimenet táblázatos formában:

|     **name \ time/op**      |    **comb.txt**    | **counting.txt** |   **bubble.txt**   |   **gnome.txt**    | **heap.txt**  |
|:---------------------------:|:------------------:|:----------------:|:------------------:|:------------------:|:-------------:|
|   Sort/input\_size=10;-12   |    37.2ns ± 2%     |   52.0ns ± 3%    |     6.5ns ± 0%     |     8.1ns ± 2%     | 243.7ns ± 3%  |
|   Sort/input\_size=19;-12   |    91.4ns ±10%     |   79.9ns ± 3%    |    12.8ns ± 8%     |    15.9ns ± 4%     | 422.7ns ± 4%  |
|   Sort/input\_size=30;-12   |     180ns ± 2%     |    228ns ± 5%    |     18ns ± 1%      |     24ns ± 2%      |  595ns ± 1%   |
|  Sort/input\_size=1000;-12  |    16.2µs ± 6%     |    3.3µs ± 2%    |     0.6µs ± 2%     |     0.8µs ± 0%     |  54.3µs ± 2%  |
| Sort/input\_size=10000;-12  |     236µs ± 2%     |    38µs ±17%     |      6µs ± 1%      |      8µs ± 1%      |  708µs ± 3%   |
| Sort/input\_size=100000;-12 |    3.14ms ± 5%     |   0.41ms ± 5%    |    0.06ms ± 1%     |    0.08ms ± 2%     |  8.27ms ± 2%  |
| Sort/input\_size=250000;-12 |    8.30ms ± 4%     |   1.29ms ±15%    |    0.15ms ± 3%     |    0.20ms ± 4%     | 21.84ms ± 2%  |
| Sort/input_size=1000000;-12 | take too much time |   4.41ms ± 2%    | take too much time | take too much time | 103.12ms ± 3% |

A fenti táblázatban láthatóak az eredmények, azonban nem mindegyik algoritmus került bele, amelyekhez készült benchmark,
mivel másfajta adatokkal dolgoztak és így ezek nem hasonlíthatóak össze.

Go-ban egészen egyszerű benchmarkot írni, minden eszköz elérhető a hivatalos könyvtárban hozzá. Fontos, hogy ha
különböző benchmarkokat akarunk összehasonlítani, akkor konzisztens adatokat célszerű használni. 