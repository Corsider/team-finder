package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"team-finder/domain"
	"team-finder/internal/utils"
)

type Database interface {
	Reference() *sql.DB
	SelectAllFromX(X string) (*sql.Rows, error)
	SelectAllFromXWhereYeqZ(X, Y, Z string) (*sql.Rows, error)
	SelectXFromYWhereZeqN(X, Y, Z, N string) (*sql.Rows, error)
	Select1FromXWhereYeqZ(X, Y, Z string) *sql.Row
	SelectXFromYWhereZeqNorMeqC(X, Y, Z, N, M, C string) (*sql.Rows, error)
	DeleteFromXWhereYeqZ(X, Y, Z string) error
	InsertIntoXYValuesZ(X, Y, Z string) error
	InsertParametrizedIntoXYValuesZ(X, Y, Z string, params ...interface{}) error
	InsertIntoXYValuesZReturningN(X, Y, Z, N string) (interface{}, error)
	InsertParametrizedIntoXYValuesZReturningN(X, Y, Z, N string, params ...interface{}) (interface{}, error)
	SelectCountFromXWhereYeqZ(X, Y string, params ...interface{}) (int, error)
	SelectCountFromXWhereYeqZorNeqM(X, Y, Z string, params ...interface{}) (int, error)
	UpdateXSetYZWhereNeqM(X, Y, Z, N, M string, decodeTo interface{}) (interface{}, error)
	UpdateNXSetYZWhereNeqM(X, N, M string, columns []string, decodeTo interface{}, params ...interface{}) (interface{}, error)
	SelectAllFromXWhereYinZandNinMandGinHOrderByO(X, Y, Z, N, M, G, H, O string, myTeam int, tags []int, from, to int, asc bool) (*sql.Rows, error)
	SelectAllFromXNinMandGinHOrderByO(X, N, M, G, H, O string, myTeam int, tags []int, from, to int, asc bool) (*sql.Rows, error)
}

type PostgresDB struct {
	DB *sql.DB
}

func (db *PostgresDB) Reference() *sql.DB {
	return db.DB
}

func NewConnection(connection string) Database {
	database, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresDB{DB: database}
}

/////////// IMPORTANT /////////////
///////////////////////////////////
///// Not final, test version /////
///////////////////////////////////
///////////////////////////////////

func (db *PostgresDB) SelectAllFromX(X string) (*sql.Rows, error) {
	//query := "select * from $1"
	return db.DB.Query("select * from " + X)
}

func (db *PostgresDB) SelectAllFromXWhereYeqZ(X, Y, Z string) (*sql.Rows, error) {
	//query := "select * from $1 where $2=$3"
	//return db.DB.Query(query, X, Y, Z)
	return db.DB.Query("select * from " + X + " where " + Y + "=" + Z)
}

func (db *PostgresDB) SelectXFromYWhereZeqN(X, Y, Z, N string) (*sql.Rows, error) {
	//query := "select $1 from $2 where $3=$4"
	return db.DB.Query("select " + X + " from " + Y + " where " + Z + "=" + N)
}

func (db *PostgresDB) Select1FromXWhereYeqZ(X, Y, Z string) *sql.Row {
	//query := "select * from $1 where $2=$3"
	return db.DB.QueryRow("select * from " + X + " where " + Y + "=" + Z)
}

func (db *PostgresDB) SelectXFromYWhereZeqNorMeqC(X, Y, Z, N, M, C string) (*sql.Rows, error) {
	//query := "select $1 from $2 where $3=$4 or $5=$6"
	return db.DB.Query("select " + X + " from " + Y + " where " + Z + "=" + N + " or " + M + "=" + C)
}

func (db *PostgresDB) DeleteFromXWhereYeqZ(X, Y, Z string) error {
	//query := "delete from $1 where $2=$3"
	return utils.Second(db.DB.Exec("delete from " + X + " where " + Y + "=" + Z))
}

func (db *PostgresDB) InsertIntoXYValuesZ(X, Y, Z string) error {
	//query := "insert into $1($2) values ($3)"
	return utils.Second(db.DB.Exec("insert into " + X + " (" + Y + ") values (" + Z + ")"))
}

