package config

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ParsingOptions holds additional configuration for parsing different env values into native types
type ParsingOptions struct {
	// Optional delimiter used when parsing strings into slices (e.g. "," for parsing a string of vals separated by ,)
	// Default: ","
	SliceDelimiter string
}

type ConfigOptions struct {
	// Optional map of variables to use instead of system environment (Mainly for WASM builds)
	// Map is organised with the key being the environment variable and value being a pointer to
	// string variables injected with `ldflags` on compilation
	EnvOverrideMap *map[string]*string

	// Optional parsing options to use when parsing env values into native types
	ParsingOptions *ParsingOptions

	// Optional logger to use for logging
	// Default: slog.New(slog.NewTextHandler(os.Stdout, nil))
	Logger *slog.Logger
}

func LoadConfig(conf any, options *ConfigOptions) error {
	var logger *slog.Logger
	if options != nil {
		if options.Logger != nil {
			logger = options.Logger
		} else {
			logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
		}

		if options.EnvOverrideMap != nil {
			logger.Debug("loading config from override map")
		}
	}

	valOfConfig := reflect.ValueOf(conf)
	return parseEnvsIntoConfig(valOfConfig.Elem(), logger, options)
}

func parseValue(confProperty *reflect.Value, rawEnvVal string, parsingOptions *ParsingOptions) error {
	propKind := confProperty.Kind()
	switch propKind {
	case reflect.String:
		confProperty.SetString(rawEnvVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(rawEnvVal, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid int value: %w", err)
		}
		confProperty.SetInt(n)
	case reflect.Bool:
		b, err := strconv.ParseBool(rawEnvVal)
		if err != nil {
			return fmt.Errorf("invalid bool value: %w", err)
		}
		confProperty.SetBool(b)
	case reflect.Slice:
		delimiter := ","
		if parsingOptions.SliceDelimiter != "" {
			delimiter = parsingOptions.SliceDelimiter
		}
		splitValues := strings.Split(rawEnvVal, delimiter)
		confProperty.Set(reflect.ValueOf(splitValues))
	default:
		return fmt.Errorf("unsupported kind")
	}

	return nil
}

func getStructTagValue(structField *reflect.StructField, key string) string {
	tag := structField.Tag.Get(key)
	parts := strings.Split(tag, ",")
	return strings.TrimSpace(parts[0])
}

func getEnvValue(keyName string, options *ConfigOptions) (string, bool) {
	if options.EnvOverrideMap != nil {
		if val, ok := (*options.EnvOverrideMap)[keyName]; ok {
			return *val, true
		}
		return "", false
	}

	return os.LookupEnv(keyName)
}

// parseEnvsIntoConfig recursively digs into given struct type and fills in properties with env values
// keyed by the tag provided in the struct declaration
func parseEnvsIntoConfig(conf reflect.Value, logger *slog.Logger, options *ConfigOptions) error {
	confType := conf.Type()

	for i := 0; i < conf.NumField(); i++ {
		confProperty := conf.Field(i)
		confPropertyTags := confType.Field(i)
		confPropertyKind := confProperty.Kind()

		switch confPropertyKind {
		case reflect.Struct:
			logger.Debug("reading sub-config struct", "field", confPropertyTags.Name)
			if err := parseEnvsIntoConfig(confProperty, logger, options); err != nil {
				return err
			}

			continue
		}

		envVarName := getStructTagValue(&confPropertyTags, "env")
		defaultValue := getStructTagValue(&confPropertyTags, "default")
		raw, ok := getEnvValue(envVarName, options)
		if !ok || raw == "" {
			if defaultValue != "" {
				logger.Debug("using default value", "field", confPropertyTags.Name, "default", defaultValue)
				raw = defaultValue
			} else {
				logger.Warn("missing required config value", "field", confPropertyTags.Name, "env", envVarName)
				return fmt.Errorf("missing env var %q", envVarName)
			}
		}

		if err := parseValue(&confProperty, raw, options.ParsingOptions); err != nil {
			logger.Error("failed to parse env value", "field", confPropertyTags.Name, "env", envVarName, "value", raw, "error", err)
			return err
		}
	}

	return nil
}
