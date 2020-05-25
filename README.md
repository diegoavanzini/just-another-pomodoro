Just Another Pomodoro made with [Go](https://golang.org/), [fyne](https://fyne.io/) toolkit, with [boltdb](https://github.com/boltdb/bolt)

```sh
go install -ldflags -H=windowsgui
```

### TODO
- ~~gestire la pausa del pomodoro~~
- ~~alert prima della fine della progress bar~~
- ~~add tests~~
- ~~position of tomatoes done at the correct hour~~ 
- ~~bundle data and images with bynary~~ 
- implement settings windows:
    1. ~~set timer duration~~ 
    2. ~~set break timer duration~~
    3. ~~continuous pomodoro~~
    4. stop at every pomodoro
    5. insert comment on every pomodoro
- ~~sync tomato with remote pair~~
- ~~add possibility to change listerener port~~
- disable pause button when in sync
- ~~validate insert ip~~
- ~~manage errors~~
- change storage from json to bolt(?)
- refactoring
- change log system
- made tests
- sync tomato with github commit
---
![alt text](img/jap2020-05-0817-15-27.png "screenshot")

