package schema
 
import (
   "github.com/facebookincubator/ent"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent/schema/edge"
)
 
// User holds the schema definition for the User entity.
type User struct {
   ent.Schema
}
 
// Fields of the User.
func (User) Fields() []ent.Field {
   return []ent.Field{
	   field.String("name").NotEmpty(),
	   field.String("email").NotEmpty(),
	   field.String("password").NotEmpty(),
   }
}
 
// Edges of the User.
func (User) Edges() []ent.Edge {
   return []ent.Edge {
      edge.From("gender", Gender.Type).Ref("users").Unique(),
		edge.From("position", Position.Type).Ref("users").Unique(),
		edge.From("title", Title.Type).Ref("users").Unique(),
   }
}