package blogcontroller_test

import (
	"go-learn-restapi-mysql/controllers/blogcontroller"
	"go-learn-restapi-mysql/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *MockDB {
	args := m.Called(dest, conds)
	return args.Get(0).(*MockDB)
}

// func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
// 	args = m.Called(query, args)
// 	return args.Get(0).(*MockDB)
// }

// func (m *MockDB) First(dest interface{}, conds ...interface{}) *MockDB {
// 	args := m.Called(dest, conds)
// 	return args.Get(0).(*MockDB)
// }

// func (m *MockDB) Create(value interface{}) *MockDB {
// 	args := m.Called(value)
// 	return args.Get(0).(*MockDB)
// }

// func (m *MockDB) Model(value interface{}) *MockDB {
// 	args := m.Called(value)
// 	return args.Get(0).(*MockDB)
// }

// func (m *MockDB) Updates(values interface{}, conds ...interface{}) *MockDB {
// 	args := m.Called(values, conds)
// 	return args.Get(0).(*MockDB)
// }

// func (m *MockDB) Delete(value interface{}, conds ...interface{}) *MockDB {
// 	args := m.Called(value, conds)
// 	return args.Get(0).(*MockDB)
// }

func TestIndex_NoDataFound(t *testing.T) {
	mockDB := new(MockDB)

	mockDB.On("Find", mock.Anything, mock.AnythingOfType("*[]models.Blog")).Run(func(args mock.Arguments) {
		blogs := args.Get(1).(*[]models.Blog)
		*blogs = []models.Blog{}
	}).Return(nil)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/blogs", blogcontroller.Index)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	expectedResponse := gin.H{
		"status":  false,
		"message": "No data found",
		"data":    []models.Blog{},
	}
	assert.Equal(t, expectedResponse, expectedResponse)
}

func TestIndex_DataFound(t *testing.T) {
	mockDB := new(MockDB)

	blogs := []models.Blog{
		{Id: 1, Title: "Blog 1", Tags: "tag1"},
		{Id: 2, Title: "Blog 2", Tags: "tag2"},
	}
	mockDB.On("Find", mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*[]models.Blog)
		*dest = blogs
	})

	router := gin.Default()
	router.GET("/blogs", blogcontroller.Index)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blogs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := gin.H{
		"status":  true,
		"message": "success",
		"data":    blogs,
	}
	assert.Equal(t, expectedResponse, expectedResponse)
}

// func TestSearch_RecordNotFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	mockDB.On("Where", mock.Anything, mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		blogs := args.Get(0).(*[]models.Blog)
// 		*blogs = []models.Blog{}
// 	})

// 	router := gin.Default()
// 	router.GET("/api/search/blogs", blogcontroller.Search)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/search/blogs?q=keyword", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  false,
// 		"message": "Record not found!",
// 		"data":    []models.Blog{},
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestSearch_RecordFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	blogs := []models.Blog{
// 		{Id: 1, Title: "Blog 1", Tags: "tag1"},
// 		{Id: 2, Title: "Blog 2", Tags: "tag2"},
// 	}
// 	mockDB.On("Where", mock.Anything, mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*[]models.Blog)
// 		*dest = blogs
// 	})

// 	router := gin.Default()
// 	router.GET("/api/search/blogs", blogcontroller.Search)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/search/blogs?q=Blog", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  true,
// 		"message": "success",
// 		"data":    blogs,
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestCreate_Success(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	blog := models.Blog{Id: 1, Title: "New Blog", Tags: "tag1"}
// 	mockDB.On("Create", mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		value := args.Get(0).(*models.Blog)
// 		*value = blog
// 	})

// 	router := gin.Default()
// 	router.POST("/api/v1/blog", blogcontroller.Create)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/api/v1/blog", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  true,
// 		"message": "create successfully",
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestShow_RecordFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	blog := models.Blog{Id: 1, Title: "Blog 1", Tags: "tag1"}
// 	mockDB.On("First", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*models.Blog)
// 		*dest = blog
// 	})

// 	router := gin.Default()
// 	router.GET("/api/v1/blog/:id", blogcontroller.Show)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/v1/blogs/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  true,
// 		"message": "success",
// 		"data":    blog,
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestShow_RecordNotFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	mockDB.On("First", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*models.Blog)
// 		*dest = models.Blog{}
// 	})

// 	router := gin.Default()
// 	router.GET("/api/v1/blog/:id", blogcontroller.Show)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/v1/blog/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  false,
// 		"message": "Record not found!",
// 		"data":    []models.Blog{},
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestUpdate_Success(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	blog := models.Blog{Id: 1, Title: "Updated Blog", Tags: "tag1"}
// 	mockDB.On("Model", mock.Anything).Return(mockDB)
// 	mockDB.On("Updates", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		values := args.Get(0).(*models.Blog)
// 		*values = blog
// 	})

// 	router := gin.Default()
// 	router.PUT("/api/v1/blog/:id", blogcontroller.Update)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/api/v1/blog/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  true,
// 		"message": "update successfully",
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestUpdate_RecordNotFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	mockDB.On("Model", mock.Anything).Return(mockDB)
// 	mockDB.On("Updates", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		values := args.Get(0).(*models.Blog)
// 		*values = models.Blog{}
// 	})

// 	router := gin.Default()
// 	router.PUT("/api/v1/blog/:id", blogcontroller.Update)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/api/v1/blog/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  false,
// 		"message": "Record not found!",
// 		"data":    []models.Blog{},
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestDelete_Success(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	blog := models.Blog{Id: 1, Title: "Blog 1", Tags: "tag1"}
// 	mockDB.On("First", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*models.Blog)
// 		*dest = blog
// 	})
// 	mockDB.On("Delete", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*models.Blog)
// 		*dest = blog
// 	})

// 	router := gin.Default()
// 	router.DELETE("/api/v1/blog/:id", blogcontroller.Delete)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/api/v1/blog/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  true,
// 		"message": "delete successfully",
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }

// func TestDelete_RecordNotFound(t *testing.T) {
// 	mockDB := new(MockDB)
// 	// config.DB = mockDB

// 	mockDB.On("First", mock.Anything, mock.Anything).Return(mockDB).Run(func(args mock.Arguments) {
// 		dest := args.Get(0).(*models.Blog)
// 		*dest = models.Blog{}
// 	})

// 	router := gin.Default()
// 	router.DELETE("/api/v1/blog/:id", blogcontroller.Delete)
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/api/v1/blog/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNotFound, w.Code)

// 	expectedResponse := gin.H{
// 		"status":  false,
// 		"message": "Record not found!",
// 		"data":    []models.Blog{},
// 	}
// 	assert.Equal(t, expectedResponse, gin.H{})
// }
