package core

type msgid struct {
	HANDSHAKE                   int32
	ECHO                        int32
	ERROR                       int32
	SPECIAL_SERVER_LIST         int32
	GAME_SERVER_LIST            int32
	GM_ID_START                 int32
	GM_GET_USER                 int32
	GM_SET_USER                 int32
	GM_GET_USER_THEME           int32
	GM_SET_USER_THEME           int32
	GM_SEND_PRIZE               int32
	GM_CHEAT                    int32
	GM_CLEAN_DISCONNECTION_DATA int32
	GM_GET_THEME                int32
	GM_KICK_PLAYER              int32
	GM_GET_ROOM                 int32
	GM_SEND_MAIL                int32
	GM_SET_ROOM                 int32
	GM_SET_GS_GROUP             int32
	GM_RESET_RANK               int32
	GM_GET_RANK                 int32
	GM_GET_RANK_ROOM            int32
	RK_ID_START                 int32
}

var MsgID = msgid{
	HANDSHAKE:                   10001,
	ECHO:                        10002,
	ERROR:                       10003,
	SPECIAL_SERVER_LIST:         10004,
	GAME_SERVER_LIST:            10005,
	GM_ID_START:                 11000,
	GM_GET_USER:                 11001,
	GM_SET_USER:                 11002,
	GM_GET_USER_THEME:           11011,
	GM_SET_USER_THEME:           11012,
	GM_SEND_PRIZE:               11013,
	GM_CHEAT:                    11014,
	GM_CLEAN_DISCONNECTION_DATA: 11015,
	GM_GET_THEME:                11016,
	GM_KICK_PLAYER:              11017,
	GM_GET_ROOM:                 11018,
	GM_SEND_MAIL:                11019,
	GM_SET_ROOM:                 11020,
	GM_SET_GS_GROUP:             11021,
	GM_RESET_RANK:               11022,
	GM_GET_RANK:                 11023,
	GM_GET_RANK_ROOM:            11024,
	RK_ID_START:                 12000,
}
