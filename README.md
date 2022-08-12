# Mongo-KRNK ![status](https://badgen.net/badge/alpha/passing/green?icon=github) ![latest release](https://badgen.net/github/release/cobearcoding/mongo-krnk)

![language](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![mongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)

Is a simple implementation of the [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver for ease of use inside different personal projects.

Considerations:
> * This library is not a replacement for **MongoDB Official Driver**
>
> * This library occupies **sync.Once** feature, so handling the connection is only done once. It is still advised to use the library as explained in the **examples**.
>
>
> This library is currently in development, major changes could be released regularly.
>
> Usage on production environments is on your own discretion.

# Installation
```terminal
 go get github.com/cobearcoding/mongo-krnk
 ```
---
### Please make sure to have installed:

* Go v1.18 +
* MongoDB Go Driver v1.10.1 +

# Usage
To use this library you should use environmental variables on your **.env** file.

Example:
> MONGO_URL=mongodb://user:password@host:port/?authSource=admin

# Important
This library uses **structs** for most of the operations, please be sure to be familiar with this data type before using this library.

## Quering a Document
---
For document query this library uses different **structs** depending the case.

### FindAll
The structure provided is just an example to know the type of data you should be loading in the **struct** to use.

This query returns a result of type **bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.

``` go
    type MongoQuery struct {
        MongoURI   string
        Database   string
        Collection string
        Page       int64
        PerPage    int64
    }
```

> Usage Example:

```go
    query := orm.Mongo.Query{
        MongoURI:   os.Getenv("MONGO_URI")
		Database:   os.Getenv("MONGO_DATABASE"),
		Collection: "example",
		Page:       0,
		PerPage:    10,
	}

    result, err := query.FindAll()
	if err != nil {
		return err
	}
```


### FindOne
The structure provided is just an example to know the type of data you should be loading in the **struct** to use.

This query returns a result of type **bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.
``` go
    type MongoQuery struct {
        MongoURI   string
        Database   string
        Collection string
        Page       int64
        PerPage    int64
    }
```

> Usage Example:

```go
    query := orm.Mongo.Query{
        MongoURI:   os.Getenv("MONGO_URI")
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
The structure provided is just an example to know the type of data you should be loading in the **struct** to use.

This query returns an array result of type **[]bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.

``` go
    type MongoQuery struct {
        MongoURI   string
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
        MongoURI:   os.Getenv("MONGO_URI")
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

## Inserting a Document
For document insertion this library uses different **structs** depending the case.

### InsertOne
The structure provided is just an example to know the type of data you should be loading in the **struct** to use.

This type of document will need a object of type **bson.D** that contains de **Key** and **Value** of the data you will want to insert.

This function returns an error of type **error** if anything went wrong during the transaction.

```go
    type MongoInsert struct {
        MongoURI   string
        Database   string
        Collection string
        Value      bson.D
    }
```

> Usage Example:

```go

    dataInsert := bson.D{
		{Key: "example_uid", Value: exampleInput.UID},
	}

    mongoInsert := orm.MongoInsert{
        MongoURI:   os.Getenv("MONGO_URL"),
        Database:   os.Getenv("MONGO_DATABASE"),
        Collection: "example",
        Value:      dataInsert,
    }

    err := mongoInsert.InsertOne()
    if err != nil {
       return err
    }
```

### Updating a Document
The structure provided is just an example to know the type of data you should be loading in the **struct** to use.

This type of document will need a object of type **bson.D** that contains de **Key** and **Value** of the data you will want to update.

This function returns an error of type **error** if anything went wrong during the transaction.

```go

    type MongoUpdate struct {
        MongoURI    string
        Database    string
        Collection  string
        FilterKey   string
        FilterValue interface{}
        Value       bson.D
    }
```
> Usage Example:

```go 
    update := bson.D{
		{Key: "example_name", Value: exampleInput.Name},
	}

	updateQuery := orm.MongoUpdate{
		MongoURI:    os.Getenv("MONGO_URL"),
		Database:    os.Getenv("MONGO_DATABASE"),
		Collection:  "example",
		FilterKey:   "example_uid",
		FilterValue: exampleInput.UID,
		Value:       update,
	}

	err := updateQuery.UpdateOne()
	if err != nil {
		return err
	}
```

## Raw Queries
Raw queries are donde for those who need extra flexibility at the time of creating a custom query or additional aggregations, etc.

### FindRaw
This function will consume a **struct** which will have less components but will allow for even more flexibility of finding documents with complex filters.

The **query** type will work as the **filter** type.

This function returns an array result of type **[]bson.M**,  for further information please refer to the  [mongoDB]("https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/upsert/") official driver documentation.

```go
    type MongoRawQuery struct {
        MongoURI   string
        Database   string
        Collection string
        Query      bson.D
    }
```

> Usage Example:

```go
    query := bson.D{
		{Key: "example_uid", Value: example.UID},
	}
	raw := orm.MongoRawQuery{
		MongoURI:   os.Getenv("MONGO_URL"),
		Database:   os.Getenv("MONGO_DATABASE"),
		Collection: "example",
		Query:      query,
	}
	results, err := raw.FindRaw()
	if err != nil {
		return err
	}
    return results
```

### UpdateRaw
This function will consume a **struct** which will have less components but will allow for even more flexibility of finding documents with complex filters.

This function will need the **$set** attribute to be explicitly typed in the **bson.D** object as in the example.

This function returns an error of type **error** if anything went wrong during the transaction.

```go

    type MongoRawUpdate struct {
        MongoURI   string
        Database   string
        Collection string
        Filter     bson.D
        Update     bson.D
    }
```

> Usage Example:

```go
    update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "example_name", Value: example.Name},
		}},
	}

	filter := bson.D{
		{Key: "example_uid", Value: example.UID},
	}

	rawUpdate := orm.MongoRawUpdate{
		MongoURI:   os.Getenv("MONGO_URL"),
		Database:   os.Getenv("MONGO_DATABASE"),
		Collection: "example",
		Filter:     filter,
		Update:     update,
	}

	err := rawUpdate.UpdateRaw()
	if err != nil {
		return err
	}

```