package helpers

type PropertySelector[TEntity any] struct {
	FieldName string
}

func NewPropertySelector[TEntity any](fieldName string) PropertySelector[TEntity] {
	// entityType := reflect.TypeOf((*TEntity)(nil)).Elem()
	// for i := 0; i < entityType.NumField(); i++ {
	// 	field := entityType.Field(i)
	// 	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
	// 		tagParts := strings.Split(jsonTag, ",")
	// 		if tagParts[0] == fieldName {
	// 			return PropertySelector[TEntity]{FieldName: fieldName}
	// 		}
	// 	}
	// }

	// if gormTag := field.Tag.Get("gorm"); gormTag != "" {
	//     tagParts := strings.Split(gormTag, ";")
	//     for _, part := range tagParts {
	//         if strings.HasPrefix(part, "column:") {
	//             columnName := strings.TrimPrefix(part, "column:")
	//             if columnName == fieldName {
	//                 return PropertySelector[TEntity]{FieldName: fieldName}
	//             }
	//         }
	//     }
	// }
	return PropertySelector[TEntity]{FieldName: fieldName}
	// panic(fmt.Errorf("field %s does not exist on %s as json tag", fieldName, entityType.Name()))
}
