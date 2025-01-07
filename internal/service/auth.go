package service

import (
	"context"
	"fmt"

	pb "corezn/api/proto"
	"corezn/internal/models"

	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	db *gorm.DB
}

func NewAuthServer(db *gorm.DB) *AuthServer {
	return &AuthServer{
		db: db,
	}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Validate input
	if err := validateRegisterRequest(req); err != nil {
		return nil, err
	}

	// Hash password
	argon := argon2.DefaultConfig()
	hash, err := argon.HashEncoded([]byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Start transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find or create organization
	org := &models.Organization{
		ID:          uuid.New(),
		VATNumber:   req.VatNumber,
		CompanyName: &req.CompanyName,
	}

	if req.CompanyName == "" {
		org.CompanyName = nil
	}

	result := tx.Where("vat_number = ? AND (company_name = ? OR (company_name IS NULL AND ? = ''))",
		req.VatNumber, req.CompanyName, req.CompanyName).
		FirstOrCreate(org)
	
	if result.Error != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to find/create organization: %v", result.Error)
	}

	// Create user
	user := &models.User{
		ID:             uuid.New(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Password:       string(hash),
		OrganizationID: org.ID,
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &pb.RegisterResponse{
		UserId:  user.ID.String(),
		Message: "User registered successfully",
	}, nil
}

func validateRegisterRequest(req *pb.RegisterRequest) error {
	if req.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}
	if req.LastName == "" {
		return fmt.Errorf("last_name is required")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	if req.VatNumber == "" {
		return fmt.Errorf("vat_number is required")
	}
	return nil
}