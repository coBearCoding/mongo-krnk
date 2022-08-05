# Mongo-KRNK
Is a simple implementation of the [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver for ease of use inside different personal projects.

Considerations:
> * This library is not a replacement for MongoDB Official Driver
>
> This library is currently in development, major changes could be released regularly.
>
> Usage on production environments is on your own discretion.

# Installation
> go get github.com/cobearcoding/mongo-krnk
---
### Please make sure to have installed:

* Go v1.18 +
* MongoDB Go Driver v1.10.1 +

# Usage
To use this library you should use environmental variables on your **.env** file. There is only one **reserverd** variable to be added to make sure the library works as expected.

Example:
> MONGO_URL=mongodb://user:password@host:port/?authSource=admin

# Important
This library uses **structs** for most of the operations, please be sure to be familiar with this data type before using this library.

## Quering a Document
---
For document query this library uses different **structs** depending the case.

### FindOne
The structure provided is just for example to know the type of data you should be loading in the **struct** to use.

This query returns a result of type **bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.
``` go
    type MongoQuery struct {
        Database   string
        Collection string
        Key        string
        Value      interface{}
        Page       int64
        PerPage    int64
    }
```

> Usage Example:

```go
    query := orm.MongoQuery{
		Database:   os.Getenv("MONGO_DATABASE"),
		Collection: "example",
		Key:        "example_id",
		Value:      exampleInput.UID,
		Page:       0,
		PerPage:    10,
	}

    result, err := query.FindOne()
	if err != nil {
		return err
	}
```

### Find
The structure provided is just for example to know the type of data you should be loading in the **struct** to use.

This query returns an array result of type **[]bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.

``` go
    type MongoQuery struct {
        Database   string
        Collection string
        Key        string
        Value      interface{}
        Page       int64
        PerPage    int64
    }
```

> Usage Example:

``` go
    mongoFind := orm.MongoQuery{
		Database:   os.Getenv("MONGO_DATABASE"),
		Collection: "example",
		Key:        "example_id",
		Value:      exampleInput.UID,
	}

	results, err := mongoFind.Find()
	if err != nil {
		return err
	}
```