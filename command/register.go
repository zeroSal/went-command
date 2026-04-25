package command

import (
	"reflect"

	"github.com/spf13/cobra"
)

func Register(
	commands []Interface,
	root *cobra.Command,
	run func(command Interface),
) *cobra.Command {
	for _, command := range commands {
		cobraCmd := &cobra.Command{
			Use:   command.GetHeader().Use,
			Short: command.GetHeader().Short,
			Long:  command.GetHeader().Long,
		}

		instanceValue := reflect.ValueOf(command).Elem()

		params := newParams()

		if baseField, ok := findField(instanceValue, "Base"); ok {
			if baseField.Kind() == reflect.Struct {
				paramsField := baseField.FieldByName("Params")
				if paramsField.IsValid() && paramsField.CanSet() {
					paramsField.Set(reflect.ValueOf(params))
				}
			}
		}

		if header := command.GetHeader(); header.Flags != nil {
			for _, boolFlag := range header.Flags.Bool {
				var ptr bool
				cobraCmd.Flags().BoolVar(&ptr, boolFlag.Name, boolFlag.Default, boolFlag.Usage)
				params.SetFlag(boolFlag.Name, &ptr)
			}
			for _, intFlag := range header.Flags.Int {
				var ptr int
				cobraCmd.Flags().IntVar(&ptr, intFlag.Name, intFlag.Default, intFlag.Usage)
				params.SetFlag(intFlag.Name, &ptr)
			}
			for _, int64Flag := range header.Flags.Int64 {
				var ptr int64
				cobraCmd.Flags().Int64Var(&ptr, int64Flag.Name, int64Flag.Default, int64Flag.Usage)
				params.SetFlag(int64Flag.Name, &ptr)
			}
			for _, stringFlag := range header.Flags.String {
				var ptr string
				cobraCmd.Flags().StringVar(&ptr, stringFlag.Name, stringFlag.Default, stringFlag.Usage)
				params.SetFlag(stringFlag.Name, &ptr)
			}
		}

		arguments := command.GetHeader().Arguments

		cobraCmd.Args = func(cmd *cobra.Command, args []string) error {
			for i, arg := range arguments {
				if i >= len(args) {
					continue
				}
				fieldName := toFieldName(arg.Name)
				if field, ok := findField(instanceValue, fieldName); ok {
					if field.Kind() == reflect.Pointer && field.Type().Elem() == reflect.TypeFor[string]() {
						ptr := args[i]
						field.Set(reflect.ValueOf(&ptr))
					}
				}
			}
			return nil
		}

		cobraCmd.Run = func(cmd *cobra.Command, args []string) {
			params.SetArgs(args)
			run(command)
		}

		root.AddCommand(cobraCmd)
	}

	return root
}

func toFieldName(name string) string {
	result := make([]byte, 0, len(name))
	for i, c := range name {
		if c == '-' {
			continue
		}
		if i == 0 || name[i-1] == '-' {
			result = append(result, byte(c-'a'+'A'))
		} else {
			result = append(result, byte(c))
		}
	}
	return string(result)
}

func findField(v reflect.Value, name string) (reflect.Value, bool) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Pointer && field.Elem().Kind() == reflect.Struct {
			if found, ok := findField(field.Elem(), name); ok {
				return found, true
			}
		}
		if v.Type().Field(i).Name == name {
			return field, true
		}
	}
	return reflect.Value{}, false
}
