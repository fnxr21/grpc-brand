package server

import (
	"context"
	"log"

	"github.com/fnxr21/brand/config"
	"github.com/fnxr21/brand/entities"
	"github.com/fnxr21/brand/protobuf/golang_protobuf_brand"
	// "github.com/fnxr21/brand/repos"
)


type BrandService struct {
	golang_protobuf_brand.UnimplementedCrudServer
	// repo repos.BrandService
}

func (c *BrandService) Create(ctx context.Context, req *golang_protobuf_brand.ProtoBrand) (*golang_protobuf_brand.BrandResponse, error) {
	ok := entities.Brand{
		ID:   uint(req.ID),
		Name: req.Name,
		Year: uint(req.Year),
	}

	log.Println(ok, "check")

	// c.repo.Create(ctx, ok)
	// if err
	if err := config.DB.Create(&ok).Error; err != nil {
		log.Println("check sini")
		return nil, err
	}
	return &golang_protobuf_brand.BrandResponse{Brand: &golang_protobuf_brand.ProtoBrand{
		ID:   uint64(ok.ID),
		Name: ok.Name,
		Year: uint32(ok.Year),
	}}, nil
}
