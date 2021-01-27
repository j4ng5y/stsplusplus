package identity

type Output string

const (
	Output_JSON       Output = "json"
	Output_YAML       Output = "yaml"
	Output_YAMLStream Output = "yaml-stream"
	Output_Text       Output = "text"
	Output_Table      Output = "table"
)
