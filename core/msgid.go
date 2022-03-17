package core

type msgid struct {
	HANDSHAKE           int32
	ECHO                int32
	ERROR               int32
	SPECIAL_SERVER_LIST int32
	GAME_SERVER_LIST    int32
}

var MsgID = msgid{
	HANDSHAKE:           10001,
	ECHO:                10002,
	ERROR:               10003,
	SPECIAL_SERVER_LIST: 10004,
	GAME_SERVER_LIST:    10005,
}
