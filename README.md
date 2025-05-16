## env

1. go version must more than 1.21

## kit-logger

1. zapx
2. logx with trace id
3. vipperx
4. gormx

   - 配置字段具体参数， 可以修改 `g.FieldOpts`
   - 配置 yaml 表关系: 生成 model 后就不应该修改(否则关联关系会丢失)
   - 为指定对象自定义模版接口:

     ```go
     info := g.Data[db.NamingStrategy.SchemaName("archived_ups_tag")] // 获取 yaml 表关系生成的model元数据
     g.ApplyInterface(func(upsTagInterface) {}, info.QueryStructMeta)
     ```
