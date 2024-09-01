package mysql

import (
  "context"
  "errors"
  "reflect"
  "time"

  "go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
  Collection(string) Collection
  CLient() CLient
}

type Collection interface {
  FindOne(context.Context, interface{}) SingleResult
  InsertOne(context.Context, interface{}) (interface{}, error)
  InsertMany(context.Context, []interface{}) ([]interface{}, error)
  Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
} 

type SingleResult interface {
  Decode(interface{}) error
}

type Cursor interface {
  Close(context.Context) error
  Next(context.Context) bool
  Decode(interface{}) error
  All(context.Context, interface{}) error
}

type Client interface {
  Database(string) Database
  Connect(context.Context) error
  Disconnect(context.Context) error
  StartSession() (mongo.Session, error)
  UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
  Ping(context.Context) error
}

type mongoClient struct {
  cl *mongo.CLient
}

type mongoDatabase struct {
  db *mongo.Database
}

type mongoCollection struct {
  coll *mongo.Collection
}

type mongoSingleResult struct {
  sr *mongo.SingleResult
}

type mongoCursor struct {
  mc *mongo.Cursor
}

type mongoSession struct {
  mongo.Session
}

type nullawareDecoder struct {
  defDecoder bsoncodec.ValueDecoder
  zeroValue reflect.Value
}

func (d *nullawareDecoder) DecodedValue(dctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
  if vr.Type() != bsontype.Null {
    return d.defDecoder.DecodedValue(dctx, vr, val)
  }

  if !val.CanSet() {
    return errors.New("value not settable")
  }

  if err := vr.ReadNull(); err != nil {
    return err
  }

  val.Set(d.zeroValue)
  return nil
}

func NewClient(connection string) (client, error) {
  time.Local = time.UTC
  c, err := mongo.NewClient(options.Client().ApplyURI(connection))

  return &mongoClient{cl: c}, err
}

func (mc *mongoClient) Ping(ctx, context.Context) error {
  return mc.clPing(ctx, readpref.Primary())
}

func (mc *mongoClient) Database(dbName string) Database {
  db := mc.cl.Database(dbName)
  return &mongoDatabase(db: db)
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.cl.Connect(ctx)
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	findResult, err := mc.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: findResult}, err
}

func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	aggregateResult, err := mc.coll.Aggregate(ctx, pipeline)
	return &mongoCursor{mc: aggregateResult}, err
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter, opts...)
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}

func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.mc.All(ctx, result)
}
