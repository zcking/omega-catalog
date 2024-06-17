package internal

import (
	"context"

	omega_catalogv1 "github.com/zcking/omega-catalog/gen/go/omega_catalog/v1"
	"github.com/zcking/omega-catalog/internal/database/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Impl) CreateCatalog(ctx context.Context, req *omega_catalogv1.CreateCatalogRequest) (*omega_catalogv1.CreateCatalogResponse, error) {
	comment := req.GetComment()
	catalog := models.Catalog{
		Name:       req.GetName(),
		Comment:    &comment,
		Properties: req.GetProperties(),
	}
	result := i.db.DB.Create(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	return &omega_catalogv1.CreateCatalogResponse{
		Catalog: &omega_catalogv1.Catalog{
			Id:         catalog.ID.String(),
			Name:       catalog.Name,
			Comment:    *catalog.Comment,
			CreatedAt:  timestamppb.New(catalog.CreatedAt),
			UpdatedAt:  timestamppb.New(catalog.UpdatedAt),
			Properties: catalog.Properties,
		},
	}, nil
}

func (i *Impl) ListCatalogs(ctx context.Context, req *omega_catalogv1.ListCatalogsRequest) (*omega_catalogv1.ListCatalogsResponse, error) {
	var catalogs []models.Catalog
	result := i.db.DB.Find(&catalogs)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert the database models to the protobuf response
	pbCatalogs := make([]*omega_catalogv1.Catalog, len(catalogs))
	for i, catalog := range catalogs {
		pbCatalogs[i] = &omega_catalogv1.Catalog{
			Id:         catalog.ID.String(),
			Name:       catalog.Name,
			Comment:    *catalog.Comment,
			CreatedAt:  timestamppb.New(catalog.CreatedAt),
			UpdatedAt:  timestamppb.New(catalog.UpdatedAt),
			Properties: catalog.Properties,
		}
	}
	return &omega_catalogv1.ListCatalogsResponse{
		Catalogs: pbCatalogs,
	}, nil
}

func (i *Impl) GetCatalog(ctx context.Context, req *omega_catalogv1.GetCatalogRequest) (*omega_catalogv1.GetCatalogResponse, error) {
	var catalog models.Catalog
	result := i.db.DB.Where("name = ?", req.GetName()).First(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	return &omega_catalogv1.GetCatalogResponse{
		Catalog: &omega_catalogv1.Catalog{
			Id:         catalog.ID.String(),
			Name:       catalog.Name,
			Comment:    *catalog.Comment,
			CreatedAt:  timestamppb.New(catalog.CreatedAt),
			UpdatedAt:  timestamppb.New(catalog.UpdatedAt),
			Properties: catalog.Properties,
		},
	}, nil
}

func (i *Impl) UpdateCatalog(ctx context.Context, req *omega_catalogv1.UpdateCatalogRequest) (*omega_catalogv1.UpdateCatalogResponse, error) {
	var catalog models.Catalog
	result := i.db.DB.Where("name = ?", req.GetName()).First(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	comment := req.GetComment()
	catalog.Comment = &comment
	catalog.Properties = req.GetProperties()
	result = i.db.DB.Save(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	return &omega_catalogv1.UpdateCatalogResponse{
		Catalog: &omega_catalogv1.Catalog{
			Id:         catalog.ID.String(),
			Name:       catalog.Name,
			Comment:    *catalog.Comment,
			CreatedAt:  timestamppb.New(catalog.CreatedAt),
			UpdatedAt:  timestamppb.New(catalog.UpdatedAt),
			Properties: catalog.Properties,
		},
	}, nil
}

func (i *Impl) DeleteCatalog(ctx context.Context, req *omega_catalogv1.DeleteCatalogRequest) (*omega_catalogv1.DeleteCatalogResponse, error) {
	var catalog models.Catalog
	result := i.db.DB.Where("name = ?", req.GetName()).First(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	result = i.db.DB.Delete(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}

	return &omega_catalogv1.DeleteCatalogResponse{}, nil
}
