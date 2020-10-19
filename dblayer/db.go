package dblayer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/vrindavanstudio/haldiayatra/utils"

	//"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	//"go.mongodb.org/mongo-driver/bson"
)

//MongoCreds credentials for mongodb connection
type MongoCreds struct {
	MongodbUser       string
	MongodbPassword   string
	MongoURI          string
	MongoDBName       string
	MembersCollection string
}

var (
	//MongoConfig configuration
	MongoConfig *MongoCreds
	//DBCredsFilePath ...
	DBCredsFilePath = "credentials.json"
)

var once sync.Once

//
//var client *mongo.Client
type DBClient struct {
	*mongo.Client
}


//GetDBConfig getting the mongodb credentials
func GetDBConfig() *MongoCreds {
	if MongoConfig != nil {
		return MongoConfig
	}

	file, err := ioutil.ReadFile(DBCredsFilePath)
	if err != nil {
		fmt.Println("FIle error", err)
	}
	MongoConfig = &MongoCreds{}
	json.Unmarshal(file, MongoConfig)

	MongoConfig.MongoURI = utils.GetEnv("MONGODB_URI", MongoConfig.MongoURI)
	MongoConfig.MongoDBName = utils.GetEnv("MONGODB_NAME", MongoConfig.MongoDBName)
	MongoConfig.MembersCollection = utils.GetEnv("MEMBERS_COLLECTION", MongoConfig.MembersCollection)

	fmt.Println("$$$$$$$$$$$$$$$$", MongoConfig)
	return MongoConfig
}

//GetConnection gives the mongo client
func GetConnection() (*DBClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var client *mongo.Client

	once.Do(func()  {
		MongoConfig := GetDBConfig()
		//uri := "mongodb+srv://vrinda:xdj3FB3MWasn1y2s@cluster0.ocshf.mongodb.net/haldiayatra?retryWrites=true&w=majority"

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(MongoConfig.MongoURI))
		})
		if err != nil {
			fmt.Println("mongo connect errrorrr")
			//panic(err)
			return nil, err
		}
		


	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged...")

    return &DBClient{
		Client: client,
	}, err
		//return client
	//
}

//CloseConnection closes the db connection
func (client *DBClient) CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	fmt.Println("connection closed")

}
