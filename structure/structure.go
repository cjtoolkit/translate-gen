package structure

type FileBase struct {
	Imports     []Import     `toml:"import"`
	Base        *Base        `toml:"base"`
	Translation *Translation `toml:"translation"`
}

type Import struct {
	Path  string `toml:"path"`
	Alias string `toml:"alias"`
}

type Value struct {
	Const string `toml:"const"`
	Value string `toml:"value"`
}

type Base struct {
	Language  string  `toml:"language"`
	CreateMap bool    `toml:"createMap"`
	Values    []Value `toml:"value"`
}

type Translation struct {
	Language     string  `toml:"language"`
	FunctionName string  `toml:"functionName"`
	Values       []Value `toml:"value"`
}
