package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
  rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(min, max int64) int64{
  return min + rand.Int63n(max-min+1)
}

func RandomString(n int)string {
  var sb strings.Builder
  k := len(alphabet)
  for i := 0; i < n; i++ {
    c := alphabet[rand.Intn(k)]
    sb.WriteByte(c)
  }
  return sb.String()
} 
func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomString(6))
}
func RandomName() string {
	return cases.Title(language.English).String(RandomString(8)) // Capitalize first letter
}
func RandomRole() string {
	roles := []string{"user", "publisher"}
	return roles[rand.Intn(len(roles))]
}
func RandomMinimumSkill() string{
  minimumskill := []string{"beginner", "intermediate", "advanced"}
  return minimumskill[rand.Intn(len(minimumskill))]
}
func RandomPassword() string {
  return fmt.Sprintf("%s%d", RandomString(8), RandInt(1000, 9999))
}

func RandomSlug()string{
  return fmt.Sprintf("devworks-%s", RandomString(4))
}
