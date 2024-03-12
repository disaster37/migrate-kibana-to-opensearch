package migrate

type Object struct {
	Attributes           map[string]any   `json:"attributes,omitempty"`
	CoreMigrationVersion string           `json:"coreMigrationVersion,omitempty"`
	CreatedAt            string           `json:"created_at,omitempty"`
	Id                   string           `json:"id,omitempty"`
	MigrationVersion     MigrationVersion `json:"migrationVersion,omitempty"`
	Type                 string           `json:"type,omitempty"`
	UpdatedAt            string           `json:"updatedAt,omitempty"`
	Version              string           `json:"version,omitempty"`
	References           []any            `json:"references,omitempty"`
	OriginId             string           `json:"originId,omitempty"`
}

type MigrationVersion struct {
	IndexPattern  string `json:"index-pattern,omitempty"`
	Search        string `json:"search,omitempty"`
	Dashboard     string `json:"dashboard,omitempty"`
	Visualization string `json:"visualization,omitempty"`
	Map           string `json:"map,omitempty"`
	Lens          string `json:"lens,omitempty"`
}
