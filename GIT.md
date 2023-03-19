# git


## add

comando in locale che aggiunge una modifica dalla working directory alla staging area di Git, una specie di lista di attesa con la raccolta delle modifiche che verranno inserite nel prossimo commit, con lo scopo finale di caricare le modifiche su Github

- `git add percorso`

- `git add .` la cartella . è la cartella corrente

- `git add nomefile` 


## commit

L'uso precendete del comando git add è necessario per selezionare le modifiche che verranno inserite per il commit successivo.

Commit è un comando in locale che crea un'istantanea delle modifiche effettuate nella cronologia di un progetto Git, applica quindi in modo effettivo le modifiche fatte con la git add

- `git commit` apre un editor di testo (il mio di default è nano)
- `git commit -m "messaggio"`

> la sintassi carina è: _verbo: descrizione delle modifiche effettuate_
>> update/add/delete: nomediqualcosa


## push 

comando che carica e pubblica le modifiche della repository locale sulla repository centrale di Github (remoto)

- `git push`


## pull

comando che scarica le modifiche fatte sulla repository di Github (remoto), è infatti usato per aggiornare la versione locale di una repository dal remoto

- `git pull`

⚠️ prima di iniziare a modificare il progetto lanciare sempre il comando git pull


## git checkout


## git merge
