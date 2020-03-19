# mongo-info
mongo 相关

### Docker Run Mongo
```
docker run -d --name mongo -p 27017-27019:27017-27019 --restart=always mongo
```

### Mongo 基础Shell命令
```
$./mongo
> show dbs   // 查看当前有哪些数据库
> db         // 当前是那个数据库
> use local  // 切换到local数据库
```
基础数据库

- admin： 从权限的角度来看，这是"root"数据库。要是将一个用户添加到这个数据库，这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器。

- local: 这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合

- config: 当Mongo用于分片设置时，config数据库在内部使用，用于保存分片的相关信息。

### 文档(Document  表中的数据):
BSON`{"site":"www.github.com", "name":"Github"}`
- 文档中的键/值对是有序的。
- 文档中的值不仅可以是在双引号里面的字符串，还可以是其他几种数据类型（甚至可以是整个嵌入的文档)。
- MongoDB区分类型和大小写。
- MongoDB的文档不能有重复的键。
- 文档的键是字符串。除了少数例外情况，键可以使用任意UTF-8字符。


### 基础操作
#### 数据库链接
`mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]`
#### 创建 删除 数据库 
```
use DatabaseName // 如果不存在就会创建
// 演示插入一些数据
db.TableName.insert({"name":"github"})

// 删除数据库
db.dropDatabase()
```
>注意: 在 MongoDB 中，集合只有在内容插入后才会创建! 就是说，创建集合(数据表)后要再插入一个文档(记录)，集合才会真正创建。
#### 集合操作
```
// 往集合中插入数据
db.TableName.insert({"name":"github"})
// 查看当前数据库中集合
show tables
// 创建集合
db.createColloection("tableName")
// 删除集合
db.集合Name.drop()
```
集合创建进阶
```
db.createCollection(name, options)
```
参数说明：
- name: 要创建的集合名称
- options: 可选参数, 指定有关内存大小及索引的选项

capped | 布尔 | （可选）如果为 true，则创建固定集合。固定集合是指有着固定 大小的集合，当达到最大值时，它会自动覆盖最早的文档。当该值为 true 时，必须指定 size 参数
-|-|-
autoIndexId | 布尔 | （可选）如为 true，自动在 _id 字段创建索引。默认为 false。 |
size | 数值 | （可选）为固定集合指定一个最大值，以千字节计（KB）。如果 capped 为 true，也需要指定该字段。 |
max  | 数值 | （可选）指定固定集合中包含文档的最大数量。 |

```
> db.createCollection("mycol", { capped : true, autoIndexId : true, size : 
   6142800, max : 10000 } )
{ "ok" : 1 }
```
#### 插入与查询
```
> db.github.insert({"name":"你爸爸还是你爸爸","message":"快叫爸爸"})
WriteResult({ "nInserted" : 1 })
> db.github.find()
{ "_id" : ObjectId("5e72e594d42fbd6445d94e9c"), "name" : "你大爷还是你大爷" }
{ "_id" : ObjectId("5e72eb57d42fbd6445d94e9e"), "name" : "你爸爸还是你爸爸", "message" : "快叫爸爸" }


#  插入单条数据

> var document = db.collection.insertOne({"a": 3})
> document
{
        "acknowledged" : true,
        "insertedId" : ObjectId("571a218011a82a1d94c02333")
}

#  插入多条数据
> var res = db.collection.insertMany([{"b": 3}, {'c': 4}])
> res
{
        "acknowledged" : true,
        "insertedIds" : [
                ObjectId("571a22a911a82a1d94c02337"),
                ObjectId("571a22a911a82a1d94c02338")
        ]
}
```
#### 更新文档
```
db.collection.update(
   <query>,
   <update>,
   {
     upsert: <boolean>,
     multi: <boolean>,
     writeConcern: <document>
   }
)
```
参数说明：
- query : update的查询条件，类似sql update查询内where后面的。
- update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
- upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
- multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
- writeConcern :可选，抛出异常的级别。
```
> db.github.update({'message':'快叫爸爸'},{$set:{'name':'我是你爸爸'}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.github.find().pretty()
{ "_id" : ObjectId("5e72e594d42fbd6445d94e9c"), "name" : "你大爷还是你大爷" }
{
	"_id" : ObjectId("5e72eb57d42fbd6445d94e9e"),
	"name" : "我是你爸爸",
	"message" : "快叫爸爸"
}

// 更新多条记录
>db.col.update({'title':'MongoDB 教程'},{$set:{'title':'MongoDB'}},{multi:true})
```

#### Save() 传入新文档进行替换
```
db.collection.save(
   <document>,
   {
     writeConcern: <document>
   }
)
```
参数说明：
- document : 文档数据。
- writeConcern :可选，抛出异常的级别。
```
// 以下实例中我们替换了 _id 为 56064f89ade2f21f36b03136 的文档数据：
>db.col.save({
    "_id" : ObjectId("56064f89ade2f21f36b03136"),
    "title" : "MongoDB",
    "description" : "MongoDB 是一个 Nosql 数据库",
    "tags" : [
            "mongodb",
            "NoSQL"
    ],
    "likes" : 110
})
```
More:
```
只更新第一条记录：

db.col.update( { "count" : { $gt : 1 } } , { $set : { "test2" : "OK"} } );
全部更新：

db.col.update( { "count" : { $gt : 3 } } , { $set : { "test2" : "OK"} },false,true );
只添加第一条：

db.col.update( { "count" : { $gt : 4 } } , { $set : { "test5" : "OK"} },true,false );
全部添加进去:

db.col.update( { "count" : { $gt : 5 } } , { $set : { "test5" : "OK"} },true,true );
全部更新：

db.col.update( { "count" : { $gt : 15 } } , { $inc : { "count" : 1} },false,true );
只更新第一条记录：

db.col.update( { "count" : { $gt : 10 } } , { $inc : { "count" : 1} },false,false );
```

#### 删除文档
```
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
```
- query :（可选）删除的文档的条件。
- justOne : （可选）如果设为 true 或 1，则只删除一个文档，如果不设置该参数，或使用默认值 false，则删除所有匹配条件的文档。
- writeConcern :（可选）抛出异常的级别。
```
>db.col.insert({title: 'MongoDB 教程', 
    description: 'MongoDB 是一个 Nosql 数据库',
    url: 'http://www.runoob.com',
    tags: ['mongodb', 'database', 'NoSQL'],
    likes: 100
})
>db.col.insert({title: 'MongoDB 教程', 
    description: 'MongoDB 是一个 Nosql 数据库',
    url: 'http://www.runoob.com',
    tags: ['mongodb', 'database', 'NoSQL'],
    likes: 100
})
>db.col.remove({'title':'MongoDB 教程'})

```