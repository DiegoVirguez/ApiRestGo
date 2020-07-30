package users

import (
	"../../../../domain/model"
	"../../../database_client"
	//"github.com/stretchr/testify/assert"
	//"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	//"github.com/jinzhu/gorm"
)

var (
	userMysqlRepository UserMysqlRepository
)
/*
func TestUserMysqlRepository_SaveProducto(t *testing.T) {
	tx := userMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var producto model.Producto
	producto, _ = producto.CreateProducto("Prueba1",3000,"caracteristicas de la prueba")
	err := userMysqlRepository.SaveProducto(&producto)

	assert.Nil(t, err)
	assert.EqualValues(t, producto.Nombre, "Prueba1", "Nombres son diferentes")
	assert.NotEqual(t, producto.Precio, 3000)
	assert.NotEqual(t, producto.Caracteristicas, "caracteristicas de la prueba")
	assert.NotNil(t, producto.Codigo, "codigo no nulo")
}
*/
var _ = Describe("save", func() {
        var producto *model.Producto
        var repository *UserMysqlRepository
	    var mock sqlmock.Sqlmock
        BeforeEach(func() {
				var err error

				db, mock, err = sqlmock.New() // mock sql.DB
				Expect(err).ShouldNot(HaveOccurred())

				repository := &UserMysqlRepository{
		                Db: database_client.GetDatabaseInstance(),
	            }

                producto = &model.Producto{
                        Nombre : "Prueba1",
                        Precio : 300,
                        Caracteristicas : "Caracteres",
                }
        })

        It("insert", func() {
                // gorm use query instead of exec
                // https://github.com/DATA-DOG/go-sqlmock/issues/118
                const sqlInsert = `
                                INSERT INTO "Productos" ("nombre","precio","caracteristicas") 
                                        VALUES ($1,$2,$3) RETURNING "Productos"."codigo"`
                const codigo = 1
                mock.ExpectBegin() // begin transaction
                mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
                        WithArgs(producto.Nombre, producto.Precio, producto.Caracteristicas).
                        WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(codigo))
                mock.ExpectCommit() // commit transaction

                Expect(producto.Codigo).Should(BeZero())

                err := repository.SaveProducto(producto)
                Expect(err).ShouldNot(HaveOccurred())

                Expect(producto.Codigo).Should(BeEquivalentTo(codigo))
        })
})
/*func TestUserMysqlRepository_Get(t *testing.T) {

	tx := userMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser("Franklin", "Carrero", "mauriciocarrero15@gmail.com", "sistemas31")
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userMysqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}
	user, err := userMysqlRepository.Get(userDb.ID)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userDb.ID, user.Id)
	assert.EqualValues(t, "Carrero", user.LastName)
}*/