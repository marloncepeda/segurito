package quotation

type Quotation struct {
    Rut float64 `json:"rut" bson:"rut"`
    Date_birth string `json:"date_birth" bson:"date_birth"`
    Email string `json:"email" bson:"email"`
    Phone string `json:"phone" bson:"phone"`
}
