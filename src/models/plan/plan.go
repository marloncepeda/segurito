package plan

type Plan struct{
    Name string `json:"name" bson:"name"`
    Insurance string `json:"insurance" bson:"insurance"`
    Mount float32 `json:"mount" bson:"price"`
    Mount_uf float32 `json":mount_uf" bson:"mount_uf"`
    Quota string `json:"quota" bson:"quota"`
    Insured_capital float32 `json:"insured_capital" bson:"insured_capital"` 
}
