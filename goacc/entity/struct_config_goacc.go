// Code generated by github.com/kasaikou/goacc, DO NOT EDIT.
// defaultTag=-
package entity

type FileConfigBuilder interface {
	Build() *FileConfig
}

// fileConfigBuilderImpl is an instance for generating an instance of FileConfig.
type fileConfigBuilderImpl struct {
	__fc *FileConfig
}

// NewFileConfigBuilder creates an FileConfigBuilder instance.
func NewFileConfigBuilder(
	filename string,
	packageName string,
	imports []ImportConfig,
	structs []StructConfig,
) FileConfigBuilder {
	__fc := &FileConfig{}

	__fc.filename = filename
	__fc.packageName = packageName
	__fc.imports = imports
	__fc.structs = structs

	return &fileConfigBuilderImpl{__fc: __fc}
}

// Build purges FileConfig instance from fileConfigBuilderImpl.
//
// If calls other method in fileConfigBuilderImpl after Purge called, it will be panic.
func (__fcb *fileConfigBuilderImpl) Build() *FileConfig {
	if __fcb == nil {
		panic("fileConfigBuilderImpl is nil")
	} else if __fcb.__fc != nil {
		__fc := __fcb.__fc
		__fcb.__fc = nil

		return __fc
	}

	panic("FileConfig has been already purged")
}

// Absolute path of Go file.
func (__fc *FileConfig) Filename() string {
	if __fc != nil {
		return __fc.filename
	}

	panic("FileConfig is nil")
}

// Package name.
func (__fc *FileConfig) PackageName() string {
	if __fc != nil {
		return __fc.packageName
	}

	panic("FileConfig is nil")
}

// Configuration of file defining Structs.
func (__fc *FileConfig) Structs() []StructConfig {
	if __fc != nil {
		return __fc.structs
	}

	panic("FileConfig is nil")
}

type ImportConfigBuilder interface {
	Build() *ImportConfig
}

// importConfigBuilderImpl is an instance for generating an instance of ImportConfig.
type importConfigBuilderImpl struct {
	__ic *ImportConfig
}

// NewImportConfigBuilder creates an ImportConfigBuilder instance.
func NewImportConfigBuilder(
	name string,
	path string,
) ImportConfigBuilder {
	__ic := &ImportConfig{}

	__ic.name = name
	__ic.path = path

	return &importConfigBuilderImpl{__ic: __ic}
}

// Build purges ImportConfig instance from importConfigBuilderImpl.
//
// If calls other method in importConfigBuilderImpl after Purge called, it will be panic.
func (__icb *importConfigBuilderImpl) Build() *ImportConfig {
	if __icb == nil {
		panic("importConfigBuilderImpl is nil")
	} else if __icb.__ic != nil {
		__ic := __icb.__ic
		__icb.__ic = nil

		return __ic
	}

	panic("ImportConfig has been already purged")
}

type StructConfigBuilder interface {
	SetDocText(docText string) StructConfigBuilder
	SetDefineFilename(defineFilename string) StructConfigBuilder
	Build() *StructConfig
}

// structConfigBuilderImpl is an instance for generating an instance of StructConfig.
type structConfigBuilderImpl struct {
	__sc *StructConfig
}

// NewStructConfigBuilder creates an StructConfigBuilder instance.
func NewStructConfigBuilder(
	name string,
	structSupports StructSupports,
	mutexFieldName string,
	enableMarshalJSON bool,
	fields []FieldConfig,
) StructConfigBuilder {
	__sc := &StructConfig{}

	__sc.name = name
	__sc.structSupports = structSupports
	__sc.mutexFieldName = mutexFieldName
	__sc.enableMarshalJSON = enableMarshalJSON
	__sc.fields = fields

	return &structConfigBuilderImpl{__sc: __sc}
}

func (__scb *structConfigBuilderImpl) SetDocText(docText string) StructConfigBuilder {
	if __scb == nil {
		panic("structConfigBuilderImpl is nil")
	} else if __scb.__sc != nil {
		__scb.__sc.docText = docText
		return __scb
	}

	panic("StructConfig has been already purged")
}

func (__scb *structConfigBuilderImpl) SetDefineFilename(defineFilename string) StructConfigBuilder {
	if __scb == nil {
		panic("structConfigBuilderImpl is nil")
	} else if __scb.__sc != nil {
		__scb.__sc.defineFilename = defineFilename
		return __scb
	}

	panic("StructConfig has been already purged")
}

// Build purges StructConfig instance from structConfigBuilderImpl.
//
// If calls other method in structConfigBuilderImpl after Purge called, it will be panic.
func (__scb *structConfigBuilderImpl) Build() *StructConfig {
	if __scb == nil {
		panic("structConfigBuilderImpl is nil")
	} else if __scb.__sc != nil {
		__sc := __scb.__sc
		__scb.__sc = nil

		return __sc
	}

	panic("StructConfig has been already purged")
}

