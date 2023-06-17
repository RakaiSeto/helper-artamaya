package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "arta01:Eliandri3!@tcp(54.169.162.177:3306)/wadb")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing

	lyric := `Saturday sunset
	We're lying on my bed with five hours to go
	Fingers entwined and so were our minds
	Cryin', "I don't want you to go"
	You wiped away your tears
	But not fears under the still and clear indigo
	You said, "Baby, don't cry, we'll be fine
	You're the one thing I swear I can't outgrow"
	My mother said the younger me was a pretending prodigy
	Well, nothing then, much has changed
	'Cause while you're wolfin' down liquor
	My soul, it gets sicker
	But I'm stickin' to the screenplay
	Gotta say I'm okay, but answer this, babe
	How is it now that somehow you're a strangеr?
	You were mine just yеsterday
	I pray the block in my airway dissipates
	And instead deters your airplane's way
	But heaven denied
	Destiny decried
	Something beautiful died
	Too soon
	But I'm letting go
	I'm givin' up the ghost
	But don't get me wrong
	I'll always love you, that's why
	I wrote you this very last song
	I guess this is where we say goodbye
	I know I'll be alright
	Someday I'll be fine
	But just not tonight (ooh)
	Plungin' into all kinds of diversions
	Like blush wine and sonorous soirées
	But even with gin and surgin' adrenaline
	I see you're all that can intoxicate
	Oceans and engines
	You're skilled at infringin' on great love affairs
	'Cause now my heart's home
	All I've known is long gone and ten thousand miles away
	And I'm not okay
	But I'm letting go
	I'm giving up the ghost
	But don't get me wrong
	I'll always love you, that's why
	I wrote you this very last song
	I guess this is where we say goodbye
	I know I'll be alright
	But just not
	Tonight was the first time I stared into seas
	Of beguiling sepia two years ago
	And the first time I learned real world superpowers lived in three words
	They revitalize my fraying bones, oh
	Now what do you do when your pillar crumbled down
	You've lost all solid ground
	Both dreams and demons drowned
	And this void's all you've found
	And doubts light it aglow?
	I have so many questions
	But I'm pouring them into the ocean
	And I'm starting up my engine
	And I'm letting go
	I'm givin' up your ghost
	It's come to a close
	I marked the end with this last song I wrote
	I'm letting go
	This is the last falsetto
	I'll ever sing to you
	My great lost love`

	split := strings.Split(strings.ReplaceAll(lyric, "\r\n", "\n"), "\n")

	for i, v := range split {
		if v == "" {
			split = append(split[:i], split[i+1:]...)
		} else if v == " "{
			split = append(split[:i], split[i+1:]...)
		}
	}

	for i := 0; i < len(split); i++ {
		query := "INSERT INTO `self_message_in` (`message`) VALUES (?)"

		insertResult, err := db.ExecContext(context.Background(), query, split[i])
		if err != nil {
			log.Fatalf("fail to insert: %s", err)
		}	
		
		id, err := insertResult.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		log.Printf("inserted id: %d", id)
	}

	fmt.Println("DONE")

    defer db.Close()
}