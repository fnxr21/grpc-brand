package handler

import (
	"context"
	"log"

	// "github.com/fnxr21/brand/config"
	"github.com/fnxr21/brand/internal/entities"
	"github.com/fnxr21/brand/internal/repository"
	"github.com/fnxr21/brand/protobuf/golang_protobuf_brand"
)

type brandServices struct {
	golang_protobuf_brand.UnimplementedCrudServer
	RepoBrand repository.BrandService
}

func HandlerBrand(BrandRepository repository.BrandService) *brandServices {
	return &brandServices{RepoBrand: BrandRepository}
}
func (c *brandServices) Create(ctx context.Context, req *golang_protobuf_brand.ProtoBrand) (*golang_protobuf_brand.BrandResponse, error) {
	ok := entities.Brand{
		ID:   uint(req.ID),
		Name: req.Name,
		Year: uint(req.Year),
	}

	log.Println(ok, "check")

	c.RepoBrand.Create(ok)

	return &golang_protobuf_brand.BrandResponse{Brand: &golang_protobuf_brand.ProtoBrand{
		ID:   uint64(ok.ID),
		Name: ok.Name,
		Year: uint32(ok.Year),
	}}, nil
}