func (__sc *StructConfig) Name() string {
	if __sc != nil {
		return __sc.name
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) DocText() string {
	if __sc != nil {
		return __sc.docText
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) DefineFilename() string {
	if __sc != nil {
		return __sc.defineFilename
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) MutexFieldName() string {
	if __sc != nil {
		return __sc.mutexFieldName
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) EnableMarshalJson() bool {
	if __sc != nil {
		return __sc.enableMarshalJSON
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) Fields() []FieldConfig {
	if __sc != nil {
		return __sc.fields
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) StructSupportsPtr() *StructSupports {
	if __sc != nil {
		return &__sc.structSupports
	}

	panic("StructConfig is nil")
}

func (__sc *StructConfig) SetDocText(docText string) {
	if __sc != nil {
		__sc.docText = docText
	}
}

func (__sc *StructConfig) SetDefineFilename(defineFilename string) {
	if __sc != nil {
		__sc.defineFilename = defineFilename
	}
}

type StructSupportsBuilder interface {
	SetHasPreNewHook(hasPreNewHook bool) StructSupportsBuilder
	SetHasPostNewHook(hasPostNewHook bool) StructSupportsBuilder
	SetHasPostNewHookError(hasPostNewHookError bool) StructSupportsBuilder
	Build() *StructSupports
}

// structSupportsBuilderImpl is an instance for generating an instance of StructSupports.
type structSupportsBuilderImpl struct {
	__ss *StructSupports
}

// NewStructSupportsBuilder creates an StructSupportsBuilder instance.
func NewStructSupportsBuilder() StructSupportsBuilder {
	__ss := &StructSupports{}

	return &structSupportsBuilderImpl{__ss: __ss}
}

func (__ssb *structSupportsBuilderImpl) SetHasPreNewHook(hasPreNewHook bool) StructSupportsBuilder {
	if __ssb == nil {
		panic("structSupportsBuilderImpl is nil")
	} else if __ssb.__ss != nil {
		__ssb.__ss.hasPreNewHook = hasPreNewHook
		return __ssb
	}

	panic("StructSupports has been already purged")
}

func (__ssb *structSupportsBuilderImpl) SetHasPostNewHook(hasPostNewHook bool) StructSupportsBuilder {
	if __ssb == nil {
		panic("structSupportsBuilderImpl is nil")
	} else if __ssb.__ss != nil {
		__ssb.__ss.hasPostNewHook = hasPostNewHook
		return __ssb
	}

	panic("StructSupports has been already purged")
}

func (__ssb *structSupportsBuilderImpl) SetHasPostNewHookError(hasPostNewHookError bool) StructSupportsBuilder {
	if __ssb == nil {
		panic("structSupportsBuilderImpl is nil")
	} else if __ssb.__ss != nil {
		__ssb.__ss.hasPostNewHookError = hasPostNewHookError
		return __ssb
	}

	panic("StructSupports has been already purged")
}

// Build purges StructSupports instance from structSupportsBuilderImpl.
//
// If calls other method in structSupportsBuilderImpl after Purge called, it will be panic.
func (__ssb *structSupportsBuilderImpl) Build() *StructSupports {
	if __ssb == nil {
		panic("structSupportsBuilderImpl is nil")
	} else if __ssb.__ss != nil {
		__ss := __ssb.__ss
		__ssb.__ss = nil

		return __ss
	}

	panic("StructSupports has been already purged")
}

func (__ss *StructSupports) HasPreNewHook() bool {
	if __ss != nil {
		return __ss.hasPreNewHook
	}

	panic("StructSupports is nil")
}

func (__ss *StructSupports) HasPostNewHook() bool {
	if __ss != nil {
		return __ss.hasPostNewHook
	}

	panic("StructSupports is nil")
}

func (__ss *StructSupports) HasPostNewHookError() bool {
	if __ss != nil {
		return __ss.hasPostNewHookError
	}

	panic("StructSupports is nil")
}

type FieldConfigBuilder interface {
	Build() *FieldConfig
}

// fieldConfigBuilderImpl is an instance for generating an instance of FieldConfig.
type fieldConfigBuilderImpl struct {
	__fc *FieldConfig
}

// NewFieldConfigBuilder creates an FieldConfigBuilder instance.
func NewFieldConfigBuilder(
	name string,
	typeName string,
	jsonTag string,
	features *FieldConfigFeatures,
) FieldConfigBuilder {
	__fc := &FieldConfig{}

	__fc.name = name
	__fc.typeName = typeName
	__fc.jsonTag = jsonTag
	__fc.features = features

	return &fieldConfigBuilderImpl{__fc: __fc}
}

// Build purges FieldConfig instance from fieldConfigBuilderImpl.
//
// If calls other method in fieldConfigBuilderImpl after Purge called, it will be panic.
func (__fcb *fieldConfigBuilderImpl) Build() *FieldConfig {
	if __fcb == nil {
		panic("fieldConfigBuilderImpl is nil")
	} else if __fcb.__fc != nil {
		__fc := __fcb.__fc
		__fcb.__fc = nil

		return __fc
	}

	panic("FieldConfig has been already purged")
}

func (__fc *FieldConfig) Name() string {
	if __fc != nil {
		return __fc.name
	}

	panic("FieldConfig is nil")
}

func (__fc *FieldConfig) DocText() string {
	if __fc != nil {
		return __fc.docText
	}

	panic("FieldConfig is nil")
}

func (__fc *FieldConfig) TypeName() string {
	if __fc != nil {
		return __fc.typeName
	}

	panic("FieldConfig is nil")
}

// If empty, it means no json tag.
func (__fc *FieldConfig) JsonTag() string {
	if __fc != nil {
		return __fc.jsonTag
	}

	panic("FieldConfig is nil")
}

func (__fc *FieldConfig) Features() *FieldConfigFeatures {
	if __fc != nil {
		return __fc.features
	}

	panic("FieldConfig is nil")
}

func (__fc *FieldConfig) SetDocText(docText string) {
	if __fc != nil {
		__fc.docText = docText
	}
}

func (__fc *FieldConfig) SetTypeName(typeName string) {
	if __fc != nil {
		__fc.typeName = typeName
	}
}

func (__fc *FieldConfig) SetFeatures(features *FieldConfigFeatures) {
	if __fc != nil {
		__fc.features = features
	}
}

type FieldConfigFeaturesBuilder interface {
	Build() *FieldConfigFeatures
}

// fieldConfigFeaturesBuilderImpl is an instance for generating an instance of FieldConfigFeatures.
type fieldConfigFeaturesBuilderImpl struct {
	__fcf *FieldConfigFeatures
}

// NewFieldConfigFeaturesBuilder creates an FieldConfigFeaturesBuilder instance.
func NewFieldConfigFeaturesBuilder(
	usesMutex bool,
	hasRequired bool,
	hasOptional bool,
	hasPtrGetter bool,
	hasGetter bool,
	hasSetter bool,
) FieldConfigFeaturesBuilder {
	__fcf := &FieldConfigFeatures{}

	__fcf.usesMutex = usesMutex
	__fcf.hasRequired = hasRequired
	__fcf.hasOptional = hasOptional
	__fcf.hasPtrGetter = hasPtrGetter
	__fcf.hasGetter = hasGetter
	__fcf.hasSetter = hasSetter

	return &fieldConfigFeaturesBuilderImpl{__fcf: __fcf}
}

// Build purges FieldConfigFeatures instance from fieldConfigFeaturesBuilderImpl.
//
// If calls other method in fieldConfigFeaturesBuilderImpl after Purge called, it will be panic.
func (__fcfb *fieldConfigFeaturesBuilderImpl) Build() *FieldConfigFeatures {
	if __fcfb == nil {
		panic("fieldConfigFeaturesBuilderImpl is nil")
	} else if __fcfb.__fcf != nil {
		__fcf := __fcfb.__fcf
		__fcfb.__fcf = nil

		return __fcf
	}

	panic("FieldConfigFeatures has been already purged")
}

func (__fcf *FieldConfigFeatures) UsesMutex() bool {
	if __fcf != nil {
		return __fcf.usesMutex
	}

	panic("FieldConfigFeatures is nil")
}

func (__fcf *FieldConfigFeatures) HasRequired() bool {
	if __fcf != nil {
		return __fcf.hasRequired
	}

	panic("FieldConfigFeatures is nil")
}

func (__fcf *FieldConfigFeatures) HasOptional() bool {
	if __fcf != nil {
		return __fcf.hasOptional
	}

	panic("FieldConfigFeatures is nil")
}

func (__fcf *FieldConfigFeatures) HasPtrGetter() bool {
	if __fcf != nil {
		return __fcf.hasPtrGetter
	}

	panic("FieldConfigFeatures is nil")
}

func (__fcf *FieldConfigFeatures) HasGetter() bool {
	if __fcf != nil {
		return __fcf.hasGetter
	}

	panic("FieldConfigFeatures is nil")
}

func (__fcf *FieldConfigFeatures) HasSetter() bool {
	if __fcf != nil {
		return __fcf.hasSetter
	}

	panic("FieldConfigFeatures is nil")
}
