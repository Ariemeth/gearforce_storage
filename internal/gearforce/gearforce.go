package gearforce

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ariemeth/gearforce_storage/internal/config"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/endpoints"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/errors"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"

	driver "github.com/arangodb/go-driver"
	ahttp "github.com/arangodb/go-driver/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	dbName      = "gearforce"
	colName     = "rosters"
	idNamespace = "011f52f8-24e4-4bb9-a9e2-e13b7fcac716"
)

var (
	namespace = uuid.MustParse(idNamespace)
)

type GearForceService interface {
	service.Hello
	service.SaveRoster
	service.GetRoster
}

type ServiceMiddleware func(GearForceService) GearForceService

func ConfigureRouteHandler(r *mux.Router, root string, dbConfig config.Database) {
	svc := newGearForceService(dbConfig)

	gfHelloHandler := endpoints.MakeHelloHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/hello", root), gfHelloHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/hello/", root), gfHelloHandler).Methods("POST")

	gfStoreHandler := endpoints.MakeSaveRosterHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/store", root), gfStoreHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/store/", root), gfStoreHandler).Methods("POST")

	gfGetHandler := endpoints.MakeGetRosterHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/{id}", root), gfGetHandler).Methods("GET")
	r.Handle(fmt.Sprintf("%s/{id}/", root), gfGetHandler).Methods("GET")
}

type gearForceService struct {
	DBConfig   config.Database
	Collection driver.Collection
}

func newGearForceService(dbConfig config.Database) GearForceService {
	svc := gearForceService{
		DBConfig: dbConfig,
	}

	// Establish a connection to the primary database
	conn, err := ahttp.NewConnection(ahttp.ConnectionConfig{
		Endpoints: []string{dbConfig.Address},
	})
	if err != nil {
		log.Fatalf("Error connecting to database %s: %s", dbConfig.Address, err.Error())
		return nil
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(dbConfig.User, dbConfig.Password),
	})
	if err != nil {
		log.Fatalf("Error creating client for database: %s", err.Error())
		return nil
	}

	ctx := context.Background()

	// Open a connection to to the primary database
	exists, err := client.DatabaseExists(ctx, dbName)
	if err != nil {
		log.Fatalf("Error checking if db %s exists on %s: %v", dbName, dbConfig.Address, err)
	}

	var db driver.Database
	if !exists {
		db, err = client.CreateDatabase(ctx, dbName, nil)
		if err != nil {
			log.Fatalf("Error creating database %s on %s: %v", dbName, dbConfig.Address, err)
		}
		log.Printf("Created database %s on %s", dbName, dbConfig.Address)
	} else {
		db, err = client.Database(ctx, dbName)
		if err != nil {
			log.Fatalf("Error opening database %s on %s: %v", dbName, dbConfig.Address, err)
		}
	}

	// Open a connection to the gearforce collection in the database
	exists, err = db.CollectionExists(ctx, colName)
	if err != nil {
		log.Fatalf("Error checking if collection %s exists in the database %s on %s: %v", colName, dbName, dbConfig.Address, err)
	}

	var col driver.Collection
	if !exists {
		col, err = db.CreateCollection(ctx, colName, nil)
		if err != nil {
			log.Fatalf("Error creating collection %s in database %s on %s: %s", colName, dbName, dbConfig.Address, err)
		}
		log.Printf("Created collection %s in %s on %s", colName, dbName, dbConfig.Address)
	} else {
		col, err = db.Collection(ctx, colName)
		if err != nil {
			log.Fatalf("Error connecting to the collection %s in the database %s on %s: %v", colName, dbName, dbConfig.Address, err)
		}
	}

	log.Printf("Connection to collection %s in %s on %s successful", colName, dbName, dbConfig.Address)
	svc.Collection = col

	return &svc
}

func (gearForceService) Hello(s string) (string, error) {

	return "hello", nil
}

// SaveRoster implements GearForceService.
func (g *gearForceService) SaveRoster(r models.Roster) (uuid.UUID, error) {
	key, _ := generateRosterID(r)

	sr := r.ToRosterStorage(key.String())

	exists, err := g.Collection.DocumentExists(context.Background(), key.String())
	if err != nil {
		log.Printf("Error checking of roster id %s already exists: %v", key.String(), err)
		return uuid.Nil, errors.ErrFromDatabase
	}
	if exists {
		return key, nil
	}

	meta, err := g.Collection.CreateDocument(context.Background(), sr)
	if err != nil {
		log.Printf("Error creating doc in database: %v", err)
		return uuid.Nil, errors.ErrCannotCreateEntry
	}

	log.Printf("Created document in collection '%s' in database '%s' with meta ID '%s' and meta key '%s'\n", colName, dbName, meta.ID, meta.Key)

	return key, nil
}

// GetRoster implements GearForceService.
func (g *gearForceService) GetRoster(id uuid.UUID) (models.Roster, error) {
	exists, err := g.Collection.DocumentExists(context.Background(), id.String())
	if err != nil {
		log.Printf("Error getting roster %s from database: %v", id.String(), err)
		return models.Roster{}, errors.ErrFromDatabase
	}

	if !exists {
		return models.Roster{}, errors.ErrIdNotFound
	}

	var result models.RosterStorage
	_, err = g.Collection.ReadDocument(context.Background(), id.String(), &result)
	if err != nil {
		log.Printf("Error getting roster %s from database: %v", id.String(), err)
		return models.Roster{}, errors.ErrFromDatabase
	}

	return models.Roster{RosterBase: result.RosterBase}, nil
}

func generateRosterID(r models.Roster) (uuid.UUID, error) {
	u, err := json.Marshal(r.RosterBase)
	if err != nil {
		return uuid.New(), err
	}

	id := uuid.NewSHA1(namespace, u)
	return id, nil
}
