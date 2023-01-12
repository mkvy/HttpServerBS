package customerrepo

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mkvy/HttpServerBS/internal/config"
	"github.com/mkvy/HttpServerBS/internal/utils"
	"github.com/mkvy/HttpServerBS/model"
	"log"
	"strings"
	"time"
)

type DBCustomerRepo struct {
	db  *sql.DB
	cfg config.Config
}

func NewPGCustomerRepository(cfg config.Config) (*DBCustomerRepo, error) {
	log.Println("Creating DB connection in DB Customer repository with driver " + cfg.Database.DriverName)
	connStr := "user=" + cfg.Database.Username + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.DBname + " sslmode=disable"
	fmt.Println(connStr)
	db, err := sql.Open(cfg.Database.DriverName, connStr)
	if err != nil {
		log.Println("Error while connecting to db")
		log.Println(err)
		return nil, utils.ErrDbConnect
	}
	if db == nil {
		log.Println("Error with database")
		panic("missing db")
	}
	return &DBCustomerRepo{db: db, cfg: cfg}, nil
}

func (repo *DBCustomerRepo) Close() error {
	err := repo.db.Close()
	return err
}

func (repo *DBCustomerRepo) Create(data model.Customer) (string, error) {
	log.Println("DBCustomerRepo Creating record")

	sqlStatement := `INSERT INTO customers VALUES ($1,$2,$3,$4,$5,$6);`
	id := uuid.New()
	res, err := repo.db.Exec(sqlStatement, id, data.Surname, data.Firstname, data.Patronym, data.Age, time.Now())
	if err != nil {
		log.Println("DBCustomerRepo Error occured while creating record: ", err)
		return "", err
	}
	rowsC, _ := res.RowsAffected()
	log.Println("rows affected: ", rowsC)
	log.Println("Created record with id: " + id.String())
	return id.String(), nil
}

func (repo *DBCustomerRepo) GetById(id string) (model.Customer, error) {
	log.Println("DBCustomerRepo GetByID record id: ", id)

	sqlStatement := `SELECT * from customers where id=$1;`

	row := repo.db.QueryRow(sqlStatement, id)
	if row == nil {
		log.Println("DBCustomerRepo Error occured while GetById record: ", utils.ErrNotFound)
		return model.Customer{}, utils.ErrNotFound
	}
	customer := model.Customer{}
	var id1 string
	err := row.Scan(&id1, &customer.Surname, &customer.Firstname, &customer.Patronym, &customer.Age, &customer.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("DBCustomerRepo Not found by id record: ", utils.ErrNotFound)
			return model.Customer{}, utils.ErrNotFound
		}
		log.Println("DBCustomerRepo GetById Error occured while scan record: ", err)
		return model.Customer{}, err
	}
	return customer, nil
}

func (repo *DBCustomerRepo) Update(data model.Customer, id string) error {
	log.Println("DBCustomerRepo Update record id: ", id)
	sqlStatement := `UPDATE customers`
	cntFields := 0
	if data.Surname != "" {
		cntFields++
		sqlStatement += ` SET surname='` + data.Surname + `',`
	}
	if data.Firstname != "" {
		cntFields++
		sqlStatement += ` SET firstname='` + data.Firstname + `',`
	}
	if data.Patronym != "" {
		cntFields++
		sqlStatement += ` SET patronym='` + data.Patronym + `',`
	}
	if data.Age != "" {
		cntFields++
		sqlStatement += ` SET age='` + data.Age + `',`
	}
	if data.DateCreated != nil {
		cntFields++
		tempt := *data.DateCreated
		sqlStatement += ` SET date_created='` + tempt.String() + `',`
	}
	if cntFields == 0 {
		log.Println("DBCustomerRepo Wrong data given to update")
		return utils.ErrWrongEntity
	}
	sqlStatement = strings.TrimSuffix(sqlStatement, ",")
	sqlStatement += ` where id=$1;`
	fmt.Println(sqlStatement)
	res, err := repo.db.Exec(sqlStatement, id)
	if err != nil {
		log.Println("DBCustomerRepo Error occured while updating record: ", err)
		return err
	}
	rowsC, err := res.RowsAffected()
	if err != nil {
		log.Println("DBCustomerRepo Error occured after updating record: ", err)
		return err
	}
	if rowsC == 0 {
		log.Println("Record not found ID: ", id)
		return utils.ErrNotFound
	}
	log.Println("rows affected: ", rowsC)
	return nil
}

func (repo *DBCustomerRepo) Delete(id string) error {
	log.Println("DBCustomerRepo Delete record id: ", id)
	res, err := repo.db.Exec(`delete from customers where id = $1;`, id)
	if err != nil {
		log.Println("DBCustomerRepo Error occured while deleting record: ", err)
		return err
	}
	rowsC, _ := res.RowsAffected()
	if rowsC == 0 {
		log.Println("Record not found ID: ", id)
		return utils.ErrNotFound
	}
	return nil
}

func (repo *DBCustomerRepo) GetBySurname(surname string) (model.Customer, error) {
	log.Println("DBCustomerRepo GetBySurname record")
	sqlStatement := `select * from customers where surname=$1;`
	row := repo.db.QueryRow(sqlStatement, surname)
	if row == nil {
		log.Println("DBCustomerRepo Error occured while GetById record: ", utils.ErrNotFound)
		return model.Customer{}, utils.ErrNotFound
	}
	customer := model.Customer{}
	var id1 string
	err := row.Scan(&id1, &customer.Surname, &customer.Firstname, &customer.Patronym, &customer.Age, &customer.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("DBCustomerRepo Not found by id record: ", utils.ErrNotFound)
			return model.Customer{}, utils.ErrNotFound
		}
		log.Println("DBCustomerRepo GetById Error occured while scan record: ", err)
		return model.Customer{}, err
	}
	return customer, nil
}
