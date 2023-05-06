记录一些写项目时的问题
事务一般不写在model层，而是写在logic中如果我们想在logic中用事务，但是事务又只能在model中使用，解决办法是将开启事务暴露给logic，即在每个Model中加一个方法，注意还要在interface中加入定义

定义一个方法：

```go
func (m *defaultUsersModel) TransCtx(ctx context.Context, fn func(ctx context.Contex, session sqlx.Session) error) error {
    return m.TransactCtx(ctx,func(ctx context.Context,s sqlx.Session) error {
        return fn(ctx,s)
    })
}
```

在logic中的用法就是把你要添加事务的数据库操作放进TransCtx中

```go
if err := l.svcCtx.RolesChangeModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			if _, err := l.svcCtx.RolesChangeModel.Insert(l.ctx, rolesChange); err != nil {
				return err
			}
			if err := l.svcCtx.UsersModel.Update(l.ctx, userInfo); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return nil, response.Error(500, err.Error())
		}
```

一些关于redis的使用

打开终端，`redis-sercer.exe`打开redis（本地设置成一直运行），`redis-cli`打开redis客户端，然后`auth 123456`登录
