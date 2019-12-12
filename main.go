package main

import "fmt"

type Duck interface {
	Duckwalk() string
}

type Dog interface {
	Bark() string
}

type Meca interface {
	Weeenween() string
}

type Chimera struct {
	soundOfWalking string
	soundOfBarking string
	soundOfMeca    string
}

func (c Chimera) Duckwalk() string {
	return c.soundOfWalking
}

func (c Chimera) Bark() string {
	return c.soundOfBarking
}

func (c Chimera) Weeenween() string {
	return c.soundOfMeca
}

func LetDuckWalk(duck Duck) string {
	return duck.Duckwalk()
}

func LetDogBark(dog Dog) string {
	return dog.Bark()
}

func LetMecaWeen(meca Meca) string {
	return meca.Weeenween()
}

func main() {
	chimera := &Chimera{
		"ぴょこぴょこ",
		"わんわん",
		"がしゃがしゃ",
	}
	fmt.Printf("これがキメラだ! -> %+v\n", chimera)

	fmt.Println(
		"アヒルなら歩いてみせよ！",
		LetDuckWalk(chimera),
		"犬なら吠えてみせよ！",
		LetDogBark(chimera),
		"メカなら",
		LetMecaWeen(chimera),
	)
}

/*
	reference
	- http://otiai10.hatenablog.com/entry/2014/01/15/210445
*/
