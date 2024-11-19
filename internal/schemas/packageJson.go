// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    packageJSONSchema, err := UnmarshalPackageJSONSchema(bytes)
//    bytes, err = packageJSONSchema.Marshal()

package schemas_package_json

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalPackageJSONSchema(data []byte) (PackageJSONSchema, error) {
	var r PackageJSONSchema
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PackageJSONSchema) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PackageJSONSchema struct {
	Name             string             `json:"name"`
	Version          string             `json:"version"`
	Private          *bool              `json:"private,omitempty"`
	Files            []string           `json:"files"`
	Scripts          map[string]string  `json:"scripts,omitempty"`
	Wireit           *Wireit            `json:"wireit,omitempty"`
	Dependencies     map[string]string  `json:"dependencies"`
	DevDependencies  map[string]string  `json:"devDependencies"`
	PeerDependencies map[string]UlldAPI `json:"peerDependencies"`
	PackageManager   *PackageManager    `json:"packageManager,omitempty"`
	Engines          *Engines           `json:"engines,omitempty"`
	Prisma           *Prisma            `json:"prisma,omitempty"`
	PublishConfig    *PublishConfig     `json:"publishConfig,omitempty"`
	Exports          *Exports           `json:"exports"`
	Type             *Type              `json:"type,omitempty"`
	Types            *string            `json:"types,omitempty"`
	Main             *string            `json:"main,omitempty"`
	UlldPluginConfig *UlldPluginConfig  `json:"ulld-pluginConfig,omitempty"`
	Module           *string            `json:"module,omitempty"`
	UlldPluginID     *string            `json:"ulld-plugin-id,omitempty"`
	License          *string            `json:"license,omitempty"`
	UlldAPI          *UlldAPI           `json:"@ulld/api,omitempty"`
	Bin              *Bin               `json:"bin,omitempty"`
	Oclif            *Oclif             `json:"oclif,omitempty"`
	UlldDeveloper    *UlldDeveloper     `json:"ulld-developer,omitempty"`
	TypesVersions    *TypesVersions     `json:"typesVersions,omitempty"`
	Prettier         *string            `json:"prettier,omitempty"`
	Xo               *Xo                `json:"xo,omitempty"`
	Ava              *Ava               `json:"ava,omitempty"`
	Repository       *string            `json:"repository,omitempty"`
}

type Ava struct {
	Extensions    Extensions `json:"extensions"`
	NodeArguments []string   `json:"nodeArguments"`
}

type Extensions struct {
	Ts  Type `json:"ts"`
	Tsx Type `json:"tsx"`
}

type Bin struct {
	Mr      *string `json:"mr,omitempty"`
	UlldDev *string `json:"ulld-dev,omitempty"`
	Ulld    *string `json:"ulld,omitempty"`
}

type Engines struct {
	Node string `json:"node"`
}

type ExportClass struct {
	Types   *string `json:"types,omitempty"`
	Import  *string `json:"import,omitempty"`
	Require *string `json:"require,omitempty"`
	Node    *string `json:"node,omitempty"`
}

type Oclif struct {
	Bin                    string                 `json:"bin"`
	Dirname                string                 `json:"dirname"`
	Commands               *CommandsUnion         `json:"commands"`
	Plugins                []string               `json:"plugins"`
	TopicSeparator         string                 `json:"topicSeparator"`
	BinAliases             []string               `json:"binAliases,omitempty"`
	Scope                  *string                `json:"scope,omitempty"`
	Theme                  *string                `json:"theme,omitempty"`
	State                  *string                `json:"state,omitempty"`
	FlexibleTaxonomy       *bool                  `json:"flexibleTaxonomy,omitempty"`
	Hooks                  *Hooks                 `json:"hooks,omitempty"`
	WarnIfUpdateAvailable  *WarnIfUpdateAvailable `json:"warn-if-update-available,omitempty"`
	AdditionalVersionFlags []string               `json:"additionalVersionFlags,omitempty"`
	AdditionalHelpFlags    []string               `json:"additionalHelpFlags,omitempty"`
}

type CommandsClass struct {
	Strategy     string   `json:"strategy"`
	Target       string   `json:"target"`
	GlobPatterns []string `json:"globPatterns"`
}

type Hooks struct {
	CommandIncomplete string `json:"command_incomplete"`
}

type WarnIfUpdateAvailable struct {
	TimeoutInDays int64  `json:"timeoutInDays"`
	Message       string `json:"message"`
}

type Prisma struct {
	Schema string `json:"schema"`
}

type PublishConfig struct {
	Access Access `json:"access"`
}

type TypesVersions struct {
	Empty Class `json:"*"`
}

type Class struct {
	IDE              []string `json:"ide"`
	LanguageSelect   []string `json:"languageSelect"`
	Modal            []string `json:"modal"`
	TextAreaEditor   []string `json:"textAreaEditor"`
	EditorTypes      []string `json:"editorTypes"`
	UtilityFunctions []string `json:"utilityFunctions"`
}

type UlldDeveloper struct {
	RouterPaths RouterPaths `json:"routerPaths"`
}

