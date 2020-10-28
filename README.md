# nmc-typhoon-db-client

从 NMC 台风数据库检索台风报文

## 编译

使用 `make` 命令编译，生成的 `nmc_typhoon_db_client` 可执行程序保存在 bin 目录下。

## 使用

`nmc_typhoon_db_client get` 命令从数据库中检索台风信息并保存到 CSV 文件中。

下面命令检索 2020 年 10 月 25 日 00 时数据，将结果保存到 test.csv 文件中。

```bash
nmc_typhoon_db_client \
    get \
    --config ./config.yaml \
    --start-time 2020102500 \
    --output-file ./test.csv
```

`--config` 指定配置文件路径，默认使用当前路径下的 `config.yaml` 文件。

`--start-time` 指定起始时次，格式 YYYYMMDDHH

`--output-file` 设置输出 CSV 文件的路径

## 更多参数

该命令还支持获取时间段内的数据。

`--end-time` 指定截止时次，默认为空。如果设置，则返回起报时次在 `[start_time, end_time]` 之间的数据。

`--forecast-hour` 指定预报时效，单位小时，默认为 `0`，也支持设置时效范围，例如 `0-120` 表示 0 到 120 小时。

下面命令检索 2020 年 10 月 25 日 00 时到 2020 年 10 月 26 日 00 时，预报时效在 0 到 120 小时的数据。

```bash
nmc_typhoon_db_client \
    get \
    --config ./config.yaml \
    --start-time 2020102500 \
    --end-time 2020102600 \
    --forecast-hour "0-120" \
    --output-file "./test.csv"
```

## 配置文件

配置文件包含数据库访问信息，格式如下：

```yaml
database:
  host: host ip
  database_name: database name
  table_name: table name

  driver: "mysql+pymysql"

  auth:
    user: user name
    password: password
```

## CSV 文件

`nmc_typhoon_db_client get` 命令生成的 CSV 文件包含表头。
列标题与数据库字段名称相同，请参看 [columns.go](./columns.go) 文件中的 `QueryColumns` 对象。

## LICENSE

Copyright &copy; 2020, Perilla Roc at nwpc-oper.

`nmc-typhoon-db-client` is licensed under [MIT License](LICENSE)