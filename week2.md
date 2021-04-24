
### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

sql.ErrNoRows 在业务里几种可能的意义：

1. 在查询用户购买记录，该用户未曾购买过，查询记录为空为正常的用户情境，不应该返回 error
2. 列表分页查询api 里所查询的分页数超过了查询结果集的总数，此时应该返回空的记录数组，而不应该返回 error
    ``` go
    func QueryPage(conditions map[string]interface{}, page, per_page int) ([]*model.UserInfo, error) {
      var db *gorm.DB
      // db = GetDB()
      list := []*model.UserInfo{}
      err := db.Model(model.UserInfo{}).
        Where(conditions).
        Offset((page-1)*per_page).
        Limit(per_page).
        Scan(&list).Error
      if err == sql.ErrNoRows { // 如果查询连接没有错误,但返回空值,业务上并不意味着出错
        return []*model.UserInfo{}, nil
      } else {
        return list, err
      }
    }
    ```


3. 所查询数据库因误操作丢失数据，此时dao 方法查询本应返回对应数据而实际返回空值，此时 error 应该添加上下文信息（比参数、查询的 sql 语句，查询时的时间戳），通过 wrap 方法返回给调用者。
    ``` go
    func QueryUser(user_id int) (*model.Wbuser, error) {
      var db *gorm.DB
      user := &model.Wbuser{}
      query := db.Model(model.Wbuser{}).First(user, "id=?", user_id)
      if err := query.Error; err == sql.ErrNoRows {
        return nil, errors.Wrapf(err, "Error when querying %s, params: %d", query, user_id)
      } else if err != nil {
        return nil, err
      } else {
        return user, nil
      }
    }

    ```
