namespace go rpc

enum ErrorCode{
    Success=0
    UnknowError=5000,
    VerifyError=5001,
    UserIsNull=5002,
    AddUserFail=5005,
    UserBeKickOff=5006,
    GoldNotEnough=5007,
    DiamondNotEnough=5009,
}

struct UserInfo{
     1: i64 userId
     2: string userName
     3: string nickName
     4: string avatarAuto
     5: i64 gold //游戏金币
     6: string token
}

struct Result{
    1:  ErrorCode code
    2: UserInfo user_obj
}

service UserService {

    Result createNewUser(1: string nickName 2:string avatarAuto 3: i64 gold )//初始金币

    //获取用户信息 BY userId
    Result getUserInfoById(1:i32 userId)

    //获取用户信息 BY token
    Result getUserInfoByken(1:string token)

    //修改用户金币
    Result modifyGoldById(1:string behavior, 2:i32 userId, 3:i64 gold)
    Result modifyGoldByToken(1:string behavior, 2:string token,3:i64 gold)
}
