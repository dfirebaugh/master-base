package main

// #include <badge-ir-game-protocol.h>
import "C"
import "fmt"

func opCodes() {
	fmt.Println(C.BADGE_IR_GAME_ADDRESS)
	fmt.Println(C.OPCODE_SET_GAME_START_TIME)
	fmt.Println(C.OPCODE_SET_GAME_DURATION)
	fmt.Println(C.OPCODE_HIT)
	fmt.Println(C.OPCODE_SET_BADGE_TEAM)
	fmt.Println(C.OPCODE_REQUEST_BADGE_DUMP)
	fmt.Println(C.OPCODE_SET_GAME_VARIANT)
	fmt.Println(C.GAME_VARIANT_FREE_FOR_ALL)
	fmt.Println(C.GAME_VARIANT_TEAM_BATTLE)
	fmt.Println(C.GAME_VARIANT_ZOMBIE)
	fmt.Println(C.GAME_VARIANT_CAPTURE_THE_BADGE)
	fmt.Println(C.OPCODE_BADGE_RECORD_COUNT)
	fmt.Println(C.OPCODE_BADGE_UPLOAD_HIT_RECORD_BADGE_ID)
	fmt.Println(C.OPCODE_GAME_ID)
	fmt.Println(C.OPCODE_BADGE_UPLOAD_HIT_RECORD_TIMESTAMP)
}
