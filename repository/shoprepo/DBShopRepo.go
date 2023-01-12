package shoprepo

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/mkvy/HttpServerBS/internal/utils"
	"github.com/mkvy/HttpServerBS/model"
	"log"
	"strconv"
	"strings"
)

type DBShopRepo struct {
	db *sql.DB
}

func NewDBShopRepository(dbConn *sql.DB) *DBShopRepo {
	return &DBShopRepo{db: dbConn}
}

func (repo *DBShopRepo) Create(data model.Shop) (string, error) {
	log.Println("DBShopRepo Creating record")

	sqlStatement := `INSERT INTO shops VALUES ($1,$2,$3,$4,$5);`
	id := uuid.New()
	res, err := repo.db.Exec(sqlStatement, id.String(), data.Name, data.Address, *data.WorkStatus, data.Owner)
	if err != nil {
		log.Println("DBShopRepo Error occured while creating record: ", err)
		return "", err
	}
	rowsC, _ := res.RowsAffected()
	log.Println("rows affected: ", rowsC)
	log.Println("Created record with id: " + id.String())
	return id.String(), nil
}

func (repo *DBShopRepo) GetById(id string) (model.Shop, error) {
	log.Println("DBShopRepo GetByID record id: ", id)

	sqlStatement := `SELECT * from shops where id=$1;`

	row := repo.db.QueryRow(sqlStatement, id)
	if row == nil {
		log.Println("DBShopRepo Error occured while GetById record: ", utils.ErrNotFound)
		return model.Shop{}, utils.ErrNotFound
	}
	customer := model.Shop{}
	var id1 string
	err := row.Scan(&id1, &customer.Name, &customer.Address, &customer.WorkStatus, &customer.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("DBShopRepo Not found by id record: ", utils.ErrNotFound)
			return model.Shop{}, utils.ErrNotFound
		}
		log.Println("DBShopRepo GetById Error occured while scan record: ", err)
		return model.Shop{}, err
	}
	return customer, nil
}

func (repo *DBShopRepo) Update(data model.Shop, id string) error {
	log.Println("DBShopRepo Update record id: ", id)
	sqlStatement := `UPDATE shops SET`
	cntFields := 0
	if data.Name != "" {
		cntFields++
		sqlStatement += ` shopname='` + data.Name + `',`
	}
	if data.Address != "" {
		cntFields++
		sqlStatement += ` address='` + data.Address + `',`
	}
	if data.WorkStatus != nil {
		cntFields++
		sqlStatement += ` work_status=` + strconv.FormatBool(*data.WorkStatus) + `,`
	}
	if data.Owner != "" {
		cntFields++
		sqlStatement += ` owner='` + data.Owner + `',`
	}
	if cntFields == 0 {
		log.Println("DBShopRepo Wrong data given to update")
		return utils.ErrWrongEntity
	}
	sqlStatement = strings.TrimSuffix(sqlStatement, ",")
	sqlStatement += ` where id=$1;`
	fmt.Println(sqlStatement)
	res, err := repo.db.Exec(sqlStatement, id)
	if err != nil {
		log.Println("DBShopRepo Error occured while updating record: ", err)
		return err
	}
	rowsC, err := res.RowsAffected()
	if err != nil {
		log.Println("DBShopRepo Error occured after updating record: ", err)
		return err
	}
	if rowsC == 0 {
		log.Println("Record not found ID: ", id)
		return utils.ErrNotFound
	}
	log.Println("rows affected: ", rowsC)
	return nil
}

func (repo *DBShopRepo) Delete(id string) error {
	log.Println("DBShopRepo Delete record id: ", id)
	res, err := repo.db.Exec(`delete from shops where id = $1;`, id)
	if err != nil {
		log.Println("DBShopRepo Error occured while deleting record: ", err)
		return err
	}
	rowsC, _ := res.RowsAffected()
	if rowsC == 0 {
		log.Println("Record not found ID: ", id)
		return utils.ErrNotFound
	}
	return nil
}

func (repo *DBShopRepo) GetByName(name string) (model.Shop, error) {
	log.Println("DBShopRepo GetByName record")
	sqlStatement := `select * from shops where shopname=$1;`
	row := repo.db.QueryRow(sqlStatement, name)
	if row == nil {
		log.Println("DBShopRepo Error occured while GetById record: ", utils.ErrNotFound)
		return model.Shop{}, utils.ErrNotFound
	}
	customer := model.Shop{}
	var id1 string
	err := row.Scan(&id1, &customer.Name, &customer.Address, &customer.WorkStatus, &customer.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("DBShopRepo Not found by id record: ", utils.ErrNotFound)
			return model.Shop{}, utils.ErrNotFound
		}
		log.Println("DBShopRepo GetById Error occured while scan record: ", err)
		return model.Shop{}, err
	}
	return customer, nil
}