func (db *PostgresDB) InsertParametrizedIntoXYValuesZ(X, Y, Z string, params ...interface{}) error {
	query := "insert into " + X + " (" + Y + ") values (" + Z + ")"
	_, err := db.DB.Exec(query, params...)
	return err
}

func (db *PostgresDB) InsertIntoXYValuesZReturningN(X, Y, Z, N string) (interface{}, error) {
	//query := "insert into $1($2) values ($3) returning $4"
	var returned interface{}
	err := db.DB.QueryRow("insert into " + X + " (" + Y + ") values (" + Z + ") returning " + N).Scan(&returned)
	return returned, err
}

func (db *PostgresDB) InsertParametrizedIntoXYValuesZReturningN(X, Y, Z, N string, params ...interface{}) (interface{}, error) {
	//query := "insert into $1($2) values ($3) returning $4"
	var returned interface{}
	query := "insert into " + X + " (" + Y + ") values (" + Z + ") returning " + N
	err := db.DB.QueryRow(query, params...).Scan(&returned)
	return returned, err
}

func (db *PostgresDB) SelectCountFromXWhereYeqZ(X, Y string, params ...interface{}) (int, error) {
	query := "select count(*) as count from " + X + " where " + Y + "=$1"
	var count int
	err := db.DB.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *PostgresDB) SelectCountFromXWhereYeqZorNeqM(X, Y, Z string, params ...interface{}) (int, error) {
	var count int
	query := "select count(*) as count from " + X + " where " + Y + "=$1 or " + Z + "=$2"
	err := db.DB.QueryRow(query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *PostgresDB) UpdateXSetYZWhereNeqM(X, Y, Z, N, M string, decodeTo interface{}) (interface{}, error) {
	query := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE %s = $2 RETURNING *", X, Y, N)
	row := db.DB.QueryRow(query, Z, M)
	switch v := decodeTo.(type) {
	case *domain.User:
		err := row.Scan(&v.UserId, &v.Name, &v.Nickname, &v.Rate, &v.Description, &v.Login, &v.Password)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func (db *PostgresDB) UpdateNXSetYZWhereNeqM(X, N, M string, columns []string, decodeTo interface{}, params ...interface{}) (interface{}, error) {
	pairs := make([]string, len(columns))
	for i, column := range columns {
		pairs[i] = fmt.Sprintf("%s = $%d", column, i+2)
	}
	log.Println(pairs)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $1 RETURNING *", X, strings.Join(pairs, ", "), N)

	params = append([]interface{}{M}, params...)

	row := db.DB.QueryRow(query, params...)
	switch v := decodeTo.(type) {
	case domain.User:
		err := row.Scan(&v.UserId, &v.Name, &v.Nickname, &v.Rate, &v.Description, &v.Login, &v.Password)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func ArrayToString(a []int, delim string) string {
	return "{" + strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]") + "}"
}

func (db *PostgresDB) SelectAllFromXWhereYinZandNinMandGinHOrderByO(X, Y, Z, N, M, G, H, O string, myTeam int, tags []int, from, to int, asc bool) (*sql.Rows, error) {
	query := "select * from %s where %s in (%s) and (%s) in (%s) and (%s) between %s order by %s"
	if asc {
		query += " ASC"
	} else {
		query += " DESC"
	}
	log.Println(fmt.Sprintf(query, X, Y, Z, N, M, G, H, O, myTeam, ArrayToString(tags, ","), from, to))
	rows, err := db.DB.Query(query, X, Y, Z, N, M, G, H, O, myTeam, ArrayToString(tags, ","), from, to)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *PostgresDB) SelectAllFromXNinMandGinHOrderByO(X, N, M, G, H, O string, myTeam int, tags []int, from, to int, asc bool) (*sql.Rows, error) {
	query := "select * from %s where (%s) in (%s) and (%s) between %s order by %s"
	if asc {
		query += " ASC"
	} else {
		query += " DESC"
	}
	//log.Println(fmt.Sprintf(query, X, N, M, G, H, O, myTeam, ArrayToString(tags, ","), from, to))
	rows, err := db.DB.Query(query, X, N, M, G, H, O, myTeam, ArrayToString(tags, ","), from, to)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
