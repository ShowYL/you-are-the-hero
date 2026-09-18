# TP - Programmation concurrente

## Installation du compilateur go

> Si vous avez déjà go installer sur votre machine, vous pouvez passer cette étape.
> Vous pouvez vérifier que vous possédez go en executant la commande `go version`

### 1. Télécharger go sous format archive

Pour télécharger le compilateur go en **archive** :
- Linux (x86-64) : https://go.dev/dl/go1.27.1.linux-amd64.tar.gz
- Windows (x86-64) : https://go.dev/dl/go1.27.1.windows-amd64.zip
- MacOS (arm64) : https://go.dev/dl/go1.27.1.darwin-arm64.pkg

> Pour installer go d'une autre manière : https://go.dev/dl/

### 2. Décompressez le répertoire `go` dans le répertoire courant du TP (où se trouve le fichier README.md)
### 3. Vérifier que go fonctionne en tapant la commande `./go/bin/go version` en étant dans le répertoire courant du TP

## Lancer les différentes démonstrations

Ces codes de démonstration permettent de comprendre rapidement la syntaxe et le multithreading en go.

1. Hello world  
   `./go/bin/go run ./demo-go/1-hello.go`
2. Variables  
   `./go/bin/go run ./demo-go/2-variable.go`
3. Boucles  
   `./go/bin/go run ./demo-go/3-boucles.go`
4. Sans concurrence  
   `./go/bin/go run ./demo-go/4-sans-concu.go`
5. Concurrence  
   `./go/bin/go run ./demo-go/5-concu.go`
6. Sémaphores (mutex)  
   `./go/bin/go run ./demo-go/6-mutex.go`

## Exercices

L'objectif des exercices est de corriger le mauvais comportement de programmation concurrente de chaque code go.
Testez le comportement sans correction du code pour visualiser le mauvais comportement du code.
Ensuite modifier le code pour rectifier le comportement du programme.

1. Exercice 1  
   `./go/bin/go run ./tp/1-exo.go`
2. Exercice 2  
   `./go/bin/go run ./tp/2-exo.go`
3. Exercice 3  
   `./go/bin/go run ./tp/3-exo.go`
4. Exercice 4  
   `./go/bin/go run ./tp/4-exo.go`

### Bonus

Pour ceux qui sont à l'aise, trois exercices un peu plus difficiles.

5. Bonus 1  
   `./go/bin/go run ./tp/5-bonus.go`
6. Bonus 2  
   `./go/bin/go run ./tp/6-bonus.go`
7. Bonus 3 (canaux, voir la doc en tête du fichier)  
   `./go/bin/go run ./tp/7-bonus.go`

## A noter

Par abus de langage nous allons appeler les coroutines de go des threads pour une question de simplification et de compréhension.

## Documentation go

Pour ceux qui souhaitent plus d'information sur le langage de programmation golang voici les liens de documentations :
https://go.dev/tour/list