type RouterPaths struct {
	Equations string `json:"equations"`
}

type UlldPluginConfig struct {
	PluginName      string           `json:"pluginName"`
	Label           string           `json:"label"`
	PluginID        string           `json:"pluginId"`
	Slot            *string          `json:"slot,omitempty"`
	Components      []Component      `json:"components"`
	Parsers         Parsers          `json:"parsers"`
	Pages           []Page           `json:"pages"`
	Events          Events           `json:"events"`
	NavigationLinks []NavigationLink `json:"navigationLinks"`
	CommandPalette  []interface{}    `json:"commandPalette"`
	Tailwind        Styles           `json:"tailwind"`
	Styles          Styles           `json:"styles"`
}

type Component struct {
	ComponentName     string        `json:"componentName"`
	Tags              []interface{} `json:"tags"`
	Slot              *string       `json:"slot,omitempty"`
	Export            string        `json:"export"`
	ComponentID       string        `json:"componentId"`
	Embeddable        []Embeddable  `json:"embeddable,omitempty"`
	ExportedPropsName *string       `json:"exportedPropsName,omitempty"`
}

type Embeddable struct {
	RegexToInclude string `json:"regexToInclude"`
	Label          string `json:"label"`
}

type Events struct {
	OnBackup  *string `json:"onBackup,omitempty"`
	OnRestore *string `json:"onRestore,omitempty"`
}

type NavigationLink struct {
	Label    string `json:"label"`
	Href     string `json:"href"`
	Category string `json:"category"`
}

type Page struct {
	Slot             string  `json:"slot"`
	Export           string  `json:"export"`
	ExportsPageProps bool    `json:"exportsPageProps"`
	TargetURL        *string `json:"targetUrl,omitempty"`
}

type Parsers struct {
	Mdx *Mdx `json:"mdx,omitempty"`
}

type Mdx struct {
	Export     string `json:"export"`
	Conditions Styles `json:"conditions"`
}

type Styles struct {
}

type Wireit struct {
	PrepareCommitWebOnly           *CreateDevelopmentDatabase `json:"prepareCommit:webOnly,omitempty"`
	PrepareWeb                     *CreateDevelopmentDatabase `json:"prepareWeb,omitempty"`
	PrepareCommitVersionEverything *CreateDevelopmentDatabase `json:"prepareCommit:versionEverything,omitempty"`
	GenerateData                   *CreateDevelopmentDatabase `json:"generateData,omitempty"`
	GenerateMdxData                *CleanBuild                `json:"generateMdxData,omitempty"`
	GenerateTypesData              *CleanBuild                `json:"generateTypesData,omitempty"`
	SyncTemplatePackage            *SyncTemplatePackage       `json:"syncTemplatePackage,omitempty"`
	Build                          *Build                     `json:"build,omitempty"`
	Rebuild                        *CreateDevelopmentDatabase `json:"rebuild,omitempty"`
	Compile                        *Bundle                    `json:"compile,omitempty"`
	DBAutoMigrate                  *CleanBuild                `json:"db:autoMigrate,omitempty"`
	DBMigrate                      *CleanBuild                `json:"db:migrate,omitempty"`
	DBGenerate                     *CleanBuild                `json:"db:generate,omitempty"`
	RunOnBuild                     *BuildTypes                `json:"runOnBuild,omitempty"`
	CleanPreBuild                  *BuildTypes                `json:"clean:preBuild,omitempty"`
	CleanPostBuild                 *BuildTypes                `json:"clean:postBuild,omitempty"`
	ToWorkspaceDeps                *Bundle                    `json:"toWorkspaceDeps,omitempty"`
	ToTemplateDeps                 *Bundle                    `json:"toTemplateDeps,omitempty"`
	Lint                           *Lint                      `json:"lint,omitempty"`
	Typecheck                      *Lint                      `json:"typecheck,omitempty"`
	GenTypes                       *Bundle                    `json:"genTypes,omitempty"`
	ClearBuild                     *CleanBuild                `json:"clearBuild,omitempty"`
	GenBuild                       *BuildTypes                `json:"genBuild,omitempty"`
	GenerateInternalLocationsData  *CleanBuild                `json:"generateInternalLocationsData,omitempty"`
	GenerateGlobalActionsData      *CleanBuild                `json:"generateGlobalActionsData,omitempty"`
	GenerateValidIconNameData      *CleanBuild                `json:"generateValidIconNameData,omitempty"`
	GatherTypes                    *Build                     `json:"gatherTypes,omitempty"`
	Transpile                      *Bundle                    `json:"transpile,omitempty"`
	CleanBuild                     *CleanBuild                `json:"cleanBuild,omitempty"`
	SetDatabaseTunnelExports       *CleanBuild                `json:"setDatabaseTunnelExports,omitempty"`
	EnsureProperDatabaseImports    *CleanBuild                `json:"ensureProperDatabaseImports,omitempty"`
	Bundle                         *Bundle                    `json:"bundle,omitempty"`
	CreateDevelopmentDatabase      *CreateDevelopmentDatabase `json:"createDevelopmentDatabase,omitempty"`
	MigrateDevelopmentDatabase     *CleanBuild                `json:"migrateDevelopmentDatabase,omitempty"`
	GenerateDefaultConfig          *BuildTypes                `json:"generateDefaultConfig,omitempty"`
	GenerateSlotMap                *BuildTypes                `json:"generateSlotMap,omitempty"`
	GenerateJSONSchemas            *BuildTypes                `json:"generateJsonSchemas,omitempty"`
	UpdateJSONSchemas              *BuildTypes                `json:"updateJsonSchemas,omitempty"`
	GenerateConfigTunnels          *BuildTypes                `json:"generateConfigTunnels,omitempty"`
	UpdateShikiThemes              *BuildTypes                `json:"updateShikiThemes,omitempty"`
	UpdateSlotMap                  *BuildTypes                `json:"updateSlotMap,omitempty"`
	ClearTypesFromTypesPackage     *BuildTypes                `json:"clearTypesFromTypesPackage,omitempty"`
	BuildTypes                     *BuildTypes                `json:"buildTypes,omitempty"`
	Wireit                         *Bundle                    `json:"wireit,omitempty"`
	SeedTestBuild                  *BuildTypes                `json:"seedTestBuild,omitempty"`
}

