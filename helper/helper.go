package helper

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	log_function "github.com/aditya3232/url-shortener/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseDataTable struct {
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// pagination datatable
type Pagination struct {
	Page          int `form:"page" json:"page" binding:"required"`
	Limit         int `form:"limit" json:"limit" binding:"required"`
	Total         int `json:"total"`
	TotalFiltered int `json:"total_filtered"`
}

func NewPagination(page int, limit int) Pagination {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

type Sort struct {
	Sort  string `form:"sort" json:"sort"`
	Order string `form:"order" json:"order"`
}

func NewSort(sort string, order string) Sort {
	if sort == "" {
		sort = "id"
	}

	if order == "" {
		order = "asc"
	}

	return Sort{
		Sort:  sort,
		Order: order,
	}
}

func APIDataTableResponse(message string, code int, pagination Pagination, data interface{}) ResponseDataTable {
	meta := Meta{
		Message: message,
		Code:    code,
	}

	jsonResponse := ResponseDataTable{
		Meta:       meta,
		Pagination: pagination,
		Data:       data,
	}

	return jsonResponse
}

func APIResponse(message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

func RecoverPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
}

func FormatError(err error) []string {
	var errors []string
	errors = append(errors, fmt.Sprintf("%s", err))
	return errors
}

func FormatErrorWithCustomMessage(message string) []string {
	var errors []string
	errors = append(errors, message)
	return errors
}

func StringToDate(dateString string) time.Time {
	StringToDate, _ := time.Parse("2006-01-02", dateString)
	return StringToDate
}

func DateToString(t time.Time) string {
	date := t.Format("2006-01-02")
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "Z", "", -1)
	return date
}

func StringToDateTime(dateTimeString string) time.Time {
	StringToDateTime, _ := time.Parse("2006-01-02 15:04:05", dateTimeString)
	return StringToDateTime

}

func DateTimeToString(t time.Time) string {
	date := t.Format("2006-01-02 15:04:05")
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "Z", "", -1)
	return date
}

func DateTimeToStringWithMilliseconds(t time.Time) string {
	date := t.Format("2006-01-02 15:04:05.000")
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "Z", "", -1)
	return date
}

func DateTimeToStringWithStrip(t time.Time) string {
	date := t.Format("2006-01-02-15-04-05.000")
	date = strings.Replace(date, "T", " ", -1)
	date = strings.Replace(date, "Z", "", -1)
	date = strings.Replace(date, ".", "-", 1)
	return date
}

func ConstructOrderClause(query *gorm.DB, sort Sort) *gorm.DB {
	if sort.Sort != "" {
		query = query.Order(fmt.Sprintf("%s %s", sort.Sort, sort.Order))
	}
	return query
}

func ConstructPaginationClause(query *gorm.DB, pagination Pagination) *gorm.DB {
	query = query.Limit(pagination.Limit)
	query = query.Offset((pagination.Page - 1) * pagination.Limit)
	return query
}

func isValidDateRange(value string) bool {
	dateRegex := regexp.MustCompile(`^startDate:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}|endDate:\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`)
	return dateRegex.MatchString(value)
}

// ==================================================advanced filter==================================================
func StrToInt(str string) int {
	var i int
	fmt.Sscanf(str, "%d", &i)
	return i
}

func InArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(haystack)
		for i := 0; i < value.Len(); i++ {
			if reflect.DeepEqual(value.Index(i).Interface(), needle) {
				return true
			}
		}
	}
	return false
}

func GetJSONTags(t reflect.Type) []string {
	var tags []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
		if field.Type.Kind() == reflect.Struct {
			embeddedTags := GetJSONTags(field.Type)
			tags = append(tags, embeddedTags...)
		}
	}
	return tags
}

func QueryParamsToMap(c *gin.Context, s interface{}) map[string]string {
	params := make(map[string]string)
	queryParams := c.Request.URL.Query()

	jsonTags := GetJSONTags(reflect.TypeOf(s))

	exclude := map[string]bool{
		"comment": true,
		"limit":   true,
		"page":    true,
		"sort":    true,
		"order":   true,
		"dir":     true,
	}
	include := []string{
		"_all_",
	}

	if len(include) > 0 {
		jsonTags = append(jsonTags, include...)
	}

	_all_ := false
	if _, ok := queryParams["_all_"]; ok {
		_all_ = true
		for _, jsonTag := range jsonTags {
			queryParams[jsonTag] = []string{queryParams["_all_"][0]}
		}
	}

	for fieldName, fieldValue := range queryParams {
		if _, ok := exclude[fieldName]; ok {
			continue
		}

		if !InArray(fieldName, jsonTags) && !_all_ {
			continue
		}

		fieldName = getGormColumnNameFromJSON(s, fieldName)
		params[fieldName] = fieldValue[0]
	}

	return params
}

