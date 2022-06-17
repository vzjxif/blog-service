package errcode

var (
	Success                   = NewError(0, "æå")
	ServerError               = NewError(10000000, "æå¡åé¨éè¯¯")
	InvalidParams             = NewError(10000001, "å¥åéè¯¯")
	NotFound                  = NewError(10000002, "æ¾ä¸å°")
	UnauthorizedAuthNotExist  = NewError(10000003, "é´æå¤±è´¥ï¼æ¾ä¸å°å¯¹åºç AppKey å AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "é´æå¤±è´¥ï¼Token éè¯¯")
	UnauthorizedTokenTimeout  = NewError(10000005, "é´æå¤±è´¥ï¼Token è¶æ¶")
	UnauthorizedTokenGenerate = NewError(10000006, "é´æå¤±è´¥ï¼Token çæå¤±è´¥")
	TooManyRequests           = NewError(10000007, "è¯·æ±è¿å¤")
)