type Build struct {
	Dependencies []string `json:"dependencies,omitempty"`
	Clean        *Clean   `json:"clean"`
	Command      *string  `json:"command,omitempty"`
	Output       []string `json:"output,omitempty"`
	Files        []string `json:"files,omitempty"`
}

type BuildTypes struct {
	Command string `json:"command"`
	Clean   bool   `json:"clean"`
}

type Bundle struct {
	Command      string   `json:"command"`
	Clean        *bool    `json:"clean,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	Output       []string `json:"output,omitempty"`
	Files        []string `json:"files,omitempty"`
}

type CleanBuild struct {
	Command string `json:"command"`
}

type CreateDevelopmentDatabase struct {
	Dependencies []string `json:"dependencies"`
}

type Lint struct {
	Command string `json:"command"`
	Clean   *bool  `json:"clean,omitempty"`
	Service *bool  `json:"service,omitempty"`
}

type SyncTemplatePackage struct {
	Command string `json:"command"`
	Clean   bool   `json:"clean"`
	DevOnly bool   `json:"devOnly"`
}

type Xo struct {
	Extends  string `json:"extends"`
	Prettier bool   `json:"prettier"`
	Rules    Rules  `json:"rules"`
}

type Rules struct {
	ReactPropTypes string `json:"react/prop-types"`
}

type Type string

const (
	Module Type = "module"
)

type PackageManager string

const (
	Pnpm904 PackageManager = "pnpm@9.0.4"
	Pnpm940 PackageManager = "pnpm@9.4.0"
)

type UlldAPI string

const (
	Empty     UlldAPI = "*"
	The001    UlldAPI = "0.0.1"
	The0126   UlldAPI = "0.12.6"
	The041    UlldAPI = "0.4.1"
	The100    UlldAPI = "1.0.0"
	The10452  UlldAPI = "10.45.2"
	The107    UlldAPI = "1.0.7"
	The1140   UlldAPI = "1.14.0"
	The1154   UlldAPI = "11.5.4"
	The117    UlldAPI = "1.1.7"
	The1423   UlldAPI = "14.2.3"
	The1778   UlldAPI = "1.77.8"
	The1820   UlldAPI = "18.2.0"
	The197    UlldAPI = "1.9.7"
	The205    UlldAPI = "2.0.5"
	The322    UlldAPI = "3.2.2"
	The341    UlldAPI = "3.4.1"
	The5170   UlldAPI = "5.17.0"
	The813    UlldAPI = "8.1.3"
	Workspace UlldAPI = "workspace:*"
)

type Access string

const (
	Public Access = "public"
)

type Exports struct {
	String   *string
	UnionMap map[string]*ExportValue
}

func (x *Exports) UnmarshalJSON(data []byte) error {
	x.UnionMap = nil
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, false, nil, true, &x.UnionMap, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Exports) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, false, nil, x.UnionMap != nil, x.UnionMap, false, nil, false)
}

type ExportValue struct {
	ExportClass *ExportClass
	String      *string
}

func (x *ExportValue) UnmarshalJSON(data []byte) error {
	x.ExportClass = nil
	var c ExportClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ExportClass = &c
	}
	return nil
}

func (x *ExportValue) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ExportClass != nil, x.ExportClass, false, nil, false, nil, false)
}

type CommandsUnion struct {
	CommandsClass *CommandsClass
	String        *string
}

func (x *CommandsUnion) UnmarshalJSON(data []byte) error {
	x.CommandsClass = nil
	var c CommandsClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.CommandsClass = &c
	}
	return nil
}

func (x *CommandsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.CommandsClass != nil, x.CommandsClass, false, nil, false, nil, false)
}

type Clean struct {
	Bool   *bool
	String *string
}

func (x *Clean) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, nil, nil, &x.Bool, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Clean) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, x.Bool, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
