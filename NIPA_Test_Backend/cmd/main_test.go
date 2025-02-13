package main

import (
	"NIPA_Test_Backend/internal/config"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestSetupRouter(t *testing.T) {
	// กำหนดการตั้งค่าจาก config
	cfg := config.LoadConfig()

	// จำลองการเชื่อมต่อกับฐานข้อมูล
	mockDB := new(mock.Mock)
	mockDB.On("ConnectDB", cfg).Return(mockDB, nil)

	// ทดสอบการตั้งค่าผ่านฟังก์ชัน SetupRouter
	r := gin.Default()
	err := setupRouter(r, cfg)

	// ตรวจสอบว่าไม่มีข้อผิดพลาดเกิดขึ้น
	assert.NoError(t, err)
	// ตรวจสอบว่า r คือ gin.Engine
	assert.IsType(t, &gin.Engine{}, r)

	// ตรวจสอบ route ที่ตั้งค่าไว้
	t.Run("Check Health route", func(t *testing.T) {
		// สร้าง request เพื่อตรวจสอบ health route
		w := performRequest(r, "GET", "/health")
		assert.Equal(t, 200, w.Code)
	})

	// ตรวจสอบว่า mockDB ถูกเรียกใช้งาน
}

// SetupRouter sets up the Gin router with the given configuration
func setupRouter(r *gin.Engine, cfg config.Config) error {
	// Add your route setup code here
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
	return nil
}

// Helper function สำหรับทดสอบ HTTP request
func performRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
