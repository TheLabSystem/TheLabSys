# TheLabSys

# For Backend

The project uses gin and gorm.
For gorm,see https://learnku.com/docs/gorm/v2 for more.


# For Code

```txt
OK                 ErrNo = 0
ParamInvalid       ErrNo = 1 // 参数不合法
UserHasExisted     ErrNo = 2 // 该 Username 已存在
UserNotExisted     ErrNo = 4 // 该 Username 用户不存在
WrongPassword      ErrNo = 5 // 密码错误
LoginRequired      ErrNo = 6 // 用户未登录
PermDenied         ErrNo = 7 // 没有操作权限
VerifyCodeNotValid ErrNo = 8 // 验证码不正确（新用户注册）
StudentNotExist    ErrNo = 9 // 学生不存在（老师维护学生列表）
MoneyNotEnough     ErrNo = 10 // 付款

UnknownError ErrNo = 255 // 未知错误
```


# For Device

```txt
Status = 2//设备可用
Status = 1//设备不可用
Status = -1//设备已损坏
```

# For Bill

```txt
Status = 2//账单未支付
Status = 1//账单已支付
Status = -1//账单已取消
```
# For UserType

```txt
UserType = 1 // 外来人员
UserType = 2 // 学生
UserType = 3 // 老师
UserType = 4 // 设备管理员
UserType = 5 // 外部接口（财务处）

UserType = 255 // 负责人
```