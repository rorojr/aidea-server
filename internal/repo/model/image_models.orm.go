package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// ImageModelN is a ImageModel object, all fields are nullable
type ImageModelN struct {
	original        *imageModelOriginal
	imageModelModel *ImageModelModel

	Id           null.Int    `json:"id"`
	ModelId      null.String `json:"model_id,omitempty"`
	ModelName    null.String `json:"model_name,omitempty"`
	Vendor       null.String `json:"vendor,omitempty"`
	RealModel    null.String `json:"real_model,omitempty"`
	Meta         null.String `json:"-"`
	PreviewImage null.String `json:"preview_image,omitempty"`
	Description  null.String `json:"description,omitempty"`
	Status       null.Int    `json:"status,omitempty"`
	CreatedAt    null.Time
	UpdatedAt    null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *ImageModelN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for ImageModel
func (inst *ImageModelN) SetModel(imageModelModel *ImageModelModel) {
	inst.imageModelModel = imageModelModel
}

// imageModelOriginal is an object which stores original ImageModel from database
type imageModelOriginal struct {
	Id           null.Int
	ModelId      null.String
	ModelName    null.String
	Vendor       null.String
	RealModel    null.String
	Meta         null.String
	PreviewImage null.String
	Description  null.String
	Status       null.Int
	CreatedAt    null.Time
	UpdatedAt    null.Time
}

// Staled identify whether the object has been modified
func (inst *ImageModelN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &imageModelOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.ModelId != inst.original.ModelId {
			return true
		}
		if inst.ModelName != inst.original.ModelName {
			return true
		}
		if inst.Vendor != inst.original.Vendor {
			return true
		}
		if inst.RealModel != inst.original.RealModel {
			return true
		}
		if inst.Meta != inst.original.Meta {
			return true
		}
		if inst.PreviewImage != inst.original.PreviewImage {
			return true
		}
		if inst.Description != inst.original.Description {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "model_id":
				if inst.ModelId != inst.original.ModelId {
					return true
				}
			case "model_name":
				if inst.ModelName != inst.original.ModelName {
					return true
				}
			case "vendor":
				if inst.Vendor != inst.original.Vendor {
					return true
				}
			case "real_model":
				if inst.RealModel != inst.original.RealModel {
					return true
				}
			case "meta":
				if inst.Meta != inst.original.Meta {
					return true
				}
			case "preview_image":
				if inst.PreviewImage != inst.original.PreviewImage {
					return true
				}
			case "description":
				if inst.Description != inst.original.Description {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *ImageModelN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &imageModelOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.ModelId != inst.original.ModelId {
			kv["model_id"] = inst.ModelId
		}
		if inst.ModelName != inst.original.ModelName {
			kv["model_name"] = inst.ModelName
		}
		if inst.Vendor != inst.original.Vendor {
			kv["vendor"] = inst.Vendor
		}
		if inst.RealModel != inst.original.RealModel {
			kv["real_model"] = inst.RealModel
		}
		if inst.Meta != inst.original.Meta {
			kv["meta"] = inst.Meta
		}
		if inst.PreviewImage != inst.original.PreviewImage {
			kv["preview_image"] = inst.PreviewImage
		}
		if inst.Description != inst.original.Description {
			kv["description"] = inst.Description
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "model_id":
				if inst.ModelId != inst.original.ModelId {
					kv["model_id"] = inst.ModelId
				}
			case "model_name":
				if inst.ModelName != inst.original.ModelName {
					kv["model_name"] = inst.ModelName
				}
			case "vendor":
				if inst.Vendor != inst.original.Vendor {
					kv["vendor"] = inst.Vendor
				}
			case "real_model":
				if inst.RealModel != inst.original.RealModel {
					kv["real_model"] = inst.RealModel
				}
			case "meta":
				if inst.Meta != inst.original.Meta {
					kv["meta"] = inst.Meta
				}
			case "preview_image":
				if inst.PreviewImage != inst.original.PreviewImage {
					kv["preview_image"] = inst.PreviewImage
				}
			case "description":
				if inst.Description != inst.original.Description {
					kv["description"] = inst.Description
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *ImageModelN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.imageModelModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.imageModelModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a image_model
func (inst *ImageModelN) Delete(ctx context.Context) error {
	if inst.imageModelModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.imageModelModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *ImageModelN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type imageModelScope struct {
	name  string
	apply func(builder query.Condition)
}

var imageModelGlobalScopes = make([]imageModelScope, 0)
var imageModelLocalScopes = make([]imageModelScope, 0)

// AddGlobalScopeForImageModel assign a global scope to a model
func AddGlobalScopeForImageModel(name string, apply func(builder query.Condition)) {
	imageModelGlobalScopes = append(imageModelGlobalScopes, imageModelScope{name: name, apply: apply})
}

// AddLocalScopeForImageModel assign a local scope to a model
func AddLocalScopeForImageModel(name string, apply func(builder query.Condition)) {
	imageModelLocalScopes = append(imageModelLocalScopes, imageModelScope{name: name, apply: apply})
}

func (m *ImageModelModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range imageModelGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range imageModelLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ImageModelModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ImageModelModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type ImageModel struct {
	Id           int64  `json:"id"`
	ModelId      string `json:"model_id,omitempty"`
	ModelName    string `json:"model_name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	RealModel    string `json:"real_model,omitempty"`
	Meta         string `json:"-"`
	PreviewImage string `json:"preview_image,omitempty"`
	Description  string `json:"description,omitempty"`
	Status       int64  `json:"status,omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (w ImageModel) ToImageModelN(allows ...string) ImageModelN {
	if len(allows) == 0 {
		return ImageModelN{

			Id:           null.IntFrom(int64(w.Id)),
			ModelId:      null.StringFrom(w.ModelId),
			ModelName:    null.StringFrom(w.ModelName),
			Vendor:       null.StringFrom(w.Vendor),
			RealModel:    null.StringFrom(w.RealModel),
			Meta:         null.StringFrom(w.Meta),
			PreviewImage: null.StringFrom(w.PreviewImage),
			Description:  null.StringFrom(w.Description),
			Status:       null.IntFrom(int64(w.Status)),
			CreatedAt:    null.TimeFrom(w.CreatedAt),
			UpdatedAt:    null.TimeFrom(w.UpdatedAt),
		}
	}

	res := ImageModelN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "model_id":
			res.ModelId = null.StringFrom(w.ModelId)
		case "model_name":
			res.ModelName = null.StringFrom(w.ModelName)
		case "vendor":
			res.Vendor = null.StringFrom(w.Vendor)
		case "real_model":
			res.RealModel = null.StringFrom(w.RealModel)
		case "meta":
			res.Meta = null.StringFrom(w.Meta)
		case "preview_image":
			res.PreviewImage = null.StringFrom(w.PreviewImage)
		case "description":
			res.Description = null.StringFrom(w.Description)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w ImageModel) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *ImageModelN) ToImageModel() ImageModel {
	return ImageModel{

		Id:           w.Id.Int64,
		ModelId:      w.ModelId.String,
		ModelName:    w.ModelName.String,
		Vendor:       w.Vendor.String,
		RealModel:    w.RealModel.String,
		Meta:         w.Meta.String,
		PreviewImage: w.PreviewImage.String,
		Description:  w.Description.String,
		Status:       w.Status.Int64,
		CreatedAt:    w.CreatedAt.Time,
		UpdatedAt:    w.UpdatedAt.Time,
	}
}

// ImageModelModel is a model which encapsulates the operations of the object
type ImageModelModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var imageModelTableName = "image_model"

// ImageModelTable return table name for ImageModel
func ImageModelTable() string {
	return imageModelTableName
}

const (
	FieldImageModelId           = "id"
	FieldImageModelModelId      = "model_id"
	FieldImageModelModelName    = "model_name"
	FieldImageModelVendor       = "vendor"
	FieldImageModelRealModel    = "real_model"
	FieldImageModelMeta         = "meta"
	FieldImageModelPreviewImage = "preview_image"
	FieldImageModelDescription  = "description"
	FieldImageModelStatus       = "status"
	FieldImageModelCreatedAt    = "created_at"
	FieldImageModelUpdatedAt    = "updated_at"
)

// ImageModelFields return all fields in ImageModel model
func ImageModelFields() []string {
	return []string{
		"id",
		"model_id",
		"model_name",
		"vendor",
		"real_model",
		"meta",
		"preview_image",
		"description",
		"status",
		"created_at",
		"updated_at",
	}
}

func SetImageModelTable(tableName string) {
	imageModelTableName = tableName
}

// NewImageModelModel create a ImageModelModel
func NewImageModelModel(db query.Database) *ImageModelModel {
	return &ImageModelModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           imageModelTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *ImageModelModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *ImageModelModel) clone() *ImageModelModel {
	return &ImageModelModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ImageModelModel) WithoutGlobalScopes(names ...string) *ImageModelModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ImageModelModel) WithLocalScopes(names ...string) *ImageModelModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *ImageModelModel) Condition(builder query.SQLBuilder) *ImageModelModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ImageModelModel) Find(ctx context.Context, id int64) (*ImageModelN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ImageModelModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ImageModelModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ImageModelModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]ImageModelN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ImageModelModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]ImageModelN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"model_id",
			"model_name",
			"vendor",
			"real_model",
			"meta",
			"preview_image",
			"description",
			"status",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "model_id":
			selectFields = append(selectFields, f)
		case "model_name":
			selectFields = append(selectFields, f)
		case "vendor":
			selectFields = append(selectFields, f)
		case "real_model":
			selectFields = append(selectFields, f)
		case "meta":
			selectFields = append(selectFields, f)
		case "preview_image":
			selectFields = append(selectFields, f)
		case "description":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*ImageModelN, []interface{}) {
		var imageModelVar ImageModelN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &imageModelVar.Id)
			case "model_id":
				scanFields = append(scanFields, &imageModelVar.ModelId)
			case "model_name":
				scanFields = append(scanFields, &imageModelVar.ModelName)
			case "vendor":
				scanFields = append(scanFields, &imageModelVar.Vendor)
			case "real_model":
				scanFields = append(scanFields, &imageModelVar.RealModel)
			case "meta":
				scanFields = append(scanFields, &imageModelVar.Meta)
			case "preview_image":
				scanFields = append(scanFields, &imageModelVar.PreviewImage)
			case "description":
				scanFields = append(scanFields, &imageModelVar.Description)
			case "status":
				scanFields = append(scanFields, &imageModelVar.Status)
			case "created_at":
				scanFields = append(scanFields, &imageModelVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &imageModelVar.UpdatedAt)
			}
		}

		return &imageModelVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	imageModels := make([]ImageModelN, 0)
	for rows.Next() {
		imageModelReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		imageModelReal.original = &imageModelOriginal{}
		_ = query.Copy(imageModelReal, imageModelReal.original)

		imageModelReal.SetModel(m)
		imageModels = append(imageModels, *imageModelReal)
	}

	return imageModels, nil
}

// First return first result for given query
func (m *ImageModelModel) First(ctx context.Context, builders ...query.SQLBuilder) (*ImageModelN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new image_model to database
func (m *ImageModelModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all image_models to database
func (m *ImageModelModel) SaveAll(ctx context.Context, imageModels []ImageModelN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, imageModel := range imageModels {
		id, err := m.Save(ctx, imageModel)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a image_model to database
func (m *ImageModelModel) Save(ctx context.Context, imageModel ImageModelN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, imageModel.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new image_model or update it when it has a id > 0
func (m *ImageModelModel) SaveOrUpdate(ctx context.Context, imageModel ImageModelN, onlyFields ...string) (id int64, updated bool, err error) {
	if imageModel.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, imageModel.Id.Int64, imageModel, onlyFields...)
		return imageModel.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, imageModel, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ImageModelModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ImageModelModel) Update(ctx context.Context, builder query.SQLBuilder, imageModel ImageModelN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, imageModel.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *ImageModelModel) UpdateById(ctx context.Context, id int64, imageModel ImageModelN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, imageModel.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *ImageModelModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *ImageModelModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
