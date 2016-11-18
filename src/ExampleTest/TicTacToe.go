package main 

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"net"
)

const(
	X_CHAR = 'X'
	O_CHAR = 'O'
	EMPTY_CHAR = '\000'
)

type TicTacToe struct{
	gameMatrix [3][3]byte
	hasWon bool
	playAsX bool
	winner byte
}



func (ticTacToe TicTacToe) printGameMatrix(){
	fmt.Printf("%c\n",ticTacToe.gameMatrix[0])
	fmt.Printf("%c\n",ticTacToe.gameMatrix[1])
	fmt.Printf("%c\n",ticTacToe.gameMatrix[2])
}

func (ticTacToe *TicTacToe) hasSomeoneWon() bool {
	var characterToUse byte
	if ticTacToe.playAsX{
		characterToUse = X_CHAR
	}else{
		characterToUse = O_CHAR
	}
	return ticTacToe.hasWonVertical(characterToUse) || ticTacToe.hasWonHorizontal(characterToUse) || ticTacToe.hasWonFirstDiagonal(characterToUse) || ticTacToe.hasWonSecondDiagonla(characterToUse)
}

func (ticTacToe *TicTacToe) isFull() bool{
	for  i := 0; i < len(ticTacToe.gameMatrix); i++ {
            for j := 0; j < len(ticTacToe.gameMatrix[i]); j++ {
                if ticTacToe.gameMatrix[i][j] == EMPTY_CHAR {
                  return false
                }
            }
        }
	return true
}

func (ticTacToe *TicTacToe) doMovement(cordX, cordY int) bool{
	var characterToUse byte
	if ticTacToe.playAsX{
		characterToUse = X_CHAR
	}else{
		characterToUse = O_CHAR
	}

	if ticTacToe.gameMatrix[cordX][cordY] == EMPTY_CHAR{
		fmt.Println("moved!")
		ticTacToe.gameMatrix[cordX][cordY] = characterToUse
		fmt.Println("Printing inside do movement")
		ticTacToe.printGameMatrix()
		ticTacToe.playAsX = !ticTacToe.playAsX
		return true
	}else{
		ticTacToe.playAsX = ticTacToe.playAsX
		return false
	}
}

func genRand(min,max int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}

func (ticTacToe TicTacToe) hasWonHorizontal(character byte) bool{
	counter := 0
        for  i := 0; i < len(ticTacToe.gameMatrix); i++ {
            for j := 0; j < len(ticTacToe.gameMatrix[i]); j++ {
                if ticTacToe.gameMatrix[i][j] == character {
                    counter++;
                    if counter == 3 {
                    	ticTacToe.winner = character
                        return true;
                        
                    }
                } else {
                    counter = 0;
                    break;
                }
            }
        }
        return false;
}

func (ticTacToe TicTacToe) hasWonVertical(character byte) bool{
	counter := 0
	 for i := 0; i < len(ticTacToe.gameMatrix); i++ {
            for j := 0; j < len(ticTacToe.gameMatrix[i]); j++ {
                if ticTacToe.gameMatrix[j][i] == character {
                    counter++;
                    if (counter == 3) {
                    	ticTacToe.winner = character
                        return true;
                    }
                } else {
                    counter = 0;
                    break;
                }
            }
        }
	 return false
}

func (ticTacToe TicTacToe) hasWonFirstDiagonal(character byte) bool{
	if ticTacToe.gameMatrix[0][0] == character && ticTacToe.gameMatrix[1][1] == character && ticTacToe.gameMatrix[2][2] == character {
		ticTacToe.winner = character
		return true
	}else{
		return false
	}
}

func (ticTacToe TicTacToe) hasWonSecondDiagonla(character byte) bool{
	if ticTacToe.gameMatrix[0][2] == character && ticTacToe.gameMatrix[1][1] == character && ticTacToe.gameMatrix[2][0] == character{
		ticTacToe.winner = character
		return true
	}else{
		return false
	}
}

func main() {
	ticTacToe := &TicTacToe{}
	//ticTacToe.gameMatrix = [3][3]byte{{X_CHAR,O_CHAR,X_CHAR},{O_CHAR,O_CHAR,O_CHAR},{X_CHAR,O_CHAR}}
	//ticTacToe.setGameMatrix(foo
	for !ticTacToe.hasSomeoneWon(){
		cords := rand.Intn(9)
		if ticTacToe.doMovement(cords/3, cords%3){
			//fmt.Println("Yay bro! this is the fucking matrix")
			//ticTacToe.printGameMatrix()
		}
		time.Sleep(1*time.Second)
	}
	fmt.Println("Winner is player ",ticTacToe.winner)
	
}