func getGormColumnNameFromJSON(s interface{}, jsonName string) string {
	// Get the type of the struct
	t := reflect.TypeOf(s)

	// Iterate over the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the value of the `json` tag
		jsonTag := field.Tag.Get("json")

		// Check if the `json` tag matches the specified name
		if jsonTag == jsonName {
			// Get the value of the `gorm` tag
			gormTag := field.Tag.Get("gorm")

			// Get the value of the `column` tag within the `gorm` tag
			switch {
			case strings.HasPrefix(gormTag, "column:"):
				columnTag := strings.Split(gormTag, ";")[0]
				return strings.TrimPrefix(columnTag, "column:")
			default:
				return jsonName
			}
		}
	}

	return jsonName
}

func ConstructWhereClause(query *gorm.DB, filter map[string]string) *gorm.DB {
	var allValue string

	// check if filter has _all_ key
	if v, ok := filter["_all_"]; ok {
		allValue = v
		delete(filter, "_all_")
	}

	for key, value := range filter {
		if value == "" {
			continue
		}

		if allValue != "" {
			value = allValue
		}

		switch {
		case isValidDateRange(value):
			dateRange := strings.Split(value, "|")
			startDate := strings.TrimPrefix(dateRange[0], "startDate:")
			endDate := strings.TrimPrefix(dateRange[1], "endDate:")
			query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", key), startDate, endDate)

		case strings.HasPrefix(value, "startDate:"):
			startDate := strings.TrimPrefix(value, "startDate:")
			query = query.Where(fmt.Sprintf("%s >= ?", key), startDate)

		case strings.HasPrefix(value, "endDate:"):
			endDate := strings.TrimPrefix(value, "endDate:")
			query = query.Where(fmt.Sprintf("%s <= ?", key), endDate)

		case strings.Contains(value, "_"):
			value = strings.Replace(value, "_", "%", -1)
			if allValue != "" {
				query = query.Or(fmt.Sprintf("%s LIKE ?", key), allValue)
			} else {
				query = query.Where(fmt.Sprintf("%s LIKE ?", key), value)
			}

		default:
			// add % to allValue, but dont loop %%value%% again
			if allValue != "" && !strings.Contains(allValue, "%") {
				allValue = fmt.Sprintf("%%%s%%", allValue)
			}

			if allValue != "" {
				query = query.Or(fmt.Sprintf("%s LIKE ?", key), allValue)
			} else {
				query = query.Where(fmt.Sprintf("%s LIKE ?", key), value)
			}
		}
	}

	// query = query.Where("deleted_at IS NULL")
	return query
}

// ==================================================advanced filter==================================================

// ==================================================file upload==================================================

// function to convert base64 to image
func RandomStringWithLength(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func Base64ToImage(base64String string) (string, error) {
	img, err := base64.StdEncoding.DecodeString(base64String[strings.IndexByte(base64String, ',')+1:])
	ext := base64String[strings.IndexByte(base64String, '/')+1 : strings.IndexByte(base64String, ';')]

	if err != nil {
		return "", err
	}

	path := RandomStringWithLength(10) + "." + ext
	err = os.WriteFile(path, img, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}

func GetMimeType(file string) string {
	buffer := make([]byte, 512)
	f, _ := os.Open(file)
	f.Read(buffer)
	f.Close()
	return http.DetectContentType(buffer)
}

func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log_function.Error(err)
	}
}

// ==================================================file upload==================================================

func CompressImageBytes(imageBytes []byte) ([]byte, error) {
	// decode image from bytes
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	// create buffer
	buf := new(bytes.Buffer)

	// encode image to buffer
	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 50})
	if err != nil {
		return nil, err
	}

	// return buffer as bytes
	return buf.Bytes(), nil
}

// is image helper
func IsImage(file *multipart.FileHeader) error {
	// open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// decode image from file
	_, _, err = image.Decode(src)
	if err != nil {
		return err
	}

	return nil
}

// convert image to jpg withoout reducce size
func ConvertImageToJpg(file *multipart.FileHeader) error {
	// open file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// decode image from file
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	// create buffer
	buf := new(bytes.Buffer)

	// encode image to buffer
	err = jpeg.Encode(buf, img, &jpeg.Options{})
	if err != nil {
		return err
	}

	// return buffer as bytes
	return nil
}

// convert file from multipart form data to base64, and compress to
func ConvertFileToBase64WithCompress(file *multipart.FileHeader) (string, error) {
	// open file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// decode image from file
	img, _, err := image.Decode(src)
	if err != nil {
		return "", err
	}

	// create buffer
	buf := new(bytes.Buffer)

	// encode image to buffer
	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 50})
	if err != nil {
		return "", err
	}

	// return buffer as bytes
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func ConvertFileToBase64(file *multipart.FileHeader) (string, error) {
	// open file
	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}

	// read file
	readFile, err := io.ReadAll(openedFile)
	if err != nil {
		return "", err
	}

	// convert file to base64
	base64String := base64.StdEncoding.EncodeToString(readFile)

	return base64String, nil
}

// encrypt url
func Encrypt(url string) string {
	encryptedUrl := base64.StdEncoding.EncodeToString([]byte(url))
	return encryptedUrl
}
