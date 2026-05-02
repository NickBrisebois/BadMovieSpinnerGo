package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func LoadConfig(conf any) error {
	valOfConfig := reflect.ValueOf(conf)
	return parseEnvsIntoConfig(valOfConfig.Elem())
}

func parseValue(confProperty *reflect.Value, confKind reflect.Kind, rawEnvVal string, envVarName string) error {
	switch confProperty.Kind() {
	case reflect.String:
		confProperty.SetString(rawEnvVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(rawEnvVal, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid int value for %q: %w", envVarName, err)
		}
		confProperty.SetInt(n)
	case reflect.Bool:
		b, err := strconv.ParseBool(rawEnvVal)
		if err != nil {
			return fmt.Errorf("invalid bool value for %q: %w", envVarName, err)
		}
		confProperty.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %v for %q", confKind, envVarName)
	}

	return nil
}

func getStructTagValue(structField *reflect.StructField, key string) string {
	tag := structField.Tag.Get(key)
	parts := strings.Split(tag, ",")
	return strings.TrimSpace(parts[0])
}

// parseEnvsIntoConfig recursively digs into given struct type and fills in properties with env values
// keyed by the tag provided in the struct declaration
func parseEnvsIntoConfig(conf reflect.Value) error {
	confType := conf.Type()

	for i := 0; i < conf.NumField(); i++ {
		confProperty := conf.Field(i)
		confPropertyTags := confType.Field(i)

		if confProperty.Kind() == reflect.Struct {
			if err := parseEnvsIntoConfig(confProperty); err != nil {
				return err
			}

			continue
		}

		envVarName := getStructTagValue(&confPropertyTags, "env")
		defaultValue := getStructTagValue(&confPropertyTags, "default")

		raw, ok := getEnvValue(envVarName)
		if !ok || raw == "" {
			if defaultValue != "" {
				raw = defaultValue
			} else {
				return fmt.Errorf("missing env var %q", envVarName)
			}
		}

		if err := parseValue(&confProperty, conf.Kind(), raw, envVarName); err != nil {
			return err
		}
	}

	return nil
}
