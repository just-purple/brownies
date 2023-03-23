# comandi rapidi da tastiera

## control+c

- _control c_ su terminale termina il comando in esecuzione


# comandi da terminale

## code

- `code .` apre la cartella corrente su vs code

## cd

- `cd percorso` campia directory


## go

- `go run percorso` esegue


# git


## add

comando in locale che aggiunge una modifica dalla working directory alla staging area di Git, una specie di lista di attesa con la raccolta delle modifiche che verranno inserite nel prossimo commit, con lo scopo finale di caricare le modifiche su Github

- `git add percorso`

- `git add .` la cartella corrente

- `git add nomefile` 


## commit

L'uso precendete del comando git add è necessario per selezionare le modifiche che verranno inserite nel commit successivo.

Commit è un comando in locale che crea un'istantanea delle modifiche effettuate nella cronologia di un progetto Git, applica quindi in modo effettivo le modifiche fatte con la git add

- `git commit` apre un editor di testo (il mio di default è nano)
- `git commit -m "messaggio"`

### Git Commit Messages Style-Guides

https://gist.github.com/rishavpandey43/84665ffe3cea76400d8e5a1ad7133a79


## push 

comando che carica e pubblica le modifiche della repository locale sulla repository centrale di Github (remoto)

- `git push`


## pull

comando che scarica le modifiche fatte sulla repository di Github (remoto), è infatti usato per aggiornare la versione locale di una repository dal remoto

- `git pull`

⚠️ prima di iniziare a modificare il progetto lanciare sempre il comando git pull


## git checkout


## git merge
