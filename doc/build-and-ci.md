# Build keretrendszer és CI beüzemelése GitHub Actions használatával

## Go tool
A Go-ban írt programok buildeléséhez és teszteléséhez a `go` tool nyújt alapértelmezett megoldásokat. Ez a tool CI-ban is könnyen használható.
A `go test <.go fileok>`, illetve `go test <package név>` parancsok segítségével tudjuk a teszteket futtatni, a `go test -bench` paranccsal pedig a benchmarkokat. Emellett a `go` tool tud a tesztekhez coverage-et mutatni a `-cover` kapcsoló használatával, tud race condition-öket ellenőrizni a `-race` kapcsolóval és van fuzzing támogatása is.

## Repo szerkezetéből jövő komplexitások
A Go (sok nyelvhez hasonlóan) package-ekbe szervezi a kódot, amit pedig a mappaszerkezetnek reflektálnia kell. Azonban ebben a repoban sok független `main` package van (mindegyik példakód külön program saját belépési ponttal). Emiatt a `go test`-et nem sikerült felparaméterezni úgy, hogy magától megtalálja a teszteket, így egy segéd scriptet kellett írni (`scripts/all_go_mains.sh`). Ez a script megkeresi az összes `main` package mappát, majd ezekben futtatja a megadott parancsokat. Ezt a scriptet hívja meg a GitHub Actions workflow.

## GitHub Actions
A weben a GitHub tutorial-ját követve összeraktam egy workflow-t, ami egy Ubuntu alapú gépre letölti a Go toolt, letölti az esetleges korábbi fordítások cache-ét, majd lefordítja, teszteli, benchamrkolja a programot, illetve a race detection-t és coverage-et is nézi. Ez a workflow a `.github/workflows/test.yml` fájlban van leírva és minden push-ra lefut.