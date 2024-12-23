package repository_test

import (
    "database/sql"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "github.com/RohithReddy35/go-angular/internal/models"
    "github.com/RohithReddy35/go-angular/internal/repository"
    _ "github.com/lib/pq"
)

var _ = Describe("UserRepository", func() {
    var (
        db   *sql.DB
        repo *repository.UserRepository
    )

    BeforeEach(func() {
        var err error
        db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=your_password dbname=users_db sslmode=disable")
        Expect(err).ToNot(HaveOccurred())

        repo = repository.NewUserRepository(db)
    })

    AfterEach(func() {
        db.Close()
    })

    Context("CRUD Operations", func() {
        It("should create a new user", func() {
            user := &models.User{
                UserName: "test_user",
                Email:    "test@example.com",
            }

            err := repo.CreateUser(user)
            Expect(err).ToNot(HaveOccurred())
            Expect(user.ID).To(BeNumerically(">", 0))
        })

        It("should fetch all users", func() {
            users, err := repo.GetAllUsers()
            Expect(err).ToNot(HaveOccurred())
            Expect(users).To(HaveLen(1)) // Assuming one user was created earlier
        })

        It("should fetch a user by ID", func() {
            user, err := repo.GetUserByID(1) // Replace with an actual ID
            Expect(err).ToNot(HaveOccurred())
            Expect(user).ToNot(BeNil())
        })

        It("should update a user", func() {
            user := &models.User{
                ID:       1,
                UserName: "updated_user",
                Email:    "updated@example.com",
            }

            err := repo.UpdateUser(user)
            Expect(err).ToNot(HaveOccurred())
        })

        It("should delete a user", func() {
            err := repo.DeleteUser(1) // Replace with an actual ID
            Expect(err).ToNot(HaveOccurred())
        })
    })
})